package route

import (
	"doduykhang/hermes-conversation/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, userController *controller.User) {
	go userController.WaitingForCreateUser()
	user := r.Group("/user")
	user.Post("/", userController.CreateUser)
	user.Get("/:roomID/:userName", userController.SearchUserNotInRoom)
}
