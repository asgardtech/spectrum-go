package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumIconsWorkflow represents a component that provides access to the Spectrum Workflow icons.
// This is a wrapper around the Adobe Spectrum Web Components icons-workflow package.
type spectrumIconsWorkflow struct {
	app.Compo

	// Properties
	PropLabel string
	PropSize  IconSize
}

// IconsWorkflow creates a new icons-workflow component.
func IconsWorkflow() *spectrumIconsWorkflow {
	return &spectrumIconsWorkflow{}
}

// Label sets the accessible label for the icon.
func (i *spectrumIconsWorkflow) Label(label string) *spectrumIconsWorkflow {
	i.PropLabel = label
	return i
}

// Size sets the size of the icons.
func (i *spectrumIconsWorkflow) Size(size IconSize) *spectrumIconsWorkflow {
	i.PropSize = size
	return i
}

// Render renders the icons-workflow component.
func (i *spectrumIconsWorkflow) Render() app.UI {
	iconsWorkflow := app.Elem("sp-icons-workflow")

	// Set attributes based on properties
	if i.PropLabel != "" {
		iconsWorkflow.Attr("label", i.PropLabel)
	}

	if i.PropSize != "" {
		iconsWorkflow.Attr("size", string(i.PropSize))
	}

	return iconsWorkflow
}

// Icon utility functions for specific Workflow icons
// These are convenience methods to create specific icon elements

// AbcIcon creates an ABC icon
func AbcIcon() app.UI {
	return app.Elem("sp-icon-abc")
}

// AdobeExpressIcon creates an Adobe Express icon
func AdobeExpressIcon() app.UI {
	return app.Elem("sp-icon-adobe-express")
}

// WorkflowAlertIcon creates an Alert icon
func WorkflowAlertIcon() app.UI {
	return app.Elem("sp-icon-alert")
}

// AlignCenterIcon creates an Align Center icon
func AlignCenterIcon() app.UI {
	return app.Elem("sp-icon-align-center")
}

// AlignLeftIcon creates an Align Left icon
func AlignLeftIcon() app.UI {
	return app.Elem("sp-icon-align-left")
}

// AlignRightIcon creates an Align Right icon
func AlignRightIcon() app.UI {
	return app.Elem("sp-icon-align-right")
}

// ArrowDownIcon creates an Arrow Down icon
func ArrowDownIcon() app.UI {
	return app.Elem("sp-icon-arrow-down")
}

// ArrowLeftIcon creates an Arrow Left icon
func ArrowLeftIcon() app.UI {
	return app.Elem("sp-icon-arrow-left")
}

// ArrowRightIcon creates an Arrow Right icon
func ArrowRightIcon() app.UI {
	return app.Elem("sp-icon-arrow-right")
}

// ArrowUpIcon creates an Arrow Up icon
func ArrowUpIcon() app.UI {
	return app.Elem("sp-icon-arrow-up")
}

// BookIcon creates a Book icon
func BookIcon() app.UI {
	return app.Elem("sp-icon-book")
}

// BrushIcon creates a Brush icon
func BrushIcon() app.UI {
	return app.Elem("sp-icon-brush")
}

// CalendarIcon creates a Calendar icon
func CalendarIcon() app.UI {
	return app.Elem("sp-icon-calendar")
}

// CameraIcon creates a Camera icon
func CameraIcon() app.UI {
	return app.Elem("sp-icon-camera")
}

// WorkflowCheckmarkIcon creates a Checkmark icon
func WorkflowCheckmarkIcon() app.UI {
	return app.Elem("sp-icon-checkmark")
}

// ClockIcon creates a Clock icon
func ClockIcon() app.UI {
	return app.Elem("sp-icon-clock")
}

// CloseIcon creates a Close/X icon
func CloseIcon() app.UI {
	return app.Elem("sp-icon-close")
}

// CodeIcon creates a Code icon
func CodeIcon() app.UI {
	return app.Elem("sp-icon-code")
}

// DocumentIcon creates a Document icon
func DocumentIcon() app.UI {
	return app.Elem("sp-icon-document")
}

// DownloadIcon creates a Download icon
func DownloadIcon() app.UI {
	return app.Elem("sp-icon-download")
}

// EditIcon creates an Edit/Pencil icon
func EditIcon() app.UI {
	return app.Elem("sp-icon-edit")
}

// FolderIcon creates a Folder icon
func FolderIcon() app.UI {
	return app.Elem("sp-icon-folder")
}

// GearIcon creates a Gear/Settings icon
func GearIcon() app.UI {
	return app.Elem("sp-icon-gear")
}

// GraphIcon creates a Graph icon
func GraphIcon() app.UI {
	return app.Elem("sp-icon-graph")
}

// HomeIcon creates a Home icon
func HomeIcon() app.UI {
	return app.Elem("sp-icon-home")
}

// ImageIcon creates an Image icon
func ImageIcon() app.UI {
	return app.Elem("sp-icon-image")
}

// WorkflowInfoIcon creates an Info icon
func WorkflowInfoIcon() app.UI {
	return app.Elem("sp-icon-info")
}

// LightbulbIcon creates a Lightbulb icon
func LightbulbIcon() app.UI {
	return app.Elem("sp-icon-lightbulb")
}

// LinkIcon creates a Link icon
func LinkIcon() app.UI {
	return app.Elem("sp-icon-link")
}

// LockIcon creates a Lock icon
func LockIcon() app.UI {
	return app.Elem("sp-icon-lock")
}

// UnlockIcon creates an Unlock icon
func UnlockIcon() app.UI {
	return app.Elem("sp-icon-unlock")
}

// MailIcon creates a Mail icon
func MailIcon() app.UI {
	return app.Elem("sp-icon-mail")
}

// PlusIcon creates a Plus icon
func PlusIcon() app.UI {
	return app.Elem("sp-icon-add")
}

// MinusIcon creates a Minus icon
func MinusIcon() app.UI {
	return app.Elem("sp-icon-remove")
}

// SearchIcon creates a Search/Magnifying Glass icon
func SearchIcon() app.UI {
	return app.Elem("sp-icon-search")
}

// StarIcon creates a Star icon
func StarIcon() app.UI {
	return app.Elem("sp-icon-star")
}

// UserIcon creates a User icon
func UserIcon() app.UI {
	return app.Elem("sp-icon-user")
}

// UsersIcon creates a Users/Group icon
func UsersIcon() app.UI {
	return app.Elem("sp-icon-users")
}

// The above are just examples. In a real implementation, all Workflow icons would be added here.
