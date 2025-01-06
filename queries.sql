-- name: GetPlayerByUsername :one
SELECT *
FROM player
WHERE username = $1;

-- name: GetPlayerByID :one
SELECT *
FROM player
WHERE id = $1;

-- name: UpdatePlayer :one
UPDATE player
SET username   = $2,
    password   = $3,
    money      = $4,
    level      = $5,
    exp        = $6,
    gold       = $7,
    token      = $8,
    hour       = $9,
    minute     = $10,
    second     = $11,
    vip        = $12,
    helper     = $13,
    is_online  = $14,
    kills      = $15,
    deaths     = $16,
    pos_x      = $17,
    pos_y      = $18,
    pos_z      = $19,
    pos_angle  = $20,
    language   = $21,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: InsertPlayer :one
INSERT INTO player (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePlayer :exec
DELETE
FROM player
WHERE id = $1;

-- name: UpdateLastLogin :exec
UPDATE player
SET last_login = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: UpdateLastPlayed :exec
UPDATE player
SET last_played = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetGuilds :many
SELECT *
FROM guild;

-- name: GetGuildByID :one
SELECT *
FROM guild
WHERE id = $1;

-- name: GetGuildByTag :one
SELECT *
FROM guild
where tag = $1;

-- name: InsertGuild :one
INSERT INTO guild (name, tag, leader_id, testimonial, color)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateGuild :one
UPDATE guild
SET testimonial = $1,
    color       = $2,
    level       = $3,
    exp         = $4,
    points      = $5,
    favors      = $6,
    is_active   = $7,
    leader_id   = $8,
    updated_at  = CURRENT_TIMESTAMP
WHERE id = $9
RETURNING *;

-- name: DeleteGuildByID :exec
DELETE
FROM guild
WHERE id = $1;