package publisher

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentCreatedEvent struct {
	CustomerName  string `json:"customer_name"`
	Shopp         string `json:"shopp"`
	TransactionID string `json:"transaction_id"`
	Data          string `json:"data"`
	Price         string `json:"price"`
}

type PaymentPublisher struct {
	channel *amqp.Channel
}

func NewPaymentPublisher(ch *amqp.Channel) *PaymentPublisher {
	return &PaymentPublisher{
		channel: ch,
	}
}

func (p *PaymentPublisher) PublishPaymentCreated(customerName string, shopp string, transactionID string, data string, price string) error {
	event := PaymentCreatedEvent{
		CustomerName:  customerName,
		Shopp:         shopp,
		TransactionID: transactionID,
		Data:          data,
		Price:         price,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",
		"payment.created",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("[publisher] mensagem publicada na fila payment.created: transaction_id=%s", transactionID)
	return nil
}
