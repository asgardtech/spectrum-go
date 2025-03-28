# sp-infield-button
Overview API
## Description
When composing complex form fields, an `<sp-infield-button>` can visually associate button functionality with other form fields to delivery enhanced capabilities to your visitors.
### Usage
    
    yarn add @spectrum-web-components/infield-button
    
Import the side effectful registration of `<sp-infield-button>` via:
    
    import '@spectrum-web-components/infield-button/sp-infield-button.js';
    
When looking to leverage the `InfieldButton` base class as a type and/or for extension purposes, do so via:
    
    import { InfieldButton } from '@spectrum-web-components/infield-button';
    
## Sizes
Small
    
    <sp-infield-button label="Add" size="s">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
Medium
    
    <sp-infield-button label="Add" size="m">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
Large
    
    <sp-infield-button label="Add" size="l">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
Extra Large
    
    <sp-infield-button label="Add" size="xl">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
## Inline buttons
Use the `inline` attribute to describe whether the `<sp-infield-button>` should be visually at the `start` or `end` of the field is associated to:
### inline="start"
    
    <sp-infield-button inline="start" label="Add">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
### inline="end"
    
    <sp-infield-button inline="end" label="Add">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
## Stacked buttons
The `block` attribute can be used to create a vertial stack of buttons. You can place buttons visually on the stack with the `start` or `end` values:
    
    <sp-infield-button block="start" label="Increment">
        <sp-icon-add size="xxs"></sp-icon-add>
    </sp-infield-button>
    <sp-infield-button block="end" label="Decrement">
        <sp-icon-remove size="xxs"></sp-icon-remove>
    </sp-infield-button>
## Disabled
An `<sp-infield-button>` with the `disabled` attribute will become non-interactive and dimmed:
    
    <sp-infield-button disabled inline="start" label="Add">
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
## Quiet
An `<sp-infield-button>` with the `quiet` attribute will feature a diminished visual presence:
    
    <sp-infield-button inline="start" label="Add" quiet>
        <sp-icon-add></sp-icon-add>
    </sp-infield-button>
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
      <td>`type`</td>
      <td>`type`</td>
      <td>`'button' | 'submit' | 'reset'`</td>
      <td>`'button'`</td>
      <td>The default behavior of the button. Possible values are: `button` (default), `submit`, and `reset`.</td>
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
      <td>text content to be displayed in the Button element</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>icon element(s) to display at the start of the button</td>
    </tr>
  </tbody>
</table>
