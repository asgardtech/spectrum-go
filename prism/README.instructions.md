# README.instructions.md

> Instructions for GPT-4o in Cursor to implement the `prism` framework (Go + WASM)

---

## 🚧 Ground Rules

* 🐹 Go is used for **both** client and server
* ⚙️ All code lives in a **flat package** called `prism`
* 🧱 Follow DDD-style layering (entities, state, behavior)
* 🚫 Avoid premature abstractions — defer packaging/internalization until needed
* ✅ Prefer functional, boring code over magic

---

## 🧭 Implementation Order (Step-by-Step)

### 📁 1. File + Structure

* Create flat files in `prism/` using these patterns:

  * `page_<name>.go` – UI page definitions (e.g. `page_login.go`)
  * `component_<name>.go` – standalone UI widgets (e.g. `component_modal.go`)
  * `entity_<name>.go` – universal business entities (e.g. `entity_user.go`)
  * `state_<name>.go` – view model & logic (e.g. `state_login.go`)
  * `util_<name>.go` – helper logic (e.g. `util_forms.go`)
  * `dev_<name>.go` – dev-only test/debug utils (e.g. `dev_fakedata.go`)

### 📦 2. Core Components

* `Page`
* `TopNav`, `SideNav`, `UserMenu`
* `Modal`, `Alert`, `Card`, `List`, `Button`, `Form`

### 🔒 3. Auth & Session

* Clerk.dev auth integration (Google, GitHub, etc)
* Login/logout flow with `page_login.go`
* Session context injection

### 📊 4. State Handling

* Define page-level state struct per screen
* Include derived helpers (e.g. `ShouldShowEmpty()`)
* Create generic `LoaderState`, `ViewModel<T>` patterns

### 🧪 5. Dev & Test Infra

* Rod-based browser E2E in `tests/e2e/`
* Fake user / session / payload generators
* Debug tools (`DebugDump()`, etc)

### 🔁 6. Hot Reload Flow

* Air watches and rebuilds
* Dev server emits WebSocket "reload"
* Browser auto-refreshes

---

## 🧩 Output Goals (in README.md)

Update the public README to include:

* How to render a page: `prism.HomePage(state)`
* How to use modals: `prism.Modal(...)`
* How to manage state: `LoginPageState{...}`
* How to add new pages
* Dev server usage instructions
* Test runner command

---

## 🧠 Style Constraints

* One struct per page for state
* One public render function per component
* No side effects inside render funcs
* UI = pure function of state

> You’re building the holy grail of boring, safe, functional UI. Stick to it.
