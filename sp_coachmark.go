package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// CoachmarkAsset represents the The type of asset to represent
type CoachmarkAsset string

// CoachmarkAsset values
const (
    CoachmarkAssetFile CoachmarkAsset = "file"
    CoachmarkAssetFolder CoachmarkAsset = "folder"
)
// CoachmarkMediatype represents the Type of media to display
type CoachmarkMediatype string

// CoachmarkMediatype values
const (
    CoachmarkMediatypeImage CoachmarkMediatype = "image"
    CoachmarkMediatypeGif CoachmarkMediatype = "gif"
)
// CoachmarkPlacement represents the Position of the coachmark relative to its target
type CoachmarkPlacement string

// CoachmarkPlacement values
const (
    CoachmarkPlacementTop CoachmarkPlacement = "top"
    CoachmarkPlacementTopStart CoachmarkPlacement = "top-start"
    CoachmarkPlacementTopEnd CoachmarkPlacement = "top-end"
    CoachmarkPlacementRight CoachmarkPlacement = "right"
    CoachmarkPlacementRightStart CoachmarkPlacement = "right-start"
    CoachmarkPlacementRightEnd CoachmarkPlacement = "right-end"
    CoachmarkPlacementBottom CoachmarkPlacement = "bottom"
    CoachmarkPlacementBottomStart CoachmarkPlacement = "bottom-start"
    CoachmarkPlacementBottomEnd CoachmarkPlacement = "bottom-end"
    CoachmarkPlacementLeft CoachmarkPlacement = "left"
    CoachmarkPlacementLeftStart CoachmarkPlacement = "left-start"
    CoachmarkPlacementLeftEnd CoachmarkPlacement = "left-end"
)

// spectrumCoachmark represents an sp-coachmark component
type spectrumCoachmark struct {
    app.Compo
    *styler[*spectrumCoachmark]

    // Properties
    // The type of asset to represent
    PropAsset CoachmarkAsset
    // Current step in a multi-step tour
    PropCurrentstep float64
    // Whether the coachmark has an asset
    PropHasasset bool
    // Type of media to display
    PropMediatype CoachmarkMediatype
    // Array of modifier key labels (e.g., 'Shift', 'Alt', 'Cmd') that can be used with the shortcut key
    PropModifierkeys []string
    // Whether the coachmark is visible or not
    PropOpen bool
    // Position of the coachmark relative to its target
    PropPlacement CoachmarkPlacement
    // Text for the primary call-to-action button
    PropPrimarycta string
    // Text for the secondary call-to-action button
    PropSecondarycta string
    // The primary key used to trigger an interaction
    PropShortcutkey string
    // URL for the image or media to display
    PropSrc string
    // Whether to show the tip/arrow pointing to the target
    PropTip bool
    // Total number of steps in a multi-step tour
    PropTotalsteps float64

    // Text content for the default slot
    PropText string

    // Content slots
    PropActionsSlot app.UI
    PropAssetSlot app.UI
    PropContentSlot app.UI
    PropTitleSlot app.UI
    PropStepCountSlot app.UI


    // Event handlers
    PropOnPrimary app.EventHandler
    PropOnSecondary app.EventHandler
}

// Coachmark creates a new sp-coachmark component
func Coachmark() *spectrumCoachmark {
    element := &spectrumCoachmark{
        PropCurrentstep: 0,
        PropHasasset: false,
        PropModifierkeys: []string{},
        PropOpen: false,
        PropPlacement: CoachmarkPlacementRight,
        PropTip: false,
        PropTotalsteps: 0,
    }

    element.styler = newStyler(element)

    return element
}

// The type of asset to represent
func (c *spectrumCoachmark) Asset(asset CoachmarkAsset) *spectrumCoachmark {
    c.PropAsset = asset
    return c
}

func (c *spectrumCoachmark) AssetFile() *spectrumCoachmark {
    return c.Asset(CoachmarkAssetFile)
}
func (c *spectrumCoachmark) AssetFolder() *spectrumCoachmark {
    return c.Asset(CoachmarkAssetFolder)
}
// Current step in a multi-step tour
func (c *spectrumCoachmark) Currentstep(currentStep float64) *spectrumCoachmark {
    c.PropCurrentstep = currentStep
    return c
}

// Whether the coachmark has an asset
func (c *spectrumCoachmark) Hasasset(hasAsset bool) *spectrumCoachmark {
    c.PropHasasset = hasAsset
    return c
}

func (c *spectrumCoachmark) SetHasasset() *spectrumCoachmark {
    return c.Hasasset(true)
}

// Type of media to display
func (c *spectrumCoachmark) Mediatype(mediaType CoachmarkMediatype) *spectrumCoachmark {
    c.PropMediatype = mediaType
    return c
}

func (c *spectrumCoachmark) MediatypeImage() *spectrumCoachmark {
    return c.Mediatype(CoachmarkMediatypeImage)
}
func (c *spectrumCoachmark) MediatypeGif() *spectrumCoachmark {
    return c.Mediatype(CoachmarkMediatypeGif)
}
// Array of modifier key labels (e.g., 'Shift', 'Alt', 'Cmd') that can be used with the shortcut key
func (c *spectrumCoachmark) Modifierkeys(modifierKeys []string) *spectrumCoachmark {
    c.PropModifierkeys = modifierKeys
    return c
}

// Whether the coachmark is visible or not
func (c *spectrumCoachmark) Open(open bool) *spectrumCoachmark {
    c.PropOpen = open
    return c
}

func (c *spectrumCoachmark) SetOpen() *spectrumCoachmark {
    return c.Open(true)
}

// Position of the coachmark relative to its target
func (c *spectrumCoachmark) Placement(placement CoachmarkPlacement) *spectrumCoachmark {
    c.PropPlacement = placement
    return c
}

func (c *spectrumCoachmark) PlacementTop() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementTop)
}
func (c *spectrumCoachmark) PlacementTopStart() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementTopStart)
}
func (c *spectrumCoachmark) PlacementTopEnd() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementTopEnd)
}
func (c *spectrumCoachmark) PlacementRight() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementRight)
}
func (c *spectrumCoachmark) PlacementRightStart() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementRightStart)
}
func (c *spectrumCoachmark) PlacementRightEnd() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementRightEnd)
}
func (c *spectrumCoachmark) PlacementBottom() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementBottom)
}
func (c *spectrumCoachmark) PlacementBottomStart() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementBottomStart)
}
func (c *spectrumCoachmark) PlacementBottomEnd() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementBottomEnd)
}
func (c *spectrumCoachmark) PlacementLeft() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementLeft)
}
func (c *spectrumCoachmark) PlacementLeftStart() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementLeftStart)
}
func (c *spectrumCoachmark) PlacementLeftEnd() *spectrumCoachmark {
    return c.Placement(CoachmarkPlacementLeftEnd)
}
// Text for the primary call-to-action button
func (c *spectrumCoachmark) Primarycta(primaryCta string) *spectrumCoachmark {
    c.PropPrimarycta = primaryCta
    return c
}

// Text for the secondary call-to-action button
func (c *spectrumCoachmark) Secondarycta(secondaryCta string) *spectrumCoachmark {
    c.PropSecondarycta = secondaryCta
    return c
}

// The primary key used to trigger an interaction
func (c *spectrumCoachmark) Shortcutkey(shortcutKey string) *spectrumCoachmark {
    c.PropShortcutkey = shortcutKey
    return c
}

// URL for the image or media to display
func (c *spectrumCoachmark) Src(src string) *spectrumCoachmark {
    c.PropSrc = src
    return c
}

// Whether to show the tip/arrow pointing to the target
func (c *spectrumCoachmark) Tip(tip bool) *spectrumCoachmark {
    c.PropTip = tip
    return c
}

func (c *spectrumCoachmark) SetTip() *spectrumCoachmark {
    return c.Tip(true)
}

// Total number of steps in a multi-step tour
func (c *spectrumCoachmark) Totalsteps(totalSteps float64) *spectrumCoachmark {
    c.PropTotalsteps = totalSteps
    return c
}


// Text sets the text content for the default slot
func (c *spectrumCoachmark) Text(text string) *spectrumCoachmark {
    c.PropText = text
    return c
}


// An sp-action-menu element outlining actions related to the tour (e.g., 'Skip tour', 'Restart tour')
func (c *spectrumCoachmark) Actions(content app.UI) *spectrumCoachmark {
    c.PropActionsSlot = content

    return c
}

// Custom media content to display in the coachmark
func (c *spectrumCoachmark) AssetContent(content app.UI) *spectrumCoachmark {
    c.PropAssetSlot = content

    return c
}

// The main content of the coachmark
func (c *spectrumCoachmark) Content(content app.UI) *spectrumCoachmark {
    c.PropContentSlot = content

    return c
}

// The title or heading of the coachmark
func (c *spectrumCoachmark) Title(content app.UI) *spectrumCoachmark {
    c.PropTitleSlot = content

    return c
}

// Override the default pagination display with custom content
func (c *spectrumCoachmark) StepCount(content app.UI) *spectrumCoachmark {
    c.PropStepCountSlot = content

    return c
}



// Announces that the primary button has been clicked
func (c *spectrumCoachmark) OnPrimary(handler app.EventHandler) *spectrumCoachmark {
    c.PropOnPrimary = handler

    return c
}

// Announces that the secondary button has been clicked
func (c *spectrumCoachmark) OnSecondary(handler app.EventHandler) *spectrumCoachmark {
    c.PropOnSecondary = handler

    return c
}


// Render renders the sp-coachmark component
func (c *spectrumCoachmark) Render() app.UI {
    element := app.Elem("sp-coachmark")

    // Set attributes
    if c.PropAsset != "" {
        element = element.Attr("asset", string(c.PropAsset))
    }
    if c.PropCurrentstep != 0 {
        element = element.Attr("currentStep", c.PropCurrentstep)
    }
    if c.PropHasasset {
        element = element.Attr("hasAsset", true)
    }
    if c.PropMediatype != "" {
        element = element.Attr("mediaType", string(c.PropMediatype))
    }
    if len(c.PropModifierkeys) > 0 {
        element = element.Attr("modifierKeys", c.PropModifierkeys)
    }
    if c.PropOpen {
        element = element.Attr("open", true)
    }
    if c.PropPlacement != "" {
        element = element.Attr("placement", string(c.PropPlacement))
    }
    if c.PropPrimarycta != "" {
        element = element.Attr("primaryCta", c.PropPrimarycta)
    }
    if c.PropSecondarycta != "" {
        element = element.Attr("secondaryCta", c.PropSecondarycta)
    }
    if c.PropShortcutkey != "" {
        element = element.Attr("shortcutKey", c.PropShortcutkey)
    }
    if c.PropSrc != "" {
        element = element.Attr("src", c.PropSrc)
    }
    if c.PropTip {
        element = element.Attr("tip", true)
    }
    if c.PropTotalsteps != 0 {
        element = element.Attr("totalSteps", c.PropTotalsteps)
    }

    // Add event handlers
    if c.PropOnPrimary != nil {
        element = element.On("primary", c.PropOnPrimary)
    }
    if c.PropOnSecondary != nil {
        element = element.On("secondary", c.PropOnSecondary)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

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
    // Add asset slot
    if c.PropAssetSlot != nil {
        slotElem := c.PropAssetSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("asset")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "asset").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add content slot
    if c.PropContentSlot != nil {
        slotElem := c.PropContentSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("content")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "content").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add title slot
    if c.PropTitleSlot != nil {
        slotElem := c.PropTitleSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("title")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "title").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add step-count slot
    if c.PropStepCountSlot != nil {
        slotElem := c.PropStepCountSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("step-count")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "step-count").
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