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
SELECT *, (
	SELECT JSON_ARRAYAGG(JSON_OBJECT('id', u.id, 'firstName', u.first_name, 'lastName', u.last_name, 'email', u.email, 'avatar', u.avatar))
    	FROM users u
    	WHERE u.id IN (
		SELECT ur.user_id
		FROM users_rooms ur
		WHERE ur.room_id = r.id
	) AND u.id != ?
) as members
FROM rooms r
WHERE r.id = ?
LIMIT 1;

/* name: GetRoomNoMembers :one */
SELECT * FROM rooms r
WHERE r.id = ?
LIMIT 1;

/* name: GetRoomThatIn :many */
SELECT * 
FROM rooms r
WHERE id IN (
	SELECT ur.room_id 
	FROM users_rooms ur
	WHERE ur.user_id = ?
)
AND r.type = "GROUP";

/* name: CheckUserInRoom :one */
SELECT * 
FROM users_rooms
WHERE user_id = ? and room_id = ?
LIMIT 1;

/* name: CheckPrivateRoomExists :one */
SELECT *
FROM rooms r
WHERE r.type = "PRIVATE" 
AND EXISTS (
	SELECT ur.user_id
	FROM users_rooms ur
	WHERE ur.room_id = r.id AND ur.user_id = ?
)
AND EXISTS (
	SELECT ur.user_id
	FROM users_rooms ur
	WHERE ur.room_id = r.id AND ur.user_id = ?
);

/* name: GetUserPrivateRoom :many */
SELECT *, (
	SELECT JSON_ARRAYAGG(JSON_OBJECT('id', u.id, 'firstName', u.first_name, 'lastName', u.last_name, 'avatar', u.avatar))
    	FROM users u
    	WHERE u.id IN (
		SELECT ur.user_id
		FROM users_rooms ur
		WHERE ur.room_id = r.id
	)
) as members
FROM rooms r
WHERE r.type = "PRIVATE" 
AND EXISTS (
	SELECT *
	FROM users_rooms ur
	WHERE ur.user_id = ?
	AND ur.room_id = r.id
)

