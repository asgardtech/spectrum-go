package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TagSize represents the visual size of a tag
type TagSize string

// Tag sizes
const (
	TagSizeS TagSize = "s"
	TagSizeM TagSize = "m"
	TagSizeL TagSize = "l"
)

// spectrumTag represents an sp-tag component
type spectrumTag struct {
	app.Compo

	// Properties
	PropText      string
	PropSize      TagSize
	PropDeletable bool
	PropDisabled  bool
	PropReadonly  bool
	PropInvalid   bool

	// Slots
	PropAvatar app.UI
	PropIcon   app.UI

	// Event handlers
	PropOnDelete app.EventHandler
}

// Tag creates a new tag component
func Tag() *spectrumTag {
	return &spectrumTag{
		PropSize: TagSizeM, // Default size is medium
	}
}

// Text sets the tag text content
func (t *spectrumTag) Text(text string) *spectrumTag {
	t.PropText = text
	return t
}

// Size sets the visual size of the tag
func (t *spectrumTag) Size(size TagSize) *spectrumTag {
	t.PropSize = size
	return t
}

// Deletable sets whether the tag is deletable
func (t *spectrumTag) Deletable(deletable bool) *spectrumTag {
	t.PropDeletable = deletable
	return t
}

// Disabled sets the disabled state of the tag
func (t *spectrumTag) Disabled(disabled bool) *spectrumTag {
	t.PropDisabled = disabled
	return t
}

// Readonly sets the readonly state of the tag
func (t *spectrumTag) Readonly(readonly bool) *spectrumTag {
	t.PropReadonly = readonly
	return t
}

// Invalid sets the invalid state of the tag
func (t *spectrumTag) Invalid(invalid bool) *spectrumTag {
	t.PropInvalid = invalid
	return t
}

// Avatar sets the avatar in the avatar slot
func (t *spectrumTag) Avatar(avatar app.UI) *spectrumTag {
	t.PropAvatar = avatar
	return t
}

// Icon sets the icon in the icon slot
func (t *spectrumTag) Icon(icon app.UI) *spectrumTag {
	t.PropIcon = icon
	return t
}

// OnDelete sets the delete event handler
func (t *spectrumTag) OnDelete(handler app.EventHandler) *spectrumTag {
	t.PropOnDelete = handler
	return t
}

// Render renders the tag component
func (t *spectrumTag) Render() app.UI {
	tag := app.Elem("sp-tag").
		Attr("size", string(t.PropSize)).
		Attr("deletable", t.PropDeletable).
		Attr("disabled", t.PropDisabled).
		Attr("readonly", t.PropReadonly).
		Attr("invalid", t.PropInvalid)

	// Add event handlers
	if t.PropOnDelete != nil {
		tag = tag.On("delete", t.PropOnDelete)
	}

	// Add slots if provided
	if t.PropAvatar != nil {
		avatar := t.PropAvatar
		if avatarWithSlot, ok := avatar.(interface{ Slot(string) app.UI }); ok {
			avatar = avatarWithSlot.Slot("avatar")
		} else {
			avatar = app.Elem("div").
				Attr("slot", "avatar").
				Body(avatar)
		}
		tag = tag.Body(avatar)
	}

	if t.PropIcon != nil {
		icon := t.PropIcon
		if iconWithSlot, ok := icon.(interface{ Slot(string) app.UI }); ok {
			icon = iconWithSlot.Slot("icon")
		} else {
			icon = app.Elem("div").
				Attr("slot", "icon").
				Body(icon)
		}
		tag = tag.Body(icon)
	}

	// Add text content
	if t.PropText != "" {
		tag = tag.Text(t.PropText)
	}

	return tag
}
