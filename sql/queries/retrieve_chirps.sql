-- name: RetrieveChirps :many 
SELECT * FROM chirps ORDER BY created_at ASC;
--
-- name: GetChirpByID :one 
SELECT * FROM chirps WHERE id = $1;
--
-- name: RetrieveChirpsByAuthor :many
SELECT * FROM chirps WHERE user_id = $1 ORDER BY created_at ASC;
