package Role

import (
	"time"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"BJS_ChainDemo/Control"
)

const default_REQTIME ="ReqTime"

type ReqTime struct {
	ReqID       string `json:"ReqID"`
	UserName     string `json:"UserName"`
	BusType      string `json:"BusType"`
	RequestTime string `json:"RequestTime"`
}

func SetResTimesOrTimeoutTimes(stub shim.ChaincodeStubInterface,reqID string,ResponseTime string)error{
	Req,err := ReqTimeGetbyReqID(stub,reqID)
	if err != nil {
		return err
	}
	timeoutTime,err:=Control.DefaultTimeoutSetting.GetTimeoutValByBusType(Req.BusType)
	if err != nil {
		return err
	}
	requested, _ := time.Parse("2006/1/2 15/4/5", Req.RequestTime)
	responsed, _ := time.Parse("2006/1/2 15/4/5", ResponseTime)
	after, _ := time.ParseDuration(string(timeoutTime)+"s")
	requested = requested.Add(after)
	if requested.Before(responsed){
		Control.DefaultUserMemory.AddTimeOut(stub,Req.UserName)
	}else{
		Control.DefaultUserMemory.AddResponse(stub,Req.UserName)
	}

	return nil

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