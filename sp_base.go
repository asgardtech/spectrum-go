package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// BaseDir represents the Sets the text direction. When not specified, inherits from document or containing sp-theme
type BaseDir string

// BaseDir values
const (
    BaseDirLtr BaseDir = "ltr"
    BaseDirRtl BaseDir = "rtl"
)

// spectrumBase represents an  component
type spectrumBase struct {
    app.Compo
    *styler[*spectrumBase]

    // Properties
    // Sets the text direction. When not specified, inherits from document or containing sp-theme
    PropDir BaseDir




}

// Base creates a new  component
func Base() *spectrumBase {
    element := &spectrumBase{
        PropDir: BaseDirLtr,
    }

    element.styler = newStyler(element)

    return element
}

// Sets the text direction. When not specified, inherits from document or containing sp-theme
func (c *spectrumBase) Dir(dir BaseDir) *spectrumBase {
    c.PropDir = dir
    return c
}

func (c *spectrumBase) DirLtr() *spectrumBase {
    return c.Dir(BaseDirLtr)
}
func (c *spectrumBase) DirRtl() *spectrumBase {
    return c.Dir(BaseDirRtl)
}





// Render renders the  component
func (c *spectrumBase) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropDir != "" {
        element = element.Attr("dir", string(c.PropDir))
    }



    element = element.Styles(c.styler.styles)

    return element
} 