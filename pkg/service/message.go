package service

import (
	"context"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/dto"
	"errors"

	dtos "github.com/dranikpg/dto-mapper"
	"github.com/google/uuid"
)

type Message interface {
	CreateMessage(request *dto.CreateMessageRequest) (*dto.MessageDTO, error)	
	GetMessageOfRoom(roomID string, userID string) ([]dto.MessageDTO, error)	
}

type message struct {
	queries *mysql.Queries
	mapper *dtos.Mapper
	auth Auth
}

func NewMessage(queries *mysql.Queries, mapper *dtos.Mapper, auth Auth) Message {
	return &message{
		queries: queries,
		mapper: mapper,
		auth: auth,
	}
}

func (s *message) CreateMessage(request *dto.CreateMessageRequest) (*dto.MessageDTO, error) {
	check, err := s.auth.CheckUserInRoom(request.UserID, request.RoomID)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("You are not allow to perform this action")
	}

	var args mysql.CreateMessageParams		
	err = s.mapper.Map(&args, request)
	if err != nil {
		return nil, err
	}
	args.ID = uuid.New().String()

	err = s.queries.CreateMessage(context.Background(), args)
	if err != nil {
		return nil, err
	}

	var dto dto.MessageDTO
	err = s.mapper.Map(&dto, &args)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
func (s *message) GetMessageOfRoom(roomID string, userID string) ([]dto.MessageDTO, error) {
	check, err := s.auth.CheckUserInRoom(userID, roomID)
	if err != nil {
		return nil, err
	}

	if !check {
		return nil, errors.New("You are not allow to perform this action")
	}

	messages, err := s.queries.GetMessageOfRoom(context.Background(), roomID)
	if err != nil {
		return nil, err
	}
	
	var dtos []dto.MessageDTO
	err = s.mapper.Map(&dtos, &messages)
	if err != nil {
		return nil, err
	}

	return dtos, nil
}
