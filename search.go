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

// spectrumSearch represents an sp-search component
type spectrumSearch struct {
	app.Compo

	// Properties
	PropSize              SearchSize
	PropAction            string
	PropMethod            SearchMethod
	PropLabel             string
	PropValue             string
	PropPlaceholder       string
	PropName              string
	PropAutocomplete      string
	PropDisabled          bool
	PropGrows             bool
	PropInvalid           bool
	PropMaxlength         int
	PropMinlength         int
	PropMultiline         bool
	PropPattern           string
	PropQuiet             bool
	PropReadonly          bool
	PropRequired          bool
	PropRows              int
	PropTabIndex          int
	PropHoldValueOnEscape bool

	// Help text slots
	PropHelpText         app.UI
	PropNegativeHelpText app.UI

	// Event handlers
	PropOnInput  app.EventHandler
	PropOnChange app.EventHandler
	PropOnSubmit app.EventHandler
}

// Search creates a new search component
func Search() *spectrumSearch {
	return &spectrumSearch{
		PropSize:        SearchSizeM, // Default size is medium
		PropPlaceholder: "Search",    // Default placeholder
		PropLabel:       "Search",    // Default accessibility label
	}
}

// Size sets the visual size of the search field
func (s *spectrumSearch) Size(size SearchSize) *spectrumSearch {
	s.PropSize = size
	return s
}

// Action sets the form action URL
func (s *spectrumSearch) Action(action string) *spectrumSearch {
	s.PropAction = action
	return s
}

// Method sets the form submission method
func (s *spectrumSearch) Method(method SearchMethod) *spectrumSearch {
	s.PropMethod = method
	return s
}

// Label sets the accessibility label
func (s *spectrumSearch) Label(label string) *spectrumSearch {
	s.PropLabel = label
	return s
}

// Value sets the input value
func (s *spectrumSearch) Value(value string) *spectrumSearch {
	s.PropValue = value
	return s
}

// Placeholder sets the placeholder text
func (s *spectrumSearch) Placeholder(placeholder string) *spectrumSearch {
	s.PropPlaceholder = placeholder
	return s
}

// Name sets the form control name
func (s *spectrumSearch) Name(name string) *spectrumSearch {
	s.PropName = name
	return s
}

// Autocomplete sets the autocomplete attribute
func (s *spectrumSearch) Autocomplete(autocomplete string) *spectrumSearch {
	s.PropAutocomplete = autocomplete
	return s
}

// Disabled sets the disabled state
func (s *spectrumSearch) Disabled(disabled bool) *spectrumSearch {
	s.PropDisabled = disabled
	return s
}

// Grows sets whether multiline search grows with content
func (s *spectrumSearch) Grows(grows bool) *spectrumSearch {
	s.PropGrows = grows
	return s
}

// Invalid sets the invalid state
func (s *spectrumSearch) Invalid(invalid bool) *spectrumSearch {
	s.PropInvalid = invalid
	return s
}

// Maxlength sets the maximum allowed length
func (s *spectrumSearch) Maxlength(maxlength int) *spectrumSearch {
	s.PropMaxlength = maxlength
	return s
}

// Minlength sets the minimum allowed length
func (s *spectrumSearch) Minlength(minlength int) *spectrumSearch {
	s.PropMinlength = minlength
	return s
}

// Multiline sets whether the search accepts multiple lines
func (s *spectrumSearch) Multiline(multiline bool) *spectrumSearch {
	s.PropMultiline = multiline
	return s
}

// Pattern sets the validation pattern
func (s *spectrumSearch) Pattern(pattern string) *spectrumSearch {
	s.PropPattern = pattern
	return s
}

// Quiet sets whether the search uses quiet styling
func (s *spectrumSearch) Quiet(quiet bool) *spectrumSearch {
	s.PropQuiet = quiet
	return s
}

// Readonly sets the readonly state
func (s *spectrumSearch) Readonly(readonly bool) *spectrumSearch {
	s.PropReadonly = readonly
	return s
}

// Required sets the required state
func (s *spectrumSearch) Required(required bool) *spectrumSearch {
	s.PropRequired = required
	return s
}

// Rows sets the number of visible rows for multiline search
func (s *spectrumSearch) Rows(rows int) *spectrumSearch {
	s.PropRows = rows
	return s
}

// TabIndex sets the tab index
func (s *spectrumSearch) TabIndex(tabIndex int) *spectrumSearch {
	s.PropTabIndex = tabIndex
	return s
}

// HoldValueOnEscape sets whether the value is held when escape key is pressed
func (s *spectrumSearch) HoldValueOnEscape(hold bool) *spectrumSearch {
	s.PropHoldValueOnEscape = hold
	return s
}

// HelpText sets the help text UI element for the help-text slot
func (s *spectrumSearch) HelpText(helpText app.UI) *spectrumSearch {
	s.PropHelpText = helpText
	return s
}

// NegativeHelpText sets the negative help text UI element for the negative-help-text slot
func (s *spectrumSearch) NegativeHelpText(helpText app.UI) *spectrumSearch {
	s.PropNegativeHelpText = helpText
	return s
}

// OnInput sets the input event handler
func (s *spectrumSearch) OnInput(handler app.EventHandler) *spectrumSearch {
	s.PropOnInput = handler
	return s
}

// OnChange sets the change event handler
func (s *spectrumSearch) OnChange(handler app.EventHandler) *spectrumSearch {
	s.PropOnChange = handler
	return s
}

// OnSubmit sets the submit event handler
func (s *spectrumSearch) OnSubmit(handler app.EventHandler) *spectrumSearch {
	s.PropOnSubmit = handler
	return s
}

// Render renders the search component
func (s *spectrumSearch) Render() app.UI {
	search := app.Elem("sp-search").
		Attr("size", string(s.PropSize)).
		Attr("action", s.PropAction).
		Attr("method", string(s.PropMethod)).
		Attr("label", s.PropLabel).
		Attr("value", s.PropValue).
		Attr("placeholder", s.PropPlaceholder).
		Attr("name", s.PropName).
		Attr("autocomplete", s.PropAutocomplete).
		Attr("disabled", s.PropDisabled).
		Attr("grows", s.PropGrows).
		Attr("invalid", s.PropInvalid).
		Attr("maxlength", s.PropMaxlength).
		Attr("minlength", s.PropMinlength).
		Attr("multiline", s.PropMultiline).
		Attr("pattern", s.PropPattern).
		Attr("quiet", s.PropQuiet).
		Attr("readonly", s.PropReadonly).
		Attr("required", s.PropRequired).
		Attr("rows", s.PropRows).
		Attr("tabindex", s.PropTabIndex).
		Attr("holdValueOnEscape", s.PropHoldValueOnEscape)

	// Add event handlers
	if s.PropOnInput != nil {
		search = search.On("input", s.PropOnInput)
	}
	if s.PropOnChange != nil {
		search = search.On("change", s.PropOnChange)
	}
	if s.PropOnSubmit != nil {
		search = search.On("submit", s.PropOnSubmit)
	}

	// Add help text slots if provided
	elements := []app.UI{}

	if s.PropHelpText != nil {
		helpTextElem := s.PropHelpText
		if helpTextWithSlot, ok := helpTextElem.(interface{ Slot(string) app.UI }); ok {
			helpTextElem = helpTextWithSlot.Slot("help-text")
		} else {
			helpTextElem = app.Elem("div").
				Attr("slot", "help-text").
				Body(helpTextElem)
		}
		elements = append(elements, helpTextElem)
	}

	if s.PropNegativeHelpText != nil {
		negativeHelpTextElem := s.PropNegativeHelpText
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
