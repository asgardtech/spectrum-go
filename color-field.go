package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumColorField represents an sp-color-field component
type spectrumColorField struct {
	app.Compo

	// Properties
	PropAutocomplete string
	PropDisabled     bool
	PropGrows        bool
	PropInvalid      bool
	PropLabel        string
	PropMaxlength    int
	PropMinlength    int
	PropMultiline    bool
	PropName         string
	PropPattern      string
	PropPlaceholder  string
	PropQuiet        bool
	PropReadonly     bool
	PropRequired     bool
	PropRows         int
	PropTabIndex     int
	PropValid        bool
	PropValue        string
	PropViewColor    bool
	PropSize         string

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// ColorField creates a new color field component
func ColorField() *spectrumColorField {
	return &spectrumColorField{
		PropLabel:     "",
		PropMaxlength: -1,
		PropMinlength: -1,
		PropRows:      -1,
	}
}

// Autocomplete sets the autocomplete attribute
func (cf *spectrumColorField) Autocomplete(autocomplete string) *spectrumColorField {
	cf.PropAutocomplete = autocomplete
	return cf
}

// Disabled sets whether the color field is disabled
func (cf *spectrumColorField) Disabled(disabled bool) *spectrumColorField {
	cf.PropDisabled = disabled
	return cf
}

// Grows sets whether the textarea will grow vertically to accommodate longer input
func (cf *spectrumColorField) Grows(grows bool) *spectrumColorField {
	cf.PropGrows = grows
	return cf
}

// Invalid sets whether the value held by the form control is invalid
func (cf *spectrumColorField) Invalid(invalid bool) *spectrumColorField {
	cf.PropInvalid = invalid
	return cf
}

// Label sets the aria-label for the color field
func (cf *spectrumColorField) Label(label string) *spectrumColorField {
	cf.PropLabel = label
	return cf
}

// Maxlength sets the maximum string length that the user can enter
func (cf *spectrumColorField) Maxlength(maxlength int) *spectrumColorField {
	cf.PropMaxlength = maxlength
	return cf
}

// Minlength sets the minimum string length that the user can enter
func (cf *spectrumColorField) Minlength(minlength int) *spectrumColorField {
	cf.PropMinlength = minlength
	return cf
}

// Multiline sets whether the form control should accept a value longer than one line
func (cf *spectrumColorField) Multiline(multiline bool) *spectrumColorField {
	cf.PropMultiline = multiline
	return cf
}

// Name sets the name of the form control
func (cf *spectrumColorField) Name(name string) *spectrumColorField {
	cf.PropName = name
	return cf
}

// Pattern sets the pattern the value must match to be valid
func (cf *spectrumColorField) Pattern(pattern string) *spectrumColorField {
	cf.PropPattern = pattern
	return cf
}

// Placeholder sets the text that appears in the form control when it has no value set
func (cf *spectrumColorField) Placeholder(placeholder string) *spectrumColorField {
	cf.PropPlaceholder = placeholder
	return cf
}

// Quiet sets whether to display the form control with no visible background
func (cf *spectrumColorField) Quiet(quiet bool) *spectrumColorField {
	cf.PropQuiet = quiet
	return cf
}

// Readonly sets whether a user can interact with the value of the form control
func (cf *spectrumColorField) Readonly(readonly bool) *spectrumColorField {
	cf.PropReadonly = readonly
	return cf
}

// Required sets whether the form control will be found to be invalid when it holds no value
func (cf *spectrumColorField) Required(required bool) *spectrumColorField {
	cf.PropRequired = required
	return cf
}

// Rows sets the specific number of rows the form control should provide in the user interface
func (cf *spectrumColorField) Rows(rows int) *spectrumColorField {
	cf.PropRows = rows
	return cf
}

// TabIndex sets the tab index
func (cf *spectrumColorField) TabIndex(tabIndex int) *spectrumColorField {
	cf.PropTabIndex = tabIndex
	return cf
}

// Valid sets whether the value held by the form control is valid
func (cf *spectrumColorField) Valid(valid bool) *spectrumColorField {
	cf.PropValid = valid
	return cf
}

// Value sets the value held by the form control
func (cf *spectrumColorField) Value(value string) *spectrumColorField {
	cf.PropValue = value
	return cf
}

// ViewColor sets whether to render the color handle
func (cf *spectrumColorField) ViewColor(viewColor bool) *spectrumColorField {
	cf.PropViewColor = viewColor
	return cf
}

// Size sets the size of the color field (s, m, l, xl)
func (cf *spectrumColorField) Size(size string) *spectrumColorField {
	cf.PropSize = size
	return cf
}

// Small sets the size to small
func (cf *spectrumColorField) Small() *spectrumColorField {
	return cf.Size("s")
}

// Medium sets the size to medium
func (cf *spectrumColorField) Medium() *spectrumColorField {
	return cf.Size("m")
}

// Large sets the size to large
func (cf *spectrumColorField) Large() *spectrumColorField {
	return cf.Size("l")
}

// ExtraLarge sets the size to extra large
func (cf *spectrumColorField) ExtraLarge() *spectrumColorField {
	return cf.Size("xl")
}

// OnChange sets the change event handler
func (cf *spectrumColorField) OnChange(handler app.EventHandler) *spectrumColorField {
	cf.PropOnChange = handler
	return cf
}

// OnInput sets the input event handler
func (cf *spectrumColorField) OnInput(handler app.EventHandler) *spectrumColorField {
	cf.PropOnInput = handler
	return cf
}

// Render renders the color field component
func (cf *spectrumColorField) Render() app.UI {
	colorField := app.Elem("sp-color-field")

	// Set attributes
	if cf.PropAutocomplete != "" {
		colorField = colorField.Attr("autocomplete", cf.PropAutocomplete)
	}
	if cf.PropDisabled {
		colorField = colorField.Attr("disabled", true)
	}
	if cf.PropGrows {
		colorField = colorField.Attr("grows", true)
	}
	if cf.PropInvalid {
		colorField = colorField.Attr("invalid", true)
	}
	if cf.PropLabel != "" {
		colorField = colorField.Attr("label", cf.PropLabel)
	}
	if cf.PropMaxlength != -1 {
		colorField = colorField.Attr("maxlength", cf.PropMaxlength)
	}
	if cf.PropMinlength != -1 {
		colorField = colorField.Attr("minlength", cf.PropMinlength)
	}
	if cf.PropMultiline {
		colorField = colorField.Attr("multiline", true)
	}
	if cf.PropName != "" {
		colorField = colorField.Attr("name", cf.PropName)
	}
	if cf.PropPattern != "" {
		colorField = colorField.Attr("pattern", cf.PropPattern)
	}
	if cf.PropPlaceholder != "" {
		colorField = colorField.Attr("placeholder", cf.PropPlaceholder)
	}
	if cf.PropQuiet {
		colorField = colorField.Attr("quiet", true)
	}
	if cf.PropReadonly {
		colorField = colorField.Attr("readonly", true)
	}
	if cf.PropRequired {
		colorField = colorField.Attr("required", true)
	}
	if cf.PropRows != -1 {
		colorField = colorField.Attr("rows", cf.PropRows)
	}
	if cf.PropTabIndex != 0 {
		colorField = colorField.Attr("tabIndex", cf.PropTabIndex)
	}
	if cf.PropValid {
		colorField = colorField.Attr("valid", true)
	}
	if cf.PropValue != "" {
		colorField = colorField.Attr("value", cf.PropValue)
	}
	if cf.PropViewColor {
		colorField = colorField.Attr("view-color", true)
	}
	if cf.PropSize != "" {
		colorField = colorField.Attr("size", cf.PropSize)
	}

	// Add event handlers
	if cf.PropOnChange != nil {
		colorField = colorField.On("change", cf.PropOnChange)
	}
	if cf.PropOnInput != nil {
		colorField = colorField.On("input", cf.PropOnInput)
	}

	return colorField
}
