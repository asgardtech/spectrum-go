package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// StatusLightSize represents the visual size of a status light
type StatusLightSize string

// Status light sizes
const (
	StatusLightSizeS  StatusLightSize = "s"
	StatusLightSizeM  StatusLightSize = "m"
	StatusLightSizeL  StatusLightSize = "l"
	StatusLightSizeXL StatusLightSize = "xl"
)

// StatusLightVariant represents the visual variant of a status light
type StatusLightVariant string

// Status light variants
const (
	StatusLightVariantNegative   StatusLightVariant = "negative"
	StatusLightVariantNotice     StatusLightVariant = "notice"
	StatusLightVariantPositive   StatusLightVariant = "positive"
	StatusLightVariantInfo       StatusLightVariant = "info"
	StatusLightVariantNeutral    StatusLightVariant = "neutral"
	StatusLightVariantYellow     StatusLightVariant = "yellow"
	StatusLightVariantFuchsia    StatusLightVariant = "fuchsia"
	StatusLightVariantIndigo     StatusLightVariant = "indigo"
	StatusLightVariantSeafoam    StatusLightVariant = "seafoam"
	StatusLightVariantChartreuse StatusLightVariant = "chartreuse"
	StatusLightVariantMagenta    StatusLightVariant = "magenta"
	StatusLightVariantCelery     StatusLightVariant = "celery"
	StatusLightVariantPurple     StatusLightVariant = "purple"
)

// SpectrumStatusLight represents an sp-status-light component
type SpectrumStatusLight struct {
	app.Compo

	// Properties
	size      StatusLightSize
	variant   StatusLightVariant
	disabled  bool
	content   string
	innerHTML string
	child     app.UI
}

// StatusLight creates a new status light component
func StatusLight() *SpectrumStatusLight {
	return &SpectrumStatusLight{
		size:    StatusLightSizeM,       // Default size is medium
		variant: StatusLightVariantInfo, // Default variant is info
	}
}

// Size sets the visual size of the status light
func (s *SpectrumStatusLight) Size(size StatusLightSize) *SpectrumStatusLight {
	s.size = size
	return s
}

// Variant sets the visual variant of the status light
func (s *SpectrumStatusLight) Variant(variant StatusLightVariant) *SpectrumStatusLight {
	s.variant = variant
	return s
}

// Disabled sets the disabled state
func (s *SpectrumStatusLight) Disabled(disabled bool) *SpectrumStatusLight {
	s.disabled = disabled
	return s
}

// Content sets the text content of the status light
func (s *SpectrumStatusLight) Content(content string) *SpectrumStatusLight {
	s.content = content
	return s
}

// InnerHTML sets the inner HTML of the status light
func (s *SpectrumStatusLight) InnerHTML(html string) *SpectrumStatusLight {
	s.innerHTML = html
	return s
}

// Child sets a UI element as the child of the status light
func (s *SpectrumStatusLight) Child(child app.UI) *SpectrumStatusLight {
	s.child = child
	return s
}

// Positive is a convenience method to set the variant to positive
func (s *SpectrumStatusLight) Positive() *SpectrumStatusLight {
	s.variant = StatusLightVariantPositive
	return s
}

// Negative is a convenience method to set the variant to negative
func (s *SpectrumStatusLight) Negative() *SpectrumStatusLight {
	s.variant = StatusLightVariantNegative
	return s
}

// Notice is a convenience method to set the variant to notice
func (s *SpectrumStatusLight) Notice() *SpectrumStatusLight {
	s.variant = StatusLightVariantNotice
	return s
}

// Info is a convenience method to set the variant to info
func (s *SpectrumStatusLight) Info() *SpectrumStatusLight {
	s.variant = StatusLightVariantInfo
	return s
}

// Neutral is a convenience method to set the variant to neutral
func (s *SpectrumStatusLight) Neutral() *SpectrumStatusLight {
	s.variant = StatusLightVariantNeutral
	return s
}

// Render renders the status light component
func (s *SpectrumStatusLight) Render() app.UI {
	statusLight := app.Elem("sp-status-light").
		Attr("size", string(s.size)).
		Attr("variant", string(s.variant)).
		Attr("disabled", s.disabled)

	// Handle content in the right order of precedence
	if s.child != nil {
		statusLight = statusLight.Body(s.child)
	} else if s.innerHTML != "" {
		statusLight = statusLight.Body(app.Raw(s.innerHTML))
	} else if s.content != "" {
		statusLight = statusLight.Body(app.Text(s.content))
	}

	return statusLight
}
