package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// TabsSize represents the visual size of tabs
type TabsSize string

// Tabs sizes
const (
	TabsSizeS TabsSize = "s"
	TabsSizeM TabsSize = "m"
	TabsSizeL TabsSize = "l"
)

// TabsDirection represents the direction of the tabs
type TabsDirection string

// Tabs directions
const (
	TabsDirectionHorizontal    TabsDirection = "horizontal"
	TabsDirectionVertical      TabsDirection = "vertical"
	TabsDirectionVerticalRight TabsDirection = "vertical-right"
)

// SpectrumTabs represents a tabbed interface with tabs and panels
type SpectrumTabs struct {
	app.Compo

	// Properties
	selected         string
	size             TabsSize
	direction        TabsDirection
	auto             bool
	compact          bool
	disabled         bool
	emphasized       bool
	quiet            bool
	enableTabsScroll bool
	label            string

	// Child elements
	tabs      []app.UI
	tabPanels []app.UI
}

// Tabs creates a new tabs component
func Tabs() *SpectrumTabs {
	return &SpectrumTabs{
		direction: TabsDirectionHorizontal,
		size:      TabsSizeM,
	}
}

// Selected sets the value of the selected tab
func (t *SpectrumTabs) Selected(value string) *SpectrumTabs {
	t.selected = value
	return t
}

// Size sets the visual size of the tabs
func (t *SpectrumTabs) Size(size TabsSize) *SpectrumTabs {
	t.size = size
	return t
}

// Direction sets the direction of the tabs
func (t *SpectrumTabs) Direction(direction TabsDirection) *SpectrumTabs {
	t.direction = direction
	return t
}

// Auto sets whether tabs are automatically activated on keyboard focus
func (t *SpectrumTabs) Auto(auto bool) *SpectrumTabs {
	t.auto = auto
	return t
}

// Compact sets whether the tabs are displayed more compactly
func (t *SpectrumTabs) Compact(compact bool) *SpectrumTabs {
	t.compact = compact
	return t
}

// Disabled sets whether the tabs component is disabled
func (t *SpectrumTabs) Disabled(disabled bool) *SpectrumTabs {
	t.disabled = disabled
	return t
}

// Emphasized sets whether the tabs should have an emphasized appearance
func (t *SpectrumTabs) Emphasized(emphasized bool) *SpectrumTabs {
	t.emphasized = emphasized
	return t
}

// Quiet sets whether the tabs have a quiet appearance without a border
func (t *SpectrumTabs) Quiet(quiet bool) *SpectrumTabs {
	t.quiet = quiet
	return t
}

// EnableTabsScroll sets whether tabs can be scrolled
func (t *SpectrumTabs) EnableTabsScroll(enable bool) *SpectrumTabs {
	t.enableTabsScroll = enable
	return t
}

// Label sets the accessible label for the tabs component
func (t *SpectrumTabs) Label(label string) *SpectrumTabs {
	t.label = label
	return t
}

// Tab adds a tab to the tabs component
func (t *SpectrumTabs) Tab(tab *SpectrumTab) *SpectrumTabs {
	t.tabs = append(t.tabs, tab)
	return t
}

// TabPanel adds a tab panel to the tabs component
func (t *SpectrumTabs) TabPanel(panel *SpectrumTabPanel) *SpectrumTabs {
	t.tabPanels = append(t.tabPanels, panel)
	return t
}

// Render renders the tabs component
func (t *SpectrumTabs) Render() app.UI {
	tabs := app.Elem("sp-tabs").
		Attr("size", string(t.size)).
		Attr("direction", string(t.direction))

	if t.selected != "" {
		tabs = tabs.Attr("selected", t.selected)
	}
	if t.auto {
		tabs = tabs.Attr("auto", "")
	}
	if t.compact {
		tabs = tabs.Attr("compact", "")
	}
	if t.disabled {
		tabs = tabs.Attr("disabled", "")
	}
	if t.emphasized {
		tabs = tabs.Attr("emphasized", "")
	}
	if t.enableTabsScroll {
		tabs = tabs.Attr("enableTabsScroll", "")
	}
	if t.quiet {
		tabs = tabs.Attr("quiet", "")
	}
	if t.label != "" {
		tabs = tabs.Attr("label", t.label)
	}

	// Add tab elements first
	elements := []app.UI{}
	elements = append(elements, t.tabs...)

	// Then add tab panels
	elements = append(elements, t.tabPanels...)

	// Add all elements to the tabs
	if len(elements) > 0 {
		tabs = tabs.Body(elements...)
	}

	return tabs
}
