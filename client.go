package hmspush

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

// HuaweiPushClient is a client structure used to send notifications
type HuaweiPushClient struct {
	clientID, clientSecret string
}

// NewClient returns a *HuaweiPushClient
func NewClient(clientID, clientSecret string) *HuaweiPushClient {
	return &HuaweiPushClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (c *HuaweiPushClient) defaultParams(params url.Values) (url.Values, error) {
	accessToken, err := RequestAccess(c.clientID, c.clientSecret)
	if err != nil {
		return params, err
	}
	if accessToken.Error != 0 {
		return params, errors.New(accessToken.ErrorDescription)
	}
	params.Add("nsp_ts", strconv.FormatInt(time.Now().Unix(), 10))
	params.Add("nsp_fmt", "JSON")
	params.Add("access_token", accessToken.AccessToken)
	return params, nil
}

// SendPush can be used to send all kinds of notifications:
//  1. SingleNotification
//  2. BatchNotification
//  3. PsSingleNotification
//  4. PsBatchNotification
func (c *HuaweiPushClient) SendPush(ctx context.Context, notification interface{}) (*PushResult, error) {
	params, err := c.defaultParams(url.Values{})
	if err != nil {
		return nil, err
	}

	switch notif := (notification).(type) {
	case *SingleNotification:
		params = notif.Form(params)
	case *BatchNotification:
		params = notif.Form(params)
	case *PsSingleNotification:
		params = notif.Form(params)
	case *PsBatchNotification:
		params = notif.Form(params)
	default:
		return nil, errors.New(fmt.Sprint("Unsupported notifiction type: ", reflect.TypeOf(notification)))
	}

	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result PushResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		tokenInstance.AccessToken = ""
		return c.SendPush(ctx, notification)
	}
	return &result, nil
}

// QueryMsgResult query send result
// But is seems that this is not work, at least for now.
// The query result is always empty
func (c *HuaweiPushClient) QueryMsgResult(ctx context.Context, requestID, token string) (*QueryMsgResult, error) {
	params := url.Values{}
	params, err := c.defaultParams(params)
	if err != nil {
		return nil, err
	}
	params.Add("request_id", requestID)
	if token != "" {
		params.Add("token", token)
	}
	params.Add("nsp_svc", queryMsgResultURL)
	bytes, err := doPost(ctx, baseAPI, params)
	if err != nil {
		return nil, err
	}
	var result QueryMsgResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Error == SessionTimeoutError || result.Error == SessionInvalidError {
		tokenInstance.AccessToken = ""
		return c.QueryMsgResult(ctx, requestID, token)
	}
	return &result, nil
}
