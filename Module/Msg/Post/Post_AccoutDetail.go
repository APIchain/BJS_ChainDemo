package Post

import (
	"BJS_ChainDemo/Log"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const BJSServerURL = "http://127.0.0.1:5000/"

type AccountReply struct {
	TransType        string `json:"transType"`
	Username         string `json:"username"`
	PostDataServer   string `json:"postDataServer"`
	ReturnDataServer string `json:"returnDataServer"`
}

func (u *AccountReply) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("AccountReply ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}

func (u *AccountReply) Post() ([]byte, error) {
	url := BJSServerURL
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
