package company

import (
	"context"
	"fmt"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/setting"
	"github.com/kodeyeen/omp"
)

func companiesApplicationAction(playerI *commons.PlayerI, tag string) {
	for _, company := range commons.Companies {
		if company.StoreModel.Tag == tag {
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
			dialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
				ep := e.Payload().(*omp.InputDialogResponseEvent)
				if ep.Response == omp.DialogResponseRight {
					return nil
				}
				if len(ep.InputText) > 80 {
					dialog.Body = fmt.Sprintf("%s\nDescription can't be larger than 80 characters.", dialogBody)
					dialog.ShowFor(playerI.Player)
					return nil
				}
				if len(company.Applications) >= setting.MaxCompanyApplications {
					playerI.SendClientMessage(
						"[Company Application]: Targeted company is not capable for more applications.",
						colors.NoteColor.Hex,
					)
					return nil
				}
				isCreated := company.CreateApplication(playerI, ep.InputText)
				if isCreated != true {
					logger.Fatal("[CompanyApplication]: Failed to save company application")
					playerI.SendClientMessage(
						"[Company Application]: Shoma dar hal hazer ozv yek company hastid.",
						colors.ErrorColor.Hex,
					)
					return nil
				}
				playerI.SendClientMessage(
					"[Company Application]: Your application was sent successfully.",
					colors.SuccessColor.Hex,
				)
				return nil
			})
			dialog.ShowFor(playerI.Player)
			return
		}
	}
	playerI.Player.SendClientMessage("[Company Application]: Company tag not found.", colors.InfoColor.Hex)
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

func companyApplicationsActions(ctx context.Context, playerI *commons.PlayerI) {
	companyMembership := playerI.GetCurrentCompanyMembership(ctx)
	if !playerI.IsInCompany() {
		playerI.SendClientMessage(
			"[Company Applications]: You are not in a company to check incoming company applications.",
			colors.ErrorColor.Hex,
		)
		return
	}
	if !playerI.HasCompanyPermission(&commons.CompanyApplicationPermissions, companyMembership.CompanyMember.Role) {
		playerI.SendClientMessage(
			"[Company Applications]: You are not allowed to access company applications.",
			colors.ErrorColor.Hex,
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
	companyApplicationsDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.TabListDialogResponseEvent)
		if ep.Response == omp.DialogResponseRight || ep.ItemNumber == 0 {
			return nil
		}
		playerID, _ := commons.StringToInt[int32](&ep.Item[0])
		applicationManagementDialog := omp.NewListDialog("Company Applications", "Select", "Close")
		applicationManagementDialog.Add("Player Stats")
		applicationManagementDialog.Add("Accept")
		applicationManagementDialog.Add("Reject")
		applicationManagementDialog.Add("Cancel")
		applicationManagementDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
			ep := e.Payload().(*omp.ListDialogResponseEvent)
			if ep.Response == omp.DialogResponseRight {
				return nil
			}
			q := database.New(database.DB)
			switch ep.Item {
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
					return nil
				}
				playerCompanyApplications, err = q.GetUserCompanyApplicationsHistory(ctx, database.GetUserCompanyApplicationsHistoryParams{
					PlayerID:  playerID,
					CompanyID: company.StoreModel.ID,
				})
				if err != nil {
					logger.Fatal("%v", err)
					commons.TechnicalIssueDialog(playerI.Player)
					return nil
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
				playerStatsDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
					ep := e.Payload().(*omp.TabListDialogResponseEvent)
					if ep.Response == omp.DialogResponseLeft {
						applicationManagementDialog.ShowFor(playerI.Player)
					}
					return nil
				})
				playerStatsDialog.ShowFor(playerI.Player)
				return nil
			case "Accept", "Reject":
				tx, _ := database.DB.Begin()
				qtx := q.WithTx(tx)
				err := qtx.AnswerCompanyApplication(ctx, database.AnswerCompanyApplicationParams{
					PlayerID:  playerID,
					CompanyID: company.StoreModel.ID,
					Accepted:  commons.If[int16](ep.Item == "Accept", enums.Accepted, enums.Rejected),
				})
				if err != nil {
					commons.TechnicalIssueDialog(playerI.Player)
					return nil
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
					return nil
				}
				err = tx.Commit()
				if err != nil {
					logger.Fatal("Failed to commit transaction: %v", err)
					err = tx.Rollback()
					logger.Fatal("Failed to rollback transaction: %v", err)
					failedDialog.ShowFor(playerI.Player)
					return nil
				}
				go companyMembership.Company.ReloadApplications()
				playerI.SendClientMessage("[Company Application]: You've successfully accepted the application.", colors.SuccessColor.Hex)
				go func() {
					for _, player := range commons.PlayersI {
						if player.StoreModel.ID == playerID {
							player.SendClientMessage("[Company Application]: You've successfully accepted into company.", colors.SuccessColor.Hex)
							player.Membership = &commons.PlayerMembership{
								CompanyMember: &companyMember.CompanyMember,
								Company:       company,
							}
							return
						}
					}
				}()
				return nil
			case "Cancel":
				return nil
			}
			return nil
		})
		applicationManagementDialog.ShowFor(playerI.Player)
		return nil
	})
	companyApplicationsDialog.ShowFor(playerI.Player)
}

func playerCompanyStats(ctx context.Context, playerI *commons.PlayerI) {
	companyMembership := playerI.GetCurrentCompanyMembership(ctx)
	if companyMembership == nil {
		playerI.SendClientMessage(
			"[Company Applications]: You are not in a company to check incoming company applications.",
			colors.ErrorColor.Hex,
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
		playerI.SendClientMessage("[Company Job]: You are already on a duty.", colors.NoteColor.Hex)
		return
	}
	companyJobsDialog := omp.NewListDialog("Company Jobs", "Select", "Close")
	for _, job := range company.CompanyJobs {
		jobID := enums.JobType(job.JobID)
		companyJobsDialog.Add(jobID.String())
	}
	companyJobsDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
		ep := e.Payload().(*omp.ListDialogResponseEvent)
		if ep.Response == omp.DialogResponseRight {
			return nil
		}
		job := enums.GetJobType(ep.Item)
		if job == enums.Unknown {
			playerI.SendClientMessage("[Company Jobs]: Invalid job type.", colors.ErrorColor.Hex)
			return nil
		}
		playerI.JoinJob(job, company)
		playerI.SendClientMessage(fmt.Sprintf("[Company Job]: You've hired into %s job successfully.", job.String()), colors.SuccessColor.Hex)
		return nil
	})
	companyJobsDialog.ShowFor(playerI.Player)
}

func companiesJobAbandonAction(playerI *commons.PlayerI) {
	if playerI.Job == nil && !playerI.Job.OnDuty {
		playerI.SendClientMessage("[Company Job]: You are not on a duty.", colors.ErrorColor.Hex)
		return
	}
	job := playerI.LeaveJob()
	playerI.SendClientMessage(fmt.Sprintf("[Company Job]: You've left the %s job.", job.Name), colors.ErrorColor.Hex)
}
