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

// SpectrumCombobox represents an sp-combobox component
// It allows users to filter lists to only the options matching a query.
type SpectrumCombobox struct {
	app.Compo

	// Properties
	id               string
	autocomplete     ComboboxAutocomplete
	disabled         bool
	grows            bool
	invalid          bool
	label            string
	maxlength        int
	minlength        int
	multiline        bool
	name             string
	open             bool
	options          []ComboboxOption
	pattern          string
	pending          bool
	pendingLabel     string
	placeholder      string
	quiet            bool
	readonly         bool
	required         bool
	rows             int
	size             ComboboxSize
	tabIndex         int
	valid            bool
	value            string
	children         []app.UI
	helpText         app.UI
	negativeHelpText app.UI
	tooltip          app.UI

	// Event handlers
	onChange app.EventHandler
	onInput  app.EventHandler
}

// Combobox creates a new combobox component
func Combobox() *SpectrumCombobox {
	return &SpectrumCombobox{
		autocomplete: ComboboxAutocompleteNone,
		pendingLabel: "Pending",
		size:         ComboboxSizeM,
	}
}

// ID sets the ID of the combobox
func (c *SpectrumCombobox) ID(id string) *SpectrumCombobox {
	c.id = id
	return c
}

// Autocomplete sets the autocomplete mode of the combobox
func (c *SpectrumCombobox) Autocomplete(autocomplete ComboboxAutocomplete) *SpectrumCombobox {
	c.autocomplete = autocomplete
	return c
}

// Disabled sets whether the combobox is disabled
func (c *SpectrumCombobox) Disabled(disabled bool) *SpectrumCombobox {
	c.disabled = disabled
	return c
}

// Grows sets whether the combobox grows to accommodate longer input
func (c *SpectrumCombobox) Grows(grows bool) *SpectrumCombobox {
	c.grows = grows
	return c
}

// Invalid sets whether the value is invalid
func (c *SpectrumCombobox) Invalid(invalid bool) *SpectrumCombobox {
	c.invalid = invalid
	return c
}

// Label sets the label of the combobox
func (c *SpectrumCombobox) Label(label string) *SpectrumCombobox {
	c.label = label
	return c
}

// Maxlength sets the maximum length of the input
func (c *SpectrumCombobox) Maxlength(maxlength int) *SpectrumCombobox {
	c.maxlength = maxlength
	return c
}

// Minlength sets the minimum length of the input
func (c *SpectrumCombobox) Minlength(minlength int) *SpectrumCombobox {
	c.minlength = minlength
	return c
}

// Multiline sets whether the combobox accepts multi-line input
func (c *SpectrumCombobox) Multiline(multiline bool) *SpectrumCombobox {
	c.multiline = multiline
	return c
}

// Name sets the name of the combobox
func (c *SpectrumCombobox) Name(name string) *SpectrumCombobox {
	c.name = name
	return c
}

// Open sets whether the listbox is visible
func (c *SpectrumCombobox) Open(open bool) *SpectrumCombobox {
	c.open = open
	return c
}

// Options sets the options for the combobox
func (c *SpectrumCombobox) Options(options []ComboboxOption) *SpectrumCombobox {
	c.options = options
	return c
}

// Pattern sets the pattern the value must match
func (c *SpectrumCombobox) Pattern(pattern string) *SpectrumCombobox {
	c.pattern = pattern
	return c
}

// Pending sets whether the items are currently loading
func (c *SpectrumCombobox) Pending(pending bool) *SpectrumCombobox {
	c.pending = pending
	return c
}

// PendingLabel sets the label for the pending state
func (c *SpectrumCombobox) PendingLabel(pendingLabel string) *SpectrumCombobox {
	c.pendingLabel = pendingLabel
	return c
}

// Placeholder sets the placeholder text
func (c *SpectrumCombobox) Placeholder(placeholder string) *SpectrumCombobox {
	c.placeholder = placeholder
	return c
}

// Quiet sets whether to display with no visible background
func (c *SpectrumCombobox) Quiet(quiet bool) *SpectrumCombobox {
	c.quiet = quiet
	return c
}

// Readonly sets whether the user can interact with the value
func (c *SpectrumCombobox) Readonly(readonly bool) *SpectrumCombobox {
	c.readonly = readonly
	return c
}

// Required sets whether the form control is required
func (c *SpectrumCombobox) Required(required bool) *SpectrumCombobox {
	c.required = required
	return c
}

// Rows sets the number of rows to display
func (c *SpectrumCombobox) Rows(rows int) *SpectrumCombobox {
	c.rows = rows
	return c
}

// Size sets the size of the combobox
func (c *SpectrumCombobox) Size(size ComboboxSize) *SpectrumCombobox {
	c.size = size
	return c
}

// Small sets the size to small
func (c *SpectrumCombobox) Small() *SpectrumCombobox {
	c.size = ComboboxSizeS
	return c
}

// Medium sets the size to medium
func (c *SpectrumCombobox) Medium() *SpectrumCombobox {
	c.size = ComboboxSizeM
	return c
}

// Large sets the size to large
func (c *SpectrumCombobox) Large() *SpectrumCombobox {
	c.size = ComboboxSizeL
	return c
}

// ExtraLarge sets the size to extra large
func (c *SpectrumCombobox) ExtraLarge() *SpectrumCombobox {
	c.size = ComboboxSizeXL
	return c
}

// TabIndex sets the tab index of the combobox
func (c *SpectrumCombobox) TabIndex(tabIndex int) *SpectrumCombobox {
	c.tabIndex = tabIndex
	return c
}

// Valid sets whether the value is valid
func (c *SpectrumCombobox) Valid(valid bool) *SpectrumCombobox {
	c.valid = valid
	return c
}

// Value sets the value of the combobox
func (c *SpectrumCombobox) Value(value string) *SpectrumCombobox {
	c.value = value
	return c
}

// Children sets the menu item children
func (c *SpectrumCombobox) Children(children ...app.UI) *SpectrumCombobox {
	c.children = children
	return c
}

// HelpText sets the help text
func (c *SpectrumCombobox) HelpText(helpText app.UI) *SpectrumCombobox {
	c.helpText = helpText
	return c
}

// NegativeHelpText sets the negative help text
func (c *SpectrumCombobox) NegativeHelpText(negativeHelpText app.UI) *SpectrumCombobox {
	c.negativeHelpText = negativeHelpText
	return c
}

// Tooltip sets the tooltip
func (c *SpectrumCombobox) Tooltip(tooltip app.UI) *SpectrumCombobox {
	c.tooltip = tooltip
	return c
}

// OnChange sets the change event handler
func (c *SpectrumCombobox) OnChange(handler app.EventHandler) *SpectrumCombobox {
	c.onChange = handler
	return c
}

// OnInput sets the input event handler
func (c *SpectrumCombobox) OnInput(handler app.EventHandler) *SpectrumCombobox {
	c.onInput = handler
	return c
}

// Render renders the combobox component
func (c *SpectrumCombobox) Render() app.UI {
	combobox := app.Elem("sp-combobox")

	// Set attributes
	if c.id != "" {
		combobox = combobox.Attr("id", c.id)
	}
	if c.autocomplete != ComboboxAutocompleteNone {
		combobox = combobox.Attr("autocomplete", string(c.autocomplete))
	}
	if c.disabled {
		combobox = combobox.Attr("disabled", true)
	}
	if c.grows {
		combobox = combobox.Attr("grows", true)
	}
	if c.invalid {
		combobox = combobox.Attr("invalid", true)
	}
	if c.label != "" {
		combobox = combobox.Attr("label", c.label)
	}
	if c.maxlength > 0 {
		combobox = combobox.Attr("maxlength", c.maxlength)
	}
	if c.minlength > 0 {
		combobox = combobox.Attr("minlength", c.minlength)
	}
	if c.multiline {
		combobox = combobox.Attr("multiline", true)
	}
	if c.name != "" {
		combobox = combobox.Attr("name", c.name)
	}
	if c.open {
		combobox = combobox.Attr("open", true)
	}
	if c.pattern != "" {
		combobox = combobox.Attr("pattern", c.pattern)
	}
	if c.pending {
		combobox = combobox.Attr("pending", true)
	}
	if c.pendingLabel != "Pending" {
		combobox = combobox.Attr("pending-label", c.pendingLabel)
	}
	if c.placeholder != "" {
		combobox = combobox.Attr("placeholder", c.placeholder)
	}
	if c.quiet {
		combobox = combobox.Attr("quiet", true)
	}
	if c.readonly {
		combobox = combobox.Attr("readonly", true)
	}
	if c.required {
		combobox = combobox.Attr("required", true)
	}
	if c.rows > 0 {
		combobox = combobox.Attr("rows", c.rows)
	}
	if c.size != ComboboxSizeM {
		combobox = combobox.Attr("size", string(c.size))
	}
	if c.tabIndex != 0 {
		combobox = combobox.Attr("tabindex", c.tabIndex)
	}
	if c.valid {
		combobox = combobox.Attr("valid", true)
	}
	if c.value != "" {
		combobox = combobox.Attr("value", c.value)
	}

	// Add event handlers
	if c.onChange != nil {
		combobox = combobox.On("change", c.onChange)
	}
	if c.onInput != nil {
		combobox = combobox.On("input", c.onInput)
	}

	// Prepare body elements
	elements := []app.UI{}

	// Add child elements
	elements = append(elements, c.children...)

	// Add slot elements
	if c.helpText != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "help-text").
			Body(c.helpText))
	}
	if c.negativeHelpText != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "negative-help-text").
			Body(c.negativeHelpText))
	}
	if c.tooltip != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "tooltip").
			Body(c.tooltip))
	}

	// Add all elements to the combobox
	if len(elements) > 0 {
		combobox = combobox.Body(elements...)
	}

	return combobox
}
