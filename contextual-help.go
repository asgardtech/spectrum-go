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

// SpectrumContextualHelp represents an sp-contextual-help component
type SpectrumContextualHelp struct {
	app.Compo

	// Properties
	label     string
	offset    interface{} // number or [number, number]
	placement ContextualHelpPlacement
	variant   ContextualHelpVariant
	maxWidth  string

	// Content elements
	heading app.UI
	content app.UI
	link    app.UI
}

// ContextualHelp creates a new contextual help component
func ContextualHelp() *SpectrumContextualHelp {
	return &SpectrumContextualHelp{
		placement: ContextualHelpPlacementBottomStart,
		variant:   ContextualHelpVariantInfo,
	}
}

// Label sets the accessible name for the action button trigger
func (c *SpectrumContextualHelp) Label(label string) *SpectrumContextualHelp {
	c.label = label
	return c
}

// Offset sets the offset of the popover
// Can be a single number for the main axis, or [mainAxis, crossAxis]
func (c *SpectrumContextualHelp) Offset(offset interface{}) *SpectrumContextualHelp {
	c.offset = offset
	return c
}

// Placement sets the placement of the popover
func (c *SpectrumContextualHelp) Placement(placement ContextualHelpPlacement) *SpectrumContextualHelp {
	c.placement = placement
	return c
}

// Variant sets the variant of the contextual help (info/help)
func (c *SpectrumContextualHelp) Variant(variant ContextualHelpVariant) *SpectrumContextualHelp {
	c.variant = variant
	return c
}

// MaxWidth sets the custom maximum width of the popover
func (c *SpectrumContextualHelp) MaxWidth(maxWidth string) *SpectrumContextualHelp {
	c.maxWidth = maxWidth
	return c
}

// Heading sets the heading content
func (c *SpectrumContextualHelp) Heading(heading app.UI) *SpectrumContextualHelp {
	c.heading = heading
	return c
}

// Content sets the main content
func (c *SpectrumContextualHelp) Content(content app.UI) *SpectrumContextualHelp {
	c.content = content
	return c
}

// Link sets the link content
func (c *SpectrumContextualHelp) Link(link app.UI) *SpectrumContextualHelp {
	c.link = link
	return c
}

// Render renders the contextual help component
func (c *SpectrumContextualHelp) Render() app.UI {
	contextualHelp := app.Elem("sp-contextual-help")

	// Set attributes
	if c.label != "" {
		contextualHelp = contextualHelp.Attr("label", c.label)
	}

	if c.offset != nil {
		switch offset := c.offset.(type) {
		case int:
			contextualHelp = contextualHelp.Attr("offset", offset)
		case []int:
			if len(offset) == 2 {
				// Convert to JSON array for [mainAxis, crossAxis]
				contextualHelp = contextualHelp.Attr("offset", offset)
			}
		}
	}

	if c.placement != ContextualHelpPlacementBottomStart {
		contextualHelp = contextualHelp.Attr("placement", string(c.placement))
	}

	if c.variant != ContextualHelpVariantInfo {
		contextualHelp = contextualHelp.Attr("variant", string(c.variant))
	}

	if c.maxWidth != "" {
		contextualHelp = contextualHelp.Style("--mod-spectrum-contextual-help-popover-maximum-width", c.maxWidth)
	}

	// Prepare body elements
	elements := []app.UI{}

	// Add slot elements
	if c.heading != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "heading").
			Body(c.heading))
	}

	if c.content != nil {
		elements = append(elements, c.content)
	}

	if c.link != nil {
		elements = append(elements, app.Elem("div").
			Attr("slot", "link").
			Body(c.link))
	}

	// Add all elements to the contextual help
	if len(elements) > 0 {
		contextualHelp = contextualHelp.Body(elements...)
	}

	return contextualHelp
}
