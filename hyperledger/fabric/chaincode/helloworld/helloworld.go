package main

import (
	"errors"
	"fmt"
	//"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// HelloWorldChaincode Example chaincode for test
type HelloWorldChaincode struct {
}

// Init function to be invoked when chaincode is deployed
func (t *HelloWorldChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("HelloWorld - Init called with function %s!\n", function)

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments.")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke function to be invoked when chaincode is invoked
func (t *HelloWorldChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("HelloWorld - Invoke called with function %s!\n", function)

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments.")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Query function to be invoked when chaincode is queried
func (t *HelloWorldChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("HelloWorld - Query called with function %s!\n", function)

	message, err := stub.GetState("hello_world")
	if err != nil {
		return nil, err
	}

	return []byte(message), nil
}

func main() {
	err := shim.Start(new(HelloWorldChaincode))
	if err != nil {
		fmt.Printf("Error starting Hello World chaincode: %s", err)
	}
}
