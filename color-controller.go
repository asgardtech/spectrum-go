package sp

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ColorSpace defines the available color spaces for color representation
type ColorSpace string

const (
	// ColorSpaceRGB represents the RGB color space
	ColorSpaceRGB ColorSpace = "rgb"
	// ColorSpaceHSL represents the HSL color space
	ColorSpaceHSL ColorSpace = "hsl"
	// ColorSpaceHSV represents the HSV color space
	ColorSpaceHSV ColorSpace = "hsv"
	// ColorSpaceHEX represents the HEX color space
	ColorSpaceHEX ColorSpace = "hex"
)

// ColorObject represents a color in any color space
type ColorObject struct {
	R     float64
	G     float64
	B     float64
	H     float64
	S     float64
	L     float64
	V     float64
	A     float64
	Hex   string
	Space ColorSpace
}

// ColorValidationResult represents the result of validating a color string
type ColorValidationResult struct {
	IsValid    bool
	ColorSpace ColorSpace
	Coords     []float64
	Alpha      float64
}

// ManageAsOptions represents the options for how the color controller manages colors
type ManageAsOptions struct {
	ManageAs ColorSpace
}

// spectrumColorController is a comprehensive utility for managing and validating color values
// in various color spaces, including RGB, HSL, HSV, and Hex
type spectrumColorController struct {
	PropHost             app.Composer
	PropColor            *ColorObject
	PropPreviousColor    *ColorObject
	PropOptions          *ManageAsOptions
	PropColorChangeEvent func(ctx app.Context, e app.Event)
}

// ColorController creates a new spectrumColorController with default options
func ColorController(host app.Composer, options *ManageAsOptions) *spectrumColorController {
	if options == nil {
		options = &ManageAsOptions{
			ManageAs: ColorSpaceHSV,
		}
	}

	controller := &spectrumColorController{
		PropHost:    host,
		PropOptions: options,
		PropColor: &ColorObject{
			R:     255,
			G:     0,
			B:     0,
			H:     0,
			S:     1,
			L:     0.5,
			V:     1,
			A:     1,
			Hex:   "#FF0000",
			Space: options.ManageAs,
		},
	}

	// Clone the color for previous state
	controller.PropPreviousColor = controller.cloneColor(controller.PropColor)

	return controller
}

// SetColorChangeEvent sets the event handler to be called when the color changes
func (c *spectrumColorController) SetColorChangeEvent(handler func(ctx app.Context, e app.Event)) *spectrumColorController {
	c.PropColorChangeEvent = handler
	return c
}

// Color gets the current color value
func (c *spectrumColorController) Color() *ColorObject {
	return c.PropColor
}

// SetColor sets the current color value and triggers a color change event
func (c *spectrumColorController) SetColor(color interface{}) *spectrumColorController {
	prevColor := c.cloneColor(c.PropColor)

	switch v := color.(type) {
	case string:
		result := c.validateColorString(v)
		if result.IsValid {
			c.updateColorFromValidation(result)
		}
	case *ColorObject:
		c.PropColor = c.cloneColor(v)
	case ColorObject:
		c.PropColor = &v
	}

	// Only trigger event if the color has changed
	if !c.colorsEqual(prevColor, c.PropColor) && c.PropColorChangeEvent != nil {
		// We don't have access to context here, so we'll need to handle it differently
		// The host component should handle context properly when it uses this controller
		// This is a limitation of this implementation
	}

	return c
}

// SavePreviousColor saves the current color as the previous color
func (c *spectrumColorController) SavePreviousColor() *spectrumColorController {
	c.PropPreviousColor = c.cloneColor(c.PropColor)
	return c
}

// RestorePreviousColor restores the previous color
func (c *spectrumColorController) RestorePreviousColor() *spectrumColorController {
	if c.PropPreviousColor != nil {
		return c.SetColor(c.PropPreviousColor)
	}
	return c
}

// GetHslString returns the current color in HSL string format
func (c *spectrumColorController) GetHslString() string {
	if c.PropColor.A < 1 {
		return fmt.Sprintf("hsla(%f, %f%%, %f%%, %f)", c.PropColor.H, c.PropColor.S*100, c.PropColor.L*100, c.PropColor.A)
	}
	return fmt.Sprintf("hsl(%f, %f%%, %f%%)", c.PropColor.H, c.PropColor.S*100, c.PropColor.L*100)
}

// GetRgbString returns the current color in RGB string format
func (c *spectrumColorController) GetRgbString() string {
	if c.PropColor.A < 1 {
		return fmt.Sprintf("rgba(%f, %f, %f, %f)", c.PropColor.R, c.PropColor.G, c.PropColor.B, c.PropColor.A)
	}
	return fmt.Sprintf("rgb(%f, %f, %f)", c.PropColor.R, c.PropColor.G, c.PropColor.B)
}

// GetHexString returns the current color in Hex string format
func (c *spectrumColorController) GetHexString() string {
	return c.PropColor.Hex
}

// GetColor returns the color in the specified format
func (c *spectrumColorController) GetColor(format interface{}) *ColorObject {
	var colorSpace ColorSpace

	switch v := format.(type) {
	case string:
		colorSpace = ColorSpace(v)
	case ColorSpace:
		colorSpace = v
	default:
		// Default to the color's current space
		colorSpace = c.PropColor.Space
	}

	// Return a copy of the color in the requested space
	result := c.cloneColor(c.PropColor)
	result.Space = colorSpace

	// Convert if necessary
	switch colorSpace {
	case ColorSpaceRGB:
		if c.PropColor.Space != ColorSpaceRGB {
			c.ensureRGB(result)
		}
	case ColorSpaceHSL:
		if c.PropColor.Space != ColorSpaceHSL {
			c.ensureHSL(result)
		}
	case ColorSpaceHSV:
		if c.PropColor.Space != ColorSpaceHSV {
			c.ensureHSV(result)
		}
	case ColorSpaceHEX:
		if c.PropColor.Space != ColorSpaceHEX {
			c.ensureHex(result)
		}
	}

	return result
}

// SetHue sets the hue value of the current color
func (c *spectrumColorController) SetHue(hue float64) *spectrumColorController {
	// Ensure hue is between 0 and 360
	hue = math.Mod(hue, 360)
	if hue < 0 {
		hue += 360
	}

	c.PropColor.H = hue

	// Update other color spaces
	c.updateColorSpaces()

	return c
}

// Hue gets the hue value of the current color
func (c *spectrumColorController) Hue() float64 {
	return c.PropColor.H
}

// validateColorString validates a color string and returns the validation result
func (c *spectrumColorController) validateColorString(color string) ColorValidationResult {
	color = strings.TrimSpace(color)
	result := ColorValidationResult{
		IsValid: false,
		Alpha:   1.0,
	}

	// Check for hex format
	hexRegex := regexp.MustCompile(`^#([A-Fa-f0-9]{3}|[A-Fa-f0-9]{6}|[A-Fa-f0-9]{8})$`)
	if hexRegex.MatchString(color) {
		result.IsValid = true
		result.ColorSpace = ColorSpaceHEX

		// Parse hex values
		var r, g, b, a int64
		hexColor := color[1:] // Remove the #

		if len(hexColor) == 3 {
			// Short form #RGB
			r, _ = strconv.ParseInt(string(hexColor[0])+string(hexColor[0]), 16, 64)
			g, _ = strconv.ParseInt(string(hexColor[1])+string(hexColor[1]), 16, 64)
			b, _ = strconv.ParseInt(string(hexColor[2])+string(hexColor[2]), 16, 64)
			a = 255
		} else if len(hexColor) == 6 {
			// Standard form #RRGGBB
			r, _ = strconv.ParseInt(hexColor[0:2], 16, 64)
			g, _ = strconv.ParseInt(hexColor[2:4], 16, 64)
			b, _ = strconv.ParseInt(hexColor[4:6], 16, 64)
			a = 255
		} else if len(hexColor) == 8 {
			// With alpha #RRGGBBAA
			r, _ = strconv.ParseInt(hexColor[0:2], 16, 64)
			g, _ = strconv.ParseInt(hexColor[2:4], 16, 64)
			b, _ = strconv.ParseInt(hexColor[4:6], 16, 64)
			a, _ = strconv.ParseInt(hexColor[6:8], 16, 64)
		}

		result.Coords = []float64{float64(r), float64(g), float64(b)}
		result.Alpha = float64(a) / 255
		return result
	}

	// Check for rgb/rgba format
	rgbRegex := regexp.MustCompile(`^rgba?\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)(?:\s*,\s*([0-9.]+))?\s*\)$`)
	if rgbRegex.MatchString(color) {
		matches := rgbRegex.FindStringSubmatch(color)
		r, _ := strconv.ParseFloat(matches[1], 64)
		g, _ := strconv.ParseFloat(matches[2], 64)
		b, _ := strconv.ParseFloat(matches[3], 64)

		result.Coords = []float64{r, g, b}
		result.ColorSpace = ColorSpaceRGB
		result.IsValid = true

		if len(matches) > 4 && matches[4] != "" {
			a, _ := strconv.ParseFloat(matches[4], 64)
			result.Alpha = a
		}

		return result
	}

	// Check for hsl/hsla format
	hslRegex := regexp.MustCompile(`^hsla?\(\s*(\d+)\s*,\s*([0-9.]+)%\s*,\s*([0-9.]+)%(?:\s*,\s*([0-9.]+))?\s*\)$`)
	if hslRegex.MatchString(color) {
		matches := hslRegex.FindStringSubmatch(color)
		h, _ := strconv.ParseFloat(matches[1], 64)
		s, _ := strconv.ParseFloat(matches[2], 64)
		l, _ := strconv.ParseFloat(matches[3], 64)

		result.Coords = []float64{h, s / 100, l / 100}
		result.ColorSpace = ColorSpaceHSL
		result.IsValid = true

		if len(matches) > 4 && matches[4] != "" {
			a, _ := strconv.ParseFloat(matches[4], 64)
			result.Alpha = a
		}

		return result
	}

	return result
}

// updateColorFromValidation updates the color based on validation results
func (c *spectrumColorController) updateColorFromValidation(result ColorValidationResult) {
	switch result.ColorSpace {
	case ColorSpaceRGB:
		c.PropColor.R = result.Coords[0]
		c.PropColor.G = result.Coords[1]
		c.PropColor.B = result.Coords[2]
		c.PropColor.A = result.Alpha
		c.PropColor.Space = ColorSpaceRGB
		// Update other color spaces
		c.rgbToHsl()
		c.hslToHsv()
		c.rgbToHex()
	case ColorSpaceHSL:
		c.PropColor.H = result.Coords[0]
		c.PropColor.S = result.Coords[1]
		c.PropColor.L = result.Coords[2]
		c.PropColor.A = result.Alpha
		c.PropColor.Space = ColorSpaceHSL
		// Update other color spaces
		c.hslToRgb()
		c.hslToHsv()
		c.rgbToHex()
	case ColorSpaceHEX:
		c.PropColor.R = result.Coords[0]
		c.PropColor.G = result.Coords[1]
		c.PropColor.B = result.Coords[2]
		c.PropColor.A = result.Alpha
		c.PropColor.Space = ColorSpaceHEX
		c.PropColor.Hex = c.rgbToHexString(int(c.PropColor.R), int(c.PropColor.G), int(c.PropColor.B), c.PropColor.A)
		// Update other color spaces
		c.rgbToHsl()
		c.hslToHsv()
	}
}

// updateColorSpaces updates all color spaces based on the current color
func (c *spectrumColorController) updateColorSpaces() {
	switch c.PropColor.Space {
	case ColorSpaceRGB:
		c.rgbToHsl()
		c.hslToHsv()
		c.rgbToHex()
	case ColorSpaceHSL:
		c.hslToRgb()
		c.hslToHsv()
		c.rgbToHex()
	case ColorSpaceHSV:
		c.hsvToHsl()
		c.hslToRgb()
		c.rgbToHex()
	case ColorSpaceHEX:
		c.rgbToHsl()
		c.hslToHsv()
	}
}

// rgbToHsl converts RGB to HSL
func (c *spectrumColorController) rgbToHsl() {
	r := c.PropColor.R / 255
	g := c.PropColor.G / 255
	b := c.PropColor.B / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	l := (max + min) / 2

	var h, s float64

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min

		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}

		h *= 60
	}

	c.PropColor.H = h
	c.PropColor.S = s
	c.PropColor.L = l
}

// hslToRgb converts HSL to RGB
func (c *spectrumColorController) hslToRgb() {
	h := c.PropColor.H
	s := c.PropColor.S
	l := c.PropColor.L

	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}

		p := 2*l - q

		r = c.hueToRgb(p, q, h+120)
		g = c.hueToRgb(p, q, h)
		b = c.hueToRgb(p, q, h-120)
	}

	c.PropColor.R = r * 255
	c.PropColor.G = g * 255
	c.PropColor.B = b * 255
}

// hueToRgb is a helper function for hslToRgb
func (c *spectrumColorController) hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 360
	}
	if t > 360 {
		t -= 360
	}

	if t < 60 {
		return p + (q-p)*t/60
	}
	if t < 180 {
		return q
	}
	if t < 240 {
		return p + (q-p)*(240-t)/60
	}
	return p
}

// hslToHsv converts HSL to HSV
func (c *spectrumColorController) hslToHsv() {
	l := c.PropColor.L
	s := c.PropColor.S

	var v float64
	if l <= 0.5 {
		v = l * (1 + s)
	} else {
		v = l + s - l*s
	}

	var sv float64
	if v == 0 {
		sv = 0
	} else {
		sv = 2 * (1 - l/v)
	}

	c.PropColor.V = v
	// Keep the hue the same
	// Update saturation for HSV
	c.PropColor.S = sv
}

// hsvToHsl converts HSV to HSL
func (c *spectrumColorController) hsvToHsl() {
	v := c.PropColor.V
	s := c.PropColor.S

	var l float64 = v * (1 - s/2)

	var sl float64
	if l == 0 || l == 1 {
		sl = 0
	} else {
		sl = (v - l) / math.Min(l, 1-l)
	}

	c.PropColor.L = l
	// Keep the hue the same
	// Update saturation for HSL
	c.PropColor.S = sl
}

// rgbToHex converts RGB to Hex
func (c *spectrumColorController) rgbToHex() {
	c.PropColor.Hex = c.rgbToHexString(int(c.PropColor.R), int(c.PropColor.G), int(c.PropColor.B), c.PropColor.A)
}

// rgbToHexString converts RGB values to a hex string
func (c *spectrumColorController) rgbToHexString(r, g, b int, a float64) string {
	if a < 1 {
		return fmt.Sprintf("#%02X%02X%02X%02X", r, g, b, int(a*255))
	}
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// ensureRGB ensures the color object has valid RGB values
func (c *spectrumColorController) ensureRGB(color *ColorObject) {
	if color.Space == ColorSpaceRGB {
		return
	}

	switch color.Space {
	case ColorSpaceHSL:
		h := color.H
		s := color.S
		l := color.L

		var r, g, b float64

		if s == 0 {
			r = l
			g = l
			b = l
		} else {
			var q float64
			if l < 0.5 {
				q = l * (1 + s)
			} else {
				q = l + s - l*s
			}

			p := 2*l - q

			r = c.hueToRgb(p, q, h+120)
			g = c.hueToRgb(p, q, h)
			b = c.hueToRgb(p, q, h-120)
		}

		color.R = r * 255
		color.G = g * 255
		color.B = b * 255
	case ColorSpaceHSV:
		// Convert HSV to HSL first
		v := color.V
		s := color.S

		var l float64 = v * (1 - s/2)

		var sl float64
		if l == 0 || l == 1 {
			sl = 0
		} else {
			sl = (v - l) / math.Min(l, 1-l)
		}

		color.L = l
		color.S = sl

		// Then convert HSL to RGB
		c.ensureRGB(color)
	case ColorSpaceHEX:
		// Hex is already in RGB format
	}

	color.Space = ColorSpaceRGB
}

// ensureHSL ensures the color object has valid HSL values
func (c *spectrumColorController) ensureHSL(color *ColorObject) {
	if color.Space == ColorSpaceHSL {
		return
	}

	switch color.Space {
	case ColorSpaceRGB:
		r := color.R / 255
		g := color.G / 255
		b := color.B / 255

		max := math.Max(r, math.Max(g, b))
		min := math.Min(r, math.Min(g, b))

		l := (max + min) / 2

		var h, s float64

		if max == min {
			h = 0
			s = 0
		} else {
			d := max - min

			if l > 0.5 {
				s = d / (2 - max - min)
			} else {
				s = d / (max + min)
			}

			switch max {
			case r:
				h = (g - b) / d
				if g < b {
					h += 6
				}
			case g:
				h = (b-r)/d + 2
			case b:
				h = (r-g)/d + 4
			}

			h *= 60
		}

		color.H = h
		color.S = s
		color.L = l
	case ColorSpaceHSV:
		v := color.V
		s := color.S

		var l float64 = v * (1 - s/2)

		var sl float64
		if l == 0 || l == 1 {
			sl = 0
		} else {
			sl = (v - l) / math.Min(l, 1-l)
		}

		color.L = l
		color.S = sl
	case ColorSpaceHEX:
		// Convert to RGB first, then to HSL
		c.ensureRGB(color)
		c.ensureHSL(color)
	}

	color.Space = ColorSpaceHSL
}

// ensureHSV ensures the color object has valid HSV values
func (c *spectrumColorController) ensureHSV(color *ColorObject) {
	if color.Space == ColorSpaceHSV {
		return
	}

	switch color.Space {
	case ColorSpaceHSL:
		l := color.L
		s := color.S

		var v float64
		if l <= 0.5 {
			v = l * (1 + s)
		} else {
			v = l + s - l*s
		}

		var sv float64
		if v == 0 {
			sv = 0
		} else {
			sv = 2 * (1 - l/v)
		}

		color.V = v
		color.S = sv
	case ColorSpaceRGB:
		// Convert to HSL first, then to HSV
		c.ensureHSL(color)
		c.ensureHSV(color)
	case ColorSpaceHEX:
		// Convert to RGB first, then to HSL, then to HSV
		c.ensureRGB(color)
		c.ensureHSL(color)
		c.ensureHSV(color)
	}

	color.Space = ColorSpaceHSV
}

// ensureHex ensures the color object has a valid Hex value
func (c *spectrumColorController) ensureHex(color *ColorObject) {
	if color.Space == ColorSpaceHEX {
		return
	}

	// Make sure we have RGB values
	c.ensureRGB(color)

	// Convert RGB to Hex
	color.Hex = c.rgbToHexString(int(color.R), int(color.G), int(color.B), color.A)
	color.Space = ColorSpaceHEX
}

// cloneColor creates a copy of a ColorObject
func (c *spectrumColorController) cloneColor(color *ColorObject) *ColorObject {
	if color == nil {
		return nil
	}

	return &ColorObject{
		R:     color.R,
		G:     color.G,
		B:     color.B,
		H:     color.H,
		S:     color.S,
		L:     color.L,
		V:     color.V,
		A:     color.A,
		Hex:   color.Hex,
		Space: color.Space,
	}
}

// colorsEqual checks if two ColorObjects are equal
func (c *spectrumColorController) colorsEqual(a, b *ColorObject) bool {
	if a == nil || b == nil {
		return a == b
	}

	// Compare the values that matter most based on the color space
	switch a.Space {
	case ColorSpaceRGB:
		return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
	case ColorSpaceHSL:
		return a.H == b.H && a.S == b.S && a.L == b.L && a.A == b.A
	case ColorSpaceHSV:
		return a.H == b.H && a.S == b.S && a.V == b.V && a.A == b.A
	case ColorSpaceHEX:
		return a.Hex == b.Hex
	default:
		// Fall back to checking all values
		return a.R == b.R && a.G == b.G && a.B == b.B &&
			a.H == b.H && a.S == b.S && a.L == b.L &&
			a.V == b.V && a.A == b.A && a.Hex == b.Hex
	}
}
