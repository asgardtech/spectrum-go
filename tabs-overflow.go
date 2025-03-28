package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTabsOverflow represents an sp-tabs-overflow component
type SpectrumTabsOverflow struct {
	app.Compo

	// Properties
	compact       bool
	labelNext     string
	labelPrevious string

	// Child component
	tabs app.UI
}

// TabsOverflow creates a new tabs overflow component
func TabsOverflow() *SpectrumTabsOverflow {
	return &SpectrumTabsOverflow{
		labelNext:     "Scroll to next tabs",     // Default label
		labelPrevious: "Scroll to previous tabs", // Default label
	}
}

// Compact sets whether the overflow is displayed compact
func (to *SpectrumTabsOverflow) Compact(compact bool) *SpectrumTabsOverflow {
	to.compact = compact
	return to
}

// LabelNext sets the accessible label for the next button
func (to *SpectrumTabsOverflow) LabelNext(label string) *SpectrumTabsOverflow {
	to.labelNext = label
	return to
}

// LabelPrevious sets the accessible label for the previous button
func (to *SpectrumTabsOverflow) LabelPrevious(label string) *SpectrumTabsOverflow {
	to.labelPrevious = label
	return to
}

// Tabs sets the tabs component to be wrapped
func (to *SpectrumTabsOverflow) Tabs(tabs *SpectrumTabs) *SpectrumTabsOverflow {
	to.tabs = tabs
	return to
}

// Child sets any UI to be wrapped (typically a SpectrumTabs)
func (to *SpectrumTabsOverflow) Child(child app.UI) *SpectrumTabsOverflow {
	to.tabs = child
	return to
}

// Render renders the tabs overflow component
func (to *SpectrumTabsOverflow) Render() app.UI {
	overflow := app.Elem("sp-tabs-overflow")

	// Set attributes
	if to.compact {
		overflow = overflow.Attr("compact", true)
	}
	if to.labelNext != "Scroll to next tabs" {
		overflow = overflow.Attr("label-next", to.labelNext)
	}
	if to.labelPrevious != "Scroll to previous tabs" {
		overflow = overflow.Attr("label-previous", to.labelPrevious)
	}

	// Add the tabs component
	if to.tabs != nil {
		overflow = overflow.Body(to.tabs)
	}

	return overflow
}
