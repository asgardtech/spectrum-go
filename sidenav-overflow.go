package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSidenavOverflow represents an sp-sidenav-overflow component
type spectrumSidenavOverflow struct {
	app.Compo

	// Properties
	PropLabel string

	// Child components
	PropChildren []*spectrumSidenavItem

	// Event handlers
	PropOnClick app.EventHandler
}

// SidenavOverflow creates a new sidenav overflow component
func SidenavOverflow() *spectrumSidenavOverflow {
	return &spectrumSidenavOverflow{}
}

// Label sets the accessible label for the overflow component
func (so *spectrumSidenavOverflow) Label(label string) *spectrumSidenavOverflow {
	so.PropLabel = label
	return so
}

// Children sets the child sidenav items
func (so *spectrumSidenavOverflow) Children(children ...*spectrumSidenavItem) *spectrumSidenavOverflow {
	so.PropChildren = children
	return so
}

// OnClick sets the click event handler
func (so *spectrumSidenavOverflow) OnClick(handler app.EventHandler) *spectrumSidenavOverflow {
	so.PropOnClick = handler
	return so
}

// AddItem adds a sidenav item to the overflow
func (so *spectrumSidenavOverflow) AddItem(item *spectrumSidenavItem) *spectrumSidenavOverflow {
	so.PropChildren = append(so.PropChildren, item)
	return so
}

// Render renders the sidenav overflow component
func (so *spectrumSidenavOverflow) Render() app.UI {
	overflow := app.Elem("sp-sidenav-overflow")

	// Set attributes
	if so.PropLabel != "" {
		overflow = overflow.Attr("label", so.PropLabel)
	}

	// Add event handlers
	if so.PropOnClick != nil {
		overflow = overflow.On("click", so.PropOnClick)
	}

	// Add children if provided
	if len(so.PropChildren) > 0 {
		// Convert spectrumSidenavItem children to app.UI
		var items []app.UI
		for _, child := range so.PropChildren {
			items = append(items, child)
		}

		// Add items to the body
		overflow = overflow.Body(items...)
	}

	return overflow
}
