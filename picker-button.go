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

// spectrumPickerButton represents an sp-picker-button component
type spectrumPickerButton struct {
	app.Compo

	// Properties
	PropActive         bool
	PropDisabled       bool
	PropDownload       string
	PropHref           string
	PropLabel          string
	PropPosition       PickerButtonPosition
	PropReferrerpolicy string
	PropRel            string
	PropRounded        bool
	PropSize           PickerButtonSize
	PropTabIndex       int
	PropTarget         string
	PropType           string
	PropOpen           bool
	PropInvalid        bool
	PropQuiet          bool

	// Slots
	PropLabelSlot app.UI
	PropIcon      app.UI

	// Event handlers
	PropOnClick app.EventHandler
}

// PickerButton creates a new picker button component
func PickerButton() *spectrumPickerButton {
	return &spectrumPickerButton{
		PropPosition: PickerButtonPositionRight, // Default position
		PropSize:     PickerButtonSizeM,         // Default size
		PropType:     "button",                  // Default type
	}
}

// Active sets whether the button is active
func (pb *spectrumPickerButton) Active(active bool) *spectrumPickerButton {
	pb.PropActive = active
	return pb
}

// Disabled sets whether the button is disabled
func (pb *spectrumPickerButton) Disabled(disabled bool) *spectrumPickerButton {
	pb.PropDisabled = disabled
	return pb
}

// Download sets the download attribute
func (pb *spectrumPickerButton) Download(download string) *spectrumPickerButton {
	pb.PropDownload = download
	return pb
}

// Href sets the URL the button points to
func (pb *spectrumPickerButton) Href(href string) *spectrumPickerButton {
	pb.PropHref = href
	return pb
}

// Label sets the accessible label for the button
func (pb *spectrumPickerButton) Label(label string) *spectrumPickerButton {
	pb.PropLabel = label
	return pb
}

// Position sets the position of the button
func (pb *spectrumPickerButton) Position(position PickerButtonPosition) *spectrumPickerButton {
	pb.PropPosition = position
	return pb
}

// Referrerpolicy sets the referrer policy for the button
func (pb *spectrumPickerButton) Referrerpolicy(policy string) *spectrumPickerButton {
	pb.PropReferrerpolicy = policy
	return pb
}

// Rel sets the relationship of the linked URL
func (pb *spectrumPickerButton) Rel(rel string) *spectrumPickerButton {
	pb.PropRel = rel
	return pb
}

// Rounded sets whether the button has rounded corners
func (pb *spectrumPickerButton) Rounded(rounded bool) *spectrumPickerButton {
	pb.PropRounded = rounded
	return pb
}

// Size sets the visual size of the button
func (pb *spectrumPickerButton) Size(size PickerButtonSize) *spectrumPickerButton {
	pb.PropSize = size
	return pb
}

// TabIndex sets the tab index of the button
func (pb *spectrumPickerButton) TabIndex(tabIndex int) *spectrumPickerButton {
	pb.PropTabIndex = tabIndex
	return pb
}

// Target sets where to display the linked URL
func (pb *spectrumPickerButton) Target(target string) *spectrumPickerButton {
	pb.PropTarget = target
	return pb
}

// Type sets the button type
func (pb *spectrumPickerButton) Type(type_ string) *spectrumPickerButton {
	pb.PropType = type_
	return pb
}

// Open sets whether the button is in open state
func (pb *spectrumPickerButton) Open(open bool) *spectrumPickerButton {
	pb.PropOpen = open
	return pb
}

// Invalid sets whether the button is in invalid state
func (pb *spectrumPickerButton) Invalid(invalid bool) *spectrumPickerButton {
	pb.PropInvalid = invalid
	return pb
}

// Quiet sets whether the button has a quiet appearance
func (pb *spectrumPickerButton) Quiet(quiet bool) *spectrumPickerButton {
	pb.PropQuiet = quiet
	return pb
}

// LabelSlot sets content for the label slot
func (pb *spectrumPickerButton) LabelSlot(label app.UI) *spectrumPickerButton {
	pb.PropLabelSlot = label
	return pb
}

// Icon sets content for the icon slot
func (pb *spectrumPickerButton) Icon(icon app.UI) *spectrumPickerButton {
	pb.PropIcon = icon
	return pb
}

// OnClick sets the click event handler
func (pb *spectrumPickerButton) OnClick(handler app.EventHandler) *spectrumPickerButton {
	pb.PropOnClick = handler
	return pb
}

// Small sets size to small (s)
func (pb *spectrumPickerButton) Small() *spectrumPickerButton {
	return pb.Size(PickerButtonSizeS)
}

// Medium sets size to medium (m)
func (pb *spectrumPickerButton) Medium() *spectrumPickerButton {
	return pb.Size(PickerButtonSizeM)
}

// Large sets size to large (l)
func (pb *spectrumPickerButton) Large() *spectrumPickerButton {
	return pb.Size(PickerButtonSizeL)
}

// ExtraLarge sets size to extra large (xl)
func (pb *spectrumPickerButton) ExtraLarge() *spectrumPickerButton {
	return pb.Size(PickerButtonSizeXL)
}

// PositionLeft sets position to left
func (pb *spectrumPickerButton) PositionLeft() *spectrumPickerButton {
	return pb.Position(PickerButtonPositionLeft)
}

// PositionRight sets position to right
func (pb *spectrumPickerButton) PositionRight() *spectrumPickerButton {
	return pb.Position(PickerButtonPositionRight)
}

// Render renders the picker button component
func (pb *spectrumPickerButton) Render() app.UI {
	pickerButton := app.Elem("sp-picker-button")

	// Set attributes
	if pb.PropActive {
		pickerButton = pickerButton.Attr("active", true)
	}
	if pb.PropDisabled {
		pickerButton = pickerButton.Attr("disabled", true)
	}
	if pb.PropDownload != "" {
		pickerButton = pickerButton.Attr("download", pb.PropDownload)
	}
	if pb.PropHref != "" {
		pickerButton = pickerButton.Attr("href", pb.PropHref)
	}
	if pb.PropLabel != "" {
		pickerButton = pickerButton.Attr("label", pb.PropLabel)
	}
	if pb.PropPosition != "" {
		pickerButton = pickerButton.Attr("position", string(pb.PropPosition))
	}
	if pb.PropReferrerpolicy != "" {
		pickerButton = pickerButton.Attr("referrerpolicy", pb.PropReferrerpolicy)
	}
	if pb.PropRel != "" {
		pickerButton = pickerButton.Attr("rel", pb.PropRel)
	}
	if pb.PropRounded {
		pickerButton = pickerButton.Attr("rounded", true)
	}
	if pb.PropSize != "" {
		pickerButton = pickerButton.Attr("size", string(pb.PropSize))
	}
	if pb.PropTabIndex != 0 {
		pickerButton = pickerButton.Attr("tabIndex", pb.PropTabIndex)
	}
	if pb.PropTarget != "" {
		pickerButton = pickerButton.Attr("target", pb.PropTarget)
	}
	if pb.PropType != "button" { // Only set if not the default
		pickerButton = pickerButton.Attr("type", pb.PropType)
	}
	if pb.PropOpen {
		pickerButton = pickerButton.Attr("open", true)
	}
	if pb.PropInvalid {
		pickerButton = pickerButton.Attr("invalid", true)
	}
	if pb.PropQuiet {
		pickerButton = pickerButton.Attr("quiet", true)
	}

	// Add event handlers
	if pb.PropOnClick != nil {
		pickerButton = pickerButton.On("click", pb.PropOnClick)
	}

	// Collect all slotted elements
	var elements []app.UI

	// Add label slot if provided
	if pb.PropLabelSlot != nil {
		label := pb.PropLabelSlot
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
	if pb.PropIcon != nil {
		icon := pb.PropIcon
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
