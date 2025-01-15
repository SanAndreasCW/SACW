package commons

import (
	"github.com/SanAndreasCW/SACW/mod/setting"
)

var (
	Companies = make(map[int32]*CompanyI, setting.MaxCompanies)
)
