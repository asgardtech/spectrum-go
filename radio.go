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

// spectrumRadio represents an sp-radio component
type spectrumRadio struct {
	app.Compo

	// Properties
	PropLabel      string
	PropValue      string
	PropSize       RadioSize
	PropChecked    bool
	PropDisabled   bool
	PropEmphasized bool
	PropInvalid    bool
	PropReadonly   bool

	// Event handlers
	PropOnChange app.EventHandler
}

// Radio creates a new radio component
func Radio() *spectrumRadio {
	return &spectrumRadio{
		PropSize: RadioSizeM, // Default size is medium
	}
}

// Label sets the radio button label text content
func (r *spectrumRadio) Label(label string) *spectrumRadio {
	r.PropLabel = label
	return r
}

// Value sets the radio button value
func (r *spectrumRadio) Value(value string) *spectrumRadio {
	r.PropValue = value
	return r
}

// Size sets the visual size of the radio button
func (r *spectrumRadio) Size(size RadioSize) *spectrumRadio {
	r.PropSize = size
	return r
}

// Checked sets the checked state of the radio button
func (r *spectrumRadio) Checked(checked bool) *spectrumRadio {
	r.PropChecked = checked
	return r
}

// Disabled sets the disabled state of the radio button
func (r *spectrumRadio) Disabled(disabled bool) *spectrumRadio {
	r.PropDisabled = disabled
	return r
}

// Emphasized sets the emphasized style of the radio button
func (r *spectrumRadio) Emphasized(emphasized bool) *spectrumRadio {
	r.PropEmphasized = emphasized
	return r
}

// Invalid sets the invalid state of the radio button
func (r *spectrumRadio) Invalid(invalid bool) *spectrumRadio {
	r.PropInvalid = invalid
	return r
}

// Readonly sets the readonly state of the radio button
func (r *spectrumRadio) Readonly(readonly bool) *spectrumRadio {
	r.PropReadonly = readonly
	return r
}

// OnChange sets the change event handler
func (r *spectrumRadio) OnChange(handler app.EventHandler) *spectrumRadio {
	r.PropOnChange = handler
	return r
}

// Render renders the radio component
func (r *spectrumRadio) Render() app.UI {
	radio := app.Elem("sp-radio").
		Attr("value", r.PropValue).
		Attr("size", string(r.PropSize)).
		Attr("checked", r.PropChecked).
		Attr("disabled", r.PropDisabled).
		Attr("emphasized", r.PropEmphasized).
		Attr("invalid", r.PropInvalid).
		Attr("readonly", r.PropReadonly)

	// Add event handlers
	if r.PropOnChange != nil {
		radio = radio.On("change", r.PropOnChange)
	}

	// Add the label as text content
	if r.PropLabel != "" {
		radio = radio.Text(r.PropLabel)
	}

	return radio
}
