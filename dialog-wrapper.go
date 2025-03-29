package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// spectrumDialogWrapper represents an sp-dialog-wrapper component
type spectrumDialogWrapper struct {
	app.Compo

	// Properties
	PropOpen               bool
	PropCancelLabel        string
	PropConfirmLabel       string
	PropDismissLabel       string
	PropDismissable        bool
	PropError              bool
	PropFooter             string
	PropHeadline           string
	PropHeadlineVisibility string
	PropHero               string
	PropHeroLabel          string
	PropMode               DialogMode
	PropNoDivider          bool
	PropResponsive         bool
	PropSecondaryLabel     string
	PropSize               DialogSize
	PropUnderlay           bool

	// Content
	PropContent []app.UI

	// Event handlers
	PropOnCancel    app.EventHandler
	PropOnClose     app.EventHandler
	PropOnConfirm   app.EventHandler
	PropOnSecondary app.EventHandler
}

// DialogWrapper creates a new dialog wrapper component
func DialogWrapper() *spectrumDialogWrapper {
	return &spectrumDialogWrapper{
		PropDismissLabel: "Close", // Default dismiss label
	}
}

// Open sets whether the dialog wrapper is open
func (dw *spectrumDialogWrapper) Open(open bool) *spectrumDialogWrapper {
	dw.PropOpen = open
	return dw
}

// CancelLabel sets the label for the cancel button
func (dw *spectrumDialogWrapper) CancelLabel(label string) *spectrumDialogWrapper {
	dw.PropCancelLabel = label
	return dw
}

// ConfirmLabel sets the label for the confirm button
func (dw *spectrumDialogWrapper) ConfirmLabel(label string) *spectrumDialogWrapper {
	dw.PropConfirmLabel = label
	return dw
}

// DismissLabel sets the label for the dismiss button
func (dw *spectrumDialogWrapper) DismissLabel(label string) *spectrumDialogWrapper {
	dw.PropDismissLabel = label
	return dw
}

// Dismissable sets whether the dialog has a dismiss button
func (dw *spectrumDialogWrapper) Dismissable(dismissable bool) *spectrumDialogWrapper {
	dw.PropDismissable = dismissable
	return dw
}

// Error sets whether the dialog displays an error state
func (dw *spectrumDialogWrapper) Error(error bool) *spectrumDialogWrapper {
	dw.PropError = error
	return dw
}

// Footer sets the footer content for the dialog
func (dw *spectrumDialogWrapper) Footer(footer string) *spectrumDialogWrapper {
	dw.PropFooter = footer
	return dw
}

// Headline sets the headline text for the dialog
func (dw *spectrumDialogWrapper) Headline(headline string) *spectrumDialogWrapper {
	dw.PropHeadline = headline
	return dw
}

// HeadlineVisibility sets the visibility of the headline
func (dw *spectrumDialogWrapper) HeadlineVisibility(visibility string) *spectrumDialogWrapper {
	dw.PropHeadlineVisibility = visibility
	return dw
}

// Hero sets the hero image URL for the dialog
func (dw *spectrumDialogWrapper) Hero(hero string) *spectrumDialogWrapper {
	dw.PropHero = hero
	return dw
}

// HeroLabel sets the accessibility label for the hero image
func (dw *spectrumDialogWrapper) HeroLabel(label string) *spectrumDialogWrapper {
	dw.PropHeroLabel = label
	return dw
}

// Mode sets the display mode of the dialog
func (dw *spectrumDialogWrapper) Mode(mode DialogMode) *spectrumDialogWrapper {
	dw.PropMode = mode
	return dw
}

// NoDivider sets whether the dialog should have a divider between header and content
func (dw *spectrumDialogWrapper) NoDivider(noDivider bool) *spectrumDialogWrapper {
	dw.PropNoDivider = noDivider
	return dw
}

// Responsive sets whether the dialog should be responsive
func (dw *spectrumDialogWrapper) Responsive(responsive bool) *spectrumDialogWrapper {
	dw.PropResponsive = responsive
	return dw
}

// SecondaryLabel sets the label for the secondary button
func (dw *spectrumDialogWrapper) SecondaryLabel(label string) *spectrumDialogWrapper {
	dw.PropSecondaryLabel = label
	return dw
}

// Size sets the visual size of the dialog
func (dw *spectrumDialogWrapper) Size(size DialogSize) *spectrumDialogWrapper {
	dw.PropSize = size
	return dw
}

// Underlay sets whether the dialog should have an underlay
func (dw *spectrumDialogWrapper) Underlay(underlay bool) *spectrumDialogWrapper {
	dw.PropUnderlay = underlay
	return dw
}

// Content adds content to the dialog's main content area
func (dw *spectrumDialogWrapper) Content(content app.UI) *spectrumDialogWrapper {
	dw.PropContent = append(dw.PropContent, content)
	return dw
}

// ContentList sets all content elements for the dialog
func (dw *spectrumDialogWrapper) ContentList(content ...app.UI) *spectrumDialogWrapper {
	dw.PropContent = content
	return dw
}

// ContentText adds text content to the dialog
func (dw *spectrumDialogWrapper) ContentText(text string) *spectrumDialogWrapper {
	dw.PropContent = append(dw.PropContent, app.P().Text(text))
	return dw
}

// OnCancel sets the cancel event handler
func (dw *spectrumDialogWrapper) OnCancel(handler app.EventHandler) *spectrumDialogWrapper {
	dw.PropOnCancel = handler
	return dw
}

// OnClose sets the close event handler
func (dw *spectrumDialogWrapper) OnClose(handler app.EventHandler) *spectrumDialogWrapper {
	dw.PropOnClose = handler
	return dw
}

// OnConfirm sets the confirm event handler
func (dw *spectrumDialogWrapper) OnConfirm(handler app.EventHandler) *spectrumDialogWrapper {
	dw.PropOnConfirm = handler
	return dw
}

// OnSecondary sets the secondary event handler
func (dw *spectrumDialogWrapper) OnSecondary(handler app.EventHandler) *spectrumDialogWrapper {
	dw.PropOnSecondary = handler
	return dw
}

// Convenience methods for setting dialog sizes

// Small sets size to small (s)
func (dw *spectrumDialogWrapper) Small() *spectrumDialogWrapper {
	return dw.Size(DialogSizeS)
}

// Medium sets size to medium (m)
func (dw *spectrumDialogWrapper) Medium() *spectrumDialogWrapper {
	return dw.Size(DialogSizeM)
}

// Large sets size to large (l)
func (dw *spectrumDialogWrapper) Large() *spectrumDialogWrapper {
	return dw.Size(DialogSizeL)
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (dw *spectrumDialogWrapper) Fullscreen() *spectrumDialogWrapper {
	return dw.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (dw *spectrumDialogWrapper) FullscreenTakeover() *spectrumDialogWrapper {
	return dw.Mode(DialogModeFullscreenTakeover)
}

// HideHeadline sets headline visibility to 'none'
func (dw *spectrumDialogWrapper) HideHeadline() *spectrumDialogWrapper {
	return dw.HeadlineVisibility("none")
}

// Render renders the dialog wrapper component
func (dw *spectrumDialogWrapper) Render() app.UI {
	dialogWrapper := app.Elem("sp-dialog-wrapper")

	// Add properties
	if dw.PropOpen {
		dialogWrapper = dialogWrapper.Attr("open", true)
	}
	if dw.PropCancelLabel != "" {
		dialogWrapper = dialogWrapper.Attr("cancel-label", dw.PropCancelLabel)
	}
	if dw.PropConfirmLabel != "" {
		dialogWrapper = dialogWrapper.Attr("confirm-label", dw.PropConfirmLabel)
	}
	if dw.PropDismissLabel != "" {
		dialogWrapper = dialogWrapper.Attr("dismiss-label", dw.PropDismissLabel)
	}
	if dw.PropDismissable {
		dialogWrapper = dialogWrapper.Attr("dismissable", true)
	}
	if dw.PropError {
		dialogWrapper = dialogWrapper.Attr("error", true)
	}
	if dw.PropFooter != "" {
		dialogWrapper = dialogWrapper.Attr("footer", dw.PropFooter)
	}
	if dw.PropHeadline != "" {
		dialogWrapper = dialogWrapper.Attr("headline", dw.PropHeadline)
	}
	if dw.PropHeadlineVisibility != "" {
		dialogWrapper = dialogWrapper.Attr("headline-visibility", dw.PropHeadlineVisibility)
	}
	if dw.PropHero != "" {
		dialogWrapper = dialogWrapper.Attr("hero", dw.PropHero)
	}
	if dw.PropHeroLabel != "" {
		dialogWrapper = dialogWrapper.Attr("hero-label", dw.PropHeroLabel)
	}
	if dw.PropMode != "" {
		dialogWrapper = dialogWrapper.Attr("mode", string(dw.PropMode))
	}
	if dw.PropNoDivider {
		dialogWrapper = dialogWrapper.Attr("no-divider", true)
	}
	if dw.PropResponsive {
		dialogWrapper = dialogWrapper.Attr("responsive", true)
	}
	if dw.PropSecondaryLabel != "" {
		dialogWrapper = dialogWrapper.Attr("secondary-label", dw.PropSecondaryLabel)
	}
	if dw.PropSize != "" {
		dialogWrapper = dialogWrapper.Attr("size", string(dw.PropSize))
	}
	if dw.PropUnderlay {
		dialogWrapper = dialogWrapper.Attr("underlay", true)
	}

	// Add event handlers
	if dw.PropOnCancel != nil {
		dialogWrapper = dialogWrapper.On("cancel", dw.PropOnCancel)
	}
	if dw.PropOnClose != nil {
		dialogWrapper = dialogWrapper.On("close", dw.PropOnClose)
	}
	if dw.PropOnConfirm != nil {
		dialogWrapper = dialogWrapper.On("confirm", dw.PropOnConfirm)
	}
	if dw.PropOnSecondary != nil {
		dialogWrapper = dialogWrapper.On("secondary", dw.PropOnSecondary)
	}

	// Add content if provided
	if len(dw.PropContent) > 0 {
		dialogWrapper = dialogWrapper.Body(dw.PropContent...)
	}

	return dialogWrapper
}
