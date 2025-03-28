package sp

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// NumberFieldSize represents the visual size of a number field
type NumberFieldSize string

// Number field sizes
const (
	NumberFieldSizeS  NumberFieldSize = "s"
	NumberFieldSizeM  NumberFieldSize = "m"
	NumberFieldSizeL  NumberFieldSize = "l"
	NumberFieldSizeXL NumberFieldSize = "xl"
)

// SpectrumNumberField represents an sp-number-field component
type SpectrumNumberField struct {
	app.Compo

	// Properties
	size          NumberFieldSize
	label         string
	value         string
	placeholder   string
	name          string
	autocomplete  string
	disabled      bool
	grows         bool
	invalid       bool
	maxlength     int
	minlength     int
	multiline     bool
	pattern       string
	quiet         bool
	readonly      bool
	required      bool
	rows          int
	tabIndex      int
	min           *float64
	max           *float64
	step          *float64
	stepModifier  int
	hideStepper   bool
	indeterminate bool
	formatOptions map[string]interface{}

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
}

// NumberField creates a new number field component
func NumberField() *SpectrumNumberField {
	return &SpectrumNumberField{
		size:          NumberFieldSizeM, // Default size is medium
		stepModifier:  10,               // Default step modifier is 10
		formatOptions: make(map[string]interface{}),
	}
}

// Size sets the visual size of the number field
func (n *SpectrumNumberField) Size(size NumberFieldSize) *SpectrumNumberField {
	n.size = size
	return n
}

// Label sets the accessibility label
func (n *SpectrumNumberField) Label(label string) *SpectrumNumberField {
	n.label = label
	return n
}

// Value sets the input value
func (n *SpectrumNumberField) Value(value string) *SpectrumNumberField {
	n.value = value
	return n
}

// Placeholder sets the placeholder text
func (n *SpectrumNumberField) Placeholder(placeholder string) *SpectrumNumberField {
	n.placeholder = placeholder
	return n
}

// Name sets the form control name
func (n *SpectrumNumberField) Name(name string) *SpectrumNumberField {
	n.name = name
	return n
}

// Autocomplete sets the autocomplete attribute
func (n *SpectrumNumberField) Autocomplete(autocomplete string) *SpectrumNumberField {
	n.autocomplete = autocomplete
	return n
}

// Disabled sets the disabled state
func (n *SpectrumNumberField) Disabled(disabled bool) *SpectrumNumberField {
	n.disabled = disabled
	return n
}

// Grows sets whether the number field grows with content
func (n *SpectrumNumberField) Grows(grows bool) *SpectrumNumberField {
	n.grows = grows
	return n
}

// Invalid sets the invalid state
func (n *SpectrumNumberField) Invalid(invalid bool) *SpectrumNumberField {
	n.invalid = invalid
	return n
}

// Maxlength sets the maximum allowed length
func (n *SpectrumNumberField) Maxlength(maxlength int) *SpectrumNumberField {
	n.maxlength = maxlength
	return n
}

// Minlength sets the minimum allowed length
func (n *SpectrumNumberField) Minlength(minlength int) *SpectrumNumberField {
	n.minlength = minlength
	return n
}

// Multiline sets whether the number field accepts multiple lines
func (n *SpectrumNumberField) Multiline(multiline bool) *SpectrumNumberField {
	n.multiline = multiline
	return n
}

// Pattern sets the validation pattern
func (n *SpectrumNumberField) Pattern(pattern string) *SpectrumNumberField {
	n.pattern = pattern
	return n
}

// Quiet sets whether the number field uses quiet styling
func (n *SpectrumNumberField) Quiet(quiet bool) *SpectrumNumberField {
	n.quiet = quiet
	return n
}

// Readonly sets the readonly state
func (n *SpectrumNumberField) Readonly(readonly bool) *SpectrumNumberField {
	n.readonly = readonly
	return n
}

// Required sets the required state
func (n *SpectrumNumberField) Required(required bool) *SpectrumNumberField {
	n.required = required
	return n
}

// Rows sets the number of visible rows
func (n *SpectrumNumberField) Rows(rows int) *SpectrumNumberField {
	n.rows = rows
	return n
}

// TabIndex sets the tab index
func (n *SpectrumNumberField) TabIndex(tabIndex int) *SpectrumNumberField {
	n.tabIndex = tabIndex
	return n
}

// Min sets the minimum allowed value
func (n *SpectrumNumberField) Min(min float64) *SpectrumNumberField {
	n.min = &min
	return n
}

// Max sets the maximum allowed value
func (n *SpectrumNumberField) Max(max float64) *SpectrumNumberField {
	n.max = &max
	return n
}

// Step sets the step value for incrementing/decrementing
func (n *SpectrumNumberField) Step(step float64) *SpectrumNumberField {
	n.step = &step
	return n
}

// StepModifier sets the multiplier for step when using modifier keys
func (n *SpectrumNumberField) StepModifier(modifier int) *SpectrumNumberField {
	n.stepModifier = modifier
	return n
}

// HideStepper sets whether to hide the stepper UI
func (n *SpectrumNumberField) HideStepper(hide bool) *SpectrumNumberField {
	n.hideStepper = hide
	return n
}

// Indeterminate sets the indeterminate state
func (n *SpectrumNumberField) Indeterminate(indeterminate bool) *SpectrumNumberField {
	n.indeterminate = indeterminate
	return n
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (n *SpectrumNumberField) FormatOptions(options map[string]interface{}) *SpectrumNumberField {
	n.formatOptions = options
	return n
}

// FormatAsPercent is a convenience method to format the value as a percentage
func (n *SpectrumNumberField) FormatAsPercent() *SpectrumNumberField {
	n.formatOptions["style"] = "percent"
	return n
}

// FormatAsCurrency is a convenience method to format the value as currency
func (n *SpectrumNumberField) FormatAsCurrency(currency string, display string) *SpectrumNumberField {
	n.formatOptions["style"] = "currency"
	n.formatOptions["currency"] = currency
	if display != "" {
		n.formatOptions["currencyDisplay"] = display
	}
	return n
}

// FormatAsUnit is a convenience method to format the value with a unit
func (n *SpectrumNumberField) FormatAsUnit(unit string, display string) *SpectrumNumberField {
	n.formatOptions["style"] = "unit"
	n.formatOptions["unit"] = unit
	if display != "" {
		n.formatOptions["unitDisplay"] = display
	}
	return n
}

// HelpText sets the help text UI element for the help-text slot
func (n *SpectrumNumberField) HelpText(helpText app.UI) *SpectrumNumberField {
	n.helpText = helpText
	return n
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (n *SpectrumNumberField) NegativeHelpText(helpText app.UI) *SpectrumNumberField {
	n.negativeHelpText = helpText
	return n
}

// OnInput sets the input event handler
func (n *SpectrumNumberField) OnInput(handler app.EventHandler) *SpectrumNumberField {
	n.onInput = handler
	return n
}

// OnChange sets the change event handler
func (n *SpectrumNumberField) OnChange(handler app.EventHandler) *SpectrumNumberField {
	n.onChange = handler
	return n
}

// Render renders the number field component
func (n *SpectrumNumberField) Render() app.UI {
	numberField := app.Elem("sp-number-field").
		Attr("size", string(n.size)).
		Attr("label", n.label).
		Attr("value", n.value).
		Attr("placeholder", n.placeholder).
		Attr("name", n.name).
		Attr("autocomplete", n.autocomplete).
		Attr("disabled", n.disabled).
		Attr("grows", n.grows).
		Attr("invalid", n.invalid).
		Attr("maxlength", n.maxlength).
		Attr("minlength", n.minlength).
		Attr("multiline", n.multiline).
		Attr("pattern", n.pattern).
		Attr("quiet", n.quiet).
		Attr("readonly", n.readonly).
		Attr("required", n.required).
		Attr("rows", n.rows).
		Attr("tabindex", n.tabIndex).
		Attr("hide-stepper", n.hideStepper).
		Attr("indeterminate", n.indeterminate).
		Attr("step-modifier", n.stepModifier)

	// Add min, max, step if set
	if n.min != nil {
		numberField = numberField.Attr("min", *n.min)
	}
	if n.max != nil {
		numberField = numberField.Attr("max", *n.max)
	}
	if n.step != nil {
		numberField = numberField.Attr("step", *n.step)
	}

	// Add format options if provided
	if len(n.formatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(n.formatOptions)
		numberField = numberField.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if n.onInput != nil {
		numberField = numberField.On("input", n.onInput)
	}
	if n.onChange != nil {
		numberField = numberField.On("change", n.onChange)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if n.helpText != nil {
		helpTextElem := n.helpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if n.negativeHelpText != nil {
		negativeHelpTextElem := n.negativeHelpText
		if negativeHelpTextWithSlot, ok := negativeHelpTextElem.(interface{ Slot(string) app.UI }); ok {
			negativeHelpTextElem = negativeHelpTextWithSlot.Slot("negative-help-text")
		} else {
			negativeHelpTextElem = app.Elem("div").
				Attr("slot", "negative-help-text").
				Body(negativeHelpTextElem)
		}
		elements = append(elements, negativeHelpTextElem)
	}

	// Add all elements to the number field
	if len(elements) > 0 {
		numberField = numberField.Body(elements...)
	}

	return numberField
}
