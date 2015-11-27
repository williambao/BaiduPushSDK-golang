// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 百度云推送 SDK go语言版本
package push

import (
	"fmt"
	"time"
)

//调用API的Client
type BaiduPushClient struct {
	apikey    string
	secretKey string
	debug     bool
}

//创建新的客户端
//
//使用方法：
//
//c = NewClient("API Key","Secret Key")
func NewClient(apikey string, secretKey string) *BaiduPushClient {
	return &BaiduPushClient{
		apikey:    apikey,
		secretKey: secretKey,
	}
}

func (c *BaiduPushClient) Debug(enabled bool) {
	c.debug = enabled
}

// Android通知消息，详见：http://push.baidu.com/doc/restapi/msg_struct
type AndroidNotification struct {
	Title                  string `json:"title"`                    //通知标题，可以为空；如果为空则设为appid对应的应用名;
	Description            string `json:"description"`              //通知文本内容，不能为空;
	NotificationBuilderId  int    `json:"notification_builder_id"`  //android客户端自定义通知样式，如果没有设置默认为0;
	NotificationBasicStyle int    `json:"notification_basic_style"` //只有notification_builder_id为0时有效，可以设置通知的基本样式包括(响铃：0x04;振动：0x02;可清除：0x01;),这是一个flag整形，每一位代表一种样式;
	OpenType               int    `json:"open_type"`                //点击通知后的行为(1：打开Url; 2：自定义行为；3：默认打开应用;)
	Url                    string `json:"url"`                      //需要打开的Url地址，open_type为1时才有效;
	//open_type为2时才有效，Android端SDK会把pkg_content字符串转换成Android Intent,
	//通过该Intent打开对应app组件，所以pkg_content字符串格式必须遵循Intent uri格式，最简单的方法可以通过Intent方法toURI()获取
	PkgContent    string                 `json:"pkg_content"`
	CustomContent map[string]interface{} `json:"custom_content"` //自定义内容，键值对，Json对象形式(可选)；在android客户端，这些键值对将以Intent中的extra进行传递。
}

// 推送消息到单台设备
func (c *BaiduPushClient) PushMsgToSingleDevice(request PushMsgToSingleDeviceRequest) (*PushMsgToSingleDeviceResponse, error) {
	var resp PushMsgToSingleDeviceJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "push", "/single_device", params, c.secretKey, &resp)
	if err == nil {
		return &resp.PushMsgToSingleDeviceResponse, nil
	} else {
		return nil, err
	}

}

// 推送广播消息
func (c *BaiduPushClient) PushMsgToAllDevice(request PushMsgToAllRequest) (*PushMsgToAllResponse, error) {
	var resp PushMsgToAllJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "push", "/all", params, c.secretKey, &resp)
	if err == nil {
		return &resp.PushMsgToAllResponse, nil
	} else {
		return nil, err
	}

}

// 推送组播消息
func (c *BaiduPushClient) PushMsgToTag(request PushMsgToTagRequest) (*PushMsgToTagResponse, error) {
	var resp PushMsgToTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "push", "/tags", params, c.secretKey, &resp)
	if err == nil {
		return &resp.PushMsgToTagResponse, nil
	} else {
		return nil, err
	}

}

// 推送消息到给定的一组设备(批量单播)
func (c *BaiduPushClient) PushBatchUniMsg(request PushBatchUniMsgRequest) (*PushBatchUniMsgResponse, error) {
	var resp PushBatchUniMsgJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "push", "/batch_device", params, c.secretKey, &resp)
	if err == nil {
		return &resp.PushBatchUniMsgResponse, nil
	} else {
		return nil, err
	}

}

// 查询消息的发送状态
func (c *BaiduPushClient) QueryMsgStatus(request QueryMsgStatusRequest) (*QueryMsgStatusResponse, error) {
	var resp QueryMsgStatusJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "report", "/query_msg_status", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryMsgStatusResponse, nil
	} else {
		return nil, err
	}

}

// 查询定时消息的发送记录
func (c *BaiduPushClient) QueryTimerRecords(request QueryTimerRecordsRequest) (*QueryTimerRecordsResponse, error) {
	var resp QueryTimerRecordsJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "report", "/query_timer_records", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryTimerRecordsResponse, nil
	} else {
		return nil, err
	}

}

// 查询指定分类主题的发送记录
func (c *BaiduPushClient) QueryTopicRecords(request QueryTopicRecordsRequest) (*QueryTopicRecordsResponse, error) {
	var resp QueryTopicRecordsJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "report", "/query_topic_records", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryTopicRecordsResponse, nil
	} else {
		return nil, err
	}

}

// 查询标签组列表
func (c *BaiduPushClient) QueryTags(request QueryTagsRequest) (*QueryTagsResponse, error) {
	var resp QueryTagsJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "app", "/query_tags", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryTagsResponse, nil
	} else {
		return nil, err
	}

}

// 创建标签组
func (c *BaiduPushClient) CreateTag(request CreateTagRequest) (*CreateTagResponse, error) {
	var resp CreateTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "app", "/create_tag", params, c.secretKey, &resp)
	if err == nil {
		return &resp.CreateTagResponse, nil
	} else {
		return nil, err
	}

}

// 删除标签组
func (c *BaiduPushClient) DeleteTag(request DeleteTagRequest) (*DeleteTagResponse, error) {
	var resp DeleteTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "app", "/del_tag", params, c.secretKey, &resp)
	if err == nil {
		return &resp.DeleteTagResponse, nil
	} else {
		return nil, err
	}

}

// 添加设备到标签组
func (c *BaiduPushClient) AddDevicesToTag(request AddDevicesToTagRequest) (*AddDevicesToTagResponse, error) {
	var resp AddDevicesToTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "tag", "/add_devices", params, c.secretKey, &resp)
	if err == nil {
		return &resp.AddDevicesToTagResponse, nil
	} else {
		return nil, err
	}

}

// 将设备从标签组中移除
func (c *BaiduPushClient) DeleteDevicesFromTag(request DeleteDevicesFromTagRequest) (*DeleteDevicesFromTagResponse, error) {
	var resp DeleteDevicesFromTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "tag", "/del_devices", params, c.secretKey, &resp)
	if err == nil {
		return &resp.DeleteDevicesFromTagResponse, nil
	} else {
		return nil, err
	}

}

// 查询标签组设备数量
func (c *BaiduPushClient) QueryDeviceNumInTag(request QueryDeviceNumInTagRequest) (*QueryDeviceNumInTagResponse, error) {
	var resp QueryDeviceNumInTagJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "tag", "/device_num", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryDeviceNumInTagResponse, nil
	} else {
		return nil, err
	}

}

// 查询定时任务列表
func (c *BaiduPushClient) QueryTimerList(request QueryTimerListRequest) (*QueryTimerListResponse, error) {
	var resp QueryTimerListJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "timer", "/query_list", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryTimerListResponse, nil
	} else {
		return nil, err
	}

}

// 取消定时任务
func (c *BaiduPushClient) CancelTimerJob(timerId string) (*CancelTimerJobResponse, error) {
	var resp CancelTimerJobResponse
	params := c.baseParams(nil)
	params.AddUnescaped("timer_id", timerId)
	err := CallApiServer("POST", API_SERVER, "timer", "/cancel", params, c.secretKey, &resp)
	if err == nil {
		return &resp, nil
	} else {
		return nil, err
	}
}

// 查询分类主题列表
func (c *BaiduPushClient) QueryTopicList(request QueryTopicListRequest) (*QueryTopicListResponse, error) {
	var resp QueryTopicListJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "topic", "/query_list", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryTopicListResponse, nil
	} else {
		return nil, err
	}

}

// 当前应用的设备统计信息
func (c *BaiduPushClient) QueryStatisticDevice() (*QueryStatisticDeviceResponse, error) {
	var resp QueryStatisticDeviceJSONResponse
	params := c.baseParams(nil)
	err := CallApiServer("POST", API_SERVER, "report", "/statistic_device", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryStatisticDeviceResponse, nil
	} else {
		return nil, err
	}

}

// 查询分类主题统计信息
func (c *BaiduPushClient) QueryStatisticTopic(request QueryStatisticTopicRequest) (*QueryStatisticTopicResponse, error) {
	var resp QueryStatisticTopicJSONResponse
	params := c.baseParams(nil)
	params = request.AddToParams(params)
	err := CallApiServer("POST", API_SERVER, "report", "/statistic_topic", params, c.secretKey, &resp)
	if err == nil {
		return &resp.QueryStatisticTopicResponse, nil
	} else {
		return nil, err
	}

}

// 构造公共参数
func (c *BaiduPushClient) baseParams(additionalParams map[string]string) *OrderedParams {
	params := NewOrderedParams()
	params.Add("apikey", c.apikey)
	params.Add("timestamp", fmt.Sprintf("%d", (time.Now().Unix())))
	for key, value := range additionalParams {
		params.Add(key, value)
	}
	return params
}
