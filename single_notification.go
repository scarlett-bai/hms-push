package hmspush

import (
	"net/url"
	"strconv"
	"time"
)

// SingleNotification is for Silently push notification to a single device
type SingleNotification struct {
	DeviceToken string // Huawei Device Token
	Message     string // Message body of the push notification, which will be sent to the device
	Priority    int32  // Priority of the notification, 0: High, 1: Normal, defaults to 1
	CacheMode   int32  // Whether the notification need to be cached in the server. 0: Won't Cache, 1: Will Cache
	MsgType     int32  // If 2 notifications (towards same device) have the same MsgType, only the latter one would be cached by Huawei
	UserType    string // For multiple users, 0: current user, 1: the primary user
	ExpireTime  string // The expire time, its format should be ISO 8601 compliant: 2013-06-03T17:30:08+08:00
}

// NewSingleNotification creates a SingleNotification with defaults
func NewSingleNotification(deviceToken, message string) *SingleNotification {
	return &SingleNotification{
		DeviceToken: deviceToken,
		Message:     message,
		Priority:    1,
		CacheMode:   1,
		MsgType:     -1,
		UserType:    "1",
		ExpireTime:  "",
	}
}

// SetMsgType msgType defaults to -1, which means Huawei will cache all msgs for that device.
// If you set msgType to any value in the range of 1 ~ 100, then Huawei will only
// cache the recent msg for that device.
func (s *SingleNotification) SetMsgType(msgType int32) *SingleNotification {
	s.MsgType = msgType
	return s
}

// SetHighPriority set the priority to High
func (s *SingleNotification) SetHighPriority() *SingleNotification {
	s.Priority = 0
	return s
}

// SetTimeToLive set a TTL value in seconds
func (s *SingleNotification) SetTimeToLive(timeToLive int64) *SingleNotification {
	expireTimeStr := time.Now().Add(time.Second * time.Duration(timeToLive)).Format(TimeFormatHuawei)
	s.ExpireTime = expireTimeStr
	return s
}

// SetCacheMode set a cacheMode
func (s *SingleNotification) SetCacheMode(cacheMode int32) *SingleNotification {
	s.CacheMode = cacheMode
	return s
}

// SetUserType set a userType
func (s *SingleNotification) SetUserType(userType string) *SingleNotification {
	s.UserType = userType
	return s
}

// Form http parameters for sending SingleNotification
func (s *SingleNotification) Form(params url.Values) url.Values {
	params.Add("nsp_svc", singleSendURL)
	params.Add("deviceToken", s.DeviceToken)
	params.Add("userType", s.UserType)
	params.Add("message", s.Message)
	params.Add("priority", strconv.FormatInt(int64(s.Priority), 10))
	params.Add("cacheMode", strconv.FormatInt(int64(s.CacheMode), 10))
	if s.MsgType > 0 {
		params.Add("msgType", strconv.FormatInt(int64(s.MsgType), 10))
	}
	if s.ExpireTime != "" {
		params.Add("expireTime", s.ExpireTime)
	}
	return params
}
