package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// RadioSize represents the visual size of a radio button
type RadioSize string

// Radio sizes
const (
	RadioSizeS  RadioSize = "s"
	RadioSizeM  RadioSize = "m"
	RadioSizeL  RadioSize = "l"
	RadioSizeXL RadioSize = "xl"
)

// SpectrumRadio represents an sp-radio component
type SpectrumRadio struct {
	app.Compo

	// Properties
	label      string
	value      string
	size       RadioSize
	checked    bool
	disabled   bool
	emphasized bool
	invalid    bool
	readonly   bool

	// Event handlers
	onChange app.EventHandler
}

// Radio creates a new radio component
func Radio() *SpectrumRadio {
	return &SpectrumRadio{
		size: RadioSizeM, // Default size is medium
	}
}

// Label sets the radio button label text content
func (r *SpectrumRadio) Label(label string) *SpectrumRadio {
	r.label = label
	return r
}

// Value sets the radio button value
func (r *SpectrumRadio) Value(value string) *SpectrumRadio {
	r.value = value
	return r
}

// Size sets the visual size of the radio button
func (r *SpectrumRadio) Size(size RadioSize) *SpectrumRadio {
	r.size = size
	return r
}

// Checked sets the checked state of the radio button
func (r *SpectrumRadio) Checked(checked bool) *SpectrumRadio {
	r.checked = checked
	return r
}

// Disabled sets the disabled state of the radio button
func (r *SpectrumRadio) Disabled(disabled bool) *SpectrumRadio {
	r.disabled = disabled
	return r
}

// Emphasized sets the emphasized style of the radio button
func (r *SpectrumRadio) Emphasized(emphasized bool) *SpectrumRadio {
	r.emphasized = emphasized
	return r
}

// Invalid sets the invalid state of the radio button
func (r *SpectrumRadio) Invalid(invalid bool) *SpectrumRadio {
	r.invalid = invalid
	return r
}

// Readonly sets the readonly state of the radio button
func (r *SpectrumRadio) Readonly(readonly bool) *SpectrumRadio {
	r.readonly = readonly
	return r
}

// OnChange sets the change event handler
func (r *SpectrumRadio) OnChange(handler app.EventHandler) *SpectrumRadio {
	r.onChange = handler
	return r
}

// Render renders the radio component
func (r *SpectrumRadio) Render() app.UI {
	radio := app.Elem("sp-radio").
		Attr("value", r.value).
		Attr("size", string(r.size)).
		Attr("checked", r.checked).
		Attr("disabled", r.disabled).
		Attr("emphasized", r.emphasized).
		Attr("invalid", r.invalid).
		Attr("readonly", r.readonly)

	// Add event handlers
	if r.onChange != nil {
		radio = radio.On("change", r.onChange)
	}

	// Add the label as text content
	if r.label != "" {
		radio = radio.Text(r.label)
	}

	return radio
}
