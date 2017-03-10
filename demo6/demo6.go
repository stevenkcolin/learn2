package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode basic struct
type SimpleChaincode struct {
}

func main() {
	fmt.Println("started logging in main()")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("failed in function main(), error is %v", err)
	}
}

//Init function
//Step1: 获得调用init()的caller, 并且保存在"admin"中
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")

	fmt.Println("started get caller's metatdat")
	adminCert, _ := stub.GetCallerMetadata()
	fmt.Printf("the administrator is [%v]", adminCert)

	// if err != nil {
	// 	return nil, errors.New("failed getting metadata")
	// }
	//
	// if len(adminCert) == 0 {
	// 	return nil, errors.New("Invalid admin certificate. Empty.")
	// }

	// stub.PutState("admin", adminCert)
	return nil, nil
}

// Invoke function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")

	return nil, nil
}

//Query funciton
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")

	switch function {
	case "getState":
		fmt.Println("started in function getState()")
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}

		key := args[0]
		result, err := stub.GetState(key)

		if err != nil {
			return nil, errors.New("Failed in function getState")
		}
		return result, err
	}
	return nil, nil
}
