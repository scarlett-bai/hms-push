package hmspush

// PushResult is the response struct of all kinds of notification
type PushResult struct {
	ResultCode int64  `json:"resultcode"`
	Message    string `json:"message"`
	RequestID  string `json:"requestID"`
	Error      string `json:"error"`
}

// QueryMsgResult is the response struct of openpush.openapi.query_msg_result
type QueryMsgResult struct {
	RequestID string `json:"request_id"`
	Result    []struct {
		Token string `json:"token"`
		// Status:
		// 0: Delivered Successfully;
		// 1: Pending to be sent;
		// 2: Overwitten by a recent msg;
		// 3: Expired;
		// 11: Deivce is offline, and Cached for it;
		// 41: Device is offline but not cache for it;
		// 12: No route to deivce and cache for it;
		// 42: No route to device and not cache for it;
		// 13: Forwarded;
		// 43: APP Uninstalled.
		Status int32 `json:"status"`
	} `json:"result"`
	Error string `json:"error"`
}
