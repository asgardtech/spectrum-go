package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumDialogWrapper represents an sp-dialog-wrapper component
type SpectrumDialogWrapper struct {
	app.Compo

	// Properties
	open               bool
	cancelLabel        string
	confirmLabel       string
	dismissLabel       string
	dismissable        bool
	error              bool
	footer             string
	headline           string
	headlineVisibility string
	hero               string
	heroLabel          string
	mode               DialogMode
	noDivider          bool
	responsive         bool
	secondaryLabel     string
	size               DialogSize
	underlay           bool

	// Content
	content []app.UI

	// Event handlers
	onCancel    app.EventHandler
	onClose     app.EventHandler
	onConfirm   app.EventHandler
	onSecondary app.EventHandler
}

// DialogWrapper creates a new dialog wrapper component
func DialogWrapper() *SpectrumDialogWrapper {
	return &SpectrumDialogWrapper{
		dismissLabel: "Close", // Default dismiss label
	}
}

// Open sets whether the dialog wrapper is open
func (dw *SpectrumDialogWrapper) Open(open bool) *SpectrumDialogWrapper {
	dw.open = open
	return dw
}

// CancelLabel sets the label for the cancel button
func (dw *SpectrumDialogWrapper) CancelLabel(label string) *SpectrumDialogWrapper {
	dw.cancelLabel = label
	return dw
}

// ConfirmLabel sets the label for the confirm button
func (dw *SpectrumDialogWrapper) ConfirmLabel(label string) *SpectrumDialogWrapper {
	dw.confirmLabel = label
	return dw
}

// DismissLabel sets the label for the dismiss button
func (dw *SpectrumDialogWrapper) DismissLabel(label string) *SpectrumDialogWrapper {
	dw.dismissLabel = label
	return dw
}

// Dismissable sets whether the dialog has a dismiss button
func (dw *SpectrumDialogWrapper) Dismissable(dismissable bool) *SpectrumDialogWrapper {
	dw.dismissable = dismissable
	return dw
}

// Error sets whether the dialog displays an error state
func (dw *SpectrumDialogWrapper) Error(error bool) *SpectrumDialogWrapper {
	dw.error = error
	return dw
}

// Footer sets the footer content for the dialog
func (dw *SpectrumDialogWrapper) Footer(footer string) *SpectrumDialogWrapper {
	dw.footer = footer
	return dw
}

// Headline sets the headline text for the dialog
func (dw *SpectrumDialogWrapper) Headline(headline string) *SpectrumDialogWrapper {
	dw.headline = headline
	return dw
}

// HeadlineVisibility sets the visibility of the headline
func (dw *SpectrumDialogWrapper) HeadlineVisibility(visibility string) *SpectrumDialogWrapper {
	dw.headlineVisibility = visibility
	return dw
}

// Hero sets the hero image URL for the dialog
func (dw *SpectrumDialogWrapper) Hero(hero string) *SpectrumDialogWrapper {
	dw.hero = hero
	return dw
}

// HeroLabel sets the accessibility label for the hero image
func (dw *SpectrumDialogWrapper) HeroLabel(label string) *SpectrumDialogWrapper {
	dw.heroLabel = label
	return dw
}

// Mode sets the display mode of the dialog
func (dw *SpectrumDialogWrapper) Mode(mode DialogMode) *SpectrumDialogWrapper {
	dw.mode = mode
	return dw
}

// NoDivider sets whether the dialog should have a divider between header and content
func (dw *SpectrumDialogWrapper) NoDivider(noDivider bool) *SpectrumDialogWrapper {
	dw.noDivider = noDivider
	return dw
}

// Responsive sets whether the dialog should be responsive
func (dw *SpectrumDialogWrapper) Responsive(responsive bool) *SpectrumDialogWrapper {
	dw.responsive = responsive
	return dw
}

// SecondaryLabel sets the label for the secondary button
func (dw *SpectrumDialogWrapper) SecondaryLabel(label string) *SpectrumDialogWrapper {
	dw.secondaryLabel = label
	return dw
}

// Size sets the visual size of the dialog
func (dw *SpectrumDialogWrapper) Size(size DialogSize) *SpectrumDialogWrapper {
	dw.size = size
	return dw
}

// Underlay sets whether the dialog should have an underlay
func (dw *SpectrumDialogWrapper) Underlay(underlay bool) *SpectrumDialogWrapper {
	dw.underlay = underlay
	return dw
}

// Content adds content to the dialog's main content area
func (dw *SpectrumDialogWrapper) Content(content app.UI) *SpectrumDialogWrapper {
	dw.content = append(dw.content, content)
	return dw
}

// ContentList sets all content elements for the dialog
func (dw *SpectrumDialogWrapper) ContentList(content ...app.UI) *SpectrumDialogWrapper {
	dw.content = content
	return dw
}

// ContentText adds text content to the dialog
func (dw *SpectrumDialogWrapper) ContentText(text string) *SpectrumDialogWrapper {
	dw.content = append(dw.content, app.P().Text(text))
	return dw
}

// OnCancel sets the cancel event handler
func (dw *SpectrumDialogWrapper) OnCancel(handler app.EventHandler) *SpectrumDialogWrapper {
	dw.onCancel = handler
	return dw
}

// OnClose sets the close event handler
func (dw *SpectrumDialogWrapper) OnClose(handler app.EventHandler) *SpectrumDialogWrapper {
	dw.onClose = handler
	return dw
}

// OnConfirm sets the confirm event handler
func (dw *SpectrumDialogWrapper) OnConfirm(handler app.EventHandler) *SpectrumDialogWrapper {
	dw.onConfirm = handler
	return dw
}

// OnSecondary sets the secondary event handler
func (dw *SpectrumDialogWrapper) OnSecondary(handler app.EventHandler) *SpectrumDialogWrapper {
	dw.onSecondary = handler
	return dw
}

// Convenience methods for setting dialog sizes

// Small sets size to small (s)
func (dw *SpectrumDialogWrapper) Small() *SpectrumDialogWrapper {
	return dw.Size(DialogSizeS)
}

// Medium sets size to medium (m)
func (dw *SpectrumDialogWrapper) Medium() *SpectrumDialogWrapper {
	return dw.Size(DialogSizeM)
}

// Large sets size to large (l)
func (dw *SpectrumDialogWrapper) Large() *SpectrumDialogWrapper {
	return dw.Size(DialogSizeL)
}

// Convenience methods for setting dialog modes

// Fullscreen sets mode to fullscreen
func (dw *SpectrumDialogWrapper) Fullscreen() *SpectrumDialogWrapper {
	return dw.Mode(DialogModeFullscreen)
}

// FullscreenTakeover sets mode to fullscreenTakeover
func (dw *SpectrumDialogWrapper) FullscreenTakeover() *SpectrumDialogWrapper {
	return dw.Mode(DialogModeFullscreenTakeover)
}

// HideHeadline sets headline visibility to 'none'
func (dw *SpectrumDialogWrapper) HideHeadline() *SpectrumDialogWrapper {
	return dw.HeadlineVisibility("none")
}

// Render renders the dialog wrapper component
func (dw *SpectrumDialogWrapper) Render() app.UI {
	dialogWrapper := app.Elem("sp-dialog-wrapper")

	// Add properties
	if dw.open {
		dialogWrapper = dialogWrapper.Attr("open", true)
	}
	if dw.cancelLabel != "" {
		dialogWrapper = dialogWrapper.Attr("cancel-label", dw.cancelLabel)
	}
	if dw.confirmLabel != "" {
		dialogWrapper = dialogWrapper.Attr("confirm-label", dw.confirmLabel)
	}
	if dw.dismissLabel != "" {
		dialogWrapper = dialogWrapper.Attr("dismiss-label", dw.dismissLabel)
	}
	if dw.dismissable {
		dialogWrapper = dialogWrapper.Attr("dismissable", true)
	}
	if dw.error {
		dialogWrapper = dialogWrapper.Attr("error", true)
	}
	if dw.footer != "" {
		dialogWrapper = dialogWrapper.Attr("footer", dw.footer)
	}
	if dw.headline != "" {
		dialogWrapper = dialogWrapper.Attr("headline", dw.headline)
	}
	if dw.headlineVisibility != "" {
		dialogWrapper = dialogWrapper.Attr("headline-visibility", dw.headlineVisibility)
	}
	if dw.hero != "" {
		dialogWrapper = dialogWrapper.Attr("hero", dw.hero)
	}
	if dw.heroLabel != "" {
		dialogWrapper = dialogWrapper.Attr("hero-label", dw.heroLabel)
	}
	if dw.mode != "" {
		dialogWrapper = dialogWrapper.Attr("mode", string(dw.mode))
	}
	if dw.noDivider {
		dialogWrapper = dialogWrapper.Attr("no-divider", true)
	}
	if dw.responsive {
		dialogWrapper = dialogWrapper.Attr("responsive", true)
	}
	if dw.secondaryLabel != "" {
		dialogWrapper = dialogWrapper.Attr("secondary-label", dw.secondaryLabel)
	}
	if dw.size != "" {
		dialogWrapper = dialogWrapper.Attr("size", string(dw.size))
	}
	if dw.underlay {
		dialogWrapper = dialogWrapper.Attr("underlay", true)
	}

	// Add event handlers
	if dw.onCancel != nil {
		dialogWrapper = dialogWrapper.On("cancel", dw.onCancel)
	}
	if dw.onClose != nil {
		dialogWrapper = dialogWrapper.On("close", dw.onClose)
	}
	if dw.onConfirm != nil {
		dialogWrapper = dialogWrapper.On("confirm", dw.onConfirm)
	}
	if dw.onSecondary != nil {
		dialogWrapper = dialogWrapper.On("secondary", dw.onSecondary)
	}

	// Add content if provided
	if len(dw.content) > 0 {
		dialogWrapper = dialogWrapper.Body(dw.content...)
	}

	return dialogWrapper
}
