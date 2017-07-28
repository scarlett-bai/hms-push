package hmspush

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// BatchNotification is for Silently push notification to mutiple devices
type BatchNotification struct {
	DeviceTokenList []string // A list of Huawei Device Tokens, at most 1000 can be set
	Message         string   // Message body of the push notification, which will be sent to the device
	CacheMode       int32    // Whether the notification need to be cached in the server. 0: Won't Cache, 1: Will Cache
	MsgType         int32    // If 2 notifications (towards same device) have the same MsgType, only the latter one would be cached by Huawei
	UserType        string   // For multiple users, 0: current user, 1: the primary user
	ExpireTime      string   // The expire time, its format should be ISO 8601 compliant: 2013-06-03T17:30:08+08:00
}

// NewBatchNotification creates a BatchNotification with default values
func NewBatchNotification(deviceTokenList []string, message string) *BatchNotification {
	return &BatchNotification{
		DeviceTokenList: deviceTokenList,
		Message:         message,
		CacheMode:       1,
		MsgType:         1,
		UserType:        "1",
		ExpireTime:      "",
	}
}

// SetMsgType msgType defaults to -1, which means Huawei will cache all msgs for that device.
// If you set msgType to any value in the range of 1 ~ 100, then Huawei will only
// cache the recent msg for that device.
func (s *BatchNotification) SetMsgType(msgType int32) *BatchNotification {
	s.MsgType = msgType
	return s
}

// SetTimeToLive set a TTL value in seconds
func (s *BatchNotification) SetTimeToLive(timeToLive int64) *BatchNotification {
	expireTimeStr := time.Now().Add(time.Second * time.Duration(timeToLive)).Format(TimeFormatHuawei)
	s.ExpireTime = expireTimeStr
	return s
}

// SetCacheMode set a cacheMode
func (s *BatchNotification) SetCacheMode(cacheMode int32) *BatchNotification {
	s.CacheMode = cacheMode
	return s
}

// SetUserType set a userType
func (s *BatchNotification) SetUserType(userType string) *BatchNotification {
	s.UserType = userType
	return s
}

// Form http parameters for sending BatchNotification
func (s *BatchNotification) Form(params url.Values) url.Values {
	params.Add("nsp_svc", batchSendURL)
	deviceTokenList := myMap(s.DeviceTokenList, func(token string) string {
		return "\"" + token + "\""
	})
	params.Add("deviceTokenList", "["+strings.Join(deviceTokenList, ",")+"]")
	params.Add("message", s.Message)
	params.Add("cacheMode", strconv.FormatInt(int64(s.CacheMode), 10))
	if s.MsgType > 0 {
		params.Add("msgType", strconv.FormatInt(int64(s.MsgType), 10))
	}
	if s.ExpireTime != "" {
		params.Add("expireTime", s.ExpireTime)
	}
	return params
}

func myMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
