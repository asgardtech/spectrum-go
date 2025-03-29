package sp

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TextDirection represents the direction of text flow (LTR or RTL)
type TextDirection string

const (
	// TextDirectionLTR is left-to-right text direction
	TextDirectionLTR TextDirection = "ltr"

	// TextDirectionRTL is right-to-left text direction
	TextDirectionRTL TextDirection = "rtl"
)

// spectrumElementMixin provides core functionality shared by Spectrum components.
// It enhances structs that embed app.Compo with additional capabilities.
type spectrumElementMixin struct {
	// Direction of text (LTR or RTL)
	PropDir TextDirection

	// Whether the direction was explicitly set
	PropDirExplicitlySet bool
}

// Dir sets the text direction of the component.
// Setting this explicitly prevents inheriting direction from document or theme.
func (m *spectrumElementMixin) Dir(dir TextDirection) *spectrumElementMixin {
	m.PropDir = dir
	m.PropDirExplicitlySet = true
	return m
}

// GetDir returns the current text direction.
func (m *spectrumElementMixin) GetDir() TextDirection {
	if m.PropDir == "" {
		return TextDirectionLTR // Default to LTR if not set
	}
	return m.PropDir
}

// IsLTR returns true if the text direction is left-to-right.
func (m *spectrumElementMixin) IsLTR() bool {
	return m.GetDir() == TextDirectionLTR
}

// IsRTL returns true if the text direction is right-to-left.
func (m *spectrumElementMixin) IsRTL() bool {
	return m.GetDir() == TextDirectionRTL
}

// DirWasExplicitlySet returns whether the direction was explicitly set.
func (m *spectrumElementMixin) DirWasExplicitlySet() bool {
	return m.PropDirExplicitlySet
}

// spectrumElement is the base struct for Spectrum Web Components.
// It embeds app.Compo for standard component functionality and
// spectrumElementMixin for Spectrum-specific features.
type spectrumElement struct {
	app.Compo
	spectrumElementMixin
}

// NewSpectrumElement creates a new spectrumElement with default values.
func NewSpectrumElement() *spectrumElement {
	return &spectrumElement{
		spectrumElementMixin: spectrumElementMixin{
			PropDir:              TextDirectionLTR,
			PropDirExplicitlySet: false,
		},
	}
}

// ApplySpectrumAttributes applies common Spectrum component attributes to an HTML element.
// This is a utility function for use in Render() methods to ensure consistent attribute handling.
func ApplySpectrumAttributes(element app.HTMLElem, mixin *spectrumElementMixin) app.HTMLElem {
	// Apply direction attribute if explicitly set
	if mixin.PropDirExplicitlySet {
		element = element.Attr("dir", string(mixin.GetDir()))
	}

	return element
}

// Observer interface for components that need to observe document/theme direction changes
type Observer interface {
	OnDirChange(newDir TextDirection)
}

// DirectionController manages text direction observation and propagation
type DirectionController struct {
	// Host component that will receive notifications
	host Observer

	// Direction manager would handle observation of document/theme
	// in a real implementation with browser APIs
}

// NewDirectionController creates a new direction controller
func NewDirectionController(host Observer) *DirectionController {
	return &DirectionController{
		host: host,
	}
}

// In a full implementation, this would use browser APIs to observe document/theme
// direction changes and notify the host component.
// For now, it's a placeholder for the functionality that would be implemented
// in a real browser environment.

// UpdateDirection would be called when document/theme direction changes
func (c *DirectionController) UpdateDirection(dir TextDirection) {
	c.host.OnDirChange(dir)
}

// This package also would include other base functionality like:
// - FocusVisibleController for managing focus visibility
// - TabIndexManager for managing component focus
// - ShadowRootManager for component shadow DOM interactions
// These would be implemented as needed for the Go-app context.
