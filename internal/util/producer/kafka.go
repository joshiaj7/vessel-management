package producer

//go:generate mockgen -source kafka.go -destination mock/kafka.go

type IKafkaProducer interface {
	ProduceWithRoutingKey() error
	Close() error
}

type KafkaProducer struct {
	ClientID string
	// Producer kafka.Producer
}

// func (p *KafkaProducer) Close() error {
// 	return p.Producer.Close()
// }

func (p *KafkaProducer) ProduceWithRoutingKey() error {

	return nil
}
