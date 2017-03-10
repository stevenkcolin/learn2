package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

type SimpleChaincode struct {
}

var myLogger = logging.MustGetLogger("asset_mgm")

func main() {
	fmt.Println("started logging in main()")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("failed in function main(), error is %v", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")
	myLogger.Debug("init chaincode")
	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")
	myLogger.Debug("start Invoke")
	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")
	myLogger.Debug("started logging in Query()")
	return nil, nil
}
