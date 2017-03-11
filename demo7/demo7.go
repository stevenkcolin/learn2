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
var availableList map[string]int

//Init comment
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")

	//started to initialize the projectState
	//step1: check whether args == 5
	if len(args) != 5 {
		return nil, errors.New("failed in args")
	}

	//step 2, intialize the project properties
	projectName = args[0]
	stub.PutState("projectName", []byte(projectName))

	projectRate, _ = strconv.Atoi(args[1])
	if projectRate <= 0 {
		return nil, errors.New("errors in args[1], it cannot be negative")
	}
	stub.PutState("projectRate", []byte(args[1]))

	projectPeriod, _ = strconv.Atoi(args[2])
	if projectPeriod <= 0 {
		return nil, errors.New("errors in args[2], it cannot be negative")
	}
	stub.PutState("projectPeriod", []byte(args[2]))

	projectGoal, _ = strconv.Atoi(args[3])
	if projectGoal <= 0 {
		return nil, errors.New("errors in args[3], it cannot be negative")
	}
	stub.PutState("projectGoal", []byte(args[3]))

	projectTimes = 1
	stub.PutState("projectTimes", []byte(strconv.Itoa(projectTimes)))

	projectBenifary = args[4]
	stub.PutState("projectBenifary", []byte(args[4]))

	projectState = "draft"
	stub.PutState("projectState", []byte(projectState))

	currentPrice = 1.0
	projectSummary = 0

	return nil, nil
} //end of Init()

// Invoke comment
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
	case "pay":
		fmt.Println("started logging in pay()")

		if len(args) != 2 {
			return nil, errors.New("failed in args")
		}
		user := args[0]
		amount := args[1]
		fmt.Printf("user is %v and amount is %v", user, amount)
		amountInt, _ := strconv.Atoi(amount)
		if amountInt <= 0 {
			return nil, errors.New("amount is negative")
		}

		// if shareList[user] == 0 {
		// 	userList = append(userList, user)
		// }
		//
		// shareList[user] += amountInt
		projectSummary += amountInt
		return nil, nil
	case "checkGoalReached":
		fmt.Println("started logging in checkGoalReached()")
		// TODO: write code for checkGoalReached
		return nil, nil
	case "checkDaoqi":
		fmt.Println("started logging in checkDaoqi()")
		// TODO: write code for checkDaoqi
		return nil, nil
	case "calculatePrice":
		fmt.Println("started logging in calculatePrice")
		return nil, nil
	case "calculateResult":
		fmt.Println("started logging in calculateResult")
		return nil, nil
	}
	return nil, nil
}

// Query comment
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Query()")
	switch function {
	case "getProjectName":
		result, err := stub.GetState("projectName")
		if err != nil {
			return nil, errors.New("error in getProjectName")
		}
		return result, nil
	case "getProjectGoal":
		result, err := stub.GetState("projectGoal")
		if err != nil {
			return nil, errors.New("error in getProjectGoal")
		}
		return result, nil
	case "getProjectSummary":
		result := strconv.Itoa(projectSummary)
		return []byte(result), nil
	case "getUserList":
		fmt.Println("started logging in getUserList")
		var result string
		for i, value := range userList {
			result += fmt.Sprintf("userList[%v] is %v ****", strconv.Itoa(i), value)
		}
		return []byte(result), nil
	case "getShareList":
		fmt.Println("started logging in getShareList")
		return nil, nil
	case "getAvailableList":
		fmt.Println("started logging in getAvailableList")
		return nil, nil
	}
	return nil, nil
}
