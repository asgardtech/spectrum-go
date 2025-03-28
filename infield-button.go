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
type SpectrumInfieldButton struct {
	app.Compo

	// Properties
	active         bool
	disabled       bool
	download       string
	href           string
	label          string
	referrerpolicy string
	rel            string
	tabIndex       int
	target         string
	buttonType     InfieldButtonType
	inline         InfieldButtonInlinePosition
	block          InfieldButtonBlockPosition
	size           InfieldButtonSize
	quiet          bool

	// Events
	onClick app.EventHandler

	// Content
	content app.UI
	icon    app.UI
}

// InfieldButton creates a new infield button component
func InfieldButton() *SpectrumInfieldButton {
	return &SpectrumInfieldButton{
		buttonType: InfieldButtonTypeButton,
		size:       InfieldButtonSizeM,
	}
}

// Active sets the active state
func (i *SpectrumInfieldButton) Active(active bool) *SpectrumInfieldButton {
	i.active = active
	return i
}

// Disabled sets the disabled state
func (i *SpectrumInfieldButton) Disabled(disabled bool) *SpectrumInfieldButton {
	i.disabled = disabled
	return i
}

// Download sets the download attribute
func (i *SpectrumInfieldButton) Download(download string) *SpectrumInfieldButton {
	i.download = download
	return i
}

// Href sets the URL that the button points to
func (i *SpectrumInfieldButton) Href(href string) *SpectrumInfieldButton {
	i.href = href
	return i
}

// Label sets the accessible label
func (i *SpectrumInfieldButton) Label(label string) *SpectrumInfieldButton {
	i.label = label
	return i
}

// ReferrerPolicy sets the referrer policy
func (i *SpectrumInfieldButton) ReferrerPolicy(policy string) *SpectrumInfieldButton {
	i.referrerpolicy = policy
	return i
}

// Rel sets the relationship attribute
func (i *SpectrumInfieldButton) Rel(rel string) *SpectrumInfieldButton {
	i.rel = rel
	return i
}

// TabIndex sets the tab index
func (i *SpectrumInfieldButton) TabIndex(index int) *SpectrumInfieldButton {
	i.tabIndex = index
	return i
}

// Target sets the target for the link
func (i *SpectrumInfieldButton) Target(target string) *SpectrumInfieldButton {
	i.target = target
	return i
}

// Type sets the button type
func (i *SpectrumInfieldButton) Type(buttonType InfieldButtonType) *SpectrumInfieldButton {
	i.buttonType = buttonType
	return i
}

// Inline sets the inline position
func (i *SpectrumInfieldButton) Inline(position InfieldButtonInlinePosition) *SpectrumInfieldButton {
	i.inline = position
	return i
}

// Block sets the block position for stacked buttons
func (i *SpectrumInfieldButton) Block(position InfieldButtonBlockPosition) *SpectrumInfieldButton {
	i.block = position
	return i
}

// Size sets the visual size
func (i *SpectrumInfieldButton) Size(size InfieldButtonSize) *SpectrumInfieldButton {
	i.size = size
	return i
}

// Quiet sets the quiet visual style
func (i *SpectrumInfieldButton) Quiet(quiet bool) *SpectrumInfieldButton {
	i.quiet = quiet
	return i
}

// OnClick sets the click event handler
func (i *SpectrumInfieldButton) OnClick(handler app.EventHandler) *SpectrumInfieldButton {
	i.onClick = handler
	return i
}

// Content sets the button text content
func (i *SpectrumInfieldButton) Content(content app.UI) *SpectrumInfieldButton {
	i.content = content
	return i
}

// Icon sets the icon element
func (i *SpectrumInfieldButton) Icon(icon app.UI) *SpectrumInfieldButton {
	i.icon = icon
	return i
}

// Render renders the infield button component
func (i *SpectrumInfieldButton) Render() app.UI {
	button := app.Elem("sp-infield-button")

	// Set attributes
	if i.active {
		button = button.Attr("active", "")
	}

	if i.disabled {
		button = button.Attr("disabled", "")
	}

	if i.download != "" {
		button = button.Attr("download", i.download)
	}

	if i.href != "" {
		button = button.Attr("href", i.href)
	}

	if i.label != "" {
		button = button.Attr("label", i.label)
	}

	if i.referrerpolicy != "" {
		button = button.Attr("referrerpolicy", i.referrerpolicy)
	}

	if i.rel != "" {
		button = button.Attr("rel", i.rel)
	}

	if i.tabIndex != 0 {
		button = button.Attr("tabindex", i.tabIndex)
	}

	if i.target != "" {
		button = button.Attr("target", i.target)
	}

	if i.buttonType != InfieldButtonTypeButton {
		button = button.Attr("type", string(i.buttonType))
	}

	if i.inline != "" {
		button = button.Attr("inline", string(i.inline))
	}

	if i.block != "" {
		button = button.Attr("block", string(i.block))
	}

	if i.size != InfieldButtonSizeM {
		button = button.Attr("size", string(i.size))
	}

	if i.quiet {
		button = button.Attr("quiet", "")
	}

	// Add event handlers
	if i.onClick != nil {
		button = button.OnClick(i.onClick)
	}

	// Add content
	elements := []app.UI{}

	if i.icon != nil {
		if iconWithSlot, ok := i.icon.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, iconWithSlot.Slot("icon"))
		} else {
			elements = append(elements, app.Div().Attr("slot", "icon").Body(i.icon))
		}
	}

	if i.content != nil {
		elements = append(elements, i.content)
	}

	if len(elements) > 0 {
		button = button.Body(elements...)
	}

	return button
}
