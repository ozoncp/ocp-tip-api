package api

import (
	"github.com/Shopify/sarama"
	"strconv"
)

func prepareMessage(topic string, tipId uint64) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(strconv.Itoa(int(tipId))),
	}
	return msg
}
