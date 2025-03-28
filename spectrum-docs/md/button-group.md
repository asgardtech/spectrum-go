# sp-button-group
Overview API
## Overview
`sp-button-group` delivers a set of buttons in horizontal or vertical orientation while ensuring the appropriate spacing between those buttons.
### Usage
    
    yarn add @spectrum-web-components/button-group
    
Import the side effectful registration of `<sp-button-group>` via:
    
    import '@spectrum-web-components/button-group/sp-button-group.js';
    
When looking to leverage the `ButtonGroup` base class as a type and/or for extension purposes, do so via:
    
    import { ButtonGroup } from '@spectrum-web-components/button-group';
    
### Options
A button group can be either horizontal or vertical in its orientation. By default, a button group is horizontal. Use vertical option when horizontal space is limited.
#### Horizontal
    
    <sp-button-group>
        <sp-button>Button 1</sp-button>
        <sp-button>Longer Button 2</sp-button>
        <sp-button>Short 3</sp-button>
    </sp-button-group>
#### Vertical
    
    <sp-button-group vertical>
        <sp-button>Button 1</sp-button>
        <sp-button>Longer Button 2</sp-button>
        <sp-button>Short 3</sp-button>
    </sp-button-group>
### Accessibility
Review the guidelines for the button children.
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
      <td>`vertical`</td>
      <td>`vertical`</td>
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
      <td>the sp-button elements that make up the group</td>
    </tr>
  </tbody>
</table>
