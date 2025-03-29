package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TooltipPlacement represents the position of the tooltip relative to its target
type TooltipPlacement string

// Tooltip placements
const (
	TooltipPlacementTop         TooltipPlacement = "top"
	TooltipPlacementTopStart    TooltipPlacement = "top-start"
	TooltipPlacementTopEnd      TooltipPlacement = "top-end"
	TooltipPlacementRight       TooltipPlacement = "right"
	TooltipPlacementRightStart  TooltipPlacement = "right-start"
	TooltipPlacementRightEnd    TooltipPlacement = "right-end"
	TooltipPlacementBottom      TooltipPlacement = "bottom"
	TooltipPlacementBottomStart TooltipPlacement = "bottom-start"
	TooltipPlacementBottomEnd   TooltipPlacement = "bottom-end"
	TooltipPlacementLeft        TooltipPlacement = "left"
	TooltipPlacementLeftStart   TooltipPlacement = "left-start"
	TooltipPlacementLeftEnd     TooltipPlacement = "left-end"
)

// TooltipVariant represents the visual variant of a tooltip
type TooltipVariant string

// Tooltip variants
const (
	TooltipVariantInfo     TooltipVariant = "info"
	TooltipVariantPositive TooltipVariant = "positive"
	TooltipVariantNegative TooltipVariant = "negative"
)

// spectrumTooltip represents an sp-tooltip component
type spectrumTooltip struct {
	app.Compo

	// Properties
	PropOpen        bool
	PropPlacement   TooltipPlacement
	PropOffset      int
	PropSelfManaged bool
	PropTipPadding  int
	PropVariant     TooltipVariant
	PropSlot        string
	PropContent     string
	PropInnerHTML   string

	// Child elements
	PropIcon  app.UI
	PropChild app.UI
}

// Tooltip creates a new tooltip component
func Tooltip() *spectrumTooltip {
	return &spectrumTooltip{}
}

// Open sets whether the tooltip is visible
func (t *spectrumTooltip) Open(open bool) *spectrumTooltip {
	t.PropOpen = open
	return t
}

// Placement sets the position of the tooltip relative to its target
func (t *spectrumTooltip) Placement(placement TooltipPlacement) *spectrumTooltip {
	t.PropPlacement = placement
	return t
}

// Offset sets the distance from the tooltip to its target
func (t *spectrumTooltip) Offset(offset int) *spectrumTooltip {
	t.PropOffset = offset
	return t
}

// SelfManaged sets whether the tooltip auto-binds to its parent element
func (t *spectrumTooltip) SelfManaged(selfManaged bool) *spectrumTooltip {
	t.PropSelfManaged = selfManaged
	return t
}

// TipPadding sets the padding for the tip of the tooltip
func (t *spectrumTooltip) TipPadding(tipPadding int) *spectrumTooltip {
	t.PropTipPadding = tipPadding
	return t
}

// Variant sets the visual variant of the tooltip
func (t *spectrumTooltip) Variant(variant TooltipVariant) *spectrumTooltip {
	t.PropVariant = variant
	return t
}

// Slot sets the slot attribute
func (t *spectrumTooltip) Slot(slot string) *spectrumTooltip {
	t.PropSlot = slot
	return t
}

// Content sets the text content of the tooltip
func (t *spectrumTooltip) Content(content string) *spectrumTooltip {
	t.PropContent = content
	return t
}

// InnerHTML sets the inner HTML of the tooltip
func (t *spectrumTooltip) InnerHTML(html string) *spectrumTooltip {
	t.PropInnerHTML = html
	return t
}

// Icon sets the icon UI element in the icon slot
func (t *spectrumTooltip) Icon(icon app.UI) *spectrumTooltip {
	t.PropIcon = icon
	return t
}

// Child sets a UI element as the child of the tooltip
func (t *spectrumTooltip) Child(child app.UI) *spectrumTooltip {
	t.PropChild = child
	return t
}

// InfoVariant is a convenience method to set the variant to info
func (t *spectrumTooltip) InfoVariant() *spectrumTooltip {
	t.PropVariant = TooltipVariantInfo
	return t
}

// PositiveVariant is a convenience method to set the variant to positive
func (t *spectrumTooltip) PositiveVariant() *spectrumTooltip {
	t.PropVariant = TooltipVariantPositive
	return t
}

// NegativeVariant is a convenience method to set the variant to negative
func (t *spectrumTooltip) NegativeVariant() *spectrumTooltip {
	t.PropVariant = TooltipVariantNegative
	return t
}

// Render renders the tooltip component
func (t *spectrumTooltip) Render() app.UI {
	tooltip := app.Elem("sp-tooltip")

	// Set attributes based on properties
	if t.PropOpen {
		tooltip = tooltip.Attr("open", true)
	}

	if t.PropPlacement != "" {
		tooltip = tooltip.Attr("placement", string(t.PropPlacement))
	}

	if t.PropOffset != 0 {
		tooltip = tooltip.Attr("offset", t.PropOffset)
	}

	if t.PropSelfManaged {
		tooltip = tooltip.Attr("self-managed", true)
	}

	if t.PropTipPadding != 0 {
		tooltip = tooltip.Attr("tip-padding", t.PropTipPadding)
	}

	if t.PropVariant != "" {
		tooltip = tooltip.Attr("variant", string(t.PropVariant))
	}

	if t.PropSlot != "" {
		tooltip = tooltip.Attr("slot", t.PropSlot)
	}

	// Handle content
	elements := []app.UI{}

	// Add icon if provided
	if t.PropIcon != nil {
		icon := t.PropIcon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		elements = append(elements, icon)
	}

	// Add content or child element
	if t.PropContent != "" {
		// Use text content if provided
		elements = append(elements, app.Text(t.PropContent))
	} else if t.PropInnerHTML != "" {
		// Use innerHTML if provided
		elements = append(elements, app.Raw(t.PropInnerHTML))
	} else if t.PropChild != nil {
		// Use child element if provided
		elements = append(elements, t.PropChild)
	}

	// Add all elements to the tooltip
	if len(elements) > 0 {
		tooltip = tooltip.Body(elements...)
	}

	return tooltip
}
