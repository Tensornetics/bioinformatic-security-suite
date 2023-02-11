package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    sc "github.com/hyperledger/fabric/protos/peer"
)

// User defines the structure of user data
type User struct {
    Name string `json:"name"`
    Role string `json:"role"`
    Salt string `json:"salt"`
    Hash string `json:"hash"`
}

// BioSecurityChaincode implements the ChaincodeInterface
type BioSecurityChaincode struct {}

// Init performs any necessary setup or initialization of the chaincode state
func (s *BioSecurityChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
    return shim.Success(nil)
}

// Invoke performs the appropriate business logic based on the function name and arguments passed in
func (s *BioSecurityChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "addUser" {
        return s.addUser(stub, args)
    } else if function == "getUser" {
        return s.getUser(stub, args)
    } else {
        return shim.Error("Invalid function name.")
    }
}

// addUser adds a new user to the state database
func (s *BioSecurityChaincode) addUser(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    // Validate user inputs
    if len(args) != 4 {
        return shim.Error("Incorrect number of arguments. Expecting 4.")
    }
    user := &User{
        Name: args[0],
        Role: args[1],
        Salt: args[2],
        Hash: args[3],
    }
    userBytes, err := json.Marshal(user)
    if err != nil {
        return shim.Error("Error marshalling user data.")
    }
    err = stub.PutState(user.Name, userBytes)
    if err != nil {
        return shim.Error(fmt.Sprintf("Error storing user data: %s", err.Error()))
    }
    return shim.Success(nil)
}

// getUser retrieves the user data from the state database
func (s *BioSecurityChaincode) getUser(stub shim.ChaincodeStubInterface, args []string) sc.Response {
    // Validate user inputs
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1.")
    }
    userBytes, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(fmt.Sprintf("Error retrieving user data: %s", err.Error()))
    }
    if userBytes == nil {
        return shim.Error("User data not found.")
    }
    return shim.Success(userBytes)
}

