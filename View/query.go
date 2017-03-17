package View

import (
	"BJS_ChainDemo/Control"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	//User
	FUNC_QUERY_USER_DETAIL = "QueryUserDetail"
)

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case FUNC_QUERY_USER_DETAIL:
		return t.QueryUserDetail(stub, function, args)
	}
	return nil, errors.New("Invalid Function Call:" + function)
}

func (t *SimpleChaincode) QueryUserDetail(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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
