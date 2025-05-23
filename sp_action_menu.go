// This file is generated by the generate_components.py script
// Do not edit this file manually

package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionMenuIcons represents the Controls icon display mode: 'only' shows only icons, 'none' shows no icons
type ActionMenuIcons string

// ActionMenuIcons values
const (
	ActionMenuIconsOnly ActionMenuIcons = "only"
	ActionMenuIconsNone ActionMenuIcons = "none"
)

// ActionMenuPlacement represents the Position of the overlay relative to the action button
type ActionMenuPlacement string

// ActionMenuPlacement values
const (
	ActionMenuPlacementTop         ActionMenuPlacement = "top"
	ActionMenuPlacementTopStart    ActionMenuPlacement = "top-start"
	ActionMenuPlacementTopEnd      ActionMenuPlacement = "top-end"
	ActionMenuPlacementRight       ActionMenuPlacement = "right"
	ActionMenuPlacementRightStart  ActionMenuPlacement = "right-start"
	ActionMenuPlacementRightEnd    ActionMenuPlacement = "right-end"
	ActionMenuPlacementBottom      ActionMenuPlacement = "bottom"
	ActionMenuPlacementBottomStart ActionMenuPlacement = "bottom-start"
	ActionMenuPlacementBottomEnd   ActionMenuPlacement = "bottom-end"
	ActionMenuPlacementLeft        ActionMenuPlacement = "left"
	ActionMenuPlacementLeftStart   ActionMenuPlacement = "left-start"
	ActionMenuPlacementLeftEnd     ActionMenuPlacement = "left-end"
)

// ActionMenuSelects represents the By default, action menu does not manage selection. Use 'single' to activate selection management.
type ActionMenuSelects string

// ActionMenuSelects values
const (
	ActionMenuSelectsSingle ActionMenuSelects = "single"
)

// ActionMenuSize represents the Size of the action menu button
type ActionMenuSize string

// ActionMenuSize values
const (
	ActionMenuSizeS  ActionMenuSize = "s"
	ActionMenuSizeM  ActionMenuSize = "m"
	ActionMenuSizeL  ActionMenuSize = "l"
	ActionMenuSizeXl ActionMenuSize = "xl"
)

// ActionMenuStaticcolor represents the Applies a specific color treatment to the button
type ActionMenuStaticcolor string

// ActionMenuStaticcolor values
const (
	ActionMenuStaticcolorWhite ActionMenuStaticcolor = "white"
	ActionMenuStaticcolorBlack ActionMenuStaticcolor = "black"
)

// spectrumActionMenu represents an sp-action-menu component
type spectrumActionMenu struct {
	app.Compo
	*styler[*spectrumActionMenu]
	*classer[*spectrumActionMenu]
	*ider[*spectrumActionMenu]

	// Properties
	// Disables the action menu
	PropDisabled bool
	// Whether the action menu has focus
	PropFocused bool
	// Forces the action menu to render as a popover on mobile instead of a tray
	PropForcepopover bool
	// Controls icon display mode: 'only' shows only icons, 'none' shows no icons
	PropIcons ActionMenuIcons
	// Marks the action menu as invalid
	PropInvalid bool
	// Accessible label for the action menu, particularly important when no visible label is provided
	PropLabel string
	// Whether the action menu overlay is open
	PropOpen bool
	// Whether the items are currently loading
	PropPending bool
	// Defines a string value that labels the action menu while it is in pending state
	PropPendinglabel string
	// Position of the overlay relative to the action button
	PropPlacement ActionMenuPlacement
	// Applies a less visually prominent style
	PropQuiet bool
	// Makes the action menu read-only
	PropReadonly bool
	// By default, action menu does not manage selection. Use 'single' to activate selection management.
	PropSelects ActionMenuSelects
	// Size of the action menu button
	PropSize ActionMenuSize
	// Applies a specific color treatment to the button
	PropStaticcolor ActionMenuStaticcolor
	// The current value of the action menu when selection is enabled
	PropValue string

	// Content for default slot
	PropBody []app.UI

	// Content slots
	PropDescriptionSlot app.UI
	PropIconSlot        app.UI
	PropLabelSlot       app.UI
	PropLabelOnlySlot   app.UI
	PropTooltipSlot     app.UI

	// Event handlers
	PropOnChange   app.EventHandler
	PropOnScroll   app.EventHandler
	PropOnSpOpened app.EventHandler
}

// IActionMenu is the interface for sp-action-menu component methods
type IActionMenu interface {
	app.UI
	Styler[IActionMenu]
	Classer[IActionMenu]
	Ider[IActionMenu]
	Disabled(bool) IActionMenu
	SetDisabled() IActionMenu
	Focused(bool) IActionMenu
	SetFocused() IActionMenu
	Forcepopover(bool) IActionMenu
	SetForcepopover() IActionMenu
	Icons(ActionMenuIcons) IActionMenu
	IconsOnly() IActionMenu
	IconsNone() IActionMenu
	Invalid(bool) IActionMenu
	SetInvalid() IActionMenu
	Label(string) IActionMenu
	Open(bool) IActionMenu
	SetOpen() IActionMenu
	Pending(bool) IActionMenu
	SetPending() IActionMenu
	Pendinglabel(string) IActionMenu
	Placement(ActionMenuPlacement) IActionMenu
	PlacementTop() IActionMenu
	PlacementTopStart() IActionMenu
	PlacementTopEnd() IActionMenu
	PlacementRight() IActionMenu
	PlacementRightStart() IActionMenu
	PlacementRightEnd() IActionMenu
	PlacementBottom() IActionMenu
	PlacementBottomStart() IActionMenu
	PlacementBottomEnd() IActionMenu
	PlacementLeft() IActionMenu
	PlacementLeftStart() IActionMenu
	PlacementLeftEnd() IActionMenu
	Quiet(bool) IActionMenu
	SetQuiet() IActionMenu
	Readonly(bool) IActionMenu
	SetReadonly() IActionMenu
	Selects(ActionMenuSelects) IActionMenu
	SelectsSingle() IActionMenu
	Size(ActionMenuSize) IActionMenu
	SizeS() IActionMenu
	SizeM() IActionMenu
	SizeL() IActionMenu
	SizeXl() IActionMenu
	Staticcolor(ActionMenuStaticcolor) IActionMenu
	StaticcolorWhite() IActionMenu
	StaticcolorBlack() IActionMenu
	Value(string) IActionMenu

	Body(...app.UI) IActionMenu
	AddToBody(app.UI) IActionMenu
	Text(string) IActionMenu

	Description(app.UI) IActionMenu
	Icon(app.UI) IActionMenu
	LabelContent(app.UI) IActionMenu
	LabelOnly(app.UI) IActionMenu
	Tooltip(app.UI) IActionMenu

	OnChange(app.EventHandler) IActionMenu
	OnScroll(app.EventHandler) IActionMenu
	OnSpOpened(app.EventHandler) IActionMenu
}

// ActionMenu An action button that triggers an overlay with menu items for activation. It can optionally manage selection of a single item.
func ActionMenu() IActionMenu {
	element := &spectrumActionMenu{
		PropDisabled:     false,
		PropFocused:      false,
		PropForcepopover: false,
		PropInvalid:      false,
		PropOpen:         false,
		PropPending:      false,
		PropPendinglabel: "Pending",
		PropPlacement:    ActionMenuPlacementBottomStart,
		PropQuiet:        false,
		PropReadonly:     false,
		PropSelects:      "",
		PropSize:         ActionMenuSizeM,
		PropValue:        "",
		PropBody:         []app.UI{},
	}

	element.styler = newStyler(element)
	element.classer = newClasser(element)
	element.ider = newIder(element)

	return element
}

// Disabled Disables the action menu
func (c *spectrumActionMenu) Disabled(disabled bool) IActionMenu {
	c.PropDisabled = disabled
	return c
}

func (c *spectrumActionMenu) SetDisabled() IActionMenu {
	return c.Disabled(true)
}

// Focused Whether the action menu has focus
func (c *spectrumActionMenu) Focused(focused bool) IActionMenu {
	c.PropFocused = focused
	return c
}

func (c *spectrumActionMenu) SetFocused() IActionMenu {
	return c.Focused(true)
}

// Forcepopover Forces the action menu to render as a popover on mobile instead of a tray
func (c *spectrumActionMenu) Forcepopover(forcePopover bool) IActionMenu {
	c.PropForcepopover = forcePopover
	return c
}

func (c *spectrumActionMenu) SetForcepopover() IActionMenu {
	return c.Forcepopover(true)
}

// Icons Controls icon display mode: 'only' shows only icons, 'none' shows no icons
func (c *spectrumActionMenu) Icons(icons ActionMenuIcons) IActionMenu {
	c.PropIcons = icons
	return c
}

func (c *spectrumActionMenu) IconsOnly() IActionMenu {
	return c.Icons(ActionMenuIconsOnly)
}
func (c *spectrumActionMenu) IconsNone() IActionMenu {
	return c.Icons(ActionMenuIconsNone)
}

// Invalid Marks the action menu as invalid
func (c *spectrumActionMenu) Invalid(invalid bool) IActionMenu {
	c.PropInvalid = invalid
	return c
}

func (c *spectrumActionMenu) SetInvalid() IActionMenu {
	return c.Invalid(true)
}

// Label Accessible label for the action menu, particularly important when no visible label is provided
func (c *spectrumActionMenu) Label(label string) IActionMenu {
	c.PropLabel = label
	return c
}

// Open Whether the action menu overlay is open
func (c *spectrumActionMenu) Open(open bool) IActionMenu {
	c.PropOpen = open
	return c
}

func (c *spectrumActionMenu) SetOpen() IActionMenu {
	return c.Open(true)
}

// Pending Whether the items are currently loading
func (c *spectrumActionMenu) Pending(pending bool) IActionMenu {
	c.PropPending = pending
	return c
}

func (c *spectrumActionMenu) SetPending() IActionMenu {
	return c.Pending(true)
}

// Pendinglabel Defines a string value that labels the action menu while it is in pending state
func (c *spectrumActionMenu) Pendinglabel(pendingLabel string) IActionMenu {
	c.PropPendinglabel = pendingLabel
	return c
}

// Placement Position of the overlay relative to the action button
func (c *spectrumActionMenu) Placement(placement ActionMenuPlacement) IActionMenu {
	c.PropPlacement = placement
	return c
}

func (c *spectrumActionMenu) PlacementTop() IActionMenu {
	return c.Placement(ActionMenuPlacementTop)
}
func (c *spectrumActionMenu) PlacementTopStart() IActionMenu {
	return c.Placement(ActionMenuPlacementTopStart)
}
func (c *spectrumActionMenu) PlacementTopEnd() IActionMenu {
	return c.Placement(ActionMenuPlacementTopEnd)
}
func (c *spectrumActionMenu) PlacementRight() IActionMenu {
	return c.Placement(ActionMenuPlacementRight)
}
func (c *spectrumActionMenu) PlacementRightStart() IActionMenu {
	return c.Placement(ActionMenuPlacementRightStart)
}
func (c *spectrumActionMenu) PlacementRightEnd() IActionMenu {
	return c.Placement(ActionMenuPlacementRightEnd)
}
func (c *spectrumActionMenu) PlacementBottom() IActionMenu {
	return c.Placement(ActionMenuPlacementBottom)
}
func (c *spectrumActionMenu) PlacementBottomStart() IActionMenu {
	return c.Placement(ActionMenuPlacementBottomStart)
}
func (c *spectrumActionMenu) PlacementBottomEnd() IActionMenu {
	return c.Placement(ActionMenuPlacementBottomEnd)
}
func (c *spectrumActionMenu) PlacementLeft() IActionMenu {
	return c.Placement(ActionMenuPlacementLeft)
}
func (c *spectrumActionMenu) PlacementLeftStart() IActionMenu {
	return c.Placement(ActionMenuPlacementLeftStart)
}
func (c *spectrumActionMenu) PlacementLeftEnd() IActionMenu {
	return c.Placement(ActionMenuPlacementLeftEnd)
}

// Quiet Applies a less visually prominent style
func (c *spectrumActionMenu) Quiet(quiet bool) IActionMenu {
	c.PropQuiet = quiet
	return c
}

func (c *spectrumActionMenu) SetQuiet() IActionMenu {
	return c.Quiet(true)
}

// Readonly Makes the action menu read-only
func (c *spectrumActionMenu) Readonly(readonly bool) IActionMenu {
	c.PropReadonly = readonly
	return c
}

func (c *spectrumActionMenu) SetReadonly() IActionMenu {
	return c.Readonly(true)
}

// Selects By default, action menu does not manage selection. Use 'single' to activate selection management.
func (c *spectrumActionMenu) Selects(selects ActionMenuSelects) IActionMenu {
	c.PropSelects = selects
	return c
}

func (c *spectrumActionMenu) SelectsSingle() IActionMenu {
	return c.Selects(ActionMenuSelectsSingle)
}

// Size Size of the action menu button
func (c *spectrumActionMenu) Size(size ActionMenuSize) IActionMenu {
	c.PropSize = size
	return c
}

func (c *spectrumActionMenu) SizeS() IActionMenu {
	return c.Size(ActionMenuSizeS)
}
func (c *spectrumActionMenu) SizeM() IActionMenu {
	return c.Size(ActionMenuSizeM)
}
func (c *spectrumActionMenu) SizeL() IActionMenu {
	return c.Size(ActionMenuSizeL)
}
func (c *spectrumActionMenu) SizeXl() IActionMenu {
	return c.Size(ActionMenuSizeXl)
}

// Staticcolor Applies a specific color treatment to the button
func (c *spectrumActionMenu) Staticcolor(staticColor ActionMenuStaticcolor) IActionMenu {
	c.PropStaticcolor = staticColor
	return c
}

func (c *spectrumActionMenu) StaticcolorWhite() IActionMenu {
	return c.Staticcolor(ActionMenuStaticcolorWhite)
}
func (c *spectrumActionMenu) StaticcolorBlack() IActionMenu {
	return c.Staticcolor(ActionMenuStaticcolorBlack)
}

// Value The current value of the action menu when selection is enabled
func (c *spectrumActionMenu) Value(value string) IActionMenu {
	c.PropValue = value
	return c
}

// Body sets the content for the default slot
func (c *spectrumActionMenu) Body(elements ...app.UI) IActionMenu {
	c.PropBody = elements
	return c
}

// AddToBody adds a UI element to the default slot
func (c *spectrumActionMenu) AddToBody(element app.UI) IActionMenu {
	c.PropBody = append(c.PropBody, element)
	return c
}

// Text sets text content for the default slot
func (c *spectrumActionMenu) Text(text string) IActionMenu {
	c.PropBody = []app.UI{app.Text(text)}
	return c
}

// The description content for the action menu
func (c *spectrumActionMenu) Description(content app.UI) IActionMenu {
	c.PropDescriptionSlot = content

	return c
}

// The icon to use for the Action Menu instead of the default
func (c *spectrumActionMenu) Icon(content app.UI) IActionMenu {
	c.PropIconSlot = content

	return c
}

// The label to use for the Action Menu (with icon space reserved)
func (c *spectrumActionMenu) LabelContent(content app.UI) IActionMenu {
	c.PropLabelSlot = content

	return c
}

// The label to use for the Action Menu (no icon space reserved)
func (c *spectrumActionMenu) LabelOnly(content app.UI) IActionMenu {
	c.PropLabelOnlySlot = content

	return c
}

// Tooltip to be applied to the Action Button
func (c *spectrumActionMenu) Tooltip(content app.UI) IActionMenu {
	c.PropTooltipSlot = content

	return c
}

// Announces that the value of the element has changed
func (c *spectrumActionMenu) OnChange(handler app.EventHandler) IActionMenu {
	c.PropOnChange = handler

	return c
}

// Fired when the menu content is scrolled
func (c *spectrumActionMenu) OnScroll(handler app.EventHandler) IActionMenu {
	c.PropOnScroll = handler

	return c
}

// Announces that the overlay has been opened
func (c *spectrumActionMenu) OnSpOpened(handler app.EventHandler) IActionMenu {
	c.PropOnSpOpened = handler

	return c
}

// Style sets a style property with a value
func (c *spectrumActionMenu) Style(key, format string, values ...any) IActionMenu {
	return c.styler.Style(key, format, values...)
}

// Styles sets multiple style properties
func (c *spectrumActionMenu) Styles(styles map[string]string) IActionMenu {
	return c.styler.Styles(styles)
}

// Class adds a class to the element
func (c *spectrumActionMenu) Class(class string) IActionMenu {
	return c.classer.Class(class)
}

// Classes adds multiple classes to the element
func (c *spectrumActionMenu) Classes(classes ...string) IActionMenu {
	return c.classer.Classes(classes...)
}

// Id sets the id of the element
func (c *spectrumActionMenu) Id(id string) IActionMenu {
	return c.ider.Id(id)
}

// Render renders the sp-action-menu component
func (c *spectrumActionMenu) Render() app.UI {
	element := app.Elem("sp-action-menu")

	// Set attributes
	if c.PropDisabled {
		element = element.Attr("disabled", true)
	}
	if c.PropFocused {
		element = element.Attr("focused", true)
	}
	if c.PropForcepopover {
		element = element.Attr("forcePopover", true)
	}
	if c.PropIcons != "" {
		element = element.Attr("icons", string(c.PropIcons))
	}
	if c.PropInvalid {
		element = element.Attr("invalid", true)
	}
	if c.PropLabel != "" {
		element = element.Attr("label", c.PropLabel)
	}
	if c.PropOpen {
		element = element.Attr("open", true)
	}
	if c.PropPending {
		element = element.Attr("pending", true)
	}
	if c.PropPendinglabel != "" {
		element = element.Attr("pendingLabel", c.PropPendinglabel)
	}
	if c.PropPlacement != "" {
		element = element.Attr("placement", string(c.PropPlacement))
	}
	if c.PropQuiet {
		element = element.Attr("quiet", true)
	}
	if c.PropReadonly {
		element = element.Attr("readonly", true)
	}
	if c.PropSelects != "" {
		element = element.Attr("selects", string(c.PropSelects))
	}
	if c.PropSize != "" {
		element = element.Attr("size", string(c.PropSize))
	}
	if c.PropStaticcolor != "" {
		element = element.Attr("staticColor", string(c.PropStaticcolor))
	}
	if c.PropValue != "" {
		element = element.Attr("value", c.PropValue)
	}

	// Add event handlers
	if c.PropOnChange != nil {
		element = element.On("change", c.PropOnChange)
	}
	if c.PropOnScroll != nil {
		element = element.On("scroll", c.PropOnScroll)
	}
	if c.PropOnSpOpened != nil {
		element = element.On("sp-opened", c.PropOnSpOpened)
	}

	// Add slots and children
	slotElements := []app.UI{}

	// Add content for default slot if specified
	if len(c.PropBody) > 0 {
		slotElements = append(slotElements, c.PropBody...)
	}

	// Add description slot
	if c.PropDescriptionSlot != nil {
		slotElem := c.PropDescriptionSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("description")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "description").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
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
	// Add label-only slot
	if c.PropLabelOnlySlot != nil {
		slotElem := c.PropLabelOnlySlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("label-only")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "label-only").
				Body(slotElem)
		}
		slotElements = append(slotElements, slotElem)
	}
	// Add tooltip slot
	if c.PropTooltipSlot != nil {
		slotElem := c.PropTooltipSlot
		if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
			slotElem = slotWithAttr.Slot("tooltip")
		} else {
			slotElem = app.Elem("div").
				Attr("slot", "tooltip").
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
