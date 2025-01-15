package company

import (
	"fmt"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/setting"
)

func companiesApplicationAction(playerI *commons.PlayerI, tag *string) {
	for _, company := range commons.Companies {
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

func companiesHistoryAction(playerI *commons.PlayerI) {
	if len(playerI.Companies) <= 0 {
		msgDialog := omp.NewMessageDialog(
			"Companies History List",
			"You don't have any company history.",
			"Ok",
			"",
		)
		msgDialog.ShowFor(playerI.Player)
		return
	}
	companyHistory := omp.NewTabListDialog("Companies History List", "Select", "Close")
	for _, company := range playerI.Companies {
		companyI := commons.Companies[company.StoreModel.ID]
		companyHistory.Add(omp.TabListItem{
			companyI.StoreModel.Name,
			companyI.StoreModel.Tag,
			string(rune(companyI.StoreModel.Multiplier.Float64)),
		})
	}
	companyHistory.ShowFor(playerI.Player)
}

func companyApplicationsActions(playerI *commons.PlayerI) {
	company := playerI.GetCurrentCompany()
	if company == nil {
		playerI.SendClientMessage(
			"[Company Applications]: You are not in a company to check incoming company applications",
			1,
		)
		return
	}
	if len(company.Applications) <= 0 {
		dialog := omp.NewMessageDialog(
			"Company Applications",
			"No applications found in the company which you are in",
			"Close",
			"",
		)
		dialog.ShowFor(playerI.Player)
		return
	}
	companyApplications := omp.NewTabListDialog("Company Applications", "Select", "Close")
	companyApplications.Add(omp.TabListItem{
		"Player Name",
		"Requested At",
		"Expire At",
	})
	for _, companyApplication := range company.Applications {
		companyApplications.Add(omp.TabListItem{
			string(companyApplication.PlayerModel.Username),
			companyApplication.StoreModel.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			companyApplication.StoreModel.ExpiredAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	companyApplications.ShowFor(playerI.Player)
}
