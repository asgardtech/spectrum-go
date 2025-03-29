package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSidenavHeading represents an sp-sidenav-heading component
type spectrumSidenavHeading struct {
	app.Compo

	// Properties
	PropLabel string

	// Slots
	PropChildren []*spectrumSidenavItem
}

// SidenavHeading creates a new sidenav heading component
func SidenavHeading() *spectrumSidenavHeading {
	return &spectrumSidenavHeading{}
}

// Label sets the heading label text
func (sh *spectrumSidenavHeading) Label(label string) *spectrumSidenavHeading {
	sh.PropLabel = label
	return sh
}

// Children sets the child sidenav items
func (sh *spectrumSidenavHeading) Children(children ...*spectrumSidenavItem) *spectrumSidenavHeading {
	sh.PropChildren = children
	return sh
}

// Render renders the sidenav heading component
func (sh *spectrumSidenavHeading) Render() app.UI {
	sidenavHeading := app.Elem("sp-sidenav-heading")

	// Set attributes
	if sh.PropLabel != "" {
		sidenavHeading = sidenavHeading.Attr("label", sh.PropLabel)
	}

	// Add child items if provided
	if len(sh.PropChildren) > 0 {
		// Convert spectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range sh.PropChildren {
			items = append(items, child)
		}

		// Add items to the body
		sidenavHeading = sidenavHeading.Body(items...)
	}

	return sidenavHeading
}
