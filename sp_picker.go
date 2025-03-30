package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerIcons represents the How icons are displayed
type PickerIcons string

// PickerIcons values
const (
    PickerIconsOnly PickerIcons = "only"
    PickerIconsNone PickerIcons = "none"
)
// PickerPlacement represents the Positioning of the dropdown menu
type PickerPlacement string

// PickerPlacement values
const (
    PickerPlacementTop PickerPlacement = "top"
    PickerPlacementTopStart PickerPlacement = "top-start"
    PickerPlacementTopEnd PickerPlacement = "top-end"
    PickerPlacementRight PickerPlacement = "right"
    PickerPlacementRightStart PickerPlacement = "right-start"
    PickerPlacementRightEnd PickerPlacement = "right-end"
    PickerPlacementBottom PickerPlacement = "bottom"
    PickerPlacementBottomStart PickerPlacement = "bottom-start"
    PickerPlacementBottomEnd PickerPlacement = "bottom-end"
    PickerPlacementLeft PickerPlacement = "left"
    PickerPlacementLeftStart PickerPlacement = "left-start"
    PickerPlacementLeftEnd PickerPlacement = "left-end"
)

// spectrumPicker represents an sp-picker component
type spectrumPicker struct {
    app.Compo
    *styler[*spectrumPicker]

    // Properties
    // Disable this control
    PropDisabled bool
    // Forces rendering as popover on mobile instead of tray
    PropForcepopover bool
    // How icons are displayed
    PropIcons PickerIcons
    // Whether the picker is invalid
    PropInvalid bool
    // Accessible label when visible label not provided
    PropLabel string
    // Whether the menu is open
    PropOpen bool
    // Whether items are currently loading
    PropPending bool
    // Label for the picker in pending state
    PropPendinglabel string
    // Positioning of the dropdown menu
    PropPlacement PickerPlacement
    // Display without visible background
    PropQuiet bool
    // Whether user can interact with the value
    PropReadonly bool
    // The value of the selected option
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots
    PropDescriptionSlot app.UI
    PropLabelSlot app.UI
    PropTooltipSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnSpOpened app.EventHandler
}

// Picker creates a new sp-picker component
func Picker() *spectrumPicker {
    element := &spectrumPicker{
        PropDisabled: false,
        PropForcepopover: false,
        PropIcons: "",
        PropInvalid: false,
        PropLabel: "",
        PropOpen: false,
        PropPending: false,
        PropPendinglabel: "Pending",
        PropPlacement: PickerPlacementBottomStart,
        PropQuiet: false,
        PropReadonly: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Disable this control
func (c *spectrumPicker) Disabled(disabled bool) *spectrumPicker {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumPicker) SetDisabled() *spectrumPicker {
    return c.Disabled(true)
}

// Forces rendering as popover on mobile instead of tray
func (c *spectrumPicker) Forcepopover(forcePopover bool) *spectrumPicker {
    c.PropForcepopover = forcePopover
    return c
}

func (c *spectrumPicker) SetForcepopover() *spectrumPicker {
    return c.Forcepopover(true)
}

// How icons are displayed
func (c *spectrumPicker) Icons(icons PickerIcons) *spectrumPicker {
    c.PropIcons = icons
    return c
}

func (c *spectrumPicker) IconsOnly() *spectrumPicker {
    return c.Icons(PickerIconsOnly)
}
func (c *spectrumPicker) IconsNone() *spectrumPicker {
    return c.Icons(PickerIconsNone)
}
// Whether the picker is invalid
func (c *spectrumPicker) Invalid(invalid bool) *spectrumPicker {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumPicker) SetInvalid() *spectrumPicker {
    return c.Invalid(true)
}

// Accessible label when visible label not provided
func (c *spectrumPicker) Label(label string) *spectrumPicker {
    c.PropLabel = label
    return c
}

// Whether the menu is open
func (c *spectrumPicker) Open(open bool) *spectrumPicker {
    c.PropOpen = open
    return c
}

func (c *spectrumPicker) SetOpen() *spectrumPicker {
    return c.Open(true)
}

// Whether items are currently loading
func (c *spectrumPicker) Pending(pending bool) *spectrumPicker {
    c.PropPending = pending
    return c
}

func (c *spectrumPicker) SetPending() *spectrumPicker {
    return c.Pending(true)
}

// Label for the picker in pending state
func (c *spectrumPicker) Pendinglabel(pendingLabel string) *spectrumPicker {
    c.PropPendinglabel = pendingLabel
    return c
}

// Positioning of the dropdown menu
func (c *spectrumPicker) Placement(placement PickerPlacement) *spectrumPicker {
    c.PropPlacement = placement
    return c
}

func (c *spectrumPicker) PlacementTop() *spectrumPicker {
    return c.Placement(PickerPlacementTop)
}
func (c *spectrumPicker) PlacementTopStart() *spectrumPicker {
    return c.Placement(PickerPlacementTopStart)
}
func (c *spectrumPicker) PlacementTopEnd() *spectrumPicker {
    return c.Placement(PickerPlacementTopEnd)
}
func (c *spectrumPicker) PlacementRight() *spectrumPicker {
    return c.Placement(PickerPlacementRight)
}
func (c *spectrumPicker) PlacementRightStart() *spectrumPicker {
    return c.Placement(PickerPlacementRightStart)
}
func (c *spectrumPicker) PlacementRightEnd() *spectrumPicker {
    return c.Placement(PickerPlacementRightEnd)
}
func (c *spectrumPicker) PlacementBottom() *spectrumPicker {
    return c.Placement(PickerPlacementBottom)
}
func (c *spectrumPicker) PlacementBottomStart() *spectrumPicker {
    return c.Placement(PickerPlacementBottomStart)
}
func (c *spectrumPicker) PlacementBottomEnd() *spectrumPicker {
    return c.Placement(PickerPlacementBottomEnd)
}
func (c *spectrumPicker) PlacementLeft() *spectrumPicker {
    return c.Placement(PickerPlacementLeft)
}
func (c *spectrumPicker) PlacementLeftStart() *spectrumPicker {
    return c.Placement(PickerPlacementLeftStart)
}
func (c *spectrumPicker) PlacementLeftEnd() *spectrumPicker {
    return c.Placement(PickerPlacementLeftEnd)
}
// Display without visible background
func (c *spectrumPicker) Quiet(quiet bool) *spectrumPicker {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumPicker) SetQuiet() *spectrumPicker {
    return c.Quiet(true)
}

// Whether user can interact with the value
func (c *spectrumPicker) Readonly(readonly bool) *spectrumPicker {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumPicker) SetReadonly() *spectrumPicker {
    return c.Readonly(true)
}

// The value of the selected option
func (c *spectrumPicker) Value(value string) *spectrumPicker {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumPicker) Text(text string) *spectrumPicker {
    c.PropText = text
    return c
}


// Description content for the picker
func (c *spectrumPicker) Description(content app.UI) *spectrumPicker {
    c.PropDescriptionSlot = content

    return c
}

// Label content for the picker
func (c *spectrumPicker) LabelContent(content app.UI) *spectrumPicker {
    c.PropLabelSlot = content

    return c
}

// Tooltip for the picker button
func (c *spectrumPicker) Tooltip(content app.UI) *spectrumPicker {
    c.PropTooltipSlot = content

    return c
}



// Value of the element has changed
func (c *spectrumPicker) OnChange(handler app.EventHandler) *spectrumPicker {
    c.PropOnChange = handler

    return c
}

// Announces that the overlay has been opened
func (c *spectrumPicker) OnSpOpened(handler app.EventHandler) *spectrumPicker {
    c.PropOnSpOpened = handler

    return c
}


// Render renders the sp-picker component
func (c *spectrumPicker) Render() app.UI {
    element := app.Elem("sp-picker")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
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
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
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