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

// SpectrumThumbnail represents a thumbnail component that can be used to display
// a preview of an image, layer, or effect.
type SpectrumThumbnail struct {
	app.Compo

	// Properties
	background string
	cover      bool
	disabled   bool
	focused    bool
	layer      bool
	selected   bool
	size       ThumbnailSize

	// Content
	children []app.UI
}

// Thumbnail creates a new thumbnail component.
func Thumbnail() *SpectrumThumbnail {
	return &SpectrumThumbnail{}
}

// Background sets the background color for letterboxing non-square content.
func (t *SpectrumThumbnail) Background(bg string) *SpectrumThumbnail {
	t.background = bg
	return t
}

// Cover sets whether the content should fill the space provided by the thumbnail.
func (t *SpectrumThumbnail) Cover(cover bool) *SpectrumThumbnail {
	t.cover = cover
	return t
}

// Disabled sets the disabled state of the thumbnail.
func (t *SpectrumThumbnail) Disabled(disabled bool) *SpectrumThumbnail {
	t.disabled = disabled
	return t
}

// Focused sets the focused state of the thumbnail.
func (t *SpectrumThumbnail) Focused(focused bool) *SpectrumThumbnail {
	t.focused = focused
	return t
}

// Layer sets whether the thumbnail is used in layer management.
func (t *SpectrumThumbnail) Layer(layer bool) *SpectrumThumbnail {
	t.layer = layer
	return t
}

// Selected sets whether the thumbnail is selected when used in layer management.
func (t *SpectrumThumbnail) Selected(selected bool) *SpectrumThumbnail {
	t.selected = selected
	return t
}

// Size sets the size of the thumbnail.
func (t *SpectrumThumbnail) Size(size ThumbnailSize) *SpectrumThumbnail {
	t.size = size
	return t
}

// Children sets the content of the thumbnail.
func (t *SpectrumThumbnail) Children(children ...app.UI) *SpectrumThumbnail {
	t.children = children
	return t
}

// Render renders the thumbnail component.
func (t *SpectrumThumbnail) Render() app.UI {
	thumbnail := app.Elem("sp-thumbnail")

	// Set attributes based on properties
	if t.background != "" {
		thumbnail.Attr("background", t.background)
	}

	if t.cover {
		thumbnail.Attr("cover", "")
	}

	if t.disabled {
		thumbnail.Attr("disabled", "")
	}

	if t.focused {
		thumbnail.Attr("focused", "")
	}

	if t.layer {
		thumbnail.Attr("layer", "")
	}

	if t.selected {
		thumbnail.Attr("selected", "")
	}

	if t.size != "" {
		thumbnail.Attr("size", string(t.size))
	}

	// Add children content
	for _, child := range t.children {
		thumbnail.Body(child)
	}

	return thumbnail
}
