-- name: GetParticipant :one
SELECT * FROM participant WHERE email = $1;

-- name: CreateParticipant :exec
INSERT INTO participant(email, track, name) VALUES ($1, $2, $3);