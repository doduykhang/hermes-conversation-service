package service

import (
	"context"
	"doduykhang/hermes-conversation/pkg/dto"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue interface {
	WaitingForMessageEvent(messageCh chan dto.CreateMessageRequest) 
	PublishAddUserToRoomEvent(userID string, roomID string) error 
	PublishRemoveUserFromRoomEvent(userID string, roomID string) error 
}

type queue struct {
	conn *amqp.Connection
}

func NewQueue(conn *amqp.Connection) Queue {
	return &queue{
		conn: conn,
	}
}

func (queue *queue) WaitingForMessageEvent(messageCh chan dto.CreateMessageRequest) {
	ch, err := queue.conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
	}
	q, err := ch.QueueDeclare(
  		"create-message", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)

	if err != nil {
		log.Panic(err)
	}
	msgs, err := ch.Consume(
  		q.Name, // queue
  		"",     // consumer
  		true,   // auto-ack
  		false,  // exclusive
  		false,  // no-local
  		false,  // no-wait
  		nil,    // args
	)

	if err != nil {
		log.Panic(err)
	}
	
	for msg := range msgs {
		var request dto.CreateMessageRequest
		err := json.Unmarshal(msg.Body, &request)
		if err != nil {
			log.Printf("Error at service.Queue.WaitingForMessageEvent, %s\n", err)
			continue
		}
		messageCh <- request
	}
}
func (queue *queue) PublishAddUserToRoomEvent(userID string, roomID string) (error) {

	ch, err := queue.conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
	}
	q, err := ch.QueueDeclare(
  		"add-user", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)

	var request struct {
		UserID string `json:"userID"`
		RoomID string `json:"roomID"`
	}

	request.RoomID = roomID
	request.UserID = userID

	body, _ := json.Marshal(&request)

	err = ch.PublishWithContext(context.Background(),
  		"",     // exchange
  		q.Name, // routing key
  		false,  // mandatory
  		false,  // immediate
  		amqp.Publishing {
    			ContentType: "text/plain",
    			Body:        []byte(body),
  		},
	)

	return err
}
func (queue *queue) PublishRemoveUserFromRoomEvent(userID string, roomID string) (error) {
	ch, err := queue.conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
	}
	q, err := ch.QueueDeclare(
  		"delete-user", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)

	var request struct {
		UserID string `json:"userID"`
		RoomID string `json:"roomID"`
	}

	request.RoomID = roomID
	request.UserID = userID

	body, _ := json.Marshal(&request)

	err = ch.PublishWithContext(context.Background(),
  		"",     // exchange
  		q.Name, // routing key
  		false,  // mandatory
  		false,  // immediate
  		amqp.Publishing {
    			ContentType: "text/plain",
    			Body:        []byte(body),
  		},
	)

	return err
}


