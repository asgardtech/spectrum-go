package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumTags represents an sp-tags component
// which is a container for sp-tag components
type spectrumTags struct {
	app.Compo

	// Children tags
	PropChildren []*spectrumTag
}

// Tags creates a new tags container component
func Tags() *spectrumTags {
	return &spectrumTags{}
}

// Children sets the child tag components
func (t *spectrumTags) Children(children ...*spectrumTag) *spectrumTags {
	t.PropChildren = children
	return t
}

// Render renders the tags container component
func (t *spectrumTags) Render() app.UI {
	tags := app.Elem("sp-tags")

	// Add all child tag elements
	if len(t.PropChildren) > 0 {
		// Convert spectrumTag children to app.UI for the Body method
		uiChildren := make([]app.UI, len(t.PropChildren))
		for i, child := range t.PropChildren {
			uiChildren[i] = child
		}
		tags = tags.Body(uiChildren...)
	}

	return tags
}
