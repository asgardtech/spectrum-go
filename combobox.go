package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ComboboxSize represents the visual size of a combobox
type ComboboxSize string

// Combobox sizes
const (
	ComboboxSizeS  ComboboxSize = "s"
	ComboboxSizeM  ComboboxSize = "m"
	ComboboxSizeL  ComboboxSize = "l"
	ComboboxSizeXL ComboboxSize = "xl"
)

// ComboboxAutocomplete represents the autocomplete mode of a combobox
type ComboboxAutocomplete string

// Combobox autocomplete modes
const (
	ComboboxAutocompleteNone ComboboxAutocomplete = "none"
	ComboboxAutocompleteList ComboboxAutocomplete = "list"
)

// ComboboxOption represents an option in a combobox
type ComboboxOption struct {
	Value    string
	ItemText string
	Disabled bool
}

// spectrumCombobox represents an sp-combobox component
// It allows users to filter lists to only the options matching a query.
type spectrumCombobox struct {
	app.Compo

	// Properties
	PropID               string
	PropAutocomplete     ComboboxAutocomplete
	PropDisabled         bool
	PropGrows            bool
	PropInvalid          bool
	PropLabel            string
	PropMaxlength        int
	PropMinlength        int
	PropMultiline        bool
	PropName             string
	PropOpen             bool
	PropOptions          []ComboboxOption
	PropPattern          string
	PropPending          bool
	PropPendingLabel     string
	PropPlaceholder      string
	PropQuiet            bool
	PropReadonly         bool
	PropRequired         bool
	PropRows             int
	PropSize             ComboboxSize
	PropTabIndex         int
	PropValid            bool
	PropValue            string
	PropChildren         []app.UI
	PropHelpText         app.UI
	PropNegativeHelpText app.UI
	PropTooltip          app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnInput  app.EventHandler
}

// Combobox creates a new combobox component
func Combobox() *spectrumCombobox {
	return &spectrumCombobox{
		PropAutocomplete: ComboboxAutocompleteNone,
		PropPendingLabel: "Pending",
		PropSize:         ComboboxSizeM,
	}
}

// ID sets the ID of the combobox
func (c *spectrumCombobox) ID(id string) *spectrumCombobox {
	c.PropID = id
	return c
}

// Autocomplete sets the autocomplete mode of the combobox
func (c *spectrumCombobox) Autocomplete(autocomplete ComboboxAutocomplete) *spectrumCombobox {
	c.PropAutocomplete = autocomplete
	return c
}

// Disabled sets whether the combobox is disabled
func (c *spectrumCombobox) Disabled(disabled bool) *spectrumCombobox {
	c.PropDisabled = disabled
	return c
}

// Grows sets whether the combobox grows to accommodate longer input
func (c *spectrumCombobox) Grows(grows bool) *spectrumCombobox {
	c.PropGrows = grows
	return c
}

// Invalid sets whether the value is invalid
func (c *spectrumCombobox) Invalid(invalid bool) *spectrumCombobox {
	c.PropInvalid = invalid
	return c
}

// Label sets the label of the combobox
func (c *spectrumCombobox) Label(label string) *spectrumCombobox {
	c.PropLabel = label
	return c
}

// Maxlength sets the maximum length of the input
func (c *spectrumCombobox) Maxlength(maxlength int) *spectrumCombobox {
	c.PropMaxlength = maxlength
	return c
}

// Minlength sets the minimum length of the input
func (c *spectrumCombobox) Minlength(minlength int) *spectrumCombobox {
	c.PropMinlength = minlength
	return c
}

// Multiline sets whether the combobox accepts multi-line input
func (c *spectrumCombobox) Multiline(multiline bool) *spectrumCombobox {
	c.PropMultiline = multiline
	return c
}

// Name sets the name of the combobox
func (c *spectrumCombobox) Name(name string) *spectrumCombobox {
	c.PropName = name
	return c
}

// Open sets whether the listbox is visible
func (c *spectrumCombobox) Open(open bool) *spectrumCombobox {
	c.PropOpen = open
	return c
}

// Options sets the options for the combobox
func (c *spectrumCombobox) Options(options []ComboboxOption) *spectrumCombobox {
	c.PropOptions = options
	return c
}

// Pattern sets the pattern the value must match
func (c *spectrumCombobox) Pattern(pattern string) *spectrumCombobox {
	c.PropPattern = pattern
	return c
}

// Pending sets whether the items are currently loading
func (c *spectrumCombobox) Pending(pending bool) *spectrumCombobox {
	c.PropPending = pending
	return c
}

// PendingLabel sets the label for the pending state
func (c *spectrumCombobox) PendingLabel(pendingLabel string) *spectrumCombobox {
	c.PropPendingLabel = pendingLabel
	return c
}

// Placeholder sets the placeholder text
func (c *spectrumCombobox) Placeholder(placeholder string) *spectrumCombobox {
	c.PropPlaceholder = placeholder
	return c
}

// Quiet sets whether to display with no visible background
func (c *spectrumCombobox) Quiet(quiet bool) *spectrumCombobox {
	c.PropQuiet = quiet
	return c
}

// Readonly sets whether the user can interact with the value
func (c *spectrumCombobox) Readonly(readonly bool) *spectrumCombobox {
	c.PropReadonly = readonly
	return c
}

// Required sets whether the form control is required
func (c *spectrumCombobox) Required(required bool) *spectrumCombobox {
	c.PropRequired = required
	return c
}

// Rows sets the number of rows to display
func (c *spectrumCombobox) Rows(rows int) *spectrumCombobox {
	c.PropRows = rows
	return c
}

// Size sets the size of the combobox
func (c *spectrumCombobox) Size(size ComboboxSize) *spectrumCombobox {
	c.PropSize = size
	return c
}

// Small sets the size to small
func (c *spectrumCombobox) Small() *spectrumCombobox {
	c.PropSize = ComboboxSizeS
	return c
}

// Medium sets the size to medium
func (c *spectrumCombobox) Medium() *spectrumCombobox {
	c.PropSize = ComboboxSizeM
	return c
}

// Large sets the size to large
func (c *spectrumCombobox) Large() *spectrumCombobox {
	c.PropSize = ComboboxSizeL
	return c
}

// ExtraLarge sets the size to extra large
func (c *spectrumCombobox) ExtraLarge() *spectrumCombobox {
	c.PropSize = ComboboxSizeXL
	return c
}

// TabIndex sets the tab index of the combobox
func (c *spectrumCombobox) TabIndex(tabIndex int) *spectrumCombobox {
	c.PropTabIndex = tabIndex
	return c
}

// Valid sets whether the value is valid
func (c *spectrumCombobox) Valid(valid bool) *spectrumCombobox {
	c.PropValid = valid
	return c
}

// Value sets the value of the combobox
func (c *spectrumCombobox) Value(value string) *spectrumCombobox {
	c.PropValue = value
	return c
}

// Children sets the menu item children
func (c *spectrumCombobox) Children(children ...app.UI) *spectrumCombobox {
	c.PropChildren = children
	return c
}

// HelpText sets the help text
func (c *spectrumCombobox) HelpText(helpText app.UI) *spectrumCombobox {
	c.PropHelpText = helpText
	return c
}

// NegativeHelpText sets the negative help text
func (c *spectrumCombobox) NegativeHelpText(negativeHelpText app.UI) *spectrumCombobox {
	c.PropNegativeHelpText = negativeHelpText
	return c
}

// Tooltip sets the tooltip
func (c *spectrumCombobox) Tooltip(tooltip app.UI) *spectrumCombobox {
	c.PropTooltip = tooltip
	return c
}

// OnChange sets the change event handler
func (c *spectrumCombobox) OnChange(handler app.EventHandler) *spectrumCombobox {
	c.PropOnChange = handler
	return c
}

// OnInput sets the input event handler
func (c *spectrumCombobox) OnInput(handler app.EventHandler) *spectrumCombobox {
	c.PropOnInput = handler
	return c
}

// Render renders the combobox component
func (c *spectrumCombobox) Render() app.UI {
	combobox := app.Elem("sp-combobox")

	// Set attributes
	if c.PropID != "" {
		combobox = combobox.Attr("id", c.PropID)
	}
	if c.PropAutocomplete != ComboboxAutocompleteNone {
		combobox = combobox.Attr("autocomplete", string(c.PropAutocomplete))
	}
	if c.PropDisabled {
		combobox = combobox.Attr("disabled", true)
	}
	if c.PropGrows {
		combobox = combobox.Attr("grows", true)
	}
	if c.PropInvalid {
		combobox = combobox.Attr("invalid", true)
	}
	if c.PropLabel != "" {
		combobox = combobox.Attr("label", c.PropLabel)
	}
	if c.PropMaxlength > 0 {
		combobox = combobox.Attr("maxlength", c.PropMaxlength)
	}
	if c.PropMinlength > 0 {
		combobox = combobox.Attr("minlength", c.PropMinlength)
	}
	if c.PropMultiline {
		combobox = combobox.Attr("multiline", true)
	}
	if c.PropName != "" {
		combobox = combobox.Attr("name", c.PropName)
	}
	if c.PropOpen {
		combobox = combobox.Attr("open", true)
	}
	if c.PropPattern != "" {
		combobox = combobox.Attr("pattern", c.PropPattern)
	}
	if c.PropPending {
		combobox = combobox.Attr("pending", true)
	}
	if c.PropPendingLabel != "Pending" {
		combobox = combobox.Attr("pending-label", c.PropPendingLabel)
	}
	if c.PropPlaceholder != "" {
		combobox = combobox.Attr("placeholder", c.PropPlaceholder)
	}
	if c.PropQuiet {
		combobox = combobox.Attr("quiet", true)
	}
	if c.PropReadonly {
		combobox = combobox.Attr("readonly", true)
	}
	if c.PropRequired {
		combobox = combobox.Attr("required", true)
	}
	if c.PropRows > 0 {
		combobox = combobox.Attr("rows", c.PropRows)
	}
	if c.PropSize != ComboboxSizeM {
		combobox = combobox.Attr("size", string(c.PropSize))
	}
	if c.PropTabIndex != 0 {
		combobox = combobox.Attr("tabindex", c.PropTabIndex)
	}
	if c.PropValid {
		combobox = combobox.Attr("valid", true)
	}
	if c.PropValue != "" {
		combobox = combobox.Attr("value", c.PropValue)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		combobox = combobox.On("change", c.PropOnChange)
	}
	if c.PropOnInput != nil {
		combobox = combobox.On("input", c.PropOnInput)
	}

	// Prepare body elements
	elements := []app.UI{}

	// Add child elements
	elements = append(elements, c.PropChildren...)

	// Add slot elements
	if c.PropHelpText != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "help-text").
			Body(c.PropHelpText))
	}
	if c.PropNegativeHelpText != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "negative-help-text").
			Body(c.PropNegativeHelpText))
	}
	if c.PropTooltip != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "tooltip").
			Body(c.PropTooltip))
	}

	// Add all elements to the combobox
	if len(elements) > 0 {
		combobox = combobox.Body(elements...)
	}

	return combobox
}
