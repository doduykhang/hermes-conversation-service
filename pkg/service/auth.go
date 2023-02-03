package service

import (
	"context"
	"database/sql"
	"doduykhang/hermes-conversation/internal/db/mysql"
	"log"
)

type Auth interface {
	CheckUserInRoom(userID string, roomID string) (bool, error)
	CheckUserOwnRoom(userID string, roomID string) (bool, error)
}

type auth struct {
	queries *mysql.Queries
}

func NewAuth(queries *mysql.Queries) Auth {
	return &auth{
		queries: queries,	
	}
}

func (s *auth) CheckUserInRoom(userID string, roomID string) (bool, error) {
	_, err := s.queries.CheckUserInRoom(context.Background(), mysql.CheckUserInRoomParams{
		UserID: userID,
		RoomID: roomID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func (s *auth) CheckUserOwnRoom(userID string, roomID string) (bool, error) {
	room, err := s.queries.GetRoomNoMembers(context.Background(), roomID)
	if err != nil {
		log.Printf("Error CheckUserOwnRoom() %s\n", err)
		return false, err
	}
	
	if room.UserID != userID {
		return false, nil
	}

	return true, nil
}
