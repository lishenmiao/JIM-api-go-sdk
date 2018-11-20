# JIM-api-go-sdk
jiguang api go sdk


_example   IM Send  Custom msg 发送自定义消息_

```
package main

import (
	"github.com/lishenmiao/JIM-api-go-sdk"
	"time"
	"fmt"
)

const (
	JIMappKey = "jiguangIM appkey"
	JIMsecret = "jiguangIM mastersecret"
)
//Custom msg_Jsonbody
type Transfer_CustomJson struct {
	Type string
	NumberText string
	HeadUrl string
	UserName string
	TransferId  string
	Createtime int64
}

func main()  {

newim:=jim.NewImClient(JIMsecret,JIMappKey)

msg:=jim.NewCustom_Msg()
msg.SetVersion(1)
msg.SetFrom_id("KWallet")
msg.SetFrom_type("admin")
msg.SetTarget_id("210012")
msg.SetTarget_type("single")
msg.SetMsg_type("custom")
msg.Create_time=time.Now().Unix()

var myjson Transfer_CustomJson
myjson.Type="Transfer"
myjson.Createtime=time.Now().Unix()
myjson.NumberText="0.001 USDT"
myjson.UserName="132106"
myjson.TransferId="TX-38388kdkdk18383"
myjson.HeadUrl="http://d.kingoapp.com/coinprice/64/825.png"

msg.SetCustomJsonBody(myjson)

msgbyte,err:=msg.ToBytes()


if err ==nil{
	res:=newim.SendCustomMsg(msgbyte)
	fmt.Println("请求结果:",res)
}else {
	fmt.Println("Json转换错误:",err)
}

}

```
