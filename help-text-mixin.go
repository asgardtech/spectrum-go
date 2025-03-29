package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// HelpTextManager manages help text elements for form controls
type HelpTextManager struct {
	// Properties
	PropID               string
	PropMode             string
	PropHelpText         app.UI
	PropNegativeHelpText app.UI
}

// NewHelpTextManager creates a new help text manager
func NewHelpTextManager(id string, mode string) *HelpTextManager {
	return &HelpTextManager{
		PropID:   id,
		PropMode: mode,
	}
}

// SetHelpText sets the help text UI element
func (h *HelpTextManager) SetHelpText(helpText app.UI) {
	h.PropHelpText = helpText
}

// SetNegativeHelpText sets the negative help text UI element
func (h *HelpTextManager) SetNegativeHelpText(negativeHelpText app.UI) {
	h.PropNegativeHelpText = negativeHelpText
}

// Render renders the appropriate help text based on the invalid state
func (h *HelpTextManager) Render(invalid bool) app.UI {
	if invalid && h.PropNegativeHelpText != nil {
		return h.PropNegativeHelpText
	}
	return h.PropHelpText
}

// GetHelpTextID returns the ID to be used for ARIA references
func (h *HelpTextManager) GetHelpTextID() string {
	return h.PropID + "-help-text"
}

// HelpTextMixin provides functionality for managing help text
// in form controls
type HelpTextMixin struct {
	// The ID attribute of the component
	PropComponentID string

	// Help text manager
	PropHelpTextManager *HelpTextManager
}

// InitHelpTextMixin initializes the help text mixin
func (h *HelpTextMixin) InitHelpTextMixin(componentID string, mode string) {
	h.PropComponentID = componentID
	h.PropHelpTextManager = NewHelpTextManager(componentID, mode)
}

// SetHelpText sets the help text UI element
func (h *HelpTextMixin) SetHelpText(helpText app.UI) {
	h.PropHelpTextManager.SetHelpText(helpText)
}

// SetNegativeHelpText sets the negative help text UI element
func (h *HelpTextMixin) SetNegativeHelpText(negativeHelpText app.UI) {
	h.PropHelpTextManager.SetNegativeHelpText(negativeHelpText)
}

// RenderHelpText renders the appropriate help text based on the invalid state
func (h *HelpTextMixin) RenderHelpText(invalid bool) app.UI {
	return h.PropHelpTextManager.Render(invalid)
}

// GetHelpTextID returns the ID to be used for ARIA references
func (h *HelpTextMixin) GetHelpTextID() string {
	return h.PropHelpTextManager.GetHelpTextID()
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
