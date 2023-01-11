package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	name        = "BioInformaticSecuritySuite"
	version     = "1.0.0"
	description = "BioInformatic Security Suite Chaincode"
)

func InitMetadata(stub shim.ChaincodeStubInterface) error {
	return stub.SetChaincodeMetadata(name, []byte(version), []byte(description))
}
