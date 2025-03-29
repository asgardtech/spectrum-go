package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DividerSize represents the visual size of a divider
type DividerSize string

// Divider sizes
const (
	DividerSizeS DividerSize = "s"
	DividerSizeM DividerSize = "m"
	DividerSizeL DividerSize = "l"
)

// spectrumDivider represents an sp-divider component
type spectrumDivider struct {
	app.Compo

	// Properties
	PropSize     DividerSize
	PropVertical bool
}

// Divider creates a new divider component
func Divider() *spectrumDivider {
	return &spectrumDivider{
		PropSize: DividerSizeM, // Default size is medium
	}
}

// Size sets the visual size of the divider
func (d *spectrumDivider) Size(size DividerSize) *spectrumDivider {
	d.PropSize = size
	return d
}

// Vertical sets whether the divider is vertically oriented
func (d *spectrumDivider) Vertical(vertical bool) *spectrumDivider {
	d.PropVertical = vertical
	return d
}

// Render renders the divider component
func (d *spectrumDivider) Render() app.UI {
	divider := app.Elem("sp-divider").
		Attr("size", string(d.PropSize)).
		Attr("vertical", d.PropVertical)

	return divider
}
