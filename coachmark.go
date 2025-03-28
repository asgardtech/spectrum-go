package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AssetType for Coachmark
type CoachmarkAssetType string

const (
	CoachmarkAssetFile   CoachmarkAssetType = "file"
	CoachmarkAssetFolder CoachmarkAssetType = "folder"
)

// SpectrumCoachmark represents an sp-coachmark component
type SpectrumCoachmark struct {
	app.Compo

	// Properties
	asset        CoachmarkAssetType
	currentStep  int
	hasAsset     bool
	item         string // Using string as a simplification for CoachmarkItem
	modifierKeys []string
	open         bool
	placement    string
	tip          bool
	totalSteps   int
	primaryCta   string
	secondaryCta string
	src          string
	mediaType    string
	imageAlt     string

	// Slots
	title       app.UI
	content     app.UI
	actions     app.UI
	assetSlot   app.UI
	coverPhoto  app.UI
	stepCount   app.UI
	description app.UI
	heading     app.UI
	children    []app.UI

	// Event handlers
	onPrimary   app.EventHandler
	onSecondary app.EventHandler
}

// Coachmark creates a new coachmark component
func Coachmark() *SpectrumCoachmark {
	return &SpectrumCoachmark{
		placement: "right", // Default placement
	}
}

// Asset sets the asset type
func (cm *SpectrumCoachmark) Asset(asset CoachmarkAssetType) *SpectrumCoachmark {
	cm.asset = asset
	return cm
}

// CurrentStep sets the current step in a tour
func (cm *SpectrumCoachmark) CurrentStep(step int) *SpectrumCoachmark {
	cm.currentStep = step
	return cm
}

// HasAsset sets whether the coachmark has an asset
func (cm *SpectrumCoachmark) HasAsset(hasAsset bool) *SpectrumCoachmark {
	cm.hasAsset = hasAsset
	return cm
}

// Item sets the coachmark item
func (cm *SpectrumCoachmark) Item(item string) *SpectrumCoachmark {
	cm.item = item
	return cm
}

// ModifierKeys sets the modifier keys
func (cm *SpectrumCoachmark) ModifierKeys(keys []string) *SpectrumCoachmark {
	cm.modifierKeys = keys
	return cm
}

// Open sets whether the coachmark is visible
func (cm *SpectrumCoachmark) Open(open bool) *SpectrumCoachmark {
	cm.open = open
	return cm
}

// Placement sets the placement of the coachmark
func (cm *SpectrumCoachmark) Placement(placement string) *SpectrumCoachmark {
	cm.placement = placement
	return cm
}

// Tip sets whether the coachmark has a tip
func (cm *SpectrumCoachmark) Tip(tip bool) *SpectrumCoachmark {
	cm.tip = tip
	return cm
}

// TotalSteps sets the total number of steps in a tour
func (cm *SpectrumCoachmark) TotalSteps(steps int) *SpectrumCoachmark {
	cm.totalSteps = steps
	return cm
}

// PrimaryCta sets the text for the primary call to action button
func (cm *SpectrumCoachmark) PrimaryCta(text string) *SpectrumCoachmark {
	cm.primaryCta = text
	return cm
}

// SecondaryCta sets the text for the secondary call to action button
func (cm *SpectrumCoachmark) SecondaryCta(text string) *SpectrumCoachmark {
	cm.secondaryCta = text
	return cm
}

// Src sets the source URL for the media
func (cm *SpectrumCoachmark) Src(src string) *SpectrumCoachmark {
	cm.src = src
	return cm
}

// MediaType sets the type of media
func (cm *SpectrumCoachmark) MediaType(mediaType string) *SpectrumCoachmark {
	cm.mediaType = mediaType
	return cm
}

// ImageAlt sets the alt text for the image
func (cm *SpectrumCoachmark) ImageAlt(alt string) *SpectrumCoachmark {
	cm.imageAlt = alt
	return cm
}

// Title sets the title content in the title slot
func (cm *SpectrumCoachmark) Title(title app.UI) *SpectrumCoachmark {
	cm.title = title
	return cm
}

// TitleText sets a text title in the title slot
func (cm *SpectrumCoachmark) TitleText(text string) *SpectrumCoachmark {
	cm.title = app.Div().Text(text)
	return cm
}

// Content sets the content in the content slot
func (cm *SpectrumCoachmark) Content(content app.UI) *SpectrumCoachmark {
	cm.content = content
	return cm
}

// ContentText sets text content in the content slot
func (cm *SpectrumCoachmark) ContentText(text string) *SpectrumCoachmark {
	cm.content = app.Div().Text(text)
	return cm
}

// Actions sets the actions menu in the actions slot
func (cm *SpectrumCoachmark) Actions(actions app.UI) *SpectrumCoachmark {
	cm.actions = actions
	return cm
}

// AssetSlot sets the asset in the asset slot
func (cm *SpectrumCoachmark) AssetSlot(asset app.UI) *SpectrumCoachmark {
	cm.assetSlot = asset
	return cm
}

// CoverPhoto sets the cover photo in the cover-photo slot
func (cm *SpectrumCoachmark) CoverPhoto(photo app.UI) *SpectrumCoachmark {
	cm.coverPhoto = photo
	return cm
}

// StepCount sets the step count in the step-count slot
func (cm *SpectrumCoachmark) StepCount(stepCount app.UI) *SpectrumCoachmark {
	cm.stepCount = stepCount
	return cm
}

// Description sets the description in the description slot
func (cm *SpectrumCoachmark) Description(description app.UI) *SpectrumCoachmark {
	cm.description = description
	return cm
}

// Heading sets the heading in the heading slot
func (cm *SpectrumCoachmark) Heading(heading app.UI) *SpectrumCoachmark {
	cm.heading = heading
	return cm
}

// Child adds a child element to the coachmark
func (cm *SpectrumCoachmark) Child(child app.UI) *SpectrumCoachmark {
	cm.children = append(cm.children, child)
	return cm
}

// Children sets the child elements of the coachmark
func (cm *SpectrumCoachmark) Children(children ...app.UI) *SpectrumCoachmark {
	cm.children = children
	return cm
}

// OnPrimary sets the primary button click event handler
func (cm *SpectrumCoachmark) OnPrimary(handler app.EventHandler) *SpectrumCoachmark {
	cm.onPrimary = handler
	return cm
}

// OnSecondary sets the secondary button click event handler
func (cm *SpectrumCoachmark) OnSecondary(handler app.EventHandler) *SpectrumCoachmark {
	cm.onSecondary = handler
	return cm
}

// Render renders the coachmark component
func (cm *SpectrumCoachmark) Render() app.UI {
	coachmark := app.Elem("sp-coachmark")

	// Set attributes
	if cm.asset != "" {
		coachmark = coachmark.Attr("asset", string(cm.asset))
	}
	if cm.currentStep > 0 {
		coachmark = coachmark.Attr("current-step", cm.currentStep)
	}
	if cm.hasAsset {
		coachmark = coachmark.Attr("has-asset", true)
	}
	if cm.item != "" {
		coachmark = coachmark.Attr("item", cm.item)
	}
	if len(cm.modifierKeys) > 0 {
		// We're simplifying here; in a real implementation,
		// you'd need to handle the JSON serialization properly
		coachmark = coachmark.Attr("modifier-keys", cm.modifierKeys)
	}
	if cm.open {
		coachmark = coachmark.Attr("open", true)
	}
	if cm.placement != "right" {
		coachmark = coachmark.Attr("placement", cm.placement)
	}
	if cm.tip {
		coachmark = coachmark.Attr("tip", true)
	}
	if cm.totalSteps > 0 {
		coachmark = coachmark.Attr("total-steps", cm.totalSteps)
	}
	if cm.primaryCta != "" {
		coachmark = coachmark.Attr("primary-cta", cm.primaryCta)
	}
	if cm.secondaryCta != "" {
		coachmark = coachmark.Attr("secondary-cta", cm.secondaryCta)
	}
	if cm.src != "" {
		coachmark = coachmark.Attr("src", cm.src)
	}
	if cm.mediaType != "" {
		coachmark = coachmark.Attr("media-type", cm.mediaType)
	}
	if cm.imageAlt != "" {
		coachmark = coachmark.Attr("image-alt", cm.imageAlt)
	}

	// Add event handlers
	if cm.onPrimary != nil {
		coachmark = coachmark.On("primary", cm.onPrimary)
	}
	if cm.onSecondary != nil {
		coachmark = coachmark.On("secondary", cm.onSecondary)
	}

	// Prepare slot elements
	elements := []app.UI{}

	// Add slotted elements
	if cm.title != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "title").Body(cm.title))
	}
	if cm.content != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "content").Body(cm.content))
	}
	if cm.actions != nil {
		if actionWithSlot, ok := cm.actions.(interface{ Slot(string) app.UI }); ok {
			elements = append(elements, actionWithSlot.Slot("actions"))
		} else {
			elements = append(elements, app.Elem("div").Attr("slot", "actions").Body(cm.actions))
		}
	}
	if cm.assetSlot != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "asset").Body(cm.assetSlot))
	}
	if cm.coverPhoto != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "cover-photo").Body(cm.coverPhoto))
	}
	if cm.stepCount != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "step-count").Body(cm.stepCount))
	}
	if cm.description != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "description").Body(cm.description))
	}
	if cm.heading != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "heading").Body(cm.heading))
	}

	// Add default slot children
	if len(cm.children) > 0 {
		elements = append(elements, cm.children...)
	}

	if len(elements) > 0 {
		coachmark = coachmark.Body(elements...)
	}

	return coachmark
}
