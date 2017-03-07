package main

import (
	"fmt"
	"math"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))

	fmt.Println("hello world!")
	fmt.Println("math.Pi is: ", math.Pi)
}
