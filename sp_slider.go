package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SliderLabelVisibility represents the Controls which labels are visible: text (just the label text), value (just the value), or none (neither)
type SliderLabelVisibility string

// SliderLabelVisibility values
const (
    SliderLabelVisibilityText SliderLabelVisibility = "text"
    SliderLabelVisibilityValue SliderLabelVisibility = "value"
    SliderLabelVisibilityNone SliderLabelVisibility = "none"
)
// SliderSize represents the Size of the slider
type SliderSize string

// SliderSize values
const (
    SliderSizeS SliderSize = "s"
    SliderSizeM SliderSize = "m"
    SliderSizeL SliderSize = "l"
    SliderSizeXl SliderSize = "xl"
)
// SliderVariant represents the The visual variant of the slider
type SliderVariant string

// SliderVariant values
const (
    SliderVariantFilled SliderVariant = "filled"
    SliderVariantTick SliderVariant = "tick"
    SliderVariantRamp SliderVariant = "ramp"
    SliderVariantRange SliderVariant = "range"
)

// spectrumSlider represents an sp-slider component
type spectrumSlider struct {
    app.Compo
    *styler[*spectrumSlider]

    // Properties
    // Set the default value of the handle. Setting this property will cause the handle to reset to the default value on double click or pressing the escape key.
    PropDefaultValue float64
    // Disable this control. It will not receive focus or events
    PropDisabled bool
    // Whether the slider is currently being dragged
    PropDragging bool
    // Whether to display a Number Field along side the slider UI
    PropEditable bool
    // The starting point for filled variant
    PropFillStart float64
    // Number formatting options as a JSON string
    PropFormatOptions string
    // Whether the stepper UI of the Number Field is hidden or not
    PropHideStepper bool
    // Whether the slider track should be highlighted
    PropHighlight bool
    // Applies indeterminate to the underlying sp-number-field when editable === true. Is removed on the next change event.
    PropIndeterminate bool
    // Accessible text label for the slider
    PropLabel string
    // Controls which labels are visible: text (just the label text), value (just the value), or none (neither)
    PropLabelVisibility SliderLabelVisibility
    // Maximum value of the slider
    PropMax float64
    // Minimum value of the slider
    PropMin float64
    // Name of the form control for form submission
    PropName string
    // Applies quiet to the underlying sp-number-field when editable === true.
    PropQuiet bool
    // Size of the slider
    PropSize SliderSize
    // Step value for increments/decrements
    PropStep float64
    // The tab index to apply to this control. See general documentation about the tabindex HTML property
    PropTabindex float64
    // Whether to display labels for tick marks
    PropTickLabels bool
    // Interval at which to display tick marks
    PropTickStep float64
    // The current value of the slider
    PropValue float64
    // The visual variant of the slider
    PropVariant SliderVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropHandleSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnInput app.EventHandler
    PropOnSpSliderHandleReady app.EventHandler
}

// Slider creates a new sp-slider component
func Slider() *spectrumSlider {
    element := &spectrumSlider{
        PropDefaultValue: 0,
        PropDisabled: false,
        PropDragging: false,
        PropFillStart: 0,
        PropHideStepper: false,
        PropHighlight: false,
        PropIndeterminate: false,
        PropLabel: "",
        PropMax: 100,
        PropMin: 0,
        PropName: "",
        PropQuiet: false,
        PropSize: SliderSizeM,
        PropStep: 1,
        PropTabindex: 0,
        PropTickLabels: false,
        PropTickStep: 0,
        PropValue: 0,
    }

    element.styler = newStyler(element)

    return element
}

// Set the default value of the handle. Setting this property will cause the handle to reset to the default value on double click or pressing the escape key.
func (c *spectrumSlider) DefaultValue(defaultValue float64) *spectrumSlider {
    c.PropDefaultValue = defaultValue
    return c
}

// Disable this control. It will not receive focus or events
func (c *spectrumSlider) Disabled(disabled bool) *spectrumSlider {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumSlider) SetDisabled() *spectrumSlider {
    return c.Disabled(true)
}

// Whether the slider is currently being dragged
func (c *spectrumSlider) Dragging(dragging bool) *spectrumSlider {
    c.PropDragging = dragging
    return c
}

func (c *spectrumSlider) SetDragging() *spectrumSlider {
    return c.Dragging(true)
}

// Whether to display a Number Field along side the slider UI
func (c *spectrumSlider) Editable(editable bool) *spectrumSlider {
    c.PropEditable = editable
    return c
}

func (c *spectrumSlider) SetEditable() *spectrumSlider {
    return c.Editable(true)
}

// The starting point for filled variant
func (c *spectrumSlider) FillStart(fillStart float64) *spectrumSlider {
    c.PropFillStart = fillStart
    return c
}

// Number formatting options as a JSON string
func (c *spectrumSlider) FormatOptions(formatOptions string) *spectrumSlider {
    c.PropFormatOptions = formatOptions
    return c
}

// Whether the stepper UI of the Number Field is hidden or not
func (c *spectrumSlider) HideStepper(hideStepper bool) *spectrumSlider {
    c.PropHideStepper = hideStepper
    return c
}

func (c *spectrumSlider) SetHideStepper() *spectrumSlider {
    return c.HideStepper(true)
}

// Whether the slider track should be highlighted
func (c *spectrumSlider) Highlight(highlight bool) *spectrumSlider {
    c.PropHighlight = highlight
    return c
}

func (c *spectrumSlider) SetHighlight() *spectrumSlider {
    return c.Highlight(true)
}

// Applies indeterminate to the underlying sp-number-field when editable === true. Is removed on the next change event.
func (c *spectrumSlider) Indeterminate(indeterminate bool) *spectrumSlider {
    c.PropIndeterminate = indeterminate
    return c
}

func (c *spectrumSlider) SetIndeterminate() *spectrumSlider {
    return c.Indeterminate(true)
}

// Accessible text label for the slider
func (c *spectrumSlider) Label(label string) *spectrumSlider {
    c.PropLabel = label
    return c
}

// Controls which labels are visible: text (just the label text), value (just the value), or none (neither)
func (c *spectrumSlider) LabelVisibility(labelVisibility SliderLabelVisibility) *spectrumSlider {
    c.PropLabelVisibility = labelVisibility
    return c
}

func (c *spectrumSlider) LabelVisibilityText() *spectrumSlider {
    return c.LabelVisibility(SliderLabelVisibilityText)
}
func (c *spectrumSlider) LabelVisibilityValue() *spectrumSlider {
    return c.LabelVisibility(SliderLabelVisibilityValue)
}
func (c *spectrumSlider) LabelVisibilityNone() *spectrumSlider {
    return c.LabelVisibility(SliderLabelVisibilityNone)
}
// Maximum value of the slider
func (c *spectrumSlider) Max(max float64) *spectrumSlider {
    c.PropMax = max
    return c
}

// Minimum value of the slider
func (c *spectrumSlider) Min(min float64) *spectrumSlider {
    c.PropMin = min
    return c
}

// Name of the form control for form submission
func (c *spectrumSlider) Name(name string) *spectrumSlider {
    c.PropName = name
    return c
}

// Applies quiet to the underlying sp-number-field when editable === true.
func (c *spectrumSlider) Quiet(quiet bool) *spectrumSlider {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumSlider) SetQuiet() *spectrumSlider {
    return c.Quiet(true)
}

// Size of the slider
func (c *spectrumSlider) Size(size SliderSize) *spectrumSlider {
    c.PropSize = size
    return c
}

func (c *spectrumSlider) SizeS() *spectrumSlider {
    return c.Size(SliderSizeS)
}
func (c *spectrumSlider) SizeM() *spectrumSlider {
    return c.Size(SliderSizeM)
}
func (c *spectrumSlider) SizeL() *spectrumSlider {
    return c.Size(SliderSizeL)
}
func (c *spectrumSlider) SizeXl() *spectrumSlider {
    return c.Size(SliderSizeXl)
}
// Step value for increments/decrements
func (c *spectrumSlider) Step(step float64) *spectrumSlider {
    c.PropStep = step
    return c
}

// The tab index to apply to this control. See general documentation about the tabindex HTML property
func (c *spectrumSlider) Tabindex(tabindex float64) *spectrumSlider {
    c.PropTabindex = tabindex
    return c
}

// Whether to display labels for tick marks
func (c *spectrumSlider) TickLabels(tickLabels bool) *spectrumSlider {
    c.PropTickLabels = tickLabels
    return c
}

func (c *spectrumSlider) SetTickLabels() *spectrumSlider {
    return c.TickLabels(true)
}

// Interval at which to display tick marks
func (c *spectrumSlider) TickStep(tickStep float64) *spectrumSlider {
    c.PropTickStep = tickStep
    return c
}

// The current value of the slider
func (c *spectrumSlider) Value(value float64) *spectrumSlider {
    c.PropValue = value
    return c
}

// The visual variant of the slider
func (c *spectrumSlider) Variant(variant SliderVariant) *spectrumSlider {
    c.PropVariant = variant
    return c
}

func (c *spectrumSlider) VariantFilled() *spectrumSlider {
    return c.Variant(SliderVariantFilled)
}
func (c *spectrumSlider) VariantTick() *spectrumSlider {
    return c.Variant(SliderVariantTick)
}
func (c *spectrumSlider) VariantRamp() *spectrumSlider {
    return c.Variant(SliderVariantRamp)
}
func (c *spectrumSlider) VariantRange() *spectrumSlider {
    return c.Variant(SliderVariantRange)
}

// Text sets the text content for the default slot
func (c *spectrumSlider) Text(text string) *spectrumSlider {
    c.PropText = text
    return c
}


// Optionally accepts two or more sp-slider-handle elements
func (c *spectrumSlider) Handle(content app.UI) *spectrumSlider {
    c.PropHandleSlot = content

    return c
}



// An alteration to the value of the element has been committed by the user.
func (c *spectrumSlider) OnChange(handler app.EventHandler) *spectrumSlider {
    c.PropOnChange = handler

    return c
}

// The value of the element has changed.
func (c *spectrumSlider) OnInput(handler app.EventHandler) *spectrumSlider {
    c.PropOnInput = handler

    return c
}

// Fired when a slider handle has been registered with the slider
func (c *spectrumSlider) OnSpSliderHandleReady(handler app.EventHandler) *spectrumSlider {
    c.PropOnSpSliderHandleReady = handler

    return c
}


// Render renders the sp-slider component
func (c *spectrumSlider) Render() app.UI {
    element := app.Elem("sp-slider")

    // Set attributes
    if c.PropDefaultValue != 0 {
        element = element.Attr("default-value", c.PropDefaultValue)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropDragging {
        element = element.Attr("dragging", true)
    }
    if c.PropEditable {
        element = element.Attr("editable", true)
    }
    if c.PropFillStart != 0 {
        element = element.Attr("fill-start", c.PropFillStart)
    }
    if c.PropFormatOptions != "" {
        element = element.Attr("format-options", c.PropFormatOptions)
    }
    if c.PropHideStepper {
        element = element.Attr("hide-stepper", true)
    }
    if c.PropHighlight {
        element = element.Attr("highlight", true)
    }
    if c.PropIndeterminate {
        element = element.Attr("indeterminate", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropLabelVisibility != "" {
        element = element.Attr("label-visibility", string(c.PropLabelVisibility))
    }
    if c.PropMax != 0 {
        element = element.Attr("max", c.PropMax)
    }
    if c.PropMin != 0 {
        element = element.Attr("min", c.PropMin)
    }
    if c.PropName != "" {
        element = element.Attr("name", c.PropName)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropStep != 0 {
        element = element.Attr("step", c.PropStep)
    }
    if c.PropTabindex != 0 {
        element = element.Attr("tabindex", c.PropTabindex)
    }
    if c.PropTickLabels {
        element = element.Attr("tick-labels", true)
    }
    if c.PropTickStep != 0 {
        element = element.Attr("tick-step", c.PropTickStep)
    }
    if c.PropValue != 0 {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnInput != nil {
        element = element.On("input", c.PropOnInput)
    }
    if c.PropOnSpSliderHandleReady != nil {
        element = element.On("sp-slider-handle-ready", c.PropOnSpSliderHandleReady)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add handle slot
    if c.PropHandleSlot != nil {
        slotElem := c.PropHandleSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("handle")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "handle").
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