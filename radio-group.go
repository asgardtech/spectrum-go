package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumRadioGroup represents an sp-radio-group component
type spectrumRadioGroup struct {
	app.Compo

	// Properties
	PropLabel      string
	PropName       string
	PropSelected   string
	PropInvalid    bool
	PropVertical   bool
	PropHorizontal bool

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Children radio buttons
	PropChildren []*spectrumRadio

	// Event handlers
	PropOnChange app.EventHandler
}

// RadioGroup creates a new radio group component
func RadioGroup() *spectrumRadioGroup {
	return &spectrumRadioGroup{}
}

// Label sets the group label text
func (rg *spectrumRadioGroup) Label(label string) *spectrumRadioGroup {
	rg.PropLabel = label
	return rg
}

// Name sets the name attribute for the radio group
func (rg *spectrumRadioGroup) Name(name string) *spectrumRadioGroup {
	rg.PropName = name
	return rg
}

// Selected sets the selected value for the radio group
func (rg *spectrumRadioGroup) Selected(value string) *spectrumRadioGroup {
	rg.PropSelected = value
	return rg
}

// Invalid sets the invalid state of the radio group
func (rg *spectrumRadioGroup) Invalid(invalid bool) *spectrumRadioGroup {
	rg.PropInvalid = invalid
	return rg
}

// Vertical sets the vertical orientation of the radio group
func (rg *spectrumRadioGroup) Vertical(vertical bool) *spectrumRadioGroup {
	rg.PropVertical = vertical
	rg.PropHorizontal = !vertical
	return rg
}

// Horizontal sets the horizontal orientation of the radio group
func (rg *spectrumRadioGroup) Horizontal(horizontal bool) *spectrumRadioGroup {
	rg.PropHorizontal = horizontal
	rg.PropVertical = !horizontal
	return rg
}

// HelpText sets the help text UI element for the help-text slot
func (rg *spectrumRadioGroup) HelpText(helpText app.UI) *spectrumRadioGroup {
	rg.PropHelpText = helpText
	return rg
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (rg *spectrumRadioGroup) NegativeHelpText(helpText app.UI) *spectrumRadioGroup {
	rg.PropNegativeHelpText = helpText
	return rg
}

// Children sets the radio button children
func (rg *spectrumRadioGroup) Children(children ...*spectrumRadio) *spectrumRadioGroup {
	rg.PropChildren = children
	return rg
}

// OnChange sets the change event handler
func (rg *spectrumRadioGroup) OnChange(handler app.EventHandler) *spectrumRadioGroup {
	rg.PropOnChange = handler
	return rg
}

// Render renders the radio group component
func (rg *spectrumRadioGroup) Render() app.UI {
	radioGroup := app.Elem("sp-radio-group").
		Attr("label", rg.PropLabel).
		Attr("name", rg.PropName).
		Attr("selected", rg.PropSelected).
		Attr("invalid", rg.PropInvalid).
		Attr("vertical", rg.PropVertical).
		Attr("horizontal", rg.PropHorizontal)

	// Add event handlers
	if rg.PropOnChange != nil {
		radioGroup = radioGroup.On("change", rg.PropOnChange)
	}

	// Create a slice to hold all child elements (radios and help text)
	elements := []app.UI{}

	// Add radio buttons
	if len(rg.PropChildren) > 0 {
		// Convert spectrumRadio children to app.UI
		for _, child := range rg.PropChildren {
			elements = append(elements, child)
		}
	}

	// Add help text slots if provided
	if rg.PropHelpText != nil {
		helpTextElem := rg.PropHelpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if rg.PropNegativeHelpText != nil {
		negativeHelpTextElem := rg.PropNegativeHelpText
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
