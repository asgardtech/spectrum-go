package prism

type App struct {
	pages       []Page
	baseURL     string
	title       string
	description string
}

func NewApp() *App {
	return &App{}
}

func (a *App) WithPage(page Page) *App {
	a.pages = append(a.pages, page)
	return a
}

func (a *App) WithBaseUrl(baseUrl string) *App {
	a.baseURL = baseUrl
	return a
}

func (a *App) WithTitle(title string) *App {
	a.title = title
	return a
}

func (a *App) WithDescription(description string) *App {
	a.description = description
	return a
}

func (a *App) Mount() {
	// TODO: implement routing with go-app once pages are fully defined.
}
