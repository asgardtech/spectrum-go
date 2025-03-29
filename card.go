package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// CardVariant represents the variant of a spectrumCard.
type CardVariant string

const (
	// CardStandard is the default card variant.
	CardStandard CardVariant = "standard"
	// CardGallery is a gallery card variant.
	CardGallery CardVariant = "gallery"
	// CardQuiet is a quiet card variant.
	CardQuiet CardVariant = "quiet"
)

// CardAsset represents the asset type of a spectrumCard.
type CardAsset string

const (
	// CardFile represents a file asset.
	CardFile CardAsset = "file"
	// CardFolder represents a folder asset.
	CardFolder CardAsset = "folder"
)

// spectrumCard represents an sp-card component.
type spectrumCard struct {
	app.Compo

	PropHeading     string
	PropSubheading  string
	PropVariant     CardVariant
	PropHorizontal  bool
	PropAsset       CardAsset
	PropHref        string
	PropTarget      string
	PropRel         string
	PropDownload    string
	PropToggles     bool
	PropFocused     bool
	PropLabel       string
	PropValue       string
	PropContent     app.UI
	PropCoverPhoto  app.UI
	PropPreview     app.UI
	PropDescription app.UI
	PropFooter      app.UI
	PropActions     app.UI
	PropChildren    []app.UI
}

// Card creates a new spectrumCard instance.
func Card() *spectrumCard {
	return &spectrumCard{
		PropVariant: CardStandard,
	}
}

// Heading sets the heading text for the card.
func (c *spectrumCard) Heading(heading string) *spectrumCard {
	c.PropHeading = heading
	return c
}

// Subheading sets the subheading text for the card.
func (c *spectrumCard) Subheading(subheading string) *spectrumCard {
	c.PropSubheading = subheading
	return c
}

// Variant sets the card variant.
func (c *spectrumCard) Variant(variant CardVariant) *spectrumCard {
	c.PropVariant = variant
	return c
}

// Horizontal sets the card to horizontal layout.
func (c *spectrumCard) Horizontal(horizontal bool) *spectrumCard {
	c.PropHorizontal = horizontal
	return c
}

// Asset sets the asset type of the card (file or folder).
func (c *spectrumCard) Asset(asset CardAsset) *spectrumCard {
	c.PropAsset = asset
	return c
}

// Href sets the URL that the card links to.
func (c *spectrumCard) Href(href string) *spectrumCard {
	c.PropHref = href
	return c
}

// Target sets where to display the linked URL.
func (c *spectrumCard) Target(target string) *spectrumCard {
	c.PropTarget = target
	return c
}

// Rel sets the relationship of the linked URL.
func (c *spectrumCard) Rel(rel string) *spectrumCard {
	c.PropRel = rel
	return c
}

// Download sets the download attribute for the card.
func (c *spectrumCard) Download(download string) *spectrumCard {
	c.PropDownload = download
	return c
}

// Toggles sets whether the card can be toggled.
func (c *spectrumCard) Toggles(toggles bool) *spectrumCard {
	c.PropToggles = toggles
	return c
}

// Focused sets whether the card is focused.
func (c *spectrumCard) Focused(focused bool) *spectrumCard {
	c.PropFocused = focused
	return c
}

// Label sets an accessible label for the card.
func (c *spectrumCard) Label(label string) *spectrumCard {
	c.PropLabel = label
	return c
}

// Value sets the value of the card.
func (c *spectrumCard) Value(value string) *spectrumCard {
	c.PropValue = value
	return c
}

// Content sets the main content of the card.
func (c *spectrumCard) Content(content app.UI) *spectrumCard {
	c.PropContent = content
	return c
}

// CoverPhoto sets the cover photo for the card.
func (c *spectrumCard) CoverPhoto(img app.UI) *spectrumCard {
	c.PropCoverPhoto = img
	return c
}

// Preview sets the preview image for the card.
func (c *spectrumCard) Preview(img app.UI) *spectrumCard {
	c.PropPreview = img
	return c
}

// Description sets the description for the card.
func (c *spectrumCard) Description(description app.UI) *spectrumCard {
	c.PropDescription = description
	return c
}

// Footer sets the footer for the card.
func (c *spectrumCard) Footer(footer app.UI) *spectrumCard {
	c.PropFooter = footer
	return c
}

// Actions sets the actions menu for the card.
func (c *spectrumCard) Actions(actions app.UI) *spectrumCard {
	c.PropActions = actions
	return c
}

// Child adds a child element to the card.
func (c *spectrumCard) Child(child app.UI) *spectrumCard {
	c.PropChildren = append(c.PropChildren, child)
	return c
}

// Children adds multiple child elements to the card.
func (c *spectrumCard) Children(children ...app.UI) *spectrumCard {
	c.PropChildren = append(c.PropChildren, children...)
	return c
}

// Render renders the component.
func (c *spectrumCard) Render() app.UI {
	cardElem := app.Elem("sp-card")

	// Set attributes
	if c.PropHeading != "" {
		cardElem.Attr("heading", c.PropHeading)
	}

	if c.PropSubheading != "" {
		cardElem.Attr("subheading", c.PropSubheading)
	}

	if c.PropVariant != CardStandard {
		cardElem.Attr("variant", string(c.PropVariant))
	}

	if c.PropHorizontal {
		cardElem.Attr("horizontal", "")
	}

	if c.PropAsset != "" {
		cardElem.Attr("asset", string(c.PropAsset))
	}

	if c.PropHref != "" {
		cardElem.Attr("href", c.PropHref)
	}

	if c.PropTarget != "" {
		cardElem.Attr("target", c.PropTarget)
	}

	if c.PropRel != "" {
		cardElem.Attr("rel", c.PropRel)
	}

	if c.PropDownload != "" {
		cardElem.Attr("download", c.PropDownload)
	}

	if c.PropToggles {
		cardElem.Attr("toggles", "")
	}

	if c.PropFocused {
		cardElem.Attr("focused", "")
	}

	if c.PropLabel != "" {
		cardElem.Attr("label", c.PropLabel)
	}

	if c.PropValue != "" {
		cardElem.Attr("value", c.PropValue)
	}

	// Prepare children elements
	elements := []app.UI{}

	// Add slotted elements
	if c.PropCoverPhoto != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "cover-photo").Body(c.PropCoverPhoto))
	}

	if c.PropPreview != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "preview").Body(c.PropPreview))
	}

	if c.PropDescription != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "description").Body(c.PropDescription))
	}

	if c.PropFooter != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "footer").Body(c.PropFooter))
	}

	if c.PropActions != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "actions").Body(c.PropActions))
	}

	if c.PropContent != nil {
		elements = append(elements, c.PropContent)
	}

	// Add custom children
	elements = append(elements, c.PropChildren...)

	// Add all elements to the card
	cardElem.Body(elements...)

	return cardElem
}
