// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: room.sql

package mysql

import (
	"context"
)

const checkUserInRoom = `-- name: CheckUserInRoom :one
SELECT user_id, room_id 
FROM users_rooms
WHERE user_id = ? and room_id = ?
LIMIT 1
`

type CheckUserInRoomParams struct {
	UserID string
	RoomID string
}

func (q *Queries) CheckUserInRoom(ctx context.Context, arg CheckUserInRoomParams) (UsersRoom, error) {
	row := q.db.QueryRowContext(ctx, checkUserInRoom, arg.UserID, arg.RoomID)
	var i UsersRoom
	err := row.Scan(&i.UserID, &i.RoomID)
	return i, err
}

const createRoom = `-- name: CreateRoom :exec
INSERT INTO rooms (id, name, user_id, type)
VALUES (?, ?, ?, ?)
`

type CreateRoomParams struct {
	ID     string
	Name   string
	UserID string
	Type   string
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) error {
	_, err := q.db.ExecContext(ctx, createRoom,
		arg.ID,
		arg.Name,
		arg.UserID,
		arg.Type,
	)
	return err
}

const deleteRoom = `-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = ?
`

func (q *Queries) DeleteRoom(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteRoom, id)
	return err
}

const getRoomByID = `-- name: GetRoomByID :one
SELECT id, name, user_id, type, created_at, updated_at, deleted_at 
FROM rooms
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetRoomByID(ctx context.Context, id string) (Room, error) {
	row := q.db.QueryRowContext(ctx, getRoomByID, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRoomThatIn = `-- name: GetRoomThatIn :many
SELECT id, name, user_id, type, created_at, updated_at, deleted_at 
FROM rooms
WHERE id IN (
	SELECT ur.room_id 
	FROM users_rooms ur
	WHERE ur.user_id = ?
)
`

func (q *Queries) GetRoomThatIn(ctx context.Context, userID string) ([]Room, error) {
	rows, err := q.db.QueryContext(ctx, getRoomThatIn, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const updateRoom = `-- name: UpdateRoom :exec
UPDATE rooms
SET name = ?
WHERE id = ?
`

type UpdateRoomParams struct {
	Name string
	ID   string
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) error {
	_, err := q.db.ExecContext(ctx, updateRoom, arg.Name, arg.ID)
	return err
}
