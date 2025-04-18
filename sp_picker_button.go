// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerButtonPosition represents the Position relative to textfield
type PickerButtonPosition string

// PickerButtonPosition values
const (
	PickerButtonPositionLeft  PickerButtonPosition = "left"
	PickerButtonPositionRight PickerButtonPosition = "right"
)

// PickerButtonType represents the Button type
type PickerButtonType string

// PickerButtonType values
const (
	PickerButtonTypeButton PickerButtonType = "button"
	PickerButtonTypeSubmit PickerButtonType = "submit"
	PickerButtonTypeReset  PickerButtonType = "reset"
)

// spectrumPickerButton represents an sp-picker-button component
type spectrumPickerButton struct {
	app.Compo
	*styler[*spectrumPickerButton]
	*classer[*spectrumPickerButton]
	*ider[*spectrumPickerButton]

	// Properties
	// Whether the button is in active state
	PropActive bool
	// Disable this control
	PropDisabled bool
	// Causes browser to treat linked URL as download
	PropDownload string
	// URL the button links to
	PropHref string
	// Accessible label for the component
	PropLabel string
	// Position relative to textfield
	PropPosition PickerButtonPosition
	// Less pronounced visual delivery
	PropQuiet bool
	// Deeply rounded corners in express system
	PropRounded bool
	// Tab index for the control
	PropTabindex float64
	// Button type
	PropType PickerButtonType

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropIconSlot  app.UI
	PropLabelSlot app.UI

	// Event handlers
	PropOnClick app.EventHandler
}

// IPickerButton is the interface for sp-picker-button component methods
type IPickerButton interface {
	app.UI
	Styler[IPickerButton]
	Classer[IPickerButton]
	Ider[IPickerButton]
	Active(bool) IPickerButton
	SetActive() IPickerButton
	Disabled(bool) IPickerButton
	SetDisabled() IPickerButton
	Download(string) IPickerButton
	Href(string) IPickerButton
	Label(string) IPickerButton
	Position(PickerButtonPosition) IPickerButton
	PositionLeft() IPickerButton
	PositionRight() IPickerButton
	Quiet(bool) IPickerButton
	SetQuiet() IPickerButton
	Rounded(bool) IPickerButton
	SetRounded() IPickerButton
	Tabindex(float64) IPickerButton
	Type(PickerButtonType) IPickerButton
	TypeButton() IPickerButton
	TypeSubmit() IPickerButton
	TypeReset() IPickerButton

	Body(...app.UI) IPickerButton
	AddToBody(app.UI) IPickerButton
	Text(string) IPickerButton

	Icon(app.UI) IPickerButton
	LabelContent(app.UI) IPickerButton

	OnClick(app.EventHandler) IPickerButton
}

// PickerButton A sub-component used in patterns like combobox to pair a button interface with a text input.
func PickerButton() IPickerButton {
	element := &spectrumPickerButton{
		PropActive:   false,
		PropDisabled: false,
		PropDownload: "",
		PropHref:     "",
		PropLabel:    "",
		PropPosition: PickerButtonPositionRight,
		PropQuiet:    false,
		PropRounded:  false,
		PropTabindex: 0,
		PropType:     PickerButtonTypeButton,
		PropBody:     []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Active Whether the button is in active state
func (c *spectrumPickerButton) Active(active bool) IPickerButton {
	c.PropActive = active
	return c
}

func (c *spectrumPickerButton) SetActive() IPickerButton {
	return c.Active(true)
}

// Disabled Disable this control
func (c *spectrumPickerButton) Disabled(disabled bool) IPickerButton {
	c.PropDisabled = disabled
	return c
}

func (c *spectrumPickerButton) SetDisabled() IPickerButton {
	return c.Disabled(true)
}

// Download Causes browser to treat linked URL as download
func (c *spectrumPickerButton) Download(download string) IPickerButton {
	c.PropDownload = download
	return c
}

// Href URL the button links to
func (c *spectrumPickerButton) Href(href string) IPickerButton {
	c.PropHref = href
	return c
}

// Label Accessible label for the component
func (c *spectrumPickerButton) Label(label string) IPickerButton {
	c.PropLabel = label
	return c
}

// Position Position relative to textfield
func (c *spectrumPickerButton) Position(position PickerButtonPosition) IPickerButton {
	c.PropPosition = position
	return c
}

func (c *spectrumPickerButton) PositionLeft() IPickerButton {
	return c.Position(PickerButtonPositionLeft)
}
func (c *spectrumPickerButton) PositionRight() IPickerButton {
	return c.Position(PickerButtonPositionRight)
}

// Quiet Less pronounced visual delivery
func (c *spectrumPickerButton) Quiet(quiet bool) IPickerButton {
	c.PropQuiet = quiet
	return c
}

func (c *spectrumPickerButton) SetQuiet() IPickerButton {
	return c.Quiet(true)
}

// Rounded Deeply rounded corners in express system
func (c *spectrumPickerButton) Rounded(rounded bool) IPickerButton {
	c.PropRounded = rounded
	return c
}

func (c *spectrumPickerButton) SetRounded() IPickerButton {
	return c.Rounded(true)
}

// Tabindex Tab index for the control
func (c *spectrumPickerButton) Tabindex(tabIndex float64) IPickerButton {
	c.PropTabindex = tabIndex
	return c
}

// Type Button type
func (c *spectrumPickerButton) Type(typeValue PickerButtonType) IPickerButton {
	c.PropType = typeValue
	return c
}

func (c *spectrumPickerButton) TypeButton() IPickerButton {
	return c.Type(PickerButtonTypeButton)
}
func (c *spectrumPickerButton) TypeSubmit() IPickerButton {
	return c.Type(PickerButtonTypeSubmit)
}
func (c *spectrumPickerButton) TypeReset() IPickerButton {
	return c.Type(PickerButtonTypeReset)
}

// Body sets the content for the default slot
func (c *spectrumPickerButton) Body(elements ...app.UI) IPickerButton {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumPickerButton) AddToBody(element app.UI) IPickerButton {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumPickerButton) Text(text string) IPickerButton {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Icon element to display
func (c *spectrumPickerButton) Icon(content app.UI) IPickerButton {
	c.PropIconSlot = content

	return c
}

// Content for the button label
func (c *spectrumPickerButton) LabelContent(content app.UI) IPickerButton {
	c.PropLabelSlot = content

	return c
}

// Fired when the button is clicked
func (c *spectrumPickerButton) OnClick(handler app.EventHandler) IPickerButton {
	c.PropOnClick = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumPickerButton) Style(key, format string, values ...any) IPickerButton {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumPickerButton) Styles(styles map[string]string) IPickerButton {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumPickerButton) Class(class string) IPickerButton {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumPickerButton) Classes(classes ...string) IPickerButton {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumPickerButton) Id(id string) IPickerButton {
	return c.ider.Id(id)
}

// Render renders the sp-picker-button component
func (c *spectrumPickerButton) Render() app.UI {
	element := app.Elem("sp-picker-button")

	// Set attributes
	if c.PropActive {
		element = element.Attr("active", true)
	}
	if c.PropDisabled {
		element = element.Attr("disabled", true)
	}
	if c.PropDownload != "" {
		element = element.Attr("download", c.PropDownload)
	}
	if c.PropHref != "" {
		element = element.Attr("href", c.PropHref)
	}
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}
	if c.PropPosition != "" {
		element = element.Attr("position", string(c.PropPosition))
	}
	if c.PropQuiet {
		element = element.Attr("quiet", true)
	}
	if c.PropRounded {
		element = element.Attr("rounded", true)
	}
	if c.PropTabindex != 0 {
		element = element.Attr("tabIndex", c.PropTabindex)
	}
	if c.PropType != "" {
		element = element.Attr("type", string(c.PropType))
	}

	// Add event handlers
	if c.PropOnClick != nil {
		element = element.On("click", c.PropOnClick)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add icon slot
	if c.PropIconSlot != nil {
		slotElem := c.PropIconSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("icon")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "icon").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add label slot
	if c.PropLabelSlot != nil {
		slotElem := c.PropLabelSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("label")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "label").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}

	// Add all elements to the component
	if len(slotElements) > 0 {
		element = element.Body(slotElements...)
	}

	// Apply styles, classes, and id
	element = element.Styles(c.styler.styles)

	// Apply classes if any
	if len(c.classer.classes) > 0 {
		element = element.Class(c.classer.classes...)
	}

	// Apply id if set
	if c.ider.id != "" {
		element = element.ID(c.ider.id)
	}

	return element
}
