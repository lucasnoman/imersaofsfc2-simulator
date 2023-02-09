package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/lucasnoman/imersaofsfc2-simulator/application/kafka"
	"github.com/lucasnoman/imersaofsfc2-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

/*
****** Para rodar: ******

Terminal 1:
docker exec -it kafka-kafka-1 bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction

Terminal 2:
docker exec -it kafka-kafka-1 bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal

Terminal 3:
go run main.go
"Após a mensagem de 'started', jogue os JSONs abaixo um por um"
*/
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
