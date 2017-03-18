package Post

import (
	"BJS_ChainDemo/Log"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"io/ioutil"
	"net/http"
	"strings"
)

const REPLYURL = "http://127.0.0.1:5000/"

const default_REQUEST_LIST = "Request"

type Request struct {
	Hash      string `json:"Hash"`
	PublicKey string `json:"PublicKey"`
	BusType   string `json:"BusType"`
}

func NewRequest(stub shim.ChaincodeStubInterface, str1 string, str2 string, str3 string) (*Request, error) {
	acc := &Request{
		Hash:      str1,
		PublicKey: str2,
		BusType:   str3,
	}

	err := acc.Put(stub)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (u *Request) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(default_REQUEST_LIST+u.PublicKey, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

func (u *Request) Del(stub shim.ChaincodeStubInterface) error {
	err := stub.DelState(default_REQUEST_LIST + u.PublicKey)
	if err != nil {
		return err
	}
	return nil
}

func (u *Request) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("Request ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}

func (u *Request) Post() ([]byte, error) {
	url := REPLYURL
	jsonx, err := u.ToJson()
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(hex.EncodeToString(jsonx))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return nil, nil
}
