// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//Rest Api请求参数构造
package push

import (
	"fmt"
	"strconv"
)

type QueryMsgStatusRequest struct {
	MsgId string
}

func (r *QueryMsgStatusRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("msg_id", r.MsgId)
	return params
}

type PushMsgToSingleDeviceRequest struct {
	ChannelId    string //唯一对应一台设备
	MsgType      int    //消息类型 0：消息；1：通知
	Message      string //消息内容，json格式, http://push.baidu.com/doc/restapi/msg_struct
	MsgExpires   int64  //相对于当前时间的消息过期时间，单位为秒 默认为5小时(18000秒)
	DeployStatus int    //设置iOS应用的部署状态，仅iOS应用推送时使用 1：开发状态；2：生产状态； 若不指定，则默认设置为生产状态。
}

func (r *PushMsgToSingleDeviceRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("channel_id", r.ChannelId)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	if r.MsgExpires != 0 {
		params.AddUnescaped("msg_expires", fmt.Sprintf("%d", (r.MsgExpires)))
	}

	if r.DeployStatus != 0 {
		params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	}
	return params
}

type PushMsgToAllRequest struct {
	MsgType      int    //消息类型 0：消息；1：通知
	Message      string //消息内容，json格式, http://push.baidu.com/doc/restapi/msg_struct
	MsgExpires   int    //相对于当前时间的消息过期时间，单位为秒 默认为5小时(18000秒)
	DeployStatus int    //设置iOS应用的部署状态，仅iOS应用推送时使用 1：开发状态；2：生产状态； 若不指定，则默认设置为生产状态。
	SendTime     int64  //定时推送，用于指定的实际发送时间 必须在当前时间60s以外，1年以内
}

func (r *PushMsgToAllRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	params.AddUnescaped("send_time", fmt.Sprintf("%d", (r.SendTime)))
	return params
}

type PushMsgToTagRequest struct {
	TagName      string //标签名 必须是已创建的
	MsgType      int    //消息类型 0：消息；1：通知
	Message      string //消息内容，json格式, http://push.baidu.com/doc/restapi/msg_struct
	MsgExpires   int    //相对于当前时间的消息过期时间，单位为秒 默认为5小时(18000秒)
	DeployStatus int    //设置iOS应用的部署状态，仅iOS应用推送时使用 1：开发状态；2：生产状态； 若不指定，则默认设置为生产状态。
	SendTime     int64  //定时推送，用于指定的实际发送时间 必须在当前时间60s以外，1年以内
}

func (r *PushMsgToTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("type", "1")
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	params.AddUnescaped("send_time", fmt.Sprintf("%d", (r.SendTime)))
	return params
}

type PushBatchUniMsgRequest struct {
	ChannelIds string //一组channel_id（最多为一万个）组成的json数组字符串
	MsgType    int    //消息类型 0：消息；1：通知
	Message    string //消息内容，json格式, http://push.baidu.com/doc/restapi/msg_struct
	MsgExpires int    //相对于当前时间的消息过期时间，单位为秒 默认为5小时(18000秒)
	TopicId    string //分类主题名称 字母、数字及下划线组成，长度限制为1~128
}

func (r *PushBatchUniMsgRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("channel_ids", r.ChannelIds)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	params.AddUnescaped("topic_id", r.TopicId)
	return params
}

type QueryTimerRecordsRequest struct {
	TimerId    string
	Start      int
	Limit      int
	RangeStart int64
	RangeEnd   int64
}

func (r *QueryTimerRecordsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("timer_id", r.TimerId)
	params.AddUnescaped("start", strconv.Itoa(r.Start))
	params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	params.AddUnescaped("range_start", fmt.Sprintf("%d", (r.RangeStart)))
	params.AddUnescaped("range_end", fmt.Sprintf("%d", (r.RangeEnd)))
	return params
}

type QueryTopicRecordsRequest struct {
	TopicId    string
	Start      int
	Limit      int
	RangeStart int64
	RangeEnd   int64
}

func (r *QueryTopicRecordsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("topic_id", r.TopicId)
	params.AddUnescaped("start", strconv.Itoa(r.Start))
	params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	params.AddUnescaped("range_start", fmt.Sprintf("%d", (r.RangeStart)))
	params.AddUnescaped("range_end", fmt.Sprintf("%d", (r.RangeEnd)))
	return params
}

type QueryTimerListRequest struct {
	TimerId string
	Start   int
	Limit   int
}

func (r *QueryTimerListRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("timer_id", r.TimerId)
	params.AddUnescaped("start", strconv.Itoa(r.Start))
	params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	return params
}

type QueryTopicListRequest struct {
	Start int
	Limit int
}

func (r *QueryTopicListRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("start", strconv.Itoa(r.Start))
	params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	return params
}

type QueryTagsRequest struct {
	TagName string
	Start   int
	Limit   int
}

func (r *QueryTagsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("start", strconv.Itoa(r.Start))
	params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	return params
}

type CreateTagRequest struct {
	TagName string
}

func (r *CreateTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type DeleteTagRequest struct {
	TagName string
}

func (r *DeleteTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type AddDevicesToTagRequest struct {
	TagName    string
	ChannelIds string
}

func (r *AddDevicesToTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("channel_ids", r.ChannelIds)
	return params
}

type DeleteDevicesFromTagRequest struct {
	TagName    string
	ChannelIds string
}

func (r *DeleteDevicesFromTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("channel_ids", r.ChannelIds)
	return params
}

type QueryDeviceNumInTagRequest struct {
	TagName string
}

func (r *QueryDeviceNumInTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type QueryStatisticTopicRequest struct {
	TopicId string
}

func (r *QueryStatisticTopicRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("topic_id", r.TopicId)
	return params
}
