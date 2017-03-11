package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {
	fmt.Println("started logging in func main()")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("failed in func main()")
	}
}

//Init comment
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")
	return nil, nil
}

//Invoke comment
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")
	return nil, nil
}

//Query comment
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")
	return nil, nil
}
