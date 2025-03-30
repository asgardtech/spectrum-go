package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ComboboxAutocomplete represents the Type of assistance when supplying values
type ComboboxAutocomplete string

// ComboboxAutocomplete values
const (
    ComboboxAutocompleteList ComboboxAutocomplete = "list"
    ComboboxAutocompleteNone ComboboxAutocomplete = "none"
)

// spectrumCombobox represents an sp-combobox component
type spectrumCombobox struct {
    app.Compo
    *styler[*spectrumCombobox]

    // Properties
    // Type of assistance when supplying values
    PropAutocomplete ComboboxAutocomplete
    // Disable this control
    PropDisabled bool
    // Whether the value is invalid
    PropInvalid bool
    // Accessible label when visible label not provided
    PropLabel string
    // Whether the listbox is visible
    PropOpen bool
    // Whether items are currently loading
    PropPending bool
    // Label for the combobox in pending state
    PropPendinglabel string
    // Text shown when no value is set
    PropPlaceholder string
    // Display without visible background
    PropQuiet bool
    // Whether field is required
    PropRequired bool
    // The value held by the control
    PropValue string

    // Text content for the default slot
    PropText string

    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI
    PropTooltipSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
}

// Combobox creates a new sp-combobox component
func Combobox() *spectrumCombobox {
    element := &spectrumCombobox{
        PropAutocomplete: ComboboxAutocompleteNone,
        PropDisabled: false,
        PropInvalid: false,
        PropLabel: "",
        PropOpen: false,
        PropPending: false,
        PropPendinglabel: "Pending",
        PropPlaceholder: "",
        PropQuiet: false,
        PropRequired: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// Type of assistance when supplying values
func (c *spectrumCombobox) Autocomplete(autocomplete ComboboxAutocomplete) *spectrumCombobox {
    c.PropAutocomplete = autocomplete
    return c
}

func (c *spectrumCombobox) AutocompleteList() *spectrumCombobox {
    return c.Autocomplete(ComboboxAutocompleteList)
}
func (c *spectrumCombobox) AutocompleteNone() *spectrumCombobox {
    return c.Autocomplete(ComboboxAutocompleteNone)
}
// Disable this control
func (c *spectrumCombobox) Disabled(disabled bool) *spectrumCombobox {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumCombobox) SetDisabled() *spectrumCombobox {
    return c.Disabled(true)
}

// Whether the value is invalid
func (c *spectrumCombobox) Invalid(invalid bool) *spectrumCombobox {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumCombobox) SetInvalid() *spectrumCombobox {
    return c.Invalid(true)
}

// Accessible label when visible label not provided
func (c *spectrumCombobox) Label(label string) *spectrumCombobox {
    c.PropLabel = label
    return c
}

// Whether the listbox is visible
func (c *spectrumCombobox) Open(open bool) *spectrumCombobox {
    c.PropOpen = open
    return c
}

func (c *spectrumCombobox) SetOpen() *spectrumCombobox {
    return c.Open(true)
}

// Whether items are currently loading
func (c *spectrumCombobox) Pending(pending bool) *spectrumCombobox {
    c.PropPending = pending
    return c
}

func (c *spectrumCombobox) SetPending() *spectrumCombobox {
    return c.Pending(true)
}

// Label for the combobox in pending state
func (c *spectrumCombobox) Pendinglabel(pendingLabel string) *spectrumCombobox {
    c.PropPendinglabel = pendingLabel
    return c
}

// Text shown when no value is set
func (c *spectrumCombobox) Placeholder(placeholder string) *spectrumCombobox {
    c.PropPlaceholder = placeholder
    return c
}

// Display without visible background
func (c *spectrumCombobox) Quiet(quiet bool) *spectrumCombobox {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumCombobox) SetQuiet() *spectrumCombobox {
    return c.Quiet(true)
}

// Whether field is required
func (c *spectrumCombobox) Required(required bool) *spectrumCombobox {
    c.PropRequired = required
    return c
}

func (c *spectrumCombobox) SetRequired() *spectrumCombobox {
    return c.Required(true)
}

// The value held by the control
func (c *spectrumCombobox) Value(value string) *spectrumCombobox {
    c.PropValue = value
    return c
}


// Text sets the text content for the default slot
func (c *spectrumCombobox) Text(text string) *spectrumCombobox {
    c.PropText = text
    return c
}


// Default or non-negative help text
func (c *spectrumCombobox) HelpText(content app.UI) *spectrumCombobox {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text shown when invalid
func (c *spectrumCombobox) NegativeHelpText(content app.UI) *spectrumCombobox {
    c.PropNegativeHelpTextSlot = content

    return c
}

// Tooltip for the picker button
func (c *spectrumCombobox) Tooltip(content app.UI) *spectrumCombobox {
    c.PropTooltipSlot = content

    return c
}



// Value has been committed by the user
func (c *spectrumCombobox) OnChange(handler app.EventHandler) *spectrumCombobox {
    c.PropOnChange = handler

    return c
}

// Value of the element has changed
func (c *spectrumCombobox) OnInput(handler app.EventHandler) *spectrumCombobox {
    c.PropOnInput = handler

    return c
}


// Render renders the sp-combobox component
func (c *spectrumCombobox) Render() app.UI {
    element := app.Elem("sp-combobox")

    // Set attributes
    if c.PropAutocomplete != "" {
        element = element.Attr("autocomplete", string(c.PropAutocomplete))
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
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
    if c.PropPlaceholder != "" {
        element = element.Attr("placeholder", c.PropPlaceholder)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropRequired {
        element = element.Attr("required", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnInput != nil {
        element = element.On("input", c.PropOnInput)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add help-text slot
    if c.PropHelpTextSlot != nil {
        slotElem := c.PropHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "help-text").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add negative-help-text slot
    if c.PropNegativeHelpTextSlot != nil {
        slotElem := c.PropNegativeHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("negative-help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "negative-help-text").
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