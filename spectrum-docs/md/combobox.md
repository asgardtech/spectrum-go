# sp-combobox
Overview API
## Overview
An `<sp-combobox>` allows users to filter lists to only the options matching a query. It's composed of a textfield, a picker button, and child menu items.
### Usage
    
    yarn add @spectrum-web-components/combobox
    
Import the side effectful registration of `<sp-combobox>` via:
    
    import '@spectrum-web-components/combobox/sp-combobox.js';
    
When looking to leverage the `Combobox` base class as a type and/or for extension purposes, do so via:
    
    import { Combobox } from '@spectrum-web-components/combobox';
    
### Anatomy
Combobox options are presented as a popup menu. Menu items can be provided via markup as `<sp-menu-item>` children, or by assigning an array to the `options` property of an `<sp-combobox>`.
#### Menu items via the `options` property
Instead of providing `<sp-menu-item>` children, you can assign an array of `ComboboxOptions` to the `options` property, and `<sp-combobox>` will create matching menu items:
    
    <sp-combobox id="color" label="Color"></sp-combobox>
    
    <script>
        document.getElementById('color').options = [
            { value: "red", itemText: "Red" },
            { value: "green", itemText: "Green" },
            { value: "blue", itemText: "Blue" }
        ];
    </script>
#### Menu items via dynamic options
When you replace the `options` Array, or add/remove `<sp-menu-item>` children, the `<sp-combobox>` will detect that change and update its popup menu contents. For example, using Lit:
    
    render() {
        return html`<sp-combobox label="Color" .options=${this.colorOptions}></sp-combobox>`;
    }
    
    mutate() {
        this.colorOptions = [
            ...this.colorOptions,
            { value: 'purple', itemText: 'Purple' }
        ];
    }
### Options
#### Sizes
Small
    
    <sp-combobox size="s" label="Color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
Medium
    
    <sp-combobox size="m" label="Color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
Large
    
    <sp-combobox size="l" label="Color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
Extra Large
    
    <sp-combobox size="xl" label="Color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
### Quiet
    
    <sp-field-label for="color">Color</sp-field-label>
    <sp-combobox id="color" quiet>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
#### Autocomplete
The text in an `<sp-combobox>` is editable, and the string the user has typed in will become the `value` of the combobox unless the user selects a different value in the popup menu.
##### None
`autocomplete="none"`
The suggested popup menu items will remain the same regardless of the currently-input value. Whenever the currently-typed input exactly matches the `value` of a popup menu item, that item is automatically selected.
    
    <sp-field-label for="color-none" autocomplete="none">Color</sp-field-label>
    <sp-combobox id="color-none" disabled>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
#### List
`autocomplete="list"`
The popup menu items are filtered to only those completing the currently-input value.
    
    <sp-field-label for="color-list" autocomplete="list">Color</sp-field-label>
    <sp-combobox id="color-list" disabled>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
## States
Disabled
    
    <sp-field-label for="color-disabled">Color</sp-field-label>
    <sp-combobox id="color-disabled" disabled>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
    <br />
    <sp-field-label for="color-disabled-item">Color</sp-field-label>
    <sp-combobox id="color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue" disabled>Blue</sp-menu-item>
    </sp-combobox>
Invalid
    
    <sp-field-label for="color-invalid">Color</sp-field-label>
    <sp-combobox id="color-invalid" invalid>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
        <sp-help-text slot="negative-help-text">
            Choose or add at least one color.
        </sp-help-text>
    </sp-combobox>
Pending
    
    <sp-field-label for="color">Color</sp-field-label>
    <sp-combobox id="color" pending>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
### Accessibility
#### Provide a label
A combobox must be labeled. Typically, you should render a visible label via `<sp-field-label>`. For exceptional cases, provide an accessible label via the `label` attribute.
    
    <sp-field-label for="color">Color</sp-field-label>
    <sp-combobox id="color">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
#### Provide help text and tooltips in the correct location
It is not currently possible to provide accessible ARIA references between elements in different shadow roots. To ensure proper association between elements, help text must be included via the `slot="help-text"` or `slot="help-text-negative"` and tooltips must be included via the `slot="tooltip"`.
See help text and tooltip for more information.
Help text
    
    <sp-field-label for="color1">Color</sp-field-label>
    <sp-combobox id="color1">
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
        <sp-help-text slot="help-text">Enter the name of a color.</sp-help-text>
    </sp-combobox>
Negative help text
    
    <sp-field-label for="color2">Color</sp-field-label>
    <sp-combobox id="color2" required>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
        <sp-help-text slot="help-text">Enter the name of a color.</sp-help-text>
        <sp-help-text slot="negative-help-text">A color is required.</sp-help-text>
    </sp-combobox>
Tooltip
    
    <sp-field-label for="color3">Color</sp-field-label>
    <sp-combobox id="color3">
        <sp-tooltip slot="tooltip">
            Color options, such as red, green, or blue.
        </sp-tooltip>
        <sp-menu-item value="red">Red</sp-menu-item>
        <sp-menu-item value="green">Green</sp-menu-item>
        <sp-menu-item value="blue">Blue</sp-menu-item>
    </sp-combobox>
#### Do not override keyboard navigation
The combobox supports both mouse and keyboard navigation. Mobile behavior is currently unspecified.
When an `<sp-combobox>` is focused, pressing the down arrow moves focus to the first menu item in the popup menu. The up and down arrows then move between available menu items.
The escape key dismisses the popup menu if open. Otherwise, it clears the combobox's textfield.
The enter key sets the `value` of the focused `<sp-combobox>`. If the popup menu is open, the value is set to the `value` of the selected menu item, returning focus back to the combobox's textfield.
For comprehensive information on combobox accessibility, see WAI ARIA Authoring Practice Guide's Menu Button Pattern.
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
      <td>`'none'`</td>
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
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the listbox is visible.</td>
    </tr>
    <tr>
      <td>`options`</td>
      <td>`options`</td>
      <td>`(ComboboxOption | MenuItem)[] | undefined`</td>
      <td>``</td>
      <td>An array of options to present in the Menu provided while typing into the input</td>
    </tr>
    <tr>
      <td>`pattern`</td>
      <td>`pattern`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Pattern the `value` must match to be valid</td>
    </tr>
    <tr>
      <td>`pending`</td>
      <td>`pending`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the items are currently loading.</td>
    </tr>
    <tr>
      <td>`pendingLabel`</td>
      <td>`pending-label`</td>
      <td>`string`</td>
      <td>`'Pending'`</td>
      <td>Defines a string value that labels the Combobox while it is in pending state.</td>
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
      <td>Supply Menu Item elements to the default slot in order to populate the available options</td>
    </tr>
    <tr>
      <td>`help-text`</td>
      <td>default or non-negative help text to associate to your form element</td>
    </tr>
    <tr>
      <td>`negative-help-text`</td>
      <td>negative help text to associate to your form element when `invalid`</td>
    </tr>
    <tr>
      <td>`tooltip`</td>
      <td>Tooltip to to be applied to the the Picker Button</td>
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
  </tbody>
</table>
