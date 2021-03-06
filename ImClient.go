package jim

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

const (
	HOST_NAME_SSL = "https://api.im.jpush.cn"
)

type ImClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewImClient(masterSecret, appKey string) *ImClient {
	//base64
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+masterSecret))
	imer := &ImClient{masterSecret, appKey, auth, HOST_NAME_SSL}
	return imer

}

func (this *ImClient) SendCustomMsg(msgbody []byte) string {
	return this.PostBody(msgbody, "/v1/messages")
}

func (this *ImClient) PostBody(msgbody []byte, patch string) string {
	req := httplib.Post(this.BaseUrl + patch)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("Authorization", this.AuthCode)
	req.Header("Content-Type", "application/json;charset=utf-8")
	req.Body(msgbody)
	res, err := req.String()
	if err == nil {
		return res
	} else {
		fmt.Println("error:", err)
		return "fail"
	}

}
