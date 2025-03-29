package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// AssetType for Coachmark
type CoachmarkAssetType string

const (
	CoachmarkAssetFile   CoachmarkAssetType = "file"
	CoachmarkAssetFolder CoachmarkAssetType = "folder"
)

// spectrumCoachmark represents an sp-coachmark component
type spectrumCoachmark struct {
	app.Compo

	// Properties
	PropAsset        CoachmarkAssetType
	PropCurrentStep  int
	PropHasAsset     bool
	PropItem         string // Using string as a simplification for CoachmarkItem
	PropModifierKeys []string
	PropOpen         bool
	PropPlacement    string
	PropTip          bool
	PropTotalSteps   int
	PropPrimaryCta   string
	PropSecondaryCta string
	PropSrc          string
	PropMediaType    string
	PropImageAlt     string

	// Slots
	PropTitle       app.UI
	PropContent     app.UI
	PropActions     app.UI
	PropAssetSlot   app.UI
	PropCoverPhoto  app.UI
	PropStepCount   app.UI
	PropDescription app.UI
	PropHeading     app.UI
	PropChildren    []app.UI

	// Event handlers
	PropOnPrimary   app.EventHandler
	PropOnSecondary app.EventHandler
}

// Coachmark creates a new coachmark component
func Coachmark() *spectrumCoachmark {
	return &spectrumCoachmark{
		PropPlacement: "right", // Default placement
	}
}

// Asset sets the asset type
func (cm *spectrumCoachmark) Asset(asset CoachmarkAssetType) *spectrumCoachmark {
	cm.PropAsset = asset
	return cm
}

// CurrentStep sets the current step in a tour
func (cm *spectrumCoachmark) CurrentStep(step int) *spectrumCoachmark {
	cm.PropCurrentStep = step
	return cm
}

// HasAsset sets whether the coachmark has an asset
func (cm *spectrumCoachmark) HasAsset(hasAsset bool) *spectrumCoachmark {
	cm.PropHasAsset = hasAsset
	return cm
}

// Item sets the coachmark item
func (cm *spectrumCoachmark) Item(item string) *spectrumCoachmark {
	cm.PropItem = item
	return cm
}

// ModifierKeys sets the modifier keys
func (cm *spectrumCoachmark) ModifierKeys(keys []string) *spectrumCoachmark {
	cm.PropModifierKeys = keys
	return cm
}

// Open sets whether the coachmark is visible
func (cm *spectrumCoachmark) Open(open bool) *spectrumCoachmark {
	cm.PropOpen = open
	return cm
}

// Placement sets the placement of the coachmark
func (cm *spectrumCoachmark) Placement(placement string) *spectrumCoachmark {
	cm.PropPlacement = placement
	return cm
}

// Tip sets whether the coachmark has a tip
func (cm *spectrumCoachmark) Tip(tip bool) *spectrumCoachmark {
	cm.PropTip = tip
	return cm
}

// TotalSteps sets the total number of steps in a tour
func (cm *spectrumCoachmark) TotalSteps(steps int) *spectrumCoachmark {
	cm.PropTotalSteps = steps
	return cm
}

// PrimaryCta sets the text for the primary call to action button
func (cm *spectrumCoachmark) PrimaryCta(text string) *spectrumCoachmark {
	cm.PropPrimaryCta = text
	return cm
}

// SecondaryCta sets the text for the secondary call to action button
func (cm *spectrumCoachmark) SecondaryCta(text string) *spectrumCoachmark {
	cm.PropSecondaryCta = text
	return cm
}

// Src sets the source URL for the media
func (cm *spectrumCoachmark) Src(src string) *spectrumCoachmark {
	cm.PropSrc = src
	return cm
}

// MediaType sets the type of media
func (cm *spectrumCoachmark) MediaType(mediaType string) *spectrumCoachmark {
	cm.PropMediaType = mediaType
	return cm
}

// ImageAlt sets the alt text for the image
func (cm *spectrumCoachmark) ImageAlt(alt string) *spectrumCoachmark {
	cm.PropImageAlt = alt
	return cm
}

// Title sets the title content in the title slot
func (cm *spectrumCoachmark) Title(title app.UI) *spectrumCoachmark {
	cm.PropTitle = title
	return cm
}

// TitleText sets a text title in the title slot
func (cm *spectrumCoachmark) TitleText(text string) *spectrumCoachmark {
	cm.PropTitle = app.Div().Text(text)
	return cm
}

// Content sets the content in the content slot
func (cm *spectrumCoachmark) Content(content app.UI) *spectrumCoachmark {
	cm.PropContent = content
	return cm
}

// ContentText sets text content in the content slot
func (cm *spectrumCoachmark) ContentText(text string) *spectrumCoachmark {
	cm.PropContent = app.Div().Text(text)
	return cm
}

// Actions sets the actions menu in the actions slot
func (cm *spectrumCoachmark) Actions(actions app.UI) *spectrumCoachmark {
	cm.PropActions = actions
	return cm
}

// AssetSlot sets the asset in the asset slot
func (cm *spectrumCoachmark) AssetSlot(asset app.UI) *spectrumCoachmark {
	cm.PropAssetSlot = asset
	return cm
}

// CoverPhoto sets the cover photo in the cover-photo slot
func (cm *spectrumCoachmark) CoverPhoto(photo app.UI) *spectrumCoachmark {
	cm.PropCoverPhoto = photo
	return cm
}

// StepCount sets the step count in the step-count slot
func (cm *spectrumCoachmark) StepCount(stepCount app.UI) *spectrumCoachmark {
	cm.PropStepCount = stepCount
	return cm
}

// Description sets the description in the description slot
func (cm *spectrumCoachmark) Description(description app.UI) *spectrumCoachmark {
	cm.PropDescription = description
	return cm
}

// Heading sets the heading in the heading slot
func (cm *spectrumCoachmark) Heading(heading app.UI) *spectrumCoachmark {
	cm.PropHeading = heading
	return cm
}

// Child adds a child element to the coachmark
func (cm *spectrumCoachmark) Child(child app.UI) *spectrumCoachmark {
	cm.PropChildren = append(cm.PropChildren, child)
	return cm
}

// Children sets the child elements of the coachmark
func (cm *spectrumCoachmark) Children(children ...app.UI) *spectrumCoachmark {
	cm.PropChildren = children
	return cm
}

// OnPrimary sets the primary button click event handler
func (cm *spectrumCoachmark) OnPrimary(handler app.EventHandler) *spectrumCoachmark {
	cm.PropOnPrimary = handler
	return cm
}

// OnSecondary sets the secondary button click event handler
func (cm *spectrumCoachmark) OnSecondary(handler app.EventHandler) *spectrumCoachmark {
	cm.PropOnSecondary = handler
	return cm
}

// Render renders the coachmark component
func (cm *spectrumCoachmark) Render() app.UI {
	coachmark := app.Elem("sp-coachmark")

	// Set attributes
	if cm.PropAsset != "" {
		coachmark = coachmark.Attr("asset", string(cm.PropAsset))
	}
	if cm.PropCurrentStep > 0 {
		coachmark = coachmark.Attr("current-step", cm.PropCurrentStep)
	}
	if cm.PropHasAsset {
		coachmark = coachmark.Attr("has-asset", true)
	}
	if cm.PropItem != "" {
		coachmark = coachmark.Attr("item", cm.PropItem)
	}
	if len(cm.PropModifierKeys) > 0 {
		// We're simplifying here; in a real implementation,
		// you'd need to handle the JSON serialization properly
		coachmark = coachmark.Attr("modifier-keys", cm.PropModifierKeys)
	}
	if cm.PropOpen {
		coachmark = coachmark.Attr("open", true)
	}
	if cm.PropPlacement != "right" {
		coachmark = coachmark.Attr("placement", cm.PropPlacement)
	}
	if cm.PropTip {
		coachmark = coachmark.Attr("tip", true)
	}
	if cm.PropTotalSteps > 0 {
		coachmark = coachmark.Attr("total-steps", cm.PropTotalSteps)
	}
	if cm.PropPrimaryCta != "" {
		coachmark = coachmark.Attr("primary-cta", cm.PropPrimaryCta)
	}
	if cm.PropSecondaryCta != "" {
		coachmark = coachmark.Attr("secondary-cta", cm.PropSecondaryCta)
	}
	if cm.PropSrc != "" {
		coachmark = coachmark.Attr("src", cm.PropSrc)
	}
	if cm.PropMediaType != "" {
		coachmark = coachmark.Attr("media-type", cm.PropMediaType)
	}
	if cm.PropImageAlt != "" {
		coachmark = coachmark.Attr("image-alt", cm.PropImageAlt)
	}

	// Set event handlers
	if cm.PropOnPrimary != nil {
		coachmark = coachmark.On("primary", cm.PropOnPrimary)
	}
	if cm.PropOnSecondary != nil {
		coachmark = coachmark.On("secondary", cm.PropOnSecondary)
	}

	// Add slotted content
	elements := []app.UI{}

	if cm.PropTitle != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "title").Body(cm.PropTitle))
	}

	if cm.PropContent != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "content").Body(cm.PropContent))
	}

	if cm.PropActions != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "actions").Body(cm.PropActions))
	}

	if cm.PropAssetSlot != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "asset").Body(cm.PropAssetSlot))
	}

	if cm.PropCoverPhoto != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "cover-photo").Body(cm.PropCoverPhoto))
	}

	if cm.PropStepCount != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "step-count").Body(cm.PropStepCount))
	}

	if cm.PropDescription != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "description").Body(cm.PropDescription))
	}

	if cm.PropHeading != nil {
		elements = append(elements, app.Elem("div").Attr("slot", "heading").Body(cm.PropHeading))
	}

	// Add children
	elements = append(elements, cm.PropChildren...)

	return coachmark.Body(elements...)
}
