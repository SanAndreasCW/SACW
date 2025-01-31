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
    vip        = $11,
    helper     = $12,
    is_online  = $13,
    kills      = $14,
    deaths     = $15,
    pos_x      = $16,
    pos_y      = $17,
    pos_z      = $18,
    pos_angle  = $19,
    language   = $20,
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

-- name: UpdateCompanyMemberInfo :exec
UPDATE company_member_info
SET level  = $3,
    hour   = $4,
    minute = $5,
    score  = $6
WHERE company_id = $1
  AND player_id = $2;

-- name: GetUserActiveCompany :one
SELECT sqlc.embed(company), sqlc.embed(company_member), sqlc.embed(company_member_info)
FROM company
         JOIN company_member ON (company.id = company_member.company_id)
         JOIN company_member_info ON (company.id = company_member_info.company_id)
WHERE company_member.player_id = $1
LIMIT 1;

-- name: GetUserCompaniesInfo :many
SELECT company.*
FROM company
         JOIN company_member_info ON (company.id = company_member_info.company_id)
WHERE company_member_info.player_id = $1;

-- name: GetCompanies :many
SELECT sqlc.embed(company), sqlc.embed(company_office)
FROM company
         JOIN company_office ON (company.id = company_office.company_id);

-- name: GetCompanyJobs :many
SELECT *
FROM company_job
WHERE company_id = $1;

-- name: UpdateCompany :exec
UPDATE company
SET balance    = $1,
    multiplier = $2
WHERE id = $3;

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
WHERE NOT EXISTS (SELECT 1
                  FROM company_member
                  WHERE player_id = $1)
RETURNING *;


-- name: GetUserJobs :many
SELECT *
FROM player_job
WHERE player_id = $1;

-- name: UpdateUserJobs :one
SELECT sqlc.embed(player_job)
FROM update_or_create_player_job($1, $2, $3) as player_job;
