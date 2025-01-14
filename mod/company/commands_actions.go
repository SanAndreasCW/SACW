package company

import (
	"fmt"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/setting"
	"github.com/SanAndreasCW/SACW/mod/types"
)

func companyApplicationAction(playerI *types.PlayerI, tag *string) {
	for _, company := range Companies {
		if company.StoreModel.Tag == *tag {
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
					dialog.ShowFor(playerI.Player)
					return true
				}
				// Insert Company Application
				if len(company.Applications) >= setting.MaxCompanyApplications {
					playerI.SendClientMessage(
						"[Company Application]: Targeted company is not capable for more applications.",
						1,
					)
					return true
				}
				isCreated := company.CreateApplication(playerI, e.InputText)
				if isCreated != true {
					logger.Fatal("[CompanyApplication]: Failed to save company application")
					playerI.SendClientMessage(
						"[Company Application]: Your application was not sent, please try again later.",
						1,
					)
					return true
				}
				playerI.SendClientMessage(
					"[Company Application]: Your application was sent successfully.",
					1,
				)
				return true
			})
			dialog.ShowFor(playerI.Player)
			return
		}
	}

	playerI.Player.SendClientMessage("[Company Application]: Company tag not found.", 1)
}
