package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionButtonSize represents the visual size of an action button
type ActionButtonSize string

// Action button sizes
const (
	ActionButtonSizeXS ActionButtonSize = "xs" // Extra Small
	ActionButtonSizeS  ActionButtonSize = "s"  // Small
	ActionButtonSizeM  ActionButtonSize = "m"  // Medium
	ActionButtonSizeL  ActionButtonSize = "l"  // Large
	ActionButtonSizeXL ActionButtonSize = "xl" // Extra Large
)

// ActionButtonStaticColor represents the static color variant for an action button
type ActionButtonStaticColor string

// Action button static colors
const (
	ActionButtonStaticColorWhite ActionButtonStaticColor = "white"
	ActionButtonStaticColorBlack ActionButtonStaticColor = "black"
)

// SpectrumActionButton represents an sp-action-button component
type SpectrumActionButton struct {
	app.Compo

	// Properties
	active         bool
	disabled       bool
	download       string
	emphasized     bool
	holdAffordance bool
	href           string
	label          string
	quiet          bool
	referrerpolicy string
	rel            string
	role           string
	selected       bool
	staticColor    ActionButtonStaticColor
	tabIndex       int
	target         string
	toggles        bool
	buttonType     string
	value          string
	size           ActionButtonSize

	// Event handlers
	onClick     app.EventHandler
	onChange    app.EventHandler
	onLongpress app.EventHandler

	// Content
	content   string
	innerHTML string
	icon      app.UI
	child     app.UI
}

// ActionButton creates a new action button component
func ActionButton() *SpectrumActionButton {
	return &SpectrumActionButton{
		role:       "button",
		buttonType: "button",
		size:       ActionButtonSizeM, // Default size is medium
	}
}

// Active sets the active state of the action button
func (a *SpectrumActionButton) Active(active bool) *SpectrumActionButton {
	a.active = active
	return a
}

// Disabled sets the disabled state of the action button
func (a *SpectrumActionButton) Disabled(disabled bool) *SpectrumActionButton {
	a.disabled = disabled
	return a
}

// Download sets the download attribute
func (a *SpectrumActionButton) Download(download string) *SpectrumActionButton {
	a.download = download
	return a
}

// Emphasized sets the emphasized visual style
func (a *SpectrumActionButton) Emphasized(emphasized bool) *SpectrumActionButton {
	a.emphasized = emphasized
	return a
}

// HoldAffordance sets whether the button should display a hold affordance
func (a *SpectrumActionButton) HoldAffordance(holdAffordance bool) *SpectrumActionButton {
	a.holdAffordance = holdAffordance
	return a
}

// Href sets the URL that the button links to
func (a *SpectrumActionButton) Href(href string) *SpectrumActionButton {
	a.href = href
	return a
}

// Label sets the accessible label for the action button
func (a *SpectrumActionButton) Label(label string) *SpectrumActionButton {
	a.label = label
	return a
}

// Quiet sets the quiet visual style
func (a *SpectrumActionButton) Quiet(quiet bool) *SpectrumActionButton {
	a.quiet = quiet
	return a
}

// Referrerpolicy sets the referrer policy for the link
func (a *SpectrumActionButton) Referrerpolicy(referrerpolicy string) *SpectrumActionButton {
	a.referrerpolicy = referrerpolicy
	return a
}

// Rel sets the relationship of the linked URL
func (a *SpectrumActionButton) Rel(rel string) *SpectrumActionButton {
	a.rel = rel
	return a
}

// Role sets the ARIA role
func (a *SpectrumActionButton) Role(role string) *SpectrumActionButton {
	a.role = role
	return a
}

// Selected sets the selected state of the action button
func (a *SpectrumActionButton) Selected(selected bool) *SpectrumActionButton {
	a.selected = selected
	return a
}

// StaticColor sets the static color variant
func (a *SpectrumActionButton) StaticColor(staticColor ActionButtonStaticColor) *SpectrumActionButton {
	a.staticColor = staticColor
	return a
}

// TabIndex sets the tab index
func (a *SpectrumActionButton) TabIndex(tabIndex int) *SpectrumActionButton {
	a.tabIndex = tabIndex
	return a
}

// Target sets where to display the linked URL
func (a *SpectrumActionButton) Target(target string) *SpectrumActionButton {
	a.target = target
	return a
}

// Toggles sets whether the button toggles its selected state
func (a *SpectrumActionButton) Toggles(toggles bool) *SpectrumActionButton {
	a.toggles = toggles
	return a
}

// Type sets the button type
func (a *SpectrumActionButton) Type(buttonType string) *SpectrumActionButton {
	a.buttonType = buttonType
	return a
}

// Value sets the button value
func (a *SpectrumActionButton) Value(value string) *SpectrumActionButton {
	a.value = value
	return a
}

// Size sets the visual size of the action button
func (a *SpectrumActionButton) Size(size ActionButtonSize) *SpectrumActionButton {
	a.size = size
	return a
}

// OnClick sets the click event handler
func (a *SpectrumActionButton) OnClick(handler app.EventHandler) *SpectrumActionButton {
	a.onClick = handler
	return a
}

// OnChange sets the change event handler
func (a *SpectrumActionButton) OnChange(handler app.EventHandler) *SpectrumActionButton {
	a.onChange = handler
	return a
}

// OnLongpress sets the longpress event handler
func (a *SpectrumActionButton) OnLongpress(handler app.EventHandler) *SpectrumActionButton {
	a.onLongpress = handler
	return a
}

// Content sets the text content of the action button
func (a *SpectrumActionButton) Content(content string) *SpectrumActionButton {
	a.content = content
	return a
}

// InnerHTML sets the inner HTML of the action button
func (a *SpectrumActionButton) InnerHTML(html string) *SpectrumActionButton {
	a.innerHTML = html
	return a
}

// Icon sets the icon UI element in the icon slot
func (a *SpectrumActionButton) Icon(icon app.UI) *SpectrumActionButton {
	a.icon = icon
	return a
}

// Child sets a UI element as the child of the action button
func (a *SpectrumActionButton) Child(child app.UI) *SpectrumActionButton {
	a.child = child
	return a
}

// Slot implements the Slot interface for ActionButton
func (a *SpectrumActionButton) Slot(slot string) app.UI {
	// Create a duplicate of the component's configuration for slot use
	actionButton := app.Elem("sp-action-button").
		Attr("slot", slot).
		Attr("active", a.active).
		Attr("disabled", a.disabled).
		Attr("emphasized", a.emphasized).
		Attr("hold-affordance", a.holdAffordance).
		Attr("quiet", a.quiet).
		Attr("selected", a.selected).
		Attr("toggles", a.toggles).
		Attr("type", a.buttonType)

	// Only add non-default attributes
	if a.download != "" {
		actionButton = actionButton.Attr("download", a.download)
	}
	if a.href != "" {
		actionButton = actionButton.Attr("href", a.href)
	}
	if a.label != "" {
		actionButton = actionButton.Attr("label", a.label)
	}
	if a.referrerpolicy != "" {
		actionButton = actionButton.Attr("referrerpolicy", a.referrerpolicy)
	}
	if a.rel != "" {
		actionButton = actionButton.Attr("rel", a.rel)
	}
	if a.role != "" && a.role != "button" {
		actionButton = actionButton.Attr("role", a.role)
	}
	if a.staticColor != "" {
		actionButton = actionButton.Attr("static-color", string(a.staticColor))
	}
	if a.tabIndex != 0 {
		actionButton = actionButton.Attr("tabindex", a.tabIndex)
	}
	if a.target != "" {
		actionButton = actionButton.Attr("target", a.target)
	}
	if a.value != "" {
		actionButton = actionButton.Attr("value", a.value)
	}
	if a.size != "" && a.size != ActionButtonSizeM {
		actionButton = actionButton.Attr("size", string(a.size))
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add icon if provided
	if a.icon != nil {
		iconElem := a.icon
		if iconWithSlot, ok := iconElem.(interface{ Slot(string) app.UI }); ok {
			iconElem = iconWithSlot.Slot("icon")
		} else {
			iconElem = app.Elem("div").
				Attr("slot", "icon").
				Body(iconElem)
		}
		elements = append(elements, iconElem)
	}

	// Add content or child
	if a.child != nil {
		elements = append(elements, a.child)
	} else if a.innerHTML != "" {
		elements = append(elements, app.Raw(a.innerHTML))
	} else if a.content != "" {
		elements = append(elements, app.Text(a.content))
	}

	// Add all elements to the action button
	if len(elements) > 0 {
		actionButton = actionButton.Body(elements...)
	}

	return actionButton
}

// Render renders the action button component
func (a *SpectrumActionButton) Render() app.UI {
	actionButton := app.Elem("sp-action-button").
		Attr("active", a.active).
		Attr("disabled", a.disabled).
		Attr("emphasized", a.emphasized).
		Attr("hold-affordance", a.holdAffordance).
		Attr("quiet", a.quiet).
		Attr("selected", a.selected).
		Attr("toggles", a.toggles).
		Attr("type", a.buttonType)

	if a.download != "" {
		actionButton = actionButton.Attr("download", a.download)
	}
	if a.href != "" {
		actionButton = actionButton.Attr("href", a.href)
	}
	if a.label != "" {
		actionButton = actionButton.Attr("label", a.label)
	}
	if a.referrerpolicy != "" {
		actionButton = actionButton.Attr("referrerpolicy", a.referrerpolicy)
	}
	if a.rel != "" {
		actionButton = actionButton.Attr("rel", a.rel)
	}
	if a.role != "" {
		actionButton = actionButton.Attr("role", a.role)
	}
	if a.staticColor != "" {
		actionButton = actionButton.Attr("static-color", string(a.staticColor))
	}
	if a.tabIndex != 0 {
		actionButton = actionButton.Attr("tabindex", a.tabIndex)
	}
	if a.target != "" {
		actionButton = actionButton.Attr("target", a.target)
	}
	if a.value != "" {
		actionButton = actionButton.Attr("value", a.value)
	}
	if a.size != "" {
		actionButton = actionButton.Attr("size", string(a.size))
	}

	// Add event handlers
	if a.onClick != nil {
		actionButton = actionButton.OnClick(a.onClick)
	}
	if a.onChange != nil {
		actionButton = actionButton.On("change", a.onChange)
	}
	if a.onLongpress != nil {
		actionButton = actionButton.On("longpress", a.onLongpress)
	}

	// Create elements array for children
	elements := []app.UI{}

	// Add icon if provided
	if a.icon != nil {
		iconElem := a.icon
		if iconWithSlot, ok := iconElem.(interface{ Slot(string) app.UI }); ok {
			iconElem = iconWithSlot.Slot("icon")
		} else {
			iconElem = app.Elem("div").
				Attr("slot", "icon").
				Body(iconElem)
		}
		elements = append(elements, iconElem)
	}

	// Add content or child
	if a.child != nil {
		elements = append(elements, a.child)
	} else if a.innerHTML != "" {
		elements = append(elements, app.Raw(a.innerHTML))
	} else if a.content != "" {
		elements = append(elements, app.Text(a.content))
	}

	// Add all elements to the action button
	if len(elements) > 0 {
		actionButton = actionButton.Body(elements...)
	}

	return actionButton
}
