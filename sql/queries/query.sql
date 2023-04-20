-- name: FindChatById :one
SELECT * FROM chats WHERE id = ?;