package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumSidenavOverflow represents an sp-sidenav-overflow component
type SpectrumSidenavOverflow struct {
	app.Compo

	// Properties
	label string

	// Child components
	children []*SpectrumSidenavItem

	// Event handlers
	onClick app.EventHandler
}

// SidenavOverflow creates a new sidenav overflow component
func SidenavOverflow() *SpectrumSidenavOverflow {
	return &SpectrumSidenavOverflow{}
}

// Label sets the accessible label for the overflow component
func (so *SpectrumSidenavOverflow) Label(label string) *SpectrumSidenavOverflow {
	so.label = label
	return so
}

// Children sets the child sidenav items
func (so *SpectrumSidenavOverflow) Children(children ...*SpectrumSidenavItem) *SpectrumSidenavOverflow {
	so.children = children
	return so
}

// OnClick sets the click event handler
func (so *SpectrumSidenavOverflow) OnClick(handler app.EventHandler) *SpectrumSidenavOverflow {
	so.onClick = handler
	return so
}

// AddItem adds a sidenav item to the overflow
func (so *SpectrumSidenavOverflow) AddItem(item *SpectrumSidenavItem) *SpectrumSidenavOverflow {
	so.children = append(so.children, item)
	return so
}

// Render renders the sidenav overflow component
func (so *SpectrumSidenavOverflow) Render() app.UI {
	overflow := app.Elem("sp-sidenav-overflow")

	// Set attributes
	if so.label != "" {
		overflow = overflow.Attr("label", so.label)
	}

	// Add event handlers
	if so.onClick != nil {
		overflow = overflow.On("click", so.onClick)
	}

	// Add children if provided
	if len(so.children) > 0 {
		// Convert SpectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range so.children {
			items = append(items, child)
		}

		// Add items to the body
		overflow = overflow.Body(items...)
	}

	return overflow
}
