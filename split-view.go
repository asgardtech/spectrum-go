package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumSplitView represents an sp-split-view component
type SpectrumSplitView struct {
	app.Compo

	// Properties
	collapsible  bool
	label        string
	primaryMax   string
	primaryMin   string
	primarySize  string
	resizable    bool
	secondaryMax string
	secondaryMin string
	splitterPos  int
	vertical     bool

	// Child components
	primaryPanel   app.UI
	secondaryPanel app.UI

	// Event handlers
	onChange app.EventHandler
}

// SplitView creates a new split view component
func SplitView() *SpectrumSplitView {
	return &SpectrumSplitView{
		label: "Resize the panels", // Default label
	}
}

// Collapsible sets whether the split view is collapsible
func (sv *SpectrumSplitView) Collapsible(collapsible bool) *SpectrumSplitView {
	sv.collapsible = collapsible
	return sv
}

// Label sets the accessible label for the split view
func (sv *SpectrumSplitView) Label(label string) *SpectrumSplitView {
	sv.label = label
	return sv
}

// PrimaryMax sets the maximum size of the primary pane
func (sv *SpectrumSplitView) PrimaryMax(max string) *SpectrumSplitView {
	sv.primaryMax = max
	return sv
}

// PrimaryMin sets the minimum size of the primary pane
func (sv *SpectrumSplitView) PrimaryMin(min string) *SpectrumSplitView {
	sv.primaryMin = min
	return sv
}

// PrimarySize sets the size of the primary pane
func (sv *SpectrumSplitView) PrimarySize(size string) *SpectrumSplitView {
	sv.primarySize = size
	return sv
}

// Resizable sets whether the split view is resizable
func (sv *SpectrumSplitView) Resizable(resizable bool) *SpectrumSplitView {
	sv.resizable = resizable
	return sv
}

// SecondaryMax sets the maximum size of the secondary pane
func (sv *SpectrumSplitView) SecondaryMax(max string) *SpectrumSplitView {
	sv.secondaryMax = max
	return sv
}

// SecondaryMin sets the minimum size of the secondary pane
func (sv *SpectrumSplitView) SecondaryMin(min string) *SpectrumSplitView {
	sv.secondaryMin = min
	return sv
}

// SplitterPos sets the position of the splitter
func (sv *SpectrumSplitView) SplitterPos(pos int) *SpectrumSplitView {
	sv.splitterPos = pos
	return sv
}

// Vertical sets whether the split view is vertical
func (sv *SpectrumSplitView) Vertical(vertical bool) *SpectrumSplitView {
	sv.vertical = vertical
	return sv
}

// PrimaryPanel sets the primary panel (first child)
func (sv *SpectrumSplitView) PrimaryPanel(panel app.UI) *SpectrumSplitView {
	sv.primaryPanel = panel
	return sv
}

// SecondaryPanel sets the secondary panel (second child)
func (sv *SpectrumSplitView) SecondaryPanel(panel app.UI) *SpectrumSplitView {
	sv.secondaryPanel = panel
	return sv
}

// OnChange sets the change event handler
func (sv *SpectrumSplitView) OnChange(handler app.EventHandler) *SpectrumSplitView {
	sv.onChange = handler
	return sv
}

// Render renders the split view component
func (sv *SpectrumSplitView) Render() app.UI {
	splitView := app.Elem("sp-split-view")

	// Set attributes
	if sv.collapsible {
		splitView = splitView.Attr("collapsible", true)
	}
	if sv.label != "" {
		splitView = splitView.Attr("label", sv.label)
	}
	if sv.primaryMax != "" {
		splitView = splitView.Attr("primary-max", sv.primaryMax)
	}
	if sv.primaryMin != "" {
		splitView = splitView.Attr("primary-min", sv.primaryMin)
	}
	if sv.primarySize != "" {
		splitView = splitView.Attr("primary-size", sv.primarySize)
	}
	if sv.resizable {
		splitView = splitView.Attr("resizable", true)
	}
	if sv.secondaryMax != "" {
		splitView = splitView.Attr("secondary-max", sv.secondaryMax)
	}
	if sv.secondaryMin != "" {
		splitView = splitView.Attr("secondary-min", sv.secondaryMin)
	}
	if sv.splitterPos != 0 {
		splitView = splitView.Attr("splitter-pos", sv.splitterPos)
	}
	if sv.vertical {
		splitView = splitView.Attr("vertical", true)
	}

	// Add event handlers
	if sv.onChange != nil {
		splitView = splitView.On("change", sv.onChange)
	}

	// Add panels
	if sv.primaryPanel != nil && sv.secondaryPanel != nil {
		splitView = splitView.Body(
			sv.primaryPanel,
			sv.secondaryPanel,
		)
	}

	return splitView
}
