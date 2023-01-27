package dto

import "encoding/json"


type MessageDTO struct {
	Audit
	ID string `json:"id"`	
	Content string `json:"content"`
	Sender json.RawMessage `json:"sender"`
}

type CreateMessageRequest struct {
	Content string `json:"content"`
	UserID string 
	RoomID string `json:"roomID"`	
}


