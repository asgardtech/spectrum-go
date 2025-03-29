package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumAccordionItem represents an sp-accordion-item component
type spectrumAccordionItem struct {
	app.Compo

	// Properties
	PropLabel    string
	PropOpen     bool
	PropDisabled bool
	PropTabIndex int

	// Content
	PropContent []app.UI
}

// AccordionItem creates a new accordion item component
func AccordionItem() *spectrumAccordionItem {
	return &spectrumAccordionItem{}
}

// Label sets the label text content
func (ai *spectrumAccordionItem) Label(label string) *spectrumAccordionItem {
	ai.PropLabel = label
	return ai
}

// Open sets the accordion item's open state
func (ai *spectrumAccordionItem) Open(open bool) *spectrumAccordionItem {
	ai.PropOpen = open
	return ai
}

// Disabled sets the accordion item's disabled state
func (ai *spectrumAccordionItem) Disabled(disabled bool) *spectrumAccordionItem {
	ai.PropDisabled = disabled
	return ai
}

// TabIndex sets the tab index for the accordion item
func (ai *spectrumAccordionItem) TabIndex(index int) *spectrumAccordionItem {
	ai.PropTabIndex = index
	return ai
}

// Content sets the content of the accordion item
func (ai *spectrumAccordionItem) Content(content ...app.UI) *spectrumAccordionItem {
	ai.PropContent = content
	return ai
}

// Render renders the accordion item component
func (ai *spectrumAccordionItem) Render() app.UI {
	accordionItem := app.Elem("sp-accordion-item").
		Attr("label", ai.PropLabel).
		Attr("open", ai.PropOpen).
		Attr("disabled", ai.PropDisabled).
		Attr("tabindex", ai.PropTabIndex)

	// Add content if provided
	if len(ai.PropContent) > 0 {
		accordionItem.Body(ai.PropContent...)
	}

	return accordionItem
}
