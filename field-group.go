package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumFieldGroup represents an sp-field-group component
type spectrumFieldGroup struct {
	app.Compo

	// Properties
	PropHorizontal bool
	PropVertical   bool
	PropInvalid    bool
	PropLabel      string
	PropID         string

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Child elements
	PropChildren []app.UI
}

// FieldGroup creates a new field group component
func FieldGroup() *spectrumFieldGroup {
	return &spectrumFieldGroup{}
}

// Horizontal sets the horizontal layout for the field group
func (f *spectrumFieldGroup) Horizontal(horizontal bool) *spectrumFieldGroup {
	f.PropHorizontal = horizontal
	return f
}

// Vertical sets the vertical layout for the field group
func (f *spectrumFieldGroup) Vertical(vertical bool) *spectrumFieldGroup {
	f.PropVertical = vertical
	return f
}

// Invalid sets the invalid state
func (f *spectrumFieldGroup) Invalid(invalid bool) *spectrumFieldGroup {
	f.PropInvalid = invalid
	return f
}

// Label sets the accessibility label
func (f *spectrumFieldGroup) Label(label string) *spectrumFieldGroup {
	f.PropLabel = label
	return f
}

// ID sets the ID attribute
func (f *spectrumFieldGroup) ID(id string) *spectrumFieldGroup {
	f.PropID = id
	return f
}

// HelpText sets the help text UI element for the help-text slot
func (f *spectrumFieldGroup) HelpText(helpText app.UI) *spectrumFieldGroup {
	f.PropHelpText = helpText
	return f
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (f *spectrumFieldGroup) NegativeHelpText(helpText app.UI) *spectrumFieldGroup {
	f.PropNegativeHelpText = helpText
	return f
}

// Children sets the child elements of the field group
func (f *spectrumFieldGroup) Children(children ...app.UI) *spectrumFieldGroup {
	f.PropChildren = children
	return f
}

// Render renders the field group component
func (f *spectrumFieldGroup) Render() app.UI {
	fieldGroup := app.Elem("sp-field-group").
		Attr("horizontal", f.PropHorizontal).
		Attr("vertical", f.PropVertical).
		Attr("invalid", f.PropInvalid).
		Attr("label", f.PropLabel)

	if f.PropID != "" {
		fieldGroup = fieldGroup.Attr("id", f.PropID)
	}

	// Add all elements
	elements := []app.UI{}

	// Add children first
	elements = append(elements, f.PropChildren...)

	// Add help text slots if provided
	if f.PropHelpText != nil {
		helpTextElem := f.PropHelpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if f.PropNegativeHelpText != nil {
		negativeHelpTextElem := f.PropNegativeHelpText
		if negativeHelpTextWithSlot, ok := negativeHelpTextElem.(interface{ Slot(string) app.UI }); ok {
			negativeHelpTextElem = negativeHelpTextWithSlot.Slot("negative-help-text")
		} else {
			negativeHelpTextElem = app.Elem("div").
				Attr("slot", "negative-help-text").
				Body(negativeHelpTextElem)
		}
		elements = append(elements, negativeHelpTextElem)
	}

	// Add all elements to the field group
	if len(elements) > 0 {
		fieldGroup = fieldGroup.Body(elements...)
	}

	return fieldGroup
}
