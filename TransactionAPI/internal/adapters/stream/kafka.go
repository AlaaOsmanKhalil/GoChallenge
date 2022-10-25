package stream

import (
	"TransactionAPI/config"
	"TransactionAPI/internal/repositories/transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"time"
)

var Kafkaconfig = kafka.ReaderConfig{}

func InitializeKafka(configurations *config.Configurations) {
	Kafkaconfig.Brokers = []string{configurations.Kafka.Broker}
	Kafkaconfig.Topic = configurations.Kafka.Topic
	Kafkaconfig.MaxBytes = configurations.Kafka.MaxBytes
}

func Produce(transaction transaction.Model) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "broker:29092", "topic_transaction", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}

func Consume(log *zap.SugaredLogger) {
	reader := kafka.NewReader(Kafkaconfig)
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			log.Fatalf(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
	}
}
