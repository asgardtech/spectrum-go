package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTopNav represents an sp-top-nav component
type spectrumTopNav struct {
	app.Compo

	// Properties
	PropIgnoreURLParts string
	PropLabel          string
	PropQuiet          bool
	PropSelected       string

	// Child components
	PropChildren []app.UI
}

// TopNav creates a new top nav component
func TopNav() *spectrumTopNav {
	return &spectrumTopNav{}
}

// IgnoreURLParts sets which parts of the URL to ignore when matching for the selected Top Nav Item
func (n *spectrumTopNav) IgnoreURLParts(parts string) *spectrumTopNav {
	n.PropIgnoreURLParts = parts
	return n
}

// Label sets the accessible label for the top nav
func (n *spectrumTopNav) Label(label string) *spectrumTopNav {
	n.PropLabel = label
	return n
}

// Quiet sets whether the top nav is displayed without a border
func (n *spectrumTopNav) Quiet(quiet bool) *spectrumTopNav {
	n.PropQuiet = quiet
	return n
}

// Selected sets the value of the selected top nav item
func (n *spectrumTopNav) Selected(selected string) *spectrumTopNav {
	n.PropSelected = selected
	return n
}

// Children sets the child components of the top nav
func (n *spectrumTopNav) Children(children ...app.UI) *spectrumTopNav {
	n.PropChildren = children
	return n
}

// AddItem adds a top nav item to the top nav
func (n *spectrumTopNav) AddItem(item *spectrumTopNavItem) *spectrumTopNav {
	n.PropChildren = append(n.PropChildren, item)
	return n
}

// AddChild adds any UI component as a child of the top nav
func (n *spectrumTopNav) AddChild(child app.UI) *spectrumTopNav {
	n.PropChildren = append(n.PropChildren, child)
	return n
}

// Render renders the top nav component
func (n *spectrumTopNav) Render() app.UI {
	topNav := app.Elem("sp-top-nav")

	// Set attributes
	if n.PropIgnoreURLParts != "" {
		topNav = topNav.Attr("ignore-url-parts", n.PropIgnoreURLParts)
	}
	if n.PropLabel != "" {
		topNav = topNav.Attr("label", n.PropLabel)
	}
	if n.PropQuiet {
		topNav = topNav.Attr("quiet", true)
	}
	if n.PropSelected != "" {
		topNav = topNav.Attr("selected", n.PropSelected)
	}

	// Add children if provided
	if len(n.PropChildren) > 0 {
		topNav = topNav.Body(n.PropChildren...)
	}

	return topNav
}
