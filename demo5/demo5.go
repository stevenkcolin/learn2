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
	switch function {
	case "putState":
		if len(args) != 2 {
			return nil, errors.New("incorrect args")
		}
		key := args[0]
		value := []byte(args[1])
		err := stub.PutState(key, value)
		if err != nil {
			return nil, err
		}
		return value, err
	case "delState":
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		key := args[0]
		err := stub.DelState(key)
		if err != nil {
			return nil, err
		}
		result := key + "has been deleted"
		return []byte(result), err

	case "createTable":
		err := stub.CreateTable("AssetsOwnership", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "Asset", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "Owner", Type: shim.ColumnDefinition_BYTES, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating AssetsOnwership table.")
		}
		return nil, nil
	}

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

	case "getState":
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		key := args[0]
		result, err := stub.GetState(key)

		if err != nil {
			return nil, errors.New("Failed in function getState")
		}
		return result, err

	case "getCallerCert":
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		result, err := stub.GetCallerCertificate()
		if err != nil {
			return nil, errors.New("Failed in function getCallerCert")
		}
		return result, err

	case "getCallerMetadata":
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		result, err := stub.GetCallerMetadata()
		if err != nil {
			return nil, errors.New("Failed in function getCallerMetadata")
		}
		return result, err

	}
	return nil, nil
}
