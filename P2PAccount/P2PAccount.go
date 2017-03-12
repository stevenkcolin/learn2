package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

var balancesOf map[string]float64

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
	if len(args) != 0 {
		return nil, errors.New("errors in args")
	}
	balancesOf = make(map[string]float64)
	balancesOf["admin"] = 10000.0
	return nil, nil
}

//Invoke comment
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")
	switch function {
	case "mint":
		fmt.Println("started logging in func mint()")
		if len(args) != 1 {
			return nil, errors.New("error in args")
		}
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return nil, errors.New("error in args")
		}
		if amount <= 0 {
			return nil, errors.New("errors in args[0], it cannot be zero or negative")
		}
		balancesOf["admin"] += amount
		return nil, nil
	}
	return nil, nil
}

//Query comment
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")
	switch function {
	case "getBalances":
		fmt.Println("started logging in func getBalances")
		result, err := json.Marshal(balancesOf)
		if err != nil {
			return nil, errors.New("errors in getBalances")
		}
		return result, nil //end of getBalances
	}
	return nil, nil
}
