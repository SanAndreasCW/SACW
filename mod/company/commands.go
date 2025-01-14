package company

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
)

func companyCommand(cmd *omp.Command) {
	playerI := auth.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, parameters := cmd.Args[0], cmd.Args[1:]
		paramsLength := len(parameters)

		switch action {
		case "application":
			if paramsLength >= 1 {
				tag := parameters[0]
				companyApplicationAction(playerI, &tag)
				return
			}
		case "history":
			companyHistoryAction(playerI)
			return
		}
		playerI.SendClientMessage("[Command Guide]: /companies :optional[actions] :optional[data...]", 1)
		return
	}
	if len(Companies) <= 0 {
		msgDialog := omp.NewMessageDialog("Companies List", "No Companies Defined Yet!", "Ok", "Close")
		msgDialog.ShowFor(playerI.Player)
		return
	}
	companiesDialog := omp.NewTabListDialog("Companies List", "Ok", "Close")
	for _, company := range Companies {
		companiesDialog.Add(omp.TabListItem{
			company.StoreModel.Name,
			company.StoreModel.Tag,
			string(rune(company.StoreModel.Multiplier.Float64)),
		})
	}
	companiesDialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
		test := e.Item
		companyApplicationAction(playerI, &test[1])
		return true
	})
	companiesDialog.ShowFor(playerI.Player)
}
