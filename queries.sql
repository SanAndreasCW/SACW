-- name: GetPlayerByUsername :one
SELECT *
FROM player
WHERE username = $1
LIMIT 1;

-- name: GetPlayerByID :one
SELECT *
FROM player
WHERE id = $1
LIMIT 1;

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

-- name: GetUserActiveCompany :one
SELECT sqlc.embed(company), sqlc.embed(company_member)
FROM company
         JOIN company_member ON (company.id = company_member.company_id)
WHERE company_member.player_id = $1
LIMIT 1;

-- name: GetUserCompaniesInfo :many
SELECT company.*
FROM company
         JOIN company_member_info ON (company.id = company_member_info.company_id)
WHERE company_member_info.player_id = $1;

-- name: GetCompanies :many
SELECT *
FROM company;

-- name: GetCompanyByID :one
SELECT *
FROM company
WHERE id = $1
LIMIT 1;

-- name: GetCompanyByTag :one
SELECT *
FROM company
WHERE tag = $1
LIMIT 1;

-- name: GetCompanyMembers :many
SELECT *
FROM company_member
WHERE company_id = $1;

-- name: InsertCompanyMembers :one
INSERT INTO company_member (player_id, company_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetCompanyMembersInfo :many
SELECT *
FROM company_member_info
WHERE company_id = $1;

-- name: GetCompanyApplications :many
SELECT sqlc.embed(company_application), sqlc.embed(player)
FROM company_application
         JOIN player ON (company_application.player_id = player.id)
WHERE company_application.company_id = $1
  AND company_application.expired_at >= CURRENT_TIMESTAMP
  AND company_application.accepted = $2
  AND NOT EXISTS (SELECT 1
                  FROM company_member
                  WHERE company_member.player_id = player.id);

-- name: GetUserCompanyApplicationsHistory :many
SELECT sqlc.embed(company_application), sqlc.embed(player), sqlc.embed(company)
FROM company_application
         JOIN player ON (company_application.player_id = player.id)
         JOIN company ON (company_application.company_id = company.id)
WHERE company_application.company_id = $1
  AND company_application.player_id = $2
  AND company_application.expired_at >= CURRENT_TIMESTAMP
  AND company_application.accepted != 0
LIMIT 4;

-- name: GetCompaniesApplications :many
SELECT *
FROM company_application
WHERE expired_at <= CURRENT_TIMESTAMP;

-- name: AnswerCompanyApplication :exec
UPDATE company_application
SET accepted    = $1,
    answered_at = CURRENT_TIMESTAMP
WHERE player_id = $2
  AND company_id = $3
  AND accepted = 0;

-- name: InsertCompanyApplication :one
INSERT INTO company_application (player_id, company_id, description)
SELECT $1, $2, $3
WHERE NOT EXISTS (
    SELECT 1 FROM company_member
    WHERE player_id = $1
)
RETURNING *;
