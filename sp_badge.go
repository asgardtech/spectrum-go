package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// BadgeFixed represents the Alters the border rounding to make the badge appear fixed to a specific edge
type BadgeFixed string

// BadgeFixed values
const (
    BadgeFixedBlockStart BadgeFixed = "block-start"
    BadgeFixedInlineEnd BadgeFixed = "inline-end"
    BadgeFixedBlockEnd BadgeFixed = "block-end"
    BadgeFixedInlineStart BadgeFixed = "inline-start"
)
// BadgeSize represents the The size of the badge
type BadgeSize string

// BadgeSize values
const (
    BadgeSizeS BadgeSize = "s"
    BadgeSizeM BadgeSize = "m"
    BadgeSizeL BadgeSize = "l"
    BadgeSizeXl BadgeSize = "xl"
)
// BadgeVariant represents the The visual variant of the badge
type BadgeVariant string

// BadgeVariant values
const (
    BadgeVariantAccent BadgeVariant = "accent"
    BadgeVariantInformative BadgeVariant = "informative"
    BadgeVariantNeutral BadgeVariant = "neutral"
    BadgeVariantPositive BadgeVariant = "positive"
    BadgeVariantNegative BadgeVariant = "negative"
    BadgeVariantNotice BadgeVariant = "notice"
    BadgeVariantSeafoam BadgeVariant = "seafoam"
    BadgeVariantIndigo BadgeVariant = "indigo"
    BadgeVariantPurple BadgeVariant = "purple"
    BadgeVariantFuchsia BadgeVariant = "fuchsia"
    BadgeVariantMagenta BadgeVariant = "magenta"
    BadgeVariantYellow BadgeVariant = "yellow"
    BadgeVariantGray BadgeVariant = "gray"
    BadgeVariantRed BadgeVariant = "red"
    BadgeVariantOrange BadgeVariant = "orange"
    BadgeVariantChartreuse BadgeVariant = "chartreuse"
    BadgeVariantCelery BadgeVariant = "celery"
    BadgeVariantGreen BadgeVariant = "green"
    BadgeVariantCyan BadgeVariant = "cyan"
    BadgeVariantBlue BadgeVariant = "blue"
)

// spectrumBadge represents an sp-badge component
type spectrumBadge struct {
    app.Compo
    *styler[*spectrumBadge]

    // Properties
    // Alters the border rounding to make the badge appear fixed to a specific edge
    PropFixed BadgeFixed
    // The size of the badge
    PropSize BadgeSize
    // The visual variant of the badge
    PropVariant BadgeVariant

    // Text content for the default slot
    PropText string

    // Content slots
    PropIconSlot app.UI


}

// Badge creates a new sp-badge component
func Badge() *spectrumBadge {
    element := &spectrumBadge{
        PropSize: BadgeSizeM,
        PropVariant: BadgeVariantInformative,
    }

    element.styler = newStyler(element)

    return element
}

// Alters the border rounding to make the badge appear fixed to a specific edge
func (c *spectrumBadge) Fixed(fixed BadgeFixed) *spectrumBadge {
    c.PropFixed = fixed
    return c
}

func (c *spectrumBadge) FixedBlockStart() *spectrumBadge {
    return c.Fixed(BadgeFixedBlockStart)
}
func (c *spectrumBadge) FixedInlineEnd() *spectrumBadge {
    return c.Fixed(BadgeFixedInlineEnd)
}
func (c *spectrumBadge) FixedBlockEnd() *spectrumBadge {
    return c.Fixed(BadgeFixedBlockEnd)
}
func (c *spectrumBadge) FixedInlineStart() *spectrumBadge {
    return c.Fixed(BadgeFixedInlineStart)
}
// The size of the badge
func (c *spectrumBadge) Size(size BadgeSize) *spectrumBadge {
    c.PropSize = size
    return c
}

func (c *spectrumBadge) SizeS() *spectrumBadge {
    return c.Size(BadgeSizeS)
}
func (c *spectrumBadge) SizeM() *spectrumBadge {
    return c.Size(BadgeSizeM)
}
func (c *spectrumBadge) SizeL() *spectrumBadge {
    return c.Size(BadgeSizeL)
}
func (c *spectrumBadge) SizeXl() *spectrumBadge {
    return c.Size(BadgeSizeXl)
}
// The visual variant of the badge
func (c *spectrumBadge) Variant(variant BadgeVariant) *spectrumBadge {
    c.PropVariant = variant
    return c
}

func (c *spectrumBadge) VariantAccent() *spectrumBadge {
    return c.Variant(BadgeVariantAccent)
}
func (c *spectrumBadge) VariantInformative() *spectrumBadge {
    return c.Variant(BadgeVariantInformative)
}
func (c *spectrumBadge) VariantNeutral() *spectrumBadge {
    return c.Variant(BadgeVariantNeutral)
}
func (c *spectrumBadge) VariantPositive() *spectrumBadge {
    return c.Variant(BadgeVariantPositive)
}
func (c *spectrumBadge) VariantNegative() *spectrumBadge {
    return c.Variant(BadgeVariantNegative)
}
func (c *spectrumBadge) VariantNotice() *spectrumBadge {
    return c.Variant(BadgeVariantNotice)
}
func (c *spectrumBadge) VariantSeafoam() *spectrumBadge {
    return c.Variant(BadgeVariantSeafoam)
}
func (c *spectrumBadge) VariantIndigo() *spectrumBadge {
    return c.Variant(BadgeVariantIndigo)
}
func (c *spectrumBadge) VariantPurple() *spectrumBadge {
    return c.Variant(BadgeVariantPurple)
}
func (c *spectrumBadge) VariantFuchsia() *spectrumBadge {
    return c.Variant(BadgeVariantFuchsia)
}
func (c *spectrumBadge) VariantMagenta() *spectrumBadge {
    return c.Variant(BadgeVariantMagenta)
}
func (c *spectrumBadge) VariantYellow() *spectrumBadge {
    return c.Variant(BadgeVariantYellow)
}
func (c *spectrumBadge) VariantGray() *spectrumBadge {
    return c.Variant(BadgeVariantGray)
}
func (c *spectrumBadge) VariantRed() *spectrumBadge {
    return c.Variant(BadgeVariantRed)
}
func (c *spectrumBadge) VariantOrange() *spectrumBadge {
    return c.Variant(BadgeVariantOrange)
}
func (c *spectrumBadge) VariantChartreuse() *spectrumBadge {
    return c.Variant(BadgeVariantChartreuse)
}
func (c *spectrumBadge) VariantCelery() *spectrumBadge {
    return c.Variant(BadgeVariantCelery)
}
func (c *spectrumBadge) VariantGreen() *spectrumBadge {
    return c.Variant(BadgeVariantGreen)
}
func (c *spectrumBadge) VariantCyan() *spectrumBadge {
    return c.Variant(BadgeVariantCyan)
}
func (c *spectrumBadge) VariantBlue() *spectrumBadge {
    return c.Variant(BadgeVariantBlue)
}

// Text sets the text content for the default slot
func (c *spectrumBadge) Text(text string) *spectrumBadge {
    c.PropText = text
    return c
}


// Optional icon that appears to the left of the label
func (c *spectrumBadge) Icon(content app.UI) *spectrumBadge {
    c.PropIconSlot = content

    return c
}




// Render renders the sp-badge component
func (c *spectrumBadge) Render() app.UI {
    element := app.Elem("sp-badge")

    // Set attributes
    if c.PropFixed != "" {
        element = element.Attr("fixed", string(c.PropFixed))
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add icon slot
    if c.PropIconSlot != nil {
        slotElem := c.PropIconSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("icon")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "icon").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 