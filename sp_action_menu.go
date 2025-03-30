package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionMenuIcons represents the Controls icon display mode: 'only' shows only icons, 'none' shows no icons
type ActionMenuIcons string

// ActionMenuIcons values
const (
    ActionMenuIconsOnly ActionMenuIcons = "only"
    ActionMenuIconsNone ActionMenuIcons = "none"
)
// ActionMenuPlacement represents the Position of the overlay relative to the action button
type ActionMenuPlacement string

// ActionMenuPlacement values
const (
    ActionMenuPlacementTop ActionMenuPlacement = "top"
    ActionMenuPlacementTopStart ActionMenuPlacement = "top-start"
    ActionMenuPlacementTopEnd ActionMenuPlacement = "top-end"
    ActionMenuPlacementRight ActionMenuPlacement = "right"
    ActionMenuPlacementRightStart ActionMenuPlacement = "right-start"
    ActionMenuPlacementRightEnd ActionMenuPlacement = "right-end"
    ActionMenuPlacementBottom ActionMenuPlacement = "bottom"
    ActionMenuPlacementBottomStart ActionMenuPlacement = "bottom-start"
    ActionMenuPlacementBottomEnd ActionMenuPlacement = "bottom-end"
    ActionMenuPlacementLeft ActionMenuPlacement = "left"
    ActionMenuPlacementLeftStart ActionMenuPlacement = "left-start"
    ActionMenuPlacementLeftEnd ActionMenuPlacement = "left-end"
)
// ActionMenuSelects represents the By default, action menu does not manage selection. Use 'single' to activate selection management.
type ActionMenuSelects string

// ActionMenuSelects values
const (
    ActionMenuSelectsSingle ActionMenuSelects = "single"
)
// ActionMenuSize represents the Size of the action menu button
type ActionMenuSize string

// ActionMenuSize values
const (
    ActionMenuSizeS ActionMenuSize = "s"
    ActionMenuSizeM ActionMenuSize = "m"
    ActionMenuSizeL ActionMenuSize = "l"
    ActionMenuSizeXl ActionMenuSize = "xl"
)
// ActionMenuStaticcolor represents the Applies a specific color treatment to the button
type ActionMenuStaticcolor string

// ActionMenuStaticcolor values
const (
    ActionMenuStaticcolorWhite ActionMenuStaticcolor = "white"
    ActionMenuStaticcolorBlack ActionMenuStaticcolor = "black"
)

// spectrumActionMenu represents an sp-action-menu component
type spectrumActionMenu struct {
    app.Compo
    *styler[*spectrumActionMenu]

    // Properties
    // Disables the action menu
    PropDisabled bool
    // Whether the action menu has focus
    PropFocused bool
    // Forces the action menu to render as a popover on mobile instead of a tray
    PropForcepopover bool
    // Controls icon display mode: 'only' shows only icons, 'none' shows no icons
    PropIcons ActionMenuIcons
    // Marks the action menu as invalid
    PropInvalid bool
    // Accessible label for the action menu, particularly important when no visible label is provided
    PropLabel string
    // Whether the action menu overlay is open
    PropOpen bool
    // Whether the items are currently loading
    PropPending bool
    // Defines a string value that labels the action menu while it is in pending state
    PropPendinglabel string
    // Position of the overlay relative to the action button
    PropPlacement ActionMenuPlacement
    // Applies a less visually prominent style
    PropQuiet bool
    // Makes the action menu read-only
    PropReadonly bool
    // By default, action menu does not manage selection. Use 'single' to activate selection management.
    PropSelects ActionMenuSelects
    // Size of the action menu button
    PropSize ActionMenuSize
    // Applies a specific color treatment to the button
    PropStaticcolor ActionMenuStaticcolor
    // The current value of the action menu when selection is enabled
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots
    PropDescriptionSlot app.UI
    PropIconSlot app.UI
    PropLabelSlot app.UI
    PropLabelOnlySlot app.UI
    PropTooltipSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnScroll app.EventHandler
    PropOnSpOpened app.EventHandler
}

// ActionMenu creates a new sp-action-menu component
func ActionMenu() *spectrumActionMenu {
    element := &spectrumActionMenu{
        PropDisabled: false,
        PropFocused: false,
        PropForcepopover: false,
        PropInvalid: false,
        PropOpen: false,
        PropPending: false,
        PropPendinglabel: "Pending",
        PropPlacement: ActionMenuPlacementBottomStart,
        PropQuiet: false,
        PropReadonly: false,
        PropSelects: "",
        PropSize: ActionMenuSizeM,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Disables the action menu
func (c *spectrumActionMenu) Disabled(disabled bool) *spectrumActionMenu {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumActionMenu) SetDisabled() *spectrumActionMenu {
    return c.Disabled(true)
}

// Whether the action menu has focus
func (c *spectrumActionMenu) Focused(focused bool) *spectrumActionMenu {
    c.PropFocused = focused
    return c
}

func (c *spectrumActionMenu) SetFocused() *spectrumActionMenu {
    return c.Focused(true)
}

// Forces the action menu to render as a popover on mobile instead of a tray
func (c *spectrumActionMenu) Forcepopover(forcePopover bool) *spectrumActionMenu {
    c.PropForcepopover = forcePopover
    return c
}

func (c *spectrumActionMenu) SetForcepopover() *spectrumActionMenu {
    return c.Forcepopover(true)
}

// Controls icon display mode: 'only' shows only icons, 'none' shows no icons
func (c *spectrumActionMenu) Icons(icons ActionMenuIcons) *spectrumActionMenu {
    c.PropIcons = icons
    return c
}

func (c *spectrumActionMenu) IconsOnly() *spectrumActionMenu {
    return c.Icons(ActionMenuIconsOnly)
}
func (c *spectrumActionMenu) IconsNone() *spectrumActionMenu {
    return c.Icons(ActionMenuIconsNone)
}
// Marks the action menu as invalid
func (c *spectrumActionMenu) Invalid(invalid bool) *spectrumActionMenu {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumActionMenu) SetInvalid() *spectrumActionMenu {
    return c.Invalid(true)
}

// Accessible label for the action menu, particularly important when no visible label is provided
func (c *spectrumActionMenu) Label(label string) *spectrumActionMenu {
    c.PropLabel = label
    return c
}

// Whether the action menu overlay is open
func (c *spectrumActionMenu) Open(open bool) *spectrumActionMenu {
    c.PropOpen = open
    return c
}

func (c *spectrumActionMenu) SetOpen() *spectrumActionMenu {
    return c.Open(true)
}

// Whether the items are currently loading
func (c *spectrumActionMenu) Pending(pending bool) *spectrumActionMenu {
    c.PropPending = pending
    return c
}

func (c *spectrumActionMenu) SetPending() *spectrumActionMenu {
    return c.Pending(true)
}

// Defines a string value that labels the action menu while it is in pending state
func (c *spectrumActionMenu) Pendinglabel(pendingLabel string) *spectrumActionMenu {
    c.PropPendinglabel = pendingLabel
    return c
}

// Position of the overlay relative to the action button
func (c *spectrumActionMenu) Placement(placement ActionMenuPlacement) *spectrumActionMenu {
    c.PropPlacement = placement
    return c
}

func (c *spectrumActionMenu) PlacementTop() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementTop)
}
func (c *spectrumActionMenu) PlacementTopStart() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementTopStart)
}
func (c *spectrumActionMenu) PlacementTopEnd() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementTopEnd)
}
func (c *spectrumActionMenu) PlacementRight() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementRight)
}
func (c *spectrumActionMenu) PlacementRightStart() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementRightStart)
}
func (c *spectrumActionMenu) PlacementRightEnd() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementRightEnd)
}
func (c *spectrumActionMenu) PlacementBottom() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementBottom)
}
func (c *spectrumActionMenu) PlacementBottomStart() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementBottomStart)
}
func (c *spectrumActionMenu) PlacementBottomEnd() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementBottomEnd)
}
func (c *spectrumActionMenu) PlacementLeft() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementLeft)
}
func (c *spectrumActionMenu) PlacementLeftStart() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementLeftStart)
}
func (c *spectrumActionMenu) PlacementLeftEnd() *spectrumActionMenu {
    return c.Placement(ActionMenuPlacementLeftEnd)
}
// Applies a less visually prominent style
func (c *spectrumActionMenu) Quiet(quiet bool) *spectrumActionMenu {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumActionMenu) SetQuiet() *spectrumActionMenu {
    return c.Quiet(true)
}

// Makes the action menu read-only
func (c *spectrumActionMenu) Readonly(readonly bool) *spectrumActionMenu {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumActionMenu) SetReadonly() *spectrumActionMenu {
    return c.Readonly(true)
}

// By default, action menu does not manage selection. Use 'single' to activate selection management.
func (c *spectrumActionMenu) Selects(selects ActionMenuSelects) *spectrumActionMenu {
    c.PropSelects = selects
    return c
}

func (c *spectrumActionMenu) SelectsSingle() *spectrumActionMenu {
    return c.Selects(ActionMenuSelectsSingle)
}
// Size of the action menu button
func (c *spectrumActionMenu) Size(size ActionMenuSize) *spectrumActionMenu {
    c.PropSize = size
    return c
}

func (c *spectrumActionMenu) SizeS() *spectrumActionMenu {
    return c.Size(ActionMenuSizeS)
}
func (c *spectrumActionMenu) SizeM() *spectrumActionMenu {
    return c.Size(ActionMenuSizeM)
}
func (c *spectrumActionMenu) SizeL() *spectrumActionMenu {
    return c.Size(ActionMenuSizeL)
}
func (c *spectrumActionMenu) SizeXl() *spectrumActionMenu {
    return c.Size(ActionMenuSizeXl)
}
// Applies a specific color treatment to the button
func (c *spectrumActionMenu) Staticcolor(staticColor ActionMenuStaticcolor) *spectrumActionMenu {
    c.PropStaticcolor = staticColor
    return c
}

func (c *spectrumActionMenu) StaticcolorWhite() *spectrumActionMenu {
    return c.Staticcolor(ActionMenuStaticcolorWhite)
}
func (c *spectrumActionMenu) StaticcolorBlack() *spectrumActionMenu {
    return c.Staticcolor(ActionMenuStaticcolorBlack)
}
// The current value of the action menu when selection is enabled
func (c *spectrumActionMenu) Value(value string) *spectrumActionMenu {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumActionMenu) Text(text string) *spectrumActionMenu {
    c.PropText = text
    return c
}


// The description content for the action menu
func (c *spectrumActionMenu) Description(content app.UI) *spectrumActionMenu {
    c.PropDescriptionSlot = content

    return c
}

// The icon to use for the Action Menu instead of the default
func (c *spectrumActionMenu) Icon(content app.UI) *spectrumActionMenu {
    c.PropIconSlot = content

    return c
}

// The label to use for the Action Menu (with icon space reserved)
func (c *spectrumActionMenu) LabelContent(content app.UI) *spectrumActionMenu {
    c.PropLabelSlot = content

    return c
}

// The label to use for the Action Menu (no icon space reserved)
func (c *spectrumActionMenu) LabelOnly(content app.UI) *spectrumActionMenu {
    c.PropLabelOnlySlot = content

    return c
}

// Tooltip to be applied to the Action Button
func (c *spectrumActionMenu) Tooltip(content app.UI) *spectrumActionMenu {
    c.PropTooltipSlot = content

    return c
}



// Announces that the value of the element has changed
func (c *spectrumActionMenu) OnChange(handler app.EventHandler) *spectrumActionMenu {
    c.PropOnChange = handler

    return c
}

// Fired when the menu content is scrolled
func (c *spectrumActionMenu) OnScroll(handler app.EventHandler) *spectrumActionMenu {
    c.PropOnScroll = handler

    return c
}

// Announces that the overlay has been opened
func (c *spectrumActionMenu) OnSpOpened(handler app.EventHandler) *spectrumActionMenu {
    c.PropOnSpOpened = handler

    return c
}


// Render renders the sp-action-menu component
func (c *spectrumActionMenu) Render() app.UI {
    element := app.Elem("sp-action-menu")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFocused {
        element = element.Attr("focused", true)
    }
    if c.PropForcepopover {
        element = element.Attr("forcePopover", true)
    }
    if c.PropIcons != "" {
        element = element.Attr("icons", string(c.PropIcons))
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropPending {
        element = element.Attr("pending", true)
    }
    if c.PropPendinglabel != "" {
        element = element.Attr("pendingLabel", c.PropPendinglabel)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropSelects != "" {
        element = element.Attr("selects", string(c.PropSelects))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropStaticcolor != "" {
        element = element.Attr("staticColor", string(c.PropStaticcolor))
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnScroll != nil {
        element = element.On("scroll", c.PropOnScroll)
    }
    if c.PropOnSpOpened != nil {
        element = element.On("sp-opened", c.PropOnSpOpened)
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
    // Add label slot
    if c.PropLabelSlot != nil {
        slotElem := c.PropLabelSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("label")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "label").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add label-only slot
    if c.PropLabelOnlySlot != nil {
        slotElem := c.PropLabelOnlySlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("label-only")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "label-only").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add tooltip slot
    if c.PropTooltipSlot != nil {
        slotElem := c.PropTooltipSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("tooltip")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "tooltip").
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