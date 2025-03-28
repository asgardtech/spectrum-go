# sp-color-slider
Overview API
## Description
An `<sp-color-slider>` lets users visually change an individual channel of a color. The background of the `<sp-color-slider>` is a visual representation of the range of values a user can select from. This can represent color properties such as hues, color channel values (such as RGB or CMYK levels), or opacity. Currently, the slider only supports leveraging the `hue` property.
### Usage
    
    yarn add @spectrum-web-components/color-slider
    
Import the side effectful registration of `<sp-color-slider>` via:
    
    import '@spectrum-web-components/color-slider/sp-color-slider.js';
    
When looking to leverage the `ColorSlider` base class as a type and/or for extension purposes, do so via:
    
    import { ColorSlider } from '@spectrum-web-components/color-slider';
    
## Color Formatting
When using the color elements, use `el.color` to access the `color` property, which should manage itself in the colour format supplied. If you supply a color in `rgb()` format, `el.color` should return the color in `rgb()` format, as well.
The current color formats supported are as follows:
  -  Hex3, Hex4, Hex6, Hex8
  -  HSV, HSVA
  -  HSL, HSLA
  -  RGB, RGBA
  -  Strings (eg "red", "blue")

**Please note for the following formats: HSV, HSVA, HSL, HSLA** When using the HSL or HSV formats, and a color's value (in HSV) is set to 0, or its luminosity (in HSL) is set to 0 or 1, the hue and saturation values may not be preserved by the element's `color` property. This is detailed in the colorjs documentation. Seperately, the element's `value` property is directly managed by the hue as represented in the interface.
## Default
    
    <sp-color-slider></sp-color-slider>
### Vertical
    
    <sp-color-slider vertical></sp-color-slider>
### Disabled
    
    <sp-color-slider disabled></sp-color-slider>
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
      <td>`An alteration to the value of the Color Slider has been committed by the user.`</td>
    </tr>
    <tr>
      <td>`input`</td>
      <td>`Event`</td>
      <td>`The value of the Color Slider has changed.`</td>
    </tr>
  </tbody>
</table>
