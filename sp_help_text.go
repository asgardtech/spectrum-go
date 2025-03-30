package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// HelpTextSize represents the The size of the help text
type HelpTextSize string

// HelpTextSize values
const (
    HelpTextSizeS HelpTextSize = "s"
    HelpTextSizeM HelpTextSize = "m"
    HelpTextSizeL HelpTextSize = "l"
    HelpTextSizeXl HelpTextSize = "xl"
)
// HelpTextVariant represents the The visual variant to apply to this help text.
type HelpTextVariant string

// HelpTextVariant values
const (
    HelpTextVariantNeutral HelpTextVariant = "neutral"
    HelpTextVariantNegative HelpTextVariant = "negative"
)

// spectrumHelpText represents an sp-help-text component
type spectrumHelpText struct {
    app.Compo
    *styler[*spectrumHelpText]

    // Properties
    // Whether the help text appears in a disabled state
    PropDisabled bool
    // Whether to display an icon as part of the help text
    PropIcon bool
    // The size of the help text
    PropSize HelpTextSize
    // The visual variant to apply to this help text.
    PropVariant HelpTextVariant

    // Text content for the default slot
    PropText string

    // Content slots


}

// HelpText creates a new sp-help-text component
func HelpText() *spectrumHelpText {
    element := &spectrumHelpText{
        PropDisabled: false,
        PropIcon: false,
        PropSize: HelpTextSizeM,
        PropVariant: HelpTextVariantNeutral,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the help text appears in a disabled state
func (c *spectrumHelpText) Disabled(disabled bool) *spectrumHelpText {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumHelpText) SetDisabled() *spectrumHelpText {
    return c.Disabled(true)
}

// Whether to display an icon as part of the help text
func (c *spectrumHelpText) Icon(icon bool) *spectrumHelpText {
    c.PropIcon = icon
    return c
}

func (c *spectrumHelpText) SetIcon() *spectrumHelpText {
    return c.Icon(true)
}

// The size of the help text
func (c *spectrumHelpText) Size(size HelpTextSize) *spectrumHelpText {
    c.PropSize = size
    return c
}

func (c *spectrumHelpText) SizeS() *spectrumHelpText {
    return c.Size(HelpTextSizeS)
}
func (c *spectrumHelpText) SizeM() *spectrumHelpText {
    return c.Size(HelpTextSizeM)
}
func (c *spectrumHelpText) SizeL() *spectrumHelpText {
    return c.Size(HelpTextSizeL)
}
func (c *spectrumHelpText) SizeXl() *spectrumHelpText {
    return c.Size(HelpTextSizeXl)
}
// The visual variant to apply to this help text.
func (c *spectrumHelpText) Variant(variant HelpTextVariant) *spectrumHelpText {
    c.PropVariant = variant
    return c
}

func (c *spectrumHelpText) VariantNeutral() *spectrumHelpText {
    return c.Variant(HelpTextVariantNeutral)
}
func (c *spectrumHelpText) VariantNegative() *spectrumHelpText {
    return c.Variant(HelpTextVariantNegative)
}

// Text sets the text content for the default slot
func (c *spectrumHelpText) Text(text string) *spectrumHelpText {
    c.PropText = text
    return c
}





// Render renders the sp-help-text component
func (c *spectrumHelpText) Render() app.UI {
    element := app.Elem("sp-help-text")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropIcon {
        element = element.Attr("icon", true)
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