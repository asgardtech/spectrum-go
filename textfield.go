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

// SpectrumTextfield represents an sp-textfield component
type SpectrumTextfield struct {
	app.Compo

	// Properties
	size         TextfieldSize
	label        string
	value        string
	placeholder  string
	inputType    TextfieldType
	name         string
	autocomplete string
	disabled     bool
	grows        bool
	invalid      bool
	maxlength    int
	minlength    int
	multiline    bool
	pattern      string
	quiet        bool
	readonly     bool
	required     bool

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
	onFocus  app.EventHandler
	onBlur   app.EventHandler
}

// Textfield creates a new textfield component
func Textfield() *SpectrumTextfield {
	return &SpectrumTextfield{
		size:      TextfieldSizeM,    // Default size is medium
		inputType: TextfieldTypeText, // Default input type is text
	}
}

// Size sets the visual size of the textfield
func (t *SpectrumTextfield) Size(size TextfieldSize) *SpectrumTextfield {
	t.size = size
	return t
}

// Label sets the accessibility label
func (t *SpectrumTextfield) Label(label string) *SpectrumTextfield {
	t.label = label
	return t
}

// Value sets the input value
func (t *SpectrumTextfield) Value(value string) *SpectrumTextfield {
	t.value = value
	return t
}

// Placeholder sets the placeholder text
func (t *SpectrumTextfield) Placeholder(placeholder string) *SpectrumTextfield {
	t.placeholder = placeholder
	return t
}

// Type sets the input type
func (t *SpectrumTextfield) Type(inputType TextfieldType) *SpectrumTextfield {
	t.inputType = inputType
	return t
}

// Name sets the form control name
func (t *SpectrumTextfield) Name(name string) *SpectrumTextfield {
	t.name = name
	return t
}

// Autocomplete sets the autocomplete attribute
func (t *SpectrumTextfield) Autocomplete(autocomplete string) *SpectrumTextfield {
	t.autocomplete = autocomplete
	return t
}

// Disabled sets the disabled state
func (t *SpectrumTextfield) Disabled(disabled bool) *SpectrumTextfield {
	t.disabled = disabled
	return t
}

// Grows sets whether multiline textfield grows with content
func (t *SpectrumTextfield) Grows(grows bool) *SpectrumTextfield {
	t.grows = grows
	return t
}

// Invalid sets the invalid state
func (t *SpectrumTextfield) Invalid(invalid bool) *SpectrumTextfield {
	t.invalid = invalid
	return t
}

// Maxlength sets the maximum allowed length
func (t *SpectrumTextfield) Maxlength(maxlength int) *SpectrumTextfield {
	t.maxlength = maxlength
	return t
}

// Minlength sets the minimum allowed length
func (t *SpectrumTextfield) Minlength(minlength int) *SpectrumTextfield {
	t.minlength = minlength
	return t
}

// Multiline sets whether the textfield accepts multiple lines
func (t *SpectrumTextfield) Multiline(multiline bool) *SpectrumTextfield {
	t.multiline = multiline
	return t
}

// Pattern sets the validation pattern
func (t *SpectrumTextfield) Pattern(pattern string) *SpectrumTextfield {
	t.pattern = pattern
	return t
}

// Quiet sets whether the textfield uses quiet styling
func (t *SpectrumTextfield) Quiet(quiet bool) *SpectrumTextfield {
	t.quiet = quiet
	return t
}

// Readonly sets the readonly state
func (t *SpectrumTextfield) Readonly(readonly bool) *SpectrumTextfield {
	t.readonly = readonly
	return t
}

// Required sets the required state
func (t *SpectrumTextfield) Required(required bool) *SpectrumTextfield {
	t.required = required
	return t
}

// HelpText sets the help text UI element for the help-text slot
func (t *SpectrumTextfield) HelpText(helpText app.UI) *SpectrumTextfield {
	t.helpText = helpText
	return t
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (t *SpectrumTextfield) NegativeHelpText(helpText app.UI) *SpectrumTextfield {
	t.negativeHelpText = helpText
	return t
}

// OnInput sets the input event handler
func (t *SpectrumTextfield) OnInput(handler app.EventHandler) *SpectrumTextfield {
	t.onInput = handler
	return t
}

// OnChange sets the change event handler
func (t *SpectrumTextfield) OnChange(handler app.EventHandler) *SpectrumTextfield {
	t.onChange = handler
	return t
}

// OnFocus sets the focus event handler
func (t *SpectrumTextfield) OnFocus(handler app.EventHandler) *SpectrumTextfield {
	t.onFocus = handler
	return t
}

// OnBlur sets the blur event handler
func (t *SpectrumTextfield) OnBlur(handler app.EventHandler) *SpectrumTextfield {
	t.onBlur = handler
	return t
}

// Render renders the textfield component
func (t *SpectrumTextfield) Render() app.UI {
	textfield := app.Elem("sp-textfield").
		Attr("size", string(t.size)).
		Attr("type", string(t.inputType)).
		Attr("label", t.label).
		Attr("value", t.value).
		Attr("placeholder", t.placeholder).
		Attr("name", t.name).
		Attr("autocomplete", t.autocomplete).
		Attr("disabled", t.disabled).
		Attr("grows", t.grows).
		Attr("invalid", t.invalid).
		Attr("maxlength", t.maxlength).
		Attr("minlength", t.minlength).
		Attr("multiline", t.multiline).
		Attr("pattern", t.pattern).
		Attr("quiet", t.quiet).
		Attr("readonly", t.readonly).
		Attr("required", t.required)

	// Add event handlers
	if t.onInput != nil {
		textfield = textfield.On("input", t.onInput)
	}
	if t.onChange != nil {
		textfield = textfield.On("change", t.onChange)
	}
	if t.onFocus != nil {
		textfield = textfield.On("focus", t.onFocus)
	}
	if t.onBlur != nil {
		textfield = textfield.On("blur", t.onBlur)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if t.helpText != nil {
		helpTextElem := t.helpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if t.negativeHelpText != nil {
		negativeHelpTextElem := t.negativeHelpText
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
