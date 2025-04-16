package consumer

import (
	"github.com/IBM/sarama"
)

type consumerGroupHandler struct{}

func NewConsunerGroupHandler() sarama.ConsumerGroupHandler {
	return consumerGroupHandler{}
}

func (obj consumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		ConsumeMessageTransaction(msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}
