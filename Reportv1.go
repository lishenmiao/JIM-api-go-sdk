package jim

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

const (
	Report_HOST_NAME_SSL = "https://report.im.jpush.cn"
)

type ReportImClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewImReport(masterSecret, appKey string) *ReportImClient {
	//base64
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+masterSecret))
	imer := &ReportImClient{masterSecret, appKey, auth, HOST_NAME_SSL}
	return imer

}

func (this *ReportImClient) GetUserMsg(username,start ,count string) string {
	return this.Get( "/v1/"+username+"/messages"+"?start="+start+"&count="+count)
}

func (this *ReportImClient) Get( patch string) string {
	req := httplib.Get(this.BaseUrl + patch)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.Header("Authorization", this.AuthCode)
	req.Header("Content-Type", "application/json;charset=utf-8")
	//req.Body(msgbody)
	res, err := req.String()
	if err == nil {
		return res
	} else {
		fmt.Println("error:", err)
		return "fail"
	}

}

//func (this *ReportImClient) PostBody(msgbody []byte, patch string) string {
//	req := httplib.Post(this.BaseUrl + patch)
//	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
//	req.Header("Authorization", this.AuthCode)
//	req.Header("Content-Type", "application/json;charset=utf-8")
//	req.Body(msgbody)
//	res, err := req.String()
//	if err == nil {
//		return res
//	} else {
//		fmt.Println("error:", err)
//		return "fail"
//	}
//
//}
