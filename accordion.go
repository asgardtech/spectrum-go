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

// SpectrumAccordion represents an sp-accordion component
type SpectrumAccordion struct {
	app.Compo

	// Properties
	size          AccordionSize
	density       AccordionDensity
	allowMultiple bool

	// Children items
	children []*SpectrumAccordionItem
}

// Accordion creates a new accordion component
func Accordion() *SpectrumAccordion {
	return &SpectrumAccordion{
		size: AccordionSizeM, // Default size is medium
	}
}

// Size sets the visual size of the accordion
func (a *SpectrumAccordion) Size(size AccordionSize) *SpectrumAccordion {
	a.size = size
	return a
}

// Density sets the density of the accordion
func (a *SpectrumAccordion) Density(density AccordionDensity) *SpectrumAccordion {
	a.density = density
	return a
}

// AllowMultiple sets whether multiple accordion items can be expanded at once
func (a *SpectrumAccordion) AllowMultiple(allow bool) *SpectrumAccordion {
	a.allowMultiple = allow
	return a
}

// Children sets the accordion items
func (a *SpectrumAccordion) Children(children ...*SpectrumAccordionItem) *SpectrumAccordion {
	a.children = children
	return a
}

// Render renders the accordion component
func (a *SpectrumAccordion) Render() app.UI {
	accordion := app.Elem("sp-accordion").
		Attr("size", string(a.size)).
		Attr("density", string(a.density)).
		Attr("allow-multiple", a.allowMultiple)

	// Add all child elements
	if len(a.children) > 0 {
		// Convert SpectrumAccordionItem children to app.UI for the Body method
		uiChildren := make([]app.UI, len(a.children))
		for i, child := range a.children {
			uiChildren[i] = child
		}
		accordion = accordion.Body(uiChildren...)
	}

	return accordion
}
