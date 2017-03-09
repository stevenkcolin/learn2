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

func main() {
	err := shim.Start(new(SimpleChaincode))

	if err != nil {
		//do nothing
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// added by chenlin@20170308
	if len(args) != 0 {
		return nil, errors.New("incorrect args")
	}
	// adminCert, err := stub.GetCallerMetadata()
	// if err != nil {
	// 	return nil, errors.New("failed getting metadata")
	// }
	// if len(adminCert) == 0 {
	// 	return nil, errors.New("invalid admin certificate. Empty")
	// }
	stub.PutState("admin", []byte("hahahahaha"))
	stub.PutState("steve", []byte("welcome"))
	stub.PutState("steve2", []byte("welcome2"))
	err := stub.CreateTable("AssetsOwnership", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Asset", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Owner", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating AssetsOnwership table.")
	}
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
			&shim.ColumnDefinition{Name: "Owner", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failed creating AssetsOnwership table.")
		}
		return nil, nil

	case "insertTable":
		if len(args) != 2 {
			return nil, errors.New("incorrect args")
		}
		asset := args[0]
		owner := args[1]
		row := shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: asset}},
				&shim.Column{Value: &shim.Column_String_{String_: owner}}}}
		ok, err := stub.InsertRow("AssetsOwnership", row)
		if !ok && err == nil {
			return nil, errors.New("asset was already assigned")
		}
		return nil, nil //end of insertTable

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
		fmt.Println(result)
		return result, err
		// return []byte("getCallerMetadata"), err
	case "getBinding":
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		result, err := stub.GetBinding()

		if err != nil {
			return nil, errors.New("Failed in function getBinding")
		}
		fmt.Println(result)
		return result, err
		// return []byte("getBinding"), err
	case "getPayload":
		if len(args) != 0 {
			return nil, errors.New("incorrect args")
		}
		result, err := stub.GetPayload()

		if err != nil {
			return nil, errors.New("Failed in function getPayload")
		}
		fmt.Println(result)
		return result, err
	// return []byte("getPayload"), err
	case "getRow":
		fmt.Println("*********************")
		fmt.Println("*********************")
		fmt.Println("*********************")
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		asset := args[0]
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: asset}}
		columns = append(columns, col1)

		row, err := stub.GetRow("AssetsOwnership", columns)
		if err != nil {
			return nil, errors.New("failed in function getRow")
		}

		val0 := row.Columns[0].GetString_()
		val1 := row.Columns[1].GetString_()
		result := "****" + val0 + "****" + val1
		return []byte(result), err
	//end of "getRow"

	case "getRowDefinition":
		if len(args) != 2 {
			return nil, errors.New("incorrect args")
		}
		asset := args[0]
		owner := args[1]

		col0 := shim.Column{Value: &shim.Column_String_{String_: asset}} //col0，type=shim.column
		col1 := shim.Column{Value: &shim.Column_String_{String_: owner}} //col1, type=shim.column
		var columns []shim.Column                                        //define columns=[]shim.column
		columns = append(columns, col0)
		columns = append(columns, col1) //完成columns的赋值

		val0 := columns[0].GetString_()
		val1 := columns[1].GetString_()
		result := val0 + "****" + val1
		return []byte(result), nil

	case "getRows":
		fmt.Println("begin logs for getRows")
		if len(args) != 1 {
			return nil, errors.New("incorrect args")
		}
		asset := args[0]
		fmt.Printf("value of args[0] is: %v\n", asset)
		var columns []shim.Column
		col0 := shim.Column{
			Value: &shim.Column_String_{String_: asset},
		}
		columns = append(columns, col0)
		fmt.Printf("value of column[0] is %v\n", columns[0].GetString_())

		jsonColumns, _ := json.Marshal(columns)
		fmt.Printf("value of columns is: %v\n", []byte(jsonColumns))

		rowChannel, _ := stub.GetRows("AssetsOwnership", columns)

		var rows []shim.Row
		for {
			select {
			case row, ok := <-rowChannel:
				if !ok {
					rowChannel = nil
				} else {
					rows = append(rows, row)
				}
			}
			if rowChannel == nil {
				break
			}
		}
		// jsonRows, err := json.Marshal(rows)
		// if err != nil {
		// 	return nil, fmt.Errorf("getRowsTableTwo operation failed. Error marshaling JSON: %s", err)
		// }

		result := strconv.Itoa(len(rows))
		return []byte(result), nil //end of function getRows()
	}

	return nil, nil
}
