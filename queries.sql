-- name: GetPlayerByUsername :one
SELECT * FROM player WHERE username = $1;

-- name: GetPlayerByID :one
SELECT * FROM player WHERE id = $1;

-- name: UpdatePlayer :one
UPDATE player 
SET username = $2, password = $3, money = $4, level = $5, exp = $6, gold = $7, token = $8, hour = $9,
    minute = $10, second = $11, vip = $12, helper = $13, is_online = $14, kills = $15, deaths = $16,
    pos_x = $17, pos_y = $18, pos_z = $19, pos_angle = $20, language = $21, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 RETURNING *;

-- name: InsertPlayer :one
INSERT INTO player (username, password) VALUES ($1, $2) RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM player WHERE id = $1;

-- name: UpdateLastLogin :exec
UPDATE player SET last_login = CURRENT_TIMESTAMP WHERE id = $1;

-- name: UpdateLastPlayed :exec
UPDATE player SET last_played = CURRENT_TIMESTAMP WHERE id = $1;
