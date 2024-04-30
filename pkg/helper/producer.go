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

func PushLikeNotificationToQueue(notifications models.LikeNotification, message []byte) error {
	// Connect to Kafka producer
	brokerUrl := []string{"kafka:9092"}

	producer, err := ConnectToProducer(brokerUrl)
	fmt.Println("err", err)
	if err != nil {
		return err // Return the error if connection to producer fails
	}
	defer producer.Close()

	// Serialize notification message to JSON
	notifications.Message = "you got one like"
	notificationJSON, err := json.Marshal(notifications)
	if err != nil {
		return err // Return error if serialization fails
	}

	msg := &sarama.ProducerMessage{
		Topic: notifications.Topic,
		Value: sarama.StringEncoder(notificationJSON), // Encode as JSON string
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err // Return error if sending message fails
	}

	fmt.Printf("like sent successfully to topic: %s\n", notifications.Topic)
	return nil
}
func PushcommentNotificationToQueue(notification models.CommentNotification, message []byte) error {
	// Connect to Kafka producer
	brokerUrl := []string{"kafka:9092"}

	producer, err := ConnectToProducer(brokerUrl)

	if err != nil {
		return err // Return the error if connection to producer fails
	}
	defer producer.Close()

	// Serialize notification message to JSON
	notification.Message = "you got a new comment"

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

	fmt.Printf("comment sent successfully to topic: %s\n", notification.Topic)
	fmt.Println("comment contents is ", notification.Content)
	return nil
}
