package helper

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/sgokul961/echo-hub-post-svc/pkg/models"
)

func PushCommentToQueue(topic string, message []byte) error {
	brokerUrl := []string{"localhost:9092"}

	producer, err := ConnectToProducer(brokerUrl)
	fmt.Println("err", err)
	if err != nil {
		return err // Return the error if connection to producer fails
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	fmt.Println("msg", msg)

	_, _, err = producer.SendMessage(msg)
	fmt.Println("err2", err)
	if err != nil {
		return err
	}

	fmt.Printf("Message sent successfully to topic: %s\n", topic)
	return nil
}

func ConnectToProducer(brokerUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerUrl, config)
	fmt.Println("conn,err", conn, err)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func PushLikeNotificationToQueue(notification models.Notification, message []byte) error {
	// Connect to Kafka producer
	brokerUrl := []string{"localhost:9092"}

	producer, err := ConnectToProducer(brokerUrl)
	fmt.Println("err", err)
	if err != nil {
		return err // Return the error if connection to producer fails
	}
	defer producer.Close()

	// Serialize notification message to JSON
	notification.Message = "you got one like"
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return err // Return error if serialization fails
	}

	msg := &sarama.ProducerMessage{
		Topic: notification.Topic,
		Value: sarama.StringEncoder(notificationJSON), // Encode as JSON string
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err // Return error if sending message fails
	}

	fmt.Printf("like sent successfully to topic: %s\n", notification.Topic)
	return nil
}
