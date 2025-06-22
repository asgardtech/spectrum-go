package prism

type FormSubmitHandler[T any] func(T)

type Form[T any] struct {
	Name             string
	Description      string
	OnSubmit         []FormSubmitHandler[T]
	SubmitButtonText string
	SubmitButtonIcon string
	CancelButtonText string
	CancelButtonIcon string
	Object           T
}

func (f *Form[T]) GetName() string {
	return f.Name
}

func (f *Form[T]) GetDescription() string {
	return f.Description
}

func (f *Form[T]) WithOnSubmit(onSubmit func(T)) *Form[T] {
	f.OnSubmit = append(f.OnSubmit, onSubmit)
	return f
}

func (f *Form[T]) WithSubmitButtonText(text string) *Form[T] {
	f.SubmitButtonText = text
	return f
}

func (f *Form[T]) WithSubmitButtonIcon(icon string) *Form[T] {
	f.SubmitButtonIcon = icon
	return f
}

func (f *Form[T]) WithCancelButtonText(text string) *Form[T] {
	f.CancelButtonText = text
	return f
}

func (f *Form[T]) WithObject(object T) *Form[T] {
	f.Object = object
	return f
}
