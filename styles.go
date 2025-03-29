package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumStyles represents a collection of Spectrum CSS custom properties
type spectrumStyles struct {
	app.Compo

	// Properties
	PropBackgroundColor string
	PropTextColor       string
	PropBorderColor     string
	PropFontSize        string
	PropFontWeight      string
	PropLineHeight      string
	PropPadding         string
	PropMargin          string
	PropBorderRadius    string
	PropBoxShadow       string
	PropOpacity         string
	PropTransition      string
	PropTransform       string
	PropDisplay         string
	PropPosition        string
	PropTop             string
	PropRight           string
	PropBottom          string
	PropLeft            string
	PropWidth           string
	PropHeight          string
	PropMinWidth        string
	PropMinHeight       string
	PropMaxWidth        string
	PropMaxHeight       string
	PropOverflow        string
	PropZIndex          string
	PropCursor          string
	PropPointerEvents   string

	// Children
	PropChildren []app.UI
}

// NewSpectrumStyles creates a new Spectrum styles component
func NewSpectrumStyles() *spectrumStyles {
	return &spectrumStyles{}
}

// BackgroundColor sets the background color
func (s *spectrumStyles) BackgroundColor(color string) *spectrumStyles {
	s.PropBackgroundColor = color
	return s
}

// TextColor sets the text color
func (s *spectrumStyles) TextColor(color string) *spectrumStyles {
	s.PropTextColor = color
	return s
}

// BorderColor sets the border color
func (s *spectrumStyles) BorderColor(color string) *spectrumStyles {
	s.PropBorderColor = color
	return s
}

// FontSize sets the font size
func (s *spectrumStyles) FontSize(size string) *spectrumStyles {
	s.PropFontSize = size
	return s
}

// FontWeight sets the font weight
func (s *spectrumStyles) FontWeight(weight string) *spectrumStyles {
	s.PropFontWeight = weight
	return s
}

// LineHeight sets the line height
func (s *spectrumStyles) LineHeight(height string) *spectrumStyles {
	s.PropLineHeight = height
	return s
}

// Padding sets the padding
func (s *spectrumStyles) Padding(padding string) *spectrumStyles {
	s.PropPadding = padding
	return s
}

// Margin sets the margin
func (s *spectrumStyles) Margin(margin string) *spectrumStyles {
	s.PropMargin = margin
	return s
}

// BorderRadius sets the border radius
func (s *spectrumStyles) BorderRadius(radius string) *spectrumStyles {
	s.PropBorderRadius = radius
	return s
}

// BoxShadow sets the box shadow
func (s *spectrumStyles) BoxShadow(shadow string) *spectrumStyles {
	s.PropBoxShadow = shadow
	return s
}

// Opacity sets the opacity
func (s *spectrumStyles) Opacity(opacity string) *spectrumStyles {
	s.PropOpacity = opacity
	return s
}

// Transition sets the transition
func (s *spectrumStyles) Transition(transition string) *spectrumStyles {
	s.PropTransition = transition
	return s
}

// Transform sets the transform
func (s *spectrumStyles) Transform(transform string) *spectrumStyles {
	s.PropTransform = transform
	return s
}

// Display sets the display property
func (s *spectrumStyles) Display(display string) *spectrumStyles {
	s.PropDisplay = display
	return s
}

// Position sets the position
func (s *spectrumStyles) Position(position string) *spectrumStyles {
	s.PropPosition = position
	return s
}

// Top sets the top position
func (s *spectrumStyles) Top(top string) *spectrumStyles {
	s.PropTop = top
	return s
}

// Right sets the right position
func (s *spectrumStyles) Right(right string) *spectrumStyles {
	s.PropRight = right
	return s
}

// Bottom sets the bottom position
func (s *spectrumStyles) Bottom(bottom string) *spectrumStyles {
	s.PropBottom = bottom
	return s
}

// Left sets the left position
func (s *spectrumStyles) Left(left string) *spectrumStyles {
	s.PropLeft = left
	return s
}

// Width sets the width
func (s *spectrumStyles) Width(width string) *spectrumStyles {
	s.PropWidth = width
	return s
}

// Height sets the height
func (s *spectrumStyles) Height(height string) *spectrumStyles {
	s.PropHeight = height
	return s
}

// MinWidth sets the minimum width
func (s *spectrumStyles) MinWidth(width string) *spectrumStyles {
	s.PropMinWidth = width
	return s
}

// MinHeight sets the minimum height
func (s *spectrumStyles) MinHeight(height string) *spectrumStyles {
	s.PropMinHeight = height
	return s
}

// MaxWidth sets the maximum width
func (s *spectrumStyles) MaxWidth(width string) *spectrumStyles {
	s.PropMaxWidth = width
	return s
}

// MaxHeight sets the maximum height
func (s *spectrumStyles) MaxHeight(height string) *spectrumStyles {
	s.PropMaxHeight = height
	return s
}

// Overflow sets the overflow property
func (s *spectrumStyles) Overflow(overflow string) *spectrumStyles {
	s.PropOverflow = overflow
	return s
}

// ZIndex sets the z-index
func (s *spectrumStyles) ZIndex(index string) *spectrumStyles {
	s.PropZIndex = index
	return s
}

// Cursor sets the cursor style
func (s *spectrumStyles) Cursor(cursor string) *spectrumStyles {
	s.PropCursor = cursor
	return s
}

// PointerEvents sets the pointer events
func (s *spectrumStyles) PointerEvents(events string) *spectrumStyles {
	s.PropPointerEvents = events
	return s
}

// Child adds a child element
func (s *spectrumStyles) Child(child app.UI) *spectrumStyles {
	s.PropChildren = append(s.PropChildren, child)
	return s
}

// Children adds multiple child elements
func (s *spectrumStyles) Children(children ...app.UI) *spectrumStyles {
	s.PropChildren = append(s.PropChildren, children...)
	return s
}

// Render renders the Spectrum styles component
func (s *spectrumStyles) Render() app.UI {
	styles := app.Elem("div")

	// Set styles
	if s.PropBackgroundColor != "" {
		styles = styles.Style("background-color", s.PropBackgroundColor)
	}
	if s.PropTextColor != "" {
		styles = styles.Style("color", s.PropTextColor)
	}
	if s.PropBorderColor != "" {
		styles = styles.Style("border-color", s.PropBorderColor)
	}
	if s.PropFontSize != "" {
		styles = styles.Style("font-size", s.PropFontSize)
	}
	if s.PropFontWeight != "" {
		styles = styles.Style("font-weight", s.PropFontWeight)
	}
	if s.PropLineHeight != "" {
		styles = styles.Style("line-height", s.PropLineHeight)
	}
	if s.PropPadding != "" {
		styles = styles.Style("padding", s.PropPadding)
	}
	if s.PropMargin != "" {
		styles = styles.Style("margin", s.PropMargin)
	}
	if s.PropBorderRadius != "" {
		styles = styles.Style("border-radius", s.PropBorderRadius)
	}
	if s.PropBoxShadow != "" {
		styles = styles.Style("box-shadow", s.PropBoxShadow)
	}
	if s.PropOpacity != "" {
		styles = styles.Style("opacity", s.PropOpacity)
	}
	if s.PropTransition != "" {
		styles = styles.Style("transition", s.PropTransition)
	}
	if s.PropTransform != "" {
		styles = styles.Style("transform", s.PropTransform)
	}
	if s.PropDisplay != "" {
		styles = styles.Style("display", s.PropDisplay)
	}
	if s.PropPosition != "" {
		styles = styles.Style("position", s.PropPosition)
	}
	if s.PropTop != "" {
		styles = styles.Style("top", s.PropTop)
	}
	if s.PropRight != "" {
		styles = styles.Style("right", s.PropRight)
	}
	if s.PropBottom != "" {
		styles = styles.Style("bottom", s.PropBottom)
	}
	if s.PropLeft != "" {
		styles = styles.Style("left", s.PropLeft)
	}
	if s.PropWidth != "" {
		styles = styles.Style("width", s.PropWidth)
	}
	if s.PropHeight != "" {
		styles = styles.Style("height", s.PropHeight)
	}
	if s.PropMinWidth != "" {
		styles = styles.Style("min-width", s.PropMinWidth)
	}
	if s.PropMinHeight != "" {
		styles = styles.Style("min-height", s.PropMinHeight)
	}
	if s.PropMaxWidth != "" {
		styles = styles.Style("max-width", s.PropMaxWidth)
	}
	if s.PropMaxHeight != "" {
		styles = styles.Style("max-height", s.PropMaxHeight)
	}
	if s.PropOverflow != "" {
		styles = styles.Style("overflow", s.PropOverflow)
	}
	if s.PropZIndex != "" {
		styles = styles.Style("z-index", s.PropZIndex)
	}
	if s.PropCursor != "" {
		styles = styles.Style("cursor", s.PropCursor)
	}
	if s.PropPointerEvents != "" {
		styles = styles.Style("pointer-events", s.PropPointerEvents)
	}

	// Add children if provided
	if len(s.PropChildren) > 0 {
		styles = styles.Body(s.PropChildren...)
	}

	return styles
}
