package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TagSize represents the The size of the tag
type TagSize string

// TagSize values
const (
    TagSizeS TagSize = "s"
    TagSizeM TagSize = "m"
    TagSizeL TagSize = "l"
)

// spectrumTag represents an sp-tag component
type spectrumTag struct {
    app.Compo
    *styler[*spectrumTag]

    // Properties
    // Whether the tag can be removed by the user
    PropDeletable bool
    // Whether the tag is disabled and cannot be interacted with
    PropDisabled bool
    // Whether the tag is in an invalid state
    PropInvalid bool
    // Whether the tag is in a readonly state
    PropReadonly bool
    // The size of the tag
    PropSize TagSize

    // Text content for the default slot
    PropText string

    // Content slots
    PropAvatarSlot app.UI
    PropIconSlot app.UI


    // Event handlers
    PropOnDelete app.EventHandler
}

// Tag creates a new sp-tag component
func Tag() *spectrumTag {
    element := &spectrumTag{
        PropDeletable: false,
        PropDisabled: false,
        PropInvalid: false,
        PropReadonly: false,
        PropSize: TagSizeM,
    }

    element.styler = newStyler(element)

    return element
}

// Whether the tag can be removed by the user
func (c *spectrumTag) Deletable(deletable bool) *spectrumTag {
    c.PropDeletable = deletable
    return c
}

func (c *spectrumTag) SetDeletable() *spectrumTag {
    return c.Deletable(true)
}

// Whether the tag is disabled and cannot be interacted with
func (c *spectrumTag) Disabled(disabled bool) *spectrumTag {
    c.PropDisabled = disabled
    return c
}

func (c *spectrumTag) SetDisabled() *spectrumTag {
    return c.Disabled(true)
}

// Whether the tag is in an invalid state
func (c *spectrumTag) Invalid(invalid bool) *spectrumTag {
    c.PropInvalid = invalid
    return c
}

func (c *spectrumTag) SetInvalid() *spectrumTag {
    return c.Invalid(true)
}

// Whether the tag is in a readonly state
func (c *spectrumTag) Readonly(readonly bool) *spectrumTag {
    c.PropReadonly = readonly
    return c
}

func (c *spectrumTag) SetReadonly() *spectrumTag {
    return c.Readonly(true)
}

// The size of the tag
func (c *spectrumTag) Size(size TagSize) *spectrumTag {
    c.PropSize = size
    return c
}

func (c *spectrumTag) SizeS() *spectrumTag {
    return c.Size(TagSizeS)
}
func (c *spectrumTag) SizeM() *spectrumTag {
    return c.Size(TagSizeM)
}
func (c *spectrumTag) SizeL() *spectrumTag {
    return c.Size(TagSizeL)
}

// Text sets the text content for the default slot
func (c *spectrumTag) Text(text string) *spectrumTag {
    c.PropText = text
    return c
}


// An avatar element to display within the Tag
func (c *spectrumTag) Avatar(content app.UI) *spectrumTag {
    c.PropAvatarSlot = content

    return c
}

// An icon element to display within the Tag
func (c *spectrumTag) Icon(content app.UI) *spectrumTag {
    c.PropIconSlot = content

    return c
}



// Fired when the delete button is clicked on a deletable tag
func (c *spectrumTag) OnDelete(handler app.EventHandler) *spectrumTag {
    c.PropOnDelete = handler

    return c
}


// Render renders the sp-tag component
func (c *spectrumTag) Render() app.UI {
    element := app.Elem("sp-tag")

    // Set attributes
    if c.PropDeletable {
        element = element.Attr("deletable", true)
    }
    if c.PropDisabled {
        element = element.Attr("disabled", true)
    }
    if c.PropInvalid {
        element = element.Attr("invalid", true)
    }
    if c.PropReadonly {
        element = element.Attr("readonly", true)
    }
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }

    // Add event handlers
    if c.PropOnDelete != nil {
        element = element.On("delete", c.PropOnDelete)
    }

    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }

    // Add avatar slot
    if c.PropAvatarSlot != nil {
        slotElem := c.PropAvatarSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("avatar")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "avatar").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
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