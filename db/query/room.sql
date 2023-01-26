/* name: CreateRoom :exec */
INSERT INTO rooms (id, name, user_id, type)
VALUES (?, ?, ?, ?);

/* name: UpdateRoom :exec */
UPDATE rooms
SET name = ?
WHERE id = ?;

/* name: DeleteRoom :exec */
DELETE FROM rooms
WHERE id = ?;

/* name: GetRoomByID :one */
SELECT * 
FROM rooms
WHERE id = ?
LIMIT 1;

/* name: GetRoomThatIn :many */
SELECT * 
FROM rooms
WHERE id IN (
	SELECT ur.room_id 
	FROM users_rooms ur
	WHERE ur.user_id = ?
);


