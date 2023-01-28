package main

import (
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/config"
	"doduykhang/hermes-conversation/pkg/controller"
	"doduykhang/hermes-conversation/pkg/route"
	"doduykhang/hermes-conversation/pkg/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

 	api := app.Group("/api", logger.New()) // /api

	db := config.NewDB("mysql", "sammy:password@/hermes_conversation?parseTime=True&loc=Local")
	
	//config
	queries := mysql.New(db)
	mapper := config.NewMapper()
	rabbitMq := config.NewRabbitMq()
	
	//service 
	queueService := service.NewQueue(rabbitMq)
	authService := service.NewAuth(queries)
	userService := service.NewUser(queries, mapper)
	roomService := service.NewRoom(queries, mapper, authService, queueService)
	messageService := service.NewMessage(queries, mapper, authService)
	
	//controller 
	userController := controller.NewUser(userService, queueService)
	roomController := controller.NewRoom(roomService)
	messageController := controller.NewMessage(messageService, queueService)

	//route
	route.UserRoute(api, userController)
	route.RoomRoute(api, roomController)
	route.MessageRoute(api, messageController)

	log.Fatal(app.Listen(":8080"))
}
