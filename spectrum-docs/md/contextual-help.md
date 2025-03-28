# sp-contextual-help
Overview API
## Description
A `sp-contextual-help` shows a user extra information about the state of either an adjacent component or an entire view. It explains a high-level topic about an experience and can point users to more information elsewhere.
### Usage
    
    yarn add @spectrum-web-components/contextual-help
    
Import the side effectful registration of `<sp-contextual-help>` via:
    
    import '@spectrum-web-components/contextual-help/sp-contextual-help.js';
    
When looking to leverage the `ContextualHelp` base class as a type and/or for extension purposes, do so via:
    
    import { ContextualHelp } from '@spectrum-web-components/contextual-help';
    
## Example
    
    <sp-contextual-help>
        <h2 slot="heading">Permission required</h2>
        Your admin must grant you permission before you can create a segment.
        <sp-link
            slot="link"
            href="https://opensource.adobe.com/spectrum-web-components/"
        >
            Request permission
        </sp-link>
    </sp-contextual-help>
## Help
Use `variant="help"` for helpful content: more detailed, in-depth guidance about a task, UI element, tool, or keyboard shortcuts. This may include an image, video, or link and should be helpful in tone.
    
    <sp-contextual-help variant="help">
        <h2 slot="heading">What is a segment?</h2>
        Segments identify who your visitors are, what devices and services they use,
        where they navigate from, and much more.
        <sp-link
            slot="link"
            href="https://opensource.adobe.com/spectrum-web-components/"
        >
            Learn more about segments
        </sp-link>
    </sp-contextual-help>
## Placement
By default a `sp-contextual-help` will render its popover at the `bottom-start` position. This can be customized using the `placement` attribute and supports all the placement options an `overlay-trigger` component supports.
    
    <sp-contextual-help placement="top-start">
        <h2 slot="heading">Permission required</h2>
        Your admin must grant you permission before you can create a segment.
        <sp-link
            slot="link"
            href="https://opensource.adobe.com/spectrum-web-components/"
        >
            Request permission
        </sp-link>
    </sp-contextual-help>
## Custom max width
Contextual help allows for a custom maximum width to be set using the `--mod-spectrum-contextual-help-popover-maximum-width` custom property.
Note: Maximum width should not be less than 100px.
    
    <sp-contextual-help
        style="--mod-spectrum-contextual-help-popover-maximum-width: 200px;"
    >
        <h2 slot="heading">Custom max width</h2>
        This is a test of the contextual help component with a custom max width of
        200px.
    </sp-contextual-help>
## Events
`sp-contextual-help` does not fire any events of its own. You can listen, however, for the `sp-open` and `sp-closed` events which are fired when the popover opens or closes.
## Accessibility
Given that the trigger is an icon-only `sp-action-button`, it is important to provide an accessible name for it, so that it can be properly announced by screen readers. By default, the `sp-contextual-help` uses an `aria-label` property with either "Informations" or "Help" as values, depending on the component's `variant`. You can customize this using the `label` attribute.
When providing headings using the `heading` slot, make sure to provide actual heading elements such as `h1`, `h2`, `h3` ... or use the `role="heading"` attribute.
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
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Provides an accessible name for the action button trigger.</td>
    </tr>
    <tr>
      <td>`offset`</td>
      <td>`offset`</td>
      <td>`number | [number, number]`</td>
      <td>`0`</td>
      <td>The `offset` property accepts either a single number, to define the offset of the Popover along the main axis from the action button, or 2-tuple, to define the offset along the main axis and the cross axis.</td>
    </tr>
    <tr>
      <td>`placement`</td>
      <td>`placement`</td>
      <td>`"top" | "top-start" | "top-end" | "right" | "right-start" | "right-end" | "bottom" | "bottom-start" | "bottom-end" | "left" | "left-start" | "left-end"`</td>
      <td>`'bottom-start'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`'info' | 'help'`</td>
      <td>`'info'`</td>
      <td>The `variant` property applies specific styling on the action button trigger.</td>
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
      <td>`Text`</td>
      <td>content to display in the popover</td>
    </tr>
    <tr>
      <td>`heading`</td>
      <td>content to display as the heading of the popover</td>
    </tr>
    <tr>
      <td>`link`</td>
      <td>link to additional informations</td>
    </tr>
  </tbody>
</table>
