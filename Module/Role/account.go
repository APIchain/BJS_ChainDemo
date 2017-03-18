package Role

import (
	"BJS_ChainDemo/Log"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const DEFAULT_USER_LIST = "User"

type Account struct {
	UserName         string `json:"Account"`
	PostDataServer   string `json:"PostDataServer"`
	ReturnDataServer string `json:"ReturnDataServer"`
	RequestTime      int64  `json:"RequestTime"`
	ResponseTime     int64  `json:"ResponseTime"`
	TimeoutTime      int64  `json:"TimeoutTime"`
	Status           bool
}

func NewAccount(stub shim.ChaincodeStubInterface, userName string, postDataServer string, returnDataServer string) (*Account, error) {
	acc := &Account{
		UserName:         userName,
		PostDataServer:   postDataServer,
		ReturnDataServer: returnDataServer,
	}

	err := acc.Put(stub)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (u *Account) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(DEFAULT_USER_LIST+u.UserName, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

func (u *Account) Del(stub shim.ChaincodeStubInterface) error {
	err := stub.DelState(DEFAULT_USER_LIST + u.UserName)
	if err != nil {
		return err
	}
	return nil
}

func (u *Account) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("User ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}
