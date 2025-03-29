package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ThumbnailSize represents the size variants available for the Thumbnail.
type ThumbnailSize string

// Thumbnail size constants.
const (
	Thumbnail50   ThumbnailSize = "50"
	Thumbnail75   ThumbnailSize = "75"
	Thumbnail100  ThumbnailSize = "100"
	Thumbnail200  ThumbnailSize = "200"
	Thumbnail300  ThumbnailSize = "300"
	Thumbnail400  ThumbnailSize = "400"
	Thumbnail500  ThumbnailSize = "500"
	Thumbnail600  ThumbnailSize = "600"
	Thumbnail700  ThumbnailSize = "700"
	Thumbnail800  ThumbnailSize = "800"
	Thumbnail900  ThumbnailSize = "900"
	Thumbnail1000 ThumbnailSize = "1000"
)

// spectrumThumbnail represents a thumbnail component that can be used to display
// a preview of an image, layer, or effect.
type spectrumThumbnail struct {
	app.Compo

	// Properties
	PropBackground string
	PropCover      bool
	PropDisabled   bool
	PropFocused    bool
	PropLayer      bool
	PropSelected   bool
	PropSize       ThumbnailSize

	// Content
	PropChildren []app.UI
}

// Thumbnail creates a new thumbnail component.
func Thumbnail() *spectrumThumbnail {
	return &spectrumThumbnail{}
}

// Background sets the background color for letterboxing non-square content.
func (t *spectrumThumbnail) Background(bg string) *spectrumThumbnail {
	t.PropBackground = bg
	return t
}

// Cover sets whether the content should fill the space provided by the thumbnail.
func (t *spectrumThumbnail) Cover(cover bool) *spectrumThumbnail {
	t.PropCover = cover
	return t
}

// Disabled sets the disabled state of the thumbnail.
func (t *spectrumThumbnail) Disabled(disabled bool) *spectrumThumbnail {
	t.PropDisabled = disabled
	return t
}

// Focused sets the focused state of the thumbnail.
func (t *spectrumThumbnail) Focused(focused bool) *spectrumThumbnail {
	t.PropFocused = focused
	return t
}

// Layer sets whether the thumbnail is used in layer management.
func (t *spectrumThumbnail) Layer(layer bool) *spectrumThumbnail {
	t.PropLayer = layer
	return t
}

// Selected sets whether the thumbnail is selected when used in layer management.
func (t *spectrumThumbnail) Selected(selected bool) *spectrumThumbnail {
	t.PropSelected = selected
	return t
}

// Size sets the size of the thumbnail.
func (t *spectrumThumbnail) Size(size ThumbnailSize) *spectrumThumbnail {
	t.PropSize = size
	return t
}

// Children sets the content of the thumbnail.
func (t *spectrumThumbnail) Children(children ...app.UI) *spectrumThumbnail {
	t.PropChildren = children
	return t
}

// Render renders the thumbnail component.
func (t *spectrumThumbnail) Render() app.UI {
	thumbnail := app.Elem("sp-thumbnail")

	// Set attributes based on properties
	if t.PropBackground != "" {
		thumbnail.Attr("background", t.PropBackground)
	}

	if t.PropCover {
		thumbnail.Attr("cover", "")
	}

	if t.PropDisabled {
		thumbnail.Attr("disabled", "")
	}

	if t.PropFocused {
		thumbnail.Attr("focused", "")
	}

	if t.PropLayer {
		thumbnail.Attr("layer", "")
	}

	if t.PropSelected {
		thumbnail.Attr("selected", "")
	}

	if t.PropSize != "" {
		thumbnail.Attr("size", string(t.PropSize))
	}

	// Add children content
	for _, child := range t.PropChildren {
		thumbnail.Body(child)
	}

	return thumbnail
}
