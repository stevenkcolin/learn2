package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var projectName string
var projectRate int
var projectPeriod int
var projectGoal int
var projectTimes int
var projectBenifary string
var projectState string
var currentPrice float64
var projectSummary int
var projectDeadline int64
var shareList map[string]int
var isProjectStarted bool
var isDue bool
var isFinished bool

var availableList map[string]float64

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
	//step1: initialize the ShareList & AvailableList
	shareList = make(map[string]int)
	availableList = make(map[string]float64)
	//step2: check whether args == 5
	if len(args) != 5 {
		return nil, errors.New("failed in args")
	}

	//step3: initialize the project properties
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
	projectState = "public"
	currentPrice = 1.0
	projectSummary = 0
	projectDeadline = 9999999999

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
	switch function {
	case "getProjectName":
		fmt.Println("started logging in func getProjectName")
		return []byte(projectName), nil //end of func getProjectName()
	case "getProjectGoal":
		fmt.Println("started logging in func getProjectGoal")
		result := strconv.Itoa(projectGoal)
		return []byte(result), nil //end of func getProjectGoal()
	case "getProjectSummary":
		fmt.Println("started logging in func getProjectSummary")
		result := strconv.Itoa(projectSummary)
		return []byte(result), nil //end of func getProjectSummary()
	case "getShareList":
		fmt.Println("started logging in func getShareList")
		var result string
		for user, amount := range shareList {
			result += "****" + user + "/" + strconv.Itoa(amount)
		}
		return []byte(result), nil //end of func getShareList

	case "getProjectDeadline":
		fmt.Println("started logging in func getProjectDeadline")
		result := strconv.Itoa(int(projectDeadline))
		return []byte(result), nil //end of func getProjectDeadline

	case "getCurrentPrice":
		fmt.Println("started logging in func getCurrentPrice")
		result := strconv.FormatFloat(currentPrice, 'E', -5, 64)
		return []byte(result), nil //end of func getCurrentPrice

	case "getAvailableList":
		fmt.Println("started logging in func getAvailableList")
		var result string
		for user, value := range availableList {
			result += "****" + user + "/" + strconv.FormatFloat(value, 'E', -5, 64)
		}
		return []byte(result), nil //end of func getAvailableList

	}
	return nil, nil
}
