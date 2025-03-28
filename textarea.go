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

// SpectrumTextarea represents an sp-textfield with multiline attribute
type SpectrumTextarea struct {
	app.Compo

	// Properties
	size         TextareaSize
	label        string
	value        string
	placeholder  string
	name         string
	autocomplete string
	disabled     bool
	grows        bool
	invalid      bool
	maxlength    int
	minlength    int
	pattern      string
	quiet        bool
	readonly     bool
	required     bool
	rows         int

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
	onFocus  app.EventHandler
	onBlur   app.EventHandler
}

// Textarea creates a new textarea component
func Textarea() *SpectrumTextarea {
	return &SpectrumTextarea{
		size: TextareaSizeM, // Default size is medium
	}
}

// Size sets the visual size of the textarea
func (t *SpectrumTextarea) Size(size TextareaSize) *SpectrumTextarea {
	t.size = size
	return t
}

// Label sets the accessibility label
func (t *SpectrumTextarea) Label(label string) *SpectrumTextarea {
	t.label = label
	return t
}

// Value sets the input value
func (t *SpectrumTextarea) Value(value string) *SpectrumTextarea {
	t.value = value
	return t
}

// Placeholder sets the placeholder text
func (t *SpectrumTextarea) Placeholder(placeholder string) *SpectrumTextarea {
	t.placeholder = placeholder
	return t
}

// Name sets the form control name
func (t *SpectrumTextarea) Name(name string) *SpectrumTextarea {
	t.name = name
	return t
}

// Autocomplete sets the autocomplete attribute
func (t *SpectrumTextarea) Autocomplete(autocomplete string) *SpectrumTextarea {
	t.autocomplete = autocomplete
	return t
}

// Disabled sets the disabled state
func (t *SpectrumTextarea) Disabled(disabled bool) *SpectrumTextarea {
	t.disabled = disabled
	return t
}

// Grows sets whether textarea grows with content
func (t *SpectrumTextarea) Grows(grows bool) *SpectrumTextarea {
	t.grows = grows
	return t
}

// Invalid sets the invalid state
func (t *SpectrumTextarea) Invalid(invalid bool) *SpectrumTextarea {
	t.invalid = invalid
	return t
}

// Maxlength sets the maximum allowed length
func (t *SpectrumTextarea) Maxlength(maxlength int) *SpectrumTextarea {
	t.maxlength = maxlength
	return t
}

// Minlength sets the minimum allowed length
func (t *SpectrumTextarea) Minlength(minlength int) *SpectrumTextarea {
	t.minlength = minlength
	return t
}

// Pattern sets the validation pattern
func (t *SpectrumTextarea) Pattern(pattern string) *SpectrumTextarea {
	t.pattern = pattern
	return t
}

// Quiet sets whether the textarea uses quiet styling
func (t *SpectrumTextarea) Quiet(quiet bool) *SpectrumTextarea {
	t.quiet = quiet
	return t
}

// Readonly sets the readonly state
func (t *SpectrumTextarea) Readonly(readonly bool) *SpectrumTextarea {
	t.readonly = readonly
	return t
}

// Required sets the required state
func (t *SpectrumTextarea) Required(required bool) *SpectrumTextarea {
	t.required = required
	return t
}

// Rows sets the number of visible rows
func (t *SpectrumTextarea) Rows(rows int) *SpectrumTextarea {
	t.rows = rows
	return t
}

// HelpText sets the help text UI element for the help-text slot
func (t *SpectrumTextarea) HelpText(helpText app.UI) *SpectrumTextarea {
	t.helpText = helpText
	return t
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (t *SpectrumTextarea) NegativeHelpText(helpText app.UI) *SpectrumTextarea {
	t.negativeHelpText = helpText
	return t
}

// OnInput sets the input event handler
func (t *SpectrumTextarea) OnInput(handler app.EventHandler) *SpectrumTextarea {
	t.onInput = handler
	return t
}

// OnChange sets the change event handler
func (t *SpectrumTextarea) OnChange(handler app.EventHandler) *SpectrumTextarea {
	t.onChange = handler
	return t
}

// OnFocus sets the focus event handler
func (t *SpectrumTextarea) OnFocus(handler app.EventHandler) *SpectrumTextarea {
	t.onFocus = handler
	return t
}

// OnBlur sets the blur event handler
func (t *SpectrumTextarea) OnBlur(handler app.EventHandler) *SpectrumTextarea {
	t.onBlur = handler
	return t
}

// Render renders the textarea component
func (t *SpectrumTextarea) Render() app.UI {
	textarea := app.Elem("sp-textfield").
		Attr("size", string(t.size)).
		Attr("multiline", true). // Always multiline for textarea
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
		Attr("pattern", t.pattern).
		Attr("quiet", t.quiet).
		Attr("readonly", t.readonly).
		Attr("required", t.required).
		Attr("rows", t.rows)

	// Add event handlers
	if t.onInput != nil {
		textarea = textarea.On("input", t.onInput)
	}
	if t.onChange != nil {
		textarea = textarea.On("change", t.onChange)
	}
	if t.onFocus != nil {
		textarea = textarea.On("focus", t.onFocus)
	}
	if t.onBlur != nil {
		textarea = textarea.On("blur", t.onBlur)
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

	// Add all elements to the textarea
	if len(elements) > 0 {
		textarea = textarea.Body(elements...)
	}

	return textarea
}
