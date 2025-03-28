package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumColorField represents an sp-color-field component
type SpectrumColorField struct {
	app.Compo

	// Properties
	autocomplete string
	disabled     bool
	grows        bool
	invalid      bool
	label        string
	maxlength    int
	minlength    int
	multiline    bool
	name         string
	pattern      string
	placeholder  string
	quiet        bool
	readonly     bool
	required     bool
	rows         int
	tabIndex     int
	valid        bool
	value        string
	viewColor    bool
	size         string

	// Event handlers
	onChange app.EventHandler
	onInput  app.EventHandler
}

// ColorField creates a new color field component
func ColorField() *SpectrumColorField {
	return &SpectrumColorField{
		label:     "",
		maxlength: -1,
		minlength: -1,
		rows:      -1,
	}
}

// Autocomplete sets the autocomplete attribute
func (cf *SpectrumColorField) Autocomplete(autocomplete string) *SpectrumColorField {
	cf.autocomplete = autocomplete
	return cf
}

// Disabled sets whether the color field is disabled
func (cf *SpectrumColorField) Disabled(disabled bool) *SpectrumColorField {
	cf.disabled = disabled
	return cf
}

// Grows sets whether the textarea will grow vertically to accommodate longer input
func (cf *SpectrumColorField) Grows(grows bool) *SpectrumColorField {
	cf.grows = grows
	return cf
}

// Invalid sets whether the value held by the form control is invalid
func (cf *SpectrumColorField) Invalid(invalid bool) *SpectrumColorField {
	cf.invalid = invalid
	return cf
}

// Label sets the aria-label for the color field
func (cf *SpectrumColorField) Label(label string) *SpectrumColorField {
	cf.label = label
	return cf
}

// Maxlength sets the maximum string length that the user can enter
func (cf *SpectrumColorField) Maxlength(maxlength int) *SpectrumColorField {
	cf.maxlength = maxlength
	return cf
}

// Minlength sets the minimum string length that the user can enter
func (cf *SpectrumColorField) Minlength(minlength int) *SpectrumColorField {
	cf.minlength = minlength
	return cf
}

// Multiline sets whether the form control should accept a value longer than one line
func (cf *SpectrumColorField) Multiline(multiline bool) *SpectrumColorField {
	cf.multiline = multiline
	return cf
}

// Name sets the name of the form control
func (cf *SpectrumColorField) Name(name string) *SpectrumColorField {
	cf.name = name
	return cf
}

// Pattern sets the pattern the value must match to be valid
func (cf *SpectrumColorField) Pattern(pattern string) *SpectrumColorField {
	cf.pattern = pattern
	return cf
}

// Placeholder sets the text that appears in the form control when it has no value set
func (cf *SpectrumColorField) Placeholder(placeholder string) *SpectrumColorField {
	cf.placeholder = placeholder
	return cf
}

// Quiet sets whether to display the form control with no visible background
func (cf *SpectrumColorField) Quiet(quiet bool) *SpectrumColorField {
	cf.quiet = quiet
	return cf
}

// Readonly sets whether a user can interact with the value of the form control
func (cf *SpectrumColorField) Readonly(readonly bool) *SpectrumColorField {
	cf.readonly = readonly
	return cf
}

// Required sets whether the form control will be found to be invalid when it holds no value
func (cf *SpectrumColorField) Required(required bool) *SpectrumColorField {
	cf.required = required
	return cf
}

// Rows sets the specific number of rows the form control should provide in the user interface
func (cf *SpectrumColorField) Rows(rows int) *SpectrumColorField {
	cf.rows = rows
	return cf
}

// TabIndex sets the tab index
func (cf *SpectrumColorField) TabIndex(tabIndex int) *SpectrumColorField {
	cf.tabIndex = tabIndex
	return cf
}

// Valid sets whether the value held by the form control is valid
func (cf *SpectrumColorField) Valid(valid bool) *SpectrumColorField {
	cf.valid = valid
	return cf
}

// Value sets the value held by the form control
func (cf *SpectrumColorField) Value(value string) *SpectrumColorField {
	cf.value = value
	return cf
}

// ViewColor sets whether to render the color handle
func (cf *SpectrumColorField) ViewColor(viewColor bool) *SpectrumColorField {
	cf.viewColor = viewColor
	return cf
}

// Size sets the size of the color field (s, m, l, xl)
func (cf *SpectrumColorField) Size(size string) *SpectrumColorField {
	cf.size = size
	return cf
}

// Small sets the size to small
func (cf *SpectrumColorField) Small() *SpectrumColorField {
	return cf.Size("s")
}

// Medium sets the size to medium
func (cf *SpectrumColorField) Medium() *SpectrumColorField {
	return cf.Size("m")
}

// Large sets the size to large
func (cf *SpectrumColorField) Large() *SpectrumColorField {
	return cf.Size("l")
}

// ExtraLarge sets the size to extra large
func (cf *SpectrumColorField) ExtraLarge() *SpectrumColorField {
	return cf.Size("xl")
}

// OnChange sets the change event handler
func (cf *SpectrumColorField) OnChange(handler app.EventHandler) *SpectrumColorField {
	cf.onChange = handler
	return cf
}

// OnInput sets the input event handler
func (cf *SpectrumColorField) OnInput(handler app.EventHandler) *SpectrumColorField {
	cf.onInput = handler
	return cf
}

// Render renders the color field component
func (cf *SpectrumColorField) Render() app.UI {
	colorField := app.Elem("sp-color-field")

	// Set attributes
	if cf.autocomplete != "" {
		colorField = colorField.Attr("autocomplete", cf.autocomplete)
	}
	if cf.disabled {
		colorField = colorField.Attr("disabled", true)
	}
	if cf.grows {
		colorField = colorField.Attr("grows", true)
	}
	if cf.invalid {
		colorField = colorField.Attr("invalid", true)
	}
	if cf.label != "" {
		colorField = colorField.Attr("label", cf.label)
	}
	if cf.maxlength != -1 {
		colorField = colorField.Attr("maxlength", cf.maxlength)
	}
	if cf.minlength != -1 {
		colorField = colorField.Attr("minlength", cf.minlength)
	}
	if cf.multiline {
		colorField = colorField.Attr("multiline", true)
	}
	if cf.name != "" {
		colorField = colorField.Attr("name", cf.name)
	}
	if cf.pattern != "" {
		colorField = colorField.Attr("pattern", cf.pattern)
	}
	if cf.placeholder != "" {
		colorField = colorField.Attr("placeholder", cf.placeholder)
	}
	if cf.quiet {
		colorField = colorField.Attr("quiet", true)
	}
	if cf.readonly {
		colorField = colorField.Attr("readonly", true)
	}
	if cf.required {
		colorField = colorField.Attr("required", true)
	}
	if cf.rows != -1 {
		colorField = colorField.Attr("rows", cf.rows)
	}
	if cf.tabIndex != 0 {
		colorField = colorField.Attr("tabIndex", cf.tabIndex)
	}
	if cf.valid {
		colorField = colorField.Attr("valid", true)
	}
	if cf.value != "" {
		colorField = colorField.Attr("value", cf.value)
	}
	if cf.viewColor {
		colorField = colorField.Attr("view-color", true)
	}
	if cf.size != "" {
		colorField = colorField.Attr("size", cf.size)
	}

	// Add event handlers
	if cf.onChange != nil {
		colorField = colorField.On("change", cf.onChange)
	}
	if cf.onInput != nil {
		colorField = colorField.On("input", cf.onInput)
	}

	return colorField
}
