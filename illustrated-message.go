package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumIllustratedMessage represents an sp-illustrated-message component
// that displays an illustration icon and a message, usually in an empty state
// or on an error page.
type SpectrumIllustratedMessage struct {
	app.Compo

	// Properties
	heading     string
	description string

	// Content slots
	illustrationSVG app.UI
	headingUI       app.UI
	descriptionUI   app.UI
}

// IllustratedMessage creates a new illustrated message component
func IllustratedMessage() *SpectrumIllustratedMessage {
	return &SpectrumIllustratedMessage{}
}

// Heading sets the heading text
func (i *SpectrumIllustratedMessage) Heading(heading string) *SpectrumIllustratedMessage {
	i.heading = heading
	return i
}

// Description sets the description text
func (i *SpectrumIllustratedMessage) Description(description string) *SpectrumIllustratedMessage {
	i.description = description
	return i
}

// IllustrationSVG sets the SVG illustration
func (i *SpectrumIllustratedMessage) IllustrationSVG(svg app.UI) *SpectrumIllustratedMessage {
	i.illustrationSVG = svg
	return i
}

// HeadingContent sets custom UI content for the heading slot
func (i *SpectrumIllustratedMessage) HeadingContent(content app.UI) *SpectrumIllustratedMessage {
	i.headingUI = content
	return i
}

// DescriptionContent sets custom UI content for the description slot
func (i *SpectrumIllustratedMessage) DescriptionContent(content app.UI) *SpectrumIllustratedMessage {
	i.descriptionUI = content
	return i
}

// Render renders the illustrated message component
func (i *SpectrumIllustratedMessage) Render() app.UI {
	// Create the base element
	element := app.Elem("sp-illustrated-message")

	// If heading is set via property, add it
	if i.heading != "" {
		element = element.Attr("heading", i.heading)
	}

	// If description is set via property, add it
	if i.description != "" {
		element = element.Attr("description", i.description)
	}

	// Add the content elements in the right order
	elements := []app.UI{}

	// Add illustration SVG to the default slot if provided
	if i.illustrationSVG != nil {
		elements = append(elements, i.illustrationSVG)
	}

	// Add heading content if provided
	if i.headingUI != nil {
		// Check if the element already has a slot attribute
		if slottable, ok := i.headingUI.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, slottable.Slot("heading"))
		} else {
			// Otherwise wrap it in a div with the slot attribute
			elements = append(elements, app.Div().Attr("slot", "heading").Body(i.headingUI))
		}
	}

	// Add description content if provided
	if i.descriptionUI != nil {
		// Check if the element already has a slot attribute
		if slottable, ok := i.descriptionUI.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, slottable.Slot("description"))
		} else {
			// Otherwise wrap it in a div with the slot attribute
			elements = append(elements, app.Div().Attr("slot", "description").Body(i.descriptionUI))
		}
	}

	// Add all elements to the component
	if len(elements) > 0 {
		element = element.Body(elements...)
	}

	return element
}
