### Publish a Message
resp, err := client.Publish(ctx, &pubsubproto.Message{
    Id:        "msg-123",
    TopicId:   "topic-1",
    Payload:   []byte("Hello, world!"),
    CreatedAt: time.Now().Unix(),
    ExpiresAt: time.Now().Add(time.Hour).Unix(),
})
if err != nil {
    log.Fatal(err)
}
fmt.Println("Published Message ID:", resp.MessageId)

### Subscribe to a Topic

stream, err := client.Subscribe(ctx, &pubsubproto.SubscribeRequest{
    SubscriptionId: "sub-1",
    VisibilityTimeout: 30,
})
if err != nil {
    log.Fatal(err)
}

for {
    msg, err := stream.Recv()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Received Message:", string(msg.Payload))
}

###Acknowledge a Message
ackResp, err := client.Acknowledge(ctx, &pubsubproto.AcknowledgeRequest{
    SubscriptionId: "sub-1",
    MessageId:      "msg-123",
})
if err != nil {
    log.Fatal(err)
}
fmt.Println("Acknowledged:", ackResp.Acknowledge)


###Modify Visibility Timeout
resp, err := client.ModifyVisibilityTimeout(ctx, &pubsubproto.ModifyVisibilityTimeoutRequest{
    MessageId: "msg-123",
    NewTimeout: 60,
})
if err != nil {
    log.Fatal(err)
}
fmt.Println("Visibility Timeout Updated:", resp.Success)
