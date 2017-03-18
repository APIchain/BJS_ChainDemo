package System

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"sync"
)

const REQID_KEY_COUNTER = "ReqIDKeyCounter"

var DefaultReqID ReqID

type ReqID struct{
	Mu sync.Mutex
	ReqID int64  `json:"ReqID"`

}

func InitReqID(stub shim.ChaincodeStubInterface) error {
	DefaultReqID = ReqID{}
	key := REQID_KEY_COUNTER
	bytexs, err := stub.GetState(key)
	if err != nil {
		return err
	}
	ReqIDNow, err := strconv.ParseInt(string(bytexs), 10, 64)
	if err != nil {
		return  err
	}
	DefaultReqID.ReqID = ReqIDNow
	return nil
}


func(rq *ReqID) GenReqID(stub shim.ChaincodeStubInterface) (string, error) {

	rq.ReqID = rq.ReqID + 1
	err:=rq.Put(stub)
	if err != nil {
		return nil, err
	}
	strx :=strconv.FormatInt(rq.ReqID, 10)
	for len(strx) < 50{
		strx = "0" +strx
	}
	rq.Put(stub)
	return strx,nil
}

func(rq *ReqID) Put(stub shim.ChaincodeStubInterface) error {
	ReqIDNowToWrite := strconv.FormatInt(rq.ReqID, 10)
	buf := bytes.NewBufferString(ReqIDNowToWrite)
	err := stub.PutState(REQID_KEY_COUNTER, buf.Bytes())
	if err != nil {
		return  err
	}
	return nil
}


func (t *ReqID) ToString() string {
	return fmt.Sprint(*t)
}
