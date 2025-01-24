package commons

import (
	"github.com/SanAndreasCW/SACW/mod/setting"
)

var (
	PlayersI = make(map[int]*PlayerI, setting.MaxPlayers)
)
