package main

import (
	"errors"
	"fmt"
	"strconv"

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

var projectName string
var projectRate int
var projectPeriod, projectGoal, projectTimes int
var projectBenfiary string
var projectState string

//Init function
//Step1: 获得调用init()的caller, 并且保存在"admin"中
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")

	//get caller's certificate, and saved to "admin" state
	fmt.Println("started get caller's metatdat")
	adminCert, err := stub.GetCallerCertificate()
	if err != nil {
		return nil, errors.New("failed getting metadata")
	}
	if len(adminCert) == 0 {
		return nil, errors.New("Invalid admin certificate. Empty.")
	}
	fmt.Printf("the administrator is [%v]", adminCert)
	stub.PutState("admin", adminCert)

	//started to initialize project
	projectName = args[0]                    //项目名称
	projectRate, _ = strconv.Atoi(args[1])   //项目年化利率
	projectPeriod, _ = strconv.Atoi(args[2]) //项目天数
	projectGoal, _ = strconv.Atoi(args[3])   //项目目标募集金额
	projectTimes = 1                         //项目期数
	projectBenfiary = args[4]                //项目受益人
	projectState = "Draft"                   //项目当前状态

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
		return result, err //end of getState
	case "getProjectState":
		fmt.Println("started in function getProjectState")
		result += projectName + "/"
		result += strconv.Itoa(projectRate) + "/"
		result += strconv.Itoa(projectPeriod) + "/"
		result += strconv.Itoa(projectGoal) + "/"
		result += strconv.Itoa(projectTimes) + "/"
		result += projectBenfiary + "/"
		result += projectState

		return result, nil

	}
	return nil, nil
}
