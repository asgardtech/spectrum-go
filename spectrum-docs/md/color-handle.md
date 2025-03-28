# sp-color-handle
Overview API
## Description
The `<sp-color-handle>` is used to select a colour on an `<sp-color-area>`, `<sp-color-slider>`, or `<sp-color-wheel>`. It functions similarly to the handle on an `<sp-slider>`.
### Usage
    
    yarn add @spectrum-web-components/color-handle
    
Import the side effectful registration of `<sp-color-handle>` via:
    
    import '@spectrum-web-components/color-handle/sp-color-handle.js';
    
When looking to leverage the `ColorHandle` base class as a type and/or for extension purposes, do so via:
    
    import { ColorHandle } from '@spectrum-web-components/color-handle';
    
## Standard
    
    <sp-color-handle></sp-color-handle>
## Disabled
    
    <sp-color-handle disabled></sp-color-handle>
## Open
When the `<sp-color-handle>` uses the `open` property, the `<sp-color-loupe>` component can be used above the handle to show the selected color that would otherwise be covered by a mouse, stylus, or finger on the down/touch state. This can be customized to appear only on finger-input, or always appear regardless of input type.
    
    <div style="height: 72px"></div>
    <sp-color-handle open></sp-color-handle>
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
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`focused`</td>
      <td>`focused`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
