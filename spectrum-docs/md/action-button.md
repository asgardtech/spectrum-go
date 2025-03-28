# sp-action-button
Overview API
## Overview
An `<sp-action-button>` represents an action a user can take.
### Usage
    
    yarn add @spectrum-web-components/action-button
    
Import the side effectful registration of `<sp-action-button>` via:
    
    import '@spectrum-web-components/action-button/sp-action-button.js';
    
When looking to leverage the `ActionButton` base class as a type and/or for extension purposes, do so via:
    
    import { ActionButton } from '@spectrum-web-components/action-button';
    
### Options
#### Sizes
Extra Small
    
    <sp-action-group size="xs">
        <sp-action-button>Edit</sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
            Edit
        </sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button hold-affordance>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
    </sp-action-group>
Small
    
    <sp-action-group size="s">
        <sp-action-button>Edit</sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
            Edit
        </sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button hold-affordance>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
    </sp-action-group>
Medium
    
    <sp-action-group size="m">
        <sp-action-button>Edit</sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
            Edit
        </sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button hold-affordance>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
    </sp-action-group>
Large
    
    <sp-action-group size="l">
        <sp-action-button>Edit</sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
            Edit
        </sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button hold-affordance>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
    </sp-action-group>
Extra Large
    
    <sp-action-group size="xl">
        <sp-action-button>Edit</sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
            Edit
        </sp-action-button>
        <sp-action-button>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button hold-affordance>
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
    </sp-action-group>
#### Variants
The `<sp-action-button>` can be customized with either or both of the `emphasized` and `quiet` attributes. These will pair with either or both of the state attributes (`selected` and `disabled`) to decide the final visual delivery of the `<sp-action-button>`. Content addressed to the `icon` slot can also be provided and will be positioned just before the rest of the the supplied button content.
Default
    
    <div
        style="display: grid; grid-template-columns: repeat(auto-fill, minmax(210px, 1fr)); gap: 2em;"
    >
        <div>
            <sp-field-label for="standard">Default</sp-field-label>
            <sp-action-group id="standard">
                <sp-action-button>Edit</sp-action-button>
                <sp-action-button>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-selected">Selected</sp-field-label>
            <sp-action-group id="standard-selected">
                <sp-action-button selected>Edit</sp-action-button>
                <sp-action-button selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled">Disabled</sp-field-label>
            <sp-action-group id="standard-disabled">
                <sp-action-button disabled>Edit</sp-action-button>
                <sp-action-button disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button disabled hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled-selected">
                Disabled + Selected
            </sp-field-label>
            <sp-action-group id="standard-disabled-selected">
                <sp-action-button disabled selected>Edit</sp-action-button>
                <sp-action-button disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button disabled selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    </div>
Quiet
    
    <div
        style="display: grid; grid-template-columns: repeat(auto-fill, minmax(210px, 1fr)); gap: 2em;"
    >
        <div>
            <sp-field-label for="standard">Default</sp-field-label>
            <sp-action-group quiet id="standard">
                <sp-action-button quiet>Edit</sp-action-button>
                <sp-action-button quiet>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button quiet>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button quiet hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-selected">Selected</sp-field-label>
            <sp-action-group quiet id="standard-selected">
                <sp-action-button quiet selected>Edit</sp-action-button>
                <sp-action-button quiet selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button quiet selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button quiet selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled">Disabled</sp-field-label>
            <sp-action-group quiet id="standard-disabled">
                <sp-action-button quiet disabled>Edit</sp-action-button>
                <sp-action-button quiet disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button quiet disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button quiet disabled hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled-selected">
                Disabled + Selected
            </sp-field-label>
            <sp-action-group quiet id="standard-disabled-selected">
                <sp-action-button quiet disabled selected>Edit</sp-action-button>
                <sp-action-button quiet disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button quiet disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button quiet disabled selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    </div>
Emphasized
    
    <div
        style="display: grid; grid-template-columns: repeat(auto-fill, minmax(210px, 1fr)); gap: 2em;"
    >
        <div>
            <sp-field-label for="standard">Default</sp-field-label>
            <sp-action-group emphasized id="standard">
                <sp-action-button emphasized>Edit</sp-action-button>
                <sp-action-button emphasized>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-selected">Selected</sp-field-label>
            <sp-action-group emphasized id="standard-selected">
                <sp-action-button emphasized selected>Edit</sp-action-button>
                <sp-action-button emphasized selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled">Disabled</sp-field-label>
            <sp-action-group emphasized id="standard-disabled">
                <sp-action-button emphasized disabled>Edit</sp-action-button>
                <sp-action-button emphasized disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized disabled hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled-selected">
                Disabled + Selected
            </sp-field-label>
            <sp-action-group emphasized id="standard-disabled-selected">
                <sp-action-button emphasized disabled selected>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized disabled selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    </div>
Emphasized + quiet
    
    <div
        style="display: grid; grid-template-columns: repeat(auto-fill, minmax(210px, 1fr)); gap: 2em;"
    >
        <div>
            <sp-field-label for="standard">Default</sp-field-label>
            <sp-action-group emphasized quiet id="standard">
                <sp-action-button emphasized quiet>Edit</sp-action-button>
                <sp-action-button emphasized quiet>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized quiet>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized quiet hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-selected">Selected</sp-field-label>
            <sp-action-group emphasized quiet id="standard-selected">
                <sp-action-button emphasized quiet selected>Edit</sp-action-button>
                <sp-action-button emphasized quiet selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized quiet selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized quiet selected hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled">Disabled</sp-field-label>
            <sp-action-group emphasized quiet id="standard-disabled">
                <sp-action-button emphasized quiet disabled>Edit</sp-action-button>
                <sp-action-button emphasized quiet disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized quiet disabled>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button emphasized quiet disabled hold-affordance>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    
        <div>
            <sp-field-label for="standard-disabled-selected">
                Disabled + Selected
            </sp-field-label>
            <sp-action-group emphasized quiet id="standard-disabled-selected">
                <sp-action-button emphasized quiet disabled selected>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized quiet disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                    Edit
                </sp-action-button>
                <sp-action-button emphasized quiet disabled selected>
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
                <sp-action-button
                    emphasized
                    quiet
                    disabled
                    selected
                    hold-affordance
                >
                    <sp-icon-edit slot="icon"></sp-icon-edit>
                </sp-action-button>
            </sp-action-group>
        </div>
    </div>
### Behaviors
#### Action button with hold affordance
The use of the `hold-affordance` attribute signifies that the `<sp-action-button>` in question will be delivered with a visual affordance outlining that special interaction with the button will dispatch a `longpress` event. Via a pointer input, this even will be dispatched when 300ms has passed after a `pointerdown` event without the presence of a `pointerup` or `pointercancel` event. Via the keyboard, an event with a code of `Space` or or `ArrowDown` while `altKey === true` will dispatch the event.
    
    <div
        style="display: grid; grid-template-columns: repeat(auto-fill, minmax(200px, 1fr)); gap: 2em;"
    >
        <overlay-trigger placement="bottom-start">
            <sp-action-button label="Edit" hold-affordance slot="trigger">
                <sp-icon-edit slot="icon"></sp-icon-edit>
            </sp-action-button>
            <sp-popover slot="longpress-content" dialog tip>
                <p style="color: var(--spectrum-neutral-content-color-default);">
                    This content is triggered by the "longpress" interaction.
                </p>
            </sp-popover>
        </overlay-trigger>
    
        <overlay-trigger placement="top">
            <sp-action-button hold-affordance quiet slot="trigger">
                Show Longpress Content
            </sp-action-button>
            <sp-popover slot="longpress-content" dialog tip>
                <p style="color: var(--spectrum-neutral-content-color-default);">
                    This content is triggered by the "longpress" interaction.
                </p>
            </sp-popover>
        </overlay-trigger>
    
        <overlay-trigger placement="top-end">
            <sp-action-button hold-affordance selected slot="trigger">
                <sp-icon-edit slot="icon"></sp-icon-edit>
                Extended Content with Longpress
            </sp-action-button>
            <sp-popover slot="longpress-content" dialog tip>
                <p style="color: var(--spectrum-neutral-content-color-default);">
                    This content is triggered by the "longpress" interaction.
                </p>
            </sp-popover>
        </overlay-trigger>
    </div>
#### Toggles
With the application of the `toggles` attribute, the button will self manage its `selected` property on `click`. When this value is updated, a cancellable `change` event will be dispatched to inform the parent application.
Default
    
    <sp-action-button toggles id="toggles-default">Toggle button</sp-action-button>
    <sp-action-button toggles selected id="toggles-default">
        Toggle button
    </sp-action-button>
Quiet
    
    <sp-action-button toggles quiet id="toggles-quiet">
        Toggle button
    </sp-action-button>
    <sp-action-button toggles quiet selected id="toggles-quiet">
        Toggle button
    </sp-action-button>
Emphasized
    
    <sp-action-button toggles emphasized id="toggles-emphasized">
        Toggle button
    </sp-action-button>
    <sp-action-button toggles emphasized selected id="toggles-emphasized">
        Toggle button
    </sp-action-button>
Emphasized + Quiet
    
    <sp-action-button toggles emphasized quiet id="toggles-emphasized-quiet">
        Toggle button
    </sp-action-button>
    <sp-action-button
        toggles
        emphasized
        quiet
        selected
        id="toggles-emphasized-quiet"
    >
        Toggle button
    </sp-action-button>
#### Handling events
Events handlers for clicks and other user actions can be registered on a `<sp-action-button>` as on a standard HTML `<button>` element.
    
    <sp-button onclick="spAlert(this, '<sp-action-button> clicked!')">
        Click me
    </sp-button>
In addition to handling events like a native `<button>` HTML element, one can also use a `<sp-action-button>` in place of the `<a>` HTML element by using the `href` and optional `target` attribute.
    
    <sp-action-button
        href="https://github.com/adobe/spectrum-web-components"
        target="_blank"
    >
        Click me
    </sp-action-button>
### Accessibility
#### Include a label
A button is required to have either a visible text label or a `label` attribute on either the `<sp-button>` itself, or on an `<sp-icon*>` element child.
#### Don't override color
Do not use custom colors for buttons. The colors of different button variations have been designed to be consistent and accessible.
#### Use static black or static white to contrast with backgrounds and images
To ensure maximum contrast with the background, use static black for light backgrounds and images, and static white for dark backgrounds and images. Avoid placing static components on top of busy images with a lot of variance in contrast.
Static black on light background
    
    <div style="background-color: #ccffee; padding: 20px">
        <sp-action-button static="black">Click me</sp-action-button>
        <sp-action-button static="black" selected>Click me</sp-action-button>
    </div>
Static white on dark background
    
    <div style="background-color: #220033; padding: 20px">
        <sp-action-button static="white">Click me</sp-action-button>
        <sp-action-button static="white" selected>Click me</sp-action-button>
    </div>
#### Clearly state the action
Make sure that an action button’s label clearly states the outcome of the action. Use the same word or phrase as found elsewhere in the experience.
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
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`holdAffordance`</td>
      <td>`hold-affordance`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>`role`</td>
      <td>`role`</td>
      <td>`string`</td>
      <td>`'button'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether an Action Button with `role='button'` should also be `aria-pressed='true'`</td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | 'black' | undefined`</td>
      <td>``</td>
      <td>The static color variant to use for the action button.</td>
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
      <td>`toggles`</td>
      <td>`toggles`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether to automatically manage the `selected` attribute on interaction and whether `aria-pressed="false"` should be used when `selected === false`</td>
    </tr>
    <tr>
      <td>`type`</td>
      <td>`type`</td>
      <td>`'button' | 'submit' | 'reset'`</td>
      <td>`'button'`</td>
      <td>The default behavior of the button. Possible values are: `button` (default), `submit`, and `reset`.</td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string`</td>
      <td>``</td>
      <td></td>
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
      <td>text label of the Action Button</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>The icon to use for Action Button</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`change`</td>
      <td>`Event`</td>
      <td>`Announces a change in the `selected` property of an action button`</td>
    </tr>
    <tr>
      <td>`longpress`</td>
      <td>`CustomEvent`</td>
      <td>`Synthesizes a "longpress" interaction that signifies a `pointerdown` event that is >=300ms or a keyboard event where code is `Space` or code is `ArrowDown` while `altKey===true`.`</td>
    </tr>
  </tbody>
</table>
