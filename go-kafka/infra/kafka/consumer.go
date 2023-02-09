// Responsável por consumir os dados da nossa fila do Kafka
package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// Obs.: esses valores após a declaração de parâmetros (*KafkaConsumer)
// indicam o tipo do retorno da função
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

// Aqui o Kafka escuta e consome as mensagens
func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	// loop inifnito
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// O canal k.MsgChan recebe a msg em cada loop. Assim o c.SubscribeTopics recebe essas mensagens
			k.MsgChan <- msg
		}
	}
}
