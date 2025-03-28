# sp-tray
Overview API
## Overview
`<sp-tray>` elements are typically used to portray information on mobile device or smaller screens.
### Usage
    
    yarn add @spectrum-web-components/tray
    
Import the side effectful registration of `<sp-tray>` via:
    
    import '@spectrum-web-components/tray/sp-tray.js';
    
When looking to leverage the `Tray` base class as a type and/or for extension purposes, do so via:
    
    import { Tray } from '@spectrum-web-components/tray';
    
### Anatomy
A tray has a single default `slot`.
Dialog
    
    <overlay-trigger type="modal">
        <sp-button slot="trigger" variant="secondary">Toggle tray</sp-button>
        <sp-tray slot="click-content">
            <sp-dialog size="s" dismissable>
                <h2 slot="heading">New Messages</h2>
                You have 5 new messages.
            </sp-dialog>
        </sp-tray>
    </overlay-trigger>
Menu
    
    <overlay-trigger type="modal">
        <sp-button slot="trigger" variant="secondary">Toggle menu</sp-button>
        <sp-tray slot="click-content">
            <sp-menu style="width: 100%">
                <sp-menu-item selected>Deselect</sp-menu-item>
                <sp-menu-item>Select Inverse</sp-menu-item>
                <sp-menu-item focused>Feather...</sp-menu-item>
                <sp-menu-item>Select and Mask...</sp-menu-item>
                <sp-menu-divider></sp-menu-divider>
                <sp-menu-item>Save Selection</sp-menu-item>
                <sp-menu-item disabled>Make Work Path</sp-menu-item>
            </sp-menu>
        </sp-tray>
    </overlay-trigger>
### Accessibility
`<sp-tray>` presents a page blocking experience and should be opened with the `Overlay` API using the `modal` interaction to ensure that the content appropriately manages the presence of other content in the tab order of the page and the availability of that content for a screen reader.
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
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
      <td>content to display within the Tray</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`close`</td>
      <td>`Event`</td>
      <td>`Announces that the Tray has been closed.`</td>
    </tr>
  </tbody>
</table>
