package services

import (
	"github.com/segmentio/kafka-go"
	"github.com/gorilla/websocket"
	"log"
)

func WriteToKafka(writer *kafka.Writer, key string, value []byte) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
	}
	return writer.WriteMessages(context.Background(), msg)
}

func ConsumeKafkaStream(conn *websocket.Conn, topic string) {
	reader := NewKafkaReader("localhost:9092", topic, "group-id")
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, msg.Value)
		if err != nil {
			log.Println("Error sending message over WebSocket:", err)
			break
		}
	}
}
