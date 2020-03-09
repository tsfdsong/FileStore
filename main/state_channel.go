package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//RPCArgs ...
type RPCArgs struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     int           `json:"id"`
}

//ChannelInfo ...
type ChannelInfo struct {
	Address   string  `json:"address"`
	Nonce     float64 `json:"nonce"`
	Oracle    string  `json:"oracle"`
	State     string  `json:"state"`
	Expire    int64   `json:"expiry"`
	Owner     string  `json:"owner"`
	IndexHash string  `json:"indexhash"`
}

//GetRPCResponse ...
func GetRPCResponse(in *RPCArgs) ([]byte, error) {
	seed := rand.NewSource(time.Now().Unix())
	in.ID = rand.New(seed).Intn(1000)

	bytes, err := json.Marshal(in)
	if err != nil {
		Errorf("Marshal para: %v \n", err)
		return nil, err
	}
	s := string(bytes)

	payload := strings.NewReader(s)

	req, err := http.NewRequest("POST", MyEnv["RPCGATEWAY"]+"/rpc", payload)
	if err != nil {
		Errorf("Post: %v \n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Errorf("DefaultClient Do: %v \n", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Errorf("ReadAll body: %v \n", err)
		return nil, err
	}

	return body, nil
}

//GetRPCResult ...
func GetRPCResult(url string, in *RPCArgs) ([]byte, error) {
	seed := rand.NewSource(time.Now().Unix())
	in.ID = rand.New(seed).Intn(1000)

	bytes, err := json.Marshal(in)
	if err != nil {
		Errorf("Marshal para: %v \n", err)
		return nil, err
	}
	s := string(bytes)

	payload := strings.NewReader(s)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		Errorf("Post: %v \n", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Errorf("DefaultClient Do: %v \n", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		Errorf("RPC response, %v", res.Status)
		return nil, fmt.Errorf("RPC response, %v", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Errorf("ReadAll body: %v \n", err)
		return nil, err
	}

	return body, nil
}

//SetUp ...
func SetUp(contractAddr, account string) (string, error) {
	var args RPCArgs
	args.Method = "setup"
	args.Params = append(args.Params, contractAddr)
	args.Params = append(args.Params, account)

	args.ID = 100

	data, err := GetRPCResponse(&args)
	if err != nil {
		return "", err
	}

	reply := &struct {
		Result struct {
			Oracle string `json:"oracle"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(data, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != "null" {
		return "", fmt.Errorf("%v", reply.Error)
	}

	oracle := reply.Result.Oracle

	Infof("SetUp:  %v", oracle)

	return oracle, nil
}

//GetChannelState ...
func GetChannelState(contractAddr string) (string, error) {
	var args RPCArgs
	args.Method = "getchannelstate"
	args.Params = append(args.Params, contractAddr)
	args.ID = 100

	data, err := GetRPCResponse(&args)
	if err != nil {
		return "", err
	}

	reply := &struct {
		Result struct {
			State string `json:"state"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(data, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != "null" {
		return "", fmt.Errorf("%v", reply.Error)
	}

	state := reply.Result.State

	Infof("getchannelstate:  %v", state)

	return state, nil
}

//GetChannelInfo ...
func GetChannelInfo(contractAddr string) (*ChannelInfo, error) {
	var args RPCArgs
	args.Method = "getchannelinfo"
	args.Params = append(args.Params, contractAddr)
	args.ID = 100

	data, err := GetRPCResponse(&args)
	if err != nil {
		return nil, err
	}

	reply := &struct {
		Result ChannelInfo `json:"result"`
		Error  string      `json:"error"`
		ID     int         `json:"id"`
	}{}

	err = json.Unmarshal(data, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != "null" {
		errStr := reply.Error
		if strings.Contains(errStr, "channel not exist") {
			return nil, fmt.Errorf("channel closed")
		} else if strings.Contains(errStr, "channel already setup") {
			return nil, fmt.Errorf("channel closed")
		} else if strings.Contains(errStr, "channel init") {
			return nil, fmt.Errorf("channel init")
		} else if strings.Contains(errStr, "channel closed") {
			return nil, fmt.Errorf("channel closed")
		}

		return nil, fmt.Errorf("%v", reply.Error)
	}

	info := reply.Result

	Infof("GetChannelInfo:  %v", info)

	return &info, nil
}

//GetNonce ...
func GetNonce(contractAddr string) (float64, error) {
	var args RPCArgs
	args.Method = "getnonce"
	args.Params = append(args.Params, contractAddr)
	args.ID = 100

	data, err := GetRPCResponse(&args)
	if err != nil {
		return 0, err
	}

	reply := &struct {
		Result struct {
			Nonce float64 `json:"nonce"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(data, reply)
	if err != nil {
		return 0, err
	}

	if reply.Error != "null" {
		return 0, fmt.Errorf("%v", reply.Error)
	}

	nonce := reply.Result.Nonce

	Infof("Getnonce:  %v", nonce)

	return nonce, nil
}

//GetIndexHash ...
func GetIndexHash(contractAddr string) (string, error) {
	var args RPCArgs
	args.Method = "gethash"
	args.Params = append(args.Params, contractAddr)
	args.ID = 100

	data, err := GetRPCResponse(&args)
	if err != nil {
		return "", err
	}

	reply := &struct {
		Result struct {
			Hash string `json:"hash"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(data, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != "null" {
		return "", fmt.Errorf("%v", reply.Error)
	}

	hash := reply.Result.Hash

	Infof("Gethash:  %v", hash)

	return hash, nil
}

//SetIndexHash ...
func SetIndexHash(contractAddr, indexhash, account string) error {
	nonce, err := GetNonce(contractAddr)
	if err != nil {
		return err
	}

	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(nonce))

	data := []byte(indexhash)
	hasher := sha256.New()
	hasher.Write(data)
	hasher.Write(buf)
	hash := hasher.Sum(nil)

	signStr, err := Sign(hash)
	if err != nil {
		return err
	}

	var args RPCArgs
	args.Method = "updatehash"
	args.Params = append(args.Params, contractAddr)
	args.Params = append(args.Params, nonce)
	args.Params = append(args.Params, indexhash)
	args.Params = append(args.Params, signStr)
	args.ID = 100

	bytes, err := GetRPCResponse(&args)
	if err != nil {
		return err
	}

	reply := &struct {
		Result struct {
			Hash string `json:"hash"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return err
	}

	if reply.Error != "null" {
		return fmt.Errorf("%v", reply.Error)
	}

	Infof("Updatehash:  %v", reply.Result.Hash)

	return nil
}

//CreateChannel ...
func CreateChannel(contractAddr, abi string) (string, error) {
	var args RPCArgs
	args.Method = "CreateChannel"
	args.Params = append(args.Params, contractAddr)
	args.Params = append(args.Params, abi)

	Infof("RPC request parameter: CreateChannel %v", contractAddr)

	bytes, err := GetRPCResult(MyEnv["RPCGATEWAY"]+"/rpc", &args)
	if err != nil {
		return "", err
	}

	reply := &struct {
		Result struct {
			ContainerID string `json:"container"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != "null" {
		return "", fmt.Errorf("%v", reply.Error)
	}

	if reply.Result.ContainerID == "" {
		return "", fmt.Errorf("create container empty")
	}

	Infof("Create channel:  %v", reply.Result.ContainerID)

	return reply.Result.ContainerID, nil
}

//SetHash ...
func SetHash(containerID, indexhash string) error {
	var args RPCArgs
	args.Method = "UpdateHash"
	args.Params = append(args.Params, indexhash)

	Infof("RPC UpdateHash:  container=%v indexhash=%v", containerID, indexhash)

	bytes, err := GetRPCResult(MyEnv["RPCGATEWAY"]+"/channel?id="+containerID, &args)
	if err != nil {
		return err
	}

	reply := &struct {
		Result struct {
			Hash string `json:"hash"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return err
	}

	if reply.Error != "null" {
		return fmt.Errorf("%v", reply.Error)
	}

	Infof("RPC Updatehash result:  %v", reply.Result.Hash)

	return nil
}

//GetHash ...
func GetHash(containerID string) (string, error) {
	var args RPCArgs
	args.Method = "GetIndexHash"

	Infof("RPC GetIndexHash, container=%v", containerID)

	bytes, err := GetRPCResult(MyEnv["RPCGATEWAY"]+"/channel?id="+containerID, &args)
	if err != nil {
		return "", err
	}

	reply := &struct {
		Result struct {
			Hash string `json:"hash"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return "", err
	}

	if reply.Error != "null" {
		return "", fmt.Errorf("%v", reply.Error)
	}

	hash := reply.Result.Hash

	Infof("RPC GetIndexHash result:  %v", hash)

	return hash, nil
}

//CloseChannel ...
func CloseChannel(id string) error {
	var args RPCArgs
	args.Method = "DestoryChannel"
	args.Params = append(args.Params, id)

	Infof("RPC DestoryChannel, container=%v", id)

	bytes, err := GetRPCResult(MyEnv["RPCGATEWAY"]+"/rpc", &args)
	if err != nil {
		return err
	}

	reply := &struct {
		Result struct {
			Result string `json:"result"`
		} `json:"result"`
		Error string `json:"error"`
		ID    int    `json:"id"`
	}{}

	err = json.Unmarshal(bytes, reply)
	if err != nil {
		return err
	}

	if reply.Error != "null" {
		return fmt.Errorf("%v", reply.Error)
	}

	if reply.Result.Result != "success" {
		return fmt.Errorf("DestoryChannel, %v", reply.Result.Result)
	}

	return nil
}
