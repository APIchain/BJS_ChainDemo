package Reply

import (
	"BJS_ChainDemo/Log"
	"encoding/json"
)

type Msg_QueryBusTypeTimeoutByType struct {
	TimeoutVal string `json:"timeoutVal"`
}

func (u *Msg_QueryBusTypeTimeoutByType) ToJson() ([]byte, error) {
	jsonRespByte, err := json.Marshal(u)
	if err != nil {
		Log.Logger.Error("Msg_QueryBusTypeTimeoutByType ToJson() failed.", err)
		return nil, err
	}
	return jsonRespByte, nil
}
