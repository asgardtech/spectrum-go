package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumSidenav represents an sp-sidenav component
type SpectrumSidenav struct {
	app.Compo

	// Properties
	disabled       bool
	label          string
	manageTabIndex bool
	tabIndex       int
	value          string
	variant        string
	defaultValue   string

	// Child components
	children []app.UI // Can be a mix of items and headings

	// Event handlers
	onChange app.EventHandler
}

// Sidenav creates a new sidenav component
func Sidenav() *SpectrumSidenav {
	return &SpectrumSidenav{}
}

// Disabled sets whether the sidenav is disabled
func (s *SpectrumSidenav) Disabled(disabled bool) *SpectrumSidenav {
	s.disabled = disabled
	return s
}

// Label sets the accessible label for the sidenav
func (s *SpectrumSidenav) Label(label string) *SpectrumSidenav {
	s.label = label
	return s
}

// ManageTabIndex sets whether the sidenav should manage tab indices of child items
func (s *SpectrumSidenav) ManageTabIndex(manage bool) *SpectrumSidenav {
	s.manageTabIndex = manage
	return s
}

// TabIndex sets the tab index of the sidenav
func (s *SpectrumSidenav) TabIndex(tabIndex int) *SpectrumSidenav {
	s.tabIndex = tabIndex
	return s
}

// Value sets the value of the sidenav
func (s *SpectrumSidenav) Value(value string) *SpectrumSidenav {
	s.value = value
	return s
}

// DefaultValue sets the default value of the sidenav
func (s *SpectrumSidenav) DefaultValue(value string) *SpectrumSidenav {
	s.defaultValue = value
	return s
}

// Variant sets the variant of the sidenav
func (s *SpectrumSidenav) Variant(variant string) *SpectrumSidenav {
	s.variant = variant
	return s
}

// MultilevelVariant sets the variant to "multilevel"
func (s *SpectrumSidenav) MultilevelVariant() *SpectrumSidenav {
	return s.Variant("multilevel")
}

// AddItem adds a sidenav item to the sidenav
func (s *SpectrumSidenav) AddItem(item *SpectrumSidenavItem) *SpectrumSidenav {
	s.children = append(s.children, item)
	return s
}

// AddHeading adds a sidenav heading to the sidenav
func (s *SpectrumSidenav) AddHeading(heading *SpectrumSidenavHeading) *SpectrumSidenav {
	s.children = append(s.children, heading)
	return s
}

// Children sets all children (items and headings) for the sidenav
func (s *SpectrumSidenav) Children(children ...app.UI) *SpectrumSidenav {
	s.children = children
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSidenav) OnChange(handler app.EventHandler) *SpectrumSidenav {
	s.onChange = handler
	return s
}

// Render renders the sidenav component
func (s *SpectrumSidenav) Render() app.UI {
	sidenav := app.Elem("sp-sidenav")

	// Set attributes
	if s.disabled {
		sidenav = sidenav.Attr("disabled", true)
	}
	if s.label != "" {
		sidenav = sidenav.Attr("label", s.label)
	}
	if s.manageTabIndex {
		sidenav = sidenav.Attr("manage-tab-index", true)
	}
	if s.tabIndex != 0 {
		sidenav = sidenav.Attr("tabIndex", s.tabIndex)
	}
	if s.value != "" {
		sidenav = sidenav.Attr("value", s.value)
	}
	if s.defaultValue != "" {
		sidenav = sidenav.Attr("defaultValue", s.defaultValue)
	}
	if s.variant != "" {
		sidenav = sidenav.Attr("variant", s.variant)
	}

	// Add event handlers
	if s.onChange != nil {
		sidenav = sidenav.On("change", s.onChange)
	}

	// Add children if provided
	if len(s.children) > 0 {
		sidenav = sidenav.Body(s.children...)
	}

	return sidenav
}
