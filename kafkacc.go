package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type consumerGroupHandler struct {
	name string
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	i := 0
	for msg := range claim.Messages() {
		fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		sess.MarkMessage(msg, "")
		i++
		//config.Consumer.Offsets.AutoCommit.Enable为false时，需要手动提交
		//每20条提交一次
		//if i%20 == 0 {
		//	sess.Commit()
		//}
	}
	return nil
}

//异步消费
func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_11_0_2
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, "test", config)
	//sarama.NewConsumerGroupFromClient()
	if err != nil {
		panic(err)
	}
	defer group.Close()

	for {
		cgHandler := consumerGroupHandler{name: "test"}
		err := group.Consume(context.Background(), []string{"testDemo"}, cgHandler)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}