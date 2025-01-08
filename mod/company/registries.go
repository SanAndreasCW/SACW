package company

import (
	"github.com/SanAndreasCW/SACW/mod/setting"
	"github.com/SanAndreasCW/SACW/mod/types"
)

var (
	Companies = make(map[int]*types.CompanyI, setting.MaxCompanies)
)
