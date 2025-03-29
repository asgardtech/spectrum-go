package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ContextualHelpPlacement represents the placement positions for the contextual help popover
type ContextualHelpPlacement string

// Contextual help placement positions
const (
	ContextualHelpPlacementTop         ContextualHelpPlacement = "top"
	ContextualHelpPlacementTopStart    ContextualHelpPlacement = "top-start"
	ContextualHelpPlacementTopEnd      ContextualHelpPlacement = "top-end"
	ContextualHelpPlacementRight       ContextualHelpPlacement = "right"
	ContextualHelpPlacementRightStart  ContextualHelpPlacement = "right-start"
	ContextualHelpPlacementRightEnd    ContextualHelpPlacement = "right-end"
	ContextualHelpPlacementBottom      ContextualHelpPlacement = "bottom"
	ContextualHelpPlacementBottomStart ContextualHelpPlacement = "bottom-start"
	ContextualHelpPlacementBottomEnd   ContextualHelpPlacement = "bottom-end"
	ContextualHelpPlacementLeft        ContextualHelpPlacement = "left"
	ContextualHelpPlacementLeftStart   ContextualHelpPlacement = "left-start"
	ContextualHelpPlacementLeftEnd     ContextualHelpPlacement = "left-end"
)

// ContextualHelpVariant represents the variant of the contextual help
type ContextualHelpVariant string

// Contextual help variants
const (
	ContextualHelpVariantInfo ContextualHelpVariant = "info"
	ContextualHelpVariantHelp ContextualHelpVariant = "help"
)

// spectrumContextualHelp represents an sp-contextual-help component
type spectrumContextualHelp struct {
	app.Compo

	// Properties
	PropLabel     string
	PropOffset    interface{} // number or [number, number]
	PropPlacement ContextualHelpPlacement
	PropVariant   ContextualHelpVariant
	PropMaxWidth  string

	// Content elements
	PropHeading app.UI
	PropContent app.UI
	PropLink    app.UI
}

// ContextualHelp creates a new contextual help component
func ContextualHelp() *spectrumContextualHelp {
	return &spectrumContextualHelp{
		PropPlacement: ContextualHelpPlacementBottomStart,
		PropVariant:   ContextualHelpVariantInfo,
	}
}

// Label sets the accessible name for the action button trigger
func (c *spectrumContextualHelp) Label(label string) *spectrumContextualHelp {
	c.PropLabel = label
	return c
}

// Offset sets the offset of the popover
// Can be a single number for the main axis, or [mainAxis, crossAxis]
func (c *spectrumContextualHelp) Offset(offset interface{}) *spectrumContextualHelp {
	c.PropOffset = offset
	return c
}

// Placement sets the placement of the popover
func (c *spectrumContextualHelp) Placement(placement ContextualHelpPlacement) *spectrumContextualHelp {
	c.PropPlacement = placement
	return c
}

// Variant sets the variant of the contextual help (info/help)
func (c *spectrumContextualHelp) Variant(variant ContextualHelpVariant) *spectrumContextualHelp {
	c.PropVariant = variant
	return c
}

// MaxWidth sets the custom maximum width of the popover
func (c *spectrumContextualHelp) MaxWidth(maxWidth string) *spectrumContextualHelp {
	c.PropMaxWidth = maxWidth
	return c
}

// Heading sets the heading content
func (c *spectrumContextualHelp) Heading(heading app.UI) *spectrumContextualHelp {
	c.PropHeading = heading
	return c
}

// Content sets the main content
func (c *spectrumContextualHelp) Content(content app.UI) *spectrumContextualHelp {
	c.PropContent = content
	return c
}

// Link sets the link content
func (c *spectrumContextualHelp) Link(link app.UI) *spectrumContextualHelp {
	c.PropLink = link
	return c
}

// Render renders the contextual help component
func (c *spectrumContextualHelp) Render() app.UI {
	contextualHelp := app.Elem("sp-contextual-help")

	// Set attributes
	if c.PropLabel != "" {
		contextualHelp = contextualHelp.Attr("label", c.PropLabel)
	}

	if c.PropOffset != nil {
		switch offset := c.PropOffset.(type) {
		case int:
			contextualHelp = contextualHelp.Attr("offset", offset)
		case []int:
			if len(offset) == 2 {
				// Convert to JSON array for [mainAxis, crossAxis]
				contextualHelp = contextualHelp.Attr("offset", offset)
			}
		}
	}

	if c.PropPlacement != ContextualHelpPlacementBottomStart {
		contextualHelp = contextualHelp.Attr("placement", string(c.PropPlacement))
	}

	if c.PropVariant != ContextualHelpVariantInfo {
		contextualHelp = contextualHelp.Attr("variant", string(c.PropVariant))
	}

	if c.PropMaxWidth != "" {
		contextualHelp = contextualHelp.Style("--mod-spectrum-contextual-help-popover-maximum-width", c.PropMaxWidth)
	}

	// Prepare body elements
	elements := []app.UI{}

	// Add slot elements
	if c.PropHeading != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "heading").
			Body(c.PropHeading))
	}

	if c.PropContent != nil {
		elements = append(elements, c.PropContent)
	}

	if c.PropLink != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "link").
			Body(c.PropLink))
	}

	// Add all elements to the contextual help
	if len(elements) > 0 {
		contextualHelp = contextualHelp.Body(elements...)
	}

	return contextualHelp
}
