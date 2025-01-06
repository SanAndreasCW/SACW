package guild

import (
	"context"
	"github.com/LosantosGW/go_LSGW/mod/auth"
	"github.com/LosantosGW/go_LSGW/mod/database"
	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/LosantosGW/go_LSGW/mod/setting"
	"github.com/LosantosGW/go_LSGW/mod/types"
	"github.com/RahRow/omp"
)

var (
	Guilds = make(map[int32]*types.GuildI, setting.MaxGuilds)
)

func onGameModeInit(_ *omp.GameModeInitEvent) bool {
	ctx := context.Background()
	q := database.New(database.DB)
	guilds, err := q.GetGuilds(ctx)

	if err != nil {
		logger.Fatal("[Guild]: Failed to load guilds: %v", err)
		omp.SendRCONCommand("exit")
		return true
	}

	for _, guild := range guilds {
		Guilds[guild.ID] = &types.GuildI{StoreModel: &guild}
	}
	return true
}

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Commands.Add("guild", func(e *omp.Command) {
		ctx := context.Background()
		q := database.New(database.DB)
		player := auth.PlayersI[e.Sender.ID()]
		if !(player.StoreModel.Money > setting.GuildCreationPrice) {
			player.SendClientMessage("Not enough money", 1)
			return
		}
		// @TODO: ask for 3 dialogs: name, tag, color
	})
}
