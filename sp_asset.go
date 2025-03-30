package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AssetVariant represents the The type of asset to represent if not providing an image
type AssetVariant string

// AssetVariant values
const (
    AssetVariantFile AssetVariant = "file"
    AssetVariantFolder AssetVariant = "folder"
)

// spectrumAsset represents an sp-asset component
type spectrumAsset struct {
    app.Compo
    *styler[*spectrumAsset]

    // Properties
    // The label to display for the asset
    PropLabel string
    // The type of asset to represent if not providing an image
    PropVariant AssetVariant

    // Text content for the default slot
    PropText string

    // Content slots


}

// Asset creates a new sp-asset component
func Asset() *spectrumAsset {
    element := &spectrumAsset{
        PropLabel: "",
        PropVariant: "",
    }

    element.styler = newStyler(element)

    return element
}

// The label to display for the asset
func (c *spectrumAsset) Label(label string) *spectrumAsset {
    c.PropLabel = label
    return c
}

// The type of asset to represent if not providing an image
func (c *spectrumAsset) Variant(variant AssetVariant) *spectrumAsset {
    c.PropVariant = variant
    return c
}

func (c *spectrumAsset) VariantFile() *spectrumAsset {
    return c.Variant(AssetVariantFile)
}
func (c *spectrumAsset) VariantFolder() *spectrumAsset {
    return c.Variant(AssetVariantFolder)
}

// Text sets the text content for the default slot
func (c *spectrumAsset) Text(text string) *spectrumAsset {
    c.PropText = text
    return c
}





// Render renders the sp-asset component
func (c *spectrumAsset) Render() app.UI {
    element := app.Elem("sp-asset")

    // Set attributes
    if c.PropLabel != "" {
        element = element.Attr("label", c.PropLabel)
    }
    if c.PropVariant != "" {
        element = element.Attr("variant", string(c.PropVariant))
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