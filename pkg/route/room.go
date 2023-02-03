package route

import (
	"doduykhang/hermes-conversation/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func RoomRoute(r fiber.Router, roomController *controller.Room) {
	room := r.Group("/room")

	groupRoom := room.Group("/group")
	groupRoom.Post("/", roomController.CreateGroupRoom)
	groupRoom.Get("/", roomController.GetUserRoom)
	groupRoom.Get("/:roomID", roomController.GetRoomByID)
	groupRoom.Post("/add-user", roomController.AddUserToRoom)
	groupRoom.Delete("/remove-user", roomController.RemoveUserFromRoom)

	privateRoom := room.Group("/private")

	privateRoom.Post("/", roomController.CreatePrivateRoom)
	privateRoom.Get("/", roomController.GetUserPrivateRoom)
}
