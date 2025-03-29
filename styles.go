package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumStyles represents a collection of Spectrum CSS custom properties
type SpectrumStyles struct {
	app.Compo

	// Properties
	backgroundColor string
	textColor       string
	borderColor     string
	fontSize        string
	fontWeight      string
	lineHeight      string
	padding         string
	margin          string
	borderRadius    string
	boxShadow       string
	opacity         string
	transition      string
	transform       string
	display         string
	position        string
	top             string
	right           string
	bottom          string
	left            string
	width           string
	height          string
	minWidth        string
	minHeight       string
	maxWidth        string
	maxHeight       string
	overflow        string
	zIndex          string
	cursor          string
	pointerEvents   string

	// Children
	children []app.UI
}

// NewSpectrumStyles creates a new Spectrum styles component
func NewSpectrumStyles() *SpectrumStyles {
	return &SpectrumStyles{}
}

// BackgroundColor sets the background color
func (s *SpectrumStyles) BackgroundColor(color string) *SpectrumStyles {
	s.backgroundColor = color
	return s
}

// TextColor sets the text color
func (s *SpectrumStyles) TextColor(color string) *SpectrumStyles {
	s.textColor = color
	return s
}

// BorderColor sets the border color
func (s *SpectrumStyles) BorderColor(color string) *SpectrumStyles {
	s.borderColor = color
	return s
}

// FontSize sets the font size
func (s *SpectrumStyles) FontSize(size string) *SpectrumStyles {
	s.fontSize = size
	return s
}

// FontWeight sets the font weight
func (s *SpectrumStyles) FontWeight(weight string) *SpectrumStyles {
	s.fontWeight = weight
	return s
}

// LineHeight sets the line height
func (s *SpectrumStyles) LineHeight(height string) *SpectrumStyles {
	s.lineHeight = height
	return s
}

// Padding sets the padding
func (s *SpectrumStyles) Padding(padding string) *SpectrumStyles {
	s.padding = padding
	return s
}

// Margin sets the margin
func (s *SpectrumStyles) Margin(margin string) *SpectrumStyles {
	s.margin = margin
	return s
}

// BorderRadius sets the border radius
func (s *SpectrumStyles) BorderRadius(radius string) *SpectrumStyles {
	s.borderRadius = radius
	return s
}

// BoxShadow sets the box shadow
func (s *SpectrumStyles) BoxShadow(shadow string) *SpectrumStyles {
	s.boxShadow = shadow
	return s
}

// Opacity sets the opacity
func (s *SpectrumStyles) Opacity(opacity string) *SpectrumStyles {
	s.opacity = opacity
	return s
}

// Transition sets the transition
func (s *SpectrumStyles) Transition(transition string) *SpectrumStyles {
	s.transition = transition
	return s
}

// Transform sets the transform
func (s *SpectrumStyles) Transform(transform string) *SpectrumStyles {
	s.transform = transform
	return s
}

// Display sets the display property
func (s *SpectrumStyles) Display(display string) *SpectrumStyles {
	s.display = display
	return s
}

// Position sets the position
func (s *SpectrumStyles) Position(position string) *SpectrumStyles {
	s.position = position
	return s
}

// Top sets the top position
func (s *SpectrumStyles) Top(top string) *SpectrumStyles {
	s.top = top
	return s
}

// Right sets the right position
func (s *SpectrumStyles) Right(right string) *SpectrumStyles {
	s.right = right
	return s
}

// Bottom sets the bottom position
func (s *SpectrumStyles) Bottom(bottom string) *SpectrumStyles {
	s.bottom = bottom
	return s
}

// Left sets the left position
func (s *SpectrumStyles) Left(left string) *SpectrumStyles {
	s.left = left
	return s
}

// Width sets the width
func (s *SpectrumStyles) Width(width string) *SpectrumStyles {
	s.width = width
	return s
}

// Height sets the height
func (s *SpectrumStyles) Height(height string) *SpectrumStyles {
	s.height = height
	return s
}

// MinWidth sets the minimum width
func (s *SpectrumStyles) MinWidth(width string) *SpectrumStyles {
	s.minWidth = width
	return s
}

// MinHeight sets the minimum height
func (s *SpectrumStyles) MinHeight(height string) *SpectrumStyles {
	s.minHeight = height
	return s
}

// MaxWidth sets the maximum width
func (s *SpectrumStyles) MaxWidth(width string) *SpectrumStyles {
	s.maxWidth = width
	return s
}

// MaxHeight sets the maximum height
func (s *SpectrumStyles) MaxHeight(height string) *SpectrumStyles {
	s.maxHeight = height
	return s
}

// Overflow sets the overflow property
func (s *SpectrumStyles) Overflow(overflow string) *SpectrumStyles {
	s.overflow = overflow
	return s
}

// ZIndex sets the z-index
func (s *SpectrumStyles) ZIndex(index string) *SpectrumStyles {
	s.zIndex = index
	return s
}

// Cursor sets the cursor style
func (s *SpectrumStyles) Cursor(cursor string) *SpectrumStyles {
	s.cursor = cursor
	return s
}

// PointerEvents sets the pointer events
func (s *SpectrumStyles) PointerEvents(events string) *SpectrumStyles {
	s.pointerEvents = events
	return s
}

// Child adds a child element
func (s *SpectrumStyles) Child(child app.UI) *SpectrumStyles {
	s.children = append(s.children, child)
	return s
}

// Children adds multiple child elements
func (s *SpectrumStyles) Children(children ...app.UI) *SpectrumStyles {
	s.children = append(s.children, children...)
	return s
}

// Render renders the Spectrum styles component
func (s *SpectrumStyles) Render() app.UI {
	styles := app.Elem("div")

	// Set styles
	if s.backgroundColor != "" {
		styles = styles.Style("background-color", s.backgroundColor)
	}
	if s.textColor != "" {
		styles = styles.Style("color", s.textColor)
	}
	if s.borderColor != "" {
		styles = styles.Style("border-color", s.borderColor)
	}
	if s.fontSize != "" {
		styles = styles.Style("font-size", s.fontSize)
	}
	if s.fontWeight != "" {
		styles = styles.Style("font-weight", s.fontWeight)
	}
	if s.lineHeight != "" {
		styles = styles.Style("line-height", s.lineHeight)
	}
	if s.padding != "" {
		styles = styles.Style("padding", s.padding)
	}
	if s.margin != "" {
		styles = styles.Style("margin", s.margin)
	}
	if s.borderRadius != "" {
		styles = styles.Style("border-radius", s.borderRadius)
	}
	if s.boxShadow != "" {
		styles = styles.Style("box-shadow", s.boxShadow)
	}
	if s.opacity != "" {
		styles = styles.Style("opacity", s.opacity)
	}
	if s.transition != "" {
		styles = styles.Style("transition", s.transition)
	}
	if s.transform != "" {
		styles = styles.Style("transform", s.transform)
	}
	if s.display != "" {
		styles = styles.Style("display", s.display)
	}
	if s.position != "" {
		styles = styles.Style("position", s.position)
	}
	if s.top != "" {
		styles = styles.Style("top", s.top)
	}
	if s.right != "" {
		styles = styles.Style("right", s.right)
	}
	if s.bottom != "" {
		styles = styles.Style("bottom", s.bottom)
	}
	if s.left != "" {
		styles = styles.Style("left", s.left)
	}
	if s.width != "" {
		styles = styles.Style("width", s.width)
	}
	if s.height != "" {
		styles = styles.Style("height", s.height)
	}
	if s.minWidth != "" {
		styles = styles.Style("min-width", s.minWidth)
	}
	if s.minHeight != "" {
		styles = styles.Style("min-height", s.minHeight)
	}
	if s.maxWidth != "" {
		styles = styles.Style("max-width", s.maxWidth)
	}
	if s.maxHeight != "" {
		styles = styles.Style("max-height", s.maxHeight)
	}
	if s.overflow != "" {
		styles = styles.Style("overflow", s.overflow)
	}
	if s.zIndex != "" {
		styles = styles.Style("z-index", s.zIndex)
	}
	if s.cursor != "" {
		styles = styles.Style("cursor", s.cursor)
	}
	if s.pointerEvents != "" {
		styles = styles.Style("pointer-events", s.pointerEvents)
	}

	// Add children if provided
	if len(s.children) > 0 {
		styles = styles.Body(s.children...)
	}

	return styles
}
