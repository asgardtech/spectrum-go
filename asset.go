package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AssetVariant represents the variant of a SpectrumAsset.
type AssetVariant string

const (
	// AssetFile represents a file asset.
	AssetFile AssetVariant = "file"
	// AssetFolder represents a folder asset.
	AssetFolder AssetVariant = "folder"
)

// SpectrumAsset represents an sp-asset component
type SpectrumAsset struct {
	app.Compo

	// Properties
	label   string
	variant AssetVariant

	// Content
	children []app.UI
}

// Asset creates a new asset component
func Asset() *SpectrumAsset {
	return &SpectrumAsset{}
}

// Label sets the label for the asset
func (a *SpectrumAsset) Label(label string) *SpectrumAsset {
	a.label = label
	return a
}

// Variant sets the asset variant (file or folder)
func (a *SpectrumAsset) Variant(variant AssetVariant) *SpectrumAsset {
	a.variant = variant
	return a
}

// File sets the variant to "file"
func (a *SpectrumAsset) File() *SpectrumAsset {
	return a.Variant(AssetFile)
}

// Folder sets the variant to "folder"
func (a *SpectrumAsset) Folder() *SpectrumAsset {
	return a.Variant(AssetFolder)
}

// Children sets the child elements of the asset
func (a *SpectrumAsset) Children(children ...app.UI) *SpectrumAsset {
	a.children = children
	return a
}

// Child adds a child element to the asset
func (a *SpectrumAsset) Child(child app.UI) *SpectrumAsset {
	a.children = append(a.children, child)
	return a
}

// Render renders the asset component
func (a *SpectrumAsset) Render() app.UI {
	asset := app.Elem("sp-asset")

	// Set attributes
	if a.label != "" {
		asset = asset.Attr("label", a.label)
	}
	if a.variant != "" {
		asset = asset.Attr("variant", string(a.variant))
	}

	// Add children if provided
	if len(a.children) > 0 {
		asset = asset.Body(a.children...)
	}

	return asset
}
