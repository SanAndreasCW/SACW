package types

import (
	"github.com/LosantosGW/go_LSGW/mod/database"
	"github.com/RahRow/omp"
)

type PlayerI struct {
	*omp.Player
	StoreModel *database.Player
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
}
