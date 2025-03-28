package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumMenu represents an sp-menu component
type SpectrumMenu struct {
	app.Compo

	// Properties
	label          string
	ignore         bool
	selects        MenuSelects
	value          string
	valueSeparator string

	// Children
	menuItems    []*SpectrumMenuItem
	menuGroups   []*SpectrumMenuGroup
	menuDividers []*SpectrumMenuDivider
	customItems  []app.UI

	// Event handlers
	onChange app.EventHandler
	onClose  app.EventHandler
}

// Menu creates a new menu component
func Menu() *SpectrumMenu {
	return &SpectrumMenu{
		valueSeparator: ",", // Default value separator
	}
}

// Label sets the label of the menu
func (m *SpectrumMenu) Label(label string) *SpectrumMenu {
	m.label = label
	return m
}

// Ignore sets whether menu should be ignored by roving tabindex controller
func (m *SpectrumMenu) Ignore(ignore bool) *SpectrumMenu {
	m.ignore = ignore
	return m
}

// Selects sets how the menu allows selection of its items
func (m *SpectrumMenu) Selects(selects MenuSelects) *SpectrumMenu {
	m.selects = selects
	return m
}

// Value sets the value of the selected item(s)
func (m *SpectrumMenu) Value(value string) *SpectrumMenu {
	m.value = value
	return m
}

// ValueSeparator sets the separator for multiple values
func (m *SpectrumMenu) ValueSeparator(separator string) *SpectrumMenu {
	m.valueSeparator = separator
	return m
}

// Items sets the menu items
func (m *SpectrumMenu) Items(items ...*SpectrumMenuItem) *SpectrumMenu {
	m.menuItems = items
	return m
}

// Groups sets the menu groups
func (m *SpectrumMenu) Groups(groups ...*SpectrumMenuGroup) *SpectrumMenu {
	m.menuGroups = groups
	return m
}

// Divider adds a menu divider
func (m *SpectrumMenu) Divider() *SpectrumMenu {
	m.menuDividers = append(m.menuDividers, MenuDivider())
	return m
}

// DividerWithSize adds a menu divider with a specific size
func (m *SpectrumMenu) DividerWithSize(size string) *SpectrumMenu {
	m.menuDividers = append(m.menuDividers, MenuDivider().Size(size))
	return m
}

// CustomItem adds a custom UI element to the menu
func (m *SpectrumMenu) CustomItem(item app.UI) *SpectrumMenu {
	m.customItems = append(m.customItems, item)
	return m
}

// OnChange sets the change event handler
func (m *SpectrumMenu) OnChange(handler app.EventHandler) *SpectrumMenu {
	m.onChange = handler
	return m
}

// OnClose sets the close event handler
func (m *SpectrumMenu) OnClose(handler app.EventHandler) *SpectrumMenu {
	m.onClose = handler
	return m
}

// SelectsInherit sets selects to "inherit"
func (m *SpectrumMenu) SelectsInherit() *SpectrumMenu {
	return m.Selects(MenuSelectsInherit)
}

// SelectsSingle sets selects to "single"
func (m *SpectrumMenu) SelectsSingle() *SpectrumMenu {
	return m.Selects(MenuSelectsSingle)
}

// SelectsMultiple sets selects to "multiple"
func (m *SpectrumMenu) SelectsMultiple() *SpectrumMenu {
	return m.Selects(MenuSelectsMultiple)
}

// Render renders the menu component
func (m *SpectrumMenu) Render() app.UI {
	menu := app.Elem("sp-menu")

	// Set attributes
	if m.label != "" {
		menu = menu.Attr("label", m.label)
	}
	if m.ignore {
		menu = menu.Attr("ignore", true)
	}
	if m.selects != "" {
		menu = menu.Attr("selects", string(m.selects))
	}
	if m.value != "" {
		menu = menu.Attr("value", m.value)
	}
	if m.valueSeparator != "," { // Only set if not the default
		menu = menu.Attr("value-separator", m.valueSeparator)
	}

	// Add event handlers
	if m.onChange != nil {
		menu = menu.On("change", m.onChange)
	}
	if m.onClose != nil {
		menu = menu.On("close", m.onClose)
	}

	// Collect all children elements
	var children []app.UI

	// Add menu items
	for _, item := range m.menuItems {
		children = append(children, item)
	}

	// Add menu groups
	for _, group := range m.menuGroups {
		children = append(children, group)
	}

	// Add menu dividers
	for _, divider := range m.menuDividers {
		children = append(children, divider)
	}

	// Add custom items
	for _, item := range m.customItems {
		children = append(children, item)
	}

	// Add all children to the menu
	if len(children) > 0 {
		menu = menu.Body(children...)
	}

	return menu
}
