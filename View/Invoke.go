package View

import (
	. "BJS_ChainDemo/Control"
	. "BJS_ChainDemo/Log"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"GoOnchain/common/log"
)

const (
	//Invoke user
	FUNC_INVOKE_USER_DELETE = "InvokeUserDelete"
	FUNC_INVOKE_USER_REGIST = "InvokeUserRegist"
	FUNC_INVOKE_USER_UPDATE = "InvokeUserUpdate"
	//Invoke Bussiness
	FUNC_INVOKE_GETDATA = "InvokeGetData"
	FUNC_INVOKE_POSTDATA = "InvokePostData"
)

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_INVOKE_USER_DELETE:
		return t.InvokeUserDelete(stub, args)
	case FUNC_INVOKE_USER_REGIST:
		return t.InvokeUserRegist(stub, args)
	case FUNC_INVOKE_USER_UPDATE:
		return t.InvokeUserUpdate(stub, args)
	case FUNC_INVOKE_GETDATA:
		return t.InvokeGetData(stub, args)
	case FUNC_INVOKE_POSTDATA:
		return t.InvokePostData(stub, args)
	}
	return nil, errors.New("Invalid Function Call:" + function)
}

func (t *SimpleChaincode) InvokeUserDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var username string
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting username of the person to query")
	}
	username = args[0]
	err := DefaultUserMemory.DeleteFromUserMemory(stub, username)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}
	return nil, nil
}

func (t *SimpleChaincode) InvokeUserRegist(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var username string
	var postDataServer string
	var returnDataServer string
	if len(args) != 3 {
		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 3 and got %d", len(args)))
	}
	username = args[0]
	postDataServer = args[1]
	returnDataServer = args[2]

	err := DefaultUserMemory.AddToUserMemory(stub, username, postDataServer, returnDataServer)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}
	return nil, nil
}

func (t *SimpleChaincode) InvokeUserUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var username string
	var postDataServer string
	var returnDataServer string
	if len(args) != 3 {
		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 3 and got %d", len(args)))
	}
	username = args[0]
	postDataServer = args[1]
	returnDataServer = args[2]

	err := DefaultUserMemory.UpdateToUserMemory(stub, username, postDataServer, returnDataServer)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}
	return nil, nil
}

func (t *SimpleChaincode) InvokeGetData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	log.Info("InvokeGetData invoke started.")
	var CHandler = NewCertHandler()
	sigma, err := stub.GetCallerMetadata()
	if err != nil {
		log.Info("Failed getting metadata")
		return nil, errors.New("Failed getting metadata")
	}
	payload, err := stub.GetPayload()
	if err != nil {
		log.Info("Failed getting payload")
		return nil, errors.New("Failed getting payload")
	}
	binding, err := stub.GetBinding()
	if err != nil {
		log.Info("Failed getting binding")
		return nil, errors.New("Failed getting binding")
	}

	Logger.Info("passed sigma [% x]", sigma)
	Logger.Info("passed payload [% x]", payload)
	Logger.Info("passed binding [% x]", binding)

	isAuthorized, err := CHandler.IsAuthorized(stub, "client")
	if !isAuthorized {
		Logger.Info("system error %v", err)
		return nil, errors.New("user is not aurthorized to assign assets")
	}

	return nil,nil
}


//func (t *SimpleChaincode) InvokeGetData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
//	var hash string
//	var publicKey string
//	var busType string
//	if len(args) != 3 {
//		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 3 and got %d", len(args)))
//	}
//	hash = args[0]
//	publicKey = args[1]
//	busType = args[2]
//
//	Control.PostRequest(hash, publicKey, busType)
//
//	account, err := attr.GetValueFrom("account", owner)
//	if err != nil {
//		fmt.Printf("Error reading account [%v] \n", err)
//		return nil, fmt.Errorf("Failed fetching recipient account. Error was [%v]", err)
//	}
//	Control.DefaultUserMemory.AddRequest()
//	return nil, nil
//}

func (t *SimpleChaincode) InvokePostData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}