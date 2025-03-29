package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AssetVariant represents the variant of a spectrumAsset.
type AssetVariant string

const (
	// AssetFile represents a file asset.
	AssetFile AssetVariant = "file"
	// AssetFolder represents a folder asset.
	AssetFolder AssetVariant = "folder"
)

// spectrumAsset represents an sp-asset component
type spectrumAsset struct {
	app.Compo

	// Properties
	PropLabel   string
	PropVariant AssetVariant

	// Content
	PropChildren []app.UI
}

// Asset creates a new asset component
func Asset() *spectrumAsset {
	return &spectrumAsset{}
}

// Label sets the label for the asset
func (a *spectrumAsset) Label(label string) *spectrumAsset {
	a.PropLabel = label
	return a
}

// Variant sets the asset variant (file or folder)
func (a *spectrumAsset) Variant(variant AssetVariant) *spectrumAsset {
	a.PropVariant = variant
	return a
}

// File sets the variant to "file"
func (a *spectrumAsset) File() *spectrumAsset {
	return a.Variant(AssetFile)
}

// Folder sets the variant to "folder"
func (a *spectrumAsset) Folder() *spectrumAsset {
	return a.Variant(AssetFolder)
}

// Children sets the child elements of the asset
func (a *spectrumAsset) Children(children ...app.UI) *spectrumAsset {
	a.PropChildren = children
	return a
}

// Child adds a child element to the asset
func (a *spectrumAsset) Child(child app.UI) *spectrumAsset {
	a.PropChildren = append(a.PropChildren, child)
	return a
}

// Render renders the asset component
func (a *spectrumAsset) Render() app.UI {
	asset := app.Elem("sp-asset")

	// Set attributes
	if a.PropLabel != "" {
		asset = asset.Attr("label", a.PropLabel)
	}
	if a.PropVariant != "" {
		asset = asset.Attr("variant", string(a.PropVariant))
	}

	// Add children if provided
	if len(a.PropChildren) > 0 {
		asset = asset.Body(a.PropChildren...)
	}

	return asset
}
