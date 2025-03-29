package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// InfieldButtonInlinePosition represents the position of the infield button relative to the field
type InfieldButtonInlinePosition string

const (
	// InfieldButtonInlineStart positions the button at the start of the field
	InfieldButtonInlineStart InfieldButtonInlinePosition = "start"
	// InfieldButtonInlineEnd positions the button at the end of the field
	InfieldButtonInlineEnd InfieldButtonInlinePosition = "end"
)

// InfieldButtonBlockPosition represents the position of the infield button in a stacked layout
type InfieldButtonBlockPosition string

const (
	// InfieldButtonBlockStart positions the button at the top of the stack
	InfieldButtonBlockStart InfieldButtonBlockPosition = "start"
	// InfieldButtonBlockEnd positions the button at the bottom of the stack
	InfieldButtonBlockEnd InfieldButtonBlockPosition = "end"
)

// InfieldButtonSize represents the visual size of the infield button
type InfieldButtonSize string

const (
	// InfieldButtonSizeS represents the small size
	InfieldButtonSizeS InfieldButtonSize = "s"
	// InfieldButtonSizeM represents the medium size
	InfieldButtonSizeM InfieldButtonSize = "m"
	// InfieldButtonSizeL represents the large size
	InfieldButtonSizeL InfieldButtonSize = "l"
	// InfieldButtonSizeXL represents the extra-large size
	InfieldButtonSizeXL InfieldButtonSize = "xl"
)

// InfieldButtonType represents the type of button
type InfieldButtonType string

const (
	// InfieldButtonTypeButton represents a normal button
	InfieldButtonTypeButton InfieldButtonType = "button"
	// InfieldButtonTypeSubmit represents a submit button
	InfieldButtonTypeSubmit InfieldButtonType = "submit"
	// InfieldButtonTypeReset represents a reset button
	InfieldButtonTypeReset InfieldButtonType = "reset"
)

// SpectrumInfieldButton represents an sp-infield-button component
// that visually associates button functionality with form fields.
type spectrumInfieldButton struct {
	app.Compo

	// Properties
	PropActive         bool
	PropDisabled       bool
	PropDownload       string
	PropHref           string
	PropLabel          string
	PropReferrerPolicy string
	PropRel            string
	PropTabIndex       int
	PropTarget         string
	PropButtonType     InfieldButtonType
	PropInline         InfieldButtonInlinePosition
	PropBlock          InfieldButtonBlockPosition
	PropSize           InfieldButtonSize
	PropQuiet          bool

	// Events
	PropOnClick app.EventHandler

	// Content
	PropContent app.UI
	PropIcon    app.UI
}

// InfieldButton creates a new infield button component
func InfieldButton() *spectrumInfieldButton {
	return &spectrumInfieldButton{
		PropButtonType: InfieldButtonTypeButton,
		PropSize:       InfieldButtonSizeM,
	}
}

// Active sets the active state
func (i *spectrumInfieldButton) Active(active bool) *spectrumInfieldButton {
	i.PropActive = active
	return i
}

// Disabled sets the disabled state
func (i *spectrumInfieldButton) Disabled(disabled bool) *spectrumInfieldButton {
	i.PropDisabled = disabled
	return i
}

// Download sets the download attribute
func (i *spectrumInfieldButton) Download(download string) *spectrumInfieldButton {
	i.PropDownload = download
	return i
}

// Href sets the URL that the button points to
func (i *spectrumInfieldButton) Href(href string) *spectrumInfieldButton {
	i.PropHref = href
	return i
}

// Label sets the accessible label
func (i *spectrumInfieldButton) Label(label string) *spectrumInfieldButton {
	i.PropLabel = label
	return i
}

// ReferrerPolicy sets the referrer policy
func (i *spectrumInfieldButton) ReferrerPolicy(policy string) *spectrumInfieldButton {
	i.PropReferrerPolicy = policy
	return i
}

// Rel sets the relationship attribute
func (i *spectrumInfieldButton) Rel(rel string) *spectrumInfieldButton {
	i.PropRel = rel
	return i
}

// TabIndex sets the tab index
func (i *spectrumInfieldButton) TabIndex(index int) *spectrumInfieldButton {
	i.PropTabIndex = index
	return i
}

// Target sets the target for the link
func (i *spectrumInfieldButton) Target(target string) *spectrumInfieldButton {
	i.PropTarget = target
	return i
}

// Type sets the button type
func (i *spectrumInfieldButton) Type(buttonType InfieldButtonType) *spectrumInfieldButton {
	i.PropButtonType = buttonType
	return i
}

// Inline sets the inline position
func (i *spectrumInfieldButton) Inline(position InfieldButtonInlinePosition) *spectrumInfieldButton {
	i.PropInline = position
	return i
}

// Block sets the block position for stacked buttons
func (i *spectrumInfieldButton) Block(position InfieldButtonBlockPosition) *spectrumInfieldButton {
	i.PropBlock = position
	return i
}

// Size sets the visual size
func (i *spectrumInfieldButton) Size(size InfieldButtonSize) *spectrumInfieldButton {
	i.PropSize = size
	return i
}

// Quiet sets the quiet visual style
func (i *spectrumInfieldButton) Quiet(quiet bool) *spectrumInfieldButton {
	i.PropQuiet = quiet
	return i
}

// OnClick sets the click event handler
func (i *spectrumInfieldButton) OnClick(handler app.EventHandler) *spectrumInfieldButton {
	i.PropOnClick = handler
	return i
}

// Content sets the button text content
func (i *spectrumInfieldButton) Content(content app.UI) *spectrumInfieldButton {
	i.PropContent = content
	return i
}

// Icon sets the icon element
func (i *spectrumInfieldButton) Icon(icon app.UI) *spectrumInfieldButton {
	i.PropIcon = icon
	return i
}

// Render renders the infield button component
func (i *spectrumInfieldButton) Render() app.UI {
	button := app.Elem("sp-infield-button")

	// Set attributes
	if i.PropActive {
		button = button.Attr("active", "")
	}

	if i.PropDisabled {
		button = button.Attr("disabled", "")
	}

	if i.PropDownload != "" {
		button = button.Attr("download", i.PropDownload)
	}

	if i.PropHref != "" {
		button = button.Attr("href", i.PropHref)
	}

	if i.PropLabel != "" {
		button = button.Attr("label", i.PropLabel)
	}

	if i.PropReferrerPolicy != "" {
		button = button.Attr("referrerpolicy", i.PropReferrerPolicy)
	}

	if i.PropRel != "" {
		button = button.Attr("rel", i.PropRel)
	}

	if i.PropTabIndex != 0 {
		button = button.Attr("tabindex", i.PropTabIndex)
	}

	if i.PropTarget != "" {
		button = button.Attr("target", i.PropTarget)
	}

	if i.PropButtonType != InfieldButtonTypeButton {
		button = button.Attr("type", string(i.PropButtonType))
	}

	if i.PropInline != "" {
		button = button.Attr("inline", string(i.PropInline))
	}

	if i.PropBlock != "" {
		button = button.Attr("block", string(i.PropBlock))
	}

	if i.PropSize != InfieldButtonSizeM {
		button = button.Attr("size", string(i.PropSize))
	}

	if i.PropQuiet {
		button = button.Attr("quiet", "")
	}

	// Add event handlers
	if i.PropOnClick != nil {
		button = button.OnClick(i.PropOnClick)
	}

	// Add content
	elements := []app.UI{}

	if i.PropIcon != nil {
		if iconWithSlot, ok := i.PropIcon.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, iconWithSlot.Slot("icon"))
		} else {
			elements = append(elements, app.Div().Attr("slot", "icon").Body(i.PropIcon))
		}
	}

	if i.PropContent != nil {
		elements = append(elements, i.PropContent)
	}

	if len(elements) > 0 {
		button = button.Body(elements...)
	}

	return button
}
