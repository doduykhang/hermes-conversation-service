package controller

import (
	"doduykhang/hermes-conversation/pkg/dto"
	"doduykhang/hermes-conversation/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	userService service.User		
}

func NewUser(userService service.User) *User {
	return &User {
		userService: userService,
	}
}

func (u *User) CreateUser(c *fiber.Ctx) error {
	var request dto.CreateUser
	if err := c.BodyParser(&request); err != nil {
        	return err
    	}
	
	response, err := u.userService.CreateUser(&request) 
	if err != nil {
        	return err
    	}
    	return c.JSON(response)
}

func (u *User) SearchUserNotInRoom(c *fiber.Ctx) error {
	var request dto.CreateUser
	if err := c.BodyParser(&request); err != nil {
        	return err
    	}
	userName := c.Params("userName")	
	roomID := c.Params("roomID")	

	response, err := u.userService.SearchForUserNotInRoom(roomID, userName) 
	if err != nil {
        	return err
    	}
    	return c.JSON(response)
}
