package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumMenuItem represents an sp-menu-item component
type spectrumMenuItem struct {
	app.Compo

	// Properties
	PropText       string
	PropValue      string
	PropActive     bool
	PropDisabled   bool
	PropDownload   string
	PropFocused    bool
	PropHasSubmenu bool
	PropHref       string
	PropLabel      string
	PropNoWrap     bool
	PropOpen       bool
	PropSelected   bool
	PropInvalid    bool

	// Event handlers
	PropOnClick app.EventHandler
	PropOnBlur  app.EventHandler
	PropOnFocus app.EventHandler

	// Slots
	PropContent     app.UI
	PropDescription app.UI
	PropIcon        app.UI
	PropSubmenu     app.UI
	PropValueSlot   app.UI
}

// MenuItem creates a new menu item component
func MenuItem() *spectrumMenuItem {
	return &spectrumMenuItem{}
}

// Text sets the text content of the menu item
func (m *spectrumMenuItem) Text(text string) *spectrumMenuItem {
	m.PropText = text
	return m
}

// Value sets the value attribute of the menu item
func (m *spectrumMenuItem) Value(value string) *spectrumMenuItem {
	m.PropValue = value
	return m
}

// Active sets whether the menu item is active or has an active descendant
func (m *spectrumMenuItem) Active(active bool) *spectrumMenuItem {
	m.PropActive = active
	return m
}

// Disabled sets whether the menu item is disabled
func (m *spectrumMenuItem) Disabled(disabled bool) *spectrumMenuItem {
	m.PropDisabled = disabled
	return m
}

// Download sets the download attribute that causes the browser to treat the linked URL as a download
func (m *spectrumMenuItem) Download(download string) *spectrumMenuItem {
	m.PropDownload = download
	return m
}

// Focused sets whether the menu item has keyboard focus
func (m *spectrumMenuItem) Focused(focused bool) *spectrumMenuItem {
	m.PropFocused = focused
	return m
}

// HasSubmenu sets whether the menu item has a submenu
func (m *spectrumMenuItem) HasSubmenu(hasSubmenu bool) *spectrumMenuItem {
	m.PropHasSubmenu = hasSubmenu
	return m
}

// Href sets the URL that the hyperlink points to
func (m *spectrumMenuItem) Href(href string) *spectrumMenuItem {
	m.PropHref = href
	return m
}

// Label sets an accessible label that describes the component
func (m *spectrumMenuItem) Label(label string) *spectrumMenuItem {
	m.PropLabel = label
	return m
}

// NoWrap sets whether menu item text content should not wrap
func (m *spectrumMenuItem) NoWrap(noWrap bool) *spectrumMenuItem {
	m.PropNoWrap = noWrap
	return m
}

// Open sets whether the submenu is open
func (m *spectrumMenuItem) Open(open bool) *spectrumMenuItem {
	m.PropOpen = open
	return m
}

// Selected sets whether the menu item is selected
func (m *spectrumMenuItem) Selected(selected bool) *spectrumMenuItem {
	m.PropSelected = selected
	return m
}

// Invalid sets whether the menu item is invalid
func (m *spectrumMenuItem) Invalid(invalid bool) *spectrumMenuItem {
	m.PropInvalid = invalid
	return m
}

// Content sets the content of the menu item
func (m *spectrumMenuItem) Content(content app.UI) *spectrumMenuItem {
	m.PropContent = content
	return m
}

// Description sets content for the description slot
func (m *spectrumMenuItem) Description(description app.UI) *spectrumMenuItem {
	m.PropDescription = description
	return m
}

// Icon sets content for the icon slot
func (m *spectrumMenuItem) Icon(icon app.UI) *spectrumMenuItem {
	m.PropIcon = icon
	return m
}

// Submenu sets content for the submenu slot
func (m *spectrumMenuItem) Submenu(submenu app.UI) *spectrumMenuItem {
	m.PropSubmenu = submenu
	return m
}

// ValueSlot sets content for the value slot (displayed at the end of the menu item)
func (m *spectrumMenuItem) ValueSlot(value app.UI) *spectrumMenuItem {
	m.PropValueSlot = value
	return m
}

// OnClick sets the click event handler
func (m *spectrumMenuItem) OnClick(handler app.EventHandler) *spectrumMenuItem {
	m.PropOnClick = handler
	return m
}

// OnBlur sets the blur event handler
func (m *spectrumMenuItem) OnBlur(handler app.EventHandler) *spectrumMenuItem {
	m.PropOnBlur = handler
	return m
}

// OnFocus sets the focus event handler
func (m *spectrumMenuItem) OnFocus(handler app.EventHandler) *spectrumMenuItem {
	m.PropOnFocus = handler
	return m
}

// Render renders the menu item component
func (m *spectrumMenuItem) Render() app.UI {
	menuItem := app.Elem("sp-menu-item")

	// Set attributes
	if m.PropValue != "" {
		menuItem = menuItem.Attr("value", m.PropValue)
	}
	if m.PropActive {
		menuItem = menuItem.Attr("active", true)
	}
	if m.PropDisabled {
		menuItem = menuItem.Attr("disabled", true)
	}
	if m.PropDownload != "" {
		menuItem = menuItem.Attr("download", m.PropDownload)
	}
	if m.PropFocused {
		menuItem = menuItem.Attr("focused", true)
	}
	if m.PropHasSubmenu {
		menuItem = menuItem.Attr("has-submenu", true)
	}
	if m.PropHref != "" {
		menuItem = menuItem.Attr("href", m.PropHref)
	}
	if m.PropLabel != "" {
		menuItem = menuItem.Attr("label", m.PropLabel)
	}
	if m.PropNoWrap {
		menuItem = menuItem.Attr("no-wrap", true)
	}
	if m.PropOpen {
		menuItem = menuItem.Attr("open", true)
	}
	if m.PropSelected {
		menuItem = menuItem.Attr("selected", true)
	}
	if m.PropInvalid {
		menuItem = menuItem.Attr("invalid", true)
	}

	// Add event handlers
	if m.PropOnClick != nil {
		menuItem = menuItem.On("click", m.PropOnClick)
	}
	if m.PropOnBlur != nil {
		menuItem = menuItem.On("blur", m.PropOnBlur)
	}
	if m.PropOnFocus != nil {
		menuItem = menuItem.On("focus", m.PropOnFocus)
	}

	// Create elements array for content and slots
	elements := []app.UI{}

	// Add text content
	if m.PropText != "" {
		elements = append(elements, app.Text(m.PropText))
	}

	// Add content UI if provided
	if m.PropContent != nil {
		elements = append(elements, m.PropContent)
	}

	// Add description slot
	if m.PropDescription != nil {
		description := m.PropDescription
		if descWithSlot, ok := description.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, descWithSlot.Slot("description"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "description").
				Body(description))
		}
	}

	// Add icon slot
	if m.PropIcon != nil {
		icon := m.PropIcon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, iconWithSlot.Slot("icon"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "icon").
				Body(icon))
		}
	}

	// Add submenu slot
	if m.PropSubmenu != nil {
		submenu := m.PropSubmenu
		if submenuWithSlot, ok := submenu.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, submenuWithSlot.Slot("submenu"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "submenu").
				Body(submenu))
		}
	}

	// Add value slot
	if m.PropValueSlot != nil {
		valueSlot := m.PropValueSlot
		if valueWithSlot, ok := valueSlot.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, valueWithSlot.Slot("value"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "value").
				Body(valueSlot))
		}
	}

	// Add all elements to the menu item
	if len(elements) > 0 {
		menuItem = menuItem.Body(elements...)
	}

	return menuItem
}
