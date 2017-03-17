package View

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"BJS_ChainDemo/Control"
)

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	Control.InitUserMemory(stub)
	return nil, nil
}
