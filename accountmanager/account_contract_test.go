package accountmanager

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_CreateAccountContract(t *testing.T) {
	kjson, err := ioutil.ReadFile("/Users/tangsong/Library/Ethereum/keystore/UTC--2019-03-21T02-58-55.366759000Z--cb8f035b48e81c79db5fadd32fae108b26c06ea6")
	if err != nil {
		t.Errorf("CreateAccountContract ReadFile: %v \n", err)
		return
	}

	key, err := keystore.DecryptKey(kjson, "123")
	if err != nil {
		t.Errorf("CreateAccountContract DecryptKey: %v \n", err)
		return
	}

	keystore := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf("CreateAccountContract private key: %v\n", keystore)

	contract, err := CreateAccountContract("http://47.100.33.107:8545", keystore)
	if err != nil {
		t.Errorf("CreateAccountContract: %v \n", err)
		return
	}

	fmt.Printf("CreateAccountContract contract addr: %v\n", contract.String())
	t.Logf("CreateAccountContract result: %v\n", contract.String())
}

func Test_SetStoreContract(t *testing.T) {
	//account keystore
	kjson, err := ioutil.ReadFile("/Users/tangsong/Library/Ethereum/keystore/UTC--2019-03-21T02-58-55.366759000Z--cb8f035b48e81c79db5fadd32fae108b26c06ea6")
	if err != nil {
		t.Errorf("SetStoreContract ReadFile: %v \n", err)
		return
	}

	key, err := keystore.DecryptKey(kjson, "123")
	if err != nil {
		t.Errorf("SetStoreContract DecryptKey: %v \n", err)
		return
	}

	keystore := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf("SetStoreContract private key: %v\n", keystore)

	//account manage contract
	manageContract := common.HexToAddress("0x17d685188E22C612E187463f434FB0482Cd496Be")
	session, err := GetAccountContract("http://192.168.1.4:8545", keystore, manageContract)
	if err != nil {
		t.Errorf("SetStoreContract: %v \n", err)
		return
	}

	storeAddr := common.HexToAddress("0xca35b7d915458ef540ade6068dfe2f44e8fa757b")
	trx, err := session.Set(storeAddr)
	if err != nil {
		t.Errorf("SetStoreContract set: %v \n", err)
		return
	}

	fmt.Printf("SetStoreContract set trx hash: %v\n", trx.Hash().String())
	t.Logf("SetStoreContract set hash result: %v\n", trx.Hash().String())
}

func Test_GetStoreContract(t *testing.T) {
	//account keystore
	kjson, err := ioutil.ReadFile("/Users/tangsong/Library/Ethereum/keystore/UTC--2019-03-21T02-58-55.366759000Z--cb8f035b48e81c79db5fadd32fae108b26c06ea6")
	if err != nil {
		t.Errorf("GetStoreContract ReadFile: %v \n", err)
		return
	}

	key, err := keystore.DecryptKey(kjson, "123")
	if err != nil {
		t.Errorf("GetStoreContract DecryptKey: %v \n", err)
		return
	}

	keystore := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf("GetStoreContract private key: %v\n", keystore)

	//account manage contract
	manageContract := common.HexToAddress("0x17d685188E22C612E187463f434FB0482Cd496Be")

	session, err := GetAccountContract("http://192.168.1.4:8545", keystore, manageContract)
	if err != nil {
		t.Errorf("GetStoreContract get account contract: %v \n", err)
		return
	}

	//account address
	account := common.HexToAddress("0xcb8f035b48e81c79db5fadd32fae108b26c06ea6")

	conaddr, err := session.Get(account)
	if err != nil {
		t.Errorf("GetStoreContract get: %v \n", err)
		return
	}

	fmt.Printf("GetStoreContract store contract address: %v\n", conaddr.String())
}

func Test_ClearStoreContract(t *testing.T) {
	//account keystore
	kjson, err := ioutil.ReadFile("/Users/tangsong/Library/Ethereum/keystore/UTC--2019-03-21T02-58-55.366759000Z--cb8f035b48e81c79db5fadd32fae108b26c06ea6")
	if err != nil {
		t.Errorf("ClearStoreContract ReadFile: %v \n", err)
		return
	}

	key, err := keystore.DecryptKey(kjson, "123")
	if err != nil {
		t.Errorf("ClearStoreContract DecryptKey: %v \n", err)
		return
	}

	keystore := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Printf("ClearStoreContract private key: %v\n", keystore)

	//account manage contract
	manageContract := common.HexToAddress("0x9c258c2FE6909b98Bc549C431B14E4ca03c13066")

	session, err := GetAccountContract("http://192.168.50.4:8545", keystore, manageContract)
	if err != nil {
		t.Errorf("ClearStoreContract get account contract: %v \n", err)
		return
	}

	//account address
	account := common.HexToAddress("0x54379E45339356f6F1817ed39a525C0874bE6109")

	trx, err := session.Clear(account)
	if err != nil {
		t.Errorf("ClearStoreContract clear: %v \n", err)
		return
	}

	fmt.Printf("ClearStoreContract clear trx hash: %v\n", trx.Hash().String())
	t.Logf("ClearStoreContract clear hash result: %v\n", trx.Hash().String())
}

func Test_SubscribeEvents(t *testing.T) {
	client, err := ethclient.Dial("http://192.168.1.4:8545")
	if err != nil {
		t.Errorf("SubscribeEvents dial: %v \n", err)
		return
	}

	contractAddress := common.HexToAddress("0x17d685188E22C612E187463f434FB0482Cd496Be")
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(0),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		t.Errorf("SubscribeEvents FilterLogs: %v \n", err)
		return
	}

	logRecordSig := []byte("Record(address,address)")
	logRecordSigHash := crypto.Keccak256Hash(logRecordSig)

	for _, vLog := range logs {

		switch vLog.Topics[0].Hex() {
		case logRecordSigHash.Hex():
			//
			store := struct {
				Account   common.Address
				StoreAddr common.Address
			}{}
			store.Account = common.HexToAddress(vLog.Topics[1].Hex())
			store.StoreAddr = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("SubscribeEvents: account=%v storeAddr=%v\n", store.Account.String(), store.StoreAddr.String())
		default:
			//
			fmt.Printf("SubscribeEvents log topic: %v \n", vLog.Topics[0].Hex())
		}

	}
}
