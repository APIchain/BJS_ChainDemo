package main

import (
	"BJS_ChainDemo/View"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(View.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
