package controller

import (
	"doduykhang/hermes-conversation/pkg/dto"
	"doduykhang/hermes-conversation/pkg/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	messageService service.Message
	queueService service.Queue
}

func NewMessage(service service.Message, queue service.Queue) *Message {
	return &Message{
		messageService: service,
		queueService: queue,
	}
}

func (m *Message) CreateMessage(c *fiber.Ctx) error {
	var request dto.CreateMessageRequest	
	if err := c.BodyParser(&request); err != nil {
		return err
	}	

	userID := string(c.Request().Header.Peek("X-User-Id"))
	request.UserID = userID
	
	response, err := m.messageService.CreateMessage(&request)
	if err != nil {
		return err
	}
	
	return c.JSON(response)
}

func (m *Message) GetMessageOfRoom(c *fiber.Ctx) error {
	roomID := c.Params("roomID")
	userID := string(c.Request().Header.Peek("X-User-Id"))
	
	response, err := m.messageService.GetMessageOfRoom(roomID, userID)
	if err != nil {
		return err
	}
	
	return c.JSON(response)
}

func (m *Message) WaitingForMessage() {
	messageCh := make(chan dto.CreateMessageRequest)
	go m.queueService.WaitingForMessageEvent(messageCh) 
	for msg := range messageCh {
		_, err := m.messageService.CreateMessage(&msg)	
		if err != nil {
			log.Printf("Error at controller.message.WaitingForMessage, %s\n", err)	
		}
	}
}
