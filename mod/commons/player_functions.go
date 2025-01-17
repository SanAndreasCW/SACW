package commons

import "github.com/RahRow/omp"

func TechnicalIssueDialog(player *omp.Player) {
	messageDialog := omp.NewMessageDialog(
		"Application Approval",
		"Application approval failed for technical reasons, please contact support.",
		"Ok",
		"",
	)
	messageDialog.ShowFor(player)
}
