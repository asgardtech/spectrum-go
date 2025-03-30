package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// CardAsset represents the Represents the card as a file or folder asset type
type CardAsset string

// CardAsset values
const (
    CardAssetFile CardAsset = "file"
    CardAssetFolder CardAsset = "folder"
)
// CardReferrerpolicy represents the How much of the referrer to send when following the link
type CardReferrerpolicy string

// CardReferrerpolicy values
const (
    CardReferrerpolicyNoReferrer CardReferrerpolicy = "no-referrer"
    CardReferrerpolicyNoReferrerWhenDowngrade CardReferrerpolicy = "no-referrer-when-downgrade"
    CardReferrerpolicyOrigin CardReferrerpolicy = "origin"
    CardReferrerpolicyOriginWhenCrossOrigin CardReferrerpolicy = "origin-when-cross-origin"
    CardReferrerpolicySameOrigin CardReferrerpolicy = "same-origin"
    CardReferrerpolicyStrictOrigin CardReferrerpolicy = "strict-origin"
    CardReferrerpolicyStrictOriginWhenCrossOrigin CardReferrerpolicy = "strict-origin-when-cross-origin"
    CardReferrerpolicyUnsafeUrl CardReferrerpolicy = "unsafe-url"
)
// CardTarget represents the Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe)
type CardTarget string

// CardTarget values
const (
    CardTarget_blank CardTarget = "_blank"
    CardTarget_parent CardTarget = "_parent"
    CardTarget_self CardTarget = "_self"
    CardTarget_top CardTarget = "_top"
)
// CardVariant represents the The visual variant of the card
type CardVariant string

// CardVariant values
const (
    CardVariantStandard CardVariant = "standard"
    CardVariantGallery CardVariant = "gallery"
    CardVariantQuiet CardVariant = "quiet"
)

// spectrumCard represents an sp-card component
type spectrumCard struct {
    app.Compo
    *styler[*spectrumCard]

    // Properties
    // Represents the card as a file or folder asset type
    PropAsset CardAsset
    // Causes the browser to treat the linked URL as a download
    PropDownload string
    // Whether the card has focus
    PropFocused bool
    // The card heading text
    PropHeading string
    // Whether the card should use a horizontal layout
    PropHorizontal bool
    // The URL that the card links to when clicked
    PropHref string
    // An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
    PropLabel string
    // How much of the referrer to send when following the link
    PropReferrerpolicy CardReferrerpolicy
    // The relationship of the linked URL as space-separated link types
    PropRel string
    // The card subheading text
    PropSubheading string
    // Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe)
    PropTarget CardTarget
    // Whether the card can be toggled between selected and unselected states
    PropToggles bool
    // The value associated with the card
    PropValue string
    // The visual variant of the card
    PropVariant CardVariant


    // Content slots
    PropActionsSlot app.UI
    PropCoverPhotoSlot app.UI
    PropDescriptionSlot app.UI
    PropFooterSlot app.UI
    PropHeadingSlot app.UI
    PropPreviewSlot app.UI
    PropSubheadingSlot app.UI


    // Event handlers
    PropOnChange app.EventHandler
    PropOnClick app.EventHandler
}

// Card creates a new sp-card component
func Card() *spectrumCard {
    element := &spectrumCard{
        PropFocused: false,
        PropHeading: "",
        PropHorizontal: false,
        PropSubheading: "",
        PropToggles: false,
        PropValue: "",
        PropVariant: CardVariantStandard,
    }

    element.styler = newStyler(element)

    return element
}

// Represents the card as a file or folder asset type
func (c *spectrumCard) Asset(asset CardAsset) *spectrumCard {
    c.PropAsset = asset
    return c
}

func (c *spectrumCard) AssetFile() *spectrumCard {
    return c.Asset(CardAssetFile)
}
func (c *spectrumCard) AssetFolder() *spectrumCard {
    return c.Asset(CardAssetFolder)
}
// Causes the browser to treat the linked URL as a download
func (c *spectrumCard) Download(download string) *spectrumCard {
    c.PropDownload = download
    return c
}

// Whether the card has focus
func (c *spectrumCard) Focused(focused bool) *spectrumCard {
    c.PropFocused = focused
    return c
}

func (c *spectrumCard) SetFocused() *spectrumCard {
    return c.Focused(true)
}

// The card heading text
func (c *spectrumCard) Heading(heading string) *spectrumCard {
    c.PropHeading = heading
    return c
}

// Whether the card should use a horizontal layout
func (c *spectrumCard) Horizontal(horizontal bool) *spectrumCard {
    c.PropHorizontal = horizontal
    return c
}

func (c *spectrumCard) SetHorizontal() *spectrumCard {
    return c.Horizontal(true)
}

// The URL that the card links to when clicked
func (c *spectrumCard) Href(href string) *spectrumCard {
    c.PropHref = href
    return c
}

// An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.
func (c *spectrumCard) Label(label string) *spectrumCard {
    c.PropLabel = label
    return c
}

// How much of the referrer to send when following the link
func (c *spectrumCard) Referrerpolicy(referrerpolicy CardReferrerpolicy) *spectrumCard {
    c.PropReferrerpolicy = referrerpolicy
    return c
}

func (c *spectrumCard) ReferrerpolicyNoReferrer() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyNoReferrer)
}
func (c *spectrumCard) ReferrerpolicyNoReferrerWhenDowngrade() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyNoReferrerWhenDowngrade)
}
func (c *spectrumCard) ReferrerpolicyOrigin() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyOrigin)
}
func (c *spectrumCard) ReferrerpolicyOriginWhenCrossOrigin() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyOriginWhenCrossOrigin)
}
func (c *spectrumCard) ReferrerpolicySameOrigin() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicySameOrigin)
}
func (c *spectrumCard) ReferrerpolicyStrictOrigin() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyStrictOrigin)
}
func (c *spectrumCard) ReferrerpolicyStrictOriginWhenCrossOrigin() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyStrictOriginWhenCrossOrigin)
}
func (c *spectrumCard) ReferrerpolicyUnsafeUrl() *spectrumCard {
    return c.Referrerpolicy(CardReferrerpolicyUnsafeUrl)
}
// The relationship of the linked URL as space-separated link types
func (c *spectrumCard) Rel(rel string) *spectrumCard {
    c.PropRel = rel
    return c
}

// The card subheading text
func (c *spectrumCard) Subheading(subheading string) *spectrumCard {
    c.PropSubheading = subheading
    return c
}

// Where to display the linked URL, as the name for a browsing context (a tab, window, or iframe)
func (c *spectrumCard) Target(target CardTarget) *spectrumCard {
    c.PropTarget = target
    return c
}

func (c *spectrumCard) Target_blank() *spectrumCard {
    return c.Target(CardTarget_blank)
}
func (c *spectrumCard) Target_parent() *spectrumCard {
    return c.Target(CardTarget_parent)
}
func (c *spectrumCard) Target_self() *spectrumCard {
    return c.Target(CardTarget_self)
}
func (c *spectrumCard) Target_top() *spectrumCard {
    return c.Target(CardTarget_top)
}
// Whether the card can be toggled between selected and unselected states
func (c *spectrumCard) Toggles(toggles bool) *spectrumCard {
    c.PropToggles = toggles
    return c
}

func (c *spectrumCard) SetToggles() *spectrumCard {
    return c.Toggles(true)
}

// The value associated with the card
func (c *spectrumCard) Value(value string) *spectrumCard {
    c.PropValue = value
    return c
}

// The visual variant of the card
func (c *spectrumCard) Variant(variant CardVariant) *spectrumCard {
    c.PropVariant = variant
    return c
}

func (c *spectrumCard) VariantStandard() *spectrumCard {
    return c.Variant(CardVariantStandard)
}
func (c *spectrumCard) VariantGallery() *spectrumCard {
    return c.Variant(CardVariantGallery)
}
func (c *spectrumCard) VariantQuiet() *spectrumCard {
    return c.Variant(CardVariantQuiet)
}


// An sp-action-menu element outlining actions to take on the represented object
func (c *spectrumCard) Actions(content app.UI) *spectrumCard {
    c.PropActionsSlot = content

    return c
}

// This is the cover photo for Default and Quiet Cards
func (c *spectrumCard) CoverPhoto(content app.UI) *spectrumCard {
    c.PropCoverPhotoSlot = content

    return c
}

// A description of the card
func (c *spectrumCard) Description(content app.UI) *spectrumCard {
    c.PropDescriptionSlot = content

    return c
}

// Footer text
func (c *spectrumCard) Footer(content app.UI) *spectrumCard {
    c.PropFooterSlot = content

    return c
}

// HTML content to be listed as the heading
func (c *spectrumCard) HeadingContent(content app.UI) *spectrumCard {
    c.PropHeadingSlot = content

    return c
}

// This is the preview image for Gallery Cards
func (c *spectrumCard) Preview(content app.UI) *spectrumCard {
    c.PropPreviewSlot = content

    return c
}

// HTML content to be listed as the subheading
func (c *spectrumCard) SubheadingContent(content app.UI) *spectrumCard {
    c.PropSubheadingSlot = content

    return c
}



// Announces a change in the selected property of a card
func (c *spectrumCard) OnChange(handler app.EventHandler) *spectrumCard {
    c.PropOnChange = handler

    return c
}

// Fired when the card is clicked
func (c *spectrumCard) OnClick(handler app.EventHandler) *spectrumCard {
    c.PropOnClick = handler

    return c
}


// Render renders the sp-card component
func (c *spectrumCard) Render() app.UI {
    element := app.Elem("sp-card")

    // Set attributes
    if c.PropAsset != "" {
        element = element.Attr("asset", string(c.PropAsset))
    }
    if c.PropDownload != "" {
        element = element.Attr("download", c.PropDownload)
    }
    if c.PropFocused {
        element = element.Attr("focused", true)
    }
    if c.PropHeading != "" {
        element = element.Attr("heading", c.PropHeading)
    }
    if c.PropHorizontal {
        element = element.Attr("horizontal", true)
    }
    if c.PropHref != "" {
        element = element.Attr("href", c.PropHref)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropReferrerpolicy != "" {
        element = element.Attr("referrerpolicy", string(c.PropReferrerpolicy))
    }
    if c.PropRel != "" {
        element = element.Attr("rel", c.PropRel)
    }
    if c.PropSubheading != "" {
        element = element.Attr("subheading", c.PropSubheading)
    }
    if c.PropTarget != "" {
        element = element.Attr("target", string(c.PropTarget))
    }
    if c.PropToggles {
        element = element.Attr("toggles", true)
    }
    if c.PropValue != "" {
        element = element.Attr("value", c.PropValue)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }
    if c.PropOnClick != nil {
        element = element.On("click", c.PropOnClick)
    }

    // Add slots and children
    slotElements := []app.UI{}


    // Add actions slot
    if c.PropActionsSlot != nil {
        slotElem := c.PropActionsSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("actions")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "actions").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add cover-photo slot
    if c.PropCoverPhotoSlot != nil {
        slotElem := c.PropCoverPhotoSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("cover-photo")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "cover-photo").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add description slot
    if c.PropDescriptionSlot != nil {
        slotElem := c.PropDescriptionSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("description")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "description").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add footer slot
    if c.PropFooterSlot != nil {
        slotElem := c.PropFooterSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("footer")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "footer").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add heading slot
    if c.PropHeadingSlot != nil {
        slotElem := c.PropHeadingSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("heading")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "heading").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add preview slot
    if c.PropPreviewSlot != nil {
        slotElem := c.PropPreviewSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("preview")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "preview").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add subheading slot
    if c.PropSubheadingSlot != nil {
        slotElem := c.PropSubheadingSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("subheading")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "subheading").
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