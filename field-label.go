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
type SpectrumFieldLabel struct {
	app.Compo

	// Properties
	size        FieldLabelSize
	sideAligned FieldLabelSideAlignment
	disabled    bool
	required    bool
	for_        string
	id          string
	content     string
	innerHTML   string
	child       app.UI
}

// FieldLabel creates a new field label component
func FieldLabel() *SpectrumFieldLabel {
	return &SpectrumFieldLabel{
		size: FieldLabelSizeM, // Default size is medium
	}
}

// Size sets the visual size of the field label
func (f *SpectrumFieldLabel) Size(size FieldLabelSize) *SpectrumFieldLabel {
	f.size = size
	return f
}

// SideAligned sets the side alignment of the field label
func (f *SpectrumFieldLabel) SideAligned(alignment FieldLabelSideAlignment) *SpectrumFieldLabel {
	f.sideAligned = alignment
	return f
}

// Disabled sets the disabled state
func (f *SpectrumFieldLabel) Disabled(disabled bool) *SpectrumFieldLabel {
	f.disabled = disabled
	return f
}

// Required sets the required state
func (f *SpectrumFieldLabel) Required(required bool) *SpectrumFieldLabel {
	f.required = required
	return f
}

// For sets the "for" attribute to associate with a form element by ID
func (f *SpectrumFieldLabel) For(for_ string) *SpectrumFieldLabel {
	f.for_ = for_
	return f
}

// ID sets the ID of the field label
func (f *SpectrumFieldLabel) ID(id string) *SpectrumFieldLabel {
	f.id = id
	return f
}

// Content sets the text content of the field label
func (f *SpectrumFieldLabel) Content(content string) *SpectrumFieldLabel {
	f.content = content
	return f
}

// InnerHTML sets the inner HTML of the field label
func (f *SpectrumFieldLabel) InnerHTML(html string) *SpectrumFieldLabel {
	f.innerHTML = html
	return f
}

// Child sets a UI element as the child of the field label
func (f *SpectrumFieldLabel) Child(child app.UI) *SpectrumFieldLabel {
	f.child = child
	return f
}

// Render renders the field label component
func (f *SpectrumFieldLabel) Render() app.UI {
	fieldLabel := app.Elem("sp-field-label").
		Attr("size", string(f.size)).
		Attr("disabled", f.disabled).
		Attr("required", f.required)

	if f.for_ != "" {
		fieldLabel = fieldLabel.Attr("for", f.for_)
	}

	if f.id != "" {
		fieldLabel = fieldLabel.Attr("id", f.id)
	}

	if f.sideAligned != "" {
		fieldLabel = fieldLabel.Attr("side-aligned", string(f.sideAligned))
	}

	// Handle content in the right order of precedence
	if f.innerHTML != "" {
		fieldLabel = fieldLabel.Body(app.Raw(f.innerHTML))
	} else if f.child != nil {
		fieldLabel = fieldLabel.Body(f.child)
	} else if f.content != "" {
		fieldLabel = fieldLabel.Body(app.Text(f.content))
	}

	return fieldLabel
}
