# sp-accordion-item
Overview API
## Description
The `<sp-accordion-item>` element represents a single item in an `<sp-accordion>` parent element. Its `label` attribute and default slot content make up the "headline" and "body" of the toggleable content item.
### Usage
    
    yarn add @spectrum-web-components/accordion
    
Import the side effectful registration of `<sp-accordion-item>` via:
    
    import '@spectrum-web-components/accordion/sp-accordion-item.js';
    
When looking to leverage the `AccordionItem` base class as a type and/or for extension purposes, do so via:
    
    import { AccordionItem } from '@spectrum-web-components/accordion';
    
## Example
    
    <sp-accordion>
        <sp-accordion-item label="Heading 1">
            <div>The content can be toggled visible.</div>
        </sp-accordion-item>
    </sp-accordion>
### Opened
    
    <sp-accordion allow-multiple>
        <sp-accordion-item open label="Heading 2">
            <div>This content is visible by default.</div>
        </sp-accordion-item>
    </sp-accordion>
### Disabled
    
    <sp-accordion allow-multiple>
        <sp-accordion-item disabled label="Heading 2">
            <div>You can not toggle this content visible.</div>
        </sp-accordion-item>
    </sp-accordion>
## Events
An `<sp-accordion-item>` element will dispatch `sp-accordion-item-toggle` events when it is opened or closed. By default, these events are dispatched to allow the parent `<sp-accordion>` to manage which of its item children can shot their children at any one time, consumers can also listen for this event and leverage `event.target.open` to ascertain the current state of the item dispatching item.
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
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Disable this control. It will not receive focus or events</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
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
      <td>The content of the item that is hidden when the item is not open</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`sp-accordion-item-toggle`</td>
      <td>`CustomEvent`</td>
      <td>`Announce that an accordion item has been toggled while allowing the event to be cancelled.`</td>
    </tr>
  </tbody>
</table>
