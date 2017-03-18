package View

import (
	"BJS_ChainDemo/Control"
	"BJS_ChainDemo/Module/system"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	Control.InitUserMemory(stub)
	Control.InitTimeoutMemory(stub)
	System.InitReqID(stub)
	return nil, nil
}
