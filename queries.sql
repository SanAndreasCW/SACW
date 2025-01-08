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
SELECT company.*
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

-- name: GetCompanyMembersInfo :many
SELECT *
FROM company_member_info
WHERE company_id = $1;

-- name: GetCompanyApplications :many
SELECT *
FROM company_application
WHERE company_id = $1
  AND expired_at <= CURRENT_TIMESTAMP;

-- name: GetCompaniesApplications :many
SELECT *
FROM company_application
WHERE expired_at <= CURRENT_TIMESTAMP;

-- name: AcceptCompanyApplication :exec
UPDATE company_application
SET accepted    = true,
    answered_at = CURRENT_TIMESTAMP
WHERE player_id = $1
  AND company_id = $2
  AND accepted = false;

-- name: InsertCompanyApplication :one
INSERT INTO company_application (player_id, company_id, description)
VALUES ($1, $2, $3)
RETURNING *;
