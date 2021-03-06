package utils

import (
	"fmt"

	"github.com/Shopify/sarama"

	"encoding/json"
	"strconv"
	"time"
)

// CreateKafkaProducer creates producer
func CreateKafkaProducer(kafkaIP string) (sarama.AsyncProducer, error) {
	dial := kafkaIP + ":9092"
	config := sarama.NewConfig()
	brokers := []string{dial}

	producer, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		fmt.Println("failed to connect with Kafka", dial)
		return nil, err
	}
	return producer, nil
}

//SendToKafkaQueue sends to kafka
func SendToKafkaQueue(producer sarama.AsyncProducer, topic string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic: topic, // destination
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(data),
	}
	select {
	case producer.Input() <- msg:
		fmt.Println("Produce message", data)
	case err := <-producer.Errors():
		fmt.Println("Failed to produce message:", err)
		return err
	}
	return nil
}
