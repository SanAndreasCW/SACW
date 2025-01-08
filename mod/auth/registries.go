package auth

import (
	"github.com/SanAndreasCW/SACW/mod/setting"
	"github.com/SanAndreasCW/SACW/mod/types"
)

var (
	PlayersI     = make(map[int]*types.PlayerI, setting.MaxPlayers)
	PlayersCache = make(map[int]*types.PlayerCache, setting.MaxPlayers)
)
