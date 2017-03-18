package View

import (
	"BJS_ChainDemo/Control"
	. "BJS_ChainDemo/Log"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"BJS_ChainDemo/Module/Msg/Post"
	"BJS_ChainDemo/Module/system"
	"BJS_ChainDemo/Module/Role"
)

const (
	//Invoke user
	FUNC_INVOKE_USER_DELETE = "InvokeUserDelete"
	FUNC_INVOKE_USER_REGIST = "InvokeUserRegist"
	FUNC_INVOKE_USER_UPDATE = "InvokeUserUpdate"
	//Invoke Bussiness
	FUNC_INVOKE_REQUEST  = "InvokeRequest"
	FUNC_INVOKE_RESPONSE = "InvokeResponse"
	//Set Timeout
	FUNC_INVOKE_SETTIMEOUT = "InvokeSetTimeout"
	//TestFunc
	TestFunc = "TestFunc"
)

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_INVOKE_USER_DELETE:
		return t.InvokeUserDelete(stub, args)
	case FUNC_INVOKE_USER_REGIST:
		return t.InvokeUserRegist(stub, args)
	case FUNC_INVOKE_USER_UPDATE:
		return t.InvokeUserUpdate(stub, args)
	case FUNC_INVOKE_REQUEST:
		return t.InvokeRequest(stub, args)
	case FUNC_INVOKE_RESPONSE:
		return t.InvokeResponse(stub, args)
	case FUNC_INVOKE_SETTIMEOUT:
		return t.InvokeSetTimeout(stub, args)
	case TestFunc:
		return t.TestFunc(stub, args)
	}

	return nil, errors.New("Invalid Function Call:" + function)
}

func (t *SimpleChaincode) InvokeUserDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var username string
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting username of the person to query")
	}
	username = args[0]
	err := Control.DefaultUserMemory.DeleteFromUserMemory(stub, username)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}
	userx,err :=Control.DefaultUserMemory.GetByUserName(username)
	if err != nil {
		return nil,err
	}

	post :=Post.AccountReply{
		TransType: "1",
		Username:userx.UserName,
		PostDataServer:userx.PostDataServer,
		ReturnDataServer:userx.ReturnDataServer,
	}
	post.Post()
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

	err := Control.DefaultUserMemory.AddToUserMemory(stub, username, postDataServer, returnDataServer)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}

	userx,err :=Control.DefaultUserMemory.GetByUserName(username)
	if err != nil {
		return nil,err
	}

	post :=Post.AccountReply{
		TransType: "2",
		Username:userx.UserName,
		PostDataServer:userx.PostDataServer,
		ReturnDataServer:userx.ReturnDataServer,
	}
	post.Post()
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

	err := Control.DefaultUserMemory.UpdateToUserMemory(stub, username, postDataServer, returnDataServer)
	if err != nil {
		Logger.Error("InvokeUserDelete Failed with username =", username, "Detail error is:", err)
	}
	userx,err :=Control.DefaultUserMemory.GetByUserName(username)
	if err != nil {
		return nil,err
	}

	post :=Post.AccountReply{
		TransType: "3",
		Username:userx.UserName,
		PostDataServer:userx.PostDataServer,
		ReturnDataServer:userx.ReturnDataServer,
	}
	post.Post()
	return nil, nil
}

func (t *SimpleChaincode) TestFunc(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//TODO: not ready
	Logger.Info("InvokeRequest invoke started.")
	var logs string
	var CHandler = Control.NewCertHandler()
	at, err := stub.ReadCertAttribute("role")
	if err != nil {
		logs = logs + fmt.Sprintf("ReadCertAttribute Failed getting metadata.%s.\n", err)
	}
	logs = logs + fmt.Sprintf("ReadCertAttribute at is%s.\n", at)
	as, err := stub.GetCallerCertificate()
	if err != nil {
		logs = logs + fmt.Sprintf("GetCallerCertificate Failed getting metadata.%s.\n", err)
	}
	logs = logs + fmt.Sprintf("GetCallerCertificate at is%s.\n", as)

	sigma, err := stub.GetCallerMetadata()
	logs = logs + fmt.Sprintf("Failed getting metadata %s.\n", sigma)
	payload, err := stub.GetPayload()
	logs = logs + fmt.Sprintf("Failed getting payload %s.\n", payload)
	binding, err := stub.GetBinding()
	logs = logs + fmt.Sprintf("Failed getting binding %s.\n", binding)

	isAuthorized1, err := CHandler.IsAuthorized(stub, "client")
	if isAuthorized1 {
		logs = logs + "client runed."
		fmt.Printf("client runed.")
		Logger.Info("system error %v", err)
		//return nil, errors.New("user is not aurthorized to assign assets")
	}

	isAuthorized2, err := CHandler.IsAuthorized(stub, "assigner")
	if isAuthorized2 {
		logs = logs + "assigner runed."
		fmt.Printf("assigner runed.")
		Logger.Info("system error %v", err)
		//return nil, errors.New("user is not aurthorized to assign assets")
	}

	return nil, errors.New(logs)
}

//func (t *SimpleChaincode) InvokeRequest(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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

func (t *SimpleChaincode) InvokeRequest(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var hash string
	var publicKey string
	var busType string
	var timestamp string
	var username string
	//TODO: username get func set

	if len(args) != 4 {
		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 6 and got %d", len(args)))
	}
	hash = args[0]
	publicKey = args[1]
	busType = args[2]
	timestamp = args[3]

	username = t.GetUserName(stub)

	reqseqno ,err:= System.DefaultReqID.GenReqID(stub)
	if err != nil {
		return nil,err
	}
	Req,err:= Post.NewRequest(stub,hash,publicKey,busType,reqseqno)
	if err != nil {
		return nil,err
	}
	Req.Post()

	rt := &Role.ReqTime{
		ReqID:       reqseqno,
		UserName:    username,
		BusType:      busType,
		RequestTime: timestamp,
	}
	rt.Put(stub)
	err= Control.DefaultUserMemory.AddRequest(stub,username)
	if err != nil {
		return nil,err
	}

	return nil, nil
}

func (t *SimpleChaincode) GetUserName(stub shim.ChaincodeStubInterface)string{
	return ""
}

func (t *SimpleChaincode) InvokeResponse(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var publicKey string
	var respCode string
	var respData string
	var reqSeq string
	var busType string
	var timestamp string

	if len(args) != 6 {
		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 6 and got %d", len(args)))
	}
	publicKey = args[0]
	respCode = args[1]
	respData = args[2]
	reqSeq = args[3]
	busType = args[4]
	timestamp = args[5]


	Req,err:= Post.NewResponse(stub,publicKey,respCode,respData,reqSeq,busType,timestamp)
	if err != nil {
		return nil,err
	}
	Req.Post()
	Control.SetResTimesOrTimeoutTimes(stub,reqSeq,timestamp)

	return nil, nil
}

func (t *SimpleChaincode) InvokeSetTimeout(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var busType string
	var timeoutVal string
	if len(args) != 2 {
		return nil, errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 2 and got %d", len(args)))
	}
	busType = args[0]
	timeoutVal = args[1]

	err := Control.DefaultTimeoutSetting.AddOrUpdatebusType(stub, busType, timeoutVal)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("AddOrUpdatebusType failed with %s.\n",busType))
	}
	return nil, nil
}
