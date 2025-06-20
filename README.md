# Spectrum Go

Spectrum-Go is a **strongly-typed Go wrapper** around the official [Adobe Spectrum Web Components](https://opensource.adobe.com/spectrum-web-components/) that runs on top of the [go-app](https://github.com/maxence-charriere/go-app) WebAssembly framework.

With Spectrum-Go you can:

* Author web front-ends **entirely in Go** – no JavaScript, TypeScript or CSS build-pipelines required.
* Compose production-ready UI primitives that follow the Adobe Spectrum design system and accessibility contract.
* Retain **static typing & IDE assistance** when declaring UI (every component, attribute, slot and event is represented by a Go method).
* Ship the same binary to desktop, mobile and server without sacrificing fast, incremental development loops.

> This README explains **what** the project does, **how** it is built and **how to contribute**. The original exhaustive component checklist is preserved further down and is referred to by the navigation links below.

---

## Table of contents

- [Spectrum Go](#spectrum-go)
  - [Table of contents](#table-of-contents)
  - [Why does this project exist?](#why-does-this-project-exist)
  - [High-level architecture](#high-level-architecture)
  - [Getting started](#getting-started)
    - [Installation](#installation)
    - [Minimal example](#minimal-example)
    - [Running the showcase locally](#running-the-showcase-locally)
  - [Project layout](#project-layout)
  - [Development workflow](#development-workflow)
  - [Code-generation pipeline](#code-generation-pipeline)
  - [Contributing](#contributing)
  - [License](#license)
  - [Components TODO](#components-todo)
  - [Tools and Utilities](#tools-and-utilities)
  - [Design Principles](#design-principles)
  - [Implementation Status](#implementation-status)
  - [Showcase Implementation](#showcase-implementation)

---

## Why does this project exist?

JavaScript frameworks evolve rapidly whereas the Go ecosystem values long-term stability.  When building PWAs in Go with `go-app` developers had to re-implement UI primitives themselves or embed large JS bundles.  Spectrum-Go bridges that gap by providing **first-class Go bindings** for every [Spectrum Web Component](https://opensource.adobe.com/spectrum-web-components/components) while preserving the original HTML semantics.

Key goals:

* **1-to-1 parity** with Spectrum Web Components (attributes, slots, events and enum values).
* **Ergonomic, chainable API** – every call returns the component so you can configure fluently.
* **Zero-runtime cost** – Render functions just emit `<sp-*>` tags, the browser downloads Adobe's already-optimised bundle from a CDN.
* **Generated, not handwritten** – bindings stay in sync with upstream docs, vastly reducing maintenance burden.

---

## High-level architecture

```mermaid
flowchart TD
  subgraph Go[Go (compile-time)]
    A[JSON definitions ↲\n(spectrum-docs/json)] --> B[Python ‑> Jinja template\n`generate_components.py`]
    B --> C[Generated `sp_*.go` files]
  end
  subgraph Runtime
    D[Your Go code] -->|imports| C
    D -->|compiles to WASM| browser[🖥️ Browser]
    browser -->|lazy-loads| cdn[(jspm.dev/@spectrum-web-components/bundle/elements.js)]
  end
```

* **Component bindings (`sp_*.go`)** – each file defines a Go struct embedding `app.Compo`, plus helper generics for CSS classes, inline styles and IDs.  The `Render()` method merely instantiates the corresponding `<sp-foo>` custom element with the configured attributes.
* **Utility mix-ins (`0_*.go`)** – small generic builders (`classer`, `styler`, `ider`) shared by every component keep the public API consistent while minimising code generation size.
* **System component (`0_system.go`)** – optional bootstrap that hooks into the DOM to turn plain `<a href="/page">` links into client-side navigation using `go-app`'s router.
* **Showcase** – full featured demo application living under `showcase/`, compiled to WASM and served by a small Go HTTP server.

---

## Getting started

### Installation

```bash
go get github.com/asgardtech/spectrum-go
```

### Minimal example

```go
package main

import (
    "github.com/maxence-charriere/go-app/v10/pkg/app"
    sp "github.com/asgardtech/spectrum-go"
)

func main() {
    app.Route("/", func() app.UI {
        return sp.Button().Variant(sp.ButtonPrimary).Label("Hello Spectrum-Go 🪄")
    })

    // optional: intercept <a> clicks and keep SPA routing
    app.Route("/", sp.System())

    app.Run()
}
```

Build & run it as usual with `go-app`:

```bash
GOOS=js GOARCH=wasm go run .
```

### Running the showcase locally

```bash
# Build the WebAssembly once
make build-wasm

# Start the http server + auto-refresh (requires air)
make run
# or, if you have air installed
make watch
```

Then navigate to http://localhost:7777 and explore every component live.

---

## Project layout

```text
spectrum-go/
├─ sp_*.go                 // generated component bindings
├─ 0_*.go                  // shared generic helpers / system bootstrap
├─ showcase/               // live demo application (WASM)
│   ├─ *.go                // individual demo pages
│   └─ web/                // compiled artefacts (app.wasm, app.css)
└─ spectrum-docs/          // doc scraping + code-gen pipeline
    ├─ download-docs.py    // pulls upstream HTML into spectrum-docs/html
    ├─ process-docs.py     // sanitises HTML -> markdown
    ├─ json/               // machine-readable component API (generated)
    ├─ generate_components.py // JSON -> Go bindings
    └─ templates/component.go.j2
```

---

## Development workflow

1. Hack on Go code as usual ‑ `go run ./showcase` will hot-reload thanks to [air](https://github.com/cosmtrek/air) if you use `make watch`.
2. When Adobe publishes a new component *or* changes an API:
   1. `cd spectrum-docs`
   2. `python3 download-docs.py` – downloads the fresh HTML docs.
   3. `python3 process-docs.py` – converts HTML ➜ markdown + extracts machine-readable data.
   4. `python3 generate_components.py` – re-creates/updates every `sp_*.go` file.
   5. Run `go vet ./... && go test ./...` and commit.

All generator scripts are deterministic; their output should always be checked-in so consumers don't need Python.

---

## Code-generation pipeline

*The following is only required when adding/updating bindings – most users can skip this section.*

1. Upstream HTML is saved under `spectrum-docs/html/` (step 1 above).
2. `process-docs.py` strips marketing content, converts tables to Markdown, and writes to `spectrum-docs/md/`.
3. During the same pass it emits canonical **JSON** describing attributes, slots and events to `spectrum-docs/json/`.
4. `generate_components.py` loads those JSON files and streams them through a Jinja2 template which produces idiomatic Go source files.
5. A final `gofmt -w` run guarantees consistent formatting.

Because everything is declarative, supporting a *new* Spectrum component typically means:

```bash
# inside spectrum-docs
python3 download-docs.py --only new-component
python3 process-docs.py
python3 generate_components.py
```

No hand-written Go is necessary 🚀

---

## Contributing

Pull requests are very welcome!  Please open an issue first if you plan to introduce a larger change (new generator features, architectural refactors, …) so we can coordinate.

* The codebase targets **Go ≥1.22**.
* `go vet ./...`, `go test ./...` and `gofmt` must pass ‑ CI will enforce this.
* Newly generated files belong in the same commit as the generator update so the repo stays buildable at every revision.

---

## License

Apache-2.0

---

## Components TODO

- [X] [Button](spectrum-docs/md/button.md) ✅
- [x] [Accordion](spectrum-docs/md/accordion.md) ✅
- [x] [Accordion Item](spectrum-docs/md/accordion-item.md) ✅
- [x] [Action Bar](spectrum-docs/md/action-bar.md) ✅
- [x] [Action Button](spectrum-docs/md/action-button.md) ✅
- [x] [Action Group](spectrum-docs/md/action-group.md) ✅
- [x] [Action Menu](spectrum-docs/md/action-menu.md) ✅
- [x] [Alert Banner](spectrum-docs/md/alert-banner.md) ✅
- [x] [Alert Dialog](spectrum-docs/md/alert-dialog.md) ✅
- [x] [Asset](spectrum-docs/md/asset.md) ✅
- [x] [Avatar](spectrum-docs/md/avatar.md) ✅
- [x] [Badge](spectrum-docs/md/badge.md) ✅
- [x] [Breadcrumb Item](spectrum-docs/md/breadcrumb-item.md) ✅
- [x] [Breadcrumbs](spectrum-docs/md/breadcrumbs.md) ✅
- [x] [Button Group](spectrum-docs/md/button-group.md) ✅
- [x] [Card](spectrum-docs/md/card.md) ✅
- [x] [Checkbox](spectrum-docs/md/checkbox.md) ✅
- [x] [Coach Indicator](spectrum-docs/md/coach-indicator.md) ✅
- [x] [Coachmark](spectrum-docs/md/coachmark.md) ✅
- [x] [Color Area](spectrum-docs/md/color-area.md) ✅
- [x] [Color Field](spectrum-docs/md/color-field.md) ✅
- [x] [Color Handle](spectrum-docs/md/color-handle.md) ✅
- [x] [Color Loupe](spectrum-docs/md/color-loupe.md) ✅
- [x] [Color Slider](spectrum-docs/md/color-slider.md) ✅
- [x] [Color Wheel](spectrum-docs/md/color-wheel.md) ✅
- [x] [Combobox](spectrum-docs/md/combobox.md) ✅
- [x] [Contextual Help](spectrum-docs/md/contextual-help.md) ✅
- [x] [Dialog](spectrum-docs/md/dialog.md), [Dialog Base](spectrum-docs/md/dialog-base.md), [Dialog Wrapper](spectrum-docs/md/dialog-wrapper.md) ✅
- [x] [Dropzone](spectrum-docs/md/dropzone.md) ✅
- [x] [Field Group](spectrum-docs/md/field-group.md) ✅
- [x] [Field Label](spectrum-docs/md/field-label.md) ✅
- [x] [Help Text](spectrum-docs/md/help-text.md) ✅
- [x] [Help Text Mixin](spectrum-docs/md/help-text-mixin.md) ✅
- [x] [Icon](spectrum-docs/md/icon.md) ✅
- [x] [Icons UI](spectrum-docs/md/icons-ui.md) ✅
- [x] [Icons Workflow](spectrum-docs/md/icons-workflow.md) ✅
- [x] [Illustrated Message](spectrum-docs/md/illustrated-message.md) ✅
- [x] [Infield Button](spectrum-docs/md/infield-button.md) ✅
- [x] [Link](spectrum-docs/md/link.md) ✅
- [x] [Menu](spectrum-docs/md/menu.md) ✅
- [x] [Menu Group](spectrum-docs/md/menu-group.md) ✅
- [x] [Menu Item](spectrum-docs/md/menu-item.md) ✅
- [x] [Meter](spectrum-docs/md/meter.md) ✅
- [x] [Number Field](spectrum-docs/md/number-field.md) ✅
- [x] [Overlay](spectrum-docs/md/overlay.md) ✅
- [x] [Overlay Trigger](spectrum-docs/md/overlay-trigger.md) ✅
- [x] [Picker](spectrum-docs/md/picker.md) ✅
- [x] [Picker Button](spectrum-docs/md/picker-button.md) ✅
- [x] [Popover](spectrum-docs/md/popover.md) ✅
- [x] [Progress Bar](spectrum-docs/md/progress-bar.md) ✅
- [x] [Progress Circle](spectrum-docs/md/progress-circle.md) ✅
- [x] [Radio](spectrum-docs/md/radio.md) ✅
- [x] [Radio Group](spectrum-docs/md/radio-group.md) ✅
- [x] [Search](spectrum-docs/md/search.md) ✅
- [x] [Sidenav](spectrum-docs/md/sidenav.md) ✅
- [x] [Sidenav Heading](spectrum-docs/md/sidenav.md) ✅
- [x] [Sidenav Item](spectrum-docs/md/sidenav.md) ✅
- [x] [Sidenav Overflow](spectrum-docs/md/sidenav-overflow.md) ✅
- [x] [Slider](spectrum-docs/md/slider.md) ✅
- [x] [Slider Handle](spectrum-docs/md/slider-handle.md) ✅
- [x] [Split View](spectrum-docs/md/split-view.md) ✅
- [x] [Status Light](spectrum-docs/md/status-light.md) ✅
- [x] [Swatch](spectrum-docs/md/swatch.md) ✅
- [x] [Swatch Group](spectrum-docs/md/swatch-group.md) ✅
- [x] [Switch](spectrum-docs/md/switch.md) ✅
- [x] [Tab](spectrum-docs/md/tab.md) ✅
- [x] [TabPanel](spectrum-docs/md/tabpanel.md) ✅
- [x] [Tabs](spectrum-docs/md/tabs.md) ✅
- [x] [Tabs Overflow](spectrum-docs/md/tabs-overflow.md) ✅
- [x] [Tag](spectrum-docs/md/tag.md) ✅
- [x] [Tags](spectrum-docs/md/tags.md) ✅
- [x] [Textarea](spectrum-docs/md/textarea.md) ✅
- [x] [Textfield](spectrum-docs/md/textfield.md) ✅
- [x] [Thumbnail](spectrum-docs/md/thumbnail.md) ✅
- [x] [Toast](spectrum-docs/md/toast.md) ✅
- [x] [Tooltip](spectrum-docs/md/tooltip.md) ✅
- [x] [Top Nav](spectrum-docs/md/top-nav.md) ✅
- [x] [Tray](spectrum-docs/md/tray.md) ✅
- [x] [Underlay](spectrum-docs/md/underlay.md) ✅

## Tools and Utilities

- [x] [Base](spectrum-docs/md/base.md) ✅ - Base class functionality
- [x] [Bundle](spectrum-docs/md/bundle.md) ✅ - Bundle management
- [x] [Color Controller](spectrum-docs/md/color-controller.md) ✅ - Color management utilities
- [x] [Core Tokens](spectrum-docs/md/core-tokens.md) ✅ - Design token system
- [x] [Dependency Manager](spectrum-docs/md/dependency-manager.md) ✅ - Manages component dependencies
- [x] [Element Resolution](spectrum-docs/md/element-resolution.md) ✅ - Element resolution utilities
- [x] [Grid](spectrum-docs/md/grid.md) ✅ - Grid layout system
- [x] [Match Media](spectrum-docs/md/match-media.md) ✅ - Media query utilities
- [x] [Opacity Checkerboard](spectrum-docs/md/opacity-checkerboard.md) ✅ - Visual aid for transparency
- [x] [Pending State](spectrum-docs/md/pending-state.md) ✅ - State management for async operations
- [x] [Reactive Controllers](spectrum-docs/md/reactive-controllers.md) ✅ - Controllers for reactive patterns
- [x] [Roving Tab Index](spectrum-docs/md/roving-tab-index.md) ✅ - Accessibility navigation utility
- [x] [Shared](spectrum-docs/md/shared.md) ✅ - Shared utilities and functionality
- [x] [Styles](spectrum-docs/md/styles.md) ✅ - Style management system
- [X] [Theme](spectrum-docs/md/theme.md) ✅ - Theming functionality
- [x] [Truncated](spectrum-docs/md/truncated.md) ✅ - Text truncation with tooltips

## Design Principles

Our implementation follows these key principles:

1. **Consistent API Design**: Every component follows the same patterns and style:
   - Type definitions for variants and constants
   - Struct representing the component with all properties
   - Constructor function named after the component
   - Fluent interface with method chaining
   - Render method that creates the web component with attributes

2. **Strong Typing**: Container components (like Tags and Accordion) accept only valid child component types for type safety.

3. **Default Values**: Components provide sensible defaults (size, treatment, etc.) while allowing customization.

4. **Separation of Concerns**: Each component is self-contained in its own file.

5. **Slot Management**: Components that accept slotted content provide clear APIs for setting those elements.

6. **Event Handling**: Components use `app.EventHandler` for proper event binding.

7. **Attribute Mapping**: Proper mapping between Go properties and HTML attributes, maintaining Adobe Spectrum naming conventions.

## Implementation Status

Currently implemented components:

1. **Button** - Full implementation with icon support and all variants
2. **Icon** - Basic implementation
3. **Accordion** - Implementation with support for different sizes and densities
4. **Accordion Item** - Implementation with support for disabled and open states
5. **Switch** - Implementation with support for all sizes and states
6. **Tag** - Implementation with icon/avatar slots and deletable state
7. **Tags** - Container implementation for Tag components
8. **Radio** - Implementation with support for sizes and states like emphasized and invalid
9. **Radio Group** - Container implementation for Radio components with orientation options and help text slots
10. **Checkbox** - Implementation with support for sizes, emphasized variant, and indeterminate state
11. **Link** - Implementation with variants, static colors, and download support
12. **Progress Bar** - Implementation with indeterminate state, labels, and size options
13. **Progress Circle** - Implementation with indeterminate state and size options
14. **Divider** - Implementation with size options and vertical orientation support
15. **Textfield** - Implementation with support for all input types, states, and help text slots
16. **Textarea** - Implementation with multiline support, grows option, and rows configuration
17. **Help Text** - Implementation with support for variants, sizes, and icon
18. **Field Label** - Implementation with size options and side alignment
19. **Field Group** - Implementation with horizontal/vertical layouts and help text slots
20. **Tooltip** - Implementation with placement options, variants, and self-managed mode
21. **Status Light**: Implemented with size options (small, medium, large, extra large) and multiple semantic variants (info, notice, positive, negative, neutral)
22. **Tab**: Implemented with value, label, vertical orientation, icon support and disabled state.
23. **Tab Panel**: Implemented with value property for association with tabs.
24. **Tabs**: Implemented with size options, direction options (horizontal, vertical, vertical-right), and features like compact, quiet, emphasized, and auto-activation. The Tabs Overflow component enables horizontal scrolling when tab content exceeds available width.
25. **Action Bar**: Implemented with emphasized styling, variants (fixed, sticky), and flexible sizing options.
26. **Action Button**: Implemented with sizes (XS to XL), quiet/emphasized variants, hold affordance, toggle behavior, and static colors.
27. **Action Group**: Implemented with horizontal/vertical layouts, single/multiple selection modes, and justified options for evenly distributed buttons.
28. **Badge**: Implemented with semantic variants (neutral, informative, positive, negative, notice), non-semantic color variants, size options, and fixed positioning support.
29. **Alert Banner**: Implemented with info/negative variants, dismissible option, action slot for buttons, and close event handling.
30. **Avatar**: Implemented with multiple size options (50-700), image source support, accessible labeling, and link functionality.
31. **Button Group**: Implemented with support for horizontal and vertical orientations to organize related buttons with proper spacing.
32. **Card**: Implemented with standard/gallery/quiet variants, horizontal layout option, asset types (file/folder), cover photos, and linking capabilities.
33. **Color Area**: Implemented with support for custom sizes, axis labels, color selection, and event handling for color changes.
34. **Search**: Implemented with size options, quiet variant, form features (action/method), and specialized behavior like holdValueOnEscape.
35. **Meter**: Implemented with size options, variants (informative, positive, notice, negative), side-label support, and static color options.
36. **Number Field**: Implemented with numeric input features, min/max ranges, step controls, stepper UI, and advanced formatting options for percentages, currency, and units.
37. **Slider**: Implemented with size options, multiple variants (filled, tick, ramp, range), label visibility control, editable mode with number field, and custom formatting.
38. **Slider Handle**: Implemented with support for multi-handle sliders, range constraining via "previous" and "next" handles, and custom formatting.
39. **Swatch**: Implemented with customizable sizes, shapes (square/rectangle), rounding options, border styles, color representation, selected state, and disabled state with event handling for selections.
40. **Swatch Group**: Implemented with selection modes (single/multiple), density options (compact/spacious), and consistent styling across grouped swatches including shape, size, rounding, and border properties.
41. **Overlay**: Implemented with various placement options, trigger interaction types (click, hover, longpress), and overlay types (auto, hint, manual, modal, page) to create contextual content on top of the base UI.
42. **Overlay Trigger**: Implemented with support for multiple content types (click, hover, longpress), various trigger models (inline, replace, modal), and flexible content placement for creating interactive overlay experiences.
43. **Popover**: Responsive popover component with customizable placement, dialog mode, and tip options.
44. **Dialog, Dialog Base, Dialog Wrapper**: Implemented with configurable modes (fullscreen, takeover), sizes, dismissable option, and support for hero images, headers, content, and footer sections.
45. **Menu, Menu Item, Menu Group**: Implemented with single/multiple selection support, grouping, dividers, header content, icons in menu items, value display, submenu capabilities, and disabled state management.
46. **Picker, Picker Button**: Implemented with dropdown menu functionality, support for icon-only and icon+text modes, pending state, customizable placement, and accessible labeling options.
47. **Sidenav**: Full-featured navigation component with support for multilevel navigation, headings, and items with selection states. Includes proper ARIA attributes and keyboard navigation support. The Sidenav Overflow component handles responsive navigation with hidden items.
48. **Breadcrumbs, Breadcrumb Item**: Implemented with support for compact display, custom root items, maximum visible items control, and overflow management. Includes link functionality, accessibility features, and customizable action menu.
49. **Split View**: Implemented with horizontal and vertical orientations, resizable panels with min/max constraints, collapsible panels, and configurable splitter position with proper accessibility features.
50. **Toast**: Implemented with multiple variants (neutral, negative, positive, info, error, warning), timeout auto-dismissal, action buttons, and proper accessibility attributes.
51. **Top Nav**: Implemented with flexible site navigation that supports URL-based selection, quiet variant, multiple items with link functionality, and accessibility labeling.
52. **Tray**: Implemented as a mobile-optimized overlay container that slides up from the bottom of the screen, supporting content like dialogs and menus with proper event handling.
53. **Underlay**: Implemented as a blocking layer that sits between overlay content and the page, providing visual separation for modal dialogs and other overlay elements.
54. **Action Menu**: Implemented as an action button that triggers an overlay with menu items for activation. Supports selection modes, custom icons, multiple sizes, and mobile-specific display options with proper accessibility features.
55. **Alert Dialog**: Implemented with multiple variants (confirmation, information, warning, error, destructive, secondary) to display important information that users need to acknowledge. Supports customizable headings, content, and multiple button configurations with proper accessibility.
56. **Asset**: Implemented to visually represent files, folders, or images in an application. File and folder representations center themselves in the provided space, while images are properly contained and centered within the element.
57. **Coach Indicator**: Implemented to show the connection between an object and its explanation in a touring mode. Supports standard and quiet variants with customizable static colors (dark, light, white, black).
58. **Coachmark**: Implemented as a temporary message that educates users through new or unfamiliar product experiences, which can be chained into a sequence to form a tour. Supports media content, shortcut keys, CTA buttons, and comprehensive event handling.
59. **Color Handle**: Implemented as a draggable handle for selecting colors on color selection components. Features include customizable color representation, disabled and focused states, and open state for displaying a color loupe.
60. **Color Loupe**: Implemented as a visual aid that shows the output color during color selection, which would otherwise be covered by a cursor, stylus, or finger. Features customizable color and visibility state.
61. **Color Slider**: Implemented as a slider that allows users to visually change an individual channel of a color. Supports color properties like hues, color channel values, and opacity with vertical orientation option, step configuration, and comprehensive event handling.
62. **Color Wheel**: Implemented as a circular control that lets users visually change an individual channel of a color on a circular track. Supports various color formats, customizable step size, disabled and focused states, and comprehensive event handling.
63. **Combobox**: Implemented as a control that allows users to filter lists to only options matching a query. Features include different autocomplete modes (none, list), multiple size variants, support for dynamic options, quiet variant, and slots for help text and tooltips.
64. **Contextual Help**: Implemented as a component that shows extra information about an adjacent component or entire view, with info and help variants. Features customizable placement options, maximum width, offset configuration, and slots for heading, content, and link.
65. **Dropzone**: Implemented as an area into which objects can be dragged and dropped to accomplish tasks like file uploads. Features support for different drop effects (copy, move, link), visual states for dragged and filled conditions, and comprehensive drag and drop event handling.
66. **Help Text Mixin**: Implemented as a utility for managing help text in form control components. Features include a manager class that handles text rendering based on validity state, API for setting help text and negative help text, and helper functions for creating form components with properly associated help text.
67. **Illustrated Message**: Implemented as a component that displays an illustration icon and a message, typically used in empty states or error pages, and inside dropzone components. Features include customizable heading, description, and support for SVG illustrations with proper slot management.
68. **Infield Button**: Implemented as a button component designed for form fields that visually associates button functionality with other form controls. Features include inline (start/end) and block (stacked) positioning, size variants (S-XL), quiet visual style, and comprehensive event handling.
69. **Thumbnail**: Implemented as a component for displaying previews of images, layers, or effects. Features include multiple size options (50px to 1000px), customizable background for letterboxing non-square content, cover mode for filling available space, layer mode with selection state for layer management, and support for focused and disabled states.
70. **Icon**: Implemented to render icons through multiple methods: from registered icon sets via name attribute, from images via src attribute, or from custom SVG elements via slotted content. Features include multiple size variants (xxs to xxl), accessible labeling that properly toggles aria-hidden, and flexible positioning through slot assignment.
71. **Icons UI**: Implemented as a wrapper for Adobe Spectrum UI icons, providing convenient access to common interface icons. Features include utility functions for specific icons like arrows, checkmarks, and alert indicators, size customization, and accessible labeling.
72. **Icons Workflow**: Implemented as a wrapper for Adobe Spectrum Workflow icons, providing access to application-specific icons for various workflows and operations. Features include convenience functions for commonly used icons like document, folder, user, edit, as well as customizable size and accessible labeling.

All components follow the same design patterns and use a fluent interface for configuration.

## Showcase Implementation

A live showcase application has been created to demonstrate the usage of various components. The showcase allows developers to see components in action and explore their different features and configurations.

Currently showcased components:

1. **Button** - Demonstrates variants (primary, secondary, CTA), sizes, icons, and states
2. **Checkbox** - Shows size variants, emphasized mode, indeterminate state, and disabled state
3. **Badge** - Displays semantic and static variants, positioning options, and sizes
4. **Radio/RadioGroup** - Shows orientation options, sizes, and states with form integration
5. **Switch** - Demonstrates size variants, emphasized style, and states
6. **Progress Bar** - Shows basic progress, indeterminate state, labels, and sizes
7. **Progress Circle** - Displays basic circular progress, indeterminate state, and sizes
8. **Link** - Shows variants (primary, secondary, quiet), disabled state, static colors, and targets
9. **Icon** - Demonstrates various icons, sizes, custom colors, and usage in components

The showcase is built using Go-App and serves as a practical reference for developers using the Spectrum Go library.