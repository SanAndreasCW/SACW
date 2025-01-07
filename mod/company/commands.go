package company

import (
	"context"
	"database/sql"
	"github.com/LosantosGW/go_LSGW/mod/auth"
	"github.com/LosantosGW/go_LSGW/mod/database"
	"github.com/RahRow/omp"
)

func companyCommand(cmd *omp.Command) {
	player := cmd.Sender
	playerI := auth.PlayersI[cmd.Sender.ID()]
	if argsCount := len(cmd.Args); argsCount > 0 {
		action, parameters := cmd.Args[0], cmd.Args[1:]
		paramsLength := len(parameters)

		switch action {
		case "application":
			if paramsLength >= 1 {
				tag := parameters[0]

				for _, company := range Companies {
					if company.StoreModel.Tag == tag {
						ctx := context.Background()
						q := database.New(database.DB)
						dialog := omp.NewInputDialog(
							"Company Application",
							"Enter company application request description.",
							"Send",
							"Close",
						)

						dialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
							if e.Response == omp.DialogResponseRight {
								return true
							}
							selectedCompany, _ := q.GetCompanyByTag(ctx, tag)
							_, _ = q.InsertCompanyApplication(ctx, database.InsertCompanyApplicationParams{
								PlayerID:    playerI.StoreModel.ID,
								CompanyID:   selectedCompany.ID,
								Description: sql.NullString{String: e.InputText, Valid: true},
							})
							playerI.SendClientMessage(
								"[Company]: Your application was sent successfully.",
								1,
							)
							return true
						})
						dialog.ShowFor(player)
						break
					}
				}
				return
			}
		}
		playerI.SendClientMessage("[Command Guide]: /companies :optional[actions] :optional[data...]", 1)
		return
	}

	namesList := ""
	if len(Companies) <= 0 {
		namesList = "No Companies Defined Yet!"
	}
	for _, company := range Companies {
		namesList += company.StoreModel.Name + ", "
	}
	dialog := omp.NewMessageDialog("Companies List", namesList, "Ok", "Close")
	dialog.On(omp.EventTypeDialogResponse, func(_ *omp.MessageDialogResponseEvent) bool {
		return true
	})
	dialog.ShowFor(player)
}
