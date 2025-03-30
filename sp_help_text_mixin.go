package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"


// spectrumHelpTextMixin represents an  component
type spectrumHelpTextMixin struct {
    app.Compo
    *styler[*spectrumHelpTextMixin]

    // Properties
    // The id attribute of the associated sp-help-text
    PropHelptextid string
    // A method that returns a TemplateResult with the help-text and negative-help-text slots
    PropRenderhelptext string


    // Content slots
    PropHelpTextSlot app.UI
    PropNegativeHelpTextSlot app.UI


}

// HelpTextMixin creates a new  component
func HelpTextMixin() *spectrumHelpTextMixin {
    element := &spectrumHelpTextMixin{
    }

    element.styler = newStyler(element)

    return element
}

// The id attribute of the associated sp-help-text
func (c *spectrumHelpTextMixin) Helptextid(helpTextId string) *spectrumHelpTextMixin {
    c.PropHelptextid = helpTextId
    return c
}

// A method that returns a TemplateResult with the help-text and negative-help-text slots
func (c *spectrumHelpTextMixin) Renderhelptext(renderHelpText string) *spectrumHelpTextMixin {
    c.PropRenderhelptext = renderHelpText
    return c
}



// Default or non-negative help text to associate to your form element
func (c *spectrumHelpTextMixin) HelpText(content app.UI) *spectrumHelpTextMixin {
    c.PropHelpTextSlot = content

    return c
}

// Negative help text to associate to your form element when invalid
func (c *spectrumHelpTextMixin) NegativeHelpText(content app.UI) *spectrumHelpTextMixin {
    c.PropNegativeHelpTextSlot = content

    return c
}




// Render renders the  component
func (c *spectrumHelpTextMixin) Render() app.UI {
    element := app.Elem("")

    // Set attributes
    if c.PropHelptextid != "" {
        element = element.Attr("helpTextId", c.PropHelptextid)
    }
    if c.PropRenderhelptext != "" {
        element = element.Attr("renderHelpText", c.PropRenderhelptext)
    }


    // Add slots and children
    slotElements := []app.UI{}


    // Add help-text slot
    if c.PropHelpTextSlot != nil {
        slotElem := c.PropHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "help-text").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }
    // Add negative-help-text slot
    if c.PropNegativeHelpTextSlot != nil {
        slotElem := c.PropNegativeHelpTextSlot
        if slotWithAttr, ok := slotElem.(interface{ Slot(string) app.UI }); ok {
            slotElem = slotWithAttr.Slot("negative-help-text")
        } else {
            slotElem = app.Elem("div").
                Attr("slot", "negative-help-text").
                Body(slotElem)
        }
        slotElements = append(slotElements, slotElem)
    }


    // Add all elements to the component
    if len(slotElements) > 0 {
        element = element.Body(slotElements...)
    }

    element = element.Styles(c.styler.styles)

    return element
} 