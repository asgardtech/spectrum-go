package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// IconsType represents the icon display mode for an action menu
type IconsType string

const (
	IconsTypeNone IconsType = "none"
	IconsTypeOnly IconsType = "only"
)

// SpectrumActionMenu represents an sp-action-menu component
type SpectrumActionMenu struct {
	app.Compo

	// Properties
	disabled     bool
	focused      bool
	forcePopover bool
	icons        IconsType
	invalid      bool
	label        string
	open         bool
	pending      bool
	pendingLabel string
	placement    string
	quiet        bool
	readonly     bool
	selects      string
	staticColor  string
	value        string

	// Content slots
	description   app.UI
	iconSlot      app.UI
	labelSlot     app.UI
	labelOnlySlot app.UI
	tooltipSlot   app.UI

	// Child components (menu items)
	children []app.UI

	// Event handlers
	onChange app.EventHandler
	onScroll app.EventHandler
	onOpened app.EventHandler
}

// ActionMenu creates a new action menu component
func ActionMenu() *SpectrumActionMenu {
	return &SpectrumActionMenu{
		placement:    "bottom-start", // Default placement
		pendingLabel: "Pending",      // Default pending label
	}
}

// Disabled sets whether the action menu is disabled
func (am *SpectrumActionMenu) Disabled(disabled bool) *SpectrumActionMenu {
	am.disabled = disabled
	return am
}

// Focused sets whether the action menu is focused
func (am *SpectrumActionMenu) Focused(focused bool) *SpectrumActionMenu {
	am.focused = focused
	return am
}

// ForcePopover forces the action menu to render as a popover on mobile instead of a tray
func (am *SpectrumActionMenu) ForcePopover(force bool) *SpectrumActionMenu {
	am.forcePopover = force
	return am
}

// Icons sets the icon display mode (none, only)
func (am *SpectrumActionMenu) Icons(icons IconsType) *SpectrumActionMenu {
	am.icons = icons
	return am
}

// IconsNone sets the icon display mode to none
func (am *SpectrumActionMenu) IconsNone() *SpectrumActionMenu {
	return am.Icons(IconsTypeNone)
}

// IconsOnly sets the icon display mode to only
func (am *SpectrumActionMenu) IconsOnly() *SpectrumActionMenu {
	return am.Icons(IconsTypeOnly)
}

// Invalid sets whether the action menu is invalid
func (am *SpectrumActionMenu) Invalid(invalid bool) *SpectrumActionMenu {
	am.invalid = invalid
	return am
}

// Label sets the accessible label for the action menu
func (am *SpectrumActionMenu) Label(label string) *SpectrumActionMenu {
	am.label = label
	return am
}

// Open sets whether the action menu is open
func (am *SpectrumActionMenu) Open(open bool) *SpectrumActionMenu {
	am.open = open
	return am
}

// Pending sets whether the items are currently loading
func (am *SpectrumActionMenu) Pending(pending bool) *SpectrumActionMenu {
	am.pending = pending
	return am
}

// PendingLabel sets the label shown while in pending state
func (am *SpectrumActionMenu) PendingLabel(label string) *SpectrumActionMenu {
	am.pendingLabel = label
	return am
}

// Placement sets the placement of the overlay
func (am *SpectrumActionMenu) Placement(placement string) *SpectrumActionMenu {
	am.placement = placement
	return am
}

// Quiet sets whether the action menu has a quiet appearance
func (am *SpectrumActionMenu) Quiet(quiet bool) *SpectrumActionMenu {
	am.quiet = quiet
	return am
}

// Readonly sets whether the action menu is readonly
func (am *SpectrumActionMenu) Readonly(readonly bool) *SpectrumActionMenu {
	am.readonly = readonly
	return am
}

// Selects sets the selection mode (single or undefined)
func (am *SpectrumActionMenu) Selects(selects string) *SpectrumActionMenu {
	am.selects = selects
	return am
}

// SelectsSingle sets the selection mode to single
func (am *SpectrumActionMenu) SelectsSingle() *SpectrumActionMenu {
	return am.Selects("single")
}

// StaticColor sets the static color of the action menu
func (am *SpectrumActionMenu) StaticColor(color string) *SpectrumActionMenu {
	am.staticColor = color
	return am
}

// StaticColorWhite sets the static color to white
func (am *SpectrumActionMenu) StaticColorWhite() *SpectrumActionMenu {
	return am.StaticColor("white")
}

// StaticColorBlack sets the static color to black
func (am *SpectrumActionMenu) StaticColorBlack() *SpectrumActionMenu {
	return am.StaticColor("black")
}

// Value sets the value of the action menu
func (am *SpectrumActionMenu) Value(value string) *SpectrumActionMenu {
	am.value = value
	return am
}

// Description sets the description content in the description slot
func (am *SpectrumActionMenu) Description(description app.UI) *SpectrumActionMenu {
	am.description = description
	return am
}

// Icon sets the icon in the icon slot
func (am *SpectrumActionMenu) Icon(icon app.UI) *SpectrumActionMenu {
	am.iconSlot = icon
	return am
}

// LabelSlot sets the content in the label slot
func (am *SpectrumActionMenu) LabelSlot(label app.UI) *SpectrumActionMenu {
	am.labelSlot = label
	return am
}

// LabelOnly sets the content in the label-only slot (no icon space reserved)
func (am *SpectrumActionMenu) LabelOnly(label app.UI) *SpectrumActionMenu {
	am.labelOnlySlot = label
	return am
}

// Tooltip sets the tooltip to be applied to the action button
func (am *SpectrumActionMenu) Tooltip(tooltip app.UI) *SpectrumActionMenu {
	am.tooltipSlot = tooltip
	return am
}

// Children sets the child components (menu items)
func (am *SpectrumActionMenu) Children(children ...app.UI) *SpectrumActionMenu {
	am.children = children
	return am
}

// AddItem adds a menu item to the action menu
func (am *SpectrumActionMenu) AddItem(item app.UI) *SpectrumActionMenu {
	am.children = append(am.children, item)
	return am
}

// OnChange sets the change event handler
func (am *SpectrumActionMenu) OnChange(handler app.EventHandler) *SpectrumActionMenu {
	am.onChange = handler
	return am
}

// OnScroll sets the scroll event handler
func (am *SpectrumActionMenu) OnScroll(handler app.EventHandler) *SpectrumActionMenu {
	am.onScroll = handler
	return am
}

// OnOpened sets the opened event handler
func (am *SpectrumActionMenu) OnOpened(handler app.EventHandler) *SpectrumActionMenu {
	am.onOpened = handler
	return am
}

// Render renders the action menu component
func (am *SpectrumActionMenu) Render() app.UI {
	actionMenu := app.Elem("sp-action-menu")

	// Set attributes
	if am.disabled {
		actionMenu = actionMenu.Attr("disabled", true)
	}
	if am.focused {
		actionMenu = actionMenu.Attr("focused", true)
	}
	if am.forcePopover {
		actionMenu = actionMenu.Attr("force-popover", true)
	}
	if am.icons != "" {
		actionMenu = actionMenu.Attr("icons", string(am.icons))
	}
	if am.invalid {
		actionMenu = actionMenu.Attr("invalid", true)
	}
	if am.label != "" {
		actionMenu = actionMenu.Attr("label", am.label)
	}
	if am.open {
		actionMenu = actionMenu.Attr("open", true)
	}
	if am.pending {
		actionMenu = actionMenu.Attr("pending", true)
	}
	if am.pendingLabel != "Pending" {
		actionMenu = actionMenu.Attr("pending-label", am.pendingLabel)
	}
	if am.placement != "bottom-start" {
		actionMenu = actionMenu.Attr("placement", am.placement)
	}
	if am.quiet {
		actionMenu = actionMenu.Attr("quiet", true)
	}
	if am.readonly {
		actionMenu = actionMenu.Attr("readonly", true)
	}
	if am.selects != "" {
		actionMenu = actionMenu.Attr("selects", am.selects)
	}
	if am.staticColor != "" {
		actionMenu = actionMenu.Attr("static-color", am.staticColor)
	}
	if am.value != "" {
		actionMenu = actionMenu.Attr("value", am.value)
	}

	// Add event handlers
	if am.onChange != nil {
		actionMenu = actionMenu.On("change", am.onChange)
	}
	if am.onScroll != nil {
		actionMenu = actionMenu.On("scroll", am.onScroll)
	}
	if am.onOpened != nil {
		actionMenu = actionMenu.On("sp-opened", am.onOpened)
	}

	// Add content slots
	elements := []app.UI{}

	if am.description != nil {
		if descWithSlot, ok := am.description.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, descWithSlot.Slot("description"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "description").
				Body(am.description))
		}
	}

	if am.iconSlot != nil {
		if iconWithSlot, ok := am.iconSlot.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, iconWithSlot.Slot("icon"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "icon").
				Body(am.iconSlot))
		}
	}

	if am.labelSlot != nil {
		if labelWithSlot, ok := am.labelSlot.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, labelWithSlot.Slot("label"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "label").
				Body(am.labelSlot))
		}
	}

	if am.labelOnlySlot != nil {
		if labelOnlyWithSlot, ok := am.labelOnlySlot.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, labelOnlyWithSlot.Slot("label-only"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "label-only").
				Body(am.labelOnlySlot))
		}
	}

	if am.tooltipSlot != nil {
		if tooltipWithSlot, ok := am.tooltipSlot.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, tooltipWithSlot.Slot("tooltip"))
		} else {
			elements = append(elements, app.Elem("div").
				Attr("slot", "tooltip").
				Body(am.tooltipSlot))
		}
	}

	// Add children
	if len(am.children) > 0 {
		elements = append(elements, am.children...)
	}

	if len(elements) > 0 {
		actionMenu = actionMenu.Body(elements...)
	}

	return actionMenu
}
