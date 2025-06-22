# Prism

Prism is a thin glue layer that combines [Go-App](https://github.com/maxence-charriere/go-app) (Go → WebAssembly SPA) with [spectrum-go](https://github.com/asgardtech/spectrum-go) (Adobe Spectrum Web Components bindings).  
It takes care of routing, layout chrome, authentication scaffolding, forms and navigation so that you can focus on business UI in pure Go.

---

## Quick-start example – a mini "Stripe" clone

Below is an end-to-end sketch that demonstrates how Prism _should_ feel when it is fully implemented.

```text
.
├── cmd/stripe/main.go        # tiny HTTP bootstrapper
└── app/
    ├── app.go               # builds & mounts the Prism SPA
    ├── login.go             # public login page
    ├── home.go              # auth-protected home
    └── dashboard.go         # auth-protected dashboard with charts
```

### `cmd/stripe/main.go`

```go
package main

import "github.com/your-org/stripe/app"

func main() {
    app.New().Serve() // serves /static/app.wasm + API endpoints
}
```

### `app/app.go`

```go
package app

import (
    "github.com/asgardtech/spectrum-go/prism"
    "github.com/maxence-charriere/go-app/v10/pkg/app"
)

func init() {
    if app.IsBrowser() {
        buildApp().Mount()
    }
}

func buildApp() *prism.App {
    return prism.NewApp().
        WithBaseURL("/").
        WithTitle("Stripe").
        WithDescription("Payments infrastructure").
        WithPages( // fictional sugar helper
            newLoginPage(),
            newHomePage(),
            newDashboardPage(),
        )
}
```

### `app/login.go`

```go
package app

import (
    "github.com/asgardtech/spectrum-go/prism"
)

type loginPage struct { prism.LoginPage }

func newLoginPage() *loginPage {
    return (&loginPage{}).
        WithUrlPath("/login").
        WithTitle("Log in – Stripe").
        WithDescription("Access your dashboard").
        WithProviders("google", "apple", "microsoft").
        WithPasswordLogin(true).
        WithWelcomeMessage("Welcome back!")
}
```

### `app/home.go`

```go
package app

import (
    "github.com/asgardtech/spectrum-go/prism"
    "github.com/maxence-charriere/go-app/v10/pkg/app"
)

type homePage struct {
    prism.Page
    app.Compo
}

func newHomePage() *homePage {
    return (&homePage{}).
        WithName("Home").
        WithUrlPath("/").
        WithDescription("Overview").
        WithIcon("home").
        WithSidenavLink(true)
}

func (h *homePage) Render() app.UI {
    return prism.NewLayout().
        WithTitle("Home – Stripe").
        Content(
            app.H1().Text("Home"),
        )
}
```

### `app/dashboard.go`

```go
package app

import (
    "github.com/asgardtech/spectrum-go/prism"
    "github.com/maxence-charriere/go-app/v10/pkg/app"
)

type dashboardPage struct {
    prism.Page
    app.Compo

    selectedTimeRange string
}

func newDashboardPage() *dashboardPage {
    return (&dashboardPage{selectedTimeRange: "7d"}).
        WithName("Dashboard").
        WithUrlPath("/dashboard").
        WithDescription("Payments & balance").
        WithIcon("dashboard").
        WithSidenavLink(true).
        WithSidenavGroup("Analytics")
}

func (d *dashboardPage) Render() app.UI {
    return prism.NewLayout().
        WithTitle("Dashboard").
        Content(
            prism.SplitView(
                d.renderFilters(),
                d.renderCharts(),
            ),
        )
}

func (d *dashboardPage) renderFilters() app.UI {
    return prism.Panel(
        prism.FieldGroup().
            Legend("Quick filters").
            Add(
                prism.Picker().
                    Label("Range").
                    Value(d.selectedTimeRange).
                    Items("1d", "7d", "30d", "90d").
                    OnChange(func(v string) {
                        d.selectedTimeRange = v
                        app.Dispatch(func() {})
                    }),
            ),
    )
}

func (d *dashboardPage) renderCharts() app.UI {
    return prism.Grid().
        Columns("1fr 1fr").
        Gaps("size-200").
        Add(
            prism.DisplayCard().
                Title("Payment volume").
                Body(prism.LineChart()),
            prism.DisplayCard().
                Title("Success rate").
                Body(prism.PieChart()),
        )
}
```

---

### `app/payments.go` (table with actions)

```go
package app

import (
    "github.com/asgardtech/spectrum-go/prism"
    "github.com/maxence-charriere/go-app/v10/pkg/app"
)

// Payment domain model
type Payment struct {
    ID       string
    Amount   int64  // in cents
    Currency string
    Status   string // "succeeded", "pending", "failed"
}

// in-memory sample data
var payments = []Payment{
    {ID: "pay_123", Amount: 4300, Currency: "USD", Status: "succeeded"},
    {ID: "pay_124", Amount: 1200, Currency: "EUR", Status: "pending"},
}

type paymentsPage struct {
    prism.Page
    app.Compo

    rows []Payment

    // UI state indices (-1 when closed)
    editingIdx      int
    pendingDeleteIdx int
}

func newPaymentsPage() *paymentsPage {
    return (&paymentsPage{rows: payments, editingIdx: -1, pendingDeleteIdx: -1}).
        WithName("Payments").
        WithUrlPath("/payments").
        WithIcon("table").
        WithSidenavLink(true)
}

func (p *paymentsPage) Render() app.UI {
    return prism.NewLayout().
        WithTitle("Payments").
        Content(
            prism.ActionButton().
                Label("Add payment").
                Icon(prism.IconPlus).
                OnClick(func(app.Context, app.Event) { p.openNew() }),
            p.renderTable(),
            p.renderEditModal(),
            p.renderDeleteDialog(),
            prism.ToastManager(),
        )
}

func (p *paymentsPage) renderTable() app.UI {
    return prism.DataTable[Payment]().
        Dense(true).
        Columns(2, 1, 1, 1). // converts to "2fr 1fr 1fr 1fr"
        Rows(p.rows).
        RowActions(func(idx int) []prism.TableAction {
            return []prism.TableAction{
                {Label: "Edit", Icon: prism.IconEdit, Run: func(app.Context) { p.openEdit(idx) }},
                {Label: "Delete", Icon: prism.IconDelete, Intent: prism.IntentDanger, Run: func(app.Context) { p.openDelete(idx) }},
            }
        })
}

func (p *paymentsPage) renderEditModal() app.UI {
    if p.editingIdx == -1 {
        return app.Div()
    }
    current := p.rows[p.editingIdx]
    return prism.ActionModal().
        Title("Edit payment").
        Dismissible(true).
        Body(
            prism.Form[Payment]().
                WithObject(current).
                WithSubmitButtonText("Save").
                WithOnSubmit(func(pay Payment) {
                    p.rows[p.editingIdx] = pay
                    prism.ToastSuccess("Payment updated")
                }),
        ).
        OnDismiss(func(app.Context) { p.editingIdx = -1 })
}

func (p *paymentsPage) renderDeleteDialog() app.UI {
    if p.pendingDeleteIdx == -1 { return app.Div() }
    id := p.rows[p.pendingDeleteIdx].ID
    return prism.ConfirmDialog().
        Title("Delete payment").
        Description("You are about to delete " + id + ". This action cannot be undone.").
        ConfirmLabel("Delete").
        ConfirmIntent(prism.IntentDanger).
        OnConfirm(func(app.Context) {
            p.rows = append(p.rows[:p.pendingDeleteIdx], p.rows[p.pendingDeleteIdx+1:]...)
            p.pendingDeleteIdx = -1
            prism.ToastSuccess("Payment deleted")
        }).
        OnCancel(func(app.Context) { p.pendingDeleteIdx = -1 })
}

func (p *paymentsPage) openDelete(idx int) { p.pendingDeleteIdx = idx }

func (p *paymentsPage) openEdit(idx int)   { p.editingIdx = idx }

func (p *paymentsPage) openNew()           { p.editingIdx = -2 } // -2 represents new row

This illustrates:

* Strongly-typed layout tokens – `GridTemplate`, `Fr`, `IntentDanger`.
* No raw strings for sizes: you can switch to `Gap(prism.Size(200))` or constants.
* CRUD flow with modal editing, delete confirmation, and toast feedback – all via Prism helpers.

---

### Where Prism helps

1. **Routing & Navigation** – `WithPages` (or individual `WithPage`) wires `app.Route` under the hood and auto-generates side-nav / top-nav links from your page metadata.
2. **Layout Chrome** – `prism.NewLayout()` wraps every page with Spectrum theme root, side-nav, top-nav, user menu and a dark/light toggle.
3. **Auth Plumbing** – `prism.LoginPage` + `prism.Auth` take care of storing a session and flipping the UI when the user logs in or out.
4. **Design Language** – You only manipulate Spectrum components (`Picker`, `Grid`, `DisplayCard`, …) in Go; Prism keeps their API as close as possible to React Spectrum so designers can read the code too.

Copy the files above, run:

```bash
go run ./cmd/stripe  # in server mode
```

then open the served URL – you have a fully functional, Spectrum-styled SPA written 100 % in Go.

Happy building with Prism!

---

## Installation

```bash
go get github.com/asgardtech/spectrum-go/prism
```

Prism depends on Go 1.22+ and the [`tinywasm`](https://tinygo.org/) tool-chain that ships with go-app. Follow go-app's documentation to set up the WASM tooling if you haven't already.

## Extending Prism

* **Bring your own pages** – embed `prism.Page` in any `app.Compo` to get navigation metadata.
* **Custom layouts** – copy `layout_page.go`, tweak the Spectrum chrome, and use `YourLayout().Content(...)` in your pages.
* **Server-side auth** – the default `prism.LoginPage` expects an HTTP JSON API at `/api/login`.  Swap it out by implementing your own login component.

## Roadmap

1. Form validation helpers (`<FormField>.Validate(...)`).
2. Server-sent events helper for realtime dashboards.
3. Static-site export (like Next.js `export`).

> Feedback and PRs are welcome – open an issue on the repo with your ideas.
