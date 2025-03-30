package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SearchSize represents the Size variant of the search field
type SearchSize string

// SearchSize values
const (
    SearchSizeS SearchSize = "s"
    SearchSizeM SearchSize = "m"
    SearchSizeL SearchSize = "l"
    SearchSizeXl SearchSize = "xl"
)

// spectrumSearch represents an sp-search component
type spectrumSearch struct {
    app.Compo
    *styler[*spectrumSearch]

    // Properties
    // The action URL for the search form
    PropAction string
    // What form of assistance should be provided when attempting to supply a value to the form control
    PropAutocomplete string
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether a form control delivered with the `multiline` attribute will change size vertically to accommodate longer input
    PropGrows bool
    // Controls whether the typed value should be held when the Escape key is pressed
    PropHoldvalueonescape bool
    // Whether the `value` held by the form control is invalid
    PropInvalid bool
    // A string applied via `aria-label` to the form control when a user visible label is not provided
    PropLabel string
    // Defines the maximum string length that the user can enter
    PropMaxlength float64
    // HTTP method to use for the form submission
    PropMethod string
    // Defines the minimum string length that the user can enter
    PropMinlength float64
    // Whether the form control should accept a value longer than one line
    PropMultiline bool
    // Name of the form control
    PropName string
    // Pattern the `value` must match to be valid
    PropPattern string
    // Text that appears in the form control when it has no value set
    PropPlaceholder string
    // Whether to display the form control with no visible background
    PropQuiet bool
    // Whether a user can interact with the value of the form control
    PropReadonly bool
    // Whether the form control will be found to be invalid when it holds no `value`
    PropRequired bool
    // The specific number of rows the form control should provide in the user interface
    PropRows float64
    // Size variant of the search field
    PropSize SearchSize
    // The tab index to apply to this control
    PropTabindex float64
    // Whether the `value` held by the form control is valid
    PropValid bool
    // The value held by the form control
    PropValue string


    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
    PropOnSubmit app.EventHandler
}

// Search creates a new sp-search component
func Search() *spectrumSearch {
    element := &spectrumSearch{
        PropAction: "",
        PropAutocomplete: "",
        PropDisabled: false,
        PropGrows: false,
        PropHoldvalueonescape: false,
        PropInvalid: false,
        PropLabel: "Search",
        PropMaxlength: -1,
        PropMethod: "",
        PropMinlength: -1,
        PropMultiline: false,
        PropName: "",
        PropPattern: "",
        PropPlaceholder: "Search",
        PropQuiet: false,
        PropReadonly: false,
        PropRequired: false,
        PropRows: -1,
        PropSize: SearchSizeM,
        PropTabindex: 0,
        PropValid: false,
        PropValue: "",
    }

    element.styler = newStyler(element)

    return element
}

// The action URL for the search form
func (c *spectrumSearch) Action(action string) *spectrumSearch {
    c.PropAction = action
    return c
}

// What form of assistance should be provided when attempting to supply a value to the form control
func (c *spectrumSearch) Autocomplete(autocomplete string) *spectrumSearch {
    c.PropAutocomplete = autocomplete
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumSearch) Disabled(disabled bool) *spectrumSearch {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSearch) SetDisabled() *spectrumSearch {
    return c.Disabled(true)
}

// Whether a form control delivered with the `multiline` attribute will change size vertically to accommodate longer input
func (c *spectrumSearch) Grows(grows bool) *spectrumSearch {
    c.PropGrows = grows
    return c
}

func (c *spectrumSearch) SetGrows() *spectrumSearch {
    return c.Grows(true)
}

// Controls whether the typed value should be held when the Escape key is pressed
func (c *spectrumSearch) Holdvalueonescape(holdValueOnEscape bool) *spectrumSearch {
    c.PropHoldvalueonescape = holdValueOnEscape
    return c
}

func (c *spectrumSearch) SetHoldvalueonescape() *spectrumSearch {
    return c.Holdvalueonescape(true)
}

// Whether the `value` held by the form control is invalid
func (c *spectrumSearch) Invalid(invalid bool) *spectrumSearch {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumSearch) SetInvalid() *spectrumSearch {
    return c.Invalid(true)
}

// A string applied via `aria-label` to the form control when a user visible label is not provided
func (c *spectrumSearch) Label(label string) *spectrumSearch {
    c.PropLabel = label
    return c
}

// Defines the maximum string length that the user can enter
func (c *spectrumSearch) Maxlength(maxlength float64) *spectrumSearch {
    c.PropMaxlength = maxlength
    return c
}

// HTTP method to use for the form submission
func (c *spectrumSearch) Method(method string) *spectrumSearch {
    c.PropMethod = method
    return c
}

// Defines the minimum string length that the user can enter
func (c *spectrumSearch) Minlength(minlength float64) *spectrumSearch {
    c.PropMinlength = minlength
    return c
}

// Whether the form control should accept a value longer than one line
func (c *spectrumSearch) Multiline(multiline bool) *spectrumSearch {
    c.PropMultiline = multiline
    return c
}

func (c *spectrumSearch) SetMultiline() *spectrumSearch {
    return c.Multiline(true)
}

// Name of the form control
func (c *spectrumSearch) Name(name string) *spectrumSearch {
    c.PropName = name
    return c
}

// Pattern the `value` must match to be valid
func (c *spectrumSearch) Pattern(pattern string) *spectrumSearch {
    c.PropPattern = pattern
    return c
}

// Text that appears in the form control when it has no value set
func (c *spectrumSearch) Placeholder(placeholder string) *spectrumSearch {
    c.PropPlaceholder = placeholder
    return c
}

// Whether to display the form control with no visible background
func (c *spectrumSearch) Quiet(quiet bool) *spectrumSearch {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumSearch) SetQuiet() *spectrumSearch {
    return c.Quiet(true)
}

// Whether a user can interact with the value of the form control
func (c *spectrumSearch) Readonly(readonly bool) *spectrumSearch {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumSearch) SetReadonly() *spectrumSearch {
    return c.Readonly(true)
}

// Whether the form control will be found to be invalid when it holds no `value`
func (c *spectrumSearch) Required(required bool) *spectrumSearch {
    c.PropRequired = required
    return c
}

func (c *spectrumSearch) SetRequired() *spectrumSearch {
    return c.Required(true)
}

// The specific number of rows the form control should provide in the user interface
func (c *spectrumSearch) Rows(rows float64) *spectrumSearch {
    c.PropRows = rows
    return c
}

// Size variant of the search field
func (c *spectrumSearch) Size(size SearchSize) *spectrumSearch {
    c.PropSize = size
    return c
}

func (c *spectrumSearch) SizeS() *spectrumSearch {
    return c.Size(SearchSizeS)
}
func (c *spectrumSearch) SizeM() *spectrumSearch {
    return c.Size(SearchSizeM)
}
func (c *spectrumSearch) SizeL() *spectrumSearch {
    return c.Size(SearchSizeL)
}
func (c *spectrumSearch) SizeXl() *spectrumSearch {
    return c.Size(SearchSizeXl)
}
// The tab index to apply to this control
func (c *spectrumSearch) Tabindex(tabIndex float64) *spectrumSearch {
    c.PropTabindex = tabIndex
    return c
}

// Whether the `value` held by the form control is valid
func (c *spectrumSearch) Valid(valid bool) *spectrumSearch {
    c.PropValid = valid
    return c
}

func (c *spectrumSearch) SetValid() *spectrumSearch {
    return c.Valid(true)
}

// The value held by the form control
func (c *spectrumSearch) Value(value string) *spectrumSearch {
    c.PropValue = value
    return c
}



// Default or non-negative help text to associate to your form element
func (c *spectrumSearch) HelpText(content app.UI) *spectrumSearch {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when `invalid`
func (c *spectrumSearch) NegativeHelpText(content app.UI) *spectrumSearch {
    c.PropNegativeHelpTextSlot = content

    return c
}



// An alteration to the value of the element has been committed by the user
func (c *spectrumSearch) OnChange(handler app.EventHandler) *spectrumSearch {
    c.PropOnChange = handler

    return c
}

// The value of the element has changed
func (c *spectrumSearch) OnInput(handler app.EventHandler) *spectrumSearch {
    c.PropOnInput = handler

    return c
}

// The search form has been submitted
func (c *spectrumSearch) OnSubmit(handler app.EventHandler) *spectrumSearch {
    c.PropOnSubmit = handler

    return c
}


// Render renders the sp-search component
func (c *spectrumSearch) Render() app.UI {
    element := app.Elem("sp-search")

    // Set attributes
    if c.PropAction != "" {
        element = element.Attr("action", c.PropAction)
    }
    if c.PropAutocomplete != "" {
        element = element.Attr("autocomplete", c.PropAutocomplete)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropGrows {
        element = element.Attr("grows", true)
    }
    if c.PropHoldvalueonescape {
        element = element.Attr("holdValueOnEscape", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMaxlength != 0 {
        element = element.Attr("maxlength", c.PropMaxlength)
    }
    if c.PropMethod != "" {
        element = element.Attr("method", c.PropMethod)
    }
    if c.PropMinlength != 0 {
        element = element.Attr("minlength", c.PropMinlength)
    }
    if c.PropMultiline {
        element = element.Attr("multiline", true)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropPattern != "" {
        element = element.Attr("pattern", c.PropPattern)
    }
    if c.PropPlaceholder != "" {
        element = element.Attr("placeholder", c.PropPlaceholder)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropRequired {
        element = element.Attr("required", true)
    }
    if c.PropRows != 0 {
        element = element.Attr("rows", c.PropRows)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabIndex", c.PropTabindex)
    }
    if c.PropValid {
        element = element.Attr("valid", true)
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
    if c.PropOnSubmit != nil {
        element = element.On("submit", c.PropOnSubmit)
    }

    // Add slots and children
    slotElements := []app.UI{}


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


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 