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

const default_RESPONSE_LIST = "Response"
const ResponseURL = "http://127.0.0.1:5000/"

type Response struct {
	PublicKey string `json:"publicKey"`
	RespCode  string `json:"respCode"`
	RespData  string `json:"respData"`
	ReqSeq    string `json:"reqSeq"`
	BusType   string `json:"busType"`
	Timestamp string `json:"timestamp"`
}

func NewResponse(stub shim.ChaincodeStubInterface, str1 string, str2 string, str3 string, str4 string, str5 string, str6 string) (*Response, error) {
	acc := &Response{
		PublicKey: str1,
		RespCode:  str2,
		RespData:  str3,
		ReqSeq:    str4,
		BusType:   str5,
		Timestamp: str6,
	}

	err := acc.Put(stub)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (u *Response) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = stub.PutState(default_RESPONSE_LIST+u.PublicKey, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

func (u *Response) Del(stub shim.ChaincodeStubInterface) error {
	err := stub.DelState(default_RESPONSE_LIST + u.PublicKey)
	if err != nil {
		return err
	}
	return nil
}

func (u *Response) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("Response ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}

func (u *Response) Post() ([]byte, error) {
	url := ResponseURL
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
