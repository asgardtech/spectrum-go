package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TableAction describes an action button per-row in DataTable.
// Run receives the current go-app Context so the handler can mutate state.

type TableAction struct {
	Label  string
	Icon   Icon
	Intent Intent
	Run    func(app.Context)
}
