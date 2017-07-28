package hmspush

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
)

const (
	testClientID     = "your-client-id"
	testClientSecret = "your-client-secret"
	testDeviceToken  = "you-device-token"
)

var testDeviceTokenList = []string{
	"you-device-token",
	"you-device-token",
}

func init() {
	Init("clientID", "clientSecret")
}

var defaultClient *HuaweiPushClient

func Init(clientID, clientSecret string) {
	defaultClient = &HuaweiPushClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func Test_RequestAccess(t *testing.T) {
	token, err := RequestAccess(testClientID, testClientSecret)
	if err != nil {
		t.Errorf("err=%v\n", err)
		return
	}
	if token.Error != 0 {
		t.Errorf("token=%+v\n", token)
		return
	}
	fmt.Printf("Test_RequestAccess -  Got Token: %+v\n", token)
}

func Test_SingleSend(t *testing.T) {
	testClient := NewClient(testClientID, testClientSecret)
	message := "This is a silent notification to a single device."
	notif := NewSingleNotification(testDeviceToken, message).SetTimeToLive(5)
	result, err := testClient.SendPush(context.Background(), notif)
	if err != nil {
		t.Errorf("err=%v\n\n", err)
		return
	}
	fmt.Printf("Test_SingleSend - result=%+v\n", result)

	// Query Msg Result
	// But it seems that the query result is aways empty, this may a bug of Huawei
	qresult, qerr := defaultClient.QueryMsgResult(context.TODO(), result.RequestID, testDeviceToken)
	if qerr != nil {
		t.Errorf("TestHuaweiPushClient_QueryResult err=%v\n\n", qerr)
		return
	}
	fmt.Printf("TestHuaweiPushClient_QueryResult - qresult=%+v\n\n", qresult)
	t.Logf("qresult=%#v\n", qresult)
}

func Test_BatchSend(t *testing.T) {
	testClient := NewClient(testClientID, testClientSecret)
	message := "This is a silent notification to a list of devices."
	notif := NewBatchNotification(testDeviceTokenList, message).SetTimeToLive(5)
	result, err := testClient.SendPush(context.Background(), notif)
	if err != nil {
		t.Errorf("err=%v\n\n", err)
		return
	}
	fmt.Printf("Test_BatchSend - result=%+v\n\n", result)
}

func Test_PsSingleSend(t *testing.T) {
	testClient := NewClient(testClientID, testClientSecret)
	message := NewAndroidMessage("notificationTitle", "This is a PS notification to a single device.")
	// PsSingleNotification uses a old version of time format, with a minutes-precision
	// So you can only set a TimeToLive greater than 60 (sec)
	notif := NewPsSingleNotification(testDeviceToken, message).SetTimeToLive(70)
	result, err := testClient.SendPush(context.Background(), notif)
	if err != nil {
		t.Errorf("err=%v\n\n", err)
		return
	}
	fmt.Printf("Test_PsSingleSend - result=%+v\n\n", result)
}

func Test_PsBatchSend(t *testing.T) {
	testClient := NewClient(testClientID, testClientSecret)
	message := NewAndroidMessage("notificationTitle", "This is a PS notification to a list of devices.")
	// PsBatchNotification uses a old version of time format, with a minutes-precision
	// So you can only set a TimeToLive greater than 60 (sec)
	notif := NewPsBatchNotification(testDeviceTokenList, message).SetTimeToLive(70)
	result, err := testClient.SendPush(context.Background(), notif)
	if err != nil {
		t.Errorf("err=%v\n\n", err)
		return
	}
	fmt.Printf("Test_PsBatchSend - result=%+v\n\n", result)
}
