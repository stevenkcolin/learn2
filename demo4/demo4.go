package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("incorrect number of args")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running" + function)

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func:" + function)

	return nil, errors.New("received unknow function" + function)
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function) //log
	switch function {
	case "read":
		return t.read(stub, args)
	case "read2":
		return t.read2(stub, args)
	case "read3":
		return t.read3(stub)
	case "helloworld":
		return t.helloworld()
	default:
		fmt.Println("query did not find func: " + function)
	} //根据输入的string类型参数function, 来决定使用哪个子函数。

	return nil, errors.New("received unknow function: " + function) //如果没有的话，则返回失败
}

func (t *SimpleChaincode) read2(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	key := args[0]
	result := []byte("helloworld" + key)
	var err error
	return result, err
}

func (t *SimpleChaincode) read3(stub shim.ChaincodeStubInterface) ([]byte, error) {
	result := []byte("****" + stub.GetTxID() + "*****")
	return result, nil
}

func (t *SimpleChaincode) readValueOfHelloworld(stub shim.ChaincodeStubInterface) ([]byte, error) {
	result, error := stub.GetState("hello_world")
	return result, error
}

func (t *SimpleChaincode) helloworld() ([]byte, error) {
	result := []byte("hello world")
	return result, nil
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	if len(args) != 1 {
		return nil, errors.New("incorrect number of arguments. expecting 1")
	}
	key = args[0]
	valAsbytes, err := stub.GetState(key)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	key = args[0]
	value = args[1]

	err = stub.PutState(key, []byte(value))

	if err != nil {
		return nil, err
	}

	return nil, nil
}
