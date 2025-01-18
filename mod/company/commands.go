package company

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/commons"
)

func companiesCommand(cmd *omp.Command) {
	playerI := commons.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, parameters := cmd.Args[0], cmd.Args[1:]
		paramsLength := len(parameters)
		switch action {
		case "application":
			if paramsLength >= 1 {
				tag := parameters[0]
				companiesApplicationAction(playerI, &tag)
				return
			}
		case "history":
			companiesHistoryAction(playerI)
			return

		case "applications":
			companyApplicationsActions(playerI)
			return
		}
		playerI.SendClientMessage("[Command Guide]: /companies :optional[actions] :optional[data...]", 1)
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
	companiesDialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
		test := e.Item
		companiesApplicationAction(playerI, &test[1])
		return true
	})
	companiesDialog.ShowFor(playerI.Player)
}

func companyCommand(cmd *omp.Command) {
	playerI := commons.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, _ := cmd.Args[0], cmd.Args[1:]
		switch action {
		case "applications":
			companyApplicationsActions(playerI)
			return
		}
	}
}
