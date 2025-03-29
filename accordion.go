package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AccordionSize represents the visual size of an accordion
type AccordionSize string

// Accordion sizes
const (
	AccordionSizeS  AccordionSize = "s"
	AccordionSizeM  AccordionSize = "m"
	AccordionSizeL  AccordionSize = "l"
	AccordionSizeXL AccordionSize = "xl"
)

// AccordionDensity represents the density of an accordion
type AccordionDensity string

// Accordion densities
const (
	AccordionDensityCompact     AccordionDensity = "compact"
	AccordionDensitySpaciousNon AccordionDensity = "spacious"
)

// spectrumAccordion represents an sp-accordion component
type spectrumAccordion struct {
	app.Compo

	// Properties
	PropSize          AccordionSize
	PropDensity       AccordionDensity
	PropAllowMultiple bool

	// Children items
	PropChildren []*spectrumAccordionItem
}

// Accordion creates a new accordion component
func Accordion() *spectrumAccordion {
	return &spectrumAccordion{
		PropSize: AccordionSizeM, // Default size is medium
	}
}

// Size sets the visual size of the accordion
func (a *spectrumAccordion) Size(size AccordionSize) *spectrumAccordion {
	a.PropSize = size
	return a
}

// Density sets the density of the accordion
func (a *spectrumAccordion) Density(density AccordionDensity) *spectrumAccordion {
	a.PropDensity = density
	return a
}

// AllowMultiple sets whether multiple accordion items can be expanded at once
func (a *spectrumAccordion) AllowMultiple(allow bool) *spectrumAccordion {
	a.PropAllowMultiple = allow
	return a
}

// Children sets the accordion items
func (a *spectrumAccordion) Children(children ...*spectrumAccordionItem) *spectrumAccordion {
	a.PropChildren = children
	return a
}

// Render renders the accordion component
func (a *spectrumAccordion) Render() app.UI {
	accordion := app.Elem("sp-accordion").
		Attr("size", string(a.PropSize)).
		Attr("density", string(a.PropDensity)).
		Attr("allow-multiple", a.PropAllowMultiple)

	// Add all child elements
	if len(a.PropChildren) > 0 {
		// Convert spectrumAccordionItem children to app.UI for the Body method
		uiChildren := make([]app.UI, len(a.PropChildren))
		for i, child := range a.PropChildren {
			uiChildren[i] = child
		}
		accordion = accordion.Body(uiChildren...)
	}

	return accordion
}
