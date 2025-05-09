package company

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/cmd"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/kodeyeen/omp"
)

func companiesCommand(cmd *cmd.Command) {
	ctx := context.Background()
	playerI := commons.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, parameters := cmd.Args[0], cmd.Args[1:]
		paramsLength := len(parameters)
		switch action {
		case "application":
			if paramsLength >= 1 {
				tag := parameters[0]
				companiesApplicationAction(playerI, tag)
				return
			}
		case "history":
			companiesHistoryAction(playerI)
			return

		case "applications":
			companyApplicationsActions(ctx, playerI)
			return
		}
		playerI.SendClientMessage("[Command Guide]: /companies :optional[actions] :optional[data...]", colors.InfoColor.Hex)
		return
	}
	if len(commons.Companies) <= 0 {
		msgDialog := omp.NewMessageDialog("Companies List", "No Companies Defined Yet!", "Ok", "Close")
		msgDialog.ShowFor(playerI.Player)
		return
	}
	companiesDialog := omp.NewTabListDialog("Companies List", "Ok", "Close")
	for _, company := range commons.Companies {
		companiesDialog.Add(omp.TabListItem{
			company.StoreModel.Name,
			company.StoreModel.Tag,
			commons.FloatToString(company.StoreModel.Multiplier),
		})
	}
	companiesDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.TabListDialogResponseEvent)
		if ep.Response == omp.DialogResponseRight {
			return nil
		}
		companiesApplicationAction(playerI, ep.Item[1])
		return nil
	})
	companiesDialog.ShowFor(playerI.Player)
}

func companyCommand(cmd *cmd.Command) {
	ctx := context.Background()
	playerI := commons.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, _ := cmd.Args[0], cmd.Args[1:]
		switch action {
		case "applications":
			companyApplicationsActions(ctx, playerI)
			return
		case "stats":
			playerCompanyStats(ctx, playerI)
			return
		}
	}
}
