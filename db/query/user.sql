/* name: GetUser :one */
SELECT * FROM users
WHERE id = ? LIMIT 1;

/* name: CreateUser :execresult */
INSERT INTO users (id, first_name, last_name, email, avatar)
VALUES (?, ?, ?, ?, ?);

/* name: UpdateUser :exec */
UPDATE users 
SET first_name = ?, last_name = ?, avatar = ?
WHERE id = ?;

/* name: DeleteUser :exec */
DELETE FROM users
WHERE id = ?;

/* name: SearchUserNotInRoom :many */
SELECT * 
FROM users u
WHERE u.id NOT IN (
	SELECT user_id
	FROM users_rooms ur 
	WHERE room_id = ?
)
AND u.email LIKE ?;

/* name: SearchUser :many */
SELECT * 
FROM users u
WHERE u.email LIKE ?
AND NOT EXISTS (
	SELECT * 
	FROM rooms r
	WHERE r.type = "PRIVATE"
	AND (
		SELECT COUNT(*)
		FROM users_rooms ur
		WHERE ur.room_id = r.id AND (ur.user_id = u.id OR ur.user_id = ?)
	) = 2
)
AND u.id != ?

