package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumSidenavHeading represents an sp-sidenav-heading component
type SpectrumSidenavHeading struct {
	app.Compo

	// Properties
	label string

	// Slots
	children []*SpectrumSidenavItem
}

// SidenavHeading creates a new sidenav heading component
func SidenavHeading() *SpectrumSidenavHeading {
	return &SpectrumSidenavHeading{}
}

// Label sets the heading label text
func (sh *SpectrumSidenavHeading) Label(label string) *SpectrumSidenavHeading {
	sh.label = label
	return sh
}

// Children sets the child sidenav items
func (sh *SpectrumSidenavHeading) Children(children ...*SpectrumSidenavItem) *SpectrumSidenavHeading {
	sh.children = children
	return sh
}

// Render renders the sidenav heading component
func (sh *SpectrumSidenavHeading) Render() app.UI {
	sidenavHeading := app.Elem("sp-sidenav-heading")

	// Set attributes
	if sh.label != "" {
		sidenavHeading = sidenavHeading.Attr("label", sh.label)
	}

	// Add child items if provided
	if len(sh.children) > 0 {
		// Convert SpectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range sh.children {
			items = append(items, child)
		}

		// Add items to the body
		sidenavHeading = sidenavHeading.Body(items...)
	}

	return sidenavHeading
}
