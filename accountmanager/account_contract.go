package accountmanager

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

//CreateAccountContract ...
func CreateAccountContract(url, keystore string) (common.Address, error) {
	client, err := GetClient(url)
	if err != nil {
		return common.Address{}, err
	}

	session, err := NewAccountSession(client, keystore)
	if err != nil {
		return common.Address{}, err
	}

	var addr common.Address
	session, addr, err = NewAccountContract(session, client)
	if err != nil {
		return common.Address{}, err
	}

	if (addr != common.Address{}) {
		return addr, nil
	}

	return common.Address{}, fmt.Errorf("Deploy account manager contract err")
}

//GetAccountContract ...
func GetAccountContract(url, keystore string, contractAddr common.Address) (*AccountManagerSession, error) {
	client, err := GetClient(url)
	if err != nil {
		return nil, err
	}

	session, err := NewAccountSession(client, keystore)
	if err != nil {
		return nil, err
	}

	//Load or Deploy contract
	session, err = LoadAccountContractByAddr(session, client, contractAddr)
	if err != nil {
		return nil, err
	}

	return session, nil
}

//NewAccountSession ...
func NewAccountSession(client *ethclient.Client, keypath string) (*AccountManagerSession, error) {
	privateKey, err := crypto.HexToECDSA(keypath)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)

	return &AccountManagerSession{
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Nonce:    nil, // nil uses nonce of pending state
			Signer:   auth.Signer,
			Value:    big.NewInt(0),
			GasPrice: nil,            // nil automatically suggests gas price
			GasLimit: uint64(300000), // 0 automatically estimates gas limit
		},
	}, nil
}

//NewAccountContract ...
func NewAccountContract(session *AccountManagerSession, client *ethclient.Client) (*AccountManagerSession, common.Address, error) {

	address, _, instance, err := DeployAccountManager(&session.TransactOpts, client)
	if err != nil {
		return nil, common.Address{}, err
	}

	session.Contract = instance

	fmt.Printf("NewAccountContract new contract: %v\n", address.Hex())

	return session, address, nil
}

//LoadAccountContractByAddr load a contract if one exists
func LoadAccountContractByAddr(session *AccountManagerSession, client *ethclient.Client, contractAddr common.Address) (*AccountManagerSession, error) {
	inst, err := NewAccountManager(contractAddr, client)
	if err != nil {
		return nil, err
	}

	session.Contract = inst

	return session, nil
}
