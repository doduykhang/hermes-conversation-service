package service

import (
	"context"
	"database/sql"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"doduykhang/hermes-conversation/pkg/constant"
	"doduykhang/hermes-conversation/pkg/dto"
	"errors"
	"log"

	dtos "github.com/dranikpg/dto-mapper"

	"github.com/google/uuid"
)

type Room interface {
	CreateGroupRoom(request *dto.CreateGroupRoomRequest) (*dto.Room, error)
	CreatePrivateRoom(request *dto.CreatePrivateRoomRequest) (*dto.RoomDTO, error)
	GetUserRoom(userID string) ([]dto.RoomDTO, error)
	GetUserPrivateRoom(userID string) ([]dto.GetUserPrivateRoomRequest, error)
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
	room, err := s.queries.GetRoomByID(context.Background(), mysql.GetRoomByIDParams{
		ID: userID,
		ID_2: roomID,
	})
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

func (s *room) CreateGroupRoom(request *dto.CreateGroupRoomRequest) (*dto.Room, error) {
	var args mysql.CreateRoomParams
	args.ID = uuid.New().String()
	args.Name = request.Name
	args.UserID = request.UserID
	args.Type = constant.GroupRoom

	err := s.queries.CreateRoom(context.Background(), args)	
	if err != nil {
		return nil, err 
	}

	err = s.AddUserToRoom(
		&dto.UserRoom{
			UserID: args.UserID,
			RoomID: args.ID,
		},
		args.UserID,
	)

	if err != nil {
		log.Printf("Error CreateGroupRoom() %s\n", err)
		return nil, err 
	}
		
	var dto dto.Room
	err = s.mapper.Map(&dto, args)
	if err != nil {
		return nil, err 
	}

	return &dto, nil
}

func (r *room) checkPrivateRoomExist(senderID string, receiverID string) (bool, error) {
	var arg mysql.CheckPrivateRoomExistsParams
	arg.UserID = senderID
	arg.UserID_2 = receiverID
	_, err := r.queries.CheckPrivateRoomExists(context.Background(), arg)
	
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (s *room) CreatePrivateRoom(request *dto.CreatePrivateRoomRequest) (*dto.RoomDTO, error) {
	check, err := s.checkPrivateRoomExist(request.UserID, request.ReceiverID)
	if err != nil {
		return nil, err
	}

	if check {
			
		return nil, errors.New("Room existed")
	}

	var args mysql.CreateRoomParams
	args.ID = uuid.New().String()
	args.Name = request.UserID + "-" + request.ReceiverID
	args.Type = constant.PrivateRoom
	args.UserID = request.UserID

	err = s.queries.CreateRoom(context.Background(), args)	

	if err != nil {
		return nil, err
	}

	err = s.AddUserToRoom(
		&dto.UserRoom{
			UserID: request.UserID,
			RoomID: args.ID,
		},
		args.UserID,
	)

	if err != nil {
		return nil, err
	}

	err = s.AddUserToRoom(
		&dto.UserRoom{
			UserID: request.ReceiverID,
			RoomID: args.ID,
		},
		args.UserID,
	)

	if err != nil {
		return nil, err
	}

	var dto dto.RoomDTO
	err = s.mapper.Map(&dto, args)
	if err != nil {
		return nil, err 
	}
	
	return &dto, nil
}

func (s *room) GetUserPrivateRoom(userID string) ([]dto.GetUserPrivateRoomRequest, error) {
	rooms, err := s.queries.GetUserPrivateRoom(context.Background(), userID)	
	if err != nil {
		return nil, err
	}

	var dtos []dto.GetUserPrivateRoomRequest
	err = s.mapper.Map(&dtos, &rooms)

	if err != nil {
		return nil, err
	}

	return dtos, nil
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
		log.Printf("Error AddUserToRoom() %s\n", err)
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
	if request.UserID == userID {
		return errors.New("Can not remove yourself from room")
	}

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
