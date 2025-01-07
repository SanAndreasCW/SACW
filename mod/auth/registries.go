package auth

import (
	"github.com/LosantosGW/go_LSGW/mod/setting"
	"github.com/LosantosGW/go_LSGW/mod/types"
)

var (
	PlayersI     = make(map[int]*types.PlayerI, setting.MaxPlayers)
	PlayersCache = make(map[int]*types.PlayerCache, setting.MaxPlayers)
)
