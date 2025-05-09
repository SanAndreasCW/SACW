// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const answerCompanyApplication = `-- name: AnswerCompanyApplication :exec
UPDATE company_application
SET accepted    = $1,
    answered_at = CURRENT_TIMESTAMP
WHERE player_id = $2
  AND company_id = $3
  AND accepted = 0
`

type AnswerCompanyApplicationParams struct {
	Accepted  int16
	PlayerID  int32
	CompanyID int32
}

func (q *Queries) AnswerCompanyApplication(ctx context.Context, arg AnswerCompanyApplicationParams) error {
	_, err := q.db.ExecContext(ctx, answerCompanyApplication, arg.Accepted, arg.PlayerID, arg.CompanyID)
	return err
}

const deletePlayer = `-- name: DeletePlayer :exec
DELETE
FROM player
WHERE id = $1
`

func (q *Queries) DeletePlayer(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePlayer, id)
	return err
}

const getCompanies = `-- name: GetCompanies :many
SELECT company.id, company.name, company.tag, company.description, company.balance, company.multiplier, company_office.id, company_office.company_id, company_office.icon_x, company_office.icon_y, company_office.pickup_x, company_office.pickup_y, company_office.pickup_z
FROM company
         JOIN company_office ON (company.id = company_office.company_id)
`

type GetCompaniesRow struct {
	Company       Company
	CompanyOffice CompanyOffice
}

func (q *Queries) GetCompanies(ctx context.Context) ([]GetCompaniesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompaniesRow
	for rows.Next() {
		var i GetCompaniesRow
		if err := rows.Scan(
			&i.Company.ID,
			&i.Company.Name,
			&i.Company.Tag,
			&i.Company.Description,
			&i.Company.Balance,
			&i.Company.Multiplier,
			&i.CompanyOffice.ID,
			&i.CompanyOffice.CompanyID,
			&i.CompanyOffice.IconX,
			&i.CompanyOffice.IconY,
			&i.CompanyOffice.PickupX,
			&i.CompanyOffice.PickupY,
			&i.CompanyOffice.PickupZ,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompaniesApplications = `-- name: GetCompaniesApplications :many
SELECT id, player_id, company_id, description, accepted, created_at, expired_at, answer, answered_at
FROM company_application
WHERE expired_at <= CURRENT_TIMESTAMP
`

func (q *Queries) GetCompaniesApplications(ctx context.Context) ([]CompanyApplication, error) {
	rows, err := q.db.QueryContext(ctx, getCompaniesApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyApplication
	for rows.Next() {
		var i CompanyApplication
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.CompanyID,
			&i.Description,
			&i.Accepted,
			&i.CreatedAt,
			&i.ExpiredAt,
			&i.Answer,
			&i.AnsweredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyApplications = `-- name: GetCompanyApplications :many
SELECT company_application.id, company_application.player_id, company_application.company_id, company_application.description, company_application.accepted, company_application.created_at, company_application.expired_at, company_application.answer, company_application.answered_at, player.id, player.username, player.password, player.money, player.level, player.exp, player.gold, player.token, player.hour, player.minute, player.vip, player.helper, player.is_online, player.kills, player.deaths, player.pos_x, player.pos_y, player.pos_z, player.pos_angle, player.language, player.last_login, player.last_played, player.created_at, player.updated_at
FROM company_application
         JOIN player ON (company_application.player_id = player.id)
WHERE company_application.company_id = $1
  AND company_application.expired_at >= CURRENT_TIMESTAMP
  AND company_application.accepted = $2
  AND NOT EXISTS (SELECT 1
                  FROM company_member
                  WHERE company_member.player_id = player.id)
`

type GetCompanyApplicationsParams struct {
	CompanyID int32
	Accepted  int16
}

type GetCompanyApplicationsRow struct {
	CompanyApplication CompanyApplication
	Player             Player
}

func (q *Queries) GetCompanyApplications(ctx context.Context, arg GetCompanyApplicationsParams) ([]GetCompanyApplicationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyApplications, arg.CompanyID, arg.Accepted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompanyApplicationsRow
	for rows.Next() {
		var i GetCompanyApplicationsRow
		if err := rows.Scan(
			&i.CompanyApplication.ID,
			&i.CompanyApplication.PlayerID,
			&i.CompanyApplication.CompanyID,
			&i.CompanyApplication.Description,
			&i.CompanyApplication.Accepted,
			&i.CompanyApplication.CreatedAt,
			&i.CompanyApplication.ExpiredAt,
			&i.CompanyApplication.Answer,
			&i.CompanyApplication.AnsweredAt,
			&i.Player.ID,
			&i.Player.Username,
			&i.Player.Password,
			&i.Player.Money,
			&i.Player.Level,
			&i.Player.Exp,
			&i.Player.Gold,
			&i.Player.Token,
			&i.Player.Hour,
			&i.Player.Minute,
			&i.Player.Vip,
			&i.Player.Helper,
			&i.Player.IsOnline,
			&i.Player.Kills,
			&i.Player.Deaths,
			&i.Player.PosX,
			&i.Player.PosY,
			&i.Player.PosZ,
			&i.Player.PosAngle,
			&i.Player.Language,
			&i.Player.LastLogin,
			&i.Player.LastPlayed,
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyByID = `-- name: GetCompanyByID :one
SELECT id, name, tag, description, balance, multiplier
FROM company
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCompanyByID(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByID, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tag,
		&i.Description,
		&i.Balance,
		&i.Multiplier,
	)
	return i, err
}

const getCompanyByTag = `-- name: GetCompanyByTag :one
SELECT id, name, tag, description, balance, multiplier
FROM company
WHERE tag = $1
LIMIT 1
`

func (q *Queries) GetCompanyByTag(ctx context.Context, tag string) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByTag, tag)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tag,
		&i.Description,
		&i.Balance,
		&i.Multiplier,
	)
	return i, err
}

const getCompanyJobs = `-- name: GetCompanyJobs :many
SELECT id, company_id, job_id, job_group
FROM company_job
WHERE company_id = $1
`

func (q *Queries) GetCompanyJobs(ctx context.Context, companyID int32) ([]CompanyJob, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyJobs, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyJob
	for rows.Next() {
		var i CompanyJob
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.JobID,
			&i.JobGroup,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyJobsCheckpoint = `-- name: GetCompanyJobsCheckpoint :many
SELECT id, company_id, job_id, type, x, y, z
FROM company_job_checkpoint
WHERE company_id = $1
`

func (q *Queries) GetCompanyJobsCheckpoint(ctx context.Context, companyID int32) ([]CompanyJobCheckpoint, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyJobsCheckpoint, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyJobCheckpoint
	for rows.Next() {
		var i CompanyJobCheckpoint
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.JobID,
			&i.Type,
			&i.X,
			&i.Y,
			&i.Z,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyMembers = `-- name: GetCompanyMembers :many
SELECT id, player_id, company_id, role
FROM company_member
WHERE company_id = $1
`

func (q *Queries) GetCompanyMembers(ctx context.Context, companyID int32) ([]CompanyMember, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyMembers, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyMember
	for rows.Next() {
		var i CompanyMember
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.CompanyID,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyMembersInfo = `-- name: GetCompanyMembersInfo :many
SELECT id, player_id, company_id, hour, minute, score, level
FROM company_member_info
WHERE company_id = $1
`

func (q *Queries) GetCompanyMembersInfo(ctx context.Context, companyID int32) ([]CompanyMemberInfo, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyMembersInfo, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyMemberInfo
	for rows.Next() {
		var i CompanyMemberInfo
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.CompanyID,
			&i.Hour,
			&i.Minute,
			&i.Score,
			&i.Level,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayerByID = `-- name: GetPlayerByID :one
SELECT id, username, password, money, level, exp, gold, token, hour, minute, vip, helper, is_online, kills, deaths, pos_x, pos_y, pos_z, pos_angle, language, last_login, last_played, created_at, updated_at
FROM player
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPlayerByID(ctx context.Context, id int32) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Money,
		&i.Level,
		&i.Exp,
		&i.Gold,
		&i.Token,
		&i.Hour,
		&i.Minute,
		&i.Vip,
		&i.Helper,
		&i.IsOnline,
		&i.Kills,
		&i.Deaths,
		&i.PosX,
		&i.PosY,
		&i.PosZ,
		&i.PosAngle,
		&i.Language,
		&i.LastLogin,
		&i.LastPlayed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlayerByUsername = `-- name: GetPlayerByUsername :one
SELECT id, username, password, money, level, exp, gold, token, hour, minute, vip, helper, is_online, kills, deaths, pos_x, pos_y, pos_z, pos_angle, language, last_login, last_played, created_at, updated_at
FROM player
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetPlayerByUsername(ctx context.Context, username string) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerByUsername, username)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Money,
		&i.Level,
		&i.Exp,
		&i.Gold,
		&i.Token,
		&i.Hour,
		&i.Minute,
		&i.Vip,
		&i.Helper,
		&i.IsOnline,
		&i.Kills,
		&i.Deaths,
		&i.PosX,
		&i.PosY,
		&i.PosZ,
		&i.PosAngle,
		&i.Language,
		&i.LastLogin,
		&i.LastPlayed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserActiveCompany = `-- name: GetUserActiveCompany :one
SELECT company.id, company.name, company.tag, company.description, company.balance, company.multiplier, company_member.id, company_member.player_id, company_member.company_id, company_member.role, company_member_info.id, company_member_info.player_id, company_member_info.company_id, company_member_info.hour, company_member_info.minute, company_member_info.score, company_member_info.level
FROM company
         JOIN company_member ON (company.id = company_member.company_id)
         JOIN company_member_info ON (company.id = company_member_info.company_id)
WHERE company_member.player_id = $1
LIMIT 1
`

type GetUserActiveCompanyRow struct {
	Company           Company
	CompanyMember     CompanyMember
	CompanyMemberInfo CompanyMemberInfo
}

func (q *Queries) GetUserActiveCompany(ctx context.Context, playerID int32) (GetUserActiveCompanyRow, error) {
	row := q.db.QueryRowContext(ctx, getUserActiveCompany, playerID)
	var i GetUserActiveCompanyRow
	err := row.Scan(
		&i.Company.ID,
		&i.Company.Name,
		&i.Company.Tag,
		&i.Company.Description,
		&i.Company.Balance,
		&i.Company.Multiplier,
		&i.CompanyMember.ID,
		&i.CompanyMember.PlayerID,
		&i.CompanyMember.CompanyID,
		&i.CompanyMember.Role,
		&i.CompanyMemberInfo.ID,
		&i.CompanyMemberInfo.PlayerID,
		&i.CompanyMemberInfo.CompanyID,
		&i.CompanyMemberInfo.Hour,
		&i.CompanyMemberInfo.Minute,
		&i.CompanyMemberInfo.Score,
		&i.CompanyMemberInfo.Level,
	)
	return i, err
}

const getUserCompaniesInfo = `-- name: GetUserCompaniesInfo :many
SELECT company.id, company.name, company.tag, company.description, company.balance, company.multiplier
FROM company
         JOIN company_member_info ON (company.id = company_member_info.company_id)
WHERE company_member_info.player_id = $1
`

func (q *Queries) GetUserCompaniesInfo(ctx context.Context, playerID int32) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompaniesInfo, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Tag,
			&i.Description,
			&i.Balance,
			&i.Multiplier,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserCompanyApplicationsHistory = `-- name: GetUserCompanyApplicationsHistory :many
SELECT company_application.id, company_application.player_id, company_application.company_id, company_application.description, company_application.accepted, company_application.created_at, company_application.expired_at, company_application.answer, company_application.answered_at, player.id, player.username, player.password, player.money, player.level, player.exp, player.gold, player.token, player.hour, player.minute, player.vip, player.helper, player.is_online, player.kills, player.deaths, player.pos_x, player.pos_y, player.pos_z, player.pos_angle, player.language, player.last_login, player.last_played, player.created_at, player.updated_at, company.id, company.name, company.tag, company.description, company.balance, company.multiplier
FROM company_application
         JOIN player ON (company_application.player_id = player.id)
         JOIN company ON (company_application.company_id = company.id)
WHERE company_application.company_id = $1
  AND company_application.player_id = $2
  AND company_application.expired_at >= CURRENT_TIMESTAMP
  AND company_application.accepted != 0
LIMIT 4
`

type GetUserCompanyApplicationsHistoryParams struct {
	CompanyID int32
	PlayerID  int32
}

type GetUserCompanyApplicationsHistoryRow struct {
	CompanyApplication CompanyApplication
	Player             Player
	Company            Company
}

func (q *Queries) GetUserCompanyApplicationsHistory(ctx context.Context, arg GetUserCompanyApplicationsHistoryParams) ([]GetUserCompanyApplicationsHistoryRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompanyApplicationsHistory, arg.CompanyID, arg.PlayerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCompanyApplicationsHistoryRow
	for rows.Next() {
		var i GetUserCompanyApplicationsHistoryRow
		if err := rows.Scan(
			&i.CompanyApplication.ID,
			&i.CompanyApplication.PlayerID,
			&i.CompanyApplication.CompanyID,
			&i.CompanyApplication.Description,
			&i.CompanyApplication.Accepted,
			&i.CompanyApplication.CreatedAt,
			&i.CompanyApplication.ExpiredAt,
			&i.CompanyApplication.Answer,
			&i.CompanyApplication.AnsweredAt,
			&i.Player.ID,
			&i.Player.Username,
			&i.Player.Password,
			&i.Player.Money,
			&i.Player.Level,
			&i.Player.Exp,
			&i.Player.Gold,
			&i.Player.Token,
			&i.Player.Hour,
			&i.Player.Minute,
			&i.Player.Vip,
			&i.Player.Helper,
			&i.Player.IsOnline,
			&i.Player.Kills,
			&i.Player.Deaths,
			&i.Player.PosX,
			&i.Player.PosY,
			&i.Player.PosZ,
			&i.Player.PosAngle,
			&i.Player.Language,
			&i.Player.LastLogin,
			&i.Player.LastPlayed,
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
			&i.Company.ID,
			&i.Company.Name,
			&i.Company.Tag,
			&i.Company.Description,
			&i.Company.Balance,
			&i.Company.Multiplier,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserJobs = `-- name: GetUserJobs :many
SELECT id, player_id, job_id, score
FROM player_job
WHERE player_id = $1
`

func (q *Queries) GetUserJobs(ctx context.Context, playerID int32) ([]PlayerJob, error) {
	rows, err := q.db.QueryContext(ctx, getUserJobs, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlayerJob
	for rows.Next() {
		var i PlayerJob
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.JobID,
			&i.Score,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCompanyApplication = `-- name: InsertCompanyApplication :one
INSERT INTO company_application (player_id, company_id, description)
SELECT $1, $2, $3
WHERE NOT EXISTS (SELECT 1
                  FROM company_member
                  WHERE player_id = $1)
RETURNING id, player_id, company_id, description, accepted, created_at, expired_at, answer, answered_at
`

type InsertCompanyApplicationParams struct {
	PlayerID    int32
	CompanyID   int32
	Description sql.NullString
}

func (q *Queries) InsertCompanyApplication(ctx context.Context, arg InsertCompanyApplicationParams) (CompanyApplication, error) {
	row := q.db.QueryRowContext(ctx, insertCompanyApplication, arg.PlayerID, arg.CompanyID, arg.Description)
	var i CompanyApplication
	err := row.Scan(
		&i.ID,
		&i.PlayerID,
		&i.CompanyID,
		&i.Description,
		&i.Accepted,
		&i.CreatedAt,
		&i.ExpiredAt,
		&i.Answer,
		&i.AnsweredAt,
	)
	return i, err
}

const insertPlayer = `-- name: InsertPlayer :one
INSERT INTO player (username, password)
VALUES ($1, $2)
RETURNING id, username, password, money, level, exp, gold, token, hour, minute, vip, helper, is_online, kills, deaths, pos_x, pos_y, pos_z, pos_angle, language, last_login, last_played, created_at, updated_at
`

type InsertPlayerParams struct {
	Username string
	Password string
}

func (q *Queries) InsertPlayer(ctx context.Context, arg InsertPlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, insertPlayer, arg.Username, arg.Password)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Money,
		&i.Level,
		&i.Exp,
		&i.Gold,
		&i.Token,
		&i.Hour,
		&i.Minute,
		&i.Vip,
		&i.Helper,
		&i.IsOnline,
		&i.Kills,
		&i.Deaths,
		&i.PosX,
		&i.PosY,
		&i.PosZ,
		&i.PosAngle,
		&i.Language,
		&i.LastLogin,
		&i.LastPlayed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :exec
UPDATE company
SET balance    = $1,
    multiplier = $2
WHERE id = $3
`

type UpdateCompanyParams struct {
	Balance    int32
	Multiplier float32
	ID         int32
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error {
	_, err := q.db.ExecContext(ctx, updateCompany, arg.Balance, arg.Multiplier, arg.ID)
	return err
}

const updateCompanyMemberInfo = `-- name: UpdateCompanyMemberInfo :exec
UPDATE company_member_info
SET level  = $3,
    hour   = $4,
    minute = $5,
    score  = $6
WHERE company_id = $1
  AND player_id = $2
`

type UpdateCompanyMemberInfoParams struct {
	CompanyID int32
	PlayerID  int32
	Level     int32
	Hour      int32
	Minute    int16
	Score     int32
}

func (q *Queries) UpdateCompanyMemberInfo(ctx context.Context, arg UpdateCompanyMemberInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateCompanyMemberInfo,
		arg.CompanyID,
		arg.PlayerID,
		arg.Level,
		arg.Hour,
		arg.Minute,
		arg.Score,
	)
	return err
}

const updateLastLogin = `-- name: UpdateLastLogin :exec
UPDATE player
SET last_login = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) UpdateLastLogin(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, updateLastLogin, id)
	return err
}

const updateLastPlayed = `-- name: UpdateLastPlayed :exec
UPDATE player
SET last_played = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) UpdateLastPlayed(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, updateLastPlayed, id)
	return err
}

const updatePlayer = `-- name: UpdatePlayer :one
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
RETURNING id, username, password, money, level, exp, gold, token, hour, minute, vip, helper, is_online, kills, deaths, pos_x, pos_y, pos_z, pos_angle, language, last_login, last_played, created_at, updated_at
`

type UpdatePlayerParams struct {
	ID       int32
	Username string
	Password string
	Money    int32
	Level    int32
	Exp      int32
	Gold     int32
	Token    int32
	Hour     int32
	Minute   int32
	Vip      int32
	Helper   int32
	IsOnline bool
	Kills    int32
	Deaths   int32
	PosX     float32
	PosY     float32
	PosZ     float32
	PosAngle float32
	Language int32
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, updatePlayer,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Money,
		arg.Level,
		arg.Exp,
		arg.Gold,
		arg.Token,
		arg.Hour,
		arg.Minute,
		arg.Vip,
		arg.Helper,
		arg.IsOnline,
		arg.Kills,
		arg.Deaths,
		arg.PosX,
		arg.PosY,
		arg.PosZ,
		arg.PosAngle,
		arg.Language,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Money,
		&i.Level,
		&i.Exp,
		&i.Gold,
		&i.Token,
		&i.Hour,
		&i.Minute,
		&i.Vip,
		&i.Helper,
		&i.IsOnline,
		&i.Kills,
		&i.Deaths,
		&i.PosX,
		&i.PosY,
		&i.PosZ,
		&i.PosAngle,
		&i.Language,
		&i.LastLogin,
		&i.LastPlayed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserJobs = `-- name: UpdateUserJobs :one
SELECT player_job.id, player_job.player_id, player_job.job_id, player_job.score
FROM update_or_create_player_job($1, $2, $3) as player_job
`

type UpdateUserJobsParams struct {
	Pid int32
	Jid int32
	Sc  int32
}

type UpdateUserJobsRow struct {
	PlayerJob PlayerJob
}

func (q *Queries) UpdateUserJobs(ctx context.Context, arg UpdateUserJobsParams) (UpdateUserJobsRow, error) {
	row := q.db.QueryRowContext(ctx, updateUserJobs, arg.Pid, arg.Jid, arg.Sc)
	var i UpdateUserJobsRow
	err := row.Scan(
		&i.PlayerJob.ID,
		&i.PlayerJob.PlayerID,
		&i.PlayerJob.JobID,
		&i.PlayerJob.Score,
	)
	return i, err
}
