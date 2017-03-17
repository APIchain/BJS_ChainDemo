package Log

import "github.com/hyperledger/fabric/core/chaincode/shim"

var Logger *shim.ChaincodeLogger

func init() {
	Logger = shim.NewLogger("ChainDemo")
	Logger.SetLevel(shim.LogDebug)
	Logger.IsEnabledFor(shim.LogDebug)
}
