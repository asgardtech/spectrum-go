package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// PickerIcons represents the icon display modes for pickers
type PickerIcons string

const (
	PickerIconsOnly PickerIcons = "only"
	PickerIconsNone PickerIcons = "none"
)

// spectrumPicker represents an sp-picker component
type spectrumPicker struct {
	app.Compo

	// Properties
	PropDisabled     bool
	PropFocused      bool
	PropForcePopover bool
	PropIcons        PickerIcons
	PropInvalid      bool
	PropLabel        string
	PropOpen         bool
	PropPending      bool
	PropPendingLabel string
	PropPlacement    OverlayPlacement
	PropQuiet        bool
	PropReadonly     bool
	PropValue        string

	// Slots
	PropDescriptionSlot app.UI
	PropLabelSlot       app.UI
	PropTooltipSlot     app.UI
	PropMenuItems       []app.UI
	PropMenuGroups      []app.UI
	PropMenuDividers    []app.UI
	PropCustomItems     []app.UI

	// Event handlers
	PropOnChange app.EventHandler
	PropOnScroll app.EventHandler
	PropOnOpened app.EventHandler
}

// Picker creates a new picker component
func Picker() *spectrumPicker {
	return &spectrumPicker{
		PropPlacement:    OverlayPlacementBottomStart, // Default placement
		PropPendingLabel: "Pending",                   // Default pending label
	}
}

// Disabled sets whether the picker is disabled
func (p *spectrumPicker) Disabled(disabled bool) *spectrumPicker {
	p.PropDisabled = disabled
	return p
}

// Focused sets whether the picker has keyboard focus
func (p *spectrumPicker) Focused(focused bool) *spectrumPicker {
	p.PropFocused = focused
	return p
}

// ForcePopover forces the Picker to render as a popover on mobile instead of a tray
func (p *spectrumPicker) ForcePopover(forcePopover bool) *spectrumPicker {
	p.PropForcePopover = forcePopover
	return p
}

// Icons sets the icon display mode for the picker
func (p *spectrumPicker) Icons(icons PickerIcons) *spectrumPicker {
	p.PropIcons = icons
	return p
}

// Invalid sets whether the picker is in an invalid state
func (p *spectrumPicker) Invalid(invalid bool) *spectrumPicker {
	p.PropInvalid = invalid
	return p
}

// Label sets the visible label text for the picker
func (p *spectrumPicker) Label(label string) *spectrumPicker {
	p.PropLabel = label
	return p
}

// Open sets whether the picker is open
func (p *spectrumPicker) Open(open bool) *spectrumPicker {
	p.PropOpen = open
	return p
}

// Pending sets whether the picker is in a pending/loading state
func (p *spectrumPicker) Pending(pending bool) *spectrumPicker {
	p.PropPending = pending
	return p
}

// PendingLabel sets the assistive text for the pending state
func (p *spectrumPicker) PendingLabel(pendingLabel string) *spectrumPicker {
	p.PropPendingLabel = pendingLabel
	return p
}

// Placement sets the placement of the picker overlay
func (p *spectrumPicker) Placement(placement OverlayPlacement) *spectrumPicker {
	p.PropPlacement = placement
	return p
}

// Quiet sets whether the picker has a quiet appearance
func (p *spectrumPicker) Quiet(quiet bool) *spectrumPicker {
	p.PropQuiet = quiet
	return p
}

// Readonly sets whether the picker is readonly
func (p *spectrumPicker) Readonly(readonly bool) *spectrumPicker {
	p.PropReadonly = readonly
	return p
}

// Value sets the value of the selected item
func (p *spectrumPicker) Value(value string) *spectrumPicker {
	p.PropValue = value
	return p
}

// Description sets the description content for the picker
func (p *spectrumPicker) Description(description app.UI) *spectrumPicker {
	p.PropDescriptionSlot = description
	return p
}

// LabelSlot sets content for the label slot
func (p *spectrumPicker) LabelSlot(label app.UI) *spectrumPicker {
	p.PropLabelSlot = label
	return p
}

// Tooltip sets content for the tooltip slot
func (p *spectrumPicker) Tooltip(tooltip app.UI) *spectrumPicker {
	p.PropTooltipSlot = tooltip
	return p
}

// Items sets the menu items for the picker
func (p *spectrumPicker) Items(items ...*spectrumMenuItem) *spectrumPicker {
	for _, item := range items {
		p.PropMenuItems = append(p.PropMenuItems, item)
	}
	return p
}

// Groups sets the menu groups for the picker
func (p *spectrumPicker) Groups(groups ...app.UI) *spectrumPicker {
	for _, group := range groups {
		p.PropMenuGroups = append(p.PropMenuGroups, group)
	}
	return p
}

// Divider adds a menu divider to the picker
func (p *spectrumPicker) Divider() *spectrumPicker {
	p.PropMenuDividers = append(p.PropMenuDividers, MenuDivider())
	return p
}

// DividerWithSize adds a menu divider with a specific size to the picker
func (p *spectrumPicker) DividerWithSize(size string) *spectrumPicker {
	p.PropMenuDividers = append(p.PropMenuDividers, MenuDivider().Size(size))
	return p
}

// CustomItem adds a custom UI element to the picker
func (p *spectrumPicker) CustomItem(item app.UI) *spectrumPicker {
	p.PropCustomItems = append(p.PropCustomItems, item)
	return p
}

// OnChange sets the change event handler
func (p *spectrumPicker) OnChange(handler app.EventHandler) *spectrumPicker {
	p.PropOnChange = handler
	return p
}

// OnScroll sets the scroll event handler
func (p *spectrumPicker) OnScroll(handler app.EventHandler) *spectrumPicker {
	p.PropOnScroll = handler
	return p
}

// OnOpened sets the opened event handler
func (p *spectrumPicker) OnOpened(handler app.EventHandler) *spectrumPicker {
	p.PropOnOpened = handler
	return p
}

// IconsOnly sets icons display mode to "only"
func (p *spectrumPicker) IconsOnly() *spectrumPicker {
	return p.Icons(PickerIconsOnly)
}

// IconsNone sets icons display mode to "none"
func (p *spectrumPicker) IconsNone() *spectrumPicker {
	return p.Icons(PickerIconsNone)
}

// Render renders the picker component
func (p *spectrumPicker) Render() app.UI {
	picker := app.Elem("sp-picker")

	// Set attributes
	if p.PropDisabled {
		picker = picker.Attr("disabled", true)
	}
	if p.PropFocused {
		picker = picker.Attr("focused", true)
	}
	if p.PropForcePopover {
		picker = picker.Attr("force-popover", true)
	}
	if p.PropIcons != "" {
		picker = picker.Attr("icons", string(p.PropIcons))
	}
	if p.PropInvalid {
		picker = picker.Attr("invalid", true)
	}
	if p.PropLabel != "" {
		picker = picker.Attr("label", p.PropLabel)
	}
	if p.PropOpen {
		picker = picker.Attr("open", true)
	}
	if p.PropPending {
		picker = picker.Attr("pending", true)
	}
	if p.PropPendingLabel != "Pending" { // Only set if not the default
		picker = picker.Attr("pending-label", p.PropPendingLabel)
	}
	if p.PropPlacement != "" {
		picker = picker.Attr("placement", string(p.PropPlacement))
	}
	if p.PropQuiet {
		picker = picker.Attr("quiet", true)
	}
	if p.PropReadonly {
		picker = picker.Attr("readonly", true)
	}
	if p.PropValue != "" {
		picker = picker.Attr("value", p.PropValue)
	}

	// Add event handlers
	if p.PropOnChange != nil {
		picker = picker.On("change", p.PropOnChange)
	}
	if p.PropOnScroll != nil {
		picker = picker.On("scroll", p.PropOnScroll)
	}
	if p.PropOnOpened != nil {
		picker = picker.On("sp-opened", p.PropOnOpened)
	}

	// Collect all slotted elements
	var elements []app.UI

	// Add description slot if provided
	if p.PropDescriptionSlot != nil {
		desc := p.PropDescriptionSlot
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
	if p.PropLabelSlot != nil {
		label := p.PropLabelSlot
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
	if p.PropTooltipSlot != nil {
		tooltip := p.PropTooltipSlot
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
	for _, item := range p.PropMenuItems {
		elements = append(elements, item)
	}

	// Add menu groups
	for _, group := range p.PropMenuGroups {
		elements = append(elements, group)
	}

	// Add menu dividers
	for _, divider := range p.PropMenuDividers {
		elements = append(elements, divider)
	}

	// Add custom items
	for _, item := range p.PropCustomItems {
		elements = append(elements, item)
	}

	// Add all elements to the picker
	if len(elements) > 0 {
		picker = picker.Body(elements...)
	}

	return picker
}
