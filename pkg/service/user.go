package service

import (
	"context"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/dto"
	"log"

	dtos "github.com/dranikpg/dto-mapper"
)

type User interface {
	CreateUser(request *dto.CreateUser) (*dto.UserDTO, error)
	SearchForUserNotInRoom(roomID string, email string) ([]dto.UserDTO, error)
	SearchUsers(email string, userID string) ([]dto.UserDTO, error)
	GetProfile(userID string) (*dto.UserDTO, error)	
}

type user struct {
	queries *mysql.Queries
	mapper *dtos.Mapper
}

func NewUser(queries *mysql.Queries, mapper *dtos.Mapper) User {
	return &user{
		queries: queries,
		mapper: mapper,
	}
}

func (s *user) CreateUser(request *dto.CreateUser) (*dto.UserDTO, error) {
	var args mysql.CreateUserParams
	err := s.mapper.Map(&args, request)

	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	_, err = s.queries.CreateUser(context.Background(), args)
	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	var dto dto.UserDTO
	dto.ID = args.ID
	err = s.mapper.Map(&dto, &args)

	if err != nil {
		log.Printf("Error at service.user.CreateUser, %s\n", err)
		return nil, err
	}

	return &dto, nil
}

func (s *user) SearchForUserNotInRoom(roomID string, email string) ([]dto.UserDTO, error) {
	users, err := s.queries.SearchUserNotInRoom(context.Background(), mysql.SearchUserNotInRoomParams{
		Email: "%" + email + "%",
		RoomID: roomID,
		
	})	
	if err != nil {
		return nil, err
	}

	var dtos []dto.UserDTO
	err = s.mapper.Map(&dtos, &users)
	if err != nil {
		return nil, err
	}
	
	return dtos, nil
}

func (s *user) SearchUsers(email string, userID string) ([]dto.UserDTO, error) {
	var args mysql.SearchUserParams
	args.Email = "%" + email + "%"
	args.UserID = userID
	args.ID = userID

	users, err := s.queries.SearchUser(context.Background(), args)	
	if err != nil {
		return nil, err
	}

	var dtos []dto.UserDTO
	err = s.mapper.Map(&dtos, &users)
	if err != nil {
		return nil, err
	}
	
	return dtos, nil
}


func (s *user) GetProfile(userID string) (*dto.UserDTO, error) {
	user, err := s.queries.GetUser(context.Background(), userID)	
	if err != nil {
		return nil, err
	}

	var dto dto.UserDTO
	err = s.mapper.Map(&dto, &user)
	if err != nil {
		return nil, err
	}
	
	return &dto, nil
}
