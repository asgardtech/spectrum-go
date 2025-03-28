# sp-checkbox
Overview API
## Description
`<sp-checkbox>` allow users to select multiple items from a list of independent options, or to mark an individual option as selected.
Should I use a checkbox or a switch? Use a switch when activating something instead of selecting.
### Usage
    
    yarn add @spectrum-web-components/checkbox
    
Import the side effectful registration of `<sp-checkbox>` via:
    
    import '@spectrum-web-components/checkbox/sp-checkbox.js';
    
When looking to leverage the `Checkbox` base class as a type and/or for extension purposes, do so via:
    
    import { Checkbox } from '@spectrum-web-components/checkbox';
    
## Sizes
Small
    
    <sp-field-group>
        <sp-checkbox size="s">Small</sp-checkbox>
        <sp-checkbox size="s" checked>Small Checked</sp-checkbox>
        <sp-checkbox size="s" indeterminate>Small Indeterminate</sp-checkbox>
    </sp-field-group>
Medium
    
    <sp-field-group>
        <sp-checkbox size="m">Medium</sp-checkbox>
        <sp-checkbox size="m" checked>Medium Checked</sp-checkbox>
        <sp-checkbox size="m" indeterminate>Medium Indeterminate</sp-checkbox>
    </sp-field-group>
Large
    
    <sp-field-group>
        <sp-checkbox size="l">Large</sp-checkbox>
        <sp-checkbox size="l" checked>Large Checked</sp-checkbox>
        <sp-checkbox size="l" indeterminate>Large Indeterminate</sp-checkbox>
    </sp-field-group>
Extra Large
    
    <sp-field-group>
        <sp-checkbox size="xl">Extra Large</sp-checkbox>
        <sp-checkbox size="xl" checked>Extra Large Checked</sp-checkbox>
        <sp-checkbox size="xl" indeterminate>Extra Large Indeterminate</sp-checkbox>
    </sp-field-group>
## Variants
### Standard checkboxes
Standard checkboxes are the default style for checkboxes. They are optimal for application panels where all visual elements are monochrome in order to direct focus to the content.
    
    <div style="display: flex; justify-content: space-between;">
        <div style="display: flex; flex-direction: column;">
            <h4 class="spectrum-Heading--subtitle1">Default</h4>
            <sp-checkbox>Web component</sp-checkbox>
            <sp-checkbox checked>Web component</sp-checkbox>
            <sp-checkbox indeterminate>Web component</sp-checkbox>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <h4 class="spectrum-Heading--subtitle1">Invalid</h4>
            <sp-checkbox invalid>Web component</sp-checkbox>
            <sp-checkbox checked invalid>Web component</sp-checkbox>
            <sp-checkbox indeterminate invalid>Web component</sp-checkbox>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <h4 class="spectrum-Heading--subtitle1">Disabled</h4>
            <sp-checkbox disabled>Web component</sp-checkbox>
            <sp-checkbox checked disabled>Web component</sp-checkbox>
            <sp-checkbox indeterminate disabled>Web component</sp-checkbox>
        </div>
    </div>
### Emphasized checkboxes
Emphasized checkboxes are a secondary style for checkboxes. The blue color provides a visual prominence that is optimal for forms, settings, lists or grids of assets, etc. where the checkboxes need to be noticed.
    
    <div style="display: flex; justify-content: space-between;">
        <div style="display: flex; flex-direction: column; justify-content: space-between;">
            <h4 class="spectrum-Heading--subtitle1">Default</h4>
            <sp-checkbox emphasized>Web component</sp-checkbox>
            <sp-checkbox emphasized checked>Web component</sp-checkbox>
            <sp-checkbox emphasized indeterminate>Web component</sp-checkbox>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <h4 class="spectrum-Heading--subtitle1">Invalid</h4>
            <sp-checkbox emphasized invalid>Web component</sp-checkbox>
            <sp-checkbox emphasized checked invalid>Web component</sp-checkbox>
            <sp-checkbox emphasized indeterminate invalid>Web component</sp-checkbox>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <h4 class="spectrum-Heading--subtitle1">Disabled</h4>
            <sp-checkbox emphasized disabled>Web component</sp-checkbox>
            <sp-checkbox emphasized checked disabled>Web component</sp-checkbox>
            <sp-checkbox emphasized indeterminate disabled>Web component</sp-checkbox>
        </div>
    </div>
### Handling events
Event handlers for clicks and other user actions can be registered on an `<sp-checkbox>` as they would a standard `<input type="checkbox">` element.
    
    <sp-checkbox
        id="checkbox-example"
        onclick="spAlert(this, '<sp-radio> clicked!')"
    >
        Web component
    </sp-checkbox>
## Accessibility
Checkboxes are accessible by default, rendered in HTML using the `<input type="checkbox">` element. When the checkbox is set as `indeterminate` or `invalid`, the appropriate ARIA state attribute will automatically be applied.
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
      <td>`checked`</td>
      <td>`checked`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`indeterminate`</td>
      <td>`indeterminate`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`invalid`</td>
      <td>`invalid`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`name`</td>
      <td>`name`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`readonly`</td>
      <td>`readonly`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabindex`</td>
      <td>`number`</td>
      <td>`0`</td>
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
      <td>content to display as the label for the Checkbox</td>
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
      <td>`Announces a change in the `checked` property of a Checkbox`</td>
    </tr>
  </tbody>
</table>
