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
var currentPrice float64
var currentSummary int

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
	fmt.Println("started to initialize project")
	if len(args) != 5 {
		return nil, errors.New("failed in args")
	}
	projectName = args[0]                    //项目名称
	projectRate, _ = strconv.Atoi(args[1])   //项目年化利率
	projectPeriod, _ = strconv.Atoi(args[2]) //项目天数
	projectGoal, _ = strconv.Atoi(args[3])   //项目目标募集金额
	projectTimes = 1                         //项目期数
	projectBenfiary = args[4]                //项目受益人
	projectState = "draft"                   //项目当前状态
	currentPrice = 1.0

	//started to initialize the table: Shares
	fmt.Println("started to create table: Shares")
	tableErr := stub.CreateTable("Shares", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "User", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Amount", Type: shim.ColumnDefinition_INT64, Key: false},
	})
	if tableErr != nil {
		return nil, errors.New("Failed creating AssetsOnwership table.")
	}

	return nil, nil
}

// Invoke function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")

	switch function {
	case "goPublic":
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		fmt.Println("started logging in goPublic()")
		projectState = "public"
		return nil, nil //end of goPublic
	case "pay":
		if len(args) != 2 {
			return nil, errors.New("failed in args")
		}
		fmt.Println("started logging in pay()")
		user := args[0]
		amount := args[1]

		fmt.Println("step1: chechk whether table contains user")
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: user}}
		columns = append(columns, col1)

		row, err := stub.GetRow("Shares", columns)
		if err != nil {
			return nil, errors.New("failed in function getRow")
		}

		if len(row.Columns) == 0 {
			fmt.Println("columns 0")
			fmt.Println("started to add a row")

			value, err := strconv.Atoi(amount)
			if err != nil {
				return nil, errors.New("errors in the amount")
			}
			currentSummary += value

			row := shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: user}},
					&shim.Column{Value: &shim.Column_Int64{Int64: int64(value)}}}}

			ok, err := stub.InsertRow("Shares", row)
			if !ok && err == nil {
				return nil, errors.New("shares was already assigned")
			}
		} else {
			fmt.Println("columns not 0")
			fmt.Println("todo we need to update")
		}

		return nil, nil //end of pay
	}

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
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		fmt.Println("started in function getProjectState")
		result := projectName + "/"
		result += strconv.Itoa(projectRate) + "/"
		result += strconv.Itoa(projectPeriod) + "/"
		result += strconv.Itoa(projectGoal) + "/"
		result += strconv.Itoa(projectTimes) + "/"
		result += projectBenfiary + "/"
		result += strconv.Itoa(currentSummary) + "/"
		result += projectState

		return []byte(result), nil //end of function getProjectState

	case "getRow":
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		fmt.Println("started logging in function getRow()")

		user := args[0]
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: user}}
		columns = append(columns, col1)

		row, err := stub.GetRow("Shares", columns)
		if err != nil {
			return nil, errors.New("failed in function getRow")
		}

		val0 := row.Columns[0].GetString_()
		val1 := row.Columns[1].GetInt64()
		result := "****" + val0 + "****" + strconv.Itoa(val1)
		return []byte(result), err //end of function getRow()
	}
	return nil, nil
}
