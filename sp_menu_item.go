package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MenuItemReferrerpolicy represents the How much of the referrer to send when following the link.
type MenuItemReferrerpolicy string

// MenuItemReferrerpolicy values
const (
    MenuItemReferrerpolicyNoReferrer MenuItemReferrerpolicy = "no-referrer"
    MenuItemReferrerpolicyNoReferrerWhenDowngrade MenuItemReferrerpolicy = "no-referrer-when-downgrade"
    MenuItemReferrerpolicyOrigin MenuItemReferrerpolicy = "origin"
    MenuItemReferrerpolicyOriginWhenCrossOrigin MenuItemReferrerpolicy = "origin-when-cross-origin"
    MenuItemReferrerpolicySameOrigin MenuItemReferrerpolicy = "same-origin"
    MenuItemReferrerpolicyStrictOrigin MenuItemReferrerpolicy = "strict-origin"
    MenuItemReferrerpolicyStrictOriginWhenCrossOrigin MenuItemReferrerpolicy = "strict-origin-when-cross-origin"
    MenuItemReferrerpolicyUnsafeUrl MenuItemReferrerpolicy = "unsafe-url"
)
// MenuItemTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
type MenuItemTarget string

// MenuItemTarget values
const (
    MenuItemTarget_blank MenuItemTarget = "_blank"
    MenuItemTarget_parent MenuItemTarget = "_parent"
    MenuItemTarget_self MenuItemTarget = "_self"
    MenuItemTarget_top MenuItemTarget = "_top"
)

// spectrumMenuItem represents an sp-menu-item component
type spectrumMenuItem struct {
    app.Compo
    *styler[*spectrumMenuItem]

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

    // Text content for the default slot
    PropText string

    // Content slots
    PropDescriptionSlot app.UI
    PropIconSlot app.UI
    PropSubmenuSlot app.UI
    PropValueSlot app.UI


    // Event handlers
    PropOnBlur app.EventHandler
    PropOnFocus app.EventHandler
    PropOnSpMenuItemAdded app.EventHandler
}

// MenuItem creates a new sp-menu-item component
func MenuItem() *spectrumMenuItem {
    element := &spectrumMenuItem{
        PropActive: false,
        PropDisabled: false,
        PropFocused: false,
        PropHasSubmenu: false,
        PropNoWrap: false,
        PropOpen: false,
        PropSelected: false,
        PropTabindex: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the menu item is active or has an active descendant
func (c *spectrumMenuItem) Active(active bool) *spectrumMenuItem {
    c.PropActive = active
    return c
}

func (c *spectrumMenuItem) SetActive() *spectrumMenuItem {
    return c.Active(true)
}

// Disable this control. It will not receive focus or events
func (c *spectrumMenuItem) Disabled(disabled bool) *spectrumMenuItem {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumMenuItem) SetDisabled() *spectrumMenuItem {
    return c.Disabled(true)
}

// Causes the browser to treat the linked URL as a download.
func (c *spectrumMenuItem) Download(download string) *spectrumMenuItem {
    c.PropDownload = download
    return c
}

// Whether the menu item has keyboard focus
func (c *spectrumMenuItem) Focused(focused bool) *spectrumMenuItem {
    c.PropFocused = focused
    return c
}

func (c *spectrumMenuItem) SetFocused() *spectrumMenuItem {
    return c.Focused(true)
}

// Whether the menu item has a submenu
func (c *spectrumMenuItem) HasSubmenu(hasSubmenu bool) *spectrumMenuItem {
    c.PropHasSubmenu = hasSubmenu
    return c
}

func (c *spectrumMenuItem) SetHasSubmenu() *spectrumMenuItem {
    return c.HasSubmenu(true)
}

// The URL that the hyperlink points to.
func (c *spectrumMenuItem) Href(href string) *spectrumMenuItem {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumMenuItem) Label(label string) *spectrumMenuItem {
    c.PropLabel = label
    return c
}

// Whether menu item text content should not wrap
func (c *spectrumMenuItem) NoWrap(noWrap bool) *spectrumMenuItem {
    c.PropNoWrap = noWrap
    return c
}

func (c *spectrumMenuItem) SetNoWrap() *spectrumMenuItem {
    return c.NoWrap(true)
}

// Whether submenu is open
func (c *spectrumMenuItem) Open(open bool) *spectrumMenuItem {
    c.PropOpen = open
    return c
}

func (c *spectrumMenuItem) SetOpen() *spectrumMenuItem {
    return c.Open(true)
}

// How much of the referrer to send when following the link.
func (c *spectrumMenuItem) Referrerpolicy(referrerpolicy MenuItemReferrerpolicy) *spectrumMenuItem {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumMenuItem) ReferrerpolicyNoReferrer() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyNoReferrer)
}
func (c *spectrumMenuItem) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumMenuItem) ReferrerpolicyOrigin() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyOriginWhenCrossOrigin() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicySameOrigin() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicySameOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyStrictOrigin() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyStrictOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumMenuItem) ReferrerpolicyUnsafeUrl() *spectrumMenuItem {
    return c.Referrerpolicy(MenuItemReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types.
func (c *spectrumMenuItem) Rel(rel string) *spectrumMenuItem {
    c.PropRel = rel
    return c
}

// Whether the menu item is selected
func (c *spectrumMenuItem) Selected(selected bool) *spectrumMenuItem {
    c.PropSelected = selected
    return c
}

func (c *spectrumMenuItem) SetSelected() *spectrumMenuItem {
    return c.Selected(true)
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumMenuItem) Tabindex(tabindex float64) *spectrumMenuItem {
    c.PropTabindex = tabindex
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe).
func (c *spectrumMenuItem) Target(target MenuItemTarget) *spectrumMenuItem {
    c.PropTarget = target
    return c
}

func (c *spectrumMenuItem) Target_blank() *spectrumMenuItem {
    return c.Target(MenuItemTarget_blank)
}
func (c *spectrumMenuItem) Target_parent() *spectrumMenuItem {
    return c.Target(MenuItemTarget_parent)
}
func (c *spectrumMenuItem) Target_self() *spectrumMenuItem {
    return c.Target(MenuItemTarget_self)
}
func (c *spectrumMenuItem) Target_top() *spectrumMenuItem {
    return c.Target(MenuItemTarget_top)
}
// Value of the menu item which is used for selection
func (c *spectrumMenuItem) Value(value string) *spectrumMenuItem {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumMenuItem) Text(text string) *spectrumMenuItem {
    c.PropText = text
    return c
}


// Description to be placed below the label of the Menu Item
func (c *spectrumMenuItem) Description(content app.UI) *spectrumMenuItem {
    c.PropDescriptionSlot = content

    return c
}

// Icon element to be placed at the start of the Menu Item
func (c *spectrumMenuItem) Icon(content app.UI) *spectrumMenuItem {
    c.PropIconSlot = content

    return c
}

// Content placed in a submenu
func (c *spectrumMenuItem) Submenu(content app.UI) *spectrumMenuItem {
    c.PropSubmenuSlot = content

    return c
}

// Content placed at the end of the Menu Item like values, keyboard shortcuts, etc.
func (c *spectrumMenuItem) ValueContent(content app.UI) *spectrumMenuItem {
    c.PropValueSlot = content

    return c
}



// Fired when the menu item loses focus
func (c *spectrumMenuItem) OnBlur(handler app.EventHandler) *spectrumMenuItem {
    c.PropOnBlur = handler

    return c
}

// Fired when the menu item receives focus
func (c *spectrumMenuItem) OnFocus(handler app.EventHandler) *spectrumMenuItem {
    c.PropOnFocus = handler

    return c
}

// Announces the item has been added so a parent menu can take ownership
func (c *spectrumMenuItem) OnSpMenuItemAdded(handler app.EventHandler) *spectrumMenuItem {
    c.PropOnSpMenuItemAdded = handler

    return c
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

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
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

    element = element.Styles(c.styler.styles)

    return element
} 