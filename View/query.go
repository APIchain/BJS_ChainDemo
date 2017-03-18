package View

import (
	"BJS_ChainDemo/Control"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	//指定某个用户
	FUNC_QUERY_USER_DETAIL_BY_USERNAME = "QueryUserDetailByUserName"
	//指定所有用户
	FUNC_QUERY_USER_DETAIL_ALL = "QueryUserDetailAll"
	//查询所有请求
	FUNC_QUERY_ALL_REQUEST = "QueryAllRequest"
	//查询所有返回
	FUNC_QUERY_ALL_RESPONSE = "QueryAllResponse"
	//查询所有请求ByublicKey
	FUNC_QUERY_REQUEST_BY_PUBLICKEY = "QueryRequestByPublicKey"
	//查询所有返回ByublicKey
	FUNC_QUERY_RESPONSE_BY_PUBLICKEY = "QueryResponseByPublicKey"

)

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_QUERY_USER_DETAIL_BY_USERNAME:
		return t.QueryUserDetailByUserName(stub, function, args)
	case FUNC_QUERY_USER_DETAIL_ALL:
		return t.QueryUserDetailAll(stub, function, args)
	case FUNC_QUERY_ALL_REQUEST:
		return t.QueryAllRequest(stub, function, args)
	case FUNC_QUERY_ALL_RESPONSE:
		return t.QueryAllResponse(stub, function, args)
	case FUNC_QUERY_REQUEST_BY_PUBLICKEY:
		return t.QueryRequestByPublicKey(stub, function, args)
	case FUNC_QUERY_RESPONSE_BY_PUBLICKEY:
		return t.QueryResponseByPublicKey(stub, function, args)
	}

	return nil, errors.New("Invalid Function Call:" + function)
}

func (t *SimpleChaincode) QueryUserDetailByUserName(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "QueryUserDetail" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var username string // Entities
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting username of the person to query")
	}
	username = args[0]
	user, err := Control.DefaultUserMemory.GetByUserName(username)
	if err != nil {
		return nil, errors.New("GetByUserName failed.")
	}
	return user.ToJson()
}

func (t *SimpleChaincode) QueryUserDetailAll(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil,nil
}
func (t *SimpleChaincode) QueryAllRequest(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil,nil
}
func (t *SimpleChaincode) QueryAllResponse(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil,nil
}
func (t *SimpleChaincode) QueryRequestByPublicKey(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil,nil
}
func (t *SimpleChaincode) QueryResponseByPublicKey(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil,nil
}