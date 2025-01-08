package types

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
)

type PlayerI struct {
	*omp.Player
	StoreModel *database.Player
	Company    *CompanyI
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
}
