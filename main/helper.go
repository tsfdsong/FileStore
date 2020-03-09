package main

import (
	"FileStore/accountmanager"
	"FileStore/storecontract"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"
)

//MyEnv ...
var MyEnv map[string]string

const envLoc = "config/.env"

const (
	//StatusChannelCode ...
	StatusChannelCode int = 900
	//StatusChannelInit ...
	StatusChannelInit int = 901
	//StatusChannelClosed ...
	StatusChannelClosed int = 902
)

//NodeKeyPrivate ...
var NodeKeyPrivate *ecdsa.PrivateKey

//NodeKeyStore ...
var NodeKeyStore string

//TransactionTrace ...
type TransactionTrace struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Gas         int    `json:"gas"`
		Failed      bool   `json:"failed"`
		ReturnValue string `json:"returnValue"`
		StructLogs  []struct {
			Pc      int    `json:"pc"`
			Op      string `json:"op"`
			Gas     int    `json:"gas"`
			GasCost int    `json:"gasCost"`
			Depth   int    `json:"depth"`
		} `json:"structLogs"`
	} `json:"result"`
}

//LoadEnv ...
func LoadEnv() {
	var err error
	if MyEnv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

// UpdateEnvFile updates our env file with a key-value pair
func UpdateEnvFile(k string, val string) {
	MyEnv[k] = val
	err := godotenv.Write(MyEnv, envLoc)
	if err != nil {
		log.Printf("failed to update %s: %v\n", envLoc, err)
	}
}

//HTTPError ...
type HTTPError struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  string `json:"result"`
}

//WriteToResponse ...
func (he *HTTPError) WriteToResponse(w http.ResponseWriter) {
	bytes, err := json.Marshal(he)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

//HandleError ...
func HandleError(code int, err error, data string, w http.ResponseWriter) {
	var result *HTTPError
	if err == nil {
		result = &HTTPError{
			RetCode: code,
			RetMsg:  "",
			Result:  data,
		}
		Infof("\n http response: %v \n\n", result)
	} else {
		result = &HTTPError{
			RetCode: code,
			RetMsg:  err.Error(),
			Result:  data,
		}
		Errorf("\n http response: %v \n\n", result)
	}

	result.WriteToResponse(w)
}

//GetFileHash ...MD5
func GetFileHash(bytes []byte) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write(bytes)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString(hasher.Sum(nil)[:16])
	return result, nil
}

//PutDHT ...
func PutDHT(bytes []byte, hash string) error {
	file, err := os.Create("./DHT/" + hash + ".json")
	if err != nil {
		Errorf("WriteToDHT create file :%v", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		Errorf("WriteToDHT create file :%v", err)
		return err
	}

	return nil
}

//PutFile ...
func PutFile(data []byte, fileName string) (string, error) {
	sh := shell.NewShell(MyEnv["IPFSGATEWAY"])
	hash, err := sh.Add(strings.NewReader(string(data)))
	if err != nil {
		Errorf("PutFile :%v", err)
		return "", err
	}

	return hash, nil
}

//PutFileWithBody ...
func PutFileWithBody(body io.ReadCloser, filename string) ([]byte, error) {
	reader, writer := io.Pipe()
	multipartWriter := multipart.NewWriter(writer)

	var response *http.Response
	done := make(chan error)
	var err error

	go func() {
		req, err := http.NewRequest("POST", MyEnv["IPFSGATEWAY"]+"file", reader)
		if err != nil {
			done <- err
			return
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		req.Header.Add("X-Dynamic-Token", MyEnv["NODETOKEN"])
		req.Close = true

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			done <- err
			return
		}

		response.Close = true

		if response.StatusCode != 201 {
			done <- fmt.Errorf("%v", response.Status)
			return
		}

		done <- nil
	}()

	fileName := fmt.Sprintf("%v", time.Now())
	fw, err := multipartWriter.CreateFormFile("file", fileName)
	if err != nil {
		Errorf("PutFileWithBody CreateFormFile: %v\n", err)
		return nil, err
	}

	_, err = io.Copy(fw, body)
	if err != nil {
		Errorf("PutFileWithBody Copy body: %v\n", err)
		return nil, err
	}

	err = multipartWriter.Close()
	if err != nil {
		Errorf("PutFileWithBody multipart close: %v\n", err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		Errorf("PutFileWithBody writer close: %v\n", err)
		return nil, err
	}

	err = <-done
	if err != nil {
		Errorf("PutFileWithBody post file: %v\n", err)
		return nil, err
	}

	defer response.Body.Close()

	respbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Errorf("PutFileWithBody readall: %v\n", err)
		return nil, err
	}

	return respbody, nil
}

//GetFileWithBody ...
func GetFileWithBody(hash string) ([]byte, error) {
	sh := shell.NewShell(MyEnv["IPFSGATEWAY"])
	filepath := "./storge/" + hash
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Errorf("GetFileWithBody open file: %v %v\n", hash, err)
		return nil, err
	}
	defer file.Close()

	err = sh.Get(hash, filepath)
	if err != nil {
		Errorf("GetFileWithBody ipfs get: %v %v\n", hash, err)
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		Errorf("GetFileWithBody read file :%v\n", err)
		return nil, err
	}

	return data, nil
}

//GetDHT ...
func GetDHT(hash string) ([]byte, error) {
	data, err := ioutil.ReadFile("./DHT/" + hash + ".json")
	if err != nil {
		Errorf("GetDHT read file :%v\n", err)
		return nil, err
	}

	return data, nil
}

//Sign ...
func Sign(data []byte) (string, error) {
	if NodeKeyPrivate == nil {
		return "", fmt.Errorf("Node private key is empty")
	}

	sig, err := crypto.Sign(data, NodeKeyPrivate)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString(sig)
	return result, nil
}

//GetPrivateKey ...
func GetPrivateKey(path, password string) (string, error) {
	kjson, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	key, err := keystore.DecryptKey(kjson, password)
	if err != nil {
		return "", err
	}

	keystore := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	return keystore, nil
}

//ValidateAccountPrivateKey ...
func ValidateAccountPrivateKey(account, keyStr string) bool {
	privateKey, err := crypto.HexToECDSA(keyStr)
	if err != nil {
		return false
	}

	publicKey := privateKey.PublicKey

	inAddress := crypto.PubkeyToAddress(publicKey)
	accountAddr := common.HexToAddress(account)

	if !reflect.DeepEqual(inAddress, accountAddr) {
		return false
	}

	return true
}

//InitStoreContract ...
func InitStoreContract(user, keystore string) (string, error) {
	result, err := storecontract.CreateStoreContract(MyEnv["GATEWAY"], keystore)
	if err != nil {
		Errorf("InitStoreContract create store contract: %v\n", err)
		return "", err
	}

	if (result == common.Address{}) {
		Errorf("InitStoreContract create store contract empty")
		return "", fmt.Errorf("InitStoreContract create store contract empty")
	}

	//set code
	storeSession, err := storecontract.GetStoreContract(MyEnv["GATEWAY"], keystore, result)
	if err != nil {
		Errorf("InitStoreContract get store contract: %v\n", err)
		return "", err
	}

	txCode := MyEnv["TXCODE"]
	var txBytes [32]byte
	hash := common.HexToHash(txCode)
	copy(txBytes[:], hash.Bytes())
	codeTrx, err := storeSession.SetCode(txBytes)
	if err != nil {
		Errorf("InitStoreContract set tx code: %v", err)
		return "", err
	}

	err = func(trx *types.Transaction) error {
		//timer
		var ch = make(chan int)
		go func() {
			chtimer := time.NewTimer(10 * time.Second)
			for {
				select {
				case <-chtimer.C:
					ch <- 1
					chtimer.Stop()
				}
			}
		}()

		client, err := ethclient.Dial(MyEnv["GATEWAY"])
		if err != nil {
			return err
		}

		timer := time.NewTimer(time.Second)
		defer timer.Stop()
		for {
			timer.Reset(time.Second)
			select {
			case <-timer.C:
				{
					receipt, err := client.TransactionReceipt(context.Background(), trx.Hash())
					if err == ethereum.NotFound {
						break
					}
					if err != nil {
						return err
					}

					Infof("InitStoreContract set tx code trx receipt: %v", receipt)
					return nil
				}
			case <-ch:
				return fmt.Errorf("InitStoreContract set tx code timeout, user=%v ", user)
			}
		}

	}(codeTrx)

	if err != nil {
		Errorf("set tx code %v", err)
		return "", err
	}

	//set account contract
	accountContractAddr := common.HexToAddress(MyEnv["ACCOUNTCONTRACTADDR"])
	session, err := accountmanager.GetAccountContract(MyEnv["GATEWAY"], keystore, accountContractAddr)
	if err != nil {
		Errorf("InitStoreContract get session: %v\n", err)
		return "", fmt.Errorf("InitStoreContract get session: %v", err)
	}

	trx, err := session.Set(result)
	if err != nil {
		Errorf("InitStoreContract set store contract address: %v\n", err)
		return "", err
	}

	err = func(trx *types.Transaction) error {
		//timer
		var ch = make(chan int)
		go func() {
			chtimer := time.NewTimer(15 * time.Second)
			for {
				select {
				case <-chtimer.C:
					ch <- 1
					chtimer.Stop()
				}
			}
		}()

		client, err := ethclient.Dial(MyEnv["GATEWAY"])
		if err != nil {
			return err
		}

		timer := time.NewTimer(time.Second)
		defer timer.Stop()
		for {
			timer.Reset(time.Second)
			select {
			case <-timer.C:
				{
					receipt, err := client.TransactionReceipt(context.Background(), trx.Hash())
					if err == ethereum.NotFound {
						break
					}
					if err != nil {
						return err
					}

					Infof("InitStoreContract save contract address trx receipt: %v", receipt)
					return nil
				}
			case <-ch:
				return fmt.Errorf("InitStoreContract save contract address timeout, user=%v", user)
			}
		}

	}(trx)

	if err != nil {
		Errorf("save contract %v", err)
		return "", err
	}

	lowstr := result.String()

	//set addr cache
	GetInstance().Put(user, lowstr)

	return lowstr, nil
}

//InitChannelState ...
func InitChannelState(user, keystore, contractAddr string) error {
	tempID, err := GetObjectInstance().GetContainer(user)
	if err == nil {
		Debugf("Get cache container: %v", tempID)
		return nil
	}

	//create state channel
	containerID, err := CreateChannel(contractAddr, storecontract.FileStoreABI)
	if err != nil {
		return err
	}

	GetObjectInstance().PutContainer(user, containerID)

	Infof("InitChannelState create channel: smart contract=%v  container=%v\n", contractAddr, containerID)

	return err
}

//CheckUserContract ...
func CheckUserContract(user, privateKey string) (string, error) {
	//get addr from cache
	addr, err := GetInstance().Get(user)
	if err == nil {
		return addr, nil
	}

	accountContractAddr := common.HexToAddress(MyEnv["ACCOUNTCONTRACTADDR"])
	session, err := accountmanager.GetAccountContract(MyEnv["GATEWAY"], privateKey, accountContractAddr)
	if err != nil {
		Errorf("CheckUserContract get account contract session: %v\n", err)
		return "", fmt.Errorf("CheckUserContract get session: %v", err)
	}

	useraccount := common.HexToAddress(user)
	conaddr, err := session.Get(useraccount)
	if err != nil {
		Errorf("CheckUserContract get user store contract address: %v\n", err)
		return "", fmt.Errorf("CheckUserContract get user=%v store contract address: %v", user, err)
	}

	//not found
	if reflect.DeepEqual(conaddr, common.HexToAddress("")) {
		return "", nil
	}

	//set addr cache
	GetInstance().Put(user, conaddr.String())

	return conaddr.String(), nil
}

//GetUserContract ...
func GetUserContract(user, privateKey string) (string, error) {
	//get addr from cache
	addr, err := GetInstance().Get(user)
	if err == nil {
		return addr, nil
	}

	accountContractAddr := common.HexToAddress(MyEnv["ACCOUNTCONTRACTADDR"])
	session, err := accountmanager.GetAccountContract(MyEnv["GATEWAY"], privateKey, accountContractAddr)
	if err != nil {
		Errorf("GetUserContract get account contract session: %v\n", err)
		return "", fmt.Errorf("GetUserContract get session: %v", err)
	}

	useraccount := common.HexToAddress(user)
	conaddr, err := session.Get(useraccount)
	if err != nil {
		Errorf("GetUserContract get user store contract address: %v\n", err)
		return "", fmt.Errorf("GetUserContract get user=%v store contract address: %v", user, err)
	}

	//not found
	if reflect.DeepEqual(conaddr, common.HexToAddress("")) {
		Errorf("GetUserContract get user store contract address empty")
		return "", fmt.Errorf("GetUserObjects get user=%v store contract address empty", user)
	}

	//set addr cache
	GetInstance().Put(user, conaddr.String())

	return conaddr.String(), nil
}

//GetUserObjects ...
func GetUserObjects(user string) (*Objects, error) {
	containerID, err := GetObjectInstance().GetContainer(user)
	if err != nil {
		Errorf("GetUserObjects get user container: %v\n", err)
		return nil, fmt.Errorf("channel closed")
	}

	path, err := GetHash(containerID)
	if err != nil {
		Errorf("GetUserObjects user=%v get index hash failed: %v", user, err)
		GetObjectInstance().PutContainer(user, "")
		return nil, fmt.Errorf("channel closed, %v", err)
	}

	Infof("GetUserObjects userid=%v indexhash=%v", user, path)

	res, err := GetObjectInstance().Get(user, path)
	if err == nil {
		return res, err
	}

	var obj *Objects
	if path == "" {
		obj = CreateObjects(user)
	} else {
		bytes, err := GetFileWithBody(path)
		if err != nil {
			Errorf("GetUserObjects get file body: %v\n", err)
			return nil, err
		}

		obj, err = UnmarshalToObjects(user, bytes)
		if err != nil {
			Errorf("GetUserObjects unmarshal to objects: %v\n", err)
			return nil, err
		}
	}

	//cache user objects
	GetObjectInstance().Put(user, path, obj)

	Debugf("GetUserObjects  userid=%v indexcontent=%v", user, obj.String())

	return obj, nil
}

//GetFileIndex ...
func GetFileIndex(user string) (string, error) {
	//get addr from cache
	var storeAddr string
	tempAddr, err := GetInstance().Get(user)
	if err == nil {
		storeAddr = tempAddr
	} else {
		accountContractAddr := common.HexToAddress(MyEnv["ACCOUNTCONTRACTADDR"])
		session, err := accountmanager.GetAccountContract(MyEnv["GATEWAY"], NodeKeyStore, accountContractAddr)
		if err != nil {
			Errorf("GetFileIndex get account contract session: %v\n", err)
			return "", fmt.Errorf("GetFileIndex get session: %v", err)
		}

		useraccount := common.HexToAddress(user)
		conaddr, err := session.Get(useraccount)
		if err != nil {
			Errorf("GetFileIndex get user store contract address: %v\n", err)
			return "", fmt.Errorf("GetFileIndex get user=%v store contract address: %v", user, err)
		}

		//not found
		if reflect.DeepEqual(conaddr, common.HexToAddress("")) {
			return "", fmt.Errorf("GetFileIndex get user=%v store contract empty", user)
		}

		GetInstance().Put(user, conaddr.String())
	}

	store := common.HexToAddress(storeAddr)

	contractSession, err := storecontract.GetStoreContract(MyEnv["GATEWAY"], NodeKeyStore, store)
	if err != nil {
		Errorf("GetFileIndex get store contract: %v\n", err)
		return "", err
	}

	hash, err := contractSession.IndexHash()
	if err != nil {
		Errorf("GetFileIndex get index hash: %v\n", err)
		return "", err
	}

	return hash, nil
}

//GetUserIndex ...
func GetUserIndex(user string) (string, *Objects, error) {
	containerID, err := GetObjectInstance().GetContainer(user)
	if err != nil {
		Errorf("GetUserIndex get user container: %v\n", err)
		return "", nil, fmt.Errorf("channel closed")
	}

	hash, err := GetHash(containerID)
	if err != nil {
		Errorf("GetUserIndex user=%v get index hash failed: %v", user, err)
		GetObjectInstance().PutContainer(user, "")
		return "", nil, fmt.Errorf("channel closed, %v", err)
	}

	Infof("GetUserIndex userid=%v indexhash=%v", user, hash)

	res, err := GetObjectInstance().Get(user, hash)
	if err == nil {
		return "", res, err
	}

	var obj *Objects
	if hash == "" {
		obj = CreateObjects(user)
	} else {
		bytes, err := GetFileWithBody(hash)
		if err != nil {
			Errorf("GetUserIndex get file body: %v\n", err)
			return "", nil, err
		}

		obj, err = UnmarshalToObjects(user, bytes)
		if err != nil {
			Errorf("GetUserIndex unmarshal to objects: %v\n", err)
			return "", nil, err
		}
	}

	//cache user objects
	GetObjectInstance().Put(user, hash, obj)

	Infof("GetUserIndex userid=%v indexhash=%v indexcontent=%v", user, hash, obj.String())

	return hash, obj, nil
}

//UpdateDHTAndContract ...
func UpdateDHTAndContract(user string, obj *Objects) error {
	bytes, err := obj.MarshalToJSON()
	if err != nil {
		Errorf("UpdateDHTAndContract marshal objects: %v\n", err)
		return err
	}

	fileName := fmt.Sprintf("%v_%v.json", user, time.Now().Unix())

	hashStr, err := PutFile(bytes, fileName)
	if err != nil {
		return fmt.Errorf("UpdateDHTAndContract upload file failed, %v", err)
	}

	Infof("UpdateDHTAndContract userid=%v indexhash=%v \n", user, hashStr)
	Debugf("\n UpdateDHTAndContract indexcontent=%v \n\n", string(bytes))

	//set hash
	containerID, err := GetObjectInstance().GetContainer(user)
	if err != nil {
		Debugf("UpdateDHTAndContract get cache container failed, %v", err)
		return fmt.Errorf("channel closed")
	}

	err = SetHash(containerID, hashStr)
	if err != nil {
		Debugf("UpdateDHTAndContract set index hash failed, %v", err)
		GetObjectInstance().PutContainer(user, "")
		return fmt.Errorf("channel closed, %v", err)
	}

	//cache user objects
	GetObjectInstance().Put(user, hashStr, obj)

	return nil
}

//CheckStateChannel ...
func CheckStateChannel(err error, w http.ResponseWriter) {
	if err == nil {
		HandleError(http.StatusOK, nil, fmt.Sprintf("check state channel success"), w)
	}

	errStr := err.Error()
	if strings.Contains(errStr, "channel not opened") {
		HandleError(StatusChannelClosed, fmt.Errorf("channel not opened"), "", w)
	} else if strings.Contains(errStr, "channel not exist") {
		HandleError(StatusChannelClosed, fmt.Errorf("channel not exist"), "", w)
	} else if strings.Contains(errStr, "channel init") {
		HandleError(StatusChannelInit, fmt.Errorf("channel init"), "", w)
	} else if strings.Contains(errStr, "channel closed") {
		HandleError(StatusChannelClosed, fmt.Errorf("channel closed"), "", w)
	} else {
		HandleError(http.StatusInternalServerError, fmt.Errorf("%v", err), "", w)
	}
}

//GetTransactStatus ...
func GetTransactStatus(url string, trx *types.Transaction) (bool, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return false, err
	}

	receipt, err := client.TransactionReceipt(context.Background(), trx.Hash())
	if err != nil {
		return false, err
	}

	// Errorf("GetTransactStatus receipt: %v\n", receipt)

	if receipt.Status == 1 {
		return true, nil
	}

	return false, fmt.Errorf("trx receipt failed: %v", trx.Hash().String())
}

//TraceTransactionStatus ...
func TraceTransactionStatus(hash string) (string, error) {
	parajson := `{"id": 1, "method": "debug_traceTransaction", "params": ["` + hash + `", {"disableStack": true, "disableMemory": true, "disableStorage": true}]}`

	Infof("TraceTransaction: %v", hash)

	jsonByte := []byte(parajson)
	req, _ := http.NewRequest("POST", MyEnv["GATEWAY"], bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Errorf("TraceTransaction Do Request error: %v\n", err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	var trace TransactionTrace
	err = json.Unmarshal(body, &trace)
	if err != nil {
		return "", err
	}

	Infof("TraceTransaction Request info: %v", trace)

	if !trace.Result.Failed {
		return "", nil
	}

	return trace.Result.ReturnValue, fmt.Errorf("TraceTransaction: %v", trace.Result.ReturnValue)
}

//GenerateToken ...
func GenerateToken(user string) (string, error) {
	if user == "" {
		return "", fmt.Errorf("generate token user is empty")
	}

	exp := time.Now().Add(24 * time.Hour).Unix()

	claims := &jwt.StandardClaims{
		Issuer:    user,
		ExpiresAt: exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(MyEnv["TOKENKEY"]))
	if err != nil {
		return "", fmt.Errorf("generate token : %v", err)
	}

	return tokenString, nil
}

//ValidateToken ...
func ValidateToken(tokenSrt string, SecretKey []byte) bool {
	token, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return false
	}

	Infof("valid token: %v", token)

	return true
}
