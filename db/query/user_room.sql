/* name: AddUserToRoom :exec */
INSERT INTO users_rooms (user_id, room_id)
VALUES (?, ?);

/* name: RemoveUserFromRoom :exec */
DELETE FROM users_rooms 
WHERE user_id = ? AND room_id = ?;
