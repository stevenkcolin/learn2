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

		//step1 : get args
		user := args[0]
		amount, _ := strconv.Atoi(args[1])
		if amount <= 0 {
			return nil, errors.New("errors in args[1],it is negative")
		}
		fmt.Printf("user is %v and amount is %v", args[0], args[1])

		//step2 : check if the userList[user] exist
		//if exist, then userList[user] += amount
		//if not exists, then userList[user] = amount
		// if shareList[user] == 0 {
		// 	fmt.Printf("the user [%v] does not exit \n", user)
		// 	userList = append(userList, user)
		// 	shareList[user] = amount
		//
		// } else {
		// 	fmt.Printf("the user [%v] exist", user)
		// 	if !GoalReached() {
		// 		shareList[user] += amount
		// 	}
		// }

		// step3: raise amount

		//// TODO: write code for pay
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
	default:
		fmt.Println("no function found")
		return nil, errors.New("no function found, recheck your function name")
	}
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
	// case "getProjectState":
	// 	fmt.Println("started logging in getProjectState()")
	// 	if len(args) != 0 {
	// 		return nil, errors.New("incorrect args")
	// 	}
	// 	result := projectName + "/"
	// 	result += strconv.Itoa(projectRate) + "/"
	// 	result += strconv.Itoa(projectPeriod) + "/"
	// 	result += strconv.Itoa(projectGoal) + "/"
	// 	result += strconv.Itoa(projectTimes) + "/"
	// 	result += projectBenifary + "/"
	// 	result += projectState + "/"
	// 	result += strconv.FormatFloat(currentPrice, 'E', -1, 64) + "/"
	// 	result += strconv.Itoa(projectSummary) + "/"
	//
	// 	return []byte(result), nil //end of getProjectState
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
	default:
		fmt.Println("no function found")
		return nil, errors.New("no function found, recheck your function name")
	}
}

func GoalReached() bool {
	return false
}
