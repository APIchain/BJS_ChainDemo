package System

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const REQID_KEY_COUNTER = "ReqIDKeyCounter"

type ReqID int64

func GenReqID(stub shim.ChaincodeStubInterface) (ReqID, error) {
	key := REQID_KEY_COUNTER
	bytexs, err := stub.GetState(key)
	if err != nil {
		return ReqID(uint64(0)), err
	}
	ReqIDNow, err := strconv.ParseInt(string(bytexs), 10, 64)
	if err != nil {
		return ReqID(uint64(0)), err
	}
	ReqIDNow = ReqIDNow + 1
	ReqIDNowToWrite := strconv.FormatInt(ReqIDNow, 10)
	if err != nil {
		return ReqID(uint64(0)), err
	}
	buf := bytes.NewBufferString(ReqIDNowToWrite)
	err = stub.PutState(key, buf.Bytes())
	if err != nil {
		return ReqID(uint64(0)), err
	}
	temp := ReqID(ReqIDNow)
	return temp, nil
}

func (t *ReqID) ToString() string {
	return fmt.Sprint(*t)
}
