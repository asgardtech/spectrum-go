package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumFocusable represents a focusable component
type spectrumFocusable struct {
	app.Compo

	// Properties
	PropTabIndex int
	PropDisabled bool
	PropOnFocus  app.EventHandler
	PropOnBlur   app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewFocusable creates a new focusable component
func NewFocusable() *spectrumFocusable {
	return &spectrumFocusable{}
}

// TabIndex sets the tab index
func (f *spectrumFocusable) TabIndex(index int) *spectrumFocusable {
	f.PropTabIndex = index
	return f
}

// Disabled sets the disabled state
func (f *spectrumFocusable) Disabled(disabled bool) *spectrumFocusable {
	f.PropDisabled = disabled
	return f
}

// OnFocus sets the focus event handler
func (f *spectrumFocusable) OnFocus(handler app.EventHandler) *spectrumFocusable {
	f.PropOnFocus = handler
	return f
}

// OnBlur sets the blur event handler
func (f *spectrumFocusable) OnBlur(handler app.EventHandler) *spectrumFocusable {
	f.PropOnBlur = handler
	return f
}

// Child adds a child element
func (f *spectrumFocusable) Child(child app.UI) *spectrumFocusable {
	f.PropChildren = append(f.PropChildren, child)
	return f
}

// Children adds multiple child elements
func (f *spectrumFocusable) Children(children ...app.UI) *spectrumFocusable {
	f.PropChildren = append(f.PropChildren, children...)
	return f
}

// Render renders the focusable component
func (f *spectrumFocusable) Render() app.UI {
	focusable := app.Elem("div")

	// Set attributes
	if f.PropTabIndex != 0 {
		focusable = focusable.Attr("tabindex", f.PropTabIndex)
	}
	if f.PropDisabled {
		focusable = focusable.Attr("disabled", true)
	}

	// Add event handlers
	if f.PropOnFocus != nil {
		focusable = focusable.On("focus", f.PropOnFocus)
	}
	if f.PropOnBlur != nil {
		focusable = focusable.On("blur", f.PropOnBlur)
	}

	// Add children if provided
	if len(f.PropChildren) > 0 {
		focusable = focusable.Body(f.PropChildren...)
	}

	return focusable
}

// spectrumLikeAnchor represents a component that behaves like an anchor element
type spectrumLikeAnchor struct {
	app.Compo

	// Properties
	PropHref     string
	PropTarget   string
	PropRel      string
	PropDownload string
	PropLabel    string
	PropOnClick  app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewLikeAnchor creates a new anchor-like component
func NewLikeAnchor() *spectrumLikeAnchor {
	return &spectrumLikeAnchor{}
}

// Href sets the href attribute
func (l *spectrumLikeAnchor) Href(href string) *spectrumLikeAnchor {
	l.PropHref = href
	return l
}

// Target sets the target attribute
func (l *spectrumLikeAnchor) Target(target string) *spectrumLikeAnchor {
	l.PropTarget = target
	return l
}

// Rel sets the rel attribute
func (l *spectrumLikeAnchor) Rel(rel string) *spectrumLikeAnchor {
	l.PropRel = rel
	return l
}

// Download sets the download attribute
func (l *spectrumLikeAnchor) Download(download string) *spectrumLikeAnchor {
	l.PropDownload = download
	return l
}

// Label sets the label attribute
func (l *spectrumLikeAnchor) Label(label string) *spectrumLikeAnchor {
	l.PropLabel = label
	return l
}

// OnClick sets the click event handler
func (l *spectrumLikeAnchor) OnClick(handler app.EventHandler) *spectrumLikeAnchor {
	l.PropOnClick = handler
	return l
}

// Child adds a child element
func (l *spectrumLikeAnchor) Child(child app.UI) *spectrumLikeAnchor {
	l.PropChildren = append(l.PropChildren, child)
	return l
}

// Children adds multiple child elements
func (l *spectrumLikeAnchor) Children(children ...app.UI) *spectrumLikeAnchor {
	l.PropChildren = append(l.PropChildren, children...)
	return l
}

// Render renders the anchor-like component
func (l *spectrumLikeAnchor) Render() app.UI {
	anchor := app.Elem("a")

	// Set attributes
	if l.PropHref != "" {
		anchor = anchor.Attr("href", l.PropHref)
	}
	if l.PropTarget != "" {
		anchor = anchor.Attr("target", l.PropTarget)
	}
	if l.PropRel != "" {
		anchor = anchor.Attr("rel", l.PropRel)
	}
	if l.PropDownload != "" {
		anchor = anchor.Attr("download", l.PropDownload)
	}
	if l.PropLabel != "" {
		anchor = anchor.Attr("label", l.PropLabel)
	}

	// Add event handler
	if l.PropOnClick != nil {
		anchor = anchor.OnClick(l.PropOnClick)
	}

	// Add children if provided
	if len(l.PropChildren) > 0 {
		anchor = anchor.Body(l.PropChildren...)
	}

	return anchor
}

// spectrumObserveSlotText represents a component that observes slot text changes
type spectrumObserveSlotText struct {
	app.Compo

	// Properties
	PropSlotName string
	PropOnChange app.EventHandler

	// Children
	PropChildren []app.UI
}

// NewObserveSlotText creates a new slot text observer component
func NewObserveSlotText() *spectrumObserveSlotText {
	return &spectrumObserveSlotText{}
}

// SlotName sets the name of the slot to observe
func (o *spectrumObserveSlotText) SlotName(name string) *spectrumObserveSlotText {
	o.PropSlotName = name
	return o
}

// OnChange sets the change event handler
func (o *spectrumObserveSlotText) OnChange(handler app.EventHandler) *spectrumObserveSlotText {
	o.PropOnChange = handler
	return o
}

// Child adds a child element
func (o *spectrumObserveSlotText) Child(child app.UI) *spectrumObserveSlotText {
	o.PropChildren = append(o.PropChildren, child)
	return o
}

// Children adds multiple child elements
func (o *spectrumObserveSlotText) Children(children ...app.UI) *spectrumObserveSlotText {
	o.PropChildren = append(o.PropChildren, children...)
	return o
}

// Render renders the slot text observer component
func (o *spectrumObserveSlotText) Render() app.UI {
	observer := app.Elem("div")

	// Set attributes
	if o.PropSlotName != "" {
		observer = observer.Attr("slot", o.PropSlotName)
	}

	// Add event handler
	if o.PropOnChange != nil {
		observer = observer.On("change", o.PropOnChange)
	}

	// Add children if provided
	if len(o.PropChildren) > 0 {
		observer = observer.Body(o.PropChildren...)
	}

	return observer
}
