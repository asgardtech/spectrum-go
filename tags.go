package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTags represents an sp-tags component
// which is a container for sp-tag components
type SpectrumTags struct {
	app.Compo

	// Children tags
	children []*SpectrumTag
}

// Tags creates a new tags container component
func Tags() *SpectrumTags {
	return &SpectrumTags{}
}

// Children sets the child tag components
func (t *SpectrumTags) Children(children ...*SpectrumTag) *SpectrumTags {
	t.children = children
	return t
}

// Render renders the tags container component
func (t *SpectrumTags) Render() app.UI {
	tags := app.Elem("sp-tags")

	// Add all child tag elements
	if len(t.children) > 0 {
		// Convert SpectrumTag children to app.UI for the Body method
		uiChildren := make([]app.UI, len(t.children))
		for i, child := range t.children {
			uiChildren[i] = child
		}
		tags = tags.Body(uiChildren...)
	}

	return tags
}
