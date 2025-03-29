package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TextfieldSize represents the visual size of a textfield
type TextfieldSize string

// Textfield sizes
const (
	TextfieldSizeS  TextfieldSize = "s"
	TextfieldSizeM  TextfieldSize = "m"
	TextfieldSizeL  TextfieldSize = "l"
	TextfieldSizeXL TextfieldSize = "xl"
)

// TextfieldType represents the input type of a textfield
type TextfieldType string

// Textfield types
const (
	TextfieldTypeText     TextfieldType = "text"
	TextfieldTypePassword TextfieldType = "password"
	TextfieldTypeEmail    TextfieldType = "email"
	TextfieldTypeURL      TextfieldType = "url"
	TextfieldTypeTel      TextfieldType = "tel"
)

// spectrumTextfield represents an sp-textfield component
type spectrumTextfield struct {
	app.Compo

	// Properties
	PropSize         TextfieldSize
	PropLabel        string
	PropValue        string
	PropPlaceholder  string
	PropInputType    TextfieldType
	PropName         string
	PropAutocomplete string
	PropDisabled     bool
	PropGrows        bool
	PropInvalid      bool
	PropMaxlength    int
	PropMinlength    int
	PropMultiline    bool
	PropPattern      string
	PropQuiet        bool
	PropReadonly     bool
	PropRequired     bool

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
	PropOnFocus  app.EventHandler
	PropOnBlur   app.EventHandler
}

// Textfield creates a new textfield component
func Textfield() *spectrumTextfield {
	return &spectrumTextfield{
		PropSize:      TextfieldSizeM,    // Default size is medium
		PropInputType: TextfieldTypeText, // Default input type is text
	}
}

// Size sets the visual size of the textfield
func (t *spectrumTextfield) Size(size TextfieldSize) *spectrumTextfield {
	t.PropSize = size
	return t
}

// Label sets the accessibility label
func (t *spectrumTextfield) Label(label string) *spectrumTextfield {
	t.PropLabel = label
	return t
}

// Value sets the input value
func (t *spectrumTextfield) Value(value string) *spectrumTextfield {
	t.PropValue = value
	return t
}

// Placeholder sets the placeholder text
func (t *spectrumTextfield) Placeholder(placeholder string) *spectrumTextfield {
	t.PropPlaceholder = placeholder
	return t
}

// Type sets the input type
func (t *spectrumTextfield) Type(inputType TextfieldType) *spectrumTextfield {
	t.PropInputType = inputType
	return t
}

// Name sets the form control name
func (t *spectrumTextfield) Name(name string) *spectrumTextfield {
	t.PropName = name
	return t
}

// Autocomplete sets the autocomplete attribute
func (t *spectrumTextfield) Autocomplete(autocomplete string) *spectrumTextfield {
	t.PropAutocomplete = autocomplete
	return t
}

// Disabled sets the disabled state
func (t *spectrumTextfield) Disabled(disabled bool) *spectrumTextfield {
	t.PropDisabled = disabled
	return t
}

// Grows sets whether multiline textfield grows with content
func (t *spectrumTextfield) Grows(grows bool) *spectrumTextfield {
	t.PropGrows = grows
	return t
}

// Invalid sets the invalid state
func (t *spectrumTextfield) Invalid(invalid bool) *spectrumTextfield {
	t.PropInvalid = invalid
	return t
}

// Maxlength sets the maximum allowed length
func (t *spectrumTextfield) Maxlength(maxlength int) *spectrumTextfield {
	t.PropMaxlength = maxlength
	return t
}

// Minlength sets the minimum allowed length
func (t *spectrumTextfield) Minlength(minlength int) *spectrumTextfield {
	t.PropMinlength = minlength
	return t
}

// Multiline sets whether the textfield accepts multiple lines
func (t *spectrumTextfield) Multiline(multiline bool) *spectrumTextfield {
	t.PropMultiline = multiline
	return t
}

// Pattern sets the validation pattern
func (t *spectrumTextfield) Pattern(pattern string) *spectrumTextfield {
	t.PropPattern = pattern
	return t
}

// Quiet sets whether the textfield uses quiet styling
func (t *spectrumTextfield) Quiet(quiet bool) *spectrumTextfield {
	t.PropQuiet = quiet
	return t
}

// Readonly sets the readonly state
func (t *spectrumTextfield) Readonly(readonly bool) *spectrumTextfield {
	t.PropReadonly = readonly
	return t
}

// Required sets the required state
func (t *spectrumTextfield) Required(required bool) *spectrumTextfield {
	t.PropRequired = required
	return t
}

// HelpText sets the help text UI element for the help-text slot
func (t *spectrumTextfield) HelpText(helpText app.UI) *spectrumTextfield {
	t.PropHelpText = helpText
	return t
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (t *spectrumTextfield) NegativeHelpText(helpText app.UI) *spectrumTextfield {
	t.PropNegativeHelpText = helpText
	return t
}

// OnInput sets the input event handler
func (t *spectrumTextfield) OnInput(handler app.EventHandler) *spectrumTextfield {
	t.PropOnInput = handler
	return t
}

// OnChange sets the change event handler
func (t *spectrumTextfield) OnChange(handler app.EventHandler) *spectrumTextfield {
	t.PropOnChange = handler
	return t
}

// OnFocus sets the focus event handler
func (t *spectrumTextfield) OnFocus(handler app.EventHandler) *spectrumTextfield {
	t.PropOnFocus = handler
	return t
}

// OnBlur sets the blur event handler
func (t *spectrumTextfield) OnBlur(handler app.EventHandler) *spectrumTextfield {
	t.PropOnBlur = handler
	return t
}

// Render renders the textfield component
func (t *spectrumTextfield) Render() app.UI {
	textfield := app.Elem("sp-textfield").
		Attr("size", string(t.PropSize)).
		Attr("type", string(t.PropInputType)).
		Attr("label", t.PropLabel).
		Attr("value", t.PropValue).
		Attr("placeholder", t.PropPlaceholder).
		Attr("name", t.PropName).
		Attr("autocomplete", t.PropAutocomplete).
		Attr("disabled", t.PropDisabled).
		Attr("grows", t.PropGrows).
		Attr("invalid", t.PropInvalid).
		Attr("maxlength", t.PropMaxlength).
		Attr("minlength", t.PropMinlength).
		Attr("multiline", t.PropMultiline).
		Attr("pattern", t.PropPattern).
		Attr("quiet", t.PropQuiet).
		Attr("readonly", t.PropReadonly).
		Attr("required", t.PropRequired)

	// Add event handlers
	if t.PropOnInput != nil {
		textfield = textfield.On("input", t.PropOnInput)
	}
	if t.PropOnChange != nil {
		textfield = textfield.On("change", t.PropOnChange)
	}
	if t.PropOnFocus != nil {
		textfield = textfield.On("focus", t.PropOnFocus)
	}
	if t.PropOnBlur != nil {
		textfield = textfield.On("blur", t.PropOnBlur)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if t.PropHelpText != nil {
		helpTextElem := t.PropHelpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if t.PropNegativeHelpText != nil {
		negativeHelpTextElem := t.PropNegativeHelpText
		if negativeHelpTextWithSlot, ok := negativeHelpTextElem.(interface{ Slot(string) app.UI }); ok {
			negativeHelpTextElem = negativeHelpTextWithSlot.Slot("negative-help-text")
		} else {
			negativeHelpTextElem = app.Elem("div").
				Attr("slot", "negative-help-text").
				Body(negativeHelpTextElem)
		}
		elements = append(elements, negativeHelpTextElem)
	}

	// Add all elements to the textfield
	if len(elements) > 0 {
		textfield = textfield.Body(elements...)
	}

	return textfield
}
