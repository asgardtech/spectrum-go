package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// StatusLightSize represents the The size of the status light
type StatusLightSize string

// StatusLightSize values
const (
    StatusLightSizeS StatusLightSize = "s"
    StatusLightSizeM StatusLightSize = "m"
    StatusLightSizeL StatusLightSize = "l"
    StatusLightSizeXl StatusLightSize = "xl"
)
// StatusLightVariant represents the The visual variant to apply to this status light.
type StatusLightVariant string

// StatusLightVariant values
const (
    StatusLightVariantPositive StatusLightVariant = "positive"
    StatusLightVariantNegative StatusLightVariant = "negative"
    StatusLightVariantNotice StatusLightVariant = "notice"
    StatusLightVariantInfo StatusLightVariant = "info"
    StatusLightVariantNeutral StatusLightVariant = "neutral"
    StatusLightVariantYellow StatusLightVariant = "yellow"
    StatusLightVariantFuchsia StatusLightVariant = "fuchsia"
    StatusLightVariantIndigo StatusLightVariant = "indigo"
    StatusLightVariantSeafoam StatusLightVariant = "seafoam"
    StatusLightVariantChartreuse StatusLightVariant = "chartreuse"
    StatusLightVariantMagenta StatusLightVariant = "magenta"
    StatusLightVariantCelery StatusLightVariant = "celery"
    StatusLightVariantPurple StatusLightVariant = "purple"
)

// spectrumStatusLight represents an sp-status-light component
type spectrumStatusLight struct {
    app.Compo
    *styler[*spectrumStatusLight]

    // Properties
    // A status light in a disabled state shows that a status exists, but is not available in that circumstance. This can be used to maintain layout continuity and communicate that a status may become available later.
    PropDisabled bool
    // The size of the status light
    PropSize StatusLightSize
    // The visual variant to apply to this status light.
    PropVariant StatusLightVariant

    // Text content for the default slot
    PropText string

    // Content slots


}

// StatusLight creates a new sp-status-light component
func StatusLight() *spectrumStatusLight {
    element := &spectrumStatusLight{
        PropDisabled: false,
        PropSize: StatusLightSizeM,
        PropVariant: StatusLightVariantInfo,
    }

    element.styler = newStyler(element)

    return element
}

// A status light in a disabled state shows that a status exists, but is not available in that circumstance. This can be used to maintain layout continuity and communicate that a status may become available later.
func (c *spectrumStatusLight) Disabled(disabled bool) *spectrumStatusLight {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumStatusLight) SetDisabled() *spectrumStatusLight {
    return c.Disabled(true)
}

// The size of the status light
func (c *spectrumStatusLight) Size(size StatusLightSize) *spectrumStatusLight {
    c.PropSize = size
    return c
}

func (c *spectrumStatusLight) SizeS() *spectrumStatusLight {
    return c.Size(StatusLightSizeS)
}
func (c *spectrumStatusLight) SizeM() *spectrumStatusLight {
    return c.Size(StatusLightSizeM)
}
func (c *spectrumStatusLight) SizeL() *spectrumStatusLight {
    return c.Size(StatusLightSizeL)
}
func (c *spectrumStatusLight) SizeXl() *spectrumStatusLight {
    return c.Size(StatusLightSizeXl)
}
// The visual variant to apply to this status light.
func (c *spectrumStatusLight) Variant(variant StatusLightVariant) *spectrumStatusLight {
    c.PropVariant = variant
    return c
}

func (c *spectrumStatusLight) VariantPositive() *spectrumStatusLight {
    return c.Variant(StatusLightVariantPositive)
}
func (c *spectrumStatusLight) VariantNegative() *spectrumStatusLight {
    return c.Variant(StatusLightVariantNegative)
}
func (c *spectrumStatusLight) VariantNotice() *spectrumStatusLight {
    return c.Variant(StatusLightVariantNotice)
}
func (c *spectrumStatusLight) VariantInfo() *spectrumStatusLight {
    return c.Variant(StatusLightVariantInfo)
}
func (c *spectrumStatusLight) VariantNeutral() *spectrumStatusLight {
    return c.Variant(StatusLightVariantNeutral)
}
func (c *spectrumStatusLight) VariantYellow() *spectrumStatusLight {
    return c.Variant(StatusLightVariantYellow)
}
func (c *spectrumStatusLight) VariantFuchsia() *spectrumStatusLight {
    return c.Variant(StatusLightVariantFuchsia)
}
func (c *spectrumStatusLight) VariantIndigo() *spectrumStatusLight {
    return c.Variant(StatusLightVariantIndigo)
}
func (c *spectrumStatusLight) VariantSeafoam() *spectrumStatusLight {
    return c.Variant(StatusLightVariantSeafoam)
}
func (c *spectrumStatusLight) VariantChartreuse() *spectrumStatusLight {
    return c.Variant(StatusLightVariantChartreuse)
}
func (c *spectrumStatusLight) VariantMagenta() *spectrumStatusLight {
    return c.Variant(StatusLightVariantMagenta)
}
func (c *spectrumStatusLight) VariantCelery() *spectrumStatusLight {
    return c.Variant(StatusLightVariantCelery)
}
func (c *spectrumStatusLight) VariantPurple() *spectrumStatusLight {
    return c.Variant(StatusLightVariantPurple)
}

// Text sets the text content for the default slot
func (c *spectrumStatusLight) Text(text string) *spectrumStatusLight {
    c.PropText = text
    return c
}





// Render renders the sp-status-light component
func (c *spectrumStatusLight) Render() app.UI {
    element := app.Elem("sp-status-light")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 