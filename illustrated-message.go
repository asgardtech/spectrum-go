package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumIllustratedMessage represents an sp-illustrated-message component
// that displays an illustration icon and a message, usually in an empty state
// or on an error page.
type spectrumIllustratedMessage struct {
	app.Compo

	// Properties
	PropHeading     string
	PropDescription string

	// Content slots
	PropIllustrationSVG app.UI
	PropHeadingUI       app.UI
	PropDescriptionUI   app.UI
}

// IllustratedMessage creates a new illustrated message component
func IllustratedMessage() *spectrumIllustratedMessage {
	return &spectrumIllustratedMessage{}
}

// Heading sets the heading text
func (i *spectrumIllustratedMessage) Heading(heading string) *spectrumIllustratedMessage {
	i.PropHeading = heading
	return i
}

// Description sets the description text
func (i *spectrumIllustratedMessage) Description(description string) *spectrumIllustratedMessage {
	i.PropDescription = description
	return i
}

// IllustrationSVG sets the SVG illustration
func (i *spectrumIllustratedMessage) IllustrationSVG(svg app.UI) *spectrumIllustratedMessage {
	i.PropIllustrationSVG = svg
	return i
}

// HeadingContent sets custom UI content for the heading slot
func (i *spectrumIllustratedMessage) HeadingContent(content app.UI) *spectrumIllustratedMessage {
	i.PropHeadingUI = content
	return i
}

// DescriptionContent sets custom UI content for the description slot
func (i *spectrumIllustratedMessage) DescriptionContent(content app.UI) *spectrumIllustratedMessage {
	i.PropDescriptionUI = content
	return i
}

// Render renders the illustrated message component
func (i *spectrumIllustratedMessage) Render() app.UI {
	// Create the base element
	element := app.Elem("sp-illustrated-message")

	// If heading is set via property, add it
	if i.PropHeading != "" {
		element = element.Attr("heading", i.PropHeading)
	}

	// If description is set via property, add it
	if i.PropDescription != "" {
		element = element.Attr("description", i.PropDescription)
	}

	// Add the content elements in the right order
	elements := []app.UI{}

	// Add illustration SVG to the default slot if provided
	if i.PropIllustrationSVG != nil {
		elements = append(elements, i.PropIllustrationSVG)
	}

	// Add heading content if provided
	if i.PropHeadingUI != nil {
		// Check if the element already has a slot attribute
		if slottable, ok := i.PropHeadingUI.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, slottable.Slot("heading"))
		} else {
			// Otherwise wrap it in a div with the slot attribute
			elements = append(elements, app.Div().Attr("slot", "heading").Body(i.PropHeadingUI))
		}
	}

	// Add description content if provided
	if i.PropDescriptionUI != nil {
		// Check if the element already has a slot attribute
		if slottable, ok := i.PropDescriptionUI.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, slottable.Slot("description"))
		} else {
			// Otherwise wrap it in a div with the slot attribute
			elements = append(elements, app.Div().Attr("slot", "description").Body(i.PropDescriptionUI))
		}
	}

	// Add all elements to the component
	if len(elements) > 0 {
		element = element.Body(elements...)
	}

	return element
}
