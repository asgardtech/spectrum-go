package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// Focusable represents a focusable component
type Focusable struct {
	app.Compo

	// Properties
	tabIndex int
	disabled bool
	onFocus  app.EventHandler
	onBlur   app.EventHandler

	// Children
	children []app.UI
}

// NewFocusable creates a new focusable component
func NewFocusable() *Focusable {
	return &Focusable{}
}

// TabIndex sets the tab index
func (f *Focusable) TabIndex(index int) *Focusable {
	f.tabIndex = index
	return f
}

// Disabled sets the disabled state
func (f *Focusable) Disabled(disabled bool) *Focusable {
	f.disabled = disabled
	return f
}

// OnFocus sets the focus event handler
func (f *Focusable) OnFocus(handler app.EventHandler) *Focusable {
	f.onFocus = handler
	return f
}

// OnBlur sets the blur event handler
func (f *Focusable) OnBlur(handler app.EventHandler) *Focusable {
	f.onBlur = handler
	return f
}

// Child adds a child element
func (f *Focusable) Child(child app.UI) *Focusable {
	f.children = append(f.children, child)
	return f
}

// Children adds multiple child elements
func (f *Focusable) Children(children ...app.UI) *Focusable {
	f.children = append(f.children, children...)
	return f
}

// Render renders the focusable component
func (f *Focusable) Render() app.UI {
	focusable := app.Elem("div")

	// Set attributes
	if f.tabIndex != 0 {
		focusable = focusable.Attr("tabindex", f.tabIndex)
	}
	if f.disabled {
		focusable = focusable.Attr("disabled", true)
	}

	// Add event handlers
	if f.onFocus != nil {
		focusable = focusable.On("focus", f.onFocus)
	}
	if f.onBlur != nil {
		focusable = focusable.On("blur", f.onBlur)
	}

	// Add children if provided
	if len(f.children) > 0 {
		focusable = focusable.Body(f.children...)
	}

	return focusable
}

// LikeAnchor represents a component that behaves like an anchor element
type LikeAnchor struct {
	app.Compo

	// Properties
	href     string
	target   string
	rel      string
	download string
	label    string
	onClick  app.EventHandler

	// Children
	children []app.UI
}

// NewLikeAnchor creates a new anchor-like component
func NewLikeAnchor() *LikeAnchor {
	return &LikeAnchor{}
}

// Href sets the href attribute
func (l *LikeAnchor) Href(href string) *LikeAnchor {
	l.href = href
	return l
}

// Target sets the target attribute
func (l *LikeAnchor) Target(target string) *LikeAnchor {
	l.target = target
	return l
}

// Rel sets the rel attribute
func (l *LikeAnchor) Rel(rel string) *LikeAnchor {
	l.rel = rel
	return l
}

// Download sets the download attribute
func (l *LikeAnchor) Download(download string) *LikeAnchor {
	l.download = download
	return l
}

// Label sets the label attribute
func (l *LikeAnchor) Label(label string) *LikeAnchor {
	l.label = label
	return l
}

// OnClick sets the click event handler
func (l *LikeAnchor) OnClick(handler app.EventHandler) *LikeAnchor {
	l.onClick = handler
	return l
}

// Child adds a child element
func (l *LikeAnchor) Child(child app.UI) *LikeAnchor {
	l.children = append(l.children, child)
	return l
}

// Children adds multiple child elements
func (l *LikeAnchor) Children(children ...app.UI) *LikeAnchor {
	l.children = append(l.children, children...)
	return l
}

// Render renders the anchor-like component
func (l *LikeAnchor) Render() app.UI {
	anchor := app.Elem("a")

	// Set attributes
	if l.href != "" {
		anchor = anchor.Attr("href", l.href)
	}
	if l.target != "" {
		anchor = anchor.Attr("target", l.target)
	}
	if l.rel != "" {
		anchor = anchor.Attr("rel", l.rel)
	}
	if l.download != "" {
		anchor = anchor.Attr("download", l.download)
	}
	if l.label != "" {
		anchor = anchor.Attr("label", l.label)
	}

	// Add event handler
	if l.onClick != nil {
		anchor = anchor.OnClick(l.onClick)
	}

	// Add children if provided
	if len(l.children) > 0 {
		anchor = anchor.Body(l.children...)
	}

	return anchor
}

// ObserveSlotText represents a component that observes slot text changes
type ObserveSlotText struct {
	app.Compo

	// Properties
	slotName string
	onChange app.EventHandler

	// Children
	children []app.UI
}

// NewObserveSlotText creates a new slot text observer component
func NewObserveSlotText() *ObserveSlotText {
	return &ObserveSlotText{}
}

// SlotName sets the name of the slot to observe
func (o *ObserveSlotText) SlotName(name string) *ObserveSlotText {
	o.slotName = name
	return o
}

// OnChange sets the change event handler
func (o *ObserveSlotText) OnChange(handler app.EventHandler) *ObserveSlotText {
	o.onChange = handler
	return o
}

// Child adds a child element
func (o *ObserveSlotText) Child(child app.UI) *ObserveSlotText {
	o.children = append(o.children, child)
	return o
}

// Children adds multiple child elements
func (o *ObserveSlotText) Children(children ...app.UI) *ObserveSlotText {
	o.children = append(o.children, children...)
	return o
}

// Render renders the slot text observer component
func (o *ObserveSlotText) Render() app.UI {
	observer := app.Elem("div")

	// Set attributes
	if o.slotName != "" {
		observer = observer.Attr("slot", o.slotName)
	}

	// Add event handler
	if o.onChange != nil {
		observer = observer.On("change", o.onChange)
	}

	// Add children if provided
	if len(o.children) > 0 {
		observer = observer.Body(o.children...)
	}

	return observer
}
