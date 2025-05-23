// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// DialogMode represents the The display mode of the dialog
type DialogMode string

// DialogMode values
const (
	DialogModeFullscreen         DialogMode = "fullscreen"
	DialogModeFullscreentakeover DialogMode = "fullscreenTakeover"
)

// DialogSize represents the Size of the dialog
type DialogSize string

// DialogSize values
const (
	DialogSizeS DialogSize = "s"
	DialogSizeM DialogSize = "m"
	DialogSizeL DialogSize = "l"
)

// spectrumDialog represents an sp-dialog component
type spectrumDialog struct {
	app.Compo
	*styler[*spectrumDialog]
	*classer[*spectrumDialog]
	*ider[*spectrumDialog]

	// Properties
	// Label for the dismiss button
	PropDismisslabel string
	// Whether the dialog has a close button
	PropDismissable bool
	// Whether the dialog represents an error state
	PropError bool
	// The display mode of the dialog
	PropMode DialogMode
	// Whether to hide the divider between header and content
	PropNodivider bool
	// Size of the dialog
	PropSize DialogSize
	// Alert dialog variant
	PropVariant string

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropButtonSlot  app.UI
	PropFooterSlot  app.UI
	PropHeadingSlot app.UI
	PropHeroSlot    app.UI

	// Event handlers
	PropOnClose app.EventHandler
}

// IDialog is the interface for sp-dialog component methods
type IDialog interface {
	app.UI
	Styler[IDialog]
	Classer[IDialog]
	Ider[IDialog]
	Dismisslabel(string) IDialog
	Dismissable(bool) IDialog
	SetDismissable() IDialog
	Error(bool) IDialog
	SetError() IDialog
	Mode(DialogMode) IDialog
	ModeFullscreen() IDialog
	ModeFullscreentakeover() IDialog
	Nodivider(bool) IDialog
	SetNodivider() IDialog
	Size(DialogSize) IDialog
	SizeS() IDialog
	SizeM() IDialog
	SizeL() IDialog
	Variant(string) IDialog

	Body(...app.UI) IDialog
	AddToBody(app.UI) IDialog
	Text(string) IDialog

	Button(app.UI) IDialog
	Footer(app.UI) IDialog
	Heading(app.UI) IDialog
	Hero(app.UI) IDialog

	OnClose(app.EventHandler) IDialog
}

// Dialog Displays important information that users need to acknowledge. They appear over the interface and block further interactions.
func Dialog() IDialog {
	element := &spectrumDialog{
		PropDismisslabel: "Close",
		PropDismissable:  false,
		PropError:        false,
		PropMode:         "",
		PropNodivider:    false,
		PropSize:         "",
		PropVariant:      "",
		PropBody:         []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Dismisslabel Label for the dismiss button
func (c *spectrumDialog) Dismisslabel(dismissLabel string) IDialog {
	c.PropDismisslabel = dismissLabel
	return c
}

// Dismissable Whether the dialog has a close button
func (c *spectrumDialog) Dismissable(dismissable bool) IDialog {
	c.PropDismissable = dismissable
	return c
}

func (c *spectrumDialog) SetDismissable() IDialog {
	return c.Dismissable(true)
}

// Error Whether the dialog represents an error state
func (c *spectrumDialog) Error(error bool) IDialog {
	c.PropError = error
	return c
}

func (c *spectrumDialog) SetError() IDialog {
	return c.Error(true)
}

// Mode The display mode of the dialog
func (c *spectrumDialog) Mode(mode DialogMode) IDialog {
	c.PropMode = mode
	return c
}

func (c *spectrumDialog) ModeFullscreen() IDialog {
	return c.Mode(DialogModeFullscreen)
}
func (c *spectrumDialog) ModeFullscreentakeover() IDialog {
	return c.Mode(DialogModeFullscreentakeover)
}

// Nodivider Whether to hide the divider between header and content
func (c *spectrumDialog) Nodivider(noDivider bool) IDialog {
	c.PropNodivider = noDivider
	return c
}

func (c *spectrumDialog) SetNodivider() IDialog {
	return c.Nodivider(true)
}

// Size Size of the dialog
func (c *spectrumDialog) Size(size DialogSize) IDialog {
	c.PropSize = size
	return c
}

func (c *spectrumDialog) SizeS() IDialog {
	return c.Size(DialogSizeS)
}
func (c *spectrumDialog) SizeM() IDialog {
	return c.Size(DialogSizeM)
}
func (c *spectrumDialog) SizeL() IDialog {
	return c.Size(DialogSizeL)
}

// Variant Alert dialog variant
func (c *spectrumDialog) Variant(variant string) IDialog {
	c.PropVariant = variant
	return c
}

// Body sets the content for the default slot
func (c *spectrumDialog) Body(elements ...app.UI) IDialog {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumDialog) AddToBody(element app.UI) IDialog {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumDialog) Text(text string) IDialog {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// Button elements addressed to this slot may be placed below the content when not delivered in a fullscreen mode
func (c *spectrumDialog) Button(content app.UI) IDialog {
	c.PropButtonSlot = content

	return c
}

// Content addressed to the footer will be placed below the main content and to the side of any [slot='button'] content
func (c *spectrumDialog) Footer(content app.UI) IDialog {
	c.PropFooterSlot = content

	return c
}

// Acts as the heading of the dialog. This should be an actual heading tag
func (c *spectrumDialog) Heading(content app.UI) IDialog {
	c.PropHeadingSlot = content

	return c
}

// Accepts a hero image to display at the top of the dialog
func (c *spectrumDialog) Hero(content app.UI) IDialog {
	c.PropHeroSlot = content

	return c
}

// Announces that the dialog has been closed
func (c *spectrumDialog) OnClose(handler app.EventHandler) IDialog {
	c.PropOnClose = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumDialog) Style(key, format string, values ...any) IDialog {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumDialog) Styles(styles map[string]string) IDialog {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumDialog) Class(class string) IDialog {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumDialog) Classes(classes ...string) IDialog {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumDialog) Id(id string) IDialog {
	return c.ider.Id(id)
}

// Render renders the sp-dialog component
func (c *spectrumDialog) Render() app.UI {
	element := app.Elem("sp-dialog")

	// Set attributes
	if c.PropDismisslabel != "" {
		element = element.Attr("dismissLabel", c.PropDismisslabel)
	}
	if c.PropDismissable {
		element = element.Attr("dismissable", true)
	}
	if c.PropError {
		element = element.Attr("error", true)
	}
	if c.PropMode != "" {
		element = element.Attr("mode", string(c.PropMode))
	}
	if c.PropNodivider {
		element = element.Attr("noDivider", true)
	}
	if c.PropSize != "" {
		element = element.Attr("size", string(c.PropSize))
	}
	if c.PropVariant != "" {
		element = element.Attr("variant", c.PropVariant)
	}

	// Add event handlers
	if c.PropOnClose != nil {
		element = element.On("close", c.PropOnClose)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add button slot
	if c.PropButtonSlot != nil {
		slotElem := c.PropButtonSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("button")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "button").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add footer slot
	if c.PropFooterSlot != nil {
		slotElem := c.PropFooterSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("footer")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "footer").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add heading slot
	if c.PropHeadingSlot != nil {
		slotElem := c.PropHeadingSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("heading")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "heading").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add hero slot
	if c.PropHeroSlot != nil {
		slotElem := c.PropHeroSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("hero")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "hero").
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
