# BaiduPushSDK-golang
 百度云推送服务器端go语言版本SDK，封装了Rest api 3.0，包含目前开放的所有API，详见：http://push.baidu.com/doc/restapi/restapi

## 安装
 `go get github.com/ChangjunZhao/BaiduPushSDK-golang`

## 使用方法

### 推送消息到客户端
```
package main

import (
	"fmt"
	"github.com/ChangjunZhao/BaiduPushSDK-golang"
)

func main() {
	// 新建客户端
	client := push.NewClient("Your API Key","Your Secret Key")
	// 构造请求
	request := &push.PushMsgToSingleDeviceRequest{ChannelId: "your channel_id", MsgType: 0, Message: "测试消息"}

	// 推送消息到指定客户端
	response, err := client.PushMsgToSingleDevice(*request)
	if err == nil {
		fmt.Println(response.MsgId)
	} else {
		fmt.Println(err)
	}
}
```

### 推送系统通知到客户端
```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChangjunZhao/BaiduPushSDK-golang"
)

func main() {
	// 新建客户端
	client := push.NewClient("Your API Key","Your Secret Key")
	// 推送通知到指定客户端
	notification := &push.AndroidNotification{Title: "测试通知", Description: "测试通知描述", NotificationBuilderId: 0, NotificationBasicStyle: 7, OpenType: 1, Url: "http://www.tsuru.cn"}
	message, _ := json.Marshal(notification)
	request := &push.PushMsgToSingleDeviceRequest{ChannelId: "your channel_id", MsgType: 1, Message: string(message)}
	response, err := client.PushMsgToSingleDevice(*request)
	if err == nil {
		fmt.Println(response.MsgId)
	} else {
		fmt.Println(err)
	}
}

## 如使用中发现问题，请通过github提交，我们会尽快完善修复，或直接pull requests
