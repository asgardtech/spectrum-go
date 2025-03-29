package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumCoreTokens represents the core design tokens used throughout Spectrum components
type spectrumCoreTokens struct {
	app.Compo

	// Properties
	PropColorScheme     string
	PropScale           string
	PropSystem          string
	PropDirection       string
	PropLanguage        string
	PropAnimationSpeed  string
	PropAnimationEasing string
	PropFontFamily      string
	PropFontSize        string
	PropLineHeight      string
	PropLetterSpacing   string
	PropFontWeight      string
	PropBorderRadius    string
	PropBorderWidth     string
	PropShadow          string
	PropSpacing         string
	PropZIndex          string
	PropOpacity         string
	PropTransition      string

	// Children
	PropChildren []app.UI
}

// NewSpectrumCoreTokens creates a new Spectrum core tokens component
func NewSpectrumCoreTokens() *spectrumCoreTokens {
	return &spectrumCoreTokens{}
}

// ColorScheme sets the color scheme
func (t *spectrumCoreTokens) ColorScheme(scheme string) *spectrumCoreTokens {
	t.PropColorScheme = scheme
	return t
}

// Scale sets the scale
func (t *spectrumCoreTokens) Scale(scale string) *spectrumCoreTokens {
	t.PropScale = scale
	return t
}

// System sets the system
func (t *spectrumCoreTokens) System(system string) *spectrumCoreTokens {
	t.PropSystem = system
	return t
}

// Direction sets the direction
func (t *spectrumCoreTokens) Direction(direction string) *spectrumCoreTokens {
	t.PropDirection = direction
	return t
}

// Language sets the language
func (t *spectrumCoreTokens) Language(language string) *spectrumCoreTokens {
	t.PropLanguage = language
	return t
}

// AnimationSpeed sets the animation speed
func (t *spectrumCoreTokens) AnimationSpeed(speed string) *spectrumCoreTokens {
	t.PropAnimationSpeed = speed
	return t
}

// AnimationEasing sets the animation easing
func (t *spectrumCoreTokens) AnimationEasing(easing string) *spectrumCoreTokens {
	t.PropAnimationEasing = easing
	return t
}

// FontFamily sets the font family
func (t *spectrumCoreTokens) FontFamily(family string) *spectrumCoreTokens {
	t.PropFontFamily = family
	return t
}

// FontSize sets the font size
func (t *spectrumCoreTokens) FontSize(size string) *spectrumCoreTokens {
	t.PropFontSize = size
	return t
}

// LineHeight sets the line height
func (t *spectrumCoreTokens) LineHeight(height string) *spectrumCoreTokens {
	t.PropLineHeight = height
	return t
}

// LetterSpacing sets the letter spacing
func (t *spectrumCoreTokens) LetterSpacing(spacing string) *spectrumCoreTokens {
	t.PropLetterSpacing = spacing
	return t
}

// FontWeight sets the font weight
func (t *spectrumCoreTokens) FontWeight(weight string) *spectrumCoreTokens {
	t.PropFontWeight = weight
	return t
}

// BorderRadius sets the border radius
func (t *spectrumCoreTokens) BorderRadius(radius string) *spectrumCoreTokens {
	t.PropBorderRadius = radius
	return t
}

// BorderWidth sets the border width
func (t *spectrumCoreTokens) BorderWidth(width string) *spectrumCoreTokens {
	t.PropBorderWidth = width
	return t
}

// Shadow sets the shadow
func (t *spectrumCoreTokens) Shadow(shadow string) *spectrumCoreTokens {
	t.PropShadow = shadow
	return t
}

// Spacing sets the spacing
func (t *spectrumCoreTokens) Spacing(spacing string) *spectrumCoreTokens {
	t.PropSpacing = spacing
	return t
}

// ZIndex sets the z-index
func (t *spectrumCoreTokens) ZIndex(index string) *spectrumCoreTokens {
	t.PropZIndex = index
	return t
}

// Opacity sets the opacity
func (t *spectrumCoreTokens) Opacity(opacity string) *spectrumCoreTokens {
	t.PropOpacity = opacity
	return t
}

// Transition sets the transition
func (t *spectrumCoreTokens) Transition(transition string) *spectrumCoreTokens {
	t.PropTransition = transition
	return t
}

// Child adds a child element
func (t *spectrumCoreTokens) Child(child app.UI) *spectrumCoreTokens {
	t.PropChildren = append(t.PropChildren, child)
	return t
}

// Children adds multiple child elements
func (t *spectrumCoreTokens) Children(children ...app.UI) *spectrumCoreTokens {
	t.PropChildren = append(t.PropChildren, children...)
	return t
}

// Render renders the Spectrum core tokens component
func (t *spectrumCoreTokens) Render() app.UI {
	tokens := app.Elem("div")

	// Set tokens
	if t.PropColorScheme != "" {
		tokens = tokens.Attr("color-scheme", t.PropColorScheme)
	}
	if t.PropScale != "" {
		tokens = tokens.Attr("scale", t.PropScale)
	}
	if t.PropSystem != "" {
		tokens = tokens.Attr("system", t.PropSystem)
	}
	if t.PropDirection != "" {
		tokens = tokens.Attr("dir", t.PropDirection)
	}
	if t.PropLanguage != "" {
		tokens = tokens.Attr("lang", t.PropLanguage)
	}
	if t.PropAnimationSpeed != "" {
		tokens = tokens.Style("--spectrum-animation-speed", t.PropAnimationSpeed)
	}
	if t.PropAnimationEasing != "" {
		tokens = tokens.Style("--spectrum-animation-easing", t.PropAnimationEasing)
	}
	if t.PropFontFamily != "" {
		tokens = tokens.Style("--spectrum-font-family", t.PropFontFamily)
	}
	if t.PropFontSize != "" {
		tokens = tokens.Style("--spectrum-font-size", t.PropFontSize)
	}
	if t.PropLineHeight != "" {
		tokens = tokens.Style("--spectrum-line-height", t.PropLineHeight)
	}
	if t.PropLetterSpacing != "" {
		tokens = tokens.Style("--spectrum-letter-spacing", t.PropLetterSpacing)
	}
	if t.PropFontWeight != "" {
		tokens = tokens.Style("--spectrum-font-weight", t.PropFontWeight)
	}
	if t.PropBorderRadius != "" {
		tokens = tokens.Style("--spectrum-border-radius", t.PropBorderRadius)
	}
	if t.PropBorderWidth != "" {
		tokens = tokens.Style("--spectrum-border-width", t.PropBorderWidth)
	}
	if t.PropShadow != "" {
		tokens = tokens.Style("--spectrum-shadow", t.PropShadow)
	}
	if t.PropSpacing != "" {
		tokens = tokens.Style("--spectrum-spacing", t.PropSpacing)
	}
	if t.PropZIndex != "" {
		tokens = tokens.Style("--spectrum-z-index", t.PropZIndex)
	}
	if t.PropOpacity != "" {
		tokens = tokens.Style("--spectrum-opacity", t.PropOpacity)
	}
	if t.PropTransition != "" {
		tokens = tokens.Style("--spectrum-transition", t.PropTransition)
	}

	// Add children if provided
	if len(t.PropChildren) > 0 {
		tokens = tokens.Body(t.PropChildren...)
	}

	return tokens
}
