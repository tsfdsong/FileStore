package storecontract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GetClient ...
func GetClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return client, nil
}

//CreateStoreContract ...
func CreateStoreContract(url, keystore string) (common.Address, error) {
	client, err := GetClient(url)
	if err != nil {
		return common.Address{}, err
	}

	session, err := NewSession(client, keystore)
	if err != nil {
		return common.Address{}, err
	}

	var addr common.Address
	session, addr, err = NewContract(session, client)
	if err != nil {
		return common.Address{}, err
	}

	if (addr != common.Address{}) {
		return addr, nil
	}

	return common.Address{}, fmt.Errorf("Deploy account manager contract err")
}

//GetStoreContract ...
func GetStoreContract(url, keystore string, contractAddr common.Address) (*FileStoreSession, error) {
	client, err := GetClient(url)
	if err != nil {
		return nil, err
	}

	session, err := NewSession(client, keystore)
	if err != nil {
		return nil, err
	}

	//Get contract
	session, err = LoadContractByAddr(session, client, contractAddr)
	if err != nil {
		return nil, err
	}

	return session, nil
}

//NewSession ...
func NewSession(client *ethclient.Client, keypath string) (*FileStoreSession, error) {
	privateKey, err := crypto.HexToECDSA(keypath)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)

	return &FileStoreSession{
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Nonce:    nil, // nil uses nonce of pending state
			Signer:   auth.Signer,
			Value:    big.NewInt(0),
			GasPrice: nil,             // nil automatically suggests gas price
			GasLimit: uint64(1000000), // 0 automatically estimates gas limit
		},
	}, nil
}

//NewContract ...
func NewContract(session *FileStoreSession, client *ethclient.Client) (*FileStoreSession, common.Address, error) {
	address, _, instance, err := DeployFileStore(&session.TransactOpts, client)
	if err != nil {
		fmt.Printf("NewContract deploy store contract: %v\n", err)
		return nil, common.Address{}, err
	}

	session.Contract = instance

	fmt.Printf("NewContract new file store contract: %v\n", address.Hex())

	return session, address, nil
}

//LoadContractByAddr load a contract if one exists
func LoadContractByAddr(session *FileStoreSession, client *ethclient.Client, contractAddr common.Address) (*FileStoreSession, error) {
	inst, err := NewFileStore(contractAddr, client)
	if err != nil {
		return nil, err
	}

	session.Contract = inst

	return session, nil
}

//CheckTransactStatus ...
func CheckTransactStatus(url string, trx *types.Transaction) (bool, error) {
	client, err := GetClient(url)
	if err != nil {
		return false, err
	}

	fmt.Printf("trx hash: %v\n", trx.Hash().String())
	receipt, err := client.TransactionReceipt(context.Background(), trx.Hash())
	if err != nil {
		return false, err
	}

	fmt.Printf("receipt: %v\n", receipt)

	if receipt.Status == 1 {
		return true, nil
	}

	return false, fmt.Errorf("trx receipt failed: %v", trx.Hash().String())
}
