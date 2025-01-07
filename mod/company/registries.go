package company

import (
	"github.com/LosantosGW/go_LSGW/mod/setting"
	"github.com/LosantosGW/go_LSGW/mod/types"
)

var (
	Companies = make(map[int]*types.CompanyI, setting.MaxCompanies)
)
