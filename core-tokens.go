package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumCoreTokens represents the core design tokens used throughout Spectrum components
type SpectrumCoreTokens struct {
	app.Compo

	// Properties
	colorScheme     string
	scale           string
	system          string
	direction       string
	language        string
	animationSpeed  string
	animationEasing string
	fontFamily      string
	fontSize        string
	lineHeight      string
	letterSpacing   string
	fontWeight      string
	borderRadius    string
	borderWidth     string
	shadow          string
	spacing         string
	zIndex          string
	opacity         string
	transition      string

	// Children
	children []app.UI
}

// NewSpectrumCoreTokens creates a new Spectrum core tokens component
func NewSpectrumCoreTokens() *SpectrumCoreTokens {
	return &SpectrumCoreTokens{}
}

// ColorScheme sets the color scheme
func (t *SpectrumCoreTokens) ColorScheme(scheme string) *SpectrumCoreTokens {
	t.colorScheme = scheme
	return t
}

// Scale sets the scale
func (t *SpectrumCoreTokens) Scale(scale string) *SpectrumCoreTokens {
	t.scale = scale
	return t
}

// System sets the system
func (t *SpectrumCoreTokens) System(system string) *SpectrumCoreTokens {
	t.system = system
	return t
}

// Direction sets the direction
func (t *SpectrumCoreTokens) Direction(direction string) *SpectrumCoreTokens {
	t.direction = direction
	return t
}

// Language sets the language
func (t *SpectrumCoreTokens) Language(language string) *SpectrumCoreTokens {
	t.language = language
	return t
}

// AnimationSpeed sets the animation speed
func (t *SpectrumCoreTokens) AnimationSpeed(speed string) *SpectrumCoreTokens {
	t.animationSpeed = speed
	return t
}

// AnimationEasing sets the animation easing
func (t *SpectrumCoreTokens) AnimationEasing(easing string) *SpectrumCoreTokens {
	t.animationEasing = easing
	return t
}

// FontFamily sets the font family
func (t *SpectrumCoreTokens) FontFamily(family string) *SpectrumCoreTokens {
	t.fontFamily = family
	return t
}

// FontSize sets the font size
func (t *SpectrumCoreTokens) FontSize(size string) *SpectrumCoreTokens {
	t.fontSize = size
	return t
}

// LineHeight sets the line height
func (t *SpectrumCoreTokens) LineHeight(height string) *SpectrumCoreTokens {
	t.lineHeight = height
	return t
}

// LetterSpacing sets the letter spacing
func (t *SpectrumCoreTokens) LetterSpacing(spacing string) *SpectrumCoreTokens {
	t.letterSpacing = spacing
	return t
}

// FontWeight sets the font weight
func (t *SpectrumCoreTokens) FontWeight(weight string) *SpectrumCoreTokens {
	t.fontWeight = weight
	return t
}

// BorderRadius sets the border radius
func (t *SpectrumCoreTokens) BorderRadius(radius string) *SpectrumCoreTokens {
	t.borderRadius = radius
	return t
}

// BorderWidth sets the border width
func (t *SpectrumCoreTokens) BorderWidth(width string) *SpectrumCoreTokens {
	t.borderWidth = width
	return t
}

// Shadow sets the shadow
func (t *SpectrumCoreTokens) Shadow(shadow string) *SpectrumCoreTokens {
	t.shadow = shadow
	return t
}

// Spacing sets the spacing
func (t *SpectrumCoreTokens) Spacing(spacing string) *SpectrumCoreTokens {
	t.spacing = spacing
	return t
}

// ZIndex sets the z-index
func (t *SpectrumCoreTokens) ZIndex(index string) *SpectrumCoreTokens {
	t.zIndex = index
	return t
}

// Opacity sets the opacity
func (t *SpectrumCoreTokens) Opacity(opacity string) *SpectrumCoreTokens {
	t.opacity = opacity
	return t
}

// Transition sets the transition
func (t *SpectrumCoreTokens) Transition(transition string) *SpectrumCoreTokens {
	t.transition = transition
	return t
}

// Child adds a child element
func (t *SpectrumCoreTokens) Child(child app.UI) *SpectrumCoreTokens {
	t.children = append(t.children, child)
	return t
}

// Children adds multiple child elements
func (t *SpectrumCoreTokens) Children(children ...app.UI) *SpectrumCoreTokens {
	t.children = append(t.children, children...)
	return t
}

// Render renders the Spectrum core tokens component
func (t *SpectrumCoreTokens) Render() app.UI {
	tokens := app.Elem("div")

	// Set tokens
	if t.colorScheme != "" {
		tokens = tokens.Attr("color-scheme", t.colorScheme)
	}
	if t.scale != "" {
		tokens = tokens.Attr("scale", t.scale)
	}
	if t.system != "" {
		tokens = tokens.Attr("system", t.system)
	}
	if t.direction != "" {
		tokens = tokens.Attr("dir", t.direction)
	}
	if t.language != "" {
		tokens = tokens.Attr("lang", t.language)
	}
	if t.animationSpeed != "" {
		tokens = tokens.Style("--spectrum-animation-speed", t.animationSpeed)
	}
	if t.animationEasing != "" {
		tokens = tokens.Style("--spectrum-animation-easing", t.animationEasing)
	}
	if t.fontFamily != "" {
		tokens = tokens.Style("--spectrum-font-family", t.fontFamily)
	}
	if t.fontSize != "" {
		tokens = tokens.Style("--spectrum-font-size", t.fontSize)
	}
	if t.lineHeight != "" {
		tokens = tokens.Style("--spectrum-line-height", t.lineHeight)
	}
	if t.letterSpacing != "" {
		tokens = tokens.Style("--spectrum-letter-spacing", t.letterSpacing)
	}
	if t.fontWeight != "" {
		tokens = tokens.Style("--spectrum-font-weight", t.fontWeight)
	}
	if t.borderRadius != "" {
		tokens = tokens.Style("--spectrum-border-radius", t.borderRadius)
	}
	if t.borderWidth != "" {
		tokens = tokens.Style("--spectrum-border-width", t.borderWidth)
	}
	if t.shadow != "" {
		tokens = tokens.Style("--spectrum-shadow", t.shadow)
	}
	if t.spacing != "" {
		tokens = tokens.Style("--spectrum-spacing", t.spacing)
	}
	if t.zIndex != "" {
		tokens = tokens.Style("--spectrum-z-index", t.zIndex)
	}
	if t.opacity != "" {
		tokens = tokens.Style("--spectrum-opacity", t.opacity)
	}
	if t.transition != "" {
		tokens = tokens.Style("--spectrum-transition", t.transition)
	}

	// Add children if provided
	if len(t.children) > 0 {
		tokens = tokens.Body(t.children...)
	}

	return tokens
}
