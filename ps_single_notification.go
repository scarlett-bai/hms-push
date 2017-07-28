package hmspush

import (
	"net/url"
	"strconv"
	"time"
)

// PsSingleNotification is for push notification to a single device (show in the notification bar)
type PsSingleNotification struct {
	DeviceToken string          // Huawei Device Token
	Android     *AndroidMessage // The Android Structure used for the notification bar
	CacheMode   int32           // Whether the notification need to be cached in the server. 0: Won't Cache, 1: Will Cache
	MsgType     int32           // If 2 notifications (towards same device) have the same MsgType, only the latter one would be cached by Huawei
	UserType    string          // For multiple users, 0: current user, 1: the primary user
	ExpireTime  string          // The expire time, its format should be ISO 8601 compliant: 2013-06-03T17:30:08+08:00
}

// NewPsSingleNotification creates a PsSingleNotification with default values
func NewPsSingleNotification(deviceToken string, android *AndroidMessage) *PsSingleNotification {
	return &PsSingleNotification{
		DeviceToken: deviceToken,
		Android:     android,
		CacheMode:   1,
		MsgType:     1,
		UserType:    "1",
		ExpireTime:  "",
	}
}

// SetMsgType msgType defaults to -1, which means Huawei will cache all msgs for that device.
// If you set msgType to any value in the range of 1 ~ 100, then Huawei will only
// cache the recent msg for that device.
func (s *PsSingleNotification) SetMsgType(msgType int32) *PsSingleNotification {
	s.MsgType = msgType
	return s
}

// SetTimeToLive set a TTL value in seconds
func (s *PsSingleNotification) SetTimeToLive(timeToLive int64) *PsSingleNotification {
	expireTimeStr := time.Now().Add(time.Second * time.Duration(timeToLive)).Format(TimeFormatHuaweiOld)
	s.ExpireTime = expireTimeStr
	return s
}

// SetCacheMode set a cacheMode
func (s *PsSingleNotification) SetCacheMode(cacheMode int32) *PsSingleNotification {
	s.CacheMode = cacheMode
	return s
}

// SetUserType set a userType
func (s *PsSingleNotification) SetUserType(userType string) *PsSingleNotification {
	s.UserType = userType
	return s
}

// Form http parameters for sending PsSingleNotification
func (s *PsSingleNotification) Form(params url.Values) url.Values {
	params.Add("nsp_svc", psSingleSendURL)
	params.Add("deviceToken", s.DeviceToken)
	params.Add("userType", s.UserType)
	params.Add("android", s.Android.String())
	params.Add("cacheMode", strconv.FormatInt(int64(s.CacheMode), 10))
	if s.MsgType > 0 {
		params.Add("msgType", strconv.FormatInt(int64(s.MsgType), 10))
	}
	if s.ExpireTime != "" {
		params.Add("expireTime", s.ExpireTime)
	}
	return params
}
