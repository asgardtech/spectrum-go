# sp-button
Overview API
## Overview
An `<sp-button>` represents an action a user can take. sp-buttons can be clicked or tapped to perform an action or to navigate to another page. Buttons in Spectrum have several variations for different uses and multiple levels of loudness for various attention-getting needs.
### Usage
    
    yarn add @spectrum-web-components/button
    
Import the side effectful registration of `<sp-button>` or `<sp-clear-button>` as follows:
    
    import '@spectrum-web-components/button/sp-button.js';
    import '@spectrum-web-components/button/sp-clear-button.js';
    import '@spectrum-web-components/button/sp-close-button.js';
    
When looking to leverage the `Button`, `ClearButton`, or `CloseButton` base classes as a type and/or for extension purposes, do so via:
    
    import { Button, ClearButton, CloseButton } from '@spectrum-web-components/button';
    
### Anatomy
#### Content
`<sp-button>` elements can be provided a visible label, a label and an icon, or just an icon.
An icon is provided by placing an icon element in the `icon` slot.
If the button is `icon-only`, a non-visible label can be provided via the `label` attribute on an `<sp-button>` or on an `<sp-icon*>` element child to appropriately fulfill the accessibility contract of the button.
Label only
    
    <sp-button variant="primary">Label only</sp-button>
Icon + label
    
    <sp-button variant="primary">
        <sp-icon-help slot="icon"></sp-icon-help>
        Icon + Label
    </sp-button>
SVG Icon + label
    
    <sp-button variant="primary">
        <svg
            slot="icon"
            viewBox="0 0 36 36"
            focusable="false"
            aria-hidden="true"
            role="img"
        >
            <path
                d="M16 36a4.407 4.407 0 0 0 4-4h-8a4.407 4.407 0 0 0 4 4zm9.143-24.615c0-3.437-3.206-4.891-7.143-5.268V3a1.079 1.079 0 0 0-1.143-1h-1.714A1.079 1.079 0 0 0 14 3v3.117c-3.937.377-7.143 1.831-7.143 5.268C6.857 26.8 2 26.111 2 28.154V30h28v-1.846C30 26 25.143 26.8 25.143 11.385z"
            ></path>
        </svg>
        SVG Icon + Label
    </sp-button>
Icon only
    
    <sp-button variant="primary" label="Icon only">
        <sp-icon-help slot="icon"></sp-icon-help>
    </sp-button>
### Options
#### Sizes
Small
    
    <sp-button size="s">Small</sp-button>
Medium
    
    <sp-button size="m">Medium</sp-button>
Large
    
    <sp-button size="l">Large</sp-button>
Extra Large
    
    <sp-button size="xl">Extra Large</sp-button>
#### Variants
There are many button variants to choose from in Spectrum. The `variant` attribute defaults to `accent`, but also accepts the following value: `accent`, `primary`, `secondary`, `negative`, `white`, and `black`. They display as follows:
Accent
    
    <sp-button-group style="min-width: max-content">
        <sp-button variant="accent">Label only</sp-button>
        <sp-button variant="accent">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button variant="accent" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
Primary
    
    <sp-button-group style="min-width: max-content">
        <sp-button variant="primary">Label only</sp-button>
        <sp-button variant="primary">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button variant="primary" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
Secondary
    
    <sp-button-group style="min-width: max-content">
        <sp-button variant="secondary">Label only</sp-button>
        <sp-button variant="secondary">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button variant="secondary" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
Negative
    
    <sp-button-group style="min-width: max-content">
        <sp-button variant="negative">Label only</sp-button>
        <sp-button variant="negative">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button variant="negative" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
Black
    
    <sp-button-group style="min-width: max-content">
        <sp-button static-color="black">Label only</sp-button>
        <sp-button static-color="black">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button static-color="black" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
White
    
    <sp-button-group style="min-width: max-content">
        <sp-button static-color="white">Label only</sp-button>
        <sp-button static-color="white">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button static-color="white" label="Icon only" icon-only>
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
#### Treatment
The `treatment` attribute accepts `fill` and `outline` as values, and defaults to `fill`. These display as follows:
Fill
    
    <sp-button-group style="min-width: max-content">
        <sp-button treatment="fill" variant="primary">Primary, Fill</sp-button>
        <sp-button treatment="fill" variant="secondary">Secondary, Fill</sp-button>
        <sp-button treatment="fill" variant="negative">Negative, Fill</sp-button>
    </sp-button-group>
Outline
    
    <sp-button-group style="min-width: max-content">
        <sp-button treatment="outline" variant="primary">
            Primary, Outline
        </sp-button>
        <sp-button treatment="outline" variant="secondary">
            Secondary, Outline
        </sp-button>
        <sp-button treatment="outline" variant="negative">
            Negative, Outline
        </sp-button>
    </sp-button-group>
Outline, black
    
    <sp-button-group
        style="background: var(--spectrum-seafoam-600); padding: 0.5em; min-width: max-content"
    >
        <sp-button treatment="outline" static-color="black">Label only</sp-button>
        <sp-button treatment="outline" static-color="black">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button
            treatment="outline"
            static-color="black"
            label="Icon only"
            icon-only
        >
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
Outline, white
    
    <sp-button-group
        style="background: var(--spectrum-seafoam-600); padding: 0.5em; min-width: max-content"
    >
        <sp-button treatment="outline" static-color="white">Label only</sp-button>
        <sp-button treatment="outline" static-color="white">
            <sp-icon-help slot="icon"></sp-icon-help>
            Icon + Label
        </sp-button>
        <sp-button
            treatment="outline"
            static-color="white"
            label="Icon only"
            icon-only
        >
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-button>
    </sp-button-group>
### States
In addition to the variant, `<sp-button>` elements support two different visual states, disabled and pending, which can be applied by adding the attribute `disabled` or `pending` respectively. All `<sp-button>` variants support these states.
#### Disabled
While disabled, `<sp-button>` elements will not respond to click events and will appear faded.
    
    <sp-button-group>
        <sp-button variant="primary">Normal</sp-button>
        <sp-button variant="primary" disabled>Disabled</sp-button>
    </sp-button-group>
#### Pending
While in pending state, `<sp-button>` elements will not respond to click events and will appear faded with an indeterminate `<sp-progress-circle>`. The `<sp-button>` element's label and icon will be hidden while in pending state.
Note: The pending state of the `<sp-button>` element is applied after a 1s delay to avoid flashing the pending state for quick actions. You can override the delay by adding custom css var `--pending-delay` to your css.
    
    <sp-button-group>
        <sp-button variant="primary">Normal</sp-button>
        <sp-button variant="primary" pending>Pending</sp-button>
    </sp-button-group>
### Behaviors
#### Handling events
Events handlers for clicks and other user actions can be registered on a `<sp-button>` as one would on a standard HTML `<button>` element.
    
    <sp-button onclick="spAlert(this, '<sp-button> clicked!')">Click me</sp-button>
In addition to handling events like a native `<button>` HTML element, one can also use a `<sp-button>` in place of the `<a>` HTML element by using the `href` and optional `target` attribute.
    
    <sp-button
        href="https://github.com/adobe/spectrum-web-components"
        target="_blank"
    >
        Click me
    </sp-button>
#### Autofocus
The `autofocus` attribute sets focus on the `<sp-button>` when the component mounts. This is useful for setting focus to a specific sp-button when a popover or dialog opens.
    
    <sp-button autofocus>Confirm</sp-button>
### Accessibility
#### Include a label
A button is required to have either a visible text label or a `label` attribute on either the `<sp-button>` itself or on an `<sp-icon*>` element child.
#### Don't override color
Do not use custom colors for buttons. The colors of different button variations have been designed to be consistent and accessible.
#### Don't mix href and non-href buttons in a set of buttons
A screen reader user will not encounter href buttons when navigating by buttons or form controls. While they can both be used in the same page problems could occur if mixing the types in close proximity to each other.
#### Use static black or static white to contrast with backgrounds and images
To ensure maximum contrast with the background, use static black for light backgrounds and images, and static white for dark backgrounds and images. Avoid placing static components on top of busy images with a lot of variance in contrast.
Static black on light background
    
    <div style="background-color: #ccffee; padding: 20px">
        <sp-button static="black">Click me</sp-button>
        <sp-button static="black" treatment="outline">Click me</sp-button>
    </div>
Static white on dark background
    
    <div style="background-color: #220033; padding: 20px">
        <sp-button static="white">Click me</sp-button>
        <sp-button static="white" treatment="outline">Click me</sp-button>
    </div>
#### Clearly state the action
Make sure that a button’s label clearly states the outcome of the action. Use the same word or phrase as found elsewhere in the experience.
## API
### Attributes and Properties
<table>
  <thead>
    <tr>
      <th>Property</th>
      <th>Attribute</th>
      <th>Type</th>
      <th>Default</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>`active`</td>
      <td>`active`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Disable this control. It will not receive focus or events</td>
    </tr>
    <tr>
      <td>`download`</td>
      <td>`download`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Causes the browser to treat the linked URL as a download.</td>
    </tr>
    <tr>
      <td>`href`</td>
      <td>`href`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The URL that the hyperlink points to.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.</td>
    </tr>
    <tr>
      <td>`noWrap`</td>
      <td>`no-wrap`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Disables text wrapping within the button component's label. Please note that this option is not a part of the design specification and should be used carefully, with consideration of this overflow behavior and the readability of the button's content.</td>
    </tr>
    <tr>
      <td>`pending`</td>
      <td>`pending`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`pendingLabel`</td>
      <td>`pending-label`</td>
      <td>`string`</td>
      <td>`'Pending'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>``</td>
      <td>Style this button to be less obvious</td>
    </tr>
    <tr>
      <td>`referrerpolicy`</td>
      <td>`referrerpolicy`</td>
      <td>`| 'no-referrer' | 'no-referrer-when-downgrade' | 'origin' | 'origin-when-cross-origin' | 'same-origin' | 'strict-origin' | 'strict-origin-when-cross-origin' | 'unsafe-url' | undefined`</td>
      <td>``</td>
      <td>How much of the referrer to send when following the link.</td>
    </tr>
    <tr>
      <td>`rel`</td>
      <td>`rel`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The relationship of the linked URL as space-separated link types.</td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'black' | 'white' | undefined`</td>
      <td>``</td>
      <td>The static color variant to use for this button.</td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
    </tr>
    <tr>
      <td>`target`</td>
      <td>`target`</td>
      <td>`'_blank' | '_parent' | '_self' | '_top' | undefined`</td>
      <td>``</td>
      <td>Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).</td>
    </tr>
    <tr>
      <td>`treatment`</td>
      <td>`treatment`</td>
      <td>`ButtonTreatments`</td>
      <td>`'fill'`</td>
      <td>The visual treatment to apply to this button.</td>
    </tr>
    <tr>
      <td>`type`</td>
      <td>`type`</td>
      <td>`'button' | 'submit' | 'reset'`</td>
      <td>`'button'`</td>
      <td>The default behavior of the button. Possible values are: `button` (default), `submit`, and `reset`.</td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`ButtonVariants`</td>
      <td>``</td>
      <td>The visual variant to apply to this button.</td>
    </tr>
  </tbody>
</table>
### Slots
<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>`default slot`</td>
      <td>text label of the Button</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>The icon to use for Button</td>
    </tr>
  </tbody>
</table>
