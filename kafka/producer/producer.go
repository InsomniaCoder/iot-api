package producer

import (
	"context"
	"fmt"

	"github.com/insomniacoder/iot-api/config"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func Produce(ctx context.Context, topic string, message string) error {

	kafKaConfig := config.Config.Kafka
	connectionString := fmt.Sprintf("%s:%d", kafKaConfig.Host, kafKaConfig.Port)

	log.Printf("publishing message %s to %s", message, connectionString)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{connectionString},
		Topic:   topic,
	})

	return w.WriteMessages(ctx, kafka.Message{
		Key:   []byte("key"),
		Value: []byte(message),
	})
}
