package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TableSize represents the Size variant for the table
type TableSize string

// TableSize values
const (
    TableSizeS TableSize = "s"
    TableSizeM TableSize = "m"
    TableSizeL TableSize = "l"
)

// spectrumTable represents an sp-table component
type spectrumTable struct {
    app.Compo
    *styler[*spectrumTable]

    // Properties
    // Size variant for the table
    PropSize TableSize
    // Whether the table has a reduced visual style
    PropQuiet bool
    // Whether the table has reduced spacing between rows
    PropDensely bool
    // Whether the table has alternating row backgrounds
    PropStriped bool
    // Whether the table header has enhanced visual prominence
    PropEmphasized bool

    // Text content for the default slot
    PropText string

    // Content slots


}

// Table creates a new sp-table component
func Table() *spectrumTable {
    element := &spectrumTable{
        PropSize: TableSizeM,
        PropQuiet: false,
        PropDensely: false,
        PropStriped: false,
        PropEmphasized: false,
    }

    element.styler = newStyler(element)

    return element
}

// Size variant for the table
func (c *spectrumTable) Size(size TableSize) *spectrumTable {
    c.PropSize = size
    return c
}

func (c *spectrumTable) SizeS() *spectrumTable {
    return c.Size(TableSizeS)
}
func (c *spectrumTable) SizeM() *spectrumTable {
    return c.Size(TableSizeM)
}
func (c *spectrumTable) SizeL() *spectrumTable {
    return c.Size(TableSizeL)
}
// Whether the table has a reduced visual style
func (c *spectrumTable) Quiet(quiet bool) *spectrumTable {
    c.PropQuiet = quiet
    return c
}

func (c *spectrumTable) SetQuiet() *spectrumTable {
    return c.Quiet(true)
}

// Whether the table has reduced spacing between rows
func (c *spectrumTable) Densely(densely bool) *spectrumTable {
    c.PropDensely = densely
    return c
}

func (c *spectrumTable) SetDensely() *spectrumTable {
    return c.Densely(true)
}

// Whether the table has alternating row backgrounds
func (c *spectrumTable) Striped(striped bool) *spectrumTable {
    c.PropStriped = striped
    return c
}

func (c *spectrumTable) SetStriped() *spectrumTable {
    return c.Striped(true)
}

// Whether the table header has enhanced visual prominence
func (c *spectrumTable) Emphasized(emphasized bool) *spectrumTable {
    c.PropEmphasized = emphasized
    return c
}

func (c *spectrumTable) SetEmphasized() *spectrumTable {
    return c.Emphasized(true)
}


// Text sets the text content for the default slot
func (c *spectrumTable) Text(text string) *spectrumTable {
    c.PropText = text
    return c
}





// Render renders the sp-table component
func (c *spectrumTable) Render() app.UI {
    element := app.Elem("sp-table")

    // Set attributes
    if c.PropSize != "" {
        element = element.Attr("size", string(c.PropSize))
    }
    if c.PropQuiet {
        element = element.Attr("quiet", true)
    }
    if c.PropDensely {
        element = element.Attr("densely", true)
    }
    if c.PropStriped {
        element = element.Attr("striped", true)
    }
    if c.PropEmphasized {
        element = element.Attr("emphasized", true)
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