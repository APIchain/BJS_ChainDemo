package Role

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const default_REQTIME ="ReqTime"

type ReqTime struct {
	ReqID       string `json:"ReqID"`
	UserName     string `json:"UserName"`
	BusType      string `json:"BusType"`
	RequestTime string `json:"RequestTime"`
}

func  ReqTimeGetbyReqID(stub shim.ChaincodeStubInterface, reqID string)(*ReqTime,error){
	ubs := new(ReqTime)
	key := default_REQTIME + reqID
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil,err
	}
	err = json.Unmarshal(bytes, ubs)
	if err != nil {
		return nil, err
	}
	return ubs,nil
}

func (u *ReqTime) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(default_REQTIME+u.ReqID, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}