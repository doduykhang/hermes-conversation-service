/* name: GetUser :one */
SELECT * FROM users
WHERE id = ? LIMIT 1;

/* name: CreateUser :execresult */
INSERT INTO users (id, first_name, last_name, user_name, avatar)
VALUES (?, ?, ?, ?, ?);

/* name: UpdateUser :exec */
UPDATE users 
SET first_name = ?, last_name = ?, avatar = ?
WHERE id = ?;

/* name: DeleteUser :exec */
DELETE FROM users
WHERE id = ?



