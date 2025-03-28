package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumFieldGroup represents an sp-field-group component
type SpectrumFieldGroup struct {
	app.Compo

	// Properties
	horizontal bool
	vertical   bool
	invalid    bool
	label      string
	id         string

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Child elements
	children []app.UI
}

// FieldGroup creates a new field group component
func FieldGroup() *SpectrumFieldGroup {
	return &SpectrumFieldGroup{}
}

// Horizontal sets the horizontal layout for the field group
func (f *SpectrumFieldGroup) Horizontal(horizontal bool) *SpectrumFieldGroup {
	f.horizontal = horizontal
	return f
}

// Vertical sets the vertical layout for the field group
func (f *SpectrumFieldGroup) Vertical(vertical bool) *SpectrumFieldGroup {
	f.vertical = vertical
	return f
}

// Invalid sets the invalid state
func (f *SpectrumFieldGroup) Invalid(invalid bool) *SpectrumFieldGroup {
	f.invalid = invalid
	return f
}

// Label sets the accessibility label
func (f *SpectrumFieldGroup) Label(label string) *SpectrumFieldGroup {
	f.label = label
	return f
}

// ID sets the ID attribute
func (f *SpectrumFieldGroup) ID(id string) *SpectrumFieldGroup {
	f.id = id
	return f
}

// HelpText sets the help text UI element for the help-text slot
func (f *SpectrumFieldGroup) HelpText(helpText app.UI) *SpectrumFieldGroup {
	f.helpText = helpText
	return f
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (f *SpectrumFieldGroup) NegativeHelpText(helpText app.UI) *SpectrumFieldGroup {
	f.negativeHelpText = helpText
	return f
}

// Children sets the child elements of the field group
func (f *SpectrumFieldGroup) Children(children ...app.UI) *SpectrumFieldGroup {
	f.children = children
	return f
}

// Render renders the field group component
func (f *SpectrumFieldGroup) Render() app.UI {
	fieldGroup := app.Elem("sp-field-group").
		Attr("horizontal", f.horizontal).
		Attr("vertical", f.vertical).
		Attr("invalid", f.invalid).
		Attr("label", f.label)

	if f.id != "" {
		fieldGroup = fieldGroup.Attr("id", f.id)
	}

	// Add all elements
	elements := []app.UI{}

	// Add children first
	elements = append(elements, f.children...)

	// Add help text slots if provided
	if f.helpText != nil {
		helpTextElem := f.helpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if f.negativeHelpText != nil {
		negativeHelpTextElem := f.negativeHelpText
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
