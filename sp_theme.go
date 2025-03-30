package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ThemeColor represents the The color theme to apply
type ThemeColor string

// ThemeColor values
const (
    ThemeColorLightest ThemeColor = "lightest"
    ThemeColorLight ThemeColor = "light"
    ThemeColorDark ThemeColor = "dark"
    ThemeColorDarkest ThemeColor = "darkest"
)
// ThemeScale represents the The scale to apply (medium for desktop, large for mobile)
type ThemeScale string

// ThemeScale values
const (
    ThemeScaleMedium ThemeScale = "medium"
    ThemeScaleLarge ThemeScale = "large"
)
// ThemeSystem represents the The design system variant to apply
type ThemeSystem string

// ThemeSystem values
const (
    ThemeSystemSpectrum ThemeSystem = "spectrum"
    ThemeSystemExpress ThemeSystem = "express"
    ThemeSystemSpectrumTwo ThemeSystem = "spectrum-two"
)
// ThemeDir represents the The text direction to apply
type ThemeDir string

// ThemeDir values
const (
    ThemeDirLtr ThemeDir = "ltr"
    ThemeDirRtl ThemeDir = "rtl"
)

// spectrumTheme represents an sp-theme component
type spectrumTheme struct {
    app.Compo
    *styler[*spectrumTheme]

    // Properties
    // The color theme to apply
    PropColor ThemeColor
    // The scale to apply (medium for desktop, large for mobile)
    PropScale ThemeScale
    // The design system variant to apply
    PropSystem ThemeSystem
    // The text direction to apply
    PropDir ThemeDir
    // The language code to apply for localization
    PropLang string

    // Text content for the default slot
    PropText string

    // Content slots


}

// Theme creates a new sp-theme component
func Theme() *spectrumTheme {
    element := &spectrumTheme{
        PropColor: ThemeColorLight,
        PropScale: ThemeScaleMedium,
        PropSystem: ThemeSystemSpectrum,
        PropDir: ThemeDirLtr,
        PropLang: "",
    }

    element.styler = newStyler(element)

    return element
}

// The color theme to apply
func (c *spectrumTheme) Color(color ThemeColor) *spectrumTheme {
    c.PropColor = color
    return c
}

func (c *spectrumTheme) ColorLightest() *spectrumTheme {
    return c.Color(ThemeColorLightest)
}
func (c *spectrumTheme) ColorLight() *spectrumTheme {
    return c.Color(ThemeColorLight)
}
func (c *spectrumTheme) ColorDark() *spectrumTheme {
    return c.Color(ThemeColorDark)
}
func (c *spectrumTheme) ColorDarkest() *spectrumTheme {
    return c.Color(ThemeColorDarkest)
}
// The scale to apply (medium for desktop, large for mobile)
func (c *spectrumTheme) Scale(scale ThemeScale) *spectrumTheme {
    c.PropScale = scale
    return c
}

func (c *spectrumTheme) ScaleMedium() *spectrumTheme {
    return c.Scale(ThemeScaleMedium)
}
func (c *spectrumTheme) ScaleLarge() *spectrumTheme {
    return c.Scale(ThemeScaleLarge)
}
// The design system variant to apply
func (c *spectrumTheme) System(system ThemeSystem) *spectrumTheme {
    c.PropSystem = system
    return c
}

func (c *spectrumTheme) SystemSpectrum() *spectrumTheme {
    return c.System(ThemeSystemSpectrum)
}
func (c *spectrumTheme) SystemExpress() *spectrumTheme {
    return c.System(ThemeSystemExpress)
}
func (c *spectrumTheme) SystemSpectrumTwo() *spectrumTheme {
    return c.System(ThemeSystemSpectrumTwo)
}
// The text direction to apply
func (c *spectrumTheme) Dir(dir ThemeDir) *spectrumTheme {
    c.PropDir = dir
    return c
}

func (c *spectrumTheme) DirLtr() *spectrumTheme {
    return c.Dir(ThemeDirLtr)
}
func (c *spectrumTheme) DirRtl() *spectrumTheme {
    return c.Dir(ThemeDirRtl)
}
// The language code to apply for localization
func (c *spectrumTheme) Lang(lang string) *spectrumTheme {
    c.PropLang = lang
    return c
}


// Text sets the text content for the default slot
func (c *spectrumTheme) Text(text string) *spectrumTheme {
    c.PropText = text
    return c
}





// Render renders the sp-theme component
func (c *spectrumTheme) Render() app.UI {
    element := app.Elem("sp-theme")

    // Set attributes
    if c.PropColor != "" {
        element = element.Attr("color", string(c.PropColor))
    }
    if c.PropScale != "" {
        element = element.Attr("scale", string(c.PropScale))
    }
    if c.PropSystem != "" {
        element = element.Attr("system", string(c.PropSystem))
    }
    if c.PropDir != "" {
        element = element.Attr("dir", string(c.PropDir))
    }
    if c.PropLang != "" {
        element = element.Attr("lang", c.PropLang)
    }


    // Add slots and children
    slotElements := []app.UI{}

    // Add text content for default slot if specified
    if c.PropText != "" {
        slotElements = append(slotElements, app.Text(c.PropText))
    }



    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 