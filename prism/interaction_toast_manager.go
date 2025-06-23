package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"time"
)

type ToastVariant string

const (
	ToastVariantInfo    ToastVariant = "info"
	ToastVariantSuccess ToastVariant = "positive"
	ToastVariantWarning ToastVariant = "warning"
	ToastVariantError   ToastVariant = "negative"
)

type Toast struct {
	ID       string
	Message  string
	Variant  ToastVariant
	Duration time.Duration
	Action   *ToastAction
}

type ToastAction struct {
	Text    string
	OnClick func(app.Context)
}

type ToastManager struct {
	app.Compo
	toasts []Toast
}

func NewToastManager() *ToastManager {
	return &ToastManager{
		toasts: make([]Toast, 0),
	}
}

func (tm *ToastManager) Show(message string, variant ToastVariant, duration time.Duration) {
	toast := Toast{
		ID:       generateID(),
		Message:  message,
		Variant:  variant,
		Duration: duration,
	}
	tm.toasts = append(tm.toasts, toast)

	if duration > 0 {
		go func() {
			time.Sleep(duration)
			tm.Remove(toast.ID)
		}()
	}
}

func (tm *ToastManager) ShowInfo(message string) {
	tm.Show(message, ToastVariantInfo, 5*time.Second)
}

func (tm *ToastManager) ShowSuccess(message string) {
	tm.Show(message, ToastVariantSuccess, 5*time.Second)
}

func (tm *ToastManager) ShowWarning(message string) {
	tm.Show(message, ToastVariantWarning, 5*time.Second)
}

func (tm *ToastManager) ShowError(message string) {
	tm.Show(message, ToastVariantError, 0) // Errors don't auto-dismiss
}

func (tm *ToastManager) ShowWithAction(message string, variant ToastVariant, action *ToastAction) {
	toast := Toast{
		ID:      generateID(),
		Message: message,
		Variant: variant,
		Action:  action,
	}
	tm.toasts = append(tm.toasts, toast)
}

func (tm *ToastManager) Remove(id string) {
	for i, toast := range tm.toasts {
		if toast.ID == id {
			tm.toasts = append(tm.toasts[:i], tm.toasts[i+1:]...)
			break
		}
	}
}

func (tm *ToastManager) Clear() {
	tm.toasts = make([]Toast, 0)
}

func (tm *ToastManager) Render() app.UI {
	if len(tm.toasts) == 0 {
		return app.Div()
	}

	return app.Div().
		Class("toast-container").
		Style("position", "fixed").
		Style("top", "16px").
		Style("right", "16px").
		Style("z-index", "1000").
		Style("display", "flex").
		Style("flex-direction", "column").
		Style("gap", "8px").
		Body(
			app.Range(tm.toasts).Slice(func(i int) app.UI {
				toast := tm.toasts[i]
				return tm.renderToast(toast)
			}),
		)
}

func (tm *ToastManager) renderToast(toast Toast) app.UI {
	return app.Div().
		Class("spectrum-Toast").
		Style("background", "#2c2c2c").
		Style("color", "white").
		Style("padding", "12px 16px").
		Style("border-radius", "4px").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "space-between").
		Style("min-width", "300px").
		Body(
			app.Span().
				Text(toast.Message),
			app.Div().
				Style("display", "flex").
				Style("gap", "8px").
				Body(
					app.If(toast.Action != nil, func() app.UI {
						return sp.Button().
							Text(toast.Action.Text).
							OnClick(func(ctx app.Context, e app.Event) {
								if toast.Action.OnClick != nil {
									toast.Action.OnClick(ctx)
								}
							})
					}),
					sp.Button().
						Text("×").
						OnClick(func(ctx app.Context, e app.Event) {
							tm.Remove(toast.ID)
						}),
				),
		)
}

func generateID() string {
	return "toast-" + string(rune(time.Now().UnixNano()))
}
