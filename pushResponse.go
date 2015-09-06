// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//Rest Api返回值封装
package push

type MsgSendInfo struct {
	MsgId        string `json:"msg_id"`
	MsgStatus    int    `json:"status"`
	SendTime     int64  `json:"send_time"`
	SuccessCount int    `json:"success"`
}

type QueryMsgStatusResponse struct {
	TotalNum     int           `json:"total_num"`
	MsgSendInfos []MsgSendInfo `json:"result"`
}

type QueryMsgStatusJSONResponse struct {
	RequestId              int                    `json:"request_id"`
	QueryMsgStatusResponse QueryMsgStatusResponse `json:"response_params"`
}

type PushMsgToSingleDeviceResponse struct {
	MsgId    string `json:"msg_id"`
	SendTime int64  `json:"send_time"`
}

type PushMsgToSingleDeviceJSONResponse struct {
	RequestId                     int                           `json:"request_id"`
	PushMsgToSingleDeviceResponse PushMsgToSingleDeviceResponse `json:"response_params"`
}

type PushMsgToAllResponse struct {
	MsgId    string
	TimerId  string
	SendTime int64
}

type PushMsgToAllJSONResponse struct {
	RequestId            int                  `json:"request_id"`
	PushMsgToAllResponse PushMsgToAllResponse `json:"response_params"`
}

type PushMsgToTagResponse struct {
	MsgId    string `json:"msg_id"`
	TimerId  string `json:"timer_id"`
	SendTime int64  `json:"send_time"`
}

type PushMsgToTagJSONResponse struct {
	RequestId            int                  `json:"request_id"`
	PushMsgToTagResponse PushMsgToTagResponse `json:"response_params"`
}

type PushBatchUniMsgResponse struct {
	MsgId    string `json:"msg_id"`
	SendTime int64  `json:"send_time"`
}

type PushBatchUniMsgJSONResponse struct {
	RequestId               int                     `json:"request_id"`
	PushBatchUniMsgResponse PushBatchUniMsgResponse `json:"response_params"`
}

type Record struct {
	MsgId    string `json:"msg_id"`
	Status   int    `json:"status"`
	SendTime int64  `json:"send_time"`
}

type QueryTimerRecordsResponse struct {
	TimerId      string   `json:"timer_id"`
	TimerRecords []Record `json:"result"`
}

type QueryTimerRecordsJSONResponse struct {
	RequestId                 int                       `json:"request_id"`
	QueryTimerRecordsResponse QueryTimerRecordsResponse `json:"response_params"`
}

type QueryTopicRecordsResponse struct {
	TopicId      string   `json:"topic_id"`
	TopicRecords []Record `json:"result"`
}

type QueryTopicRecordsJSONResponse struct {
	RequestId                 int                       `json:"request_id"`
	QueryTopicRecordsResponse QueryTopicRecordsResponse `json:"response_params"`
}

type TimerResultInfo struct {
	timerId   string `json:"timer_id"`
	sendTime  int64  `json:"send_time"`
	message   string `json:"msg"`
	msgTpye   int    `json:"msg_type"`
	rangeType int    `json:"range_type"`
}

type QueryTimerListResponse struct {
	TotalNum         int               `json:"total_num"`
	TimerResultInfos []TimerResultInfo `json:"result"`
}

type QueryTimerListJSONResponse struct {
	RequestId              int                    `json:"request_id"`
	QueryTimerListResponse QueryTimerListResponse `json:"response_params"`
}

type TopicResultInfo struct {
	TopicId          string `json:"topic_id"`
	FirstPushTime    int64  `json:"ctime"`
	LastPushTime     int64  `json:"mtime"`
	TotalPushDevsNum int    `json:"push_cnt"`
	TotalAckDevsNum  int    `json:"ack_cnt"`
}

type QueryTopicListResponse struct {
	TotalNum         int               `json:"total_num"`
	TopicResultInfos []TopicResultInfo `json:"topics"`
}

type QueryTopicListJSONResponse struct {
	RequestId              int                    `json:"request_id"`
	QueryTopicListResponse QueryTopicListResponse `json:"response_params"`
}

type TagInfo struct {
	TagId      string `json:"tid"`
	TagName    string `json:"tag"`
	Info       string `json:"info"`
	CreateTime int64  `json:"createtime"`
}

type QueryTagsResponse struct {
	TotalNum int       `json:"total_num"`
	TagsInfo []TagInfo `json:"tags"`
}

type QueryTagsJSONResponse struct {
	RequestId         int               `json:"request_id"`
	QueryTagsResponse QueryTagsResponse `json:"response_params"`
}

type CreateTagResponse struct {
	TagName string `json:"tag"`
	Result  int    `json:"result"`
}

type CreateTagJSONResponse struct {
	RequestId         int               `json:"request_id"`
	CreateTagResponse CreateTagResponse `json:"response_params"`
}

type DeleteTagResponse struct {
	TagName string `json:"tag"`
	Result  int    `json:"result"`
}

type DeleteTagJSONResponse struct {
	RequestId         int               `json:"request_id"`
	DeleteTagResponse DeleteTagResponse `json:"response_params"`
}

type DeviceInfo struct {
	ChannelId string `json:"channel_id"`
	Result    int    `json:"result"`
}

type AddDevicesToTagResponse struct {
	DevicesInfoAfterAdded []DeviceInfo `json:"devices"`
}

type AddDevicesToTagJSONResponse struct {
	RequestId               int                     `json:"request_id"`
	AddDevicesToTagResponse AddDevicesToTagResponse `json:"response_params"`
}

type DeleteDevicesFromTagResponse struct {
	DevicesInfoAfterDel []DeviceInfo `json:"devices"`
}

type DeleteDevicesFromTagJSONResponse struct {
	RequestId                    int                          `json:"request_id"`
	DeleteDevicesFromTagResponse DeleteDevicesFromTagResponse `json:"response_params"`
}

type QueryDeviceNumInTagResponse struct {
	DeviceNum int `json:"device_num"`
}

type QueryDeviceNumInTagJSONResponse struct {
	RequestId                   int                         `json:"request_id"`
	QueryDeviceNumInTagResponse QueryDeviceNumInTagResponse `json:"response_params"`
}

type DeviceStatUnit struct {
	NewTerm    int `json:"new_term"`
	DelTerm    int `json:"del_term"`
	OnlineTerm int `json:"online_term"`
	AddupTerm  int `json:"addup_term"`
	TotalTerm  int `json:"total_term"`
}

type QueryStatisticDeviceResponse struct {
	TotalNum int                       `json:"total_num"`
	Result   map[string]DeviceStatUnit `json:"result"`
}

type QueryStatisticDeviceJSONResponse struct {
	RequestId                    int                          `json:"request_id"`
	QueryStatisticDeviceResponse QueryStatisticDeviceResponse `json:"response_params"`
}

type TopicStatUnit struct {
	Ack int `json:"ack"`
}

type QueryStatisticTopicResponse struct {
	TotalNum int                      `json:"total_num"`
	Result   map[string]TopicStatUnit `json:"result"`
}

type QueryStatisticTopicJSONResponse struct {
	RequestId                   int                         `json:"request_id"`
	QueryStatisticTopicResponse QueryStatisticTopicResponse `json:"response_params"`
}

type CancelTimerJobResponse struct {
	RequestId int `json:"request_id"`
}

// 服务器端错误返回信息
//
// 详见：http://push.baidu.com/doc/restapi/error_code
type ErrorResponse struct {
	RequestId int    `json:"request_id"` //请求ID
	Code      int    `json:"error_code"` //服务端错误码
	Message   string `json:"error_msg"`  //错误消息
}
