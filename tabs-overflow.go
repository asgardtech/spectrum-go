package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TabsOverflowSide represents the side of the tabs where overflow is positioned
type TabsOverflowSide string

// Tabs overflow sides
const (
	TabsOverflowSideStart TabsOverflowSide = "start"
	TabsOverflowSideEnd   TabsOverflowSide = "end"
)

// spectrumTabsOverflow represents a component that handles overflow of tabs
type spectrumTabsOverflow struct {
	app.Compo

	// Properties
	PropSide            TabsOverflowSide
	PropCompact         bool
	PropDisabled        bool
	PropLabelPrevious   string
	PropLabelNext       string
	PropLabelSelect     string
	PropOverflowOnClick func(ctx app.Context, e app.Event)
	PropPreviousOnClick func(ctx app.Context, e app.Event)
	PropNextOnClick     func(ctx app.Context, e app.Event)
}

// TabsOverflow creates a new tabs overflow component
func TabsOverflow() *spectrumTabsOverflow {
	return &spectrumTabsOverflow{
		PropSide: TabsOverflowSideEnd,
	}
}

// Side sets the side where the overflow is positioned
func (o *spectrumTabsOverflow) Side(side TabsOverflowSide) *spectrumTabsOverflow {
	o.PropSide = side
	return o
}

// Compact sets whether the overflow is displayed compactly
func (o *spectrumTabsOverflow) Compact(compact bool) *spectrumTabsOverflow {
	o.PropCompact = compact
	return o
}

// Disabled sets whether the overflow is disabled
func (o *spectrumTabsOverflow) Disabled(disabled bool) *spectrumTabsOverflow {
	o.PropDisabled = disabled
	return o
}

// LabelPrevious sets the label for the previous button
func (o *spectrumTabsOverflow) LabelPrevious(label string) *spectrumTabsOverflow {
	o.PropLabelPrevious = label
	return o
}

// LabelNext sets the label for the next button
func (o *spectrumTabsOverflow) LabelNext(label string) *spectrumTabsOverflow {
	o.PropLabelNext = label
	return o
}

// LabelSelect sets the label for the select button
func (o *spectrumTabsOverflow) LabelSelect(label string) *spectrumTabsOverflow {
	o.PropLabelSelect = label
	return o
}

// OnOverflow sets the handler for overflow click events
func (o *spectrumTabsOverflow) OnOverflow(handler func(ctx app.Context, e app.Event)) *spectrumTabsOverflow {
	o.PropOverflowOnClick = handler
	return o
}

// OnPrevious sets the handler for previous button click events
func (o *spectrumTabsOverflow) OnPrevious(handler func(ctx app.Context, e app.Event)) *spectrumTabsOverflow {
	o.PropPreviousOnClick = handler
	return o
}

// OnNext sets the handler for next button click events
func (o *spectrumTabsOverflow) OnNext(handler func(ctx app.Context, e app.Event)) *spectrumTabsOverflow {
	o.PropNextOnClick = handler
	return o
}

// Render renders the tabs overflow component
func (o *spectrumTabsOverflow) Render() app.UI {
	overflow := app.Elem("sp-tabs-overflow")

	if o.PropSide != "" {
		overflow = overflow.Attr("side", string(o.PropSide))
	}
	if o.PropCompact {
		overflow = overflow.Attr("compact", "")
	}
	if o.PropDisabled {
		overflow = overflow.Attr("disabled", "")
	}
	if o.PropLabelPrevious != "" {
		overflow = overflow.Attr("label-previous", o.PropLabelPrevious)
	}
	if o.PropLabelNext != "" {
		overflow = overflow.Attr("label-next", o.PropLabelNext)
	}
	if o.PropLabelSelect != "" {
		overflow = overflow.Attr("label-select", o.PropLabelSelect)
	}
	if o.PropOverflowOnClick != nil {
		overflow = overflow.On("overflow", o.PropOverflowOnClick)
	}
	if o.PropPreviousOnClick != nil {
		overflow = overflow.On("previous", o.PropPreviousOnClick)
	}
	if o.PropNextOnClick != nil {
		overflow = overflow.On("next", o.PropNextOnClick)
	}

	return overflow
}
