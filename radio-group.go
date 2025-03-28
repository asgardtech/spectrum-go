package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumRadioGroup represents an sp-radio-group component
type SpectrumRadioGroup struct {
	app.Compo

	// Properties
	label      string
	name       string
	selected   string
	invalid    bool
	vertical   bool
	horizontal bool

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Children radio buttons
	children []*SpectrumRadio

	// Event handlers
	onChange app.EventHandler
}

// RadioGroup creates a new radio group component
func RadioGroup() *SpectrumRadioGroup {
	return &SpectrumRadioGroup{}
}

// Label sets the group label text
func (rg *SpectrumRadioGroup) Label(label string) *SpectrumRadioGroup {
	rg.label = label
	return rg
}

// Name sets the name attribute for the radio group
func (rg *SpectrumRadioGroup) Name(name string) *SpectrumRadioGroup {
	rg.name = name
	return rg
}

// Selected sets the selected value for the radio group
func (rg *SpectrumRadioGroup) Selected(value string) *SpectrumRadioGroup {
	rg.selected = value
	return rg
}

// Invalid sets the invalid state of the radio group
func (rg *SpectrumRadioGroup) Invalid(invalid bool) *SpectrumRadioGroup {
	rg.invalid = invalid
	return rg
}

// Vertical sets the vertical orientation of the radio group
func (rg *SpectrumRadioGroup) Vertical(vertical bool) *SpectrumRadioGroup {
	rg.vertical = vertical
	rg.horizontal = !vertical
	return rg
}

// Horizontal sets the horizontal orientation of the radio group
func (rg *SpectrumRadioGroup) Horizontal(horizontal bool) *SpectrumRadioGroup {
	rg.horizontal = horizontal
	rg.vertical = !horizontal
	return rg
}

// HelpText sets the help text UI element for the help-text slot
func (rg *SpectrumRadioGroup) HelpText(helpText app.UI) *SpectrumRadioGroup {
	rg.helpText = helpText
	return rg
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (rg *SpectrumRadioGroup) NegativeHelpText(helpText app.UI) *SpectrumRadioGroup {
	rg.negativeHelpText = helpText
	return rg
}

// Children sets the radio button children
func (rg *SpectrumRadioGroup) Children(children ...*SpectrumRadio) *SpectrumRadioGroup {
	rg.children = children
	return rg
}

// OnChange sets the change event handler
func (rg *SpectrumRadioGroup) OnChange(handler app.EventHandler) *SpectrumRadioGroup {
	rg.onChange = handler
	return rg
}

// Render renders the radio group component
func (rg *SpectrumRadioGroup) Render() app.UI {
	radioGroup := app.Elem("sp-radio-group").
		Attr("label", rg.label).
		Attr("name", rg.name).
		Attr("selected", rg.selected).
		Attr("invalid", rg.invalid).
		Attr("vertical", rg.vertical).
		Attr("horizontal", rg.horizontal)

	// Add event handlers
	if rg.onChange != nil {
		radioGroup = radioGroup.On("change", rg.onChange)
	}

	// Create a slice to hold all child elements (radios and help text)
	elements := []app.UI{}

	// Add radio buttons
	if len(rg.children) > 0 {
		// Convert SpectrumRadio children to app.UI
		for _, child := range rg.children {
			elements = append(elements, child)
		}
	}

	// Add help text slots if provided
	if rg.helpText != nil {
		helpTextElem := rg.helpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if rg.negativeHelpText != nil {
		negativeHelpTextElem := rg.negativeHelpText
		if negativeHelpTextWithSlot, ok := negativeHelpTextElem.(interface{ Slot(string) app.UI }); ok {
			negativeHelpTextElem = negativeHelpTextWithSlot.Slot("negative-help-text")
		} else {
			negativeHelpTextElem = app.Elem("div").
				Attr("slot", "negative-help-text").
				Body(negativeHelpTextElem)
		}
		elements = append(elements, negativeHelpTextElem)
	}

	// Add all elements to the radio group
	if len(elements) > 0 {
		radioGroup = radioGroup.Body(elements...)
	}

	return radioGroup
}
