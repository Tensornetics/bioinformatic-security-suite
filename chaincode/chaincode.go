package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type BioSecurityChaincode struct {
}

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Salt string `json:"salt"`
	Hash string `json:"hash"`
}

func (s *BioSecurityChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *BioSecurityChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "addUser" {
		return s.addUser(stub, args)
	} else if fn == "getUser" {
		return s.getUser(stub, args)
	}

	return shim.Error("Invalid function name")
}

func (s *BioSecurityChaincode) addUser(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments")
	}

	user := &User{
		Name: args[0],
		Role: args[1],
		Salt: args[2],
		Hash: args[3],
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error("Error marshalling user")
	}
	stub.PutState(user.Name, userBytes)
	return shim.Success(nil)
}
