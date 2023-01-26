/* name: CreateMessage :exec */
INSERT INTO messages (id, content, room_id, user_id)
VALUES (?, ?, ?, ?);

/* name: UpdateMessage :exec */
UPDATE messages
SET content = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

/* name: DeleteMessage :exec */
UPDATE messages
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;

/* name: GetMessageOfRoom :many */
select *, (
	SELECT JSON_OBJECT('id', u.id, 'firstName', u.first_name, 'lastName', u.last_name, 'avatar', u.avatar)
    	FROM users u
    	WHERE u.id = m.user_id
) as user
from messages m
where room_id = ?;
