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

// spectrumTabs represents a tabbed interface with tabs and panels
type spectrumTabs struct {
	app.Compo

	// Properties
	PropSelected         string
	PropSize             TabsSize
	PropDirection        TabsDirection
	PropAuto             bool
	PropCompact          bool
	PropDisabled         bool
	PropEmphasized       bool
	PropQuiet            bool
	PropEnableTabsScroll bool
	PropLabel            string

	// Child elements
	PropTabs      []app.UI
	PropTabPanels []app.UI
}

// Tabs creates a new tabs component
func Tabs() *spectrumTabs {
	return &spectrumTabs{
		PropDirection: TabsDirectionHorizontal,
		PropSize:      TabsSizeM,
	}
}

// Selected sets the value of the selected tab
func (t *spectrumTabs) Selected(value string) *spectrumTabs {
	t.PropSelected = value
	return t
}

// Size sets the visual size of the tabs
func (t *spectrumTabs) Size(size TabsSize) *spectrumTabs {
	t.PropSize = size
	return t
}

// Direction sets the direction of the tabs
func (t *spectrumTabs) Direction(direction TabsDirection) *spectrumTabs {
	t.PropDirection = direction
	return t
}

// Auto sets whether tabs are automatically activated on keyboard focus
func (t *spectrumTabs) Auto(auto bool) *spectrumTabs {
	t.PropAuto = auto
	return t
}

// Compact sets whether the tabs are displayed more compactly
func (t *spectrumTabs) Compact(compact bool) *spectrumTabs {
	t.PropCompact = compact
	return t
}

// Disabled sets whether the tabs component is disabled
func (t *spectrumTabs) Disabled(disabled bool) *spectrumTabs {
	t.PropDisabled = disabled
	return t
}

// Emphasized sets whether the tabs should have an emphasized appearance
func (t *spectrumTabs) Emphasized(emphasized bool) *spectrumTabs {
	t.PropEmphasized = emphasized
	return t
}

// Quiet sets whether the tabs have a quiet appearance without a border
func (t *spectrumTabs) Quiet(quiet bool) *spectrumTabs {
	t.PropQuiet = quiet
	return t
}

// EnableTabsScroll sets whether tabs can be scrolled
func (t *spectrumTabs) EnableTabsScroll(enable bool) *spectrumTabs {
	t.PropEnableTabsScroll = enable
	return t
}

// Label sets the accessible label for the tabs component
func (t *spectrumTabs) Label(label string) *spectrumTabs {
	t.PropLabel = label
	return t
}

// Tab adds a tab to the tabs component
func (t *spectrumTabs) Tab(tab *spectrumTab) *spectrumTabs {
	t.PropTabs = append(t.PropTabs, tab)
	return t
}

// TabPanel adds a tab panel to the tabs component
func (t *spectrumTabs) TabPanel(panel *spectrumTabPanel) *spectrumTabs {
	t.PropTabPanels = append(t.PropTabPanels, panel)
	return t
}

// Render renders the tabs component
func (t *spectrumTabs) Render() app.UI {
	tabs := app.Elem("sp-tabs").
		Attr("size", string(t.PropSize)).
		Attr("direction", string(t.PropDirection))

	if t.PropSelected != "" {
		tabs = tabs.Attr("selected", t.PropSelected)
	}
	if t.PropAuto {
		tabs = tabs.Attr("auto", "")
	}
	if t.PropCompact {
		tabs = tabs.Attr("compact", "")
	}
	if t.PropDisabled {
		tabs = tabs.Attr("disabled", "")
	}
	if t.PropEmphasized {
		tabs = tabs.Attr("emphasized", "")
	}
	if t.PropEnableTabsScroll {
		tabs = tabs.Attr("enableTabsScroll", "")
	}
	if t.PropQuiet {
		tabs = tabs.Attr("quiet", "")
	}
	if t.PropLabel != "" {
		tabs = tabs.Attr("label", t.PropLabel)
	}

	// Add tab elements first
	elements := []app.UI{}
	elements = append(elements, t.PropTabs...)

	// Then add tab panels
	elements = append(elements, t.PropTabPanels...)

	// Add all elements to the tabs
	if len(elements) > 0 {
		tabs = tabs.Body(elements...)
	}

	return tabs
}
