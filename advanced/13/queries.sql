-- name: CreateParticipant :exec
INSERT INTO participant(track, email) VALUES ($1, $2);
