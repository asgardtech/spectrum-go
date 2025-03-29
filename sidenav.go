package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSidenav represents an sp-sidenav component
type spectrumSidenav struct {
	app.Compo

	// Properties
	PropDisabled       bool
	PropLabel          string
	PropManageTabIndex bool
	PropTabIndex       int
	PropValue          string
	PropVariant        string
	PropDefaultValue   string

	// Child components
	PropChildren []app.UI // Can be a mix of items and headings

	// Event handlers
	PropOnChange app.EventHandler
}

// Sidenav creates a new sidenav component
func Sidenav() *spectrumSidenav {
	return &spectrumSidenav{}
}

// Disabled sets whether the sidenav is disabled
func (s *spectrumSidenav) Disabled(disabled bool) *spectrumSidenav {
	s.PropDisabled = disabled
	return s
}

// Label sets the accessible label for the sidenav
func (s *spectrumSidenav) Label(label string) *spectrumSidenav {
	s.PropLabel = label
	return s
}

// ManageTabIndex sets whether the sidenav should manage tab indices of child items
func (s *spectrumSidenav) ManageTabIndex(manage bool) *spectrumSidenav {
	s.PropManageTabIndex = manage
	return s
}

// TabIndex sets the tab index of the sidenav
func (s *spectrumSidenav) TabIndex(tabIndex int) *spectrumSidenav {
	s.PropTabIndex = tabIndex
	return s
}

// Value sets the value of the sidenav
func (s *spectrumSidenav) Value(value string) *spectrumSidenav {
	s.PropValue = value
	return s
}

// DefaultValue sets the default value of the sidenav
func (s *spectrumSidenav) DefaultValue(value string) *spectrumSidenav {
	s.PropDefaultValue = value
	return s
}

// Variant sets the variant of the sidenav
func (s *spectrumSidenav) Variant(variant string) *spectrumSidenav {
	s.PropVariant = variant
	return s
}

// MultilevelVariant sets the variant to "multilevel"
func (s *spectrumSidenav) MultilevelVariant() *spectrumSidenav {
	return s.Variant("multilevel")
}

// AddItem adds a sidenav item to the sidenav
func (s *spectrumSidenav) AddItem(item *spectrumSidenavItem) *spectrumSidenav {
	s.PropChildren = append(s.PropChildren, item)
	return s
}

// AddHeading adds a sidenav heading to the sidenav
func (s *spectrumSidenav) AddHeading(heading *spectrumSidenavHeading) *spectrumSidenav {
	s.PropChildren = append(s.PropChildren, heading)
	return s
}

// Children sets all children (items and headings) for the sidenav
func (s *spectrumSidenav) Children(children ...app.UI) *spectrumSidenav {
	s.PropChildren = children
	return s
}

// OnChange sets the change event handler
func (s *spectrumSidenav) OnChange(handler app.EventHandler) *spectrumSidenav {
	s.PropOnChange = handler
	return s
}

// Render renders the sidenav component
func (s *spectrumSidenav) Render() app.UI {
	sidenav := app.Elem("sp-sidenav")

	// Set attributes
	if s.PropDisabled {
		sidenav = sidenav.Attr("disabled", true)
	}
	if s.PropLabel != "" {
		sidenav = sidenav.Attr("label", s.PropLabel)
	}
	if s.PropManageTabIndex {
		sidenav = sidenav.Attr("manage-tab-index", true)
	}
	if s.PropTabIndex != 0 {
		sidenav = sidenav.Attr("tabIndex", s.PropTabIndex)
	}
	if s.PropValue != "" {
		sidenav = sidenav.Attr("value", s.PropValue)
	}
	if s.PropDefaultValue != "" {
		sidenav = sidenav.Attr("defaultValue", s.PropDefaultValue)
	}
	if s.PropVariant != "" {
		sidenav = sidenav.Attr("variant", s.PropVariant)
	}

	// Add event handlers
	if s.PropOnChange != nil {
		sidenav = sidenav.On("change", s.PropOnChange)
	}

	// Add children if provided
	if len(s.PropChildren) > 0 {
		sidenav = sidenav.Body(s.PropChildren...)
	}

	return sidenav
}
