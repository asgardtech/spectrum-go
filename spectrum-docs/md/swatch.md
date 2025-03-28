# sp-swatch
Overview API
## Description
An `<sp-swatch>` shows a small sample of a fill — such as a color, gradient, texture, or material — that is intended to be applied to an object.
### Usage
    
    yarn add @spectrum-web-components/swatch
    
Import the side effectful registration of `<sp-swatch>` via:
    
    import '@spectrum-web-components/swatch/sp-swatch.js';
    
When looking to leverage the `Swatch` base class as a type and/or for extension purposes, do so via:
    
    import { Swatch } from '@spectrum-web-components/swatch';
    
## Sizes
Extra Small
    
    <div style="display: flex;gap: 5px;">
        <sp-swatch color="rgb(255 0 0 / 0.7)" size="xs"></sp-swatch>
        <sp-swatch rounding="none" color="rgb(255 0 0 / 0.7)" size="xs"></sp-swatch>
        <sp-swatch rounding="full" color="rgb(255 0 0 / 0.7)" size="xs"></sp-swatch>
        <sp-swatch border="light" color="rgb(255 0 0 / 0.7)" size="xs"></sp-swatch>
        <sp-swatch border="none" color="rgb(255 0 0 / 0.7)" size="xs"></sp-swatch>
        <sp-swatch nothing size="xs"></sp-swatch>
        <sp-swatch
            shape="rectangle"
            color="rgb(255 0 0 / 0.7)"
            size="xs"
        ></sp-swatch>
        <sp-swatch
            shape="rectangle"
            disabled
            color="rgb(255 0 0 / 0.7)"
            size="xs"
        ></sp-swatch>
        <sp-swatch
            rounding="full"
            shape="rectangle"
            mixed-value
            size="xs"
        ></sp-swatch>
    </div>
Small
    
    <div style="display: flex;gap: 5px;">
        <sp-swatch color="rgb(255 0 0 / 0.7)" size="s"></sp-swatch>
        <sp-swatch rounding="none" color="rgb(255 0 0 / 0.7)" size="s"></sp-swatch>
        <sp-swatch rounding="full" color="rgb(255 0 0 / 0.7)" size="s"></sp-swatch>
        <sp-swatch border="light" color="rgb(255 0 0 / 0.7)" size="s"></sp-swatch>
        <sp-swatch border="none" color="rgb(255 0 0 / 0.7)" size="s"></sp-swatch>
        <sp-swatch nothing size="s"></sp-swatch>
        <sp-swatch
            shape="rectangle"
            color="rgb(255 0 0 / 0.7)"
            size="s"
        ></sp-swatch>
        <sp-swatch
            shape="rectangle"
            disabled
            color="rgb(255 0 0 / 0.7)"
            size="s"
        ></sp-swatch>
        <sp-swatch
            rounding="full"
            shape="rectangle"
            mixed-value
            size="s"
        ></sp-swatch>
    </div>
Medium
    
    <div style="display: flex;gap: 5px;">
        <sp-swatch color="rgb(255 0 0 / 0.7)" size="m"></sp-swatch>
        <sp-swatch rounding="none" color="rgb(255 0 0 / 0.7)" size="m"></sp-swatch>
        <sp-swatch rounding="full" color="rgb(255 0 0 / 0.7)" size="m"></sp-swatch>
        <sp-swatch border="light" color="rgb(255 0 0 / 0.7)" size="m"></sp-swatch>
        <sp-swatch border="none" color="rgb(255 0 0 / 0.7)" size="m"></sp-swatch>
        <sp-swatch nothing size="m"></sp-swatch>
        <sp-swatch
            shape="rectangle"
            color="rgb(255 0 0 / 0.7)"
            size="m"
        ></sp-swatch>
        <sp-swatch
            shape="rectangle"
            disabled
            color="rgb(255 0 0 / 0.7)"
            size="m"
        ></sp-swatch>
        <sp-swatch
            rounding="full"
            shape="rectangle"
            mixed-value
            size="m"
        ></sp-swatch>
    </div>
Large
    
    <div style="display: flex;gap: 5px;">
        <sp-swatch color="rgb(255 0 0 / 0.7)" size="l"></sp-swatch>
        <sp-swatch rounding="none" color="rgb(255 0 0 / 0.7)" size="l"></sp-swatch>
        <sp-swatch rounding="full" color="rgb(255 0 0 / 0.7)" size="l"></sp-swatch>
        <sp-swatch border="light" color="rgb(255 0 0 / 0.7)" size="l"></sp-swatch>
        <sp-swatch border="none" color="rgb(255 0 0 / 0.7)" size="l"></sp-swatch>
        <sp-swatch nothing size="l"></sp-swatch>
        <sp-swatch
            shape="rectangle"
            color="rgb(255 0 0 / 0.7)"
            size="l"
        ></sp-swatch>
        <sp-swatch
            shape="rectangle"
            disabled
            color="rgb(255 0 0 / 0.7)"
            size="l"
        ></sp-swatch>
        <sp-swatch
            rounding="full"
            shape="rectangle"
            mixed-value
            size="l"
        ></sp-swatch>
    </div>
## Modifying attributes
An `<sp-swatch>` element can be modified by the following attributes/properties to customize its delivery as desired for your use case: `border`, `color`, `disabled`, `mixedValue` (accepted as the `mixed-value` attribute), `nothing`, `rounding`, and `shape`. Use these in concert with each other for a variety of final visual deliveries.
### Border
The `border` attribute/property is not required and when applied accepts the values of `none` or `light`.
    
    <sp-swatch-group>
        <sp-swatch color="rgb(255 0 0 / 0.7)"></sp-swatch>
        <sp-swatch color="rgb(255 0 0 / 0.7)" border="light"></sp-swatch>
        <sp-swatch color="rgb(255 0 0 / 0.7)" border="none"></sp-swatch>
    </sp-swatch-group>
### Color
The `color` attribute/property determines the color value that the `<sp-swatch>` element will deliver.
    
    <sp-swatch-group>
        <sp-swatch color="rgb(255 0 0 / 0.7)"></sp-swatch>
        <sp-swatch color="orange"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
### Disabled
The `disabled` attribute/property determines prevents interaction on the `<sp-swatch>` element.
    
    <sp-swatch-group>
        <sp-swatch disabled color="rgb(255 0 0 / 0.7)"></sp-swatch>
        <sp-swatch disabled color="orange"></sp-swatch>
        <sp-swatch disabled color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
### Mixed value
The `mixed-value` attribute and `mixedValue` property outline when an `<sp-swatch>` element represents more than one color.
    
    <sp-swatch-group selects="multiple">
        <sp-swatch mixed-value></sp-swatch>
        <sp-swatch mixed-value rounding="full"></sp-swatch>
        <sp-swatch mixed-value shape="rectangle"></sp-swatch>
    </sp-swatch-group>
Please note that the `aria-checked="mixed"` value only applies when the swatch is in a group with `selects="multiple"`
### Nothing
The `nothing` attribute/property outlines that the `<sp-swatch>` represents no color or that it represents "transparent".
    
    <sp-swatch-group>
        <sp-swatch nothing></sp-swatch>
        <sp-swatch nothing rounding="full"></sp-swatch>
        <sp-swatch nothing shape="rectangle"></sp-swatch>
    </sp-swatch-group>
### Rounding
The `rounding` attribute/property is not required and when applied accepts the values of `none` or `full`.
    
    <sp-swatch-group>
        <sp-swatch color="rgb(255 0 0 / 0.7)"></sp-swatch>
        <sp-swatch color="rgb(255 0 0 / 0.7)" rounding="none"></sp-swatch>
        <sp-swatch color="rgb(255 0 0 / 0.7)" rounding="full"></sp-swatch>
    </sp-swatch-group>
### Shape
The `shape` attribute/property is not required and when applied accepts the values of `rectangle`.
    
    <sp-swatch-group>
        <sp-swatch color="rgb(255 0 0 / 0.7)"></sp-swatch>
        <sp-swatch color="rgb(255 0 0 / 0.7)" shape="rectangle"></sp-swatch>
    </sp-swatch-group>
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
      <td>`border`</td>
      <td>`border`</td>
      <td>`SwatchBorder`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`color`</td>
      <td>`color`</td>
      <td>`string`</td>
      <td>`''`</td>
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
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`mixedValue`</td>
      <td>`mixed-value`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`nothing`</td>
      <td>`nothing`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`role`</td>
      <td>`role`</td>
      <td>`string`</td>
      <td>`'button'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`rounding`</td>
      <td>`rounding`</td>
      <td>`SwatchRounding`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`shape`</td>
      <td>`shape`</td>
      <td>`SwatchShape`</td>
      <td>``</td>
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
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`change`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
