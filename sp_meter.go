package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MeterSize represents the Size variant for the meter
type MeterSize string

// MeterSize values
const (
    MeterSizeS MeterSize = "s"
    MeterSizeM MeterSize = "m"
    MeterSizeL MeterSize = "l"
)
// MeterVariant represents the The visual style to apply to this meter
type MeterVariant string

// MeterVariant values
const (
    MeterVariantPositive MeterVariant = "positive"
    MeterVariantNegative MeterVariant = "negative"
    MeterVariantNotice MeterVariant = "notice"
    MeterVariantNeutral MeterVariant = "neutral"
)

// spectrumMeter represents an sp-meter component
type spectrumMeter struct {
    app.Compo
    *styler[*spectrumMeter]

    // Properties
    // The accessible label for this meter
    PropLabel string
    // The maximum value for this meter
    PropMax float64
    // The minimum value for this meter
    PropMin float64
    // Whether to display the meter with a neutral color instead of the default positive color
    PropNeutral bool
    // Whether to display the meter with the notice color
    PropNotice bool
    // Whether the meter is to be displayed over a background
    PropOverbackground bool
    // A percentage value (0-100) for the meter; calculated from value, min, and max if not provided
    PropPercentage float64
    // Whether to display the meter with the positive color
    PropPositive bool
    // Optional side label to display next to the meter value
    PropSideLabel string
    // Size variant for the meter
    PropSize MeterSize
    // The current value for this meter
    PropValue float64
    // The visual style to apply to this meter
    PropVariant MeterVariant


    // Content slots
    PropLabelSlot app.UI
    PropValueSlot app.UI


}

// Meter creates a new sp-meter component
func Meter() *spectrumMeter {
    element := &spectrumMeter{
        PropLabel: "",
        PropMax: 100,
        PropMin: 0,
        PropNeutral: false,
        PropNotice: false,
        PropOverbackground: false,
        PropPercentage: 0,
        PropPositive: true,
        PropSideLabel: "",
        PropSize: MeterSizeM,
        PropValue: 0,
        PropVariant: MeterVariantPositive,
    }

    element.styler = newStyler(element)

    return element
}

// The accessible label for this meter
func (c *spectrumMeter) Label(label string) *spectrumMeter {
    c.PropLabel = label
    return c
}

// The maximum value for this meter
func (c *spectrumMeter) Max(max float64) *spectrumMeter {
    c.PropMax = max
    return c
}

// The minimum value for this meter
func (c *spectrumMeter) Min(min float64) *spectrumMeter {
    c.PropMin = min
    return c
}

// Whether to display the meter with a neutral color instead of the default positive color
func (c *spectrumMeter) Neutral(neutral bool) *spectrumMeter {
    c.PropNeutral = neutral
    return c
}

func (c *spectrumMeter) SetNeutral() *spectrumMeter {
    return c.Neutral(true)
}

// Whether to display the meter with the notice color
func (c *spectrumMeter) Notice(notice bool) *spectrumMeter {
    c.PropNotice = notice
    return c
}

func (c *spectrumMeter) SetNotice() *spectrumMeter {
    return c.Notice(true)
}

// Whether the meter is to be displayed over a background
func (c *spectrumMeter) Overbackground(overBackground bool) *spectrumMeter {
    c.PropOverbackground = overBackground
    return c
}

func (c *spectrumMeter) SetOverbackground() *spectrumMeter {
    return c.Overbackground(true)
}

// A percentage value (0-100) for the meter; calculated from value, min, and max if not provided
func (c *spectrumMeter) Percentage(percentage float64) *spectrumMeter {
    c.PropPercentage = percentage
    return c
}

// Whether to display the meter with the positive color
func (c *spectrumMeter) Positive(positive bool) *spectrumMeter {
    c.PropPositive = positive
    return c
}

func (c *spectrumMeter) SetPositive() *spectrumMeter {
    return c.Positive(true)
}

// Optional side label to display next to the meter value
func (c *spectrumMeter) SideLabel(sideLabel string) *spectrumMeter {
    c.PropSideLabel = sideLabel
    return c
}

// Size variant for the meter
func (c *spectrumMeter) Size(size MeterSize) *spectrumMeter {
    c.PropSize = size
    return c
}

func (c *spectrumMeter) SizeS() *spectrumMeter {
    return c.Size(MeterSizeS)
}
func (c *spectrumMeter) SizeM() *spectrumMeter {
    return c.Size(MeterSizeM)
}
func (c *spectrumMeter) SizeL() *spectrumMeter {
    return c.Size(MeterSizeL)
}
// The current value for this meter
func (c *spectrumMeter) Value(value float64) *spectrumMeter {
    c.PropValue = value
    return c
}

// The visual style to apply to this meter
func (c *spectrumMeter) Variant(variant MeterVariant) *spectrumMeter {
    c.PropVariant = variant
    return c
}

func (c *spectrumMeter) VariantPositive() *spectrumMeter {
    return c.Variant(MeterVariantPositive)
}
func (c *spectrumMeter) VariantNegative() *spectrumMeter {
    return c.Variant(MeterVariantNegative)
}
func (c *spectrumMeter) VariantNotice() *spectrumMeter {
    return c.Variant(MeterVariantNotice)
}
func (c *spectrumMeter) VariantNeutral() *spectrumMeter {
    return c.Variant(MeterVariantNeutral)
}


// Content for the meter label
func (c *spectrumMeter) LabelContent(content app.UI) *spectrumMeter {
    c.PropLabelSlot = content

    return c
}

// Content to display as the meter value
func (c *spectrumMeter) ValueContent(content app.UI) *spectrumMeter {
    c.PropValueSlot = content

    return c
}




// Render renders the sp-meter component
func (c *spectrumMeter) Render() app.UI {
    element := app.Elem("sp-meter")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropMax != 0 {
        element = element.Attr("max", c.PropMax)
    }
    if c.PropMin != 0 {
        element = element.Attr("min", c.PropMin)
    }
    if c.PropNeutral {
        element = element.Attr("neutral", true)
    }
    if c.PropNotice {
        element = element.Attr("notice", true)
    }
    if c.PropOverbackground {
        element = element.Attr("overBackground", true)
    }
    if c.PropPercentage != 0 {
        element = element.Attr("percentage", c.PropPercentage)
    }
    if c.PropPositive {
        element = element.Attr("positive", true)
    }
    if c.PropSideLabel != "" {
        element = element.Attr("side-label", c.PropSideLabel)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropValue != 0 {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }


    // Add slots and children
    slotElements := []app.UI{}


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