package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type BioInfo struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Hash       string `json:"hash"`
	Status     string `json:"status"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryBioInfo" {
		return s.queryBioInfo(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createBioInfo" {
		return s.createBioInfo(APIstub, args)
	} else if function == "changeBioInfoStatus" {
		return s.changeBioInfoStatus(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryBioInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	bioInfoAsBytes, _ := APIstub.GetState(args[0])
	if bioInfoAsBytes == nil {
		return shim.Error("Could not locate bioinfo")
	}
	return shim.Success(bioInfoAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	bioInfos := []BioInfo{
		BioInfo{Identifier: "BI001", Name: "Sequence A", Owner: "John Doe", Hash: "hash1", Status: "approved"},
		BioInfo{Identifier: "BI002", Name: "Sequence B", Owner: "Jane Doe", Hash: "hash2", Status: "pending"},
		BioInfo{Identifier: "BI003", Name: "Sequence C", Owner: "Bob Smith", Hash: "hash3", Status: "rejected"},
	}

	i := 0
	for i < len(bioInfos) {
		fmt.Println("i is ", i)
		bioInfoAsBytes, _ := json.Marshal(bioInfos[i])
		APIstub.PutState("BI"+strconv.Itoa(i+1), bioInfoAsBytes)
		fmt.Println("Added", bioInfos[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createBioInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var bioInfo = BioInfo{Identifier: args[0], Name: args[1], Owner: args[2], Hash: args[3], Status: args[4]}

	bioInfoAsBytes, _ := json.Marshal(bioInfo)
	APIstub.PutState(args[0], bioInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeBioInfoStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	bioInfoAsBytes, _ := APIstub.GetState(args[0])
	bioInfo := BioInfo{}

	json.Unmarshal(bioInfoAsBytes, &bioInfo)
	bioInfo.Status = args[1]

	bioInfoAsBytes, _ = json.Marshal(bioInfo)
	APIstub.PutState(args[0], bioInfoAsBytes)

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
