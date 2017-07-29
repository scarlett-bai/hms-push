# hms-push
华为推送服务HMS版  
A simple, fast [HMS Huawei PUSH](http://developer.huawei.com/consumer/cn/wiki/index.php?title=HMS%E6%9C%8D%E5%8A%A1%E4%BB%8B%E7%BB%8D-PUSH%E6%9C%8D%E5%8A%A1) in Golang  

This is an updated version of yilee's code (https://github.com/yilee/huawei-push), but specific to the new Huawei PUSH (HMS). 

HMS promises us a higher delivery rate, a smaller client-SDK, less battery-consumption, and a BULLSHIT documentation. ￣へ￣  

## APIs Supported:
RequestAccessToken  
SingleSend  
BatchSend  
PsSingleSend  
PsBatchSend  
QueryMsgResult  

## Usage
go get github.com/terry-xiaoyu/hms-push

```Go
// create a client
testClient := NewClient(testClientID, testClientSecret)

// the message body
message := "This is a silent notification to a single device."

// create a silent notification towards a single device
// and set the time to live to 300 seconds
notif := NewSingleNotification(testDeviceToken, message).SetTimeToLive(300)
result, err := testClient.SendPush(context.Background(), notif)
if err != nil {
    panic(fmt.Sprintf("Push failed! Error: %v", err))
}
fmt.Println("Push success, Result:", result)
```

More examples can be got from [client_test.go](https://github.com/terry-xiaoyu/hms-push/blob/master/client_test.go)

## TO DO

A production ready HMS provider will come soon.  
It will uses a simple Redis LIST as its input queue, TOML based config file, Metrics and Logs.
