package company

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/database"
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
						dialogBody := fmt.Sprintf(
							"Enter company application request description for %s company.",
							company.StoreModel.Name,
						)
						dialog := omp.NewInputDialog(
							"Company Application",
							dialogBody,
							"Send",
							"Close",
						)

						dialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
							if e.Response == omp.DialogResponseRight {
								return true
							}
							if len(e.InputText) > 80 {
								dialog.Body = fmt.Sprintf("%s\nDescription can't be larger than 80 characters.", dialogBody)
								dialog.ShowFor(player)
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
