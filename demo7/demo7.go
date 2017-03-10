package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {
	fmt.Println("started logging in main()")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("failed in function main()")
	}
}

var projectName string
var projectRate int
var projectPeriod int
var projectGoal int
var projectTimes int
var projectBenifary string
var projectState string
var currentPrice float64
var projectSummary int
var userList []string
var shareList map[string]int

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")

	//started to initialize the projectState
	//step1: check whether args == 5
	if len(args) != 5 {
		return nil, errors.New("failed in args")
	}

	//step 2, intialize the project properties
	projectName = args[0]
	projectRate, _ = strconv.Atoi(args[1])
	if projectRate <= 0 {
		return nil, errors.New("errors in args[1], it cannot be negative")
	}

	projectPeriod, _ = strconv.Atoi(args[2])
	if projectPeriod <= 0 {
		return nil, errors.New("errors in args[2], it cannot be negative")
	}

	projectGoal, _ = strconv.Atoi(args[3])
	if projectGoal <= 0 {
		return nil, errors.New("errors in args[3], it cannot be negative")
	}

	projectTimes = 1
	projectBenifary = args[4]
	projectState = "draft"
	currentPrice = 1.0
	projectSummary = 0

	return nil, nil
} //end of Init()

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")
	switch function {
	case "goPublic":
		fmt.Println("started logging in goPublic")
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		projectState = "public"
		return nil, nil //end of goPublic
	}
	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")
	switch function {
	case "getProjectState":
		fmt.Println("started logging in getProjectState()")
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		result := projectName + "/"
		result += strconv.Itoa(projectRate) + "/"
		result += strconv.Itoa(projectPeriod) + "/"
		result += strconv.Itoa(projectGoal) + "/"
		result += strconv.Itoa(projectTimes) + "/"
		result += projectBenifary + "/"
		result += projectState + "/"
		result += strconv.FormatFloat(currentPrice, 'E', -1, 64) + "/"
		result += strconv.Itoa(projectSummary) + "/"

		return []byte(result), nil //end of getProjectState
	}
	return nil, nil
}
