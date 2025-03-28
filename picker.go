package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerIcons represents the icon display modes for pickers
type PickerIcons string

const (
	PickerIconsOnly PickerIcons = "only"
	PickerIconsNone PickerIcons = "none"
)

// SpectrumPicker represents an sp-picker component
type SpectrumPicker struct {
	app.Compo

	// Properties
	disabled     bool
	focused      bool
	forcePopover bool
	icons        PickerIcons
	invalid      bool
	label        string
	open         bool
	pending      bool
	pendingLabel string
	placement    OverlayPlacement
	quiet        bool
	readonly     bool
	value        string

	// Slots
	descriptionSlot app.UI
	labelSlot       app.UI
	tooltipSlot     app.UI
	menuItems       []*SpectrumMenuItem
	menuGroups      []*SpectrumMenuGroup
	menuDividers    []*SpectrumMenuDivider
	customItems     []app.UI

	// Event handlers
	onChange app.EventHandler
	onScroll app.EventHandler
	onOpened app.EventHandler
}

// Picker creates a new picker component
func Picker() *SpectrumPicker {
	return &SpectrumPicker{
		placement:    OverlayPlacementBottomStart, // Default placement
		pendingLabel: "Pending",                   // Default pending label
	}
}

// Disabled sets whether the picker is disabled
func (p *SpectrumPicker) Disabled(disabled bool) *SpectrumPicker {
	p.disabled = disabled
	return p
}

// Focused sets whether the picker has keyboard focus
func (p *SpectrumPicker) Focused(focused bool) *SpectrumPicker {
	p.focused = focused
	return p
}

// ForcePopover forces the Picker to render as a popover on mobile instead of a tray
func (p *SpectrumPicker) ForcePopover(forcePopover bool) *SpectrumPicker {
	p.forcePopover = forcePopover
	return p
}

// Icons sets the icon display mode for the picker
func (p *SpectrumPicker) Icons(icons PickerIcons) *SpectrumPicker {
	p.icons = icons
	return p
}

// Invalid sets whether the picker is in an invalid state
func (p *SpectrumPicker) Invalid(invalid bool) *SpectrumPicker {
	p.invalid = invalid
	return p
}

// Label sets the visible label text for the picker
func (p *SpectrumPicker) Label(label string) *SpectrumPicker {
	p.label = label
	return p
}

// Open sets whether the picker is open
func (p *SpectrumPicker) Open(open bool) *SpectrumPicker {
	p.open = open
	return p
}

// Pending sets whether the picker is in a pending/loading state
func (p *SpectrumPicker) Pending(pending bool) *SpectrumPicker {
	p.pending = pending
	return p
}

// PendingLabel sets the assistive text for the pending state
func (p *SpectrumPicker) PendingLabel(pendingLabel string) *SpectrumPicker {
	p.pendingLabel = pendingLabel
	return p
}

// Placement sets the placement of the picker overlay
func (p *SpectrumPicker) Placement(placement OverlayPlacement) *SpectrumPicker {
	p.placement = placement
	return p
}

// Quiet sets whether the picker has a quiet appearance
func (p *SpectrumPicker) Quiet(quiet bool) *SpectrumPicker {
	p.quiet = quiet
	return p
}

// Readonly sets whether the picker is readonly
func (p *SpectrumPicker) Readonly(readonly bool) *SpectrumPicker {
	p.readonly = readonly
	return p
}

// Value sets the value of the selected item
func (p *SpectrumPicker) Value(value string) *SpectrumPicker {
	p.value = value
	return p
}

// Description sets the description content for the picker
func (p *SpectrumPicker) Description(description app.UI) *SpectrumPicker {
	p.descriptionSlot = description
	return p
}

// LabelSlot sets content for the label slot
func (p *SpectrumPicker) LabelSlot(label app.UI) *SpectrumPicker {
	p.labelSlot = label
	return p
}

// Tooltip sets content for the tooltip slot
func (p *SpectrumPicker) Tooltip(tooltip app.UI) *SpectrumPicker {
	p.tooltipSlot = tooltip
	return p
}

// Items sets the menu items for the picker
func (p *SpectrumPicker) Items(items ...*SpectrumMenuItem) *SpectrumPicker {
	p.menuItems = items
	return p
}

// Groups sets the menu groups for the picker
func (p *SpectrumPicker) Groups(groups ...*SpectrumMenuGroup) *SpectrumPicker {
	p.menuGroups = groups
	return p
}

// Divider adds a menu divider to the picker
func (p *SpectrumPicker) Divider() *SpectrumPicker {
	p.menuDividers = append(p.menuDividers, MenuDivider())
	return p
}

// DividerWithSize adds a menu divider with a specific size to the picker
func (p *SpectrumPicker) DividerWithSize(size string) *SpectrumPicker {
	p.menuDividers = append(p.menuDividers, MenuDivider().Size(size))
	return p
}

// CustomItem adds a custom UI element to the picker
func (p *SpectrumPicker) CustomItem(item app.UI) *SpectrumPicker {
	p.customItems = append(p.customItems, item)
	return p
}

// OnChange sets the change event handler
func (p *SpectrumPicker) OnChange(handler app.EventHandler) *SpectrumPicker {
	p.onChange = handler
	return p
}

// OnScroll sets the scroll event handler
func (p *SpectrumPicker) OnScroll(handler app.EventHandler) *SpectrumPicker {
	p.onScroll = handler
	return p
}

// OnOpened sets the opened event handler
func (p *SpectrumPicker) OnOpened(handler app.EventHandler) *SpectrumPicker {
	p.onOpened = handler
	return p
}

// IconsOnly sets icons display mode to "only"
func (p *SpectrumPicker) IconsOnly() *SpectrumPicker {
	return p.Icons(PickerIconsOnly)
}

// IconsNone sets icons display mode to "none"
func (p *SpectrumPicker) IconsNone() *SpectrumPicker {
	return p.Icons(PickerIconsNone)
}

// Render renders the picker component
func (p *SpectrumPicker) Render() app.UI {
	picker := app.Elem("sp-picker")

	// Set attributes
	if p.disabled {
		picker = picker.Attr("disabled", true)
	}
	if p.focused {
		picker = picker.Attr("focused", true)
	}
	if p.forcePopover {
		picker = picker.Attr("force-popover", true)
	}
	if p.icons != "" {
		picker = picker.Attr("icons", string(p.icons))
	}
	if p.invalid {
		picker = picker.Attr("invalid", true)
	}
	if p.label != "" {
		picker = picker.Attr("label", p.label)
	}
	if p.open {
		picker = picker.Attr("open", true)
	}
	if p.pending {
		picker = picker.Attr("pending", true)
	}
	if p.pendingLabel != "Pending" { // Only set if not the default
		picker = picker.Attr("pending-label", p.pendingLabel)
	}
	if p.placement != "" {
		picker = picker.Attr("placement", string(p.placement))
	}
	if p.quiet {
		picker = picker.Attr("quiet", true)
	}
	if p.readonly {
		picker = picker.Attr("readonly", true)
	}
	if p.value != "" {
		picker = picker.Attr("value", p.value)
	}

	// Add event handlers
	if p.onChange != nil {
		picker = picker.On("change", p.onChange)
	}
	if p.onScroll != nil {
		picker = picker.On("scroll", p.onScroll)
	}
	if p.onOpened != nil {
		picker = picker.On("sp-opened", p.onOpened)
	}

	// Collect all slotted elements
	var elements []app.UI

	// Add description slot if provided
	if p.descriptionSlot != nil {
		desc := p.descriptionSlot
		if descWithSlot, ok := desc.(interface{ Slot(string) app.UI }); ok {
			desc = descWithSlot.Slot("description")
		} else {
			desc = app.Elem("div").
				Attr("slot", "description").
				Body(desc)
		}
		elements = append(elements, desc)
	}

	// Add label slot if provided
	if p.labelSlot != nil {
		label := p.labelSlot
		if labelWithSlot, ok := label.(interface{ Slot(string) app.UI }); ok {
			label = labelWithSlot.Slot("label")
		} else {
			label = app.Elem("div").
				Attr("slot", "label").
				Body(label)
		}
		elements = append(elements, label)
	}

	// Add tooltip slot if provided
	if p.tooltipSlot != nil {
		tooltip := p.tooltipSlot
		if tooltipWithSlot, ok := tooltip.(interface{ Slot(string) app.UI }); ok {
			tooltip = tooltipWithSlot.Slot("tooltip")
		} else {
			tooltip = app.Elem("div").
				Attr("slot", "tooltip").
				Body(tooltip)
		}
		elements = append(elements, tooltip)
	}

	// Add menu items
	for _, item := range p.menuItems {
		elements = append(elements, item)
	}

	// Add menu groups
	for _, group := range p.menuGroups {
		elements = append(elements, group)
	}

	// Add menu dividers
	for _, divider := range p.menuDividers {
		elements = append(elements, divider)
	}

	// Add custom items
	for _, item := range p.customItems {
		elements = append(elements, item)
	}

	// Add all elements to the picker
	if len(elements) > 0 {
		picker = picker.Body(elements...)
	}

	return picker
}
