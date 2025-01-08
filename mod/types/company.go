package types

import (
	"github.com/SanAndreasCW/SACW/mod/database"
)

type CompanyI struct {
	StoreModel   *database.Company
	Applications []*database.CompanyApplication
	Members      []*PlayerI
}

type UserCompanyI struct {
	*CompanyI
	Valid bool
}
