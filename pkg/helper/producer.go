package helper

import (
	"fmt"

	"github.com/IBM/sarama"
)

func PushCommentToQueue(topic string, message []byte) error {
	brokerUrl := []string{"localhost:9092"}
	//brokerURL := "kafka:9092"

	//producer, err := ConnectToProducer(brokerUrl)

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
	// partition, offset, err := producer.SendMessage(msg)
	// if err != nil {
	// 	return err // Return the error if sending message fails
	// }
	// fmt.Printf("Message stored in topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	// return nil // Return nil to indicate success

	// Send message asynchronously
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
