# sp-color-area
Overview API
## Description
An `<sp-color-area>` allows users to visually select two properties of a color simultaneously. It's commonly used together with a color slider or color wheel.
### Usage
    
    yarn add @spectrum-web-components/color-area
    
Import the side effectful registration of `<sp-color-area>` via:
    
    import '@spectrum-web-components/color-area/sp-color-area.js';
    
When looking to leverage the `ColorArea` base class as a type and/or for extension purposes, do so via:
    
    import { ColorArea } from '@spectrum-web-components/color-area';
    
## Color Formatting
When using the color elements, use `el.color` to access the `color` property, which should manage itself in the colour format supplied. For example, If you supply a color in `rgb()` format, `el.color` should return the color in `rgb()` format, as well. In ColorArea, colours are formatted as hex values.
The current color formats supported are as follows:
  -  Hex3, Hex4, Hex6, Hex8
  -  HSV, HSVA
  -  HSL, HSLA
  -  RGB, RGBA
  -  Strings (eg "red", "blue")

## Standard
    
    <sp-color-area></sp-color-area>
## Variants
### Disabled
An `<sp-color-area>` in a disabled state shows that an input exists, but is not available in that circumstance. This can be used to maintain layout continuity and communicate that the area may become available later.
    
    <sp-color-area disabled></sp-color-area>
### Sized
An `<sp-color-area>`’s height and width can be customized appropriately for its context.
    
    <sp-color-area
        style="
            width: 72px; 
            height: 72px"
    ></sp-color-area>
## Labels
An `<sp-color-area>` renders accessible labels for each axis: *"saturation"* and *"luminosity"*. Specify `label-x` and `label-y` attributes to override these defaults.
    
    <sp-color-area
        label-x="Color intensity"
        label-y="Color brightness"
    ></sp-color-area>
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
      <td>`ColorTypes`</td>
      <td>``</td>
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
      <td>`hue`</td>
      <td>`hue`</td>
      <td>`number`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`labelX`</td>
      <td>`label-x`</td>
      <td>`string`</td>
      <td>`'saturation'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`labelY`</td>
      <td>`label-y`</td>
      <td>`string`</td>
      <td>`'luminosity'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`step`</td>
      <td>`step`</td>
      <td>`number`</td>
      <td>`0.01`</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`ColorTypes`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`x`</td>
      <td>`x`</td>
      <td>`number`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`y`</td>
      <td>`y`</td>
      <td>`number`</td>
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
      <td>`gradient`</td>
      <td>a custom gradient visually outlining the available color values</td>
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
      <td>`An alteration to the value of the Color Area has been committed by the user.`</td>
    </tr>
    <tr>
      <td>`input`</td>
      <td>`Event`</td>
      <td>`The value of the Color Area has changed.`</td>
    </tr>
  </tbody>
</table>
