# sp-slider-handle
Overview API
## Overview
Some advanced slider uses require more than one handle. One example of this is the range slider. `<sp-slider>` supports multiple handles via the `<sp-slider-handle>` sub-component, although it would be very rare to ever require more than two handles.
`<sp-slider-handle>` is unnecessary for single-handle sliders. Always slot two or more `<sp-slider-handle>` components together. To customize the properties of a single-handle slider (`normalization`, `value`, etc), set them on the `<sp-slider>` element directly.
### Usage
    
    yarn add @spectrum-web-components/slider
    
Import the side effectful registration of `<sp-slider>` and `<sp-slider-handle>` via:
    
    import '@spectrum-web-components/slider/sp-slider.js';
    import '@spectrum-web-components/slider/sp-slider-handle.js';
    
## Anatomy
When two or more sliders are present, the `label` attribute can be used to identify each handle to assistive technology. For form submission, the `name` property is a unique identifier for each handle. Each handle will also have its own `value` based on its position on the slider.
    
    <sp-slider variant="range" step="1" min="0" max="255">
        Output Levels
        <sp-slider-handle
            slot="handle"
            name="min"
            label="Minimum"
            value="5"
        ></sp-slider-handle>
        <sp-slider-handle
            slot="handle"
            name="max"
            label="Maximum"
            value="250"
        ></sp-slider-handle>
    </sp-slider>
For slider handles that have the same numeric range, you can specify `min="previous"` or `max="next"` to constrain handles by the values of their `previous/nextElementSiblings`. Keep in mind that the *first* slider handle with not have a `previous` handle to be its `min` and the *last* slider handle will not have a `next` handle to be its `max`.
    
    <sp-slider step="1" min="0" max="255">
        Output Levels
        <sp-slider-handle
            slot="handle"
            name="low"
            label="Low"
            value="5"
            max="next"
        ></sp-slider-handle>
        <sp-slider-handle
            slot="handle"
            name="mid"
            label="Mid"
            value="100"
            min="previous"
            max="next"
        ></sp-slider-handle>
        <sp-slider-handle
            slot="handle"
            name="high"
            label="High"
            value="250"
            min="previous"
        ></sp-slider-handle>
    </sp-slider>
### Accessibility
#### Include a label
For sliders with more than one handle, each slider handle should have a label. A slider without a label is ambiguous and not accessible.
Review the accessibility guidelines for the slider.
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
      <td>`defaultValue`</td>
      <td>`default-value`</td>
      <td>`number`</td>
      <td>``</td>
      <td>Set the default value of the handle. Setting this property will cause the handle to reset to the default value on double click or pressing the `escape` key.</td>
    </tr>
    <tr>
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Disable this control. It will not receive focus or events</td>
    </tr>
    <tr>
      <td>`dragging`</td>
      <td>`dragging`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`formatOptions`</td>
      <td>`format-options`</td>
      <td>`NumberFormatOptions | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`highlight`</td>
      <td>`highlight`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`max`</td>
      <td>`max`</td>
      <td>`number | 'next' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`min`</td>
      <td>`min`</td>
      <td>`number | 'previous' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`name`</td>
      <td>`name`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`step`</td>
      <td>`step`</td>
      <td>`number | undefined`</td>
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
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`number`</td>
      <td>``</td>
      <td>By default, the value of a Slider Handle will be halfway between its `min` and `max` values, or the `min` value when `max` is less than `min`.</td>
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
      <td>`An alteration to the value of the element has been committed by the user.`</td>
    </tr>
    <tr>
      <td>`input`</td>
      <td>`Event`</td>
      <td>`The value of the element has changed.`</td>
    </tr>
    <tr>
      <td>`sp-slider-handle-ready`</td>
      <td>`CustomEvent`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
