package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// MeterSize represents the visual size of a meter
type MeterSize string

// Meter sizes
const (
	MeterSizeS  MeterSize = "s"
	MeterSizeM  MeterSize = "m"
	MeterSizeL  MeterSize = "l"
	MeterSizeXL MeterSize = "xl"
)

// MeterVariant represents the visual variant of a meter
type MeterVariant string

// Meter variants
const (
	MeterVariantInformative MeterVariant = "" // Default/informative variant
	MeterVariantPositive    MeterVariant = "positive"
	MeterVariantNotice      MeterVariant = "notice"
	MeterVariantNegative    MeterVariant = "negative"
)

// MeterStaticColor represents static color options for meter
type MeterStaticColor string

// Meter static colors
const (
	MeterStaticColorWhite MeterStaticColor = "white"
)

// spectrumMeter represents an sp-meter component
type spectrumMeter struct {
	app.Compo

	// Properties
	PropSize        MeterSize
	PropVariant     MeterVariant
	PropProgress    float64
	PropLabel       string
	PropSideLabel   bool
	PropStaticColor MeterStaticColor

	// Content
	PropText      string
	PropInnerHTML string
	PropChild     app.UI
}

// Meter creates a new meter component
func Meter() *spectrumMeter {
	return &spectrumMeter{
		PropSize:    MeterSizeM,              // Default size is medium
		PropVariant: MeterVariantInformative, // Default variant is informative
	}
}

// Size sets the visual size of the meter
func (m *spectrumMeter) Size(size MeterSize) *spectrumMeter {
	m.PropSize = size
	return m
}

// Variant sets the visual variant of the meter
func (m *spectrumMeter) Variant(variant MeterVariant) *spectrumMeter {
	m.PropVariant = variant
	return m
}

// Progress sets the progress value (0-100)
func (m *spectrumMeter) Progress(progress float64) *spectrumMeter {
	m.PropProgress = progress
	return m
}

// Label sets the accessibility label
func (m *spectrumMeter) Label(label string) *spectrumMeter {
	m.PropLabel = label
	return m
}

// SideLabel sets whether to display the label on the side
func (m *spectrumMeter) SideLabel(sideLabel bool) *spectrumMeter {
	m.PropSideLabel = sideLabel
	return m
}

// StaticColor sets the static color
func (m *spectrumMeter) StaticColor(color MeterStaticColor) *spectrumMeter {
	m.PropStaticColor = color
	return m
}

// Text sets the text content
func (m *spectrumMeter) Text(text string) *spectrumMeter {
	m.PropText = text
	return m
}

// InnerHTML sets the inner HTML content
func (m *spectrumMeter) InnerHTML(html string) *spectrumMeter {
	m.PropInnerHTML = html
	return m
}

// Child sets the child UI element
func (m *spectrumMeter) Child(child app.UI) *spectrumMeter {
	m.PropChild = child
	return m
}

// Convenience methods for setting variants

// Positive sets the variant to positive
func (m *spectrumMeter) Positive() *spectrumMeter {
	return m.Variant(MeterVariantPositive)
}

// Notice sets the variant to notice
func (m *spectrumMeter) Notice() *spectrumMeter {
	return m.Variant(MeterVariantNotice)
}

// Negative sets the variant to negative
func (m *spectrumMeter) Negative() *spectrumMeter {
	return m.Variant(MeterVariantNegative)
}

// Render renders the meter component
func (m *spectrumMeter) Render() app.UI {
	meter := app.Elem("sp-meter").
		Attr("size", string(m.PropSize)).
		Attr("progress", m.PropProgress).
		Attr("side-label", m.PropSideLabel)

	// Only add variant if it's not the default informative variant
	if m.PropVariant != MeterVariantInformative {
		meter = meter.Attr("variant", string(m.PropVariant))
	}

	if m.PropLabel != "" {
		meter = meter.Attr("label", m.PropLabel)
	}

	if m.PropStaticColor != "" {
		meter = meter.Attr("static-color", string(m.PropStaticColor))
	}

	// Handle content
	elements := []app.UI{}

	if m.PropInnerHTML != "" {
		elements = append(elements, app.Raw(m.PropInnerHTML))
	} else if m.PropText != "" {
		elements = append(elements, app.Text(m.PropText))
	} else if m.PropChild != nil {
		elements = append(elements, m.PropChild)
	}

	// Add elements to the meter
	if len(elements) > 0 {
		meter = meter.Body(elements...)
	}

	return meter
}
