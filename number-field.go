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

// spectrumNumberField represents an sp-number-field component
type spectrumNumberField struct {
	app.Compo

	// Properties
	PropSize          NumberFieldSize
	PropLabel         string
	PropValue         string
	PropPlaceholder   string
	PropName          string
	PropAutocomplete  string
	PropDisabled      bool
	PropGrows         bool
	PropInvalid       bool
	PropMaxlength     int
	PropMinlength     int
	PropMultiline     bool
	PropPattern       string
	PropQuiet         bool
	PropReadonly      bool
	PropRequired      bool
	PropRows          int
	PropTabIndex      int
	PropMin           *float64
	PropMax           *float64
	PropStep          *float64
	PropStepModifier  int
	PropHideStepper   bool
	PropIndeterminate bool
	PropFormatOptions map[string]interface{}

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
}

// NumberField creates a new number field component
func NumberField() *spectrumNumberField {
	return &spectrumNumberField{
		PropSize:          NumberFieldSizeM, // Default size is medium
		PropStepModifier:  10,               // Default step modifier is 10
		PropFormatOptions: make(map[string]interface{}),
	}
}

// Size sets the visual size of the number field
func (n *spectrumNumberField) Size(size NumberFieldSize) *spectrumNumberField {
	n.PropSize = size
	return n
}

// Label sets the accessibility label
func (n *spectrumNumberField) Label(label string) *spectrumNumberField {
	n.PropLabel = label
	return n
}

// Value sets the input value
func (n *spectrumNumberField) Value(value string) *spectrumNumberField {
	n.PropValue = value
	return n
}

// Placeholder sets the placeholder text
func (n *spectrumNumberField) Placeholder(placeholder string) *spectrumNumberField {
	n.PropPlaceholder = placeholder
	return n
}

// Name sets the form control name
func (n *spectrumNumberField) Name(name string) *spectrumNumberField {
	n.PropName = name
	return n
}

// Autocomplete sets the autocomplete attribute
func (n *spectrumNumberField) Autocomplete(autocomplete string) *spectrumNumberField {
	n.PropAutocomplete = autocomplete
	return n
}

// Disabled sets the disabled state
func (n *spectrumNumberField) Disabled(disabled bool) *spectrumNumberField {
	n.PropDisabled = disabled
	return n
}

// Grows sets whether the number field grows with content
func (n *spectrumNumberField) Grows(grows bool) *spectrumNumberField {
	n.PropGrows = grows
	return n
}

// Invalid sets the invalid state
func (n *spectrumNumberField) Invalid(invalid bool) *spectrumNumberField {
	n.PropInvalid = invalid
	return n
}

// Maxlength sets the maximum allowed length
func (n *spectrumNumberField) Maxlength(maxlength int) *spectrumNumberField {
	n.PropMaxlength = maxlength
	return n
}

// Minlength sets the minimum allowed length
func (n *spectrumNumberField) Minlength(minlength int) *spectrumNumberField {
	n.PropMinlength = minlength
	return n
}

// Multiline sets whether the number field accepts multiple lines
func (n *spectrumNumberField) Multiline(multiline bool) *spectrumNumberField {
	n.PropMultiline = multiline
	return n
}

// Pattern sets the validation pattern
func (n *spectrumNumberField) Pattern(pattern string) *spectrumNumberField {
	n.PropPattern = pattern
	return n
}

// Quiet sets whether the number field uses quiet styling
func (n *spectrumNumberField) Quiet(quiet bool) *spectrumNumberField {
	n.PropQuiet = quiet
	return n
}

// Readonly sets the readonly state
func (n *spectrumNumberField) Readonly(readonly bool) *spectrumNumberField {
	n.PropReadonly = readonly
	return n
}

// Required sets the required state
func (n *spectrumNumberField) Required(required bool) *spectrumNumberField {
	n.PropRequired = required
	return n
}

// Rows sets the number of visible rows
func (n *spectrumNumberField) Rows(rows int) *spectrumNumberField {
	n.PropRows = rows
	return n
}

// TabIndex sets the tab index
func (n *spectrumNumberField) TabIndex(tabIndex int) *spectrumNumberField {
	n.PropTabIndex = tabIndex
	return n
}

// Min sets the minimum allowed value
func (n *spectrumNumberField) Min(min float64) *spectrumNumberField {
	n.PropMin = &min
	return n
}

// Max sets the maximum allowed value
func (n *spectrumNumberField) Max(max float64) *spectrumNumberField {
	n.PropMax = &max
	return n
}

// Step sets the step value for incrementing/decrementing
func (n *spectrumNumberField) Step(step float64) *spectrumNumberField {
	n.PropStep = &step
	return n
}

// StepModifier sets the multiplier for step when using modifier keys
func (n *spectrumNumberField) StepModifier(modifier int) *spectrumNumberField {
	n.PropStepModifier = modifier
	return n
}

// HideStepper sets whether to hide the stepper UI
func (n *spectrumNumberField) HideStepper(hide bool) *spectrumNumberField {
	n.PropHideStepper = hide
	return n
}

// Indeterminate sets the indeterminate state
func (n *spectrumNumberField) Indeterminate(indeterminate bool) *spectrumNumberField {
	n.PropIndeterminate = indeterminate
	return n
}

// FormatOptions sets the Intl.NumberFormatOptions for formatting the displayed value
func (n *spectrumNumberField) FormatOptions(options map[string]interface{}) *spectrumNumberField {
	n.PropFormatOptions = options
	return n
}

// FormatAsPercent is a convenience method to format the value as a percentage
func (n *spectrumNumberField) FormatAsPercent() *spectrumNumberField {
	n.PropFormatOptions["style"] = "percent"
	return n
}

// FormatAsCurrency is a convenience method to format the value as currency
func (n *spectrumNumberField) FormatAsCurrency(currency string, display string) *spectrumNumberField {
	n.PropFormatOptions["style"] = "currency"
	n.PropFormatOptions["currency"] = currency
	if display != "" {
		n.PropFormatOptions["currencyDisplay"] = display
	}
	return n
}

// FormatAsUnit is a convenience method to format the value with a unit
func (n *spectrumNumberField) FormatAsUnit(unit string, display string) *spectrumNumberField {
	n.PropFormatOptions["style"] = "unit"
	n.PropFormatOptions["unit"] = unit
	if display != "" {
		n.PropFormatOptions["unitDisplay"] = display
	}
	return n
}

// HelpText sets the help text UI element for the help-text slot
func (n *spectrumNumberField) HelpText(helpText app.UI) *spectrumNumberField {
	n.PropHelpText = helpText
	return n
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (n *spectrumNumberField) NegativeHelpText(helpText app.UI) *spectrumNumberField {
	n.PropNegativeHelpText = helpText
	return n
}

// OnInput sets the input event handler
func (n *spectrumNumberField) OnInput(handler app.EventHandler) *spectrumNumberField {
	n.PropOnInput = handler
	return n
}

// OnChange sets the change event handler
func (n *spectrumNumberField) OnChange(handler app.EventHandler) *spectrumNumberField {
	n.PropOnChange = handler
	return n
}

// Render renders the number field component
func (n *spectrumNumberField) Render() app.UI {
	numberField := app.Elem("sp-number-field").
		Attr("size", string(n.PropSize)).
		Attr("label", n.PropLabel).
		Attr("value", n.PropValue).
		Attr("placeholder", n.PropPlaceholder).
		Attr("name", n.PropName).
		Attr("autocomplete", n.PropAutocomplete).
		Attr("disabled", n.PropDisabled).
		Attr("grows", n.PropGrows).
		Attr("invalid", n.PropInvalid).
		Attr("maxlength", n.PropMaxlength).
		Attr("minlength", n.PropMinlength).
		Attr("multiline", n.PropMultiline).
		Attr("pattern", n.PropPattern).
		Attr("quiet", n.PropQuiet).
		Attr("readonly", n.PropReadonly).
		Attr("required", n.PropRequired).
		Attr("rows", n.PropRows).
		Attr("tabindex", n.PropTabIndex).
		Attr("hide-stepper", n.PropHideStepper).
		Attr("indeterminate", n.PropIndeterminate).
		Attr("step-modifier", n.PropStepModifier)

	// Add min, max, step if set
	if n.PropMin != nil {
		numberField = numberField.Attr("min", *n.PropMin)
	}
	if n.PropMax != nil {
		numberField = numberField.Attr("max", *n.PropMax)
	}
	if n.PropStep != nil {
		numberField = numberField.Attr("step", *n.PropStep)
	}

	// Add format options if provided
	if len(n.PropFormatOptions) > 0 {
		formatOptionsJSON, _ := json.Marshal(n.PropFormatOptions)
		numberField = numberField.Attr("format-options", string(formatOptionsJSON))
	}

	// Add event handlers
	if n.PropOnInput != nil {
		numberField = numberField.On("input", n.PropOnInput)
	}
	if n.PropOnChange != nil {
		numberField = numberField.On("change", n.PropOnChange)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if n.PropHelpText != nil {
		helpTextElem := n.PropHelpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if n.PropNegativeHelpText != nil {
		negativeHelpTextElem := n.PropNegativeHelpText
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
