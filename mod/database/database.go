package database

import (
	"database/sql"
	_ "database/sql/driver"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {}
