package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// HelpTextManager manages help text elements for form controls
type HelpTextManager struct {
	// Properties
	id               string
	mode             string
	helpText         app.UI
	negativeHelpText app.UI
}

// NewHelpTextManager creates a new help text manager
func NewHelpTextManager(id string, mode string) *HelpTextManager {
	return &HelpTextManager{
		id:   id,
		mode: mode,
	}
}

// SetHelpText sets the help text UI element
func (h *HelpTextManager) SetHelpText(helpText app.UI) {
	h.helpText = helpText
}

// SetNegativeHelpText sets the negative help text UI element
func (h *HelpTextManager) SetNegativeHelpText(negativeHelpText app.UI) {
	h.negativeHelpText = negativeHelpText
}

// Render renders the appropriate help text based on the invalid state
func (h *HelpTextManager) Render(invalid bool) app.UI {
	if invalid && h.negativeHelpText != nil {
		return h.negativeHelpText
	}
	return h.helpText
}

// GetHelpTextID returns the ID to be used for ARIA references
func (h *HelpTextManager) GetHelpTextID() string {
	return h.id + "-help-text"
}

// HelpTextMixin provides functionality for managing help text
// in form controls
type HelpTextMixin struct {
	// The ID attribute of the component
	componentID string

	// Help text manager
	helpTextManager *HelpTextManager
}

// InitHelpTextMixin initializes the help text mixin
func (h *HelpTextMixin) InitHelpTextMixin(componentID string, mode string) {
	h.componentID = componentID
	h.helpTextManager = NewHelpTextManager(componentID, mode)
}

// SetHelpText sets the help text UI element
func (h *HelpTextMixin) SetHelpText(helpText app.UI) {
	h.helpTextManager.SetHelpText(helpText)
}

// SetNegativeHelpText sets the negative help text UI element
func (h *HelpTextMixin) SetNegativeHelpText(negativeHelpText app.UI) {
	h.helpTextManager.SetNegativeHelpText(negativeHelpText)
}

// RenderHelpText renders the appropriate help text based on the invalid state
func (h *HelpTextMixin) RenderHelpText(invalid bool) app.UI {
	return h.helpTextManager.Render(invalid)
}

// GetHelpTextID returns the ID to be used for ARIA references
func (h *HelpTextMixin) GetHelpTextID() string {
	return h.helpTextManager.GetHelpTextID()
}

// MakeFormComponent is a helper function to create a form component with help text slots
func MakeFormComponent(elem app.UI, helpText app.UI, negativeHelpText app.UI, invalid bool) app.UI {
	// Create elements
	elements := []app.UI{elem}

	// Add the appropriate help text based on the invalid state
	if invalid && negativeHelpText != nil {
		if helpTextWithSlot, ok := negativeHelpText.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, helpTextWithSlot.Slot("negative-help-text"))
		} else {
			elements = append(elements,
				app.Div().
					Attr("slot", "negative-help-text").
					Body(negativeHelpText),
			)
		}
	} else if helpText != nil {
		if helpTextWithSlot, ok := helpText.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, helpTextWithSlot.Slot("help-text"))
		} else {
			elements = append(elements,
				app.Div().
					Attr("slot", "help-text").
					Body(helpText),
			)
		}
	}

	// Create a div with all elements
	return app.Div().Body(elements...)
}
