package controller

import (
	"doduykhang/hermes-conversation/pkg/dto"
	"doduykhang/hermes-conversation/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type Room struct {
	roomService service.Room
}

func NewRoom(roomService service.Room) *Room {
	return &Room {
		roomService: roomService,
	}
}

func (r *Room) CreateGroupRoom(c *fiber.Ctx) error {
	var request dto.CreateGroupRoomRequest	
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}
	request.UserID = userID

	response, err := r.roomService.CreateGroupRoom(&request)
	if err != nil {
		return err
	}
	return c.JSON(response)
}

func (r *Room) GetUserRoom(c *fiber.Ctx) error {
	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}

	response, err := r.roomService.GetUserRoom(userID)
	if err != nil {
		return err
	}
	return c.JSON(response)
}

func (r *Room) AddUserToRoom(c *fiber.Ctx) error {
	var request dto.UserRoom
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}

	err := r.roomService.AddUserToRoom(&request, userID)
	if err != nil {
		return err
	}
	return c.Send([]byte("ok"))
}

func (r *Room) RemoveUserFromRoom(c *fiber.Ctx) error {
	var request dto.UserRoom
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}

	err := r.roomService.RemoveUserFromRoom(&request, userID)
	if err != nil {
		return err
	}
	return c.Send([]byte("ok"))
}

func (r *Room) GetRoomByID(c *fiber.Ctx) error {
	roomID := c.Params("roomID")
	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}

	response, err := r.roomService.GetRoomById(roomID, userID)
	if err != nil {
		return err
	}
	return c.JSON(response)
}
