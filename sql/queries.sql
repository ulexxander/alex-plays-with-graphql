
-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: CreateMessage :one
INSERT INTO messages (user_id, chat_id, content)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAllMessages :many
SELECT * FROM messages;

-- name: CreateChat :one
INSERT INTO chats (title)
VALUES ($1) RETURNING *;

-- name: GetAllChats :many
SELECT * FROM chats;

-- name: GetChatByID :one
SELECT * FROM chats WHERE chat_id = $1;

-- name: CountUserMessages :one
SELECT COUNT(message_id) FROM messages WHERE user_id = $1;

-- name: CountChatMessages :one
SELECT COUNT(message_id) FROM messages WHERE chat_id = $1;

-- name: GetChatMessages :many
SELECT * FROM messages WHERE chat_id = $1;

-- name: CountChatMembers :one
SELECT COUNT(user_id) FROM members WHERE chat_id = $1;

-- name: GetChatMembers :many
SELECT users.* FROM members
  INNER JOIN users
  ON members.user_id = users.user_id
WHERE chat_id = $1;

-- name: CreateChatMember :exec
INSERT INTO members (user_id, chat_id)
VALUES ($1, $2);
