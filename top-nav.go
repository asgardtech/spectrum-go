package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTopNav represents an sp-top-nav component
type SpectrumTopNav struct {
	app.Compo

	// Properties
	ignoreURLParts string
	label          string
	quiet          bool
	selected       string

	// Child components
	children []app.UI
}

// TopNav creates a new top nav component
func TopNav() *SpectrumTopNav {
	return &SpectrumTopNav{}
}

// IgnoreURLParts sets which parts of the URL to ignore when matching for the selected Top Nav Item
func (n *SpectrumTopNav) IgnoreURLParts(parts string) *SpectrumTopNav {
	n.ignoreURLParts = parts
	return n
}

// Label sets the accessible label for the top nav
func (n *SpectrumTopNav) Label(label string) *SpectrumTopNav {
	n.label = label
	return n
}

// Quiet sets whether the top nav is displayed without a border
func (n *SpectrumTopNav) Quiet(quiet bool) *SpectrumTopNav {
	n.quiet = quiet
	return n
}

// Selected sets the value of the selected top nav item
func (n *SpectrumTopNav) Selected(selected string) *SpectrumTopNav {
	n.selected = selected
	return n
}

// Children sets the child components of the top nav
func (n *SpectrumTopNav) Children(children ...app.UI) *SpectrumTopNav {
	n.children = children
	return n
}

// AddItem adds a top nav item to the top nav
func (n *SpectrumTopNav) AddItem(item *SpectrumTopNavItem) *SpectrumTopNav {
	n.children = append(n.children, item)
	return n
}

// AddChild adds any UI component as a child of the top nav
func (n *SpectrumTopNav) AddChild(child app.UI) *SpectrumTopNav {
	n.children = append(n.children, child)
	return n
}

// Render renders the top nav component
func (n *SpectrumTopNav) Render() app.UI {
	topNav := app.Elem("sp-top-nav")

	// Set attributes
	if n.ignoreURLParts != "" {
		topNav = topNav.Attr("ignore-url-parts", n.ignoreURLParts)
	}
	if n.label != "" {
		topNav = topNav.Attr("label", n.label)
	}
	if n.quiet {
		topNav = topNav.Attr("quiet", true)
	}
	if n.selected != "" {
		topNav = topNav.Attr("selected", n.selected)
	}

	// Add children if provided
	if len(n.children) > 0 {
		topNav = topNav.Body(n.children...)
	}

	return topNav
}
