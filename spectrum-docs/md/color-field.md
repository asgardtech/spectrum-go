# sp-color-field
Overview API
## Description
`<sp-color-field>` elements are textfields that allow users to input custom color values. Color formats supported are `HEX, RGB, HSL, HSV, and shorthand HEX`
## Usage
    
    yarn add @spectrum-web-components/color-field
    
Import the side effectful registration of `<sp-color-field>` via:
    
    import '@spectrum-web-components/color-field/sp-color-field.js';
    
When looking to leverage the `ColorField` base class as a type and/or for extension purposes, do so via:
    
    import { ColorField } from '@spectrum-web-components/color-field';
    
## Sizes
Small
    
    <sp-color-field size="s" value="#ffff00"></sp-color-field>
Medium
    
    <sp-color-field size="m" value="#ffff00"></sp-color-field>
Large
    
    <sp-color-field size="l" value="#ffff00"></sp-color-field>
Xtra Large
    
    <sp-color-field size="xl" value="#ffff00"></sp-color-field>
## View Color
When `view-color` is true, the color handle will be rendered. This is useful for development and debugging purposes.
    
    <sp-color-field view-color value="#f00"></sp-color-field>
## Read Only
A readonly color field
    
    <sp-color-field readonly value="#ffff00"></sp-color-field>
## Quiet
A Quiet color field
    
    <sp-color-field quiet value="#e6e600"></sp-color-field>
## Invalid Input
If the input value is not a valid color, `<sp-color-field>` will not accept it.
    
    <sp-color-field value="not a color"></sp-color-field>
## Valid Input
If the input value is a valid color, the `<sp-color-field>` will accept it and the color handle will be updated to reflect the new color.
`<sp-color-field>` component accepts color values in various formats: `HEX, RGB, HSL, HSV, and shorthand HEX`
  -  **HEX** : A hexadecimal color is specified with: `#RRGGBB`. `RR` (red), `GG` (green) and `BB` (blue) are hexadecimal integers between `00` and `FF` specifying the intensity of the color.

    <sp-color-field value="#ff0000"></sp-color-field>
  -  **Shorthand HEX** : Shorthand hexadecimal color values are also supported. `#RGB` is a shorthand for `#RRGGBB`. In the shorthand form, `R` (red), `G` (green), and `B` (blue) are hexadecimal characters between `0` and `F`. Each character is repeated to create the full 6-digit color code. For example, `#123` would expand to `#112233`.

    <sp-color-field view-color value="#f00"></sp-color-field>
  -  **RGB** : An RGB color value is specified with: rgb(red, green, blue). Each parameter defines the intensity of the color with a value between 0 and 255.

    <sp-color-field view-color value="rgb(0,2555,0)"></sp-color-field>
  -  **RGBA** : An RGBA color value is specified with: `rgba(red, green, blue, alpha)`. The `alpha` parameter is a number between 0.0 (fully transparent) and 1.0 (fully opaque).

    <sp-color-field view-color value="rgba(0,255,0,0.3)"></sp-color-field>
  -  **HSL** : An HSL color value is specified with: hsl(hue, saturation, lightness). Hue is a degree on the color wheel from 0 to 360. 0 is red, 120 is green, and 240 is blue. Saturation and lightness are percentages.

    <sp-color-field view-color value="hsl(234, 70%, 50%)"></sp-color-field>
  -  **HSV** : An HSV color value is specified with: hsv(hue, saturation, value). Hue is a degree on the color wheel from 0 to 360. 0 is red, 120 is green, and 240 is blue. Saturation and value are percentages.

    <sp-color-field view-color value="hsv(0, 70%, 50%)"></sp-color-field>
## Events
The sp-color-field component fires a change event when the color value is changed. You can listen for this event to react to changes in the color value.
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
      <td>`autocomplete`</td>
      <td>`autocomplete`</td>
      <td>`| 'list' | 'none' | HTMLInputElement['autocomplete'] | HTMLTextAreaElement['autocomplete'] | undefined`</td>
      <td>``</td>
      <td>What form of assistance should be provided when attempting to supply a value to the form control</td>
    </tr>
    <tr>
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Disable this control. It will not receive focus or events</td>
    </tr>
    <tr>
      <td>`grows`</td>
      <td>`grows`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether a form control delivered with the `multiline` attribute will change size vertically to accomodate longer input</td>
    </tr>
    <tr>
      <td>`invalid`</td>
      <td>`invalid`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the `value` held by the form control is invalid.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td>A string applied via `aria-label` to the form control when a user visible label is not provided.</td>
    </tr>
    <tr>
      <td>`maxlength`</td>
      <td>`maxlength`</td>
      <td>`number`</td>
      <td>`-1`</td>
      <td>Defines the maximum string length that the user can enter</td>
    </tr>
    <tr>
      <td>`minlength`</td>
      <td>`minlength`</td>
      <td>`number`</td>
      <td>`-1`</td>
      <td>Defines the minimum string length that the user can enter</td>
    </tr>
    <tr>
      <td>`multiline`</td>
      <td>`multiline`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the form control should accept a value longer than one line</td>
    </tr>
    <tr>
      <td>`name`</td>
      <td>`name`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Name of the form control.</td>
    </tr>
    <tr>
      <td>`pattern`</td>
      <td>`pattern`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Pattern the `value` must match to be valid</td>
    </tr>
    <tr>
      <td>`placeholder`</td>
      <td>`placeholder`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td>Text that appears in the form control when it has no value set</td>
    </tr>
    <tr>
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether to display the form control with no visible background</td>
    </tr>
    <tr>
      <td>`readonly`</td>
      <td>`readonly`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether a user can interact with the value of the form control</td>
    </tr>
    <tr>
      <td>`required`</td>
      <td>`required`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the form control will be found to be invalid when it holds no `value`</td>
    </tr>
    <tr>
      <td>`rows`</td>
      <td>`rows`</td>
      <td>`number`</td>
      <td>`-1`</td>
      <td>The specific number of rows the form control should provide in the user interface</td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
    </tr>
    <tr>
      <td>`valid`</td>
      <td>`valid`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the `value` held by the form control is valid.</td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string | number`</td>
      <td>``</td>
      <td>The value held by the form control</td>
    </tr>
    <tr>
      <td>`viewColor`</td>
      <td>`view-color`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>`An alteration to the value of the color-field has been committed by the user.`</td>
    </tr>
    <tr>
      <td>`input`</td>
      <td>`Event`</td>
      <td>`The value of the color-field has changed.`</td>
    </tr>
  </tbody>
</table>
