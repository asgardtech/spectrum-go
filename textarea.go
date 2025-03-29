package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TextareaSize represents the visual size of a textarea
type TextareaSize string

// Textarea sizes
const (
	TextareaSizeS  TextareaSize = "s"
	TextareaSizeM  TextareaSize = "m"
	TextareaSizeL  TextareaSize = "l"
	TextareaSizeXL TextareaSize = "xl"
)

// spectrumTextarea represents an sp-textfield with multiline attribute
type spectrumTextarea struct {
	app.Compo

	// Properties
	PropSize         TextareaSize
	PropLabel        string
	PropValue        string
	PropPlaceholder  string
	PropName         string
	PropAutocomplete string
	PropDisabled     bool
	PropGrows        bool
	PropInvalid      bool
	PropMaxlength    int
	PropMinlength    int
	PropPattern      string
	PropQuiet        bool
	PropReadonly     bool
	PropRequired     bool
	PropRows         int

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
	PropOnFocus  app.EventHandler
	PropOnBlur   app.EventHandler
}

// Textarea creates a new textarea component
func Textarea() *spectrumTextarea {
	return &spectrumTextarea{
		PropSize: TextareaSizeM, // Default size is medium
	}
}

// Size sets the visual size of the textarea
func (t *spectrumTextarea) Size(size TextareaSize) *spectrumTextarea {
	t.PropSize = size
	return t
}

// Label sets the accessibility label
func (t *spectrumTextarea) Label(label string) *spectrumTextarea {
	t.PropLabel = label
	return t
}

// Value sets the input value
func (t *spectrumTextarea) Value(value string) *spectrumTextarea {
	t.PropValue = value
	return t
}

// Placeholder sets the placeholder text
func (t *spectrumTextarea) Placeholder(placeholder string) *spectrumTextarea {
	t.PropPlaceholder = placeholder
	return t
}

// Name sets the form control name
func (t *spectrumTextarea) Name(name string) *spectrumTextarea {
	t.PropName = name
	return t
}

// Autocomplete sets the autocomplete attribute
func (t *spectrumTextarea) Autocomplete(autocomplete string) *spectrumTextarea {
	t.PropAutocomplete = autocomplete
	return t
}

// Disabled sets the disabled state
func (t *spectrumTextarea) Disabled(disabled bool) *spectrumTextarea {
	t.PropDisabled = disabled
	return t
}

// Grows sets whether textarea grows with content
func (t *spectrumTextarea) Grows(grows bool) *spectrumTextarea {
	t.PropGrows = grows
	return t
}

// Invalid sets the invalid state
func (t *spectrumTextarea) Invalid(invalid bool) *spectrumTextarea {
	t.PropInvalid = invalid
	return t
}

// Maxlength sets the maximum allowed length
func (t *spectrumTextarea) Maxlength(maxlength int) *spectrumTextarea {
	t.PropMaxlength = maxlength
	return t
}

// Minlength sets the minimum allowed length
func (t *spectrumTextarea) Minlength(minlength int) *spectrumTextarea {
	t.PropMinlength = minlength
	return t
}

// Pattern sets the validation pattern
func (t *spectrumTextarea) Pattern(pattern string) *spectrumTextarea {
	t.PropPattern = pattern
	return t
}

// Quiet sets whether the textarea uses quiet styling
func (t *spectrumTextarea) Quiet(quiet bool) *spectrumTextarea {
	t.PropQuiet = quiet
	return t
}

// Readonly sets the readonly state
func (t *spectrumTextarea) Readonly(readonly bool) *spectrumTextarea {
	t.PropReadonly = readonly
	return t
}

// Required sets the required state
func (t *spectrumTextarea) Required(required bool) *spectrumTextarea {
	t.PropRequired = required
	return t
}

// Rows sets the number of visible rows
func (t *spectrumTextarea) Rows(rows int) *spectrumTextarea {
	t.PropRows = rows
	return t
}

// HelpText sets the help text UI element for the help-text slot
func (t *spectrumTextarea) HelpText(helpText app.UI) *spectrumTextarea {
	t.PropHelpText = helpText
	return t
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (t *spectrumTextarea) NegativeHelpText(helpText app.UI) *spectrumTextarea {
	t.PropNegativeHelpText = helpText
	return t
}

// OnInput sets the input event handler
func (t *spectrumTextarea) OnInput(handler app.EventHandler) *spectrumTextarea {
	t.PropOnInput = handler
	return t
}

// OnChange sets the change event handler
func (t *spectrumTextarea) OnChange(handler app.EventHandler) *spectrumTextarea {
	t.PropOnChange = handler
	return t
}

// OnFocus sets the focus event handler
func (t *spectrumTextarea) OnFocus(handler app.EventHandler) *spectrumTextarea {
	t.PropOnFocus = handler
	return t
}

// OnBlur sets the blur event handler
func (t *spectrumTextarea) OnBlur(handler app.EventHandler) *spectrumTextarea {
	t.PropOnBlur = handler
	return t
}

// Render renders the textarea component
func (t *spectrumTextarea) Render() app.UI {
	textarea := app.Elem("sp-textfield").
		Attr("size", string(t.PropSize)).
		Attr("multiline", true). // Always multiline for textarea
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
		Attr("pattern", t.PropPattern).
		Attr("quiet", t.PropQuiet).
		Attr("readonly", t.PropReadonly).
		Attr("required", t.PropRequired).
		Attr("rows", t.PropRows)

	// Add event handlers
	if t.PropOnInput != nil {
		textarea = textarea.On("input", t.PropOnInput)
	}
	if t.PropOnChange != nil {
		textarea = textarea.On("change", t.PropOnChange)
	}
	if t.PropOnFocus != nil {
		textarea = textarea.On("focus", t.PropOnFocus)
	}
	if t.PropOnBlur != nil {
		textarea = textarea.On("blur", t.PropOnBlur)
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

	// Add all elements to the textarea
	if len(elements) > 0 {
		textarea = textarea.Body(elements...)
	}

	return textarea
}
