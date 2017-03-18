package View

import (
	"BJS_ChainDemo/Control"
	"BJS_ChainDemo/Module/Reply"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	//查找指定用户所有信息
	FUNC_QUERY_USER_DETAIL_BY_USERNAME = "QueryUserDetailByUserName"
	//指定所有用户
	FUNC_QUERY_BUSTYPE_TIMEOUT_BYTYPE = "QueryBusTypeTimeoutByType"
)

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_QUERY_USER_DETAIL_BY_USERNAME:
		return t.QueryUserDetailByUserName(stub, function, args)
	case FUNC_QUERY_BUSTYPE_TIMEOUT_BYTYPE:
		return t.QueryBusTypeTimeoutByType(stub, function, args)
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

func (t *SimpleChaincode) QueryBusTypeTimeoutByType(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "QueryBusTypeTimeoutByType" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var txtype string // Entities
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting busType.")
	}
	txtype = args[0]

	if exist := Control.DefaultTimeoutSetting.CheckExist(txtype); exist {
		return nil, errors.New("This txtype does not exist.")
	}

	user := Control.DefaultTimeoutSetting.GetTimeoutValByBusType(txtype)

	msgReturn := &Reply.Msg_QueryBusTypeTimeoutByType{
		TimeoutVal: *user,
	}

	return msgReturn.ToJson()
}
