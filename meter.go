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

// SpectrumMeter represents an sp-meter component
type SpectrumMeter struct {
	app.Compo

	// Properties
	size        MeterSize
	variant     MeterVariant
	progress    float64
	label       string
	sideLabel   bool
	staticColor MeterStaticColor

	// Content
	text      string
	innerHTML string
	child     app.UI
}

// Meter creates a new meter component
func Meter() *SpectrumMeter {
	return &SpectrumMeter{
		size:    MeterSizeM,              // Default size is medium
		variant: MeterVariantInformative, // Default variant is informative
	}
}

// Size sets the visual size of the meter
func (m *SpectrumMeter) Size(size MeterSize) *SpectrumMeter {
	m.size = size
	return m
}

// Variant sets the visual variant of the meter
func (m *SpectrumMeter) Variant(variant MeterVariant) *SpectrumMeter {
	m.variant = variant
	return m
}

// Progress sets the progress value (0-100)
func (m *SpectrumMeter) Progress(progress float64) *SpectrumMeter {
	m.progress = progress
	return m
}

// Label sets the accessibility label
func (m *SpectrumMeter) Label(label string) *SpectrumMeter {
	m.label = label
	return m
}

// SideLabel sets whether to display the label on the side
func (m *SpectrumMeter) SideLabel(sideLabel bool) *SpectrumMeter {
	m.sideLabel = sideLabel
	return m
}

// StaticColor sets the static color
func (m *SpectrumMeter) StaticColor(color MeterStaticColor) *SpectrumMeter {
	m.staticColor = color
	return m
}

// Text sets the text content
func (m *SpectrumMeter) Text(text string) *SpectrumMeter {
	m.text = text
	return m
}

// InnerHTML sets the inner HTML content
func (m *SpectrumMeter) InnerHTML(html string) *SpectrumMeter {
	m.innerHTML = html
	return m
}

// Child sets the child UI element
func (m *SpectrumMeter) Child(child app.UI) *SpectrumMeter {
	m.child = child
	return m
}

// Convenience methods for setting variants

// Positive sets the variant to positive
func (m *SpectrumMeter) Positive() *SpectrumMeter {
	return m.Variant(MeterVariantPositive)
}

// Notice sets the variant to notice
func (m *SpectrumMeter) Notice() *SpectrumMeter {
	return m.Variant(MeterVariantNotice)
}

// Negative sets the variant to negative
func (m *SpectrumMeter) Negative() *SpectrumMeter {
	return m.Variant(MeterVariantNegative)
}

// Render renders the meter component
func (m *SpectrumMeter) Render() app.UI {
	meter := app.Elem("sp-meter").
		Attr("size", string(m.size)).
		Attr("progress", m.progress).
		Attr("side-label", m.sideLabel)

	// Only add variant if it's not the default informative variant
	if m.variant != MeterVariantInformative {
		meter = meter.Attr("variant", string(m.variant))
	}

	if m.label != "" {
		meter = meter.Attr("label", m.label)
	}

	if m.staticColor != "" {
		meter = meter.Attr("static-color", string(m.staticColor))
	}

	// Handle content
	elements := []app.UI{}

	if m.innerHTML != "" {
		elements = append(elements, app.Raw(m.innerHTML))
	} else if m.text != "" {
		elements = append(elements, app.Text(m.text))
	} else if m.child != nil {
		elements = append(elements, m.child)
	}

	// Add elements to the meter
	if len(elements) > 0 {
		meter = meter.Body(elements...)
	}

	return meter
}
