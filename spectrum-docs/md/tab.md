# sp-tab
Overview API
## Description
An `<sp-tab>` element surfaces a `label` attribute to serve as its default text content when none is available in the DOM. An icon may be assigned to the element via the `icon` slot; e.g. `<sp-tab><svg slot="icon" aria-label="Tab w/ Icon">...</svg></sp-tab>`. Associate an `<sp-tab>` element with the `<sp-tab-panel>` that represents its content with the `value` attribute.
### Usage
    
    yarn add @spectrum-web-components/tabs
    
Import the side effectful registration of `<sp-tab>` via:
    
    import '@spectrum-web-components/tabs/sp-tab.js';
    
When looking to leverage the `Tab` base class as a type and/or for extension purposes, do so via:
    
    import {
        Tab,
    } from '@spectrum-web-components/tabs';
    
## Examples
    
    <sp-tabs selected="1">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab value="2">Tab 2</sp-tab>
        <sp-tab label="Tab 3" value="3">
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab vertical value="4">
            Tab 4
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab-panel value="1">
            Content for Tab 1 which uses an attribute for its content delivery
        </sp-tab-panel>
        <sp-tab-panel value="2">
            Content for Tab 2 which uses its text content directly
        </sp-tab-panel>
        <sp-tab-panel value="3">
            Content for Tab 3 which uses an attribute with a
            `[slot="icon"]`
            child
        </sp-tab-panel>
        <sp-tab-panel value="4">
            Content for Tab 4 which uses both its text content and a
            `[slot="icon"]`
            child displayed using the
            `[vertical]`
            attribute to define their alignment
        </sp-tab-panel>
    </sp-tabs>
    
    <sp-tabs selected="1" direction="vertical">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab value="2">Tab 2</sp-tab>
        <sp-tab label="Tab 3" value="3">
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab vertical value="4">
            Tab 4
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab-panel value="1">
            Content for Tab 1 which uses an attribute for its content delivery
        </sp-tab-panel>
        <sp-tab-panel value="2">
            Content for Tab 2 which uses its text content directly
        </sp-tab-panel>
        <sp-tab-panel value="3">
            Content for Tab 3 which uses an attribute with a
            `[slot="icon"]`
            child
        </sp-tab-panel>
        <sp-tab-panel value="4">
            Content for Tab 4 which uses both its text content and a
            `[slot="icon"]`
            child displayed using the
            `[vertical]`
            attribute to define their alignment
        </sp-tab-panel>
    </sp-tabs>
### Disabled
When an `<sp-tab>` element is given the `disabled` attribute it will prevent visitor from selecting that tab and its contents. The ability to select other tabs and their content will go unimpeaded.
    
    <sp-tabs selected="2">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3" disabled></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1 is selectable</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2 is selected</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3 is not selectable</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4 is selectable</sp-tab-panel>
    </sp-tabs>
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
      <td>`disabled`</td>
      <td>`disabled`</td>
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
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string`</td>
      <td>`''`</td>
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
      <td>`default slot`</td>
      <td>text label of the Tab</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>The icon that appears on the left of the label</td>
    </tr>
  </tbody>
</table>
