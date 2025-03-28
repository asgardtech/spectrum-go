package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// Selects represents selection mode for menu groups
type MenuSelects string

const (
	MenuSelectsInherit  MenuSelects = "inherit"
	MenuSelectsSingle   MenuSelects = "single"
	MenuSelectsMultiple MenuSelects = "multiple"
)

// SpectrumMenuGroup represents an sp-menu-group component
type SpectrumMenuGroup struct {
	app.Compo

	// Properties
	label          string
	ignore         bool
	selects        MenuSelects
	value          string
	valueSeparator string

	// Slots
	header   app.UI
	children []*SpectrumMenuItem

	// Event handlers
	onChange app.EventHandler
	onClose  app.EventHandler
}

// MenuGroup creates a new menu group component
func MenuGroup() *SpectrumMenuGroup {
	return &SpectrumMenuGroup{
		valueSeparator: ",", // Default value separator
	}
}

// Label sets the label of the menu group
func (m *SpectrumMenuGroup) Label(label string) *SpectrumMenuGroup {
	m.label = label
	return m
}

// Ignore sets whether menu should be ignored by roving tabindex controller
func (m *SpectrumMenuGroup) Ignore(ignore bool) *SpectrumMenuGroup {
	m.ignore = ignore
	return m
}

// Selects sets how the menu allows selection of its items
func (m *SpectrumMenuGroup) Selects(selects MenuSelects) *SpectrumMenuGroup {
	m.selects = selects
	return m
}

// Value sets the value of the selected item(s)
func (m *SpectrumMenuGroup) Value(value string) *SpectrumMenuGroup {
	m.value = value
	return m
}

// ValueSeparator sets the separator for multiple values
func (m *SpectrumMenuGroup) ValueSeparator(separator string) *SpectrumMenuGroup {
	m.valueSeparator = separator
	return m
}

// Header sets the content for the header slot
func (m *SpectrumMenuGroup) Header(header app.UI) *SpectrumMenuGroup {
	m.header = header
	return m
}

// HeaderText sets text content for the header slot
func (m *SpectrumMenuGroup) HeaderText(text string) *SpectrumMenuGroup {
	m.header = app.Text(text)
	return m
}

// Children sets the menu items
func (m *SpectrumMenuGroup) Children(children ...*SpectrumMenuItem) *SpectrumMenuGroup {
	m.children = children
	return m
}

// OnChange sets the change event handler
func (m *SpectrumMenuGroup) OnChange(handler app.EventHandler) *SpectrumMenuGroup {
	m.onChange = handler
	return m
}

// OnClose sets the close event handler
func (m *SpectrumMenuGroup) OnClose(handler app.EventHandler) *SpectrumMenuGroup {
	m.onClose = handler
	return m
}

// SelectsInherit sets selects to "inherit"
func (m *SpectrumMenuGroup) SelectsInherit() *SpectrumMenuGroup {
	return m.Selects(MenuSelectsInherit)
}

// SelectsSingle sets selects to "single"
func (m *SpectrumMenuGroup) SelectsSingle() *SpectrumMenuGroup {
	return m.Selects(MenuSelectsSingle)
}

// SelectsMultiple sets selects to "multiple"
func (m *SpectrumMenuGroup) SelectsMultiple() *SpectrumMenuGroup {
	return m.Selects(MenuSelectsMultiple)
}

// Render renders the menu group component
func (m *SpectrumMenuGroup) Render() app.UI {
	menuGroup := app.Elem("sp-menu-group")

	// Set attributes
	if m.label != "" {
		menuGroup = menuGroup.Attr("label", m.label)
	}
	if m.ignore {
		menuGroup = menuGroup.Attr("ignore", true)
	}
	if m.selects != "" {
		menuGroup = menuGroup.Attr("selects", string(m.selects))
	}
	if m.value != "" {
		menuGroup = menuGroup.Attr("value", m.value)
	}
	if m.valueSeparator != "," { // Only set if not the default
		menuGroup = menuGroup.Attr("value-separator", m.valueSeparator)
	}

	// Add event handlers
	if m.onChange != nil {
		menuGroup = menuGroup.On("change", m.onChange)
	}
	if m.onClose != nil {
		menuGroup = menuGroup.On("close", m.onClose)
	}

	// Add header if provided
	if m.header != nil {
		header := m.header
		if headerWithSlot, ok := header.(interface{ Slot(string) app.UI }); ok {
			header = headerWithSlot.Slot("header")
		} else {
			header = app.Elem("div").
				Attr("slot", "header").
				Body(header)
		}
		menuGroup = menuGroup.Body(header)
	}

	// Add menu items
	if len(m.children) > 0 {
		// Convert SpectrumMenuItem children to app.UI
		var items []app.UI
		for _, child := range m.children {
			items = append(items, child)
		}

		// Add items to the body if there are already elements in the body (header)
		if m.header != nil {
			menuGroup = menuGroup.Body(items...)
		} else {
			menuGroup = menuGroup.Body(items...)
		}
	}

	return menuGroup
}
