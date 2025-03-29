package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DropEffect describes the various visual effects that can be shown during drag and drop
type DropEffect string

const (
	// DropEffectCopy indicates a copy operation
	DropEffectCopy DropEffect = "copy"
	// DropEffectMove indicates a move operation
	DropEffectMove DropEffect = "move"
	// DropEffectLink indicates a link operation
	DropEffectLink DropEffect = "link"
	// DropEffectNone indicates no operation
	DropEffectNone DropEffect = "none"
)

// spectrumDropzone represents an area into which an object can be dragged and dropped
// to accomplish a task.
type spectrumDropzone struct {
	app.Compo

	PropDragged    bool
	PropFilled     bool
	PropDropEffect DropEffect
	PropChildren   []app.UI

	// Event handlers
	PropOnDragLeaveHandler    app.EventHandler
	PropOnDragOverHandler     app.EventHandler
	PropOnDropHandler         app.EventHandler
	PropOnShouldAcceptHandler app.EventHandler
}

// Dropzone creates a new dropzone element.
func Dropzone() *spectrumDropzone {
	return &spectrumDropzone{}
}

// Dragged sets whether files are currently being dragged over the dropzone.
func (d *spectrumDropzone) Dragged(dragged bool) *spectrumDropzone {
	d.PropDragged = dragged
	return d
}

// Filled sets whether the dropzone is in a filled state.
func (d *spectrumDropzone) Filled(filled bool) *spectrumDropzone {
	d.PropFilled = filled
	return d
}

// DropEffect sets the visual feedback given to the user during a drag and drop operation.
func (d *spectrumDropzone) DropEffect(effect DropEffect) *spectrumDropzone {
	d.PropDropEffect = effect
	return d
}

// OnDragLeave sets the handler for when dragged files have been moved out of the UI
// without having been dropped.
func (d *spectrumDropzone) OnDragLeave(handler app.EventHandler) *spectrumDropzone {
	d.PropOnDragLeaveHandler = handler
	return d
}

// OnDragOver sets the handler for when files have been dragged over the UI,
// but not yet dropped.
func (d *spectrumDropzone) OnDragOver(handler app.EventHandler) *spectrumDropzone {
	d.PropOnDragOverHandler = handler
	return d
}

// OnDrop sets the handler for when dragged files have been dropped on the UI.
func (d *spectrumDropzone) OnDrop(handler app.EventHandler) *spectrumDropzone {
	d.PropOnDropHandler = handler
	return d
}

// OnShouldAccept sets the handler for confirming whether or not a file dropped
// on the UI should be accepted.
func (d *spectrumDropzone) OnShouldAccept(handler app.EventHandler) *spectrumDropzone {
	d.PropOnShouldAcceptHandler = handler
	return d
}

// Children sets the child elements of the dropzone.
func (d *spectrumDropzone) Children(children ...app.UI) *spectrumDropzone {
	d.PropChildren = children
	return d
}

// Render renders the dropzone component.
func (d *spectrumDropzone) Render() app.UI {
	element := app.Div().
		Class("spectrum-Dropzone").
		DataSet("spectrum-dropzone", "")

	if d.PropDragged {
		element.Attr("dragged", "")
	}

	if d.PropFilled {
		element.Attr("filled", "")
	}

	if d.PropDropEffect != "" {
		element.Attr("drop-effect", string(d.PropDropEffect))
	}

	if d.PropOnDragLeaveHandler != nil {
		element.OnDragLeave(d.PropOnDragLeaveHandler)
	}

	if d.PropOnDragOverHandler != nil {
		element.OnDragOver(d.PropOnDragOverHandler)
	}

	if d.PropOnDropHandler != nil {
		element.OnDrop(d.PropOnDropHandler)
	}

	if d.PropOnShouldAcceptHandler != nil {
		// For custom events, we need to handle them during the standard drag events
		// We'll use the dragenter event to check if we should accept the dragged items
		element.OnDragEnter(func(ctx app.Context, e app.Event) {
			// Call the shouldAccept handler with the original event
			d.PropOnShouldAcceptHandler(ctx, e)
		})
	}

	element.Body(d.PropChildren...)

	return element
}
