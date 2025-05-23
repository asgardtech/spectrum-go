// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// FieldLabelSideAligned represents the Controls the alignment of the label when displayed inline with the form control
type FieldLabelSideAligned string

// FieldLabelSideAligned values
const (
	FieldLabelSideAlignedStart FieldLabelSideAligned = "start"
	FieldLabelSideAlignedEnd   FieldLabelSideAligned = "end"
)

// FieldLabelSize represents the The size of the field label
type FieldLabelSize string

// FieldLabelSize values
const (
	FieldLabelSizeS  FieldLabelSize = "s"
	FieldLabelSizeM  FieldLabelSize = "m"
	FieldLabelSizeL  FieldLabelSize = "l"
	FieldLabelSizeXl FieldLabelSize = "xl"
)

// spectrumFieldLabel represents an sp-field-label component
type spectrumFieldLabel struct {
	app.Compo
	*styler[*spectrumFieldLabel]
	*classer[*spectrumFieldLabel]
	*ider[*spectrumFieldLabel]

	// Properties
	// Whether the field label appears in a disabled state
	PropDisabled bool
	// The ID of the form element that this label is associated with
	PropFor string
	// Whether the field label indicates a required field
	PropRequired bool
	// Controls the alignment of the label when displayed inline with the form control
	PropSideAligned FieldLabelSideAligned
	// The size of the field label
	PropSize FieldLabelSize

	// Content for default slot
	PropBody []app.UI

	// Content slots

}

// IFieldLabel is the interface for sp-field-label component methods
type IFieldLabel interface {
	app.UI
	Styler[IFieldLabel]
	Classer[IFieldLabel]
	Ider[IFieldLabel]
	Disabled(bool) IFieldLabel
	SetDisabled() IFieldLabel
	For(string) IFieldLabel
	Required(bool) IFieldLabel
	SetRequired() IFieldLabel
	SideAligned(FieldLabelSideAligned) IFieldLabel
	SideAlignedStart() IFieldLabel
	SideAlignedEnd() IFieldLabel
	Size(FieldLabelSize) IFieldLabel
	SizeS() IFieldLabel
	SizeM() IFieldLabel
	SizeL() IFieldLabel
	SizeXl() IFieldLabel

	Body(...app.UI) IFieldLabel
	AddToBody(app.UI) IFieldLabel
	Text(string) IFieldLabel
}

// FieldLabel A field label provides accessible labelling for form elements. Use the for attribute to outline the id of an element in the same DOM tree to which it should associate itself.
func FieldLabel() IFieldLabel {
	element := &spectrumFieldLabel{
		PropDisabled: false,
		PropFor:      "",
		PropRequired: false,
		PropSize:     FieldLabelSizeM,
		PropBody:     []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Disabled Whether the field label appears in a disabled state
func (c *spectrumFieldLabel) Disabled(disabled bool) IFieldLabel {
	c.PropDisabled = disabled
	return c
}

func (c *spectrumFieldLabel) SetDisabled() IFieldLabel {
	return c.Disabled(true)
}

// For The ID of the form element that this label is associated with
func (c *spectrumFieldLabel) For(forValue string) IFieldLabel {
	c.PropFor = forValue
	return c
}

// Required Whether the field label indicates a required field
func (c *spectrumFieldLabel) Required(required bool) IFieldLabel {
	c.PropRequired = required
	return c
}

func (c *spectrumFieldLabel) SetRequired() IFieldLabel {
	return c.Required(true)
}

// SideAligned Controls the alignment of the label when displayed inline with the form control
func (c *spectrumFieldLabel) SideAligned(sideAligned FieldLabelSideAligned) IFieldLabel {
	c.PropSideAligned = sideAligned
	return c
}

func (c *spectrumFieldLabel) SideAlignedStart() IFieldLabel {
	return c.SideAligned(FieldLabelSideAlignedStart)
}
func (c *spectrumFieldLabel) SideAlignedEnd() IFieldLabel {
	return c.SideAligned(FieldLabelSideAlignedEnd)
}

// Size The size of the field label
func (c *spectrumFieldLabel) Size(size FieldLabelSize) IFieldLabel {
	c.PropSize = size
	return c
}

func (c *spectrumFieldLabel) SizeS() IFieldLabel {
	return c.Size(FieldLabelSizeS)
}
func (c *spectrumFieldLabel) SizeM() IFieldLabel {
	return c.Size(FieldLabelSizeM)
}
func (c *spectrumFieldLabel) SizeL() IFieldLabel {
	return c.Size(FieldLabelSizeL)
}
func (c *spectrumFieldLabel) SizeXl() IFieldLabel {
	return c.Size(FieldLabelSizeXl)
}

// Body sets the content for the default slot
func (c *spectrumFieldLabel) Body(elements ...app.UI) IFieldLabel {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumFieldLabel) AddToBody(element app.UI) IFieldLabel {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumFieldLabel) Text(text string) IFieldLabel {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Style sets a style property with a value
func (c *spectrumFieldLabel) Style(key, format string, values ...any) IFieldLabel {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumFieldLabel) Styles(styles map[string]string) IFieldLabel {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumFieldLabel) Class(class string) IFieldLabel {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumFieldLabel) Classes(classes ...string) IFieldLabel {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumFieldLabel) Id(id string) IFieldLabel {
	return c.ider.Id(id)
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

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add all elements to the component
	if len(slotElements) > 0 {
		element = element.Body(slotElements...)
	}

	// Apply styles, classes, and id
	element = element.Styles(c.styler.styles)

	// Apply classes if any
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}

	// Apply id if set
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}

	return element
}
