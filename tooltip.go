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

// SpectrumTooltip represents an sp-tooltip component
type SpectrumTooltip struct {
	app.Compo

	// Properties
	open        bool
	placement   TooltipPlacement
	offset      int
	selfManaged bool
	tipPadding  int
	variant     TooltipVariant
	slot        string
	content     string
	innerHTML   string

	// Child elements
	icon  app.UI
	child app.UI
}

// Tooltip creates a new tooltip component
func Tooltip() *SpectrumTooltip {
	return &SpectrumTooltip{}
}

// Open sets whether the tooltip is visible
func (t *SpectrumTooltip) Open(open bool) *SpectrumTooltip {
	t.open = open
	return t
}

// Placement sets the position of the tooltip relative to its target
func (t *SpectrumTooltip) Placement(placement TooltipPlacement) *SpectrumTooltip {
	t.placement = placement
	return t
}

// Offset sets the distance from the tooltip to its target
func (t *SpectrumTooltip) Offset(offset int) *SpectrumTooltip {
	t.offset = offset
	return t
}

// SelfManaged sets whether the tooltip auto-binds to its parent element
func (t *SpectrumTooltip) SelfManaged(selfManaged bool) *SpectrumTooltip {
	t.selfManaged = selfManaged
	return t
}

// TipPadding sets the padding for the tip of the tooltip
func (t *SpectrumTooltip) TipPadding(tipPadding int) *SpectrumTooltip {
	t.tipPadding = tipPadding
	return t
}

// Variant sets the visual variant of the tooltip
func (t *SpectrumTooltip) Variant(variant TooltipVariant) *SpectrumTooltip {
	t.variant = variant
	return t
}

// Slot sets the slot attribute
func (t *SpectrumTooltip) Slot(slot string) *SpectrumTooltip {
	t.slot = slot
	return t
}

// Content sets the text content of the tooltip
func (t *SpectrumTooltip) Content(content string) *SpectrumTooltip {
	t.content = content
	return t
}

// InnerHTML sets the inner HTML of the tooltip
func (t *SpectrumTooltip) InnerHTML(html string) *SpectrumTooltip {
	t.innerHTML = html
	return t
}

// Icon sets the icon UI element in the icon slot
func (t *SpectrumTooltip) Icon(icon app.UI) *SpectrumTooltip {
	t.icon = icon
	return t
}

// Child sets a UI element as the child of the tooltip
func (t *SpectrumTooltip) Child(child app.UI) *SpectrumTooltip {
	t.child = child
	return t
}

// InfoVariant is a convenience method to set the variant to info
func (t *SpectrumTooltip) InfoVariant() *SpectrumTooltip {
	t.variant = TooltipVariantInfo
	return t
}

// PositiveVariant is a convenience method to set the variant to positive
func (t *SpectrumTooltip) PositiveVariant() *SpectrumTooltip {
	t.variant = TooltipVariantPositive
	return t
}

// NegativeVariant is a convenience method to set the variant to negative
func (t *SpectrumTooltip) NegativeVariant() *SpectrumTooltip {
	t.variant = TooltipVariantNegative
	return t
}

// Render renders the tooltip component
func (t *SpectrumTooltip) Render() app.UI {
	tooltip := app.Elem("sp-tooltip").
		Attr("open", t.open).
		Attr("self-managed", t.selfManaged)

	if t.placement != "" {
		tooltip = tooltip.Attr("placement", string(t.placement))
	}

	if t.offset != 0 {
		tooltip = tooltip.Attr("offset", t.offset)
	}

	if t.tipPadding != 0 {
		tooltip = tooltip.Attr("tipPadding", t.tipPadding)
	}

	if t.variant != "" {
		tooltip = tooltip.Attr("variant", string(t.variant))
	}

	if t.slot != "" {
		tooltip = tooltip.Attr("slot", t.slot)
	}

	// Add elements
	elements := []app.UI{}

	// Add icon if provided
	if t.icon != nil {
		iconElem := t.icon
		if iconWithSlot, ok := iconElem.(interface{ Slot(string) app.UI }); ok {
			iconElem = iconWithSlot.Slot("icon")
		} else {
			iconElem = app.Elem("div").
				Attr("slot", "icon").
				Body(iconElem)
		}
		elements = append(elements, iconElem)
	}

	// Handle content in the right order of precedence
	if t.child != nil {
		elements = append(elements, t.child)
	} else if t.innerHTML != "" {
		elements = append(elements, app.Raw(t.innerHTML))
	} else if t.content != "" {
		elements = append(elements, app.Text(t.content))
	}

	// Add all elements to the tooltip
	if len(elements) > 0 {
		tooltip = tooltip.Body(elements...)
	}

	return tooltip
}
