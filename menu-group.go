package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumMenuGroup represents an sp-menu-group component
type spectrumMenuGroup struct {
	app.Compo

	// Properties
	PropLabel          string
	PropIgnore         bool
	PropSelects        string
	PropValue          string
	PropValueSeparator string

	// Slots
	PropHeader   app.UI
	PropChildren []app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnClose  app.EventHandler
}

// MenuGroup creates a new menu group component
func MenuGroup() *spectrumMenuGroup {
	return &spectrumMenuGroup{}
}

// Label sets the menu group label
func (m *spectrumMenuGroup) Label(label string) *spectrumMenuGroup {
	m.PropLabel = label
	return m
}

// Ignore sets whether the menu group should ignore keyboard focus navigation
func (m *spectrumMenuGroup) Ignore(ignore bool) *spectrumMenuGroup {
	m.PropIgnore = ignore
	return m
}

// Selects sets the selection mode
func (m *spectrumMenuGroup) Selects(selects string) *spectrumMenuGroup {
	m.PropSelects = selects
	return m
}

// Value sets the selected value(s)
func (m *spectrumMenuGroup) Value(value string) *spectrumMenuGroup {
	m.PropValue = value
	return m
}

// ValueSeparator sets the separator character for multiple selection values
func (m *spectrumMenuGroup) ValueSeparator(separator string) *spectrumMenuGroup {
	m.PropValueSeparator = separator
	return m
}

// Header sets the header slot content
func (m *spectrumMenuGroup) Header(header app.UI) *spectrumMenuGroup {
	m.PropHeader = header
	return m
}

// HeaderText sets a simple text header
func (m *spectrumMenuGroup) HeaderText(text string) *spectrumMenuGroup {
	return m.Header(app.Text(text))
}

// Children adds menu items to the menu group
func (m *spectrumMenuGroup) Children(children ...app.UI) *spectrumMenuGroup {
	m.PropChildren = append(m.PropChildren, children...)
	return m
}

// OnChange sets the change event handler
func (m *spectrumMenuGroup) OnChange(handler app.EventHandler) *spectrumMenuGroup {
	m.PropOnChange = handler
	return m
}

// OnClose sets the close event handler
func (m *spectrumMenuGroup) OnClose(handler app.EventHandler) *spectrumMenuGroup {
	m.PropOnClose = handler
	return m
}

// SelectsInherit sets the selection mode to inherit
func (m *spectrumMenuGroup) SelectsInherit() *spectrumMenuGroup {
	return m.Selects("inherit")
}

// SelectsSingle sets the selection mode to single
func (m *spectrumMenuGroup) SelectsSingle() *spectrumMenuGroup {
	return m.Selects("single")
}

// SelectsMultiple sets the selection mode to multiple
func (m *spectrumMenuGroup) SelectsMultiple() *spectrumMenuGroup {
	return m.Selects("multiple")
}

// Render renders the menu group component
func (m *spectrumMenuGroup) Render() app.UI {
	menuGroup := app.Elem("sp-menu-group")

	// Set attributes
	if m.PropLabel != "" {
		menuGroup = menuGroup.Attr("label", m.PropLabel)
	}
	if m.PropIgnore {
		menuGroup = menuGroup.Attr("ignore", true)
	}
	if m.PropSelects != "" {
		menuGroup = menuGroup.Attr("selects", m.PropSelects)
	}
	if m.PropValue != "" {
		menuGroup = menuGroup.Attr("value", m.PropValue)
	}
	if m.PropValueSeparator != "" {
		menuGroup = menuGroup.Attr("value-separator", m.PropValueSeparator)
	}

	// Add event handlers
	if m.PropOnChange != nil {
		menuGroup = menuGroup.On("change", m.PropOnChange)
	}
	if m.PropOnClose != nil {
		menuGroup = menuGroup.On("close", m.PropOnClose)
	}

	// Create elements array for slots and children
	elements := []app.UI{}

	// Add header if provided
	if m.PropHeader != nil {
		header := m.PropHeader
		if headerWithSlot, ok := header.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, headerWithSlot.Slot("header"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "header").
				Body(header))
		}
	}

	// Add children
	for _, child := range m.PropChildren {
		elements = append(elements, child)
	}

	// Add all elements to the menu group
	if len(elements) > 0 {
		menuGroup = menuGroup.Body(elements...)
	}

	return menuGroup
}
