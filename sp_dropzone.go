package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DropzoneDropeffect represents the Controls the feedback (typically visual) the user is given during a drag and drop operation
type DropzoneDropeffect string

// DropzoneDropeffect values
const (
    DropzoneDropeffectCopy DropzoneDropeffect = "copy"
    DropzoneDropeffectMove DropzoneDropeffect = "move"
    DropzoneDropeffectLink DropzoneDropeffect = "link"
    DropzoneDropeffectNone DropzoneDropeffect = "none"
)

// spectrumDropzone represents an sp-dropzone component
type spectrumDropzone struct {
    app.Compo
    *styler[*spectrumDropzone]

    // Properties
    // Indicates that files are currently being dragged over the dropzone.
    PropIsdragged bool
    // Controls the feedback (typically visual) the user is given during a drag and drop operation
    PropDropeffect DropzoneDropeffect
    // Set this property to indicate that the component is in a filled state.
    PropIsfilled bool

    // Text content for the default slot
    PropText string

    // Content slots


    // Event handlers
    PropOnSpDropzoneDragleave app.EventHandler
    PropOnSpDropzoneDragover app.EventHandler
    PropOnSpDropzoneDrop app.EventHandler
    PropOnSpDropzoneShouldAccept app.EventHandler
}

// Dropzone creates a new sp-dropzone component
func Dropzone() *spectrumDropzone {
    element := &spectrumDropzone{
        PropIsdragged: false,
        PropDropeffect: DropzoneDropeffectCopy,
        PropIsfilled: false,
    }

    element.styler = newStyler(element)

    return element
}

// Indicates that files are currently being dragged over the dropzone.
func (c *spectrumDropzone) Isdragged(isDragged bool) *spectrumDropzone {
    c.PropIsdragged = isDragged
    return c
}

func (c *spectrumDropzone) SetIsdragged() *spectrumDropzone {
    return c.Isdragged(true)
}

// Controls the feedback (typically visual) the user is given during a drag and drop operation
func (c *spectrumDropzone) Dropeffect(dropEffect DropzoneDropeffect) *spectrumDropzone {
    c.PropDropeffect = dropEffect
    return c
}

func (c *spectrumDropzone) DropeffectCopy() *spectrumDropzone {
    return c.Dropeffect(DropzoneDropeffectCopy)
}
func (c *spectrumDropzone) DropeffectMove() *spectrumDropzone {
    return c.Dropeffect(DropzoneDropeffectMove)
}
func (c *spectrumDropzone) DropeffectLink() *spectrumDropzone {
    return c.Dropeffect(DropzoneDropeffectLink)
}
func (c *spectrumDropzone) DropeffectNone() *spectrumDropzone {
    return c.Dropeffect(DropzoneDropeffectNone)
}
// Set this property to indicate that the component is in a filled state.
func (c *spectrumDropzone) Isfilled(isFilled bool) *spectrumDropzone {
    c.PropIsfilled = isFilled
    return c
}

func (c *spectrumDropzone) SetIsfilled() *spectrumDropzone {
    return c.Isfilled(true)
}


// Text sets the text content for the default slot
func (c *spectrumDropzone) Text(text string) *spectrumDropzone {
    c.PropText = text
    return c
}




// Announces when dragged files have been moved out of the UI without having been dropped.
func (c *spectrumDropzone) OnSpDropzoneDragleave(handler app.EventHandler) *spectrumDropzone {
    c.PropOnSpDropzoneDragleave = handler

    return c
}

// Announces when files have been dragged over the UI, but not yet dropped.
func (c *spectrumDropzone) OnSpDropzoneDragover(handler app.EventHandler) *spectrumDropzone {
    c.PropOnSpDropzoneDragover = handler

    return c
}

// Announces when dragged files have been dropped on the UI.
func (c *spectrumDropzone) OnSpDropzoneDrop(handler app.EventHandler) *spectrumDropzone {
    c.PropOnSpDropzoneDrop = handler

    return c
}

// A cancellable event that confirms whether or not a file dropped on the UI should be accepted.
func (c *spectrumDropzone) OnSpDropzoneShouldAccept(handler app.EventHandler) *spectrumDropzone {
    c.PropOnSpDropzoneShouldAccept = handler

    return c
}


// Render renders the sp-dropzone component
func (c *spectrumDropzone) Render() app.UI {
    element := app.Elem("sp-dropzone")

    // Set attributes
    if c.PropIsdragged {
        element = element.Attr("isDragged", true)
    }
    if c.PropDropeffect != "" {
        element = element.Attr("dropEffect", string(c.PropDropeffect))
    }
    if c.PropIsfilled {
        element = element.Attr("isFilled", true)
    }

    // Add event handlers
    if c.PropOnSpDropzoneDragleave != nil {
        element = element.On("sp-dropzone-dragleave", c.PropOnSpDropzoneDragleave)
    }
    if c.PropOnSpDropzoneDragover != nil {
        element = element.On("sp-dropzone-dragover", c.PropOnSpDropzoneDragover)
    }
    if c.PropOnSpDropzoneDrop != nil {
        element = element.On("sp-dropzone-drop", c.PropOnSpDropzoneDrop)
    }
    if c.PropOnSpDropzoneShouldAccept != nil {
        element = element.On("sp-dropzone-should-accept", c.PropOnSpDropzoneShouldAccept)
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