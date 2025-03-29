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

// spectrumActionButton represents an sp-action-button component
type spectrumActionButton struct {
	app.Compo

	// Properties
	PropActive         bool
	PropDisabled       bool
	PropDownload       string
	PropEmphasized     bool
	PropHoldAffordance bool
	PropHref           string
	PropLabel          string
	PropQuiet          bool
	PropReferrerpolicy string
	PropRel            string
	PropRole           string
	PropSelected       bool
	PropStaticColor    ActionButtonStaticColor
	PropTabIndex       int
	PropTarget         string
	PropToggles        bool
	PropButtonType     string
	PropValue          string
	PropSize           ActionButtonSize

	// Event handlers
	PropOnClick     app.EventHandler
	PropOnChange    app.EventHandler
	PropOnLongpress app.EventHandler

	// Content
	PropContent   string
	PropInnerHTML string
	PropIcon      app.UI
	PropChild     app.UI
}

// ActionButton creates a new action button component
func ActionButton() *spectrumActionButton {
	return &spectrumActionButton{
		PropRole:       "button",
		PropButtonType: "button",
		PropSize:       ActionButtonSizeM, // Default size is medium
	}
}

// Active sets the active state of the action button
func (a *spectrumActionButton) Active(active bool) *spectrumActionButton {
	a.PropActive = active
	return a
}

// Disabled sets the disabled state of the action button
func (a *spectrumActionButton) Disabled(disabled bool) *spectrumActionButton {
	a.PropDisabled = disabled
	return a
}

// Download sets the download attribute
func (a *spectrumActionButton) Download(download string) *spectrumActionButton {
	a.PropDownload = download
	return a
}

// Emphasized sets the emphasized visual style
func (a *spectrumActionButton) Emphasized(emphasized bool) *spectrumActionButton {
	a.PropEmphasized = emphasized
	return a
}

// HoldAffordance sets whether the button should display a hold affordance
func (a *spectrumActionButton) HoldAffordance(holdAffordance bool) *spectrumActionButton {
	a.PropHoldAffordance = holdAffordance
	return a
}

// Href sets the URL that the button links to
func (a *spectrumActionButton) Href(href string) *spectrumActionButton {
	a.PropHref = href
	return a
}

// Label sets the accessible label for the action button
func (a *spectrumActionButton) Label(label string) *spectrumActionButton {
	a.PropLabel = label
	return a
}

// Quiet sets the quiet visual style
func (a *spectrumActionButton) Quiet(quiet bool) *spectrumActionButton {
	a.PropQuiet = quiet
	return a
}

// Referrerpolicy sets the referrer policy for the link
func (a *spectrumActionButton) Referrerpolicy(referrerpolicy string) *spectrumActionButton {
	a.PropReferrerpolicy = referrerpolicy
	return a
}

// Rel sets the relationship of the linked URL
func (a *spectrumActionButton) Rel(rel string) *spectrumActionButton {
	a.PropRel = rel
	return a
}

// Role sets the ARIA role
func (a *spectrumActionButton) Role(role string) *spectrumActionButton {
	a.PropRole = role
	return a
}

// Selected sets the selected state of the action button
func (a *spectrumActionButton) Selected(selected bool) *spectrumActionButton {
	a.PropSelected = selected
	return a
}

// StaticColor sets the static color variant
func (a *spectrumActionButton) StaticColor(staticColor ActionButtonStaticColor) *spectrumActionButton {
	a.PropStaticColor = staticColor
	return a
}

// TabIndex sets the tab index
func (a *spectrumActionButton) TabIndex(tabIndex int) *spectrumActionButton {
	a.PropTabIndex = tabIndex
	return a
}

// Target sets where to display the linked URL
func (a *spectrumActionButton) Target(target string) *spectrumActionButton {
	a.PropTarget = target
	return a
}

// Toggles sets whether the button toggles its selected state
func (a *spectrumActionButton) Toggles(toggles bool) *spectrumActionButton {
	a.PropToggles = toggles
	return a
}

// Type sets the button type
func (a *spectrumActionButton) Type(buttonType string) *spectrumActionButton {
	a.PropButtonType = buttonType
	return a
}

// Value sets the button value
func (a *spectrumActionButton) Value(value string) *spectrumActionButton {
	a.PropValue = value
	return a
}

// Size sets the visual size of the action button
func (a *spectrumActionButton) Size(size ActionButtonSize) *spectrumActionButton {
	a.PropSize = size
	return a
}

// OnClick sets the click event handler
func (a *spectrumActionButton) OnClick(handler app.EventHandler) *spectrumActionButton {
	a.PropOnClick = handler
	return a
}

// OnChange sets the change event handler
func (a *spectrumActionButton) OnChange(handler app.EventHandler) *spectrumActionButton {
	a.PropOnChange = handler
	return a
}

// OnLongpress sets the longpress event handler
func (a *spectrumActionButton) OnLongpress(handler app.EventHandler) *spectrumActionButton {
	a.PropOnLongpress = handler
	return a
}

// Content sets the text content of the action button
func (a *spectrumActionButton) Content(content string) *spectrumActionButton {
	a.PropContent = content
	return a
}

// InnerHTML sets the inner HTML of the action button
func (a *spectrumActionButton) InnerHTML(html string) *spectrumActionButton {
	a.PropInnerHTML = html
	return a
}

// Icon sets the icon UI element in the icon slot
func (a *spectrumActionButton) Icon(icon app.UI) *spectrumActionButton {
	a.PropIcon = icon
	return a
}

// Child sets a UI element as the child of the action button
func (a *spectrumActionButton) Child(child app.UI) *spectrumActionButton {
	a.PropChild = child
	return a
}

// Render renders the action button component
func (a *spectrumActionButton) Render() app.UI {
	actionButton := app.Elem("sp-action-button").
		Attr("active", a.PropActive).
		Attr("disabled", a.PropDisabled).
		Attr("emphasized", a.PropEmphasized).
		Attr("hold-affordance", a.PropHoldAffordance).
		Attr("quiet", a.PropQuiet).
		Attr("selected", a.PropSelected).
		Attr("toggles", a.PropToggles).
		Attr("type", a.PropButtonType).
		OnClick(a.PropOnClick).
		OnChange(a.PropOnChange).
		Body(
			a.PropIcon,
			a.PropChild,
			app.Text(a.PropContent),
		).
		On("longpress", a.PropOnLongpress)

	if a.PropDownload != "" {
		actionButton = actionButton.Attr("download", a.PropDownload)
	}

	if a.PropHref != "" {
		actionButton = actionButton.Attr("href", a.PropHref)
	}

	if a.PropLabel != "" {
		actionButton = actionButton.Attr("label", a.PropLabel)
	}

	if a.PropReferrerpolicy != "" {
		actionButton = actionButton.Attr("referrerpolicy", a.PropReferrerpolicy)
	}

	if a.PropRel != "" {
		actionButton = actionButton.Attr("rel", a.PropRel)
	}

	if a.PropRole != "" {
		actionButton = actionButton.Attr("role", a.PropRole)
	}

	if a.PropStaticColor != "" {
		actionButton = actionButton.Attr("static-color", string(a.PropStaticColor))
	}

	if a.PropTabIndex != 0 {
		actionButton = actionButton.Attr("tabindex", a.PropTabIndex)
	}

	if a.PropTarget != "" {
		actionButton = actionButton.Attr("target", a.PropTarget)
	}

	if a.PropValue != "" {
		actionButton = actionButton.Attr("value", a.PropValue)
	}

	if a.PropSize != ActionButtonSizeM {
		actionButton = actionButton.Attr("size", string(a.PropSize))
	}

	return actionButton
}
