package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumSplitView represents an sp-split-view component
type spectrumSplitView struct {
	app.Compo

	// Properties
	PropCollapsible  bool
	PropLabel        string
	PropPrimaryMax   string
	PropPrimaryMin   string
	PropPrimarySize  string
	PropResizable    bool
	PropSecondaryMax string
	PropSecondaryMin string
	PropSplitterPos  int
	PropVertical     bool

	// Child components
	PropPrimaryPanel   app.UI
	PropSecondaryPanel app.UI

	// Event handlers
	PropOnChange app.EventHandler
}

// SplitView creates a new split view component
func SplitView() *spectrumSplitView {
	return &spectrumSplitView{
		PropLabel: "Resize the panels", // Default label
	}
}

// Collapsible sets whether the split view is collapsible
func (sv *spectrumSplitView) Collapsible(collapsible bool) *spectrumSplitView {
	sv.PropCollapsible = collapsible
	return sv
}

// Label sets the accessible label for the split view
func (sv *spectrumSplitView) Label(label string) *spectrumSplitView {
	sv.PropLabel = label
	return sv
}

// PrimaryMax sets the maximum size of the primary pane
func (sv *spectrumSplitView) PrimaryMax(max string) *spectrumSplitView {
	sv.PropPrimaryMax = max
	return sv
}

// PrimaryMin sets the minimum size of the primary pane
func (sv *spectrumSplitView) PrimaryMin(min string) *spectrumSplitView {
	sv.PropPrimaryMin = min
	return sv
}

// PrimarySize sets the size of the primary pane
func (sv *spectrumSplitView) PrimarySize(size string) *spectrumSplitView {
	sv.PropPrimarySize = size
	return sv
}

// Resizable sets whether the split view is resizable
func (sv *spectrumSplitView) Resizable(resizable bool) *spectrumSplitView {
	sv.PropResizable = resizable
	return sv
}

// SecondaryMax sets the maximum size of the secondary pane
func (sv *spectrumSplitView) SecondaryMax(max string) *spectrumSplitView {
	sv.PropSecondaryMax = max
	return sv
}

// SecondaryMin sets the minimum size of the secondary pane
func (sv *spectrumSplitView) SecondaryMin(min string) *spectrumSplitView {
	sv.PropSecondaryMin = min
	return sv
}

// SplitterPos sets the position of the splitter
func (sv *spectrumSplitView) SplitterPos(pos int) *spectrumSplitView {
	sv.PropSplitterPos = pos
	return sv
}

// Vertical sets whether the split view is vertical
func (sv *spectrumSplitView) Vertical(vertical bool) *spectrumSplitView {
	sv.PropVertical = vertical
	return sv
}

// PrimaryPanel sets the primary panel (first child)
func (sv *spectrumSplitView) PrimaryPanel(panel app.UI) *spectrumSplitView {
	sv.PropPrimaryPanel = panel
	return sv
}

// SecondaryPanel sets the secondary panel (second child)
func (sv *spectrumSplitView) SecondaryPanel(panel app.UI) *spectrumSplitView {
	sv.PropSecondaryPanel = panel
	return sv
}

// OnChange sets the change event handler
func (sv *spectrumSplitView) OnChange(handler app.EventHandler) *spectrumSplitView {
	sv.PropOnChange = handler
	return sv
}

// Render renders the split view component
func (sv *spectrumSplitView) Render() app.UI {
	splitView := app.Elem("sp-split-view")

	// Set attributes
	if sv.PropCollapsible {
		splitView = splitView.Attr("collapsible", true)
	}
	if sv.PropLabel != "" {
		splitView = splitView.Attr("label", sv.PropLabel)
	}
	if sv.PropPrimaryMax != "" {
		splitView = splitView.Attr("primary-max", sv.PropPrimaryMax)
	}
	if sv.PropPrimaryMin != "" {
		splitView = splitView.Attr("primary-min", sv.PropPrimaryMin)
	}
	if sv.PropPrimarySize != "" {
		splitView = splitView.Attr("primary-size", sv.PropPrimarySize)
	}
	if sv.PropResizable {
		splitView = splitView.Attr("resizable", true)
	}
	if sv.PropSecondaryMax != "" {
		splitView = splitView.Attr("secondary-max", sv.PropSecondaryMax)
	}
	if sv.PropSecondaryMin != "" {
		splitView = splitView.Attr("secondary-min", sv.PropSecondaryMin)
	}
	if sv.PropSplitterPos != 0 {
		splitView = splitView.Attr("splitter-pos", sv.PropSplitterPos)
	}
	if sv.PropVertical {
		splitView = splitView.Attr("vertical", true)
	}

	// Add event handlers
	if sv.PropOnChange != nil {
		splitView = splitView.On("change", sv.PropOnChange)
	}

	// Add panels
	if sv.PropPrimaryPanel != nil && sv.PropSecondaryPanel != nil {
		splitView = splitView.Body(
			sv.PropPrimaryPanel,
			sv.PropSecondaryPanel,
		)
	}

	return splitView
}
