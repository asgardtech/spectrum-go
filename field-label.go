package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// FieldLabelSize represents the visual size of a field label
type FieldLabelSize string

// Field label sizes
const (
	FieldLabelSizeS  FieldLabelSize = "s"
	FieldLabelSizeM  FieldLabelSize = "m"
	FieldLabelSizeL  FieldLabelSize = "l"
	FieldLabelSizeXL FieldLabelSize = "xl"
)

// FieldLabelSideAlignment represents the side alignment of a field label
type FieldLabelSideAlignment string

// Field label side alignments
const (
	FieldLabelSideAlignmentStart FieldLabelSideAlignment = "start"
	FieldLabelSideAlignmentEnd   FieldLabelSideAlignment = "end"
)

// SpectrumFieldLabel represents an sp-field-label component
type spectrumFieldLabel struct {
	app.Compo

	// Properties
	PropSize        FieldLabelSize
	PropSideAligned FieldLabelSideAlignment
	PropDisabled    bool
	PropRequired    bool
	PropFor         string
	PropID          string
	PropContent     string
	PropInnerHTML   string
	PropChild       app.UI
}

// FieldLabel creates a new field label component
func FieldLabel() *spectrumFieldLabel {
	return &spectrumFieldLabel{
		PropSize: FieldLabelSizeM, // Default size is medium
	}
}

// Size sets the visual size of the field label
func (f *spectrumFieldLabel) Size(size FieldLabelSize) *spectrumFieldLabel {
	f.PropSize = size
	return f
}

// SideAligned sets the side alignment of the field label
func (f *spectrumFieldLabel) SideAligned(alignment FieldLabelSideAlignment) *spectrumFieldLabel {
	f.PropSideAligned = alignment
	return f
}

// Disabled sets the disabled state
func (f *spectrumFieldLabel) Disabled(disabled bool) *spectrumFieldLabel {
	f.PropDisabled = disabled
	return f
}

// Required sets the required state
func (f *spectrumFieldLabel) Required(required bool) *spectrumFieldLabel {
	f.PropRequired = required
	return f
}

// For sets the "for" attribute to associate with a form element by ID
func (f *spectrumFieldLabel) For(for_ string) *spectrumFieldLabel {
	f.PropFor = for_
	return f
}

// ID sets the ID of the field label
func (f *spectrumFieldLabel) ID(id string) *spectrumFieldLabel {
	f.PropID = id
	return f
}

// Content sets the text content of the field label
func (f *spectrumFieldLabel) Content(content string) *spectrumFieldLabel {
	f.PropContent = content
	return f
}

// InnerHTML sets the inner HTML of the field label
func (f *spectrumFieldLabel) InnerHTML(html string) *spectrumFieldLabel {
	f.PropInnerHTML = html
	return f
}

// Child sets a UI element as the child of the field label
func (f *spectrumFieldLabel) Child(child app.UI) *spectrumFieldLabel {
	f.PropChild = child
	return f
}

// Render renders the field label component
func (f *spectrumFieldLabel) Render() app.UI {
	fieldLabel := app.Elem("sp-field-label").
		Attr("size", string(f.PropSize)).
		Attr("disabled", f.PropDisabled).
		Attr("required", f.PropRequired)

	if f.PropFor != "" {
		fieldLabel = fieldLabel.Attr("for", f.PropFor)
	}

	if f.PropID != "" {
		fieldLabel = fieldLabel.Attr("id", f.PropID)
	}

	if f.PropSideAligned != "" {
		fieldLabel = fieldLabel.Attr("side-aligned", string(f.PropSideAligned))
	}

	// Handle content in the right order of precedence
	if f.PropInnerHTML != "" {
		fieldLabel = fieldLabel.Body(app.Raw(f.PropInnerHTML))
	} else if f.PropChild != nil {
		fieldLabel = fieldLabel.Body(f.PropChild)
	} else if f.PropContent != "" {
		fieldLabel = fieldLabel.Body(app.Text(f.PropContent))
	}

	return fieldLabel
}
