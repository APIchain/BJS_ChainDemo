package Control

import (
	"BJS_ChainDemo/Module/Role"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const default_Timeout_LIST = "Timeout"

var DefaultTimeoutSetting TimeoutSetting

type TimeoutSetting struct {
	List []Role.Timeout `json:"TimeList"`
}

func InitTimeoutMemory(stub shim.ChaincodeStubInterface) error {
	ubs := new(TimeoutSetting)
	key := default_Timeout_LIST
	bytes, err := stub.GetState(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, ubs)
	if err != nil {
		return err
	}
	DefaultTimeoutSetting = *ubs
	return nil
}

func (u *TimeoutSetting) put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(default_Timeout_LIST, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

func (u *TimeoutSetting) GetTimeoutValByBusType(txtype string) *string {
	for _, v := range DefaultTimeoutSetting.List {
		if txtype == v.BusType {
			return &v.TimeoutVal
		}
	}
	return nil
}

func (u *TimeoutSetting) CheckExist(busType string) bool {
	for _, v := range DefaultTimeoutSetting.List {
		if v.BusType == busType {
			return true
		}

	}
	return false
}

func (u *TimeoutSetting) AddOrUpdatebusType(stub shim.ChaincodeStubInterface, busType string, timeoutVal string) error {
	if u.CheckExist(busType) {
		for _, v := range u.List {
			if v.BusType == busType {
				v.TimeoutVal = timeoutVal
			}
		}

	} else {
		timeout := Role.Timeout{
			BusType:    busType,
			TimeoutVal: timeoutVal,
		}
		u.List = append(u.List, timeout)
	}
	u.put(stub)

	return nil
}
