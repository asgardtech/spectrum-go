package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// CardVariant represents the variant of a SpectrumCard.
type CardVariant string

const (
	// CardStandard is the default card variant.
	CardStandard CardVariant = "standard"
	// CardGallery is a gallery card variant.
	CardGallery CardVariant = "gallery"
	// CardQuiet is a quiet card variant.
	CardQuiet CardVariant = "quiet"
)

// CardAsset represents the asset type of a SpectrumCard.
type CardAsset string

const (
	// CardFile represents a file asset.
	CardFile CardAsset = "file"
	// CardFolder represents a folder asset.
	CardFolder CardAsset = "folder"
)

// SpectrumCard represents an sp-card component.
type SpectrumCard struct {
	app.Compo

	heading     string
	subheading  string
	variant     CardVariant
	horizontal  bool
	asset       CardAsset
	href        string
	target      string
	rel         string
	download    string
	toggles     bool
	focused     bool
	label       string
	value       string
	content     app.UI
	coverPhoto  app.UI
	preview     app.UI
	description app.UI
	footer      app.UI
	actions     app.UI
	children    []app.UI
}

// Card creates a new SpectrumCard instance.
func Card() *SpectrumCard {
	return &SpectrumCard{
		variant: CardStandard,
	}
}

// Heading sets the heading text for the card.
func (c *SpectrumCard) Heading(heading string) *SpectrumCard {
	c.heading = heading
	return c
}

// Subheading sets the subheading text for the card.
func (c *SpectrumCard) Subheading(subheading string) *SpectrumCard {
	c.subheading = subheading
	return c
}

// Variant sets the card variant.
func (c *SpectrumCard) Variant(variant CardVariant) *SpectrumCard {
	c.variant = variant
	return c
}

// Horizontal sets the card to horizontal layout.
func (c *SpectrumCard) Horizontal(horizontal bool) *SpectrumCard {
	c.horizontal = horizontal
	return c
}

// Asset sets the asset type of the card (file or folder).
func (c *SpectrumCard) Asset(asset CardAsset) *SpectrumCard {
	c.asset = asset
	return c
}

// Href sets the URL that the card links to.
func (c *SpectrumCard) Href(href string) *SpectrumCard {
	c.href = href
	return c
}

// Target sets where to display the linked URL.
func (c *SpectrumCard) Target(target string) *SpectrumCard {
	c.target = target
	return c
}

// Rel sets the relationship of the linked URL.
func (c *SpectrumCard) Rel(rel string) *SpectrumCard {
	c.rel = rel
	return c
}

// Download sets the download attribute for the card.
func (c *SpectrumCard) Download(download string) *SpectrumCard {
	c.download = download
	return c
}

// Toggles sets whether the card can be toggled.
func (c *SpectrumCard) Toggles(toggles bool) *SpectrumCard {
	c.toggles = toggles
	return c
}

// Focused sets whether the card is focused.
func (c *SpectrumCard) Focused(focused bool) *SpectrumCard {
	c.focused = focused
	return c
}

// Label sets an accessible label for the card.
func (c *SpectrumCard) Label(label string) *SpectrumCard {
	c.label = label
	return c
}

// Value sets the value of the card.
func (c *SpectrumCard) Value(value string) *SpectrumCard {
	c.value = value
	return c
}

// Content sets the main content of the card.
func (c *SpectrumCard) Content(content app.UI) *SpectrumCard {
	c.content = content
	return c
}

// CoverPhoto sets the cover photo for the card.
func (c *SpectrumCard) CoverPhoto(img app.UI) *SpectrumCard {
	c.coverPhoto = img
	return c
}

// Preview sets the preview image for the card.
func (c *SpectrumCard) Preview(img app.UI) *SpectrumCard {
	c.preview = img
	return c
}

// Description sets the description for the card.
func (c *SpectrumCard) Description(description app.UI) *SpectrumCard {
	c.description = description
	return c
}

// Footer sets the footer for the card.
func (c *SpectrumCard) Footer(footer app.UI) *SpectrumCard {
	c.footer = footer
	return c
}

// Actions sets the actions menu for the card.
func (c *SpectrumCard) Actions(actions app.UI) *SpectrumCard {
	c.actions = actions
	return c
}

// Child adds a child element to the card.
func (c *SpectrumCard) Child(child app.UI) *SpectrumCard {
	c.children = append(c.children, child)
	return c
}

// Children adds multiple child elements to the card.
func (c *SpectrumCard) Children(children ...app.UI) *SpectrumCard {
	c.children = append(c.children, children...)
	return c
}

// Render renders the component.
func (c *SpectrumCard) Render() app.UI {
	cardElem := app.Elem("sp-card")

	// Set attributes
	if c.heading != "" {
		cardElem.Attr("heading", c.heading)
	}

	if c.subheading != "" {
		cardElem.Attr("subheading", c.subheading)
	}

	if c.variant != CardStandard {
		cardElem.Attr("variant", string(c.variant))
	}

	if c.horizontal {
		cardElem.Attr("horizontal", "")
	}

	if c.asset != "" {
		cardElem.Attr("asset", string(c.asset))
	}

	if c.href != "" {
		cardElem.Attr("href", c.href)
	}

	if c.target != "" {
		cardElem.Attr("target", c.target)
	}

	if c.rel != "" {
		cardElem.Attr("rel", c.rel)
	}

	if c.download != "" {
		cardElem.Attr("download", c.download)
	}

	if c.toggles {
		cardElem.Attr("toggles", "")
	}

	if c.focused {
		cardElem.Attr("focused", "")
	}

	if c.label != "" {
		cardElem.Attr("label", c.label)
	}

	if c.value != "" {
		cardElem.Attr("value", c.value)
	}

	// Prepare children elements
	elements := []app.UI{}

	// Add slotted elements
	if c.coverPhoto != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "cover-photo").Body(c.coverPhoto))
	}

	if c.preview != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "preview").Body(c.preview))
	}

	if c.description != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "description").Body(c.description))
	}

	if c.footer != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "footer").Body(c.footer))
	}

	if c.actions != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "actions").Body(c.actions))
	}

	if c.content != nil {
		elements = append(elements, c.content)
	}

	// Add regular children
	elements = append(elements, c.children...)

	return cardElem.Body(elements...)
}
