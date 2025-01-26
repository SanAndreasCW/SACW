package company

import (
	"context"
	"fmt"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/setting"
)

func companiesApplicationAction(playerI *commons.PlayerI, tag *string) {
	for _, company := range commons.Companies {
		if company.StoreModel.Tag == *tag {
			if playerI.Membership != nil {
				companyStatsDialog(company).ShowFor(playerI.Player)
				return
			}
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
						colors.NoteHex,
					)
					return true
				}
				isCreated := company.CreateApplication(playerI, e.InputText)
				if isCreated != true {
					logger.Fatal("[CompanyApplication]: Failed to save company application")
					playerI.SendClientMessage(
						"[Company Application]: Shoma dar hal hazer ozv yek company hastid.",
						colors.ErrorHex,
					)
					return true
				}
				playerI.SendClientMessage(
					"[Company Application]: Your application was sent successfully.",
					colors.SuccessHex,
				)
				return true
			})
			dialog.ShowFor(playerI.Player)
			return
		}
	}
	playerI.Player.SendClientMessage("[Company Application]: Company tag not found.", colors.InfoHex)
}

func companiesHistoryAction(playerI *commons.PlayerI) {
	if len(playerI.CompaniesHistory) <= 0 {
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
	for _, company := range playerI.CompaniesHistory {
		companyI := commons.Companies[company.StoreModel.ID]
		companyHistory.Add(omp.TabListItem{
			companyI.StoreModel.Name,
			companyI.StoreModel.Tag,
			commons.FloatToString(companyI.StoreModel.Multiplier),
		})
	}
	companyHistory.ShowFor(playerI.Player)
}

func companyApplicationsActions(playerI *commons.PlayerI) {
	companyMembership := playerI.GetCurrentCompanyMembership()
	if !playerI.IsInCompany() {
		playerI.SendClientMessage(
			"[Company Applications]: You are not in a company to check incoming company applications.",
			colors.ErrorHex,
		)
		return
	}
	if !playerI.HasCompanyPermission(&commons.CompanyApplicationPermissions, companyMembership.CompanyMember.Role) {
		playerI.SendClientMessage(
			"[Company Applications]: You are not allowed to access company applications.",
			colors.ErrorHex,
		)
		return
	}
	company := commons.Companies[companyMembership.Company.StoreModel.ID]
	if len(company.Applications) <= 0 {
		dialog := omp.NewMessageDialog(
			"Company Applications",
			"No applications found in the company which you are in.",
			"Close",
			"",
		)
		dialog.ShowFor(playerI.Player)
		return
	}
	companyApplicationsDialog := omp.NewTabListDialog("Company Applications", "Select", "Close")
	companyApplicationsDialog.Add(omp.TabListItem{
		"Player ID",
		"Player Name",
		"Requested At",
		"Expire At",
	})
	for _, companyApplication := range company.Applications {
		playerModel := *companyApplication.PlayerModel
		companyApplicationsDialog.Add(omp.TabListItem{
			commons.IntToString(playerModel.ID),
			playerModel.Username,
			companyApplication.StoreModel.CreatedAt.Format("2006-01-02 15:04:05"),
			companyApplication.StoreModel.ExpiredAt.Format("2006-01-02 15:04:05"),
		})
	}
	companyApplicationsDialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
		if e.Response == omp.DialogResponseRight || e.ItemNumber == 0 {
			return true
		}
		playerID, _ := commons.StringToInt[int32](&e.Item[0])
		applicationManagementDialog := omp.NewListDialog("Company Applications", "Select", "Close")
		applicationManagementDialog.Add("Player Stats")
		applicationManagementDialog.Add("Accept")
		applicationManagementDialog.Add("Reject")
		applicationManagementDialog.Add("Cancel")
		applicationManagementDialog.On(omp.EventTypeDialogResponse, func(e *omp.ListDialogResponseEvent) bool {
			if e.Response == omp.DialogResponseRight {
				return true
			}
			ctx := context.Background()
			q := database.New(database.DB)
			switch e.Item {
			case "Player Stats":
				var (
					playerDB                  database.Player
					err                       error
					playerCompanyApplications []database.GetUserCompanyApplicationsHistoryRow
				)
				playerDB, err = q.GetPlayerByID(ctx, playerID)
				if err != nil {
					logger.Fatal("%v", err)
					commons.TechnicalIssueDialog(playerI.Player)
					return true
				}
				playerCompanyApplications, err = q.GetUserCompanyApplicationsHistory(ctx, database.GetUserCompanyApplicationsHistoryParams{
					PlayerID:  playerID,
					CompanyID: company.StoreModel.ID,
				})
				if err != nil {
					logger.Fatal("%v", err)
					commons.TechnicalIssueDialog(playerI.Player)
					return true
				}
				playerStatsDialog := omp.NewTabListDialog(
					"Application Player Stats",
					"<< Back",
					"Close",
				)
				playerStatsDialog.Add(omp.TabListItem{
					"Player Name",
					playerDB.Username,
				})
				playerStatsDialog.Add(omp.TabListItem{
					"Player Level",
					commons.IntToString[int32](playerDB.Level),
				})
				playerStatsDialog.Add(omp.TabListItem{
					"Last Login",
					playerDB.LastLogin.Time.Format("2006-01-02 15:04:05"),
				})
				playerStatsDialog.Add(omp.TabListItem{
					"Last Played",
					playerDB.LastPlayed.Time.Format("2006-01-02 15:04:05"),
				})
				for _, playerCompanyApplication := range playerCompanyApplications {
					playerStatsDialog.Add(omp.TabListItem{
						fmt.Sprintf(
							"%s|%s", playerCompanyApplication.Company.Name,
							commons.If(
								playerCompanyApplication.CompanyApplication.Accepted == enums.Accepted,
								"Accepted",
								"Rejected",
							),
						),
						playerCompanyApplication.CompanyApplication.Answer.String,
					})
				}
				playerStatsDialog.On(omp.EventTypeDialogResponse, func(e *omp.TabListDialogResponseEvent) bool {
					if e.Response == omp.DialogResponseLeft {
						applicationManagementDialog.ShowFor(playerI.Player)
					}
					return true
				})
				playerStatsDialog.ShowFor(playerI.Player)
				return true
			case "Accept", "Reject":
				tx, _ := database.DB.Begin()
				qtx := q.WithTx(tx)
				err := qtx.AnswerCompanyApplication(ctx, database.AnswerCompanyApplicationParams{
					PlayerID:  playerID,
					CompanyID: company.StoreModel.ID,
					Accepted:  commons.If[int16](e.Item == "Accept", enums.Accepted, enums.Rejected),
				})
				if err != nil {
					commons.TechnicalIssueDialog(playerI.Player)
					return true
				}
				failedDialog := omp.NewMessageDialog(
					"Failed to accept application",
					"Failed to accept this answer.",
					"Ok",
					"",
				)
				companyMember, err := qtx.InsertCompanyMembers(ctx, database.InsertCompanyMembersParams{
					CompanyID: company.StoreModel.ID,
					PlayerID:  playerID,
				})
				if err != nil {
					logger.Fatal("Failed to insert company member %v", err)
					failedDialog.Body = "Player is already in another company."
					failedDialog.ShowFor(playerI.Player)
					return true
				}
				err = tx.Commit()
				if err != nil {
					logger.Fatal("Failed to commit transaction: %v", err)
					err = tx.Rollback()
					logger.Fatal("Failed to rollback transaction: %v", err)
					failedDialog.ShowFor(playerI.Player)
					return true
				}
				go companyMembership.Company.ReloadApplications()
				playerI.SendClientMessage("[Company Application]: You've successfully accepted the application.", colors.SuccessHex)
				go func() {
					for _, player := range commons.PlayersI {
						if player.StoreModel.ID == playerID {
							player.SendClientMessage("[Company Application]: You've successfully accepted into company.", colors.SuccessHex)
							player.Membership = &commons.PlayerMembership{
								CompanyMember: &companyMember.CompanyMember,
								Company:       company,
							}
							return
						}
					}
				}()
				return true
			case "Cancel":
				return true
			}
			return true
		})
		applicationManagementDialog.ShowFor(playerI.Player)
		return true
	})
	companyApplicationsDialog.ShowFor(playerI.Player)
}

func playerCompanyStats(playerI *commons.PlayerI) {
	companyMembership := playerI.GetCurrentCompanyMembership()
	if companyMembership == nil {
		playerI.SendClientMessage(
			"[Company Applications]: You are not in a company to check incoming company applications.",
			colors.ErrorHex,
		)
		return
	}
	companyI := commons.Companies[companyMembership.Company.StoreModel.ID]
	statsDialog := companyStatsDialog(companyI)
	statsDialog.ShowFor(playerI.Player)
	return
}

func companyStatsDialog(companyI *commons.CompanyI) *omp.TabListDialog {
	statsDialog := omp.NewTabListDialog("Company Stats", "Ok", "")
	statsDialog.Add(omp.TabListItem{
		"Name",
		companyI.StoreModel.Name,
	})
	statsDialog.Add(omp.TabListItem{
		"Tag",
		companyI.StoreModel.Tag,
	})
	statsDialog.Add(omp.TabListItem{
		"Balance",
		commons.IntToString(companyI.StoreModel.Balance),
	})
	statsDialog.Add(omp.TabListItem{
		"Multiplier",
		commons.FloatToString(companyI.StoreModel.Multiplier),
	})
	return statsDialog
}

func companiesJobsAction(playerI *commons.PlayerI, company *commons.CompanyI) {
	if playerI.Job != nil && playerI.Job.OnDuty {
		playerI.SendClientMessage("[Company Job]: You are already on a duty.", colors.NoteHex)
		return
	}
	companyJobsDialog := omp.NewListDialog("Company Jobs", "Select", "Close")
	companyJobsDialog.Add("Delivery")
	companyJobsDialog.On(omp.EventTypeDialogResponse, func(e *omp.ListDialogResponseEvent) bool {
		if e.Response == omp.DialogResponseRight {
			return true
		}
		switch e.Item {
		case enums.Delivery.String():
			playerI.JoinJob(enums.Delivery, company)
			playerI.SendClientMessage(fmt.Sprintf("[Company Job]: You've hired into %s job successfully.", enums.Delivery.String()), colors.SuccessHex)
			return true
		}
		return true
	})
	companyJobsDialog.ShowFor(playerI.Player)
}

func companiesJobAbandonAction(playerI *commons.PlayerI) {
	if playerI.Job == nil && !playerI.Job.OnDuty {
		playerI.SendClientMessage("[Company Job]: You are not on a duty.", colors.ErrorHex)
		return
	}
	job := playerI.LeaveJob()
	playerI.SendClientMessage(fmt.Sprintf("[Company Job]: You've left the %s job.", job.Name), colors.ErrorHex)
}
