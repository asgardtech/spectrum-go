package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionGroupSelects represents the Determines selection behavior: 'single' for radio-button behavior, 'multiple' for checkbox behavior
type ActionGroupSelects string

// ActionGroupSelects values
const (
    ActionGroupSelectsSingle ActionGroupSelects = "single"
    ActionGroupSelectsMultiple ActionGroupSelects = "multiple"
)
// ActionGroupSize represents the Size of the action buttons in the group
type ActionGroupSize string

// ActionGroupSize values
const (
    ActionGroupSizeXs ActionGroupSize = "xs"
    ActionGroupSizeS ActionGroupSize = "s"
    ActionGroupSizeM ActionGroupSize = "m"
    ActionGroupSizeL ActionGroupSize = "l"
    ActionGroupSizeXl ActionGroupSize = "xl"
)
// ActionGroupStaticcolor represents the Applies a specific color treatment to the buttons
type ActionGroupStaticcolor string

// ActionGroupStaticcolor values
const (
    ActionGroupStaticcolorWhite ActionGroupStaticcolor = "white"
    ActionGroupStaticcolorBlack ActionGroupStaticcolor = "black"
)

// spectrumActionGroup represents an sp-action-group component
type spectrumActionGroup struct {
    app.Compo
    *styler[*spectrumActionGroup]

    // Properties
    // Visually joins buttons together to clarify their relationship
    PropCompact bool
    // Applies visual emphasis to the buttons in the group
    PropEmphasized bool
    // Fills available horizontal space and evenly distributes that space across child buttons
    PropJustified bool
    // Accessible label for the action group
    PropLabel string
    // Applies a less visually prominent style to the buttons
    PropQuiet bool
    // Determines selection behavior: 'single' for radio-button behavior, 'multiple' for checkbox behavior
    PropSelects ActionGroupSelects
    // JSON string representation of selected button values array
    PropSelected string
    // Size of the action buttons in the group
    PropSize ActionGroupSize
    // Applies a specific color treatment to the buttons
    PropStaticcolor ActionGroupStaticcolor
    // Arranges buttons vertically instead of horizontally
    PropVertical bool

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnChange app.EventHandler
}

// ActionGroup creates a new sp-action-group component
func ActionGroup() *spectrumActionGroup {
    element := &spectrumActionGroup{
        PropCompact: false,
        PropEmphasized: false,
        PropJustified: false,
        PropLabel: "",
        PropQuiet: false,
        PropSelects: "",
        PropSelected: "[]",
        PropSize: ActionGroupSizeM,
        PropStaticcolor: "",
        PropVertical: false,
    }

    element.styler = newStyler(element)

    return element
}

// Visually joins buttons together to clarify their relationship
func (c *spectrumActionGroup) Compact(compact bool) *spectrumActionGroup {
    c.PropCompact = compact
    return c
}

func (c *spectrumActionGroup) SetCompact() *spectrumActionGroup {
    return c.Compact(true)
}

// Applies visual emphasis to the buttons in the group
func (c *spectrumActionGroup) Emphasized(emphasized bool) *spectrumActionGroup {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumActionGroup) SetEmphasized() *spectrumActionGroup {
    return c.Emphasized(true)
}

// Fills available horizontal space and evenly distributes that space across child buttons
func (c *spectrumActionGroup) Justified(justified bool) *spectrumActionGroup {
    c.PropJustified = justified
    return c
}

func (c *spectrumActionGroup) SetJustified() *spectrumActionGroup {
    return c.Justified(true)
}

// Accessible label for the action group
func (c *spectrumActionGroup) Label(label string) *spectrumActionGroup {
    c.PropLabel = label
    return c
}

// Applies a less visually prominent style to the buttons
func (c *spectrumActionGroup) Quiet(quiet bool) *spectrumActionGroup {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumActionGroup) SetQuiet() *spectrumActionGroup {
    return c.Quiet(true)
}

// Determines selection behavior: 'single' for radio-button behavior, 'multiple' for checkbox behavior
func (c *spectrumActionGroup) Selects(selects ActionGroupSelects) *spectrumActionGroup {
    c.PropSelects = selects
    return c
}

func (c *spectrumActionGroup) SelectsSingle() *spectrumActionGroup {
    return c.Selects(ActionGroupSelectsSingle)
}
func (c *spectrumActionGroup) SelectsMultiple() *spectrumActionGroup {
    return c.Selects(ActionGroupSelectsMultiple)
}
// JSON string representation of selected button values array
func (c *spectrumActionGroup) Selected(selected string) *spectrumActionGroup {
    c.PropSelected = selected
    return c
}

// Size of the action buttons in the group
func (c *spectrumActionGroup) Size(size ActionGroupSize) *spectrumActionGroup {
    c.PropSize = size
    return c
}

func (c *spectrumActionGroup) SizeXs() *spectrumActionGroup {
    return c.Size(ActionGroupSizeXs)
}
func (c *spectrumActionGroup) SizeS() *spectrumActionGroup {
    return c.Size(ActionGroupSizeS)
}
func (c *spectrumActionGroup) SizeM() *spectrumActionGroup {
    return c.Size(ActionGroupSizeM)
}
func (c *spectrumActionGroup) SizeL() *spectrumActionGroup {
    return c.Size(ActionGroupSizeL)
}
func (c *spectrumActionGroup) SizeXl() *spectrumActionGroup {
    return c.Size(ActionGroupSizeXl)
}
// Applies a specific color treatment to the buttons
func (c *spectrumActionGroup) Staticcolor(staticColor ActionGroupStaticcolor) *spectrumActionGroup {
    c.PropStaticcolor = staticColor
    return c
}

func (c *spectrumActionGroup) StaticcolorWhite() *spectrumActionGroup {
    return c.Staticcolor(ActionGroupStaticcolorWhite)
}
func (c *spectrumActionGroup) StaticcolorBlack() *spectrumActionGroup {
    return c.Staticcolor(ActionGroupStaticcolorBlack)
}
// Arranges buttons vertically instead of horizontally
func (c *spectrumActionGroup) Vertical(vertical bool) *spectrumActionGroup {
    c.PropVertical = vertical
    return c
}

func (c *spectrumActionGroup) SetVertical() *spectrumActionGroup {
    return c.Vertical(true)
}


// Text sets the text content for the default slot
func (c *spectrumActionGroup) Text(text string) *spectrumActionGroup {
    c.PropText = text
    return c
}




// Announces that selection state has been changed by user
func (c *spectrumActionGroup) OnChange(handler app.EventHandler) *spectrumActionGroup {
    c.PropOnChange = handler

    return c
}


// Render renders the sp-action-group component
func (c *spectrumActionGroup) Render() app.UI {
    element := app.Elem("sp-action-group")

    // Set attributes
    if c.PropCompact {
        element = element.Attr("compact", true)
    }
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
    }
    if c.PropJustified {
        element = element.Attr("justified", true)
    }
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropSelects != "" {
        element = element.Attr("selects", string(c.PropSelects))
    }
    if c.PropSelected != "" {
        element = element.Attr("selected", c.PropSelected)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropStaticcolor != "" {
        element = element.Attr("staticColor", string(c.PropStaticcolor))
    }
    if c.PropVertical {
        element = element.Attr("vertical", true)
    }

    // Add event handlers
    if c.PropOnChange != nil {
        element = element.On("change", c.PropOnChange)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 