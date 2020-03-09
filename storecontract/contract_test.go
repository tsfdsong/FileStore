package storecontract

import (
	"fmt"
	"testing"
)

func Test_CreateStoreContract(t *testing.T) {
	contract, err := CreateStoreContract("http://127.0.0.1:8545", "")
	if err != nil {
		t.Errorf("CreateStoreContract: %v \n", err)
		return
	}

	fmt.Printf("CreateStoreContract addr: %v\n", contract.String())
	t.Logf("CreateStoreContract result: %v\n", contract.String())
}
