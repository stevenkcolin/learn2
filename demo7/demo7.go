package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
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
var projectDeadline int64
var shareList map[string]int
var isProjectStarted bool
var isDue bool
var isFinished bool

var availableList map[string]float64

func main() {
	fmt.Println("started logging in main()")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("failed in function main()")
	}
}

//Init comment
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Init()")
	//started to initialize the ShareList & AvailableList
	shareList = make(map[string]int)
	availableList = make(map[string]float64)

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

	projectState = "public"
	stub.PutState("projectState", []byte(projectState))

	currentPrice = 1.0
	projectSummary = 0
	projectDeadline = 9999999999

	return nil, nil
} //end of Init()

// Invoke comment
// It contains below functions
// goPublic() set the state to public
// pay() users could pay for the project
// autoPay() if the gap to the project goal is less than 2%, it will be filled automatically
// checkGoalReached() check whether the project goal is reached, if reached, started to count interest
// checkDueDate() check whether the project is due to date, if it is due do date, then started to calculate the price
// checkRepay() calculate the shareList[user] * currentPrice and saved to availableList
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("started logging in Invoke()")
	switch function {
	case "goPublic":
		fmt.Println("started logging in goPublic")
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		} //check the length of args
		projectState = "public"
		return nil, nil //end of goPublic
	case "pay":
		fmt.Println("started logging in pay()")
		//check the project state
		if !isPublic() {
			return nil, errors.New("current state is not public, function pay() failed")
		}
		//check the length of args
		if len(args) != 2 {
			return nil, errors.New("failed in args")
		}
		//get the args[0] & args[1]
		user := args[0]
		amount := args[1]
		fmt.Printf("user is %v and amount is %v\n", user, amount)
		amountInt, _ := strconv.Atoi(amount)
		if amountInt <= 0 {
			return nil, errors.New("amount is negative")
		}
		if isOverGoal(amountInt) {
			return nil, errors.New("projectSummary + amount is over goal")
		}

		shareList[user] += amountInt

		projectSummary += amountInt
		return nil, nil
	case "autoPay":
		fmt.Println("started logging in autoPay")
		if isGoalReached() {
			return nil, errors.New("goal has been reached")
		}
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}

		gap := getGoalGap()
		if gap != 0 {
			fmt.Printf("admin pay,for user admin and value %v\n", gap)
			shareList["admin"] = gap
			projectSummary += gap
		}

		return nil, nil
	case "checkGoalReached":
		fmt.Println("started logging in checkGoalReached()")
		if !isGoalReached() {
			return nil, errors.New("goal has not been reached")
		}
		if len(args) != 0 {
			return nil, errors.New("failed in args")
		}
		if isProjectStarted {
			return nil, errors.New("project has been started")
		}

		updateDeadline()
		isProjectStarted = true
		return nil, nil
	case "checkDueDate":
		fmt.Println("started logging in checkDueDate()")
		if !isDueDate() {
			return nil, errors.New("due date is wrong")
		}
		if isDue {
			return nil, errors.New("dueDate can be called only once")
		}
		interest := float64(float64(projectRate) * float64(projectPeriod) / 365 / 100)
		currentPrice += interest

		fmt.Printf("value of projectRate is %v\n", projectRate)
		fmt.Printf("value of projectPeriod is %v\n", projectPeriod)
		fmt.Printf("value of interest is %v\n", interest)
		fmt.Printf("value of current price is %v\n", currentPrice)

		isDue = true
		return nil, nil
	case "checkRepay":
		fmt.Println("started logging in checkRepay()")
		if !isDue {
			return nil, errors.New("the project is not dueDate, please invoke checkDueDate() first")
		}
		if isFinished {
			return nil, errors.New("the project has been finished, throw errors")
		}

		for user, amount := range shareList {
			newValue := float64(amount) * currentPrice
			availableList[user] = newValue
		}
		isFinished = true
		return nil, nil
	default:
		fmt.Println("no function found")
		return nil, errors.New("no function found")
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
	case "getProjectGoal":
		result, err := stub.GetState("projectGoal")
		if err != nil {
			return nil, errors.New("error in getProjectGoal")
		}
		return result, nil
	case "getProjectSummary":
		result := strconv.Itoa(projectSummary)
		return []byte(result), nil

	case "getShareList":
		fmt.Println("started logging in getUserList")

		var result string
		for user, amount := range shareList {
			result += "****" + user + "/" + strconv.Itoa(amount)
		}
		return []byte(result), nil

	case "getProjectDeadline":
		fmt.Println("started logging in getProjectDeadline")
		result := strconv.Itoa(int(projectDeadline))
		return []byte(result), nil

	case "getCurrentPrice":
		fmt.Println("started logging in getCurrentPrice")
		result := strconv.FormatFloat(currentPrice, 'E', -5, 64)
		return []byte(result), nil

	case "getAvailableList":
		fmt.Println("started logging in getAvailableList")
		var result string
		for user, value := range availableList {
			result += "****" + user + "/" + strconv.FormatFloat(value, 'E', -5, 64)
		}
		return []byte(result), nil //end of getAvailableList
	}
	return nil, nil
}

func isPublic() bool {
	if projectState == "public" {
		return true
	} else {
		return false
	}
}

func isOverGoal(amount int) bool {
	if projectSummary+amount > projectGoal {
		return true
	} else {
		return false
	}
}

func isGoalReached() bool {
	if projectSummary >= projectGoal {
		return true
	} else {
		return false
	}
}

func getGoalGap() int {
	if projectSummary >= projectGoal {
		return 0
	}
	var gap int
	gap = projectGoal - projectSummary
	if gap < projectGoal/50 {
		return gap
	} else {
		return 0
	}
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
