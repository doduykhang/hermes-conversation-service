package service

import (
	"context"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/constant"
	"doduykhang/hermes-conversation/pkg/dto"
	"errors"

	dtos "github.com/dranikpg/dto-mapper"

	"github.com/google/uuid"
)

type Room interface {
	CreateGroupRoom(request *dto.CreateGroupRoomRequest) (*dto.RoomDTO, error)
	CreatePrivateRoom(request *dto.CreatePrivateRoomRequest) (*dto.RoomDTO, error)
	GetUserRoom(userID string) ([]dto.RoomDTO, error)
	AddUserToRoom(request *dto.UserRoom, userID string) (error)
	RemoveUserFromRoom(request *dto.UserRoom, userID string) (error)
	GetRoomById(roomID string, userID string) (*dto.RoomDTO, error)
}

type room struct {
	queries *mysql.Queries
	mapper *dtos.Mapper
	auth Auth
	queue Queue
	userService User	
}

func NewRoom(queries *mysql.Queries, mapper *dtos.Mapper, auth Auth, queue Queue, userService User) Room {
	return &room{
		queries: queries,
		mapper: mapper,
		auth: auth,
		queue: queue,
		userService: userService,
	}
}

func (s *room) GetRoomById(roomID string, userID string) (*dto.RoomDTO, error) {
	check, err := s.auth.CheckUserInRoom(userID, roomID)
	if err != nil {
		return nil, err 
	}

	if !check {
		return nil, errors.New("You are not allowed here, get out")
	}
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

	s.AddUserToRoom(
		&dto.UserRoom{
			UserID: args.UserID,
			RoomID: args.ID,
		},
		args.UserID,
	)
		
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

func (s *room) AddUserToRoom(request *dto.UserRoom, userID string) (error) {	
	check, err := s.auth.CheckUserOwnRoom(userID, request.RoomID)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("You are not allow to do this action")
	}

	var args mysql.AddUserToRoomParams
	args.UserID = request.UserID
	args.RoomID = request.RoomID
	err = s.queue.PublishAddUserToRoomEvent(request.UserID, request.RoomID)
	if err != nil {
		return err
	}
	return s.queries.AddUserToRoom(context.Background(), args)
}

func (s *room) RemoveUserFromRoom(request *dto.UserRoom, userID string) (error) {
	check, err := s.auth.CheckUserOwnRoom(userID, request.RoomID)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("You are not allow to do this action")
	}

	var args mysql.RemoveUserFromRoomParams
	args.UserID = request.UserID
	args.RoomID = request.RoomID
	err = s.queue.PublishRemoveUserFromRoomEvent(request.UserID, request.RoomID)
	if err != nil {
		return err
	}
	return s.queries.RemoveUserFromRoom(context.Background(), args)
}
