package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumMenuItem represents an sp-menu-item component
type SpectrumMenuItem struct {
	app.Compo

	// Properties
	text       string
	value      string
	active     bool
	disabled   bool
	download   string
	focused    bool
	hasSubmenu bool
	href       string
	label      string
	noWrap     bool
	open       bool
	selected   bool
	invalid    bool

	// Event handlers
	onClick app.EventHandler
	onBlur  app.EventHandler
	onFocus app.EventHandler

	// Slots
	content     app.UI
	description app.UI
	icon        app.UI
	submenu     app.UI
	valueSlot   app.UI
}

// MenuItem creates a new menu item component
func MenuItem() *SpectrumMenuItem {
	return &SpectrumMenuItem{}
}

// Text sets the text content of the menu item
func (m *SpectrumMenuItem) Text(text string) *SpectrumMenuItem {
	m.text = text
	return m
}

// Value sets the value attribute of the menu item
func (m *SpectrumMenuItem) Value(value string) *SpectrumMenuItem {
	m.value = value
	return m
}

// Active sets whether the menu item is active or has an active descendant
func (m *SpectrumMenuItem) Active(active bool) *SpectrumMenuItem {
	m.active = active
	return m
}

// Disabled sets whether the menu item is disabled
func (m *SpectrumMenuItem) Disabled(disabled bool) *SpectrumMenuItem {
	m.disabled = disabled
	return m
}

// Download sets the download attribute that causes the browser to treat the linked URL as a download
func (m *SpectrumMenuItem) Download(download string) *SpectrumMenuItem {
	m.download = download
	return m
}

// Focused sets whether the menu item has keyboard focus
func (m *SpectrumMenuItem) Focused(focused bool) *SpectrumMenuItem {
	m.focused = focused
	return m
}

// HasSubmenu sets whether the menu item has a submenu
func (m *SpectrumMenuItem) HasSubmenu(hasSubmenu bool) *SpectrumMenuItem {
	m.hasSubmenu = hasSubmenu
	return m
}

// Href sets the URL that the hyperlink points to
func (m *SpectrumMenuItem) Href(href string) *SpectrumMenuItem {
	m.href = href
	return m
}

// Label sets an accessible label that describes the component
func (m *SpectrumMenuItem) Label(label string) *SpectrumMenuItem {
	m.label = label
	return m
}

// NoWrap sets whether menu item text content should not wrap
func (m *SpectrumMenuItem) NoWrap(noWrap bool) *SpectrumMenuItem {
	m.noWrap = noWrap
	return m
}

// Open sets whether the submenu is open
func (m *SpectrumMenuItem) Open(open bool) *SpectrumMenuItem {
	m.open = open
	return m
}

// Selected sets whether the menu item is selected
func (m *SpectrumMenuItem) Selected(selected bool) *SpectrumMenuItem {
	m.selected = selected
	return m
}

// Invalid sets whether the menu item is invalid
func (m *SpectrumMenuItem) Invalid(invalid bool) *SpectrumMenuItem {
	m.invalid = invalid
	return m
}

// Content sets the content of the menu item
func (m *SpectrumMenuItem) Content(content app.UI) *SpectrumMenuItem {
	m.content = content
	return m
}

// Description sets content for the description slot
func (m *SpectrumMenuItem) Description(description app.UI) *SpectrumMenuItem {
	m.description = description
	return m
}

// Icon sets content for the icon slot
func (m *SpectrumMenuItem) Icon(icon app.UI) *SpectrumMenuItem {
	m.icon = icon
	return m
}

// Submenu sets content for the submenu slot
func (m *SpectrumMenuItem) Submenu(submenu app.UI) *SpectrumMenuItem {
	m.submenu = submenu
	return m
}

// ValueSlot sets content for the value slot (displayed at the end of the menu item)
func (m *SpectrumMenuItem) ValueSlot(value app.UI) *SpectrumMenuItem {
	m.valueSlot = value
	return m
}

// OnClick sets the click event handler
func (m *SpectrumMenuItem) OnClick(handler app.EventHandler) *SpectrumMenuItem {
	m.onClick = handler
	return m
}

// OnBlur sets the blur event handler
func (m *SpectrumMenuItem) OnBlur(handler app.EventHandler) *SpectrumMenuItem {
	m.onBlur = handler
	return m
}

// OnFocus sets the focus event handler
func (m *SpectrumMenuItem) OnFocus(handler app.EventHandler) *SpectrumMenuItem {
	m.onFocus = handler
	return m
}

// Render renders the menu item component
func (m *SpectrumMenuItem) Render() app.UI {
	menuItem := app.Elem("sp-menu-item")

	// Set attributes
	if m.value != "" {
		menuItem = menuItem.Attr("value", m.value)
	}
	if m.active {
		menuItem = menuItem.Attr("active", true)
	}
	if m.disabled {
		menuItem = menuItem.Attr("disabled", true)
	}
	if m.download != "" {
		menuItem = menuItem.Attr("download", m.download)
	}
	if m.focused {
		menuItem = menuItem.Attr("focused", true)
	}
	if m.hasSubmenu {
		menuItem = menuItem.Attr("has-submenu", true)
	}
	if m.href != "" {
		menuItem = menuItem.Attr("href", m.href)
	}
	if m.label != "" {
		menuItem = menuItem.Attr("label", m.label)
	}
	if m.noWrap {
		menuItem = menuItem.Attr("no-wrap", true)
	}
	if m.open {
		menuItem = menuItem.Attr("open", true)
	}
	if m.selected {
		menuItem = menuItem.Attr("selected", true)
	}
	if m.invalid {
		menuItem = menuItem.Attr("invalid", true)
	}

	// Add event handlers
	if m.onClick != nil {
		menuItem = menuItem.OnClick(m.onClick)
	}
	if m.onBlur != nil {
		menuItem = menuItem.OnBlur(m.onBlur)
	}
	if m.onFocus != nil {
		menuItem = menuItem.OnFocus(m.onFocus)
	}

	// Handle slots
	elements := []app.UI{}

	// Add icon if provided
	if m.icon != nil {
		icon := m.icon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		elements = append(elements, icon)
	}

	// Add description if provided
	if m.description != nil {
		desc := m.description
		if descWithSlot, ok := desc.(interface{ Slot(string) app.UI }); ok {
			desc = descWithSlot.Slot("description")
		} else {
			desc = app.Elem("div").
				Attr("slot", "description").
				Body(desc)
		}
		elements = append(elements, desc)
	}

	// Add value if provided
	if m.valueSlot != nil {
		val := m.valueSlot
		if valWithSlot, ok := val.(interface{ Slot(string) app.UI }); ok {
			val = valWithSlot.Slot("value")
		} else {
			val = app.Elem("div").
				Attr("slot", "value").
				Body(val)
		}
		elements = append(elements, val)
	}

	// Add submenu if provided
	if m.submenu != nil {
		sub := m.submenu
		if subWithSlot, ok := sub.(interface{ Slot(string) app.UI }); ok {
			sub = subWithSlot.Slot("submenu")
		} else {
			sub = app.Elem("div").
				Attr("slot", "submenu").
				Body(sub)
		}
		elements = append(elements, sub)
	}

	// Add content or text if provided
	if m.content != nil {
		elements = append(elements, m.content)
	} else if m.text != "" {
		menuItem = menuItem.Text(m.text)
	}

	if len(elements) > 0 {
		menuItem = menuItem.Body(elements...)
	}

	return menuItem
}
