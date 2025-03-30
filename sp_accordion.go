package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AccordionDensity represents the Sets the spacing between the content to borders of an accordion item
type AccordionDensity string

// AccordionDensity values
const (
    AccordionDensityCompact AccordionDensity = "compact"
    AccordionDensitySpacious AccordionDensity = "spacious"
)
// AccordionSize represents the The size of the accordion
type AccordionSize string

// AccordionSize values
const (
    AccordionSizeS AccordionSize = "s"
    AccordionSizeM AccordionSize = "m"
    AccordionSizeL AccordionSize = "l"
    AccordionSizeXl AccordionSize = "xl"
)

// spectrumAccordion represents an sp-accordion component
type spectrumAccordion struct {
    app.Compo
    *styler[*spectrumAccordion]

    // Properties
    // Allows multiple accordion items to be opened at the same time
    PropAllowMultiple bool
    // Sets the spacing between the content to borders of an accordion item
    PropDensity AccordionDensity
    // The size of the accordion
    PropSize AccordionSize


    // Content slots

    // Child components
    PropChildren []app.UI

}

// Accordion creates a new sp-accordion component
func Accordion() *spectrumAccordion {
    element := &spectrumAccordion{
        PropAllowMultiple: false,
        PropSize: AccordionSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Allows multiple accordion items to be opened at the same time
func (c *spectrumAccordion) AllowMultiple(allowMultiple bool) *spectrumAccordion {
    c.PropAllowMultiple = allowMultiple
    return c
}

func (c *spectrumAccordion) SetAllowMultiple() *spectrumAccordion {
    return c.AllowMultiple(true)
}

// Sets the spacing between the content to borders of an accordion item
func (c *spectrumAccordion) Density(density AccordionDensity) *spectrumAccordion {
    c.PropDensity = density
    return c
}

func (c *spectrumAccordion) DensityCompact() *spectrumAccordion {
    return c.Density(AccordionDensityCompact)
}
func (c *spectrumAccordion) DensitySpacious() *spectrumAccordion {
    return c.Density(AccordionDensitySpacious)
}
// The size of the accordion
func (c *spectrumAccordion) Size(size AccordionSize) *spectrumAccordion {
    c.PropSize = size
    return c
}

func (c *spectrumAccordion) SizeS() *spectrumAccordion {
    return c.Size(AccordionSizeS)
}
func (c *spectrumAccordion) SizeM() *spectrumAccordion {
    return c.Size(AccordionSizeM)
}
func (c *spectrumAccordion) SizeL() *spectrumAccordion {
    return c.Size(AccordionSizeL)
}
func (c *spectrumAccordion) SizeXl() *spectrumAccordion {
    return c.Size(AccordionSizeXl)
}



// Children sets the child components
func (c *spectrumAccordion) Children(children ...app.UI) *spectrumAccordion {
    c.PropChildren = children

    return c
}

// AddChild adds a child component
func (c *spectrumAccordion) AddChild(child app.UI) *spectrumAccordion {
    c.PropChildren = append(c.PropChildren, child)

    return c
}



// Render renders the sp-accordion component
func (c *spectrumAccordion) Render() app.UI {
    element := app.Elem("sp-accordion")

    // Set attributes
    if c.PropAllowMultiple {
        element = element.Attr("allow-multiple", true)
    }
    if c.PropDensity != "" {
        element = element.Attr("density", string(c.PropDensity))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }


    // Add slots and children
    slotElements := []app.UI{}



    // Add children
    if len(c.PropChildren) > 0 {
        for _, child := range c.PropChildren {
            slotElements = append(slotElements, child)
        }
    }

    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 