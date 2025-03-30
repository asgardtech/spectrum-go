package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconsWorkflowSize represents the Size of the icon
type IconsWorkflowSize string

// IconsWorkflowSize values
const (
    IconsWorkflowSizeXxs IconsWorkflowSize = "xxs"
    IconsWorkflowSizeXs IconsWorkflowSize = "xs"
    IconsWorkflowSizeS IconsWorkflowSize = "s"
    IconsWorkflowSizeM IconsWorkflowSize = "m"
    IconsWorkflowSizeL IconsWorkflowSize = "l"
    IconsWorkflowSizeXl IconsWorkflowSize = "xl"
    IconsWorkflowSizeXxl IconsWorkflowSize = "xxl"
)

// spectrumIconsWorkflow represents an sp-icon-workflow component
type spectrumIconsWorkflow struct {
    app.Compo
    *styler[*spectrumIconsWorkflow]

    // Properties
    // Accessibility label for the icon
    PropLabel string
    // Size of the icon
    PropSize IconsWorkflowSize




}

// IconsWorkflow creates a new sp-icon-workflow component
func IconsWorkflow() *spectrumIconsWorkflow {
    element := &spectrumIconsWorkflow{
        PropLabel: "",
        PropSize: "",
    }

    element.styler = newStyler(element)

    return element
}

// Accessibility label for the icon
func (c *spectrumIconsWorkflow) Label(label string) *spectrumIconsWorkflow {
    c.PropLabel = label
    return c
}

// Size of the icon
func (c *spectrumIconsWorkflow) Size(size IconsWorkflowSize) *spectrumIconsWorkflow {
    c.PropSize = size
    return c
}

func (c *spectrumIconsWorkflow) SizeXxs() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeXxs)
}
func (c *spectrumIconsWorkflow) SizeXs() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeXs)
}
func (c *spectrumIconsWorkflow) SizeS() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeS)
}
func (c *spectrumIconsWorkflow) SizeM() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeM)
}
func (c *spectrumIconsWorkflow) SizeL() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeL)
}
func (c *spectrumIconsWorkflow) SizeXl() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeXl)
}
func (c *spectrumIconsWorkflow) SizeXxl() *spectrumIconsWorkflow {
    return c.Size(IconsWorkflowSizeXxl)
}





// Render renders the sp-icon-workflow component
func (c *spectrumIconsWorkflow) Render() app.UI {
    element := app.Elem("sp-icon-workflow")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }



    element = element.Styles(c.styler.styles)

    return element
} 