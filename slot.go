package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type slot interface {
	Slot(name string) app.UI
}

type noOpSlot struct{}

func (noOpSlot) Slot(name string) app.UI {
	return nil
}

func isSlot(v any) bool {
	_, ok := v.(slot)
	return ok
}

func toSlot(v any) slot {
	if !isSlot(v) {
		return noOpSlot{}
	}

	return v.(slot)
}
