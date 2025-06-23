# 🧠 The Holy Grail of UI Frameworks (Go + WASM Edition)

A pragmatic, composable, Go-native approach to building scalable, fast, and enjoyable UI frameworks for fullstack SaaS apps. This is your blueprint to build once, ship fast, and stay sane.

---

## 🏗️ Core Principles

### 1. **Predictability**

* Naming, props, structure are consistent across all components.
* If it works one way in one place, it works the same everywhere.

### 2. **Composability**

* Every piece can be combined or nested.
* Components accept children and callbacks — no special wrappers needed.

### 3. **Progressive Complexity**

* Use simple defaults, then override or extend when needed.
* `Form()` works out of the box, `FormWithLayout()` gives you control.

### 4. **Convention over Configuration**

* Layouts, spacing, inputs, alerts — all have defaults that “just work.”
* You override only when necessary.

### 5. **Context Awareness**

* Components adapt based on context (modals, pages, etc).
* `Button` inside `Modal` has proper spacing without manual hacks.

### 6. **Visual & Semantic Hierarchy**

* Pages → Sections → Cards → Fields.
* Logical mental model that matches visual output.

---

## ✅ Attributes of Great Frameworks

| Attribute              | Why It Matters                                               |
| ---------------------- | ------------------------------------------------------------ |
| Flat mental model      | Everything is in `prism.*`, no package hell                  |
| Low boilerplate        | Most pages/components under 100 lines                        |
| Built-in UX primitives | Modals, toasts, filters, pagination, etc. are first-class    |
| Opinionated layout     | Consistent sectioning and spacing                            |
| Theme extensibility    | Global overrides + local tweaks                              |
| Anti-fragile scaling   | Your 2-page app and 30-page app use the same building blocks |

---

## 🧱 State Management

### 🔒 Global State — Keep it Minimal

Only include things like:

* Current user
* Feature flags
* Current org/account
* Theme/lang

**NOT** modal state, forms, tabs, etc.

### 🧠 Page State — Explicit and Local

```go
type ExportListPageState struct {
  Filters   ExportFilters
  Payouts   []Payout
  Selected  *Payout
  Loading   bool
  ShowModal bool
}
```

Pass this down. Treat it like scoped DDD.

### ⚙️ Component State — Own It or Bubble It

* Local state stays in the component.
* If it must integrate with parent: `onToggle()`, `onSubmit(data)`.

### 🧩 ViewModels for Shared Patterns

```go
type DataTableViewModel[T any] struct {
  Items     []T
  Filters   map[string]string
  Loading   bool
  ModalOpen bool
  ModalData *T
}
```

---

## 🔁 Derived State via Helpers

```go
func (s *ExportListPageState) ShouldShowEmpty() bool {
  return len(s.Payouts) == 0 && !s.Loading
}
```

Clean declarative rendering, no inline madness.

---

## 🧠 Developer Constraints & Guardrails

| Rule                                      | Reason                                      |
| ----------------------------------------- | ------------------------------------------- |
| Global state must be read-only            | Prevents spaghetti                          |
| Page state = one struct                   | Easy to trace, test, debug                  |
| Derived state lives in helpers            | Keeps templates readable                    |
| No magic modal visibility                 | Always managed by state struct              |
| Use nil-safe pointer patterns             | Avoid runtime panics                        |
| Never mutate deeply nested state directly | Wrap in `SetX()` or `WithX()` style helpers |

---

## 😱 Easy to Forget But Critical

### 1. **Loading & Error Boundaries**

* `LoaderState` struct with fallback rendering
* `RenderWithState(state, content)` to handle all 3 states

### 2. **Permission Guards**

* `if !user.CanView(x) { return Render403() }`
* Wrap components in `IfAuthorized(...)`

### 3. **Empty States**

* `RenderEmptyState(icon, title, message, actionBtn)`
* First-time UX = retention

### 4. **Email Handling**

* Welcome, passwordless, export-ready notifications
* Wrap all sending behind `SendEmail()` abstraction

### 5. **Logging & Audit**

* `Log(event, actor, target)` — stash for later
* `AuditLog{UserID, Action, Target, At}`

### 6. **URL State & Deep Linking**

* Parse `?payout_id=xyz` on page load
* Open modal or scroll to item

### 7. **Session Expiry UX**

* Token timeout = "you’ve been logged out" overlay
* Don’t crash silently

### 8. **Panic Recovery**

* `defer func() { recover(); RenderSomethingWentWrong() }`
* Production apps must fail gracefully

### 9. **Debug Tools for Devs**

* `DebugDump(v)`
* `prism_dev.go` for logging/test helpers
* Fake user generators for local state mocking

---

## 🪜 Progressive App Scaling

### Level 1: Static Pages

* `HomePage`, `LoginPage`, `About`, etc.

### Level 2: CRUD UI with Forms & Tables

* `DataTable`, `EntityEditor`, `Modal`, etc.

### Level 3: Auth, Billing, Multi-User

* Clerk, Stripe, user settings

### Level 4: Org Support, Roles, Invites

* `User`, `Organization`, `AuditLog`, `Role`

### Level 5: Real-time + Multi-tenant SaaS

* Event streaming, audit, analytics, scoped billing

---

## 🏁 Summary

> The holy grail is a framework that:
>
> * Helps you ship faster, not slower
> * Encourages composition, not complexity
> * Makes you *want* to build another app

Keep everything flat in `prism/` until you feel pain — then extract. Trust your nose. Ship first, abstract second.

**Make boring, safe, functional things. Just like a CSV.**
