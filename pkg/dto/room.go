package dto

import "encoding/json"

type RoomDTO struct {
	Audit
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Members json.RawMessage `json:"members"`
}

type Room struct {
	Audit
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateGroupRoomRequest struct {
	Name string `json:"name"`
	UserID string
}

type CreatePrivateRoomRequest struct {
	UserID string
	ReceiverID string `json:"receiverID"`
}

type GetUserRoomRequest struct {
	UserID string
}

type GetRoomRequest struct {
	UserID string
}

type UserRoom struct {
	UserID string `json:"userID"`
	RoomID string `json:"roomID"`
}

type GetUserPrivateRoomRequest struct {
	Audit
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserID    string `json:"userID"`
	Type      string `json:"type"`
	Members   json.RawMessage `json:"members"`
}
