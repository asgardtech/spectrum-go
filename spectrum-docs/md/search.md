# sp-search
Overview API
## Description
The `<sp-search>` element delivers a single input field with a "reset" button as well as a mangifying glass icon with which to power search interactions.
### Usage
    
    yarn add @spectrum-web-components/search
    
Import the side effectful registration of `<sp-search>` via:
    
    import '@spectrum-web-components/search/sp-search.js';
    
When looking to leverage the `Search` base class as a type and/or for extension purposes, do so via:
    
    import { Search } from '@spectrum-web-components/search';
    
## Sizes
Small
    
    <sp-search size="s"></sp-search>
    <sp-search size="s" disabled></sp-search>
Medium
    
    <sp-search></sp-search>
    <sp-search disabled></sp-search>
Large
    
    <sp-search size="l"></sp-search>
    <sp-search size="l" disabled></sp-search>
Extra Large
    
    <sp-search size="xl"></sp-search>
    <sp-search size="xl" disabled></sp-search>
## Variants
### Quiet
    
    <sp-search quiet></sp-search>
    <sp-search quiet disabled></sp-search>
## Events
The `submit` event fires when the `<sp-search>` is submitted. This is a non-`composed` event inline with what you would expect a `<form />`{:target="_blank"} to fire. If you choose to prevent the default of this event, the default action for the underlying `<form />` element will also be prevented, which will allow you to handle the search action in javascript.
## Properties
### holdValueOnEscape
  -  Type: `boolean`
  -  Default: `false`

If `holdValueOnEscape` controls whether the typed value should be held (i.e., not cleared or reset) when the Escape key is pressed. If set to true, pressing the Escape key will not affect the value in the search field.
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
      <td>`action`</td>
      <td>`action`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
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
      <td>`holdValueOnEscape`</td>
      <td>`holdValueOnEscape`</td>
      <td>`boolean`</td>
      <td>``</td>
      <td></td>
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
      <td>`'Search'`</td>
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
      <td>`method`</td>
      <td>`method`</td>
      <td>`'get' | 'post' | 'dialog' | undefined`</td>
      <td>``</td>
      <td></td>
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
      <td>`'Search'`</td>
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
      <td>`help-text`</td>
      <td>default or non-negative help text to associate to your form element</td>
    </tr>
    <tr>
      <td>`negative-help-text`</td>
      <td>negative help text to associate to your form element when `invalid`</td>
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
      <td>`submit`</td>
      <td>`Event`</td>
      <td>`The search form has been submitted.`</td>
    </tr>
  </tbody>
</table>
