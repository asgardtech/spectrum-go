package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconsType represents the icon display mode for an action menu
type IconsType string

const (
	IconsTypeNone IconsType = "none"
	IconsTypeOnly IconsType = "only"
)

// spectrumActionMenu represents an sp-action-menu component
type spectrumActionMenu struct {
	app.Compo

	// Properties
	PropDisabled     bool
	PropFocused      bool
	PropForcePopover bool
	PropIcons        IconsType
	PropInvalid      bool
	PropLabel        string
	PropOpen         bool
	PropPending      bool
	PropPendingLabel string
	PropPlacement    string
	PropQuiet        bool
	PropReadonly     bool
	PropSelects      string
	PropStaticColor  string
	PropValue        string

	// Content slots
	PropDescription   app.UI
	PropIconSlot      app.UI
	PropLabelSlot     app.UI
	PropLabelOnlySlot app.UI
	PropTooltipSlot   app.UI

	// Child components (menu items)
	PropChildren []app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnScroll app.EventHandler
	PropOnOpened app.EventHandler
}

// ActionMenu creates a new action menu component
func ActionMenu() *spectrumActionMenu {
	return &spectrumActionMenu{
		PropPlacement:    "bottom-start", // Default placement
		PropPendingLabel: "Pending",      // Default pending label
	}
}

// Disabled sets whether the action menu is disabled
func (am *spectrumActionMenu) Disabled(disabled bool) *spectrumActionMenu {
	am.PropDisabled = disabled
	return am
}

// Focused sets whether the action menu is focused
func (am *spectrumActionMenu) Focused(focused bool) *spectrumActionMenu {
	am.PropFocused = focused
	return am
}

// ForcePopover forces the action menu to render as a popover on mobile instead of a tray
func (am *spectrumActionMenu) ForcePopover(force bool) *spectrumActionMenu {
	am.PropForcePopover = force
	return am
}

// Icons sets the icon display mode (none, only)
func (am *spectrumActionMenu) Icons(icons IconsType) *spectrumActionMenu {
	am.PropIcons = icons
	return am
}

// IconsNone sets the icon display mode to none
func (am *spectrumActionMenu) IconsNone() *spectrumActionMenu {
	return am.Icons(IconsTypeNone)
}

// IconsOnly sets the icon display mode to only
func (am *spectrumActionMenu) IconsOnly() *spectrumActionMenu {
	return am.Icons(IconsTypeOnly)
}

// Invalid sets whether the action menu is invalid
func (am *spectrumActionMenu) Invalid(invalid bool) *spectrumActionMenu {
	am.PropInvalid = invalid
	return am
}

// Label sets the accessible label for the action menu
func (am *spectrumActionMenu) Label(label string) *spectrumActionMenu {
	am.PropLabel = label
	return am
}

// Open sets whether the action menu is open
func (am *spectrumActionMenu) Open(open bool) *spectrumActionMenu {
	am.PropOpen = open
	return am
}

// Pending sets whether the items are currently loading
func (am *spectrumActionMenu) Pending(pending bool) *spectrumActionMenu {
	am.PropPending = pending
	return am
}

// PendingLabel sets the label shown while in pending state
func (am *spectrumActionMenu) PendingLabel(label string) *spectrumActionMenu {
	am.PropPendingLabel = label
	return am
}

// Placement sets the placement of the overlay
func (am *spectrumActionMenu) Placement(placement string) *spectrumActionMenu {
	am.PropPlacement = placement
	return am
}

// Quiet sets whether the action menu has a quiet appearance
func (am *spectrumActionMenu) Quiet(quiet bool) *spectrumActionMenu {
	am.PropQuiet = quiet
	return am
}

// Readonly sets whether the action menu is readonly
func (am *spectrumActionMenu) Readonly(readonly bool) *spectrumActionMenu {
	am.PropReadonly = readonly
	return am
}

// Selects sets the selection mode (single or undefined)
func (am *spectrumActionMenu) Selects(selects string) *spectrumActionMenu {
	am.PropSelects = selects
	return am
}

// SelectsSingle sets the selection mode to single
func (am *spectrumActionMenu) SelectsSingle() *spectrumActionMenu {
	return am.Selects("single")
}

// StaticColor sets the static color of the action menu
func (am *spectrumActionMenu) StaticColor(color string) *spectrumActionMenu {
	am.PropStaticColor = color
	return am
}

// StaticColorWhite sets the static color to white
func (am *spectrumActionMenu) StaticColorWhite() *spectrumActionMenu {
	return am.StaticColor("white")
}

// StaticColorBlack sets the static color to black
func (am *spectrumActionMenu) StaticColorBlack() *spectrumActionMenu {
	return am.StaticColor("black")
}

// Value sets the value of the action menu
func (am *spectrumActionMenu) Value(value string) *spectrumActionMenu {
	am.PropValue = value
	return am
}

// Description sets the description content in the description slot
func (am *spectrumActionMenu) Description(description app.UI) *spectrumActionMenu {
	am.PropDescription = description
	return am
}

// Icon sets the icon in the icon slot
func (am *spectrumActionMenu) Icon(icon app.UI) *spectrumActionMenu {
	am.PropIconSlot = icon
	return am
}

// LabelSlot sets the content in the label slot
func (am *spectrumActionMenu) LabelSlot(label app.UI) *spectrumActionMenu {
	am.PropLabelSlot = label
	return am
}

// LabelOnly sets the content in the label-only slot (no icon space reserved)
func (am *spectrumActionMenu) LabelOnly(label app.UI) *spectrumActionMenu {
	am.PropLabelOnlySlot = label
	return am
}

// Tooltip sets the tooltip to be applied to the action button
func (am *spectrumActionMenu) Tooltip(tooltip app.UI) *spectrumActionMenu {
	am.PropTooltipSlot = tooltip
	return am
}

// Children sets the child components (menu items)
func (am *spectrumActionMenu) Children(children ...app.UI) *spectrumActionMenu {
	am.PropChildren = children
	return am
}

// AddItem adds a menu item to the action menu
func (am *spectrumActionMenu) AddItem(item app.UI) *spectrumActionMenu {
	am.PropChildren = append(am.PropChildren, item)
	return am
}

// OnChange sets the change event handler
func (am *spectrumActionMenu) OnChange(handler app.EventHandler) *spectrumActionMenu {
	am.PropOnChange = handler
	return am
}

// OnScroll sets the scroll event handler
func (am *spectrumActionMenu) OnScroll(handler app.EventHandler) *spectrumActionMenu {
	am.PropOnScroll = handler
	return am
}

// OnOpened sets the opened event handler
func (am *spectrumActionMenu) OnOpened(handler app.EventHandler) *spectrumActionMenu {
	am.PropOnOpened = handler
	return am
}

// Render renders the action menu component
func (am *spectrumActionMenu) Render() app.UI {
	actionMenu := app.Elem("sp-action-menu")

	// Set attributes
	if am.PropDisabled {
		actionMenu = actionMenu.Attr("disabled", true)
	}
	if am.PropFocused {
		actionMenu = actionMenu.Attr("focused", true)
	}
	if am.PropForcePopover {
		actionMenu = actionMenu.Attr("force-popover", true)
	}
	if am.PropIcons != "" {
		actionMenu = actionMenu.Attr("icons", string(am.PropIcons))
	}
	if am.PropInvalid {
		actionMenu = actionMenu.Attr("invalid", true)
	}
	if am.PropLabel != "" {
		actionMenu = actionMenu.Attr("label", am.PropLabel)
	}
	if am.PropOpen {
		actionMenu = actionMenu.Attr("open", true)
	}
	if am.PropPending {
		actionMenu = actionMenu.Attr("pending", true)
	}
	if am.PropPendingLabel != "" {
		actionMenu = actionMenu.Attr("pending-label", am.PropPendingLabel)
	}
	if am.PropPlacement != "" {
		actionMenu = actionMenu.Attr("placement", am.PropPlacement)
	}
	if am.PropQuiet {
		actionMenu = actionMenu.Attr("quiet", true)
	}
	if am.PropReadonly {
		actionMenu = actionMenu.Attr("readonly", true)
	}
	if am.PropSelects != "" {
		actionMenu = actionMenu.Attr("selects", am.PropSelects)
	}
	if am.PropStaticColor != "" {
		actionMenu = actionMenu.Attr("static-color", am.PropStaticColor)
	}
	if am.PropValue != "" {
		actionMenu = actionMenu.Attr("value", am.PropValue)
	}

	// Add event handlers
	if am.PropOnChange != nil {
		actionMenu = actionMenu.On("change", am.PropOnChange)
	}
	if am.PropOnScroll != nil {
		actionMenu = actionMenu.On("scroll", am.PropOnScroll)
	}
	if am.PropOnOpened != nil {
		actionMenu = actionMenu.On("opened", am.PropOnOpened)
	}

	// Add slots
	slotElements := []app.UI{}

	// Add description slot
	if am.PropDescription != nil {
		descriptionElem := am.PropDescription
		if descWithSlot, ok := descriptionElem.(interface{ Slot(string) app.UI }); ok {
			descriptionElem = descWithSlot.Slot("description")
		} else {
			descriptionElem = app.Elem("div").
				Attr("slot", "description").
				Body(descriptionElem)
		}
		slotElements = append(slotElements, descriptionElem)
	}

	// Add icon slot
	if am.PropIconSlot != nil {
		iconElem := am.PropIconSlot
		if iconWithSlot, ok := iconElem.(interface{ Slot(string) app.UI }); ok {
			iconElem = iconWithSlot.Slot("icon")
		} else {
			iconElem = app.Elem("div").
				Attr("slot", "icon").
				Body(iconElem)
		}
		slotElements = append(slotElements, iconElem)
	}

	// Add label slot
	if am.PropLabelSlot != nil {
		labelElem := am.PropLabelSlot
		if labelWithSlot, ok := labelElem.(interface{ Slot(string) app.UI }); ok {
			labelElem = labelWithSlot.Slot("label")
		} else {
			labelElem = app.Elem("div").
				Attr("slot", "label").
				Body(labelElem)
		}
		slotElements = append(slotElements, labelElem)
	}

	// Add label-only slot
	if am.PropLabelOnlySlot != nil {
		labelOnlyElem := am.PropLabelOnlySlot
		if labelOnlyWithSlot, ok := labelOnlyElem.(interface{ Slot(string) app.UI }); ok {
			labelOnlyElem = labelOnlyWithSlot.Slot("label-only")
		} else {
			labelOnlyElem = app.Elem("div").
				Attr("slot", "label-only").
				Body(labelOnlyElem)
		}
		slotElements = append(slotElements, labelOnlyElem)
	}

	// Add tooltip slot
	if am.PropTooltipSlot != nil {
		tooltipElem := am.PropTooltipSlot
		if tooltipWithSlot, ok := tooltipElem.(interface{ Slot(string) app.UI }); ok {
			tooltipElem = tooltipWithSlot.Slot("tooltip")
		} else {
			tooltipElem = app.Elem("div").
				Attr("slot", "tooltip").
				Body(tooltipElem)
		}
		slotElements = append(slotElements, tooltipElem)
	}

	// Add child elements (menu items)
	if len(am.PropChildren) > 0 {
		for _, child := range am.PropChildren {
			slotElements = append(slotElements, child)
		}
	}

	// Add all elements to the action menu
	if len(slotElements) > 0 {
		actionMenu = actionMenu.Body(slotElements...)
	}

	return actionMenu
}
