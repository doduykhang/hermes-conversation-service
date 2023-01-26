package route

import (
	"doduykhang/hermes-conversation/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func RoomRoute(r fiber.Router, roomController *controller.Room) {
	user := r.Group("/room")
	user.Post("/group", roomController.CreateGroupRoom)
	user.Get("/group", roomController.GetUserRoom)
	user.Get("/group/:roomID", roomController.GetRoomByID)
	user.Post("/group/add-user", roomController.AddUserToRoom)
	user.Delete("/group/remove-user", roomController.RemoveUserFromRoom)
}
