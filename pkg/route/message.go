package route

import (
	"doduykhang/hermes-conversation/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func MessageRoute(r fiber.Router, messageController *controller.Message) {
	go messageController.WaitingForMessage()
	message := r.Group("/message")
	message.Get("/group/:roomID", messageController.GetMessageOfRoom)
	message.Post("/", messageController.CreateMessage)
}
