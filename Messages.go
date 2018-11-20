package jim

import (
	"encoding/json"
	//"fmt"
)

type Custom_MsgType struct {
	Version         int64       `json:"version"`
	Target_type     string      `json:"target_type"`
	Target_id       string      `json:"target_id"`
	From_type       string      `json:"from_type"`
	From_id         string      `json:"from_id"`
	Msg_type        string      `json:"msg_type"`
	Create_time     int64       `json:"create_time"`
	From_platform   string      `json:"from_platform"`
	No_offline      bool        `json:"no_offline"`
	No_notification bool        `json:"no_notification"`
	Msg_body        interface{} `json:"msg_body"`
}

func NewCustom_Msg() *Custom_MsgType {
	msg := &Custom_MsgType{}
	return msg
}

func (this *Custom_MsgType) SetVersion(version int64) {
	this.Version = version

}

func (this *Custom_MsgType) SetTarget_type(target_type string) {
	this.Target_type = target_type

}
func (this *Custom_MsgType) SetTarget_id(target_id string) {
	this.Target_id = target_id

}
func (this *Custom_MsgType) SetFrom_type(from_type string) {
	this.From_type = from_type

}

func (this *Custom_MsgType) SetFrom_id(from_id string) {
	this.From_id = from_id

}

func (this *Custom_MsgType) SetMsg_type(msg_type string) {
	this.Msg_type = msg_type

}
func (this *Custom_MsgType) SetNo_offline(c bool) {
	this.No_offline = c

}
func (this *Custom_MsgType) SetNo_notification(c bool) {
	this.No_notification = c

}
func (this *Custom_MsgType) SetCustomJsonBody(customJson interface{}) {
	this.Msg_body = customJson

}
func (this *Custom_MsgType) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	//fmt.Println("ToBytes:",string(content))
	return content, nil
}
