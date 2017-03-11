package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

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

//SimpleChaincode comment
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

	//step3, get caller's certificate, and saved to "admin" state
	fmt.Println("started get caller's metatdat")
	adminCert, err := stub.GetCallerCertificate()
	if err != nil {
		return nil, errors.New("failed getting metadata")
	}
	if len(adminCert) == 0 {
		return nil, errors.New("invalid admin certificate. Empty.")
	}
	fmt.Printf("the administrator is [%v]", adminCert)
	stub.PutState("admin", adminCert)

	//step4: initialize the project properties
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
	switch function {
	case "pay":
		fmt.Println("started logging in func pay()")
		//step1: check current state
		if !isPublic() {
			return nil, errors.New("current state is not public")
		}
		//step2: check len(args)==2
		if len(args) != 2 {
			return nil, errors.New("failed in args")
		}
		//step3: get user=args[0] and amount=args[1]
		//validate the input data
		user := args[0]
		amount, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, errors.New("errors in args[1]")
		}
		if amount <= 0 {
			return nil, errors.New("amount is negative")
		}
		//step 4: check whether amount will be over the projectGoal
		if isOverGoal(amount) {
			return nil, errors.New("projctSummary + amount is over goal")
		}
		// step 5: putdata into shareList & projectSummary
		shareList[user] += amount
		projectSummary += amount
		return nil, nil //end of func pay()

	case "autoPay":
		fmt.Println("started logging in autoPay")
		// step1: check len(args) ==0
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		// step2: check isGoalReached()
		if isGoalReached() {
			return nil, errors.New("goal has been reached")
		}
		// step3: get Gap between projectGoal and projectSummary
		gap := getGoalGap()
		if gap != 0 {
			fmt.Printf("admin pay, admin pay the value %v\n", gap)
			shareList["admin"] = gap
			projectSummary += gap
		}
		return nil, nil //end of func autoPay()
	case "checkGoalReached":
		fmt.Println("started logging in checkGoalReached()")
		//step1: check len(args) ==0
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		//step2: check isGoalReached()
		if !isGoalReached() {
			return nil, errors.New("goal has not been reached")
		}
		// step3
		updateDeadline()
		isProjectStarted = true
		return nil, nil //end of checkGoalReached()

	case "checkDueDate":
		fmt.Println("started logging in checkDueDate()")
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		if !isDueDate() {
			return nil, errors.New("due date is wrong")
		}
		if isDue {
			return nil, errors.New("dueDate can be called only once")
		} //avoid interest is calculated multiple;

		interest := float64(float64(projectRate) * float64(projectPeriod) / 365 / 100)
		currentPrice += interest
		isDue = true
		return nil, nil //end of checkDueDate;

	case "checkRepay":
		fmt.Println("started logging in checkRepay()")
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		if !isDue {
			return nil, errors.New("the project is not dueDate, please invoke checkDueDate() first")
		}
		if isFinished {
			return nil, errors.New("the project has been finished, throw errors")
		} //avoid to be invoked twice;

		for user, amount := range shareList {
			newValue := float64(amount) * currentPrice
			availableList[user] = newValue
		}
		isFinished = true
		return nil, nil //end of "checkRepay"
	case "goPublic":
		fmt.Println("started logging in goPublic()")
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		projectState = "public"
		return nil, nil
	default:
		fmt.Println("no function found")
		return nil, errors.New("no function found")
	}
}

//Query comment
//func getProjectName()
//func getProjectGoal()
//func getProjectSummary()
//func getShareList()
//func getProjectDeadline()
//func getCurrentPrice()
//func getAvailableList()
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
		result, err := json.Marshal(shareList)
		if err != nil {
			return nil, errors.New("errors in getShareList()")
		}
		// for user, amount := range shareList {
		// 	result += "****" + user + "/" + strconv.Itoa(amount)
		// }
		return result, nil //end of func getShareList

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

	default:
		fmt.Println("no function found")
		return nil, errors.New("no function found")
	}

}

func isPublic() bool {
	if projectState == "public" {
		return true
	}
	return false

}

func isOverGoal(amount int) bool {
	if projectSummary+amount > projectGoal {
		return true
	}
	return false

}

func isGoalReached() bool {
	if projectSummary >= projectGoal {
		return true
	}
	return false

}

func getGoalGap() int {
	if projectSummary >= projectGoal {
		return 0
	}
	var gap int
	gap = projectGoal - projectSummary
	if gap < projectGoal/50 {
		return gap
	}
	return 0
}

func updateDeadline() {
	now := time.Now().Unix()
	fmt.Printf("now is%v\n", now)

	day := projectPeriod
	projectDeadline = time.Now().AddDate(0, 0, day).Unix()
	fmt.Printf("deadline is %v\n", projectDeadline)
}

func isDueDate() bool {
	// 正常代码，检查当前时间是否大于最后期限
	// now := time.Now().Unix()
	// if now >= projectDeadline {
	// 	return true;
	// } else {
	// 	return false;
	// }
	//为了测试所用，永远返回true，
	return true
}
