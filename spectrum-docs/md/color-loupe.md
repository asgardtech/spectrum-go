# sp-color-loupe
Overview API
## Description
An `<sp-color-loupe>` shows the output color that would otherwise be covered by a cursor, stylus, or finger during color selection.
### Usage
    
    yarn add @spectrum-web-components/color-loupe
    
Import the side effectful registration of `<sp-color-loupe>` via:
    
    import '@spectrum-web-components/color-loupe/sp-color-loupe.js';
    
When looking to leverage the `ColorLoupe` base class as a type and/or for extension purposes, do so via:
    
    import { ColorLoupe } from '@spectrum-web-components/color-loupe';
    
## Example
    
    <div style="padding: 100px 0 0;">
        <div style="position:relative">
            <sp-color-loupe open="" dir="ltr"></sp-color-loupe>
        </div>
    </div>
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
      <td>`color`</td>
      <td>`color`</td>
      <td>`string`</td>
      <td>`'rgba(255, 0, 0, 0.5)'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
  </tbody>
</table>
