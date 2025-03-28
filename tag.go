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

// SpectrumTag represents an sp-tag component
type SpectrumTag struct {
	app.Compo

	// Properties
	text      string
	size      TagSize
	deletable bool
	disabled  bool
	readonly  bool
	invalid   bool

	// Slots
	avatar app.UI
	icon   app.UI

	// Event handlers
	onDelete app.EventHandler
}

// Tag creates a new tag component
func Tag() *SpectrumTag {
	return &SpectrumTag{
		size: TagSizeM, // Default size is medium
	}
}

// Text sets the tag text content
func (t *SpectrumTag) Text(text string) *SpectrumTag {
	t.text = text
	return t
}

// Size sets the visual size of the tag
func (t *SpectrumTag) Size(size TagSize) *SpectrumTag {
	t.size = size
	return t
}

// Deletable sets whether the tag is deletable
func (t *SpectrumTag) Deletable(deletable bool) *SpectrumTag {
	t.deletable = deletable
	return t
}

// Disabled sets the disabled state of the tag
func (t *SpectrumTag) Disabled(disabled bool) *SpectrumTag {
	t.disabled = disabled
	return t
}

// Readonly sets the readonly state of the tag
func (t *SpectrumTag) Readonly(readonly bool) *SpectrumTag {
	t.readonly = readonly
	return t
}

// Invalid sets the invalid state of the tag
func (t *SpectrumTag) Invalid(invalid bool) *SpectrumTag {
	t.invalid = invalid
	return t
}

// Avatar sets the avatar in the avatar slot
func (t *SpectrumTag) Avatar(avatar app.UI) *SpectrumTag {
	t.avatar = avatar
	return t
}

// Icon sets the icon in the icon slot
func (t *SpectrumTag) Icon(icon app.UI) *SpectrumTag {
	t.icon = icon
	return t
}

// OnDelete sets the delete event handler
func (t *SpectrumTag) OnDelete(handler app.EventHandler) *SpectrumTag {
	t.onDelete = handler
	return t
}

// Render renders the tag component
func (t *SpectrumTag) Render() app.UI {
	tag := app.Elem("sp-tag").
		Attr("size", string(t.size)).
		Attr("deletable", t.deletable).
		Attr("disabled", t.disabled).
		Attr("readonly", t.readonly).
		Attr("invalid", t.invalid)

	// Add event handlers
	if t.onDelete != nil {
		tag = tag.On("delete", t.onDelete)
	}

	// Add slots if provided
	if t.avatar != nil {
		avatar := t.avatar
		if avatarWithSlot, ok := avatar.(interface{ Slot(string) app.UI }); ok {
			avatar = avatarWithSlot.Slot("avatar")
		} else {
			avatar = app.Elem("div").
				Attr("slot", "avatar").
				Body(avatar)
		}
		tag = tag.Body(avatar)
	}

	if t.icon != nil {
		icon := t.icon
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
	if t.text != "" {
		tag = tag.Text(t.text)
	}

	return tag
}
