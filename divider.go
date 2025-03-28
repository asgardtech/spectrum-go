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

// SpectrumDivider represents an sp-divider component
type SpectrumDivider struct {
	app.Compo

	// Properties
	size     DividerSize
	vertical bool
}

// Divider creates a new divider component
func Divider() *SpectrumDivider {
	return &SpectrumDivider{
		size: DividerSizeM, // Default size is medium
	}
}

// Size sets the visual size of the divider
func (d *SpectrumDivider) Size(size DividerSize) *SpectrumDivider {
	d.size = size
	return d
}

// Vertical sets whether the divider is vertically oriented
func (d *SpectrumDivider) Vertical(vertical bool) *SpectrumDivider {
	d.vertical = vertical
	return d
}

// Render renders the divider component
func (d *SpectrumDivider) Render() app.UI {
	divider := app.Elem("sp-divider").
		Attr("size", string(d.size)).
		Attr("vertical", d.vertical)

	return divider
}
