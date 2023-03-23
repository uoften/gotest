package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

// 初始化MQTT服务
func NewClient() {
	if client == nil {
		opts := mqtt.NewClientOptions()
		opts.AddBroker("tcp://broker.emqx.io:1883") // 这个中转服务器不需要任何账号密码
		opts.SetClientID("go_mqtt_client1")
		// opts.SetUsername("")
		// opts.SetPassword("")
		opts.OnConnect = func(c mqtt.Client) {
			fmt.Println("MQTT链接成功！")
		}
		opts.OnConnectionLost = func(c mqtt.Client, err error) {
			fmt.Println("MQTT断开链接！", err.Error())
			fmt.Println("尝试重新链接！")
			NewClient()
		}
		client = mqtt.NewClient(opts)
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	// 订阅事件
	for _, subscribe := range subscribes {
		subscribe()
	}
}
// 发布消息 ClientSend("topic/publish/example", 2, false, `{"code":0, "msg":"hello world!"}`)
func ClientSend(topic string, qos byte, retained bool, payload interface{}) error {
	if token := client.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		fmt.Println("消息发布失败！", token.Error())
		return token.Error()
	}
	return nil
}

// 订阅消息
func ClientSubscribe(topic string, qos byte, callback mqtt.MessageHandler, err func(error)) {
	if token := client.Subscribe(topic, qos, func(c mqtt.Client, msg mqtt.Message) {
		callback(c, msg)
	}); token.Wait() && token.Error() != nil {
		err(token.Error())
	}
}

//发布消息
func pushMsg()  {
	err := ClientSend("topic/publish/example", 2, false, `{"code":0, "msg":"hello world!"}`)
	fmt.Println(err)
}

// 订阅消息
var subscribes = []func(){
	// 直接写方法
	func() {
		ClientSubscribe("topic/subscribe/example", 1, func(c mqtt.Client, msg mqtt.Message) {
			fmt.Println("subscribe Msg：", string(msg.Payload()))
		}, func(err error) {
			fmt.Println(err.Error())
		})
	},
	// 调用
	subscribeExample2,
}

func subscribeExample2() {
	ClientSubscribe("topic/subscribe/example2", 1, func(c mqtt.Client, msg mqtt.Message) {
		fmt.Println("subscribe Msg2：", string(msg.Payload()))
	}, func(err error) {
		fmt.Println(err.Error())
	})
}