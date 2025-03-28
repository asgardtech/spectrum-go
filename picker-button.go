package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerButtonSize represents the visual size of a picker button
type PickerButtonSize string

// Picker button sizes
const (
	PickerButtonSizeS  PickerButtonSize = "s"
	PickerButtonSizeM  PickerButtonSize = "m"
	PickerButtonSizeL  PickerButtonSize = "l"
	PickerButtonSizeXL PickerButtonSize = "xl"
)

// PickerButtonPosition represents the position of a picker button
type PickerButtonPosition string

// Picker button positions
const (
	PickerButtonPositionLeft  PickerButtonPosition = "left"
	PickerButtonPositionRight PickerButtonPosition = "right"
)

// SpectrumPickerButton represents an sp-picker-button component
type SpectrumPickerButton struct {
	app.Compo

	// Properties
	active         bool
	disabled       bool
	download       string
	href           string
	label          string
	position       PickerButtonPosition
	referrerpolicy string
	rel            string
	rounded        bool
	size           PickerButtonSize
	tabIndex       int
	target         string
	type_          string
	open           bool
	invalid        bool
	quiet          bool

	// Slots
	labelSlot app.UI
	icon      app.UI

	// Event handlers
	onClick app.EventHandler
}

// PickerButton creates a new picker button component
func PickerButton() *SpectrumPickerButton {
	return &SpectrumPickerButton{
		position: PickerButtonPositionRight, // Default position
		size:     PickerButtonSizeM,         // Default size
		type_:    "button",                  // Default type
	}
}

// Active sets whether the button is active
func (pb *SpectrumPickerButton) Active(active bool) *SpectrumPickerButton {
	pb.active = active
	return pb
}

// Disabled sets whether the button is disabled
func (pb *SpectrumPickerButton) Disabled(disabled bool) *SpectrumPickerButton {
	pb.disabled = disabled
	return pb
}

// Download sets the download attribute
func (pb *SpectrumPickerButton) Download(download string) *SpectrumPickerButton {
	pb.download = download
	return pb
}

// Href sets the URL the button points to
func (pb *SpectrumPickerButton) Href(href string) *SpectrumPickerButton {
	pb.href = href
	return pb
}

// Label sets the accessible label for the button
func (pb *SpectrumPickerButton) Label(label string) *SpectrumPickerButton {
	pb.label = label
	return pb
}

// Position sets the position of the button
func (pb *SpectrumPickerButton) Position(position PickerButtonPosition) *SpectrumPickerButton {
	pb.position = position
	return pb
}

// Referrerpolicy sets the referrer policy for the button
func (pb *SpectrumPickerButton) Referrerpolicy(policy string) *SpectrumPickerButton {
	pb.referrerpolicy = policy
	return pb
}

// Rel sets the relationship of the linked URL
func (pb *SpectrumPickerButton) Rel(rel string) *SpectrumPickerButton {
	pb.rel = rel
	return pb
}

// Rounded sets whether the button has rounded corners
func (pb *SpectrumPickerButton) Rounded(rounded bool) *SpectrumPickerButton {
	pb.rounded = rounded
	return pb
}

// Size sets the visual size of the button
func (pb *SpectrumPickerButton) Size(size PickerButtonSize) *SpectrumPickerButton {
	pb.size = size
	return pb
}

// TabIndex sets the tab index of the button
func (pb *SpectrumPickerButton) TabIndex(tabIndex int) *SpectrumPickerButton {
	pb.tabIndex = tabIndex
	return pb
}

// Target sets where to display the linked URL
func (pb *SpectrumPickerButton) Target(target string) *SpectrumPickerButton {
	pb.target = target
	return pb
}

// Type sets the button type
func (pb *SpectrumPickerButton) Type(type_ string) *SpectrumPickerButton {
	pb.type_ = type_
	return pb
}

// Open sets whether the button is in open state
func (pb *SpectrumPickerButton) Open(open bool) *SpectrumPickerButton {
	pb.open = open
	return pb
}

// Invalid sets whether the button is in invalid state
func (pb *SpectrumPickerButton) Invalid(invalid bool) *SpectrumPickerButton {
	pb.invalid = invalid
	return pb
}

// Quiet sets whether the button has a quiet appearance
func (pb *SpectrumPickerButton) Quiet(quiet bool) *SpectrumPickerButton {
	pb.quiet = quiet
	return pb
}

// LabelSlot sets content for the label slot
func (pb *SpectrumPickerButton) LabelSlot(label app.UI) *SpectrumPickerButton {
	pb.labelSlot = label
	return pb
}

// Icon sets content for the icon slot
func (pb *SpectrumPickerButton) Icon(icon app.UI) *SpectrumPickerButton {
	pb.icon = icon
	return pb
}

// OnClick sets the click event handler
func (pb *SpectrumPickerButton) OnClick(handler app.EventHandler) *SpectrumPickerButton {
	pb.onClick = handler
	return pb
}

// Small sets size to small (s)
func (pb *SpectrumPickerButton) Small() *SpectrumPickerButton {
	return pb.Size(PickerButtonSizeS)
}

// Medium sets size to medium (m)
func (pb *SpectrumPickerButton) Medium() *SpectrumPickerButton {
	return pb.Size(PickerButtonSizeM)
}

// Large sets size to large (l)
func (pb *SpectrumPickerButton) Large() *SpectrumPickerButton {
	return pb.Size(PickerButtonSizeL)
}

// ExtraLarge sets size to extra large (xl)
func (pb *SpectrumPickerButton) ExtraLarge() *SpectrumPickerButton {
	return pb.Size(PickerButtonSizeXL)
}

// PositionLeft sets position to left
func (pb *SpectrumPickerButton) PositionLeft() *SpectrumPickerButton {
	return pb.Position(PickerButtonPositionLeft)
}

// PositionRight sets position to right
func (pb *SpectrumPickerButton) PositionRight() *SpectrumPickerButton {
	return pb.Position(PickerButtonPositionRight)
}

// Render renders the picker button component
func (pb *SpectrumPickerButton) Render() app.UI {
	pickerButton := app.Elem("sp-picker-button")

	// Set attributes
	if pb.active {
		pickerButton = pickerButton.Attr("active", true)
	}
	if pb.disabled {
		pickerButton = pickerButton.Attr("disabled", true)
	}
	if pb.download != "" {
		pickerButton = pickerButton.Attr("download", pb.download)
	}
	if pb.href != "" {
		pickerButton = pickerButton.Attr("href", pb.href)
	}
	if pb.label != "" {
		pickerButton = pickerButton.Attr("label", pb.label)
	}
	if pb.position != "" {
		pickerButton = pickerButton.Attr("position", string(pb.position))
	}
	if pb.referrerpolicy != "" {
		pickerButton = pickerButton.Attr("referrerpolicy", pb.referrerpolicy)
	}
	if pb.rel != "" {
		pickerButton = pickerButton.Attr("rel", pb.rel)
	}
	if pb.rounded {
		pickerButton = pickerButton.Attr("rounded", true)
	}
	if pb.size != "" {
		pickerButton = pickerButton.Attr("size", string(pb.size))
	}
	if pb.tabIndex != 0 {
		pickerButton = pickerButton.Attr("tabIndex", pb.tabIndex)
	}
	if pb.target != "" {
		pickerButton = pickerButton.Attr("target", pb.target)
	}
	if pb.type_ != "button" { // Only set if not the default
		pickerButton = pickerButton.Attr("type", pb.type_)
	}
	if pb.open {
		pickerButton = pickerButton.Attr("open", true)
	}
	if pb.invalid {
		pickerButton = pickerButton.Attr("invalid", true)
	}
	if pb.quiet {
		pickerButton = pickerButton.Attr("quiet", true)
	}

	// Add event handlers
	if pb.onClick != nil {
		pickerButton = pickerButton.OnClick(pb.onClick)
	}

	// Collect all slotted elements
	var elements []app.UI

	// Add label slot if provided
	if pb.labelSlot != nil {
		label := pb.labelSlot
		if labelWithSlot, ok := label.(interface{ Slot(string) app.UI }); ok {
			label = labelWithSlot.Slot("label")
		} else {
			label = app.Elem("span").
				Attr("slot", "label").
				Body(label)
		}
		elements = append(elements, label)
	}

	// Add icon if provided
	if pb.icon != nil {
		icon := pb.icon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		elements = append(elements, icon)
	}

	// Add all elements to the button
	if len(elements) > 0 {
		pickerButton = pickerButton.Body(elements...)
	}

	return pickerButton
}
