package Control

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
	"BJS_ChainDemo/Module/Role"
)

func SetResTimesOrTimeoutTimes(stub shim.ChaincodeStubInterface,reqID string,ResponseTime string)error{
	Req,err := Role.ReqTimeGetbyReqID(stub,reqID)
	if err != nil {
		return err
	}
	timeoutTime,err:=DefaultTimeoutSetting.GetTimeoutValByBusType(Req.BusType)
	if err != nil {
		return err
	}
	requested, _ := time.Parse("2006/1/2 15/4/5", Req.RequestTime)
	responsed, _ := time.Parse("2006/1/2 15/4/5", ResponseTime)
	after, _ := time.ParseDuration(string(timeoutTime)+"s")
	requested = requested.Add(after)
	if requested.Before(responsed){
		DefaultUserMemory.AddTimeOut(stub,Req.UserName)
	}else{
		DefaultUserMemory.AddResponse(stub,Req.UserName)
	}

	return nil

}