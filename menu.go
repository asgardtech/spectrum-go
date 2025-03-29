package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumMenu represents an sp-menu component
type spectrumMenu struct {
	app.Compo

	// Properties
	PropLabel          string
	PropIgnore         bool
	PropSelects        string
	PropValue          string
	PropValueSeparator string

	// Collections
	PropMenuItems    []app.UI
	PropMenuGroups   []app.UI
	PropMenuDividers []app.UI
	PropCustomItems  []app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnClose  app.EventHandler
}

// Menu creates a new menu component
func Menu() *spectrumMenu {
	return &spectrumMenu{}
}

// Label sets the menu label
func (m *spectrumMenu) Label(label string) *spectrumMenu {
	m.PropLabel = label
	return m
}

// Ignore sets whether the menu should ignore keyboard focus navigation
func (m *spectrumMenu) Ignore(ignore bool) *spectrumMenu {
	m.PropIgnore = ignore
	return m
}

// Selects sets the selection mode
func (m *spectrumMenu) Selects(selects string) *spectrumMenu {
	m.PropSelects = selects
	return m
}

// Value sets the selected value(s)
func (m *spectrumMenu) Value(value string) *spectrumMenu {
	m.PropValue = value
	return m
}

// ValueSeparator sets the separator character for multiple selection values
func (m *spectrumMenu) ValueSeparator(separator string) *spectrumMenu {
	m.PropValueSeparator = separator
	return m
}

// Items adds menu items to the menu
func (m *spectrumMenu) Items(items ...*spectrumMenuItem) *spectrumMenu {
	for _, item := range items {
		m.PropMenuItems = append(m.PropMenuItems, item)
	}
	return m
}

// Groups adds menu groups to the menu
func (m *spectrumMenu) Groups(groups ...app.UI) *spectrumMenu {
	for _, group := range groups {
		m.PropMenuGroups = append(m.PropMenuGroups, group)
	}
	return m
}

// Divider adds a menu divider
func (m *spectrumMenu) Divider(divider app.UI) *spectrumMenu {
	m.PropMenuDividers = append(m.PropMenuDividers, divider)
	return m
}

// CustomItem adds a custom item
func (m *spectrumMenu) CustomItem(item app.UI) *spectrumMenu {
	m.PropCustomItems = append(m.PropCustomItems, item)
	return m
}

// OnChange sets the change event handler
func (m *spectrumMenu) OnChange(handler app.EventHandler) *spectrumMenu {
	m.PropOnChange = handler
	return m
}

// OnClose sets the close event handler
func (m *spectrumMenu) OnClose(handler app.EventHandler) *spectrumMenu {
	m.PropOnClose = handler
	return m
}

// SelectsNone sets the selection mode to none
func (m *spectrumMenu) SelectsNone() *spectrumMenu {
	return m.Selects("none")
}

// SelectsSingle sets the selection mode to single
func (m *spectrumMenu) SelectsSingle() *spectrumMenu {
	return m.Selects("single")
}

// SelectsMultiple sets the selection mode to multiple
func (m *spectrumMenu) SelectsMultiple() *spectrumMenu {
	return m.Selects("multiple")
}

// Render renders the menu component
func (m *spectrumMenu) Render() app.UI {
	menu := app.Elem("sp-menu")

	// Set attributes
	if m.PropLabel != "" {
		menu = menu.Attr("label", m.PropLabel)
	}
	if m.PropIgnore {
		menu = menu.Attr("ignore", true)
	}
	if m.PropSelects != "" {
		menu = menu.Attr("selects", m.PropSelects)
	}
	if m.PropValue != "" {
		menu = menu.Attr("value", m.PropValue)
	}
	if m.PropValueSeparator != "" {
		menu = menu.Attr("value-separator", m.PropValueSeparator)
	}

	// Add event handlers
	if m.PropOnChange != nil {
		menu = menu.On("change", m.PropOnChange)
	}
	if m.PropOnClose != nil {
		menu = menu.On("close", m.PropOnClose)
	}

	// Gather all menu items, groups, dividers and custom items
	var elements []app.UI

	// Add menu items
	for _, item := range m.PropMenuItems {
		elements = append(elements, item)
	}

	// Add menu groups
	for _, group := range m.PropMenuGroups {
		elements = append(elements, group)
	}

	// Add menu dividers
	for _, divider := range m.PropMenuDividers {
		elements = append(elements, divider)
	}

	// Add custom items
	for _, item := range m.PropCustomItems {
		elements = append(elements, item)
	}

	// Add all elements to the menu
	if len(elements) > 0 {
		menu = menu.Body(elements...)
	}

	return menu
}
