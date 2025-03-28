package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// BadgeSize represents the visual size of a badge
type BadgeSize string

// Badge sizes
const (
	BadgeSizeS  BadgeSize = "s"  // Small
	BadgeSizeM  BadgeSize = "m"  // Medium
	BadgeSizeL  BadgeSize = "l"  // Large
	BadgeSizeXL BadgeSize = "xl" // Extra Large
)

// BadgeVariant represents the visual variant of a badge
type BadgeVariant string

// Badge semantic variants
const (
	BadgeVariantAccent      BadgeVariant = "accent"
	BadgeVariantNeutral     BadgeVariant = "neutral"
	BadgeVariantInformative BadgeVariant = "informative"
	BadgeVariantPositive    BadgeVariant = "positive"
	BadgeVariantNegative    BadgeVariant = "negative"
	BadgeVariantNotice      BadgeVariant = "notice"
)

// Badge non-semantic variants (maintained for backwards compatibility)
const (
	BadgeVariantSeafoam    BadgeVariant = "seafoam"
	BadgeVariantIndigo     BadgeVariant = "indigo"
	BadgeVariantPurple     BadgeVariant = "purple"
	BadgeVariantFuchsia    BadgeVariant = "fuchsia"
	BadgeVariantMagenta    BadgeVariant = "magenta"
	BadgeVariantYellow     BadgeVariant = "yellow"
	BadgeVariantGray       BadgeVariant = "gray"
	BadgeVariantRed        BadgeVariant = "red"
	BadgeVariantOrange     BadgeVariant = "orange"
	BadgeVariantChartreuse BadgeVariant = "chartreuse"
	BadgeVariantCelery     BadgeVariant = "celery"
	BadgeVariantGreen      BadgeVariant = "green"
	BadgeVariantCyan       BadgeVariant = "cyan"
	BadgeVariantBlue       BadgeVariant = "blue"
)

// BadgeFixed represents the fixed position of a badge
type BadgeFixed string

// Badge fixed positions
const (
	BadgeFixedBlockStart  BadgeFixed = "block-start"
	BadgeFixedInlineEnd   BadgeFixed = "inline-end"
	BadgeFixedBlockEnd    BadgeFixed = "block-end"
	BadgeFixedInlineStart BadgeFixed = "inline-start"
)

// SpectrumBadge represents an sp-badge component
type SpectrumBadge struct {
	app.Compo

	// Properties
	variant BadgeVariant
	size    BadgeSize
	fixed   BadgeFixed

	// Content
	content   string
	innerHTML string
	icon      app.UI
	child     app.UI
}

// Badge creates a new badge component
func Badge() *SpectrumBadge {
	return &SpectrumBadge{
		variant: BadgeVariantInformative, // Default variant is informative
		size:    BadgeSizeM,              // Default size is medium
	}
}

// Variant sets the visual variant of the badge
func (b *SpectrumBadge) Variant(variant BadgeVariant) *SpectrumBadge {
	b.variant = variant
	return b
}

// Size sets the visual size of the badge
func (b *SpectrumBadge) Size(size BadgeSize) *SpectrumBadge {
	b.size = size
	return b
}

// Fixed sets the fixed position of the badge
func (b *SpectrumBadge) Fixed(fixed BadgeFixed) *SpectrumBadge {
	b.fixed = fixed
	return b
}

// Content sets the text content of the badge
func (b *SpectrumBadge) Content(content string) *SpectrumBadge {
	b.content = content
	return b
}

// InnerHTML sets the inner HTML of the badge
func (b *SpectrumBadge) InnerHTML(html string) *SpectrumBadge {
	b.innerHTML = html
	return b
}

// Icon sets the icon UI element in the icon slot
func (b *SpectrumBadge) Icon(icon app.UI) *SpectrumBadge {
	b.icon = icon
	return b
}

// Child sets a UI element as the child of the badge
func (b *SpectrumBadge) Child(child app.UI) *SpectrumBadge {
	b.child = child
	return b
}

// Render renders the badge component
func (b *SpectrumBadge) Render() app.UI {
	badge := app.Elem("sp-badge")

	// Set attributes conditionally
	if b.variant != BadgeVariantInformative {
		badge = badge.Attr("variant", string(b.variant))
	}

	if b.size != BadgeSizeM {
		badge = badge.Attr("size", string(b.size))
	}

	if b.fixed != "" {
		badge = badge.Attr("fixed", string(b.fixed))
	}

	// Create an elements slice for body content
	elements := []app.UI{}

	// Add icon if provided
	if b.icon != nil {
		iconElem := b.icon
		if iconWithSlot, ok := iconElem.(interface{ Slot(string) app.UI }); ok {
			iconElem = iconWithSlot.Slot("icon")
		} else {
			iconElem = app.Elem("div").
				Attr("slot", "icon").
				Body(iconElem)
		}
		elements = append(elements, iconElem)
	}

	// Add content or child based on precedence
	if b.child != nil {
		elements = append(elements, b.child)
	} else if b.innerHTML != "" {
		elements = append(elements, app.Raw(b.innerHTML))
	} else if b.content != "" {
		elements = append(elements, app.Text(b.content))
	}

	// Add all elements to the badge
	if len(elements) > 0 {
		badge = badge.Body(elements...)
	}

	return badge
}
