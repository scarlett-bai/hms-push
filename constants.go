package hmspush

import "time"

const (
	accessTokenAPI    = "https://login.vmall.com/oauth2/token"
	baseAPI           = "https://api.vmall.com/rest.php"
	singleSendURL     = "openpush.message.single_send"
	batchSendURL      = "openpush.message.batch_send"
	psSingleSendURL   = "openpush.message.psSingleSend"
	psBatchSendURL    = "openpush.message.psBatchSend"
	queryMsgResultURL = "openpush.openapi.query_msg_result"
)

const (
	// NoPermission to send message to these tmIDs, may need to resend using
	// an updated accessToken
	NoPermission        = 20203
	SessionTimeoutError = "session timeout"
	SessionInvalidError = "invalid session"
)

const (
	// TimeFormatHuawei "ISO 8601: 2013-06-03T17:30:08+08:00"
	TimeFormatHuawei = time.RFC3339
	// TimeFormatHuaweiOld is the time format of old version of Huawei Push
	// psSingleSend and psBatchSend use this time format, this may be a bug in Huawei
	TimeFormatHuaweiOld = "2006-01-02 15:04"
)
