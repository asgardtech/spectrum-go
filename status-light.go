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

// spectrumStatusLight represents an sp-status-light component
type spectrumStatusLight struct {
	app.Compo

	// Properties
	PropSize      StatusLightSize
	PropVariant   StatusLightVariant
	PropDisabled  bool
	PropContent   string
	PropInnerHTML string
	PropChild     app.UI
}

// StatusLight creates a new status light component
func StatusLight() *spectrumStatusLight {
	return &spectrumStatusLight{
		PropSize:    StatusLightSizeM,       // Default size is medium
		PropVariant: StatusLightVariantInfo, // Default variant is info
	}
}

// Size sets the visual size of the status light
func (s *spectrumStatusLight) Size(size StatusLightSize) *spectrumStatusLight {
	s.PropSize = size
	return s
}

// Variant sets the visual variant of the status light
func (s *spectrumStatusLight) Variant(variant StatusLightVariant) *spectrumStatusLight {
	s.PropVariant = variant
	return s
}

// Disabled sets the disabled state
func (s *spectrumStatusLight) Disabled(disabled bool) *spectrumStatusLight {
	s.PropDisabled = disabled
	return s
}

// Content sets the text content of the status light
func (s *spectrumStatusLight) Content(content string) *spectrumStatusLight {
	s.PropContent = content
	return s
}

// InnerHTML sets the inner HTML of the status light
func (s *spectrumStatusLight) InnerHTML(html string) *spectrumStatusLight {
	s.PropInnerHTML = html
	return s
}

// Child sets a UI element as the child of the status light
func (s *spectrumStatusLight) Child(child app.UI) *spectrumStatusLight {
	s.PropChild = child
	return s
}

// Positive is a convenience method to set the variant to positive
func (s *spectrumStatusLight) Positive() *spectrumStatusLight {
	s.PropVariant = StatusLightVariantPositive
	return s
}

// Negative is a convenience method to set the variant to negative
func (s *spectrumStatusLight) Negative() *spectrumStatusLight {
	s.PropVariant = StatusLightVariantNegative
	return s
}

// Notice is a convenience method to set the variant to notice
func (s *spectrumStatusLight) Notice() *spectrumStatusLight {
	s.PropVariant = StatusLightVariantNotice
	return s
}

// Info is a convenience method to set the variant to info
func (s *spectrumStatusLight) Info() *spectrumStatusLight {
	s.PropVariant = StatusLightVariantInfo
	return s
}

// Neutral is a convenience method to set the variant to neutral
func (s *spectrumStatusLight) Neutral() *spectrumStatusLight {
	s.PropVariant = StatusLightVariantNeutral
	return s
}

// Render renders the status light component
func (s *spectrumStatusLight) Render() app.UI {
	statusLight := app.Elem("sp-status-light").
		Attr("size", string(s.PropSize)).
		Attr("variant", string(s.PropVariant)).
		Attr("disabled", s.PropDisabled)

	// Handle content in the right order of precedence
	if s.PropChild != nil {
		statusLight = statusLight.Body(s.PropChild)
	} else if s.PropInnerHTML != "" {
		statusLight = statusLight.Body(app.Raw(s.PropInnerHTML))
	} else if s.PropContent != "" {
		statusLight = statusLight.Body(app.Text(s.PropContent))
	}

	return statusLight
}
