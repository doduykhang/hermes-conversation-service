// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: message.sql

package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (id, content, room_id, user_id)
VALUES (?, ?, ?, ?)
`

type CreateMessageParams struct {
	ID      string
	Content string
	RoomID  string
	UserID  string
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.ExecContext(ctx, createMessage,
		arg.ID,
		arg.Content,
		arg.RoomID,
		arg.UserID,
	)
	return err
}

const deleteMessage = `-- name: DeleteMessage :exec
UPDATE messages
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?
`

func (q *Queries) DeleteMessage(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteMessage, id)
	return err
}

const getMessageOfRoom = `-- name: GetMessageOfRoom :many
select id, content, user_id, room_id, created_at, updated_at, deleted_at, (
	SELECT JSON_OBJECT('id', u.id, 'firstName', u.first_name, 'lastName', u.last_name, 'avatar', u.avatar)
    	FROM users u
    	WHERE u.id = m.user_id
) as user
from messages m
where room_id = ?
`

type GetMessageOfRoomRow struct {
	ID        string
	Content   string
	UserID    string
	RoomID    string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
	User      json.RawMessage
}

func (q *Queries) GetMessageOfRoom(ctx context.Context, roomID string) ([]GetMessageOfRoomRow, error) {
	rows, err := q.db.QueryContext(ctx, getMessageOfRoom, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessageOfRoomRow
	for rows.Next() {
		var i GetMessageOfRoomRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserID,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.User,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMessage = `-- name: UpdateMessage :exec
UPDATE messages
SET content = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateMessageParams struct {
	Content string
	ID      string
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) error {
	_, err := q.db.ExecContext(ctx, updateMessage, arg.Content, arg.ID)
	return err
}