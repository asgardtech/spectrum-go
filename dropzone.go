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

// SpectrumDropzone represents an area into which an object can be dragged and dropped
// to accomplish a task.
type SpectrumDropzone struct {
	app.Compo

	dragged    bool
	filled     bool
	dropEffect DropEffect
	children   []app.UI

	// Event handlers
	onDragLeaveHandler    app.EventHandler
	onDragOverHandler     app.EventHandler
	onDropHandler         app.EventHandler
	onShouldAcceptHandler app.EventHandler
}

// Dropzone creates a new dropzone element.
func Dropzone() *SpectrumDropzone {
	return &SpectrumDropzone{}
}

// Dragged sets whether files are currently being dragged over the dropzone.
func (d *SpectrumDropzone) Dragged(dragged bool) *SpectrumDropzone {
	d.dragged = dragged
	return d
}

// Filled sets whether the dropzone is in a filled state.
func (d *SpectrumDropzone) Filled(filled bool) *SpectrumDropzone {
	d.filled = filled
	return d
}

// DropEffect sets the visual feedback given to the user during a drag and drop operation.
func (d *SpectrumDropzone) DropEffect(effect DropEffect) *SpectrumDropzone {
	d.dropEffect = effect
	return d
}

// OnDragLeave sets the handler for when dragged files have been moved out of the UI
// without having been dropped.
func (d *SpectrumDropzone) OnDragLeave(handler app.EventHandler) *SpectrumDropzone {
	d.onDragLeaveHandler = handler
	return d
}

// OnDragOver sets the handler for when files have been dragged over the UI,
// but not yet dropped.
func (d *SpectrumDropzone) OnDragOver(handler app.EventHandler) *SpectrumDropzone {
	d.onDragOverHandler = handler
	return d
}

// OnDrop sets the handler for when dragged files have been dropped on the UI.
func (d *SpectrumDropzone) OnDrop(handler app.EventHandler) *SpectrumDropzone {
	d.onDropHandler = handler
	return d
}

// OnShouldAccept sets the handler for confirming whether or not a file dropped
// on the UI should be accepted.
func (d *SpectrumDropzone) OnShouldAccept(handler app.EventHandler) *SpectrumDropzone {
	d.onShouldAcceptHandler = handler
	return d
}

// Children sets the child elements of the dropzone.
func (d *SpectrumDropzone) Children(children ...app.UI) *SpectrumDropzone {
	d.children = children
	return d
}

// Render renders the dropzone component.
func (d *SpectrumDropzone) Render() app.UI {
	element := app.Div().
		Class("spectrum-Dropzone").
		DataSet("spectrum-dropzone", "")

	if d.dragged {
		element.Attr("dragged", "")
	}

	if d.filled {
		element.Attr("filled", "")
	}

	if d.dropEffect != "" {
		element.Attr("drop-effect", string(d.dropEffect))
	}

	if d.onDragLeaveHandler != nil {
		element.OnDragLeave(d.onDragLeaveHandler)
	}

	if d.onDragOverHandler != nil {
		element.OnDragOver(d.onDragOverHandler)
	}

	if d.onDropHandler != nil {
		element.OnDrop(d.onDropHandler)
	}

	if d.onShouldAcceptHandler != nil {
		// For custom events, we need to handle them during the standard drag events
		// We'll use the dragenter event to check if we should accept the dragged items
		element.OnDragEnter(func(ctx app.Context, e app.Event) {
			// Call the shouldAccept handler with the original event
			d.onShouldAcceptHandler(ctx, e)
		})
	}

	element.Body(d.children...)

	return element
}
