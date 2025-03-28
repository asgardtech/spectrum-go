package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SearchSize represents the visual size of a search field
type SearchSize string

// Search sizes
const (
	SearchSizeS  SearchSize = "s"
	SearchSizeM  SearchSize = "m"
	SearchSizeL  SearchSize = "l"
	SearchSizeXL SearchSize = "xl"
)

// SearchMethod represents the method to use when submitting the search form
type SearchMethod string

// Search methods
const (
	SearchMethodGet    SearchMethod = "get"
	SearchMethodPost   SearchMethod = "post"
	SearchMethodDialog SearchMethod = "dialog"
)

// SpectrumSearch represents an sp-search component
type SpectrumSearch struct {
	app.Compo

	// Properties
	size              SearchSize
	action            string
	method            SearchMethod
	label             string
	value             string
	placeholder       string
	name              string
	autocomplete      string
	disabled          bool
	grows             bool
	invalid           bool
	maxlength         int
	minlength         int
	multiline         bool
	pattern           string
	quiet             bool
	readonly          bool
	required          bool
	rows              int
	tabIndex          int
	holdValueOnEscape bool

	// Help text slots
	helpText         app.UI
	negativeHelpText app.UI

	// Event handlers
	onInput  app.EventHandler
	onChange app.EventHandler
	onSubmit app.EventHandler
}

// Search creates a new search component
func Search() *SpectrumSearch {
	return &SpectrumSearch{
		size:        SearchSizeM, // Default size is medium
		placeholder: "Search",    // Default placeholder
		label:       "Search",    // Default accessibility label
	}
}

// Size sets the visual size of the search field
func (s *SpectrumSearch) Size(size SearchSize) *SpectrumSearch {
	s.size = size
	return s
}

// Action sets the form action URL
func (s *SpectrumSearch) Action(action string) *SpectrumSearch {
	s.action = action
	return s
}

// Method sets the form submission method
func (s *SpectrumSearch) Method(method SearchMethod) *SpectrumSearch {
	s.method = method
	return s
}

// Label sets the accessibility label
func (s *SpectrumSearch) Label(label string) *SpectrumSearch {
	s.label = label
	return s
}

// Value sets the input value
func (s *SpectrumSearch) Value(value string) *SpectrumSearch {
	s.value = value
	return s
}

// Placeholder sets the placeholder text
func (s *SpectrumSearch) Placeholder(placeholder string) *SpectrumSearch {
	s.placeholder = placeholder
	return s
}

// Name sets the form control name
func (s *SpectrumSearch) Name(name string) *SpectrumSearch {
	s.name = name
	return s
}

// Autocomplete sets the autocomplete attribute
func (s *SpectrumSearch) Autocomplete(autocomplete string) *SpectrumSearch {
	s.autocomplete = autocomplete
	return s
}

// Disabled sets the disabled state
func (s *SpectrumSearch) Disabled(disabled bool) *SpectrumSearch {
	s.disabled = disabled
	return s
}

// Grows sets whether multiline search grows with content
func (s *SpectrumSearch) Grows(grows bool) *SpectrumSearch {
	s.grows = grows
	return s
}

// Invalid sets the invalid state
func (s *SpectrumSearch) Invalid(invalid bool) *SpectrumSearch {
	s.invalid = invalid
	return s
}

// Maxlength sets the maximum allowed length
func (s *SpectrumSearch) Maxlength(maxlength int) *SpectrumSearch {
	s.maxlength = maxlength
	return s
}

// Minlength sets the minimum allowed length
func (s *SpectrumSearch) Minlength(minlength int) *SpectrumSearch {
	s.minlength = minlength
	return s
}

// Multiline sets whether the search accepts multiple lines
func (s *SpectrumSearch) Multiline(multiline bool) *SpectrumSearch {
	s.multiline = multiline
	return s
}

// Pattern sets the validation pattern
func (s *SpectrumSearch) Pattern(pattern string) *SpectrumSearch {
	s.pattern = pattern
	return s
}

// Quiet sets whether the search uses quiet styling
func (s *SpectrumSearch) Quiet(quiet bool) *SpectrumSearch {
	s.quiet = quiet
	return s
}

// Readonly sets the readonly state
func (s *SpectrumSearch) Readonly(readonly bool) *SpectrumSearch {
	s.readonly = readonly
	return s
}

// Required sets the required state
func (s *SpectrumSearch) Required(required bool) *SpectrumSearch {
	s.required = required
	return s
}

// Rows sets the number of visible rows for multiline search
func (s *SpectrumSearch) Rows(rows int) *SpectrumSearch {
	s.rows = rows
	return s
}

// TabIndex sets the tab index
func (s *SpectrumSearch) TabIndex(tabIndex int) *SpectrumSearch {
	s.tabIndex = tabIndex
	return s
}

// HoldValueOnEscape sets whether the value is held when escape key is pressed
func (s *SpectrumSearch) HoldValueOnEscape(hold bool) *SpectrumSearch {
	s.holdValueOnEscape = hold
	return s
}

// HelpText sets the help text UI element for the help-text slot
func (s *SpectrumSearch) HelpText(helpText app.UI) *SpectrumSearch {
	s.helpText = helpText
	return s
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (s *SpectrumSearch) NegativeHelpText(helpText app.UI) *SpectrumSearch {
	s.negativeHelpText = helpText
	return s
}

// OnInput sets the input event handler
func (s *SpectrumSearch) OnInput(handler app.EventHandler) *SpectrumSearch {
	s.onInput = handler
	return s
}

// OnChange sets the change event handler
func (s *SpectrumSearch) OnChange(handler app.EventHandler) *SpectrumSearch {
	s.onChange = handler
	return s
}

// OnSubmit sets the submit event handler
func (s *SpectrumSearch) OnSubmit(handler app.EventHandler) *SpectrumSearch {
	s.onSubmit = handler
	return s
}

// Render renders the search component
func (s *SpectrumSearch) Render() app.UI {
	search := app.Elem("sp-search").
		Attr("size", string(s.size)).
		Attr("action", s.action).
		Attr("method", string(s.method)).
		Attr("label", s.label).
		Attr("value", s.value).
		Attr("placeholder", s.placeholder).
		Attr("name", s.name).
		Attr("autocomplete", s.autocomplete).
		Attr("disabled", s.disabled).
		Attr("grows", s.grows).
		Attr("invalid", s.invalid).
		Attr("maxlength", s.maxlength).
		Attr("minlength", s.minlength).
		Attr("multiline", s.multiline).
		Attr("pattern", s.pattern).
		Attr("quiet", s.quiet).
		Attr("readonly", s.readonly).
		Attr("required", s.required).
		Attr("rows", s.rows).
		Attr("tabindex", s.tabIndex).
		Attr("holdValueOnEscape", s.holdValueOnEscape)

	// Add event handlers
	if s.onInput != nil {
		search = search.On("input", s.onInput)
	}
	if s.onChange != nil {
		search = search.On("change", s.onChange)
	}
	if s.onSubmit != nil {
		search = search.On("submit", s.onSubmit)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if s.helpText != nil {
		helpTextElem := s.helpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if s.negativeHelpText != nil {
		negativeHelpTextElem := s.negativeHelpText
		if negativeHelpTextWithSlot, ok := negativeHelpTextElem.(interface{ Slot(string) app.UI }); ok {
			negativeHelpTextElem = negativeHelpTextWithSlot.Slot("negative-help-text")
		} else {
			negativeHelpTextElem = app.Elem("div").
				Attr("slot", "negative-help-text").
				Body(negativeHelpTextElem)
		}
		elements = append(elements, negativeHelpTextElem)
	}

	// Add all elements to the search
	if len(elements) > 0 {
		search = search.Body(elements...)
	}

	return search
}
