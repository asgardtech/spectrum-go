package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// FieldLabelSideAligned represents the Controls the alignment of the label when displayed inline with the form control
type FieldLabelSideAligned string

// FieldLabelSideAligned values
const (
    FieldLabelSideAlignedStart FieldLabelSideAligned = "start"
    FieldLabelSideAlignedEnd FieldLabelSideAligned = "end"
)
// FieldLabelSize represents the The size of the field label
type FieldLabelSize string

// FieldLabelSize values
const (
    FieldLabelSizeS FieldLabelSize = "s"
    FieldLabelSizeM FieldLabelSize = "m"
    FieldLabelSizeL FieldLabelSize = "l"
    FieldLabelSizeXl FieldLabelSize = "xl"
)

// spectrumFieldLabel represents an sp-field-label component
type spectrumFieldLabel struct {
    app.Compo
    *styler[*spectrumFieldLabel]

    // Properties
    // Whether the field label appears in a disabled state
    PropDisabled bool
    // The ID of the form element that this label is associated with
    PropFor string
    // Unique identifier for the field label element
    PropId string
    // Whether the field label indicates a required field
    PropRequired bool
    // Controls the alignment of the label when displayed inline with the form control
    PropSideAligned FieldLabelSideAligned
    // The size of the field label
    PropSize FieldLabelSize

    // Text content for the default slot
    PropText string

    // Content slots


}

// FieldLabel creates a new sp-field-label component
func FieldLabel() *spectrumFieldLabel {
    element := &spectrumFieldLabel{
        PropDisabled: false,
        PropFor: "",
        PropId: "",
        PropRequired: false,
        PropSize: FieldLabelSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the field label appears in a disabled state
func (c *spectrumFieldLabel) Disabled(disabled bool) *spectrumFieldLabel {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumFieldLabel) SetDisabled() *spectrumFieldLabel {
    return c.Disabled(true)
}

// The ID of the form element that this label is associated with
func (c *spectrumFieldLabel) For(forValue string) *spectrumFieldLabel {
    c.PropFor = forValue
    return c
}

// Unique identifier for the field label element
func (c *spectrumFieldLabel) Id(id string) *spectrumFieldLabel {
    c.PropId = id
    return c
}

// Whether the field label indicates a required field
func (c *spectrumFieldLabel) Required(required bool) *spectrumFieldLabel {
    c.PropRequired = required
    return c
}

func (c *spectrumFieldLabel) SetRequired() *spectrumFieldLabel {
    return c.Required(true)
}

// Controls the alignment of the label when displayed inline with the form control
func (c *spectrumFieldLabel) SideAligned(sideAligned FieldLabelSideAligned) *spectrumFieldLabel {
    c.PropSideAligned = sideAligned
    return c
}

func (c *spectrumFieldLabel) SideAlignedStart() *spectrumFieldLabel {
    return c.SideAligned(FieldLabelSideAlignedStart)
}
func (c *spectrumFieldLabel) SideAlignedEnd() *spectrumFieldLabel {
    return c.SideAligned(FieldLabelSideAlignedEnd)
}
// The size of the field label
func (c *spectrumFieldLabel) Size(size FieldLabelSize) *spectrumFieldLabel {
    c.PropSize = size
    return c
}

func (c *spectrumFieldLabel) SizeS() *spectrumFieldLabel {
    return c.Size(FieldLabelSizeS)
}
func (c *spectrumFieldLabel) SizeM() *spectrumFieldLabel {
    return c.Size(FieldLabelSizeM)
}
func (c *spectrumFieldLabel) SizeL() *spectrumFieldLabel {
    return c.Size(FieldLabelSizeL)
}
func (c *spectrumFieldLabel) SizeXl() *spectrumFieldLabel {
    return c.Size(FieldLabelSizeXl)
}

// Text sets the text content for the default slot
func (c *spectrumFieldLabel) Text(text string) *spectrumFieldLabel {
    c.PropText = text
    return c
}





// Render renders the sp-field-label component
func (c *spectrumFieldLabel) Render() app.UI {
    element := app.Elem("sp-field-label")

    // Set attributes
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropFor != "" {
        element = element.Attr("for", c.PropFor)
    }
    if c.PropId != "" {
        element = element.Attr("id", c.PropId)
    }
    if c.PropRequired {
        element = element.Attr("required", true)
    }
    if c.PropSideAligned != "" {
        element = element.Attr("side-aligned", string(c.PropSideAligned))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
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