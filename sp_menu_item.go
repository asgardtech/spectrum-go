// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MenuItemReferrerpolicy represents the How much of the referrer to send when following the link.
type MenuItemReferrerpolicy string

// MenuItemReferrerpolicy values
const (
	MenuItemReferrerpolicyNoReferrer                  MenuItemReferrerpolicy = "no-referrer"
	MenuItemReferrerpolicyNoReferrerWhenDowngrade     MenuItemReferrerpolicy = "no-referrer-when-downgrade"
	MenuItemReferrerpolicyOrigin                      MenuItemReferrerpolicy = "origin"
	MenuItemReferrerpolicyOriginWhenCrossOrigin       MenuItemReferrerpolicy = "origin-when-cross-origin"
	MenuItemReferrerpolicySameOrigin                  MenuItemReferrerpolicy = "same-origin"
	MenuItemReferrerpolicyStrictOrigin                MenuItemReferrerpolicy = "strict-origin"
	MenuItemReferrerpolicyStrictOriginWhenCrossOrigin MenuItemReferrerpolicy = "strict-origin-when-cross-origin"
	MenuItemReferrerpolicyUnsafeUrl                   MenuItemReferrerpolicy = "unsafe-url"
)

// MenuItemTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type MenuItemTarget string

// MenuItemTarget values
const (
	MenuItemTarget_blank  MenuItemTarget = "_blank"
	MenuItemTarget_parent MenuItemTarget = "_parent"
	MenuItemTarget_self   MenuItemTarget = "_self"
	MenuItemTarget_top    MenuItemTarget = "_top"
)

// spectrumMenuItem represents an sp-menu-item component
type spectrumMenuItem struct {
	app.Compo
	*styler[*spectrumMenuItem]
	*classer[*spectrumMenuItem]
	*ider[*spectrumMenuItem]

	// Properties
	// Whether the menu item is active or has an active descendant
	PropActive bool
	// Disable this control. It will not receive focus or events
	PropDisabled bool
	// Causes the browser to treat the linked URL as a download.
	PropDownload string
	// Whether the menu item has keyboard focus
	PropFocused bool
	// Whether the menu item has a submenu
	PropHasSubmenu bool
	// The URL that the hyperlink points to.
	PropHref string
	// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
	PropLabel string
	// Whether menu item text content should not wrap
	PropNoWrap bool
	// Whether submenu is open
	PropOpen bool
	// How much of the referrer to send when following the link.
	PropReferrerpolicy MenuItemReferrerpolicy
	// The relationship of the linked URL as space-separated link types.
	PropRel string
	// Whether the menu item is selected
	PropSelected bool
	// The tab index to apply to this control. See general documentation about the tabindex HTML property
	PropTabindex float64
	// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
	PropTarget MenuItemTarget
	// Value of the menu item which is used for selection
	PropValue string

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropDescriptionSlot app.UI
	PropIconSlot        app.UI
	PropSubmenuSlot     app.UI
	PropValueSlot       app.UI

	// Event handlers
	PropOnBlur            app.EventHandler
	PropOnFocus           app.EventHandler
	PropOnSpMenuItemAdded app.EventHandler
	PropOnClick           app.EventHandler
}

// IMenuItem is the interface for sp-menu-item component methods
type IMenuItem interface {
	app.UI
	Styler[IMenuItem]
	Classer[IMenuItem]
	Ider[IMenuItem]
	Active(bool) IMenuItem
	SetActive() IMenuItem
	Disabled(bool) IMenuItem
	SetDisabled() IMenuItem
	Download(string) IMenuItem
	Focused(bool) IMenuItem
	SetFocused() IMenuItem
	HasSubmenu(bool) IMenuItem
	SetHasSubmenu() IMenuItem
	Href(string) IMenuItem
	Label(string) IMenuItem
	NoWrap(bool) IMenuItem
	SetNoWrap() IMenuItem
	Open(bool) IMenuItem
	SetOpen() IMenuItem
	Referrerpolicy(MenuItemReferrerpolicy) IMenuItem
	ReferrerpolicyNoReferrer() IMenuItem
	ReferrerpolicyNoReferrerWhenDowngrade() IMenuItem
	ReferrerpolicyOrigin() IMenuItem
	ReferrerpolicyOriginWhenCrossOrigin() IMenuItem
	ReferrerpolicySameOrigin() IMenuItem
	ReferrerpolicyStrictOrigin() IMenuItem
	ReferrerpolicyStrictOriginWhenCrossOrigin() IMenuItem
	ReferrerpolicyUnsafeUrl() IMenuItem
	Rel(string) IMenuItem
	Selected(bool) IMenuItem
	SetSelected() IMenuItem
	Tabindex(float64) IMenuItem
	Target(MenuItemTarget) IMenuItem
	Target_blank() IMenuItem
	Target_parent() IMenuItem
	Target_self() IMenuItem
	Target_top() IMenuItem
	Value(string) IMenuItem

	Body(...app.UI) IMenuItem
	AddToBody(app.UI) IMenuItem
	Text(string) IMenuItem

	Description(app.UI) IMenuItem
	Icon(app.UI) IMenuItem
	Submenu(app.UI) IMenuItem
	ValueContent(app.UI) IMenuItem

	OnBlur(app.EventHandler) IMenuItem
	OnFocus(app.EventHandler) IMenuItem
	OnSpMenuItemAdded(app.EventHandler) IMenuItem
	OnClick(app.EventHandler) IMenuItem
}

// MenuItem For use within an sp-menu element, a menu-item represents a single item in a menu.
func MenuItem() IMenuItem {
	element := &spectrumMenuItem{
		PropActive:     false,
		PropDisabled:   false,
		PropFocused:    false,
		PropHasSubmenu: false,
		PropNoWrap:     false,
		PropOpen:       false,
		PropSelected:   false,
		PropTabindex:   0,
		PropBody:       []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Active Whether the menu item is active or has an active descendant
func (c *spectrumMenuItem) Active(active bool) IMenuItem {
	c.PropActive = active
	return c
}

func (c *spectrumMenuItem) SetActive() IMenuItem {
	return c.Active(true)
}

// Disabled Disable this control. It will not receive focus or events
func (c *spectrumMenuItem) Disabled(disabled bool) IMenuItem {
	c.PropDisabled = disabled
	return c
}

func (c *spectrumMenuItem) SetDisabled() IMenuItem {
	return c.Disabled(true)
}

// Download Causes the browser to treat the linked URL as a download.
func (c *spectrumMenuItem) Download(download string) IMenuItem {
	c.PropDownload = download
	return c
}

// Focused Whether the menu item has keyboard focus
func (c *spectrumMenuItem) Focused(focused bool) IMenuItem {
	c.PropFocused = focused
	return c
}

func (c *spectrumMenuItem) SetFocused() IMenuItem {
	return c.Focused(true)
}

// HasSubmenu Whether the menu item has a submenu
func (c *spectrumMenuItem) HasSubmenu(hasSubmenu bool) IMenuItem {
	c.PropHasSubmenu = hasSubmenu
	return c
}

func (c *spectrumMenuItem) SetHasSubmenu() IMenuItem {
	return c.HasSubmenu(true)
}

// Href The URL that the hyperlink points to.
func (c *spectrumMenuItem) Href(href string) IMenuItem {
	c.PropHref = href
	return c
}

// Label An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumMenuItem) Label(label string) IMenuItem {
	c.PropLabel = label
	return c
}

// NoWrap Whether menu item text content should not wrap
func (c *spectrumMenuItem) NoWrap(noWrap bool) IMenuItem {
	c.PropNoWrap = noWrap
	return c
}

func (c *spectrumMenuItem) SetNoWrap() IMenuItem {
	return c.NoWrap(true)
}

// Open Whether submenu is open
func (c *spectrumMenuItem) Open(open bool) IMenuItem {
	c.PropOpen = open
	return c
}

func (c *spectrumMenuItem) SetOpen() IMenuItem {
	return c.Open(true)
}

// Referrerpolicy How much of the referrer to send when following the link.
func (c *spectrumMenuItem) Referrerpolicy(referrerpolicy MenuItemReferrerpolicy) IMenuItem {
	c.PropReferrerpolicy = referrerpolicy
	return c
}

func (c *spectrumMenuItem) ReferrerpolicyNoReferrer() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyNoReferrer)
}
func (c *spectrumMenuItem) ReferrerpolicyNoReferrerWhenDowngrade() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumMenuItem) ReferrerpolicyOrigin() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyOriginWhenCrossOrigin() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicySameOrigin() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicySameOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyStrictOrigin() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyStrictOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyStrictOriginWhenCrossOrigin() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyUnsafeUrl() IMenuItem {
	return c.Referrerpolicy(MenuItemReferrerpolicyUnsafeUrl)
}

// Rel The relationship of the linked URL as space-separated link types.
func (c *spectrumMenuItem) Rel(rel string) IMenuItem {
	c.PropRel = rel
	return c
}

// Selected Whether the menu item is selected
func (c *spectrumMenuItem) Selected(selected bool) IMenuItem {
	c.PropSelected = selected
	return c
}

func (c *spectrumMenuItem) SetSelected() IMenuItem {
	return c.Selected(true)
}

// Tabindex The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumMenuItem) Tabindex(tabindex float64) IMenuItem {
	c.PropTabindex = tabindex
	return c
}

// Target Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumMenuItem) Target(target MenuItemTarget) IMenuItem {
	c.PropTarget = target
	return c
}

func (c *spectrumMenuItem) Target_blank() IMenuItem {
	return c.Target(MenuItemTarget_blank)
}
func (c *spectrumMenuItem) Target_parent() IMenuItem {
	return c.Target(MenuItemTarget_parent)
}
func (c *spectrumMenuItem) Target_self() IMenuItem {
	return c.Target(MenuItemTarget_self)
}
func (c *spectrumMenuItem) Target_top() IMenuItem {
	return c.Target(MenuItemTarget_top)
}

// Value Value of the menu item which is used for selection
func (c *spectrumMenuItem) Value(value string) IMenuItem {
	c.PropValue = value
	return c
}

// Body sets the content for the default slot
func (c *spectrumMenuItem) Body(elements ...app.UI) IMenuItem {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumMenuItem) AddToBody(element app.UI) IMenuItem {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumMenuItem) Text(text string) IMenuItem {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Description to be placed below the label of the Menu Item
func (c *spectrumMenuItem) Description(content app.UI) IMenuItem {
	c.PropDescriptionSlot = content

	return c
}

// Icon element to be placed at the start of the Menu Item
func (c *spectrumMenuItem) Icon(content app.UI) IMenuItem {
	c.PropIconSlot = content

	return c
}

// Content placed in a submenu
func (c *spectrumMenuItem) Submenu(content app.UI) IMenuItem {
	c.PropSubmenuSlot = content

	return c
}

// Content placed at the end of the Menu Item like values, keyboard shortcuts, etc.
func (c *spectrumMenuItem) ValueContent(content app.UI) IMenuItem {
	c.PropValueSlot = content

	return c
}

// Fired when the menu item loses focus
func (c *spectrumMenuItem) OnBlur(handler app.EventHandler) IMenuItem {
	c.PropOnBlur = handler

	return c
}

// Fired when the menu item receives focus
func (c *spectrumMenuItem) OnFocus(handler app.EventHandler) IMenuItem {
	c.PropOnFocus = handler

	return c
}

// Announces the item has been added so a parent menu can take ownership
func (c *spectrumMenuItem) OnSpMenuItemAdded(handler app.EventHandler) IMenuItem {
	c.PropOnSpMenuItemAdded = handler

	return c
}

// Fired when the menu item is clicked
func (c *spectrumMenuItem) OnClick(handler app.EventHandler) IMenuItem {
	c.PropOnClick = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumMenuItem) Style(key, format string, values ...any) IMenuItem {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumMenuItem) Styles(styles map[string]string) IMenuItem {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumMenuItem) Class(class string) IMenuItem {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumMenuItem) Classes(classes ...string) IMenuItem {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumMenuItem) Id(id string) IMenuItem {
	return c.ider.Id(id)
}

// Render renders the sp-menu-item component
func (c *spectrumMenuItem) Render() app.UI {
	element := app.Elem("sp-menu-item")

	// Set attributes
	if c.PropActive {
		element = element.Attr("active", true)
	}
	if c.PropDisabled {
		element = element.Attr("disabled", true)
	}
	if c.PropDownload != "" {
		element = element.Attr("download", c.PropDownload)
	}
	if c.PropFocused {
		element = element.Attr("focused", true)
	}
	if c.PropHasSubmenu {
		element = element.Attr("has-submenu", true)
	}
	if c.PropHref != "" {
		element = element.Attr("href", c.PropHref)
	}
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}
	if c.PropNoWrap {
		element = element.Attr("no-wrap", true)
	}
	if c.PropOpen {
		element = element.Attr("open", true)
	}
	if c.PropReferrerpolicy != "" {
		element = element.Attr("referrerpolicy", string(c.PropReferrerpolicy))
	}
	if c.PropRel != "" {
		element = element.Attr("rel", c.PropRel)
	}
	if c.PropSelected {
		element = element.Attr("selected", true)
	}
	if c.PropTabindex != 0 {
		element = element.Attr("tabindex", c.PropTabindex)
	}
	if c.PropTarget != "" {
		element = element.Attr("target", string(c.PropTarget))
	}
	if c.PropValue != "" {
		element = element.Attr("value", c.PropValue)
	}

	// Add event handlers
	if c.PropOnBlur != nil {
		element = element.On("blur", c.PropOnBlur)
	}
	if c.PropOnFocus != nil {
		element = element.On("focus", c.PropOnFocus)
	}
	if c.PropOnSpMenuItemAdded != nil {
		element = element.On("sp-menu-item-added", c.PropOnSpMenuItemAdded)
	}
	if c.PropOnClick != nil {
		element = element.On("click", c.PropOnClick)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add description slot
	if c.PropDescriptionSlot != nil {
		slotElem := c.PropDescriptionSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("description")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "description").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add icon slot
	if c.PropIconSlot != nil {
		slotElem := c.PropIconSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("icon")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "icon").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add submenu slot
	if c.PropSubmenuSlot != nil {
		slotElem := c.PropSubmenuSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("submenu")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "submenu").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add value slot
	if c.PropValueSlot != nil {
		slotElem := c.PropValueSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("value")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "value").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}

	// Add all elements to the component
	if len(slotElements) > 0 {
		element = element.Body(slotElements...)
	}

	// Apply styles, classes, and id
	element = element.Styles(c.styler.styles)

	// Apply classes if any
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}

	// Apply id if set
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}

	return element
}
