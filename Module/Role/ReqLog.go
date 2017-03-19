package Role

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const default_ReqLog ="ReqLog"

type ReqLog struct {
	ReqID       string `json:"ReqID"`
	UserName     string `json:"UserName"`
	BusType      string `json:"BusType"`
	RequestTime string `json:"RequestTime"`
}

func  ReqLogGetbyReqID(stub shim.ChaincodeStubInterface, reqID string)(*ReqLog,error){
	ubs := new(ReqLog)
	key := default_ReqLog + reqID
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

func (u *ReqLog) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(default_ReqLog+u.ReqID, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}