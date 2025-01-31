package database

import "context"

const insertCompanyMembers = `-- name: InsertCompanyMembers :one
SELECT company_member, company_member_info FROM update_or_create_player_membership($1, $2);`

type InsertCompanyMembersParams struct {
	PlayerID  int32
	CompanyID int32
}

type InsertCompanyMembersRow struct {
	CompanyMember     CompanyMember
	CompanyMemberInfo CompanyMemberInfo
}

func (q *Queries) InsertCompanyMembers(ctx context.Context, arg InsertCompanyMembersParams) (InsertCompanyMembersRow, error) {
	row := q.db.QueryRowContext(ctx, insertCompanyMembers, arg.PlayerID, arg.CompanyID)
	var i InsertCompanyMembersRow
	err := row.Scan(
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
