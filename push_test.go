// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 由于百度不提供API沙箱测试环境，本示例仅说明简单的单元测试方法
package push

import (
	"encoding/json"
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type PUSHTestSuite struct {
	client *BaiduPushClient
}

var _ = check.Suite(&PUSHTestSuite{NewClient(
	"fH5CGHFKliIad9LGE7RwWDLs",
	"nO29YHEp3RbhizVyVIrAd2MtYIYvxtpp",
)})

// 测试查询当前应用的设备统计信息
func (s *PUSHTestSuite) TestQueryStatisticDevice(c *check.C) {
	_, err := s.client.QueryStatisticDevice()
	c.Assert(err, check.IsNil)
}

// 测试推送消息到单个客户端
func (s *PUSHTestSuite) TestPushMsgToSingleDevice(c *check.C) {
	// 构造请求
	request := &PushMsgToSingleDeviceRequest{ChannelId: "4548209656360317009", MsgType: 0, Message: "测试消息"}
	// 推送消息到指定客户端
	_, err := s.client.PushMsgToSingleDevice(*request)
	c.Assert(err, check.IsNil)
}

// 测试推送系统通知到单个客户端
func (s *PUSHTestSuite) TestPushNotificationToSingleDevice(c *check.C) {
	// 推送消息到指定客户端
	notification := AndroidNotification{Title: "测试通知", Description: "测试通知描述", NotificationBuilderId: 0, NotificationBasicStyle: 7, OpenType: 1, Url: "http://www.tsuru.cn"}
	message, _ := json.Marshal(notification)
	request := &PushMsgToSingleDeviceRequest{ChannelId: "4548209656360317009", MsgType: 1, Message: string(message)}
	_, err := s.client.PushMsgToSingleDevice(*request)
	c.Assert(err, check.IsNil)
}
