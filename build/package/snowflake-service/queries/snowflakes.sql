-- name: Create :exec
INSERT INTO snowflakes (id) VALUES (?);

-- name: Get :one
SELECT * FROM snowflakes WHERE id = ? LIMIT 1;
