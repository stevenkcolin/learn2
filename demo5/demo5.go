package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))

	if err != nil {
		//do nothing
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case "getTxID":
		txID := stub.GetTxID()
		result := []byte(txID)
		return result, nil
	case "getTxTimestamp":
		time, err := stub.GetTxTimestamp()
		result := []byte(time.String()) //时间转换为字符串，time.String()
		return result, err
	case "getStringArgs":
		strList := stub.GetStringArgs()
		var result string
		for index := 0; index < len(strList); index++ {
			result += "***" + strList[index] + "***"
		}
		return []byte(result), nil
	case "putState":
		if len(args) != 2 {
			return nil, errors.New("incorrect args")
		}
		key := args[0]
		value := []byte(args[1])
		err := stub.PutState(key, value)
		return nil, err
	case "getState":
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		key := args[0]
		result, err := stub.GetState(key)
		return result, err
	}
	return nil, nil
}
