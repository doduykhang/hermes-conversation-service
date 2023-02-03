package controller

import (
	"doduykhang/hermes-conversation/pkg/dto"
	"doduykhang/hermes-conversation/pkg/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	userService service.User		
	queueService service.Queue
}

func NewUser(userService service.User, queueService service.Queue) *User {
	return &User {
		userService: userService,
		queueService: queueService,
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
	email := c.Params("email")	
	roomID := c.Params("roomID")	

	response, err := u.userService.SearchForUserNotInRoom(roomID, email) 
	if err != nil {
        	return err
    	}
    	return c.JSON(response)
}

func (u *User) SearchUser(c *fiber.Ctx) error {
	email := c.Params("email")	
	userID := string(c.Request().Header.Peek("X-User-Id"))

	response, err := u.userService.SearchUsers(email, userID) 
	if err != nil {
        	return err
    	}
    	return c.JSON(response)
}


func (u *User) GetUser(c *fiber.Ctx) error {
	userID := string(c.Request().Header.Peek("X-User-Id"))
	if userID == "" {
		return fiber.ErrUnprocessableEntity
	}

	response, err := u.userService.GetProfile(userID) 
	if err != nil {
        	return err
    	}
    	return c.JSON(response)
}

func (u *User) WaitingForCreateUser() {
	messageCh := make(chan dto.CreateUser)
	go u.queueService.WaitingForCreateUserEvent(messageCh) 
	for msg := range messageCh {
		_, err := u.userService.CreateUser(&msg)	
		if err != nil {
			log.Printf("Error at controller.user.WaitingForCreateUser, %s\n", err)	
		}
	}
}
