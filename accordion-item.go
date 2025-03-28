package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumAccordionItem represents an sp-accordion-item component
type SpectrumAccordionItem struct {
	app.Compo

	// Properties
	label    string
	open     bool
	disabled bool
	tabIndex int

	// Content
	content []app.UI
}

// AccordionItem creates a new accordion item component
func AccordionItem() *SpectrumAccordionItem {
	return &SpectrumAccordionItem{}
}

// Label sets the label text content
func (ai *SpectrumAccordionItem) Label(label string) *SpectrumAccordionItem {
	ai.label = label
	return ai
}

// Open sets the accordion item's open state
func (ai *SpectrumAccordionItem) Open(open bool) *SpectrumAccordionItem {
	ai.open = open
	return ai
}

// Disabled sets the accordion item's disabled state
func (ai *SpectrumAccordionItem) Disabled(disabled bool) *SpectrumAccordionItem {
	ai.disabled = disabled
	return ai
}

// TabIndex sets the tab index for the accordion item
func (ai *SpectrumAccordionItem) TabIndex(index int) *SpectrumAccordionItem {
	ai.tabIndex = index
	return ai
}

// Content sets the content of the accordion item
func (ai *SpectrumAccordionItem) Content(content ...app.UI) *SpectrumAccordionItem {
	ai.content = content
	return ai
}

// Render renders the accordion item component
func (ai *SpectrumAccordionItem) Render() app.UI {
	accordionItem := app.Elem("sp-accordion-item").
		Attr("label", ai.label).
		Attr("open", ai.open).
		Attr("disabled", ai.disabled).
		Attr("tabindex", ai.tabIndex)

	// Add content if provided
	if len(ai.content) > 0 {
		accordionItem.Body(ai.content...)
	}

	return accordionItem
}
