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

// spectrumBadge represents an sp-badge component
type spectrumBadge struct {
	app.Compo

	// Properties
	PropVariant BadgeVariant
	PropSize    BadgeSize
	PropFixed   BadgeFixed

	// Content
	PropContent   string
	PropInnerHTML string
	PropIcon      app.UI
	PropChild     app.UI
}

// Badge creates a new badge component
func Badge() *spectrumBadge {
	return &spectrumBadge{
		PropVariant: BadgeVariantInformative, // Default variant is informative
		PropSize:    BadgeSizeM,              // Default size is medium
	}
}

// Variant sets the visual variant of the badge
func (b *spectrumBadge) Variant(variant BadgeVariant) *spectrumBadge {
	b.PropVariant = variant
	return b
}

// Size sets the visual size of the badge
func (b *spectrumBadge) Size(size BadgeSize) *spectrumBadge {
	b.PropSize = size
	return b
}

// Fixed sets the fixed position of the badge
func (b *spectrumBadge) Fixed(fixed BadgeFixed) *spectrumBadge {
	b.PropFixed = fixed
	return b
}

// Content sets the text content of the badge
func (b *spectrumBadge) Content(content string) *spectrumBadge {
	b.PropContent = content
	return b
}

// InnerHTML sets the inner HTML of the badge
func (b *spectrumBadge) InnerHTML(html string) *spectrumBadge {
	b.PropInnerHTML = html
	return b
}

// Icon sets the icon UI element in the icon slot
func (b *spectrumBadge) Icon(icon app.UI) *spectrumBadge {
	b.PropIcon = icon
	return b
}

// Child sets a UI element as the child of the badge
func (b *spectrumBadge) Child(child app.UI) *spectrumBadge {
	b.PropChild = child
	return b
}

// Render renders the badge component
func (b *spectrumBadge) Render() app.UI {
	badge := app.Elem("sp-badge")

	// Set attributes conditionally
	if b.PropVariant != BadgeVariantInformative {
		badge = badge.Attr("variant", string(b.PropVariant))
	}

	if b.PropSize != BadgeSizeM {
		badge = badge.Attr("size", string(b.PropSize))
	}

	if b.PropFixed != "" {
		badge = badge.Attr("fixed", string(b.PropFixed))
	}

	// Create an elements slice for body content
	elements := []app.UI{}

	// Add icon if provided
	if b.PropIcon != nil {
		iconElem := b.PropIcon
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
	if b.PropChild != nil {
		elements = append(elements, b.PropChild)
	} else if b.PropInnerHTML != "" {
		elements = append(elements, app.Raw(b.PropInnerHTML))
	} else if b.PropContent != "" {
		elements = append(elements, app.Text(b.PropContent))
	}

	// Add all elements to the badge
	if len(elements) > 0 {
		badge = badge.Body(elements...)
	}

	return badge
}
