package dto

import "time"

type Audit struct {
	CreatedAt time.Time `json:"createdAt"`	
	UpdatedAt time.Time `json:"updatedAt,omitempty"`	
	DeletedAt time.Time `json:"deletedAt"`	
}
