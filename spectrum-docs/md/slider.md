# sp-slider
Overview API
## Overview
`<sp-slider>` allows users to quickly select a value within a range. They should be used when the upper and lower bounds of the range are invariable.
### Usage
    
    yarn add @spectrum-web-components/slider
    
Import the side effectful registration of `<sp-slider>` via:
    
    import '@spectrum-web-components/slider/sp-slider.js';
    
When leveraging the `editable` attribute, the `@spectrum-web-components/number-field/sp-number-field.js` dependency will be asynchronously loaded via a dynamic import to reduce JS payload for applications not leveraging this feature. In the case that you would like to import those transverse dependencies statically, import the side effectful registration of `<sp-slider>` as follows:
    
    import '@spectrum-web-components/slider/sync/sp-slider.js';
    
When looking to leverage the `Slider` base class as a type and/or for extension purposes, do so via:
    
    import { Slider } from '@spectrum-web-components/slider';
    
### Anatomy
    
    <sp-slider label="Slider Label"></sp-slider>
    <sp-slider label="Slider Label - Disabled" disabled></sp-slider>
#### Label visibility
By default, an `<sp-slider>` element has both a "text" label and a "value" label. The "value" label is output by the element itself based on its state, but the "text" label must be supplied by the consuming developer in order for the `<sp-slider>` to be delivered in an accessible manner.
Either or both of these can be suppressed visually as needed by your application UI, while still fulfilling their role in delivering a quality accessibility tree to the browser. This delivery is controlled by the `label-visibility` attribute (or `labelVisibility` property) which accepts `text`, `value`, or `none` as values.
Use `label-visibility="text"` to suppress the "value" label, use `label-visibility="value"` to suppress the "text" label, or use `label-visibility="none"` to suppress the "text" label. In each case outlined above the content for both labels will be made available to screen readers, so be sure to manage the content delivered to visitors in that context.
Text
    
    <sp-slider
        label="No visible value label"
        label-visibility="text"
        value="50"
    ></sp-slider>
Value
    
    <sp-slider label="No visible text label" label-visibility="value"></sp-slider>
None
    
    <sp-slider label="No visible labels" label-visibility="none"></sp-slider>
#### Handle slot
The slider also optionally accepts two or more `<sp-slider-handle>` elements with `slot="handle"`:
    
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
### Options
#### Sizes
Small
    
    <sp-slider label="Slider Label" size="s"></sp-slider>
    <sp-slider label="Slider Label - Editable" editable size="s"></sp-slider>
Medium
    
    <sp-slider label="Slider Label"></sp-slider>
    <sp-slider label="Slider Label - Editable" editable></sp-slider>
Large
    
    <sp-slider label="Slider Label" size="l"></sp-slider>
    <sp-slider label="Slider Label - Editable" editable size="l"></sp-slider>
Extra Large
    
    <sp-slider label="Slider Label" size="xl"></sp-slider>
    <sp-slider label="Slider Label - Editable" editable size="xl"></sp-slider>
#### Editable
An `<sp-slider>` element can be paired with an `<sp-number-field>` element via the `editable` attribute. The `<sp-number-field>` will be passed all of the standard options from the `<sp-slider>` element (e.g. `min`, `max`, `formatOptions`, etc.) and will also accept the `hide-stepper` attribute in order to prevent the display of its stepper UI.
The `quiet` attribute applies Quiet styling to the number field when a slider is `editable`.
Editable
    
    <sp-slider
        label="Hours of the day (editable)"
        editable
        max="24"
        min="0"
        value="7.25"
        step="0.25"
        style="--spectrum-slider-editable-number-field-width: 100px;"
        format-options='{
            "style": "unit",
            "unit": "hour"
        }'
    ></sp-slider>
Editable, hide-stepper
    
    <sp-slider
        label="Angle (editable)"
        editable
        hide-stepper
        min="0"
        max="360"
        format-options='{
            "style": "unit",
            "unit": "degree",
            "unitDisplay": "narrow"
        }'
    ></sp-slider>
Editable, quiet
    
    <sp-slider quiet editable value="50"></sp-slider>
    <sp-slider quiet editable hide-stepper value="50"></sp-slider>
#### Filled
Use `variant="filled"` to add a filled style to the slider from a starting point to the current value. By default the starting point is the `min` value.
    
    <sp-slider
        label="Slider Label"
        max="1"
        variant="filled"
        min="0"
        value=".5"
        step="0.01"
    ></sp-slider>
    <sp-slider
        label="Slider Label - Disabled"
        max="1"
        variant="filled"
        min="0"
        value=".5"
        step="0.01"
        disabled
    ></sp-slider>
##### fill-start
When both `fill-start` and `variant="filled"` are used in `<sp-slider>`, the `fill-start` property defines the starting point. If fill start does not have a number associated with it, the starting point will be the `value`.
Any number (including `0`) can be used as a fill-start value. If a custom normalization function is provided, it will also normalize all fill-related params.
fill-start
    
    <sp-slider
        label="Slider Label"
        max="1"
        fill-start
        variant="filled"
        min="0"
        value=".5"
        step="0.01"
    ></sp-slider>
    <sp-slider
        label="Slider Label"
        max="1"
        fill-start
        variant="filled"
        min="0"
        value=".8"
        step="0.01"
        disabled
    ></sp-slider>
fill-start > value
    
    <sp-slider
        id="fill-start-slider"
        label="fill-start greater than value"
        max="1"
        min="0"
        value=".3"
        step="0.1"
        fill-start="0.7"
        variant="filled"
    ></sp-slider>
fill-start < value
    
    <sp-slider
        id="fill-start-slider"
        label="Fill Start less than Value"
        max="1"
        min="0"
        value=".7"
        step="0.1"
        fill-start="0.25"
        variant="filled"
    ></sp-slider>
fill-start="0" and negative min
    
    <sp-slider
        label="fill-start set to 0"
        max="1"
        min="-1"
        value=".7"
        step="0.1"
        fill-start="0"
        variant="filled"
    ></sp-slider>
#### Tick
With `variant="tick"`, ticks are applied at intervals defined with the `tick-step` attribute. The `tick-labels` attribute will apply labels to the ticks.
Tick with No Labels
    
    <sp-slider label="Slider Label" variant="tick" tick-step="5"></sp-slider>
    <sp-slider
        label="Slider Label - Disabled"
        variant="tick"
        tick-step="5"
        disabled
    ></sp-slider>
Tick with Labels
    
    <sp-slider
        label="Slider Label"
        variant="tick"
        tick-step="5"
        tick-labels
    ></sp-slider>
    <sp-slider
        label="Slider Label - Disabled"
        variant="tick"
        tick-step="5"
        tick-labels
        disabled
    ></sp-slider>
#### Ramp
With `variant="ramp"`, the slider with increases as it approaches the `max` position.
    
    <sp-slider label="Slider Label" variant="ramp"></sp-slider>
    <sp-slider label="Slider Label - Disabled" variant="ramp" disabled></sp-slider>
### Range
The `"range"` variant along with two handles to create a range slider. (See slider handle.)
    
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
#### Format options for labels
An `<sp-slider>` or `<sp-slider-handle>` element will process its numeric value with `new Intl.NumberFormat(this.resolvedLanguage, this.formatOptions).format(this.value)` in order to prepare it for visual delivery in the input. In order to customize this processing supply your own `Intl.NumberFormatOptions` via the `formatOptions` property, or `format-options` attribute as seen below.
`this.resolvedLanguage` represents the language in which the `<sp-slider>` or `<sp-slider-handle>` element is currently being delivered. By default, this value will represent the language established by the `lang` attribute on the root `<html>` element while falling back to `navigator.language` when that is not present. This value can be customized via a language context provided by a parent element that listens for the `sp-language-context` event and supplies updated language settings to the `callback` function contained therein. Applications leveraging the `<sp-theme>` element to manage the visual delivery or text direction of their content will also be provided a reactive context for supplying language information to its descendants.
    
    <sp-slider
        min="0"
        max="1"
        step="0.01"
        value="0.5"
        label="Slider Label"
        format-options='{
            "style": "percent"
        }'
    ></sp-slider>
More advanced formatting is available by specifying a formatting function to the `getAriaHandleText` property on an `sp-slider` or `sp-slider-handle`. Or, for a multi-handle slider, you can format the combined value label for all handles by passing a formatting function to the `getAriaValueText` property on the parent `sp-slider`.
##### Units not included in `Intl.NumberFormatOptions`
While `Intl.NumberFormatOptions` does support a wide range of units, it is possible to encounter units (e.g. the graphics units of `pixel`, `pixels`, `points`, etc.) that are not supported therein. When this occurs, an `<sp-slider>` element will attempt to polyfill support for this unit. See the following example delivering `{ style: "unit", unit: "px" }` below:
    
    <sp-slider
        style="width: 200px"
        value="500"
        format-options='{
            "style": "unit",
            "unit": "px"
        }'
    >
        Document width in pixels
    </sp-slider>
Note: the polyfilling done here is very simplistic and is triggered by supplying options that would otherwise cause the `Intl.NumberFormat()` call to throw an error. Once the unsupporting unit of `px` causes the construction of the object to throw, a backup formatter/parser pair will be created without the supplied unit data. When the `style` is set to `unit`, the `unit` value will be adopted as the *static* unit display. This means that neither pluralization nor translation will be handled within the `<sp-number-field>` element itself. If pluralization or translation is important to the delivered interface, please be sure to handle passing those strings into to element via the `formatOptions` property reactively to the value of the element or locale of that page in question.
### Behavior
#### Default value
Slider will reset to its `default-value` when the user double clicks on the slider handle or if the user presses the `escape` key when the slider handle is focused.
    
    <sp-slider value="50" default-value="20"></sp-slider>
Note: If a slider with `default-value` attribute is contained in a modal and the slider handle is focused then the following interaction will occur on pressing the `escape` key:
  -  If the slider value is different from the default value then the slider value will be reset to the default value and the modal will not be closed.
  -  If the slider value is equal to the default value then the modal will be closed.

#### Advanced normalization
By default, `sp-slider` assumes a linear scale between the `min` and `max` values. For advanced applications, it is sometimes necessary to specify a custom "normalization."
Normalization is the process of converting a slider to a value between 0 and 1 where 0 represents the minimum and 1 represents the maximum. See the "Three Handles Complex" example in the playground.
### States
#### Disabled
    
    <sp-slider disabled value="50"></sp-slider>
    <sp-slider disabled quiet value="50"></sp-slider>
    <sp-slider disabled editable value="50"></sp-slider>
    <sp-slider disabled variant="filled" value="50"></sp-slider>
    <sp-slider disabled variant="ticks" value="50"></sp-slider>
    <sp-slider disabled variant="ramp" value="50"></sp-slider>
#### Indeterminate
The indeterminate attribute will be passed to the internal `<sp-number-field>` element and alter its visual delivery until a change has been made to the `<sp-slider>` element at which point the `change` event that is dispatched can be understood as always removing the indeterminate attribute from the `<sp-slider>`.
    
    <sp-slider indeterminate value="50"></sp-slider>
    <sp-slider indeterminate quiet value="50"></sp-slider>
    <sp-slider indeterminate editable value="50"></sp-slider>
    <sp-slider indeterminate variant="filled" value="50"></sp-slider>
    <sp-slider indeterminate variant="ticks" value="50"></sp-slider>
    <sp-slider indeterminate variant="ramp" value="50"></sp-slider>
### Events
Like the `<input type="range">` element after which the `<sp-slider>` is fashioned, it will dispatch `input` events in a stream culminating with a `change` event (representing the final commit of the `value` to the element) once the user has finished interacting with the element. Both other these events can access the `value` of their dispatching target via `event.target.value`. In this way, a steaming listener pattern similar to the following can prove useful:
    
    const slider = document.querySelector('sp-slider');
    
    const endListener = ({ target }) => {
        target.addEventListener('input', startListener);
        target.removeEventListener('input', streamListener);
        target.removeEventListener('change', endListener);
        console.log(target.value);
    };
    
    const streamListener = ({ target }) => {
        console.log(target.value);
    };
    
    const startListener = ({ target }) => {
        target.removeEventListener('input', startListener);
        target.addEventListener('input', streamListener);
        target.addEventListener('change', endListener);
        console.log(target.value);
    };
    
    slider.addEventListener('input', startListener);
### Accessibility
#### Include labels
Every slider should have a label. A slider without a label is ambiguous and not accessible. Write the label in sentence case.
In rare cases where context is sufficient and a label doesn't require visibility, make sure to have the design reviewed and approved by an accessibility expert. Use `label-visibility` to set which labels should remain visible, and non-visible labels will still be read by assistive technology.
#### Keyboard navigation
The `Tab` and `Shift+Tab` keys are used to navigate to and set focus on the slider control. The `Arrow Up (↑)` and `Arrow Down (↓)` keys are used to increment the slider value, respectively.
Because multiple sliders are often used on the same page, the number field in the `editable` slider is designed to not be keyboard focusable via `Tab` or `Shift+Tab` keys. (See WAI ARIA APG's Keyboard Navigation Inside Components.) Since the slider itself can already be incremented via the arrow keys, a roving tabindex controller cannot be used within the slider; therefore, the number field will not be keyboard navigable.
Review the accessibility guidelines for the slider handle.
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
      <td>`editable`</td>
      <td>`editable`</td>
      <td>`boolean`</td>
      <td>``</td>
      <td>Whether to display a Number Field along side the slider UI</td>
    </tr>
    <tr>
      <td>`fillStart`</td>
      <td>`fill-start`</td>
      <td>`number | boolean | undefined`</td>
      <td>``</td>
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
      <td>`hideStepper`</td>
      <td>`hide-stepper`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the stepper UI of the Number Field is hidden or not</td>
    </tr>
    <tr>
      <td>`highlight`</td>
      <td>`highlight`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`indeterminate`</td>
      <td>`indeterminate`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Applies `indeterminate` to the underlying `sp-number-field` when `editable === true`. Is removed on the next `change` event.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`labelVisibility`</td>
      <td>`label-visibility`</td>
      <td>`'text' | 'value' | 'none' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`max`</td>
      <td>`max`</td>
      <td>`number | 'next' | undefined`</td>
      <td>`100`</td>
      <td></td>
    </tr>
    <tr>
      <td>`min`</td>
      <td>`min`</td>
      <td>`number | 'previous' | undefined`</td>
      <td>`0`</td>
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
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Applies `quiet` to the underlying `sp-number-field` when `editable === true`.</td>
    </tr>
    <tr>
      <td>`step`</td>
      <td>`step`</td>
      <td>`number | undefined`</td>
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
      <td>`tickLabels`</td>
      <td>`tick-labels`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`tickStep`</td>
      <td>`tick-step`</td>
      <td>`number`</td>
      <td>`0`</td>
      <td></td>
    </tr>
    <tr>
      <td>`type`</td>
      <td>`type`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`number`</td>
      <td>``</td>
      <td>By default, the value of a Slider Handle will be halfway between its `min` and `max` values, or the `min` value when `max` is less than `min`.</td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`string`</td>
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
      <td>`default slot`</td>
      <td>@deprecated Text label for the Slider. Use the `label` property instead.</td>
    </tr>
    <tr>
      <td>`handle`</td>
      <td>optionally accepts two or more sp-slider-handle elements</td>
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
