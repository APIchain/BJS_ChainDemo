package Reply

import (
	"BJS_ChainDemo/Log"
	"encoding/json"
)

type Reply_ResponseNotifier struct {
	TimeoutVal string `json:"timeoutVal"`
}

func (u *Reply_ResponseNotifier) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("Reply_ResponseNotifier ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}
