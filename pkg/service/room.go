package service

import (
	"context"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/constant"
	"doduykhang/hermes-conversation/pkg/dto"

	dtos "github.com/dranikpg/dto-mapper"

	"github.com/google/uuid"
)

type Room interface {
	CreateGroupRoom(request *dto.CreateGroupRoomRequest) (*dto.RoomDTO, error)
	CreatePrivateRoom(request *dto.CreatePrivateRoomRequest) (*dto.RoomDTO, error)
	GetUserRoom(userID string) ([]dto.RoomDTO, error)
	AddUserToRoom(request *dto.UserRoom) (error)
	RemoveUserFromRoom(request *dto.UserRoom) (error)
	GetRoomById(roomID string) (*dto.RoomDTO, error)
}

type room struct {
	queries *mysql.Queries
	mapper *dtos.Mapper
}


func NewRoom(queries *mysql.Queries, mapper *dtos.Mapper) Room {
	return &room{
		queries: queries,
		mapper: mapper,
	}
}

func (s *room) GetRoomById(roomID string) (*dto.RoomDTO, error) {
	room, err := s.queries.GetRoomByID(context.Background(), roomID)
	if err != nil {
		return nil, err 
	}
	var dto dto.RoomDTO
	err = s.mapper.Map(&dto, room)
	if err != nil {
		return nil, err 
	}
	return &dto, nil
}

func (s *room) CreateGroupRoom(request *dto.CreateGroupRoomRequest) (*dto.RoomDTO, error) {
	var args mysql.CreateRoomParams
	args.ID = uuid.New().String()
	args.Name = request.Name
	args.UserID = request.UserID
	args.Type = constant.GroupRoom

	err := s.queries.CreateRoom(context.Background(), args)	
	if err != nil {
		return nil, err 
	}

	
	err = s.queries.AddUserToRoom(context.Background(), mysql.AddUserToRoomParams{
		UserID: request.UserID,
		RoomID: args.ID,
	})

	var dto dto.RoomDTO
	err = s.mapper.Map(&dto, args)
	if err != nil {
		return nil, err 
	}

	return &dto, nil
}

func (room) CreatePrivateRoom(request *dto.CreatePrivateRoomRequest) (*dto.RoomDTO, error) {
	panic("not implemented") // TODO: Implement
}

func (s *room) GetUserRoom(userID string) ([]dto.RoomDTO, error) {
	rooms, err := s.queries.GetRoomThatIn(context.Background(), userID)	
	if err != nil {
		return nil, err 
	}
	
	var dto []dto.RoomDTO
	err = s.mapper.Map(&dto, &rooms)
	if err != nil {
		return nil, err 
	}

	return dto, nil
}

func (s *room) AddUserToRoom(request *dto.UserRoom) (error) {	
	var args mysql.AddUserToRoomParams
	args.UserID = request.UserID
	args.RoomID = request.RoomID
	return s.queries.AddUserToRoom(context.Background(), args)
}

func (s *room) RemoveUserFromRoom(request *dto.UserRoom) (error) {
	var args mysql.RemoveUserFromRoomParams
	args.UserID = request.UserID
	args.RoomID = request.RoomID
	return s.queries.RemoveUserFromRoom(context.Background(), args)
}