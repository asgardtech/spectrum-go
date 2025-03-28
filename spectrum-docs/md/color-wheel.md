# sp-color-wheel
Overview API
## Description
An `<sp-color-wheel>` lets users visually change an individual channel of a color on a circular track.
### Usage
    
    yarn add @spectrum-web-components/color-wheel
    
Import the side effectful registration of `<sp-color-wheel>` via:
    
    import '@spectrum-web-components/color-wheel/sp-color-wheel.js';
    
When looking to leverage the `ColorWheel` base class as a type and/or for extension purposes, do so via:
    
    import { ColorWheel } from '@spectrum-web-components/color-wheel';
    
## Color Formatting
When using the color elements, use `el.color` to access the `color` property, which should manage itself in the color format supplied. If you supply a color in `rgb()` format, `el.color` should return the color in `rgb()` format, as well.
The current color formats supported are as follows:
  -  Hex3, Hex4, Hex6, Hex8
  -  HSV, HSVA
  -  HSL, HSLA
  -  RGB, RGBA
  -  Strings (eg "red", "blue")

**Please note for the following formats: HSV, HSVA, HSL, HSLA** When using the HSL or HSV formats, and a color's value (in HSV) is set to 0, or its luminosity (in HSL) is set to 0 or 1, the hue and saturation values may not be preserved by the element's `color` property. This is detailed in the colorjs documentation. Seperately, the element's `value` property is directly managed by the hue as represented in the interface.
## Example
    
    <sp-color-wheel></sp-color-wheel>
### Disabled
A color wheel in a disabled state shows that an input exists, but is not available in that circumstance. This can be used to maintain layout continuity and communicate that the wheel may become available later.
    
    <sp-color-wheel disabled></sp-color-wheel>
## Variants
### Sized
An `<sp-color-wheel>`’s size can be customized appropriately for its context. By default, the size is size-2400 (192 px on desktop, 240 px on mobile).
    
    <sp-color-wheel style="width: 300px; height: 300px;"></sp-color-wheel>
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
      <td>Disable this control. It will not receive focus or events</td>
    </tr>
    <tr>
      <td>`focused`</td>
      <td>`focused`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`'hue'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`step`</td>
      <td>`step`</td>
      <td>`number`</td>
      <td>`1`</td>
      <td></td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
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
      <td>`An alteration to the value of the Color Wheel has been committed by the user.`</td>
    </tr>
    <tr>
      <td>`input`</td>
      <td>`Event`</td>
      <td>`The value of the Color Wheel has changed.`</td>
    </tr>
  </tbody>
</table>
