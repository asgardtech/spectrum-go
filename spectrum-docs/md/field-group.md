# sp-field-group
Overview API
## Overview
An `<sp-field-group>` element is used to layout a group of fields, usually `<sp-checkbox>` elements. It can be leveraged for `vertical` or `horizontal` organization of the fields that are supplied as its children.
### Usage
    
    yarn add @spectrum-web-components/field-group
Import the side effectful registration of `<sp-field-group>` via:
    
    import '@spectrum-web-components/field-group/sp-field-group.js';
When looking to leverage the `FieldGroup` base class as a type and/or for extension purposes, do so via:
    
    import { FieldGroup } from '@spectrum-web-components/field-group';
### Anatomy
    
    <sp-field-group horizontal label="Choose from horizontally placed options">
        <sp-checkbox>Checkbox 1</sp-checkbox>
        <sp-checkbox>Checkbox 2</sp-checkbox>
        <sp-checkbox checked>Checkbox 3</sp-checkbox>
        <sp-checkbox>Checkbox 4</sp-checkbox>
        <sp-checkbox>Checkbox 5</sp-checkbox>
    </sp-field-group>
#### Label
A field group must have a label in order to be accessible. A label can be provided either via the `label` attribute, like the previous example or with an `<sp-field-label>` element.
    
    <sp-field-label for="horizontal">
        Choose from horizontally placed options
    </sp-field-label>
    <sp-field-group horizontal id="horizontal">
        <sp-checkbox>Checkbox 1</sp-checkbox>
        <sp-checkbox>Checkbox 2</sp-checkbox>
        <sp-checkbox checked>Checkbox 3</sp-checkbox>
        <sp-checkbox>Checkbox 4</sp-checkbox>
        <sp-checkbox>Checkbox 5</sp-checkbox>
    </sp-field-group>
#### Help text
Help text can be accessibly associated with an `<sp-field-group>` element by using the `help-text` or `negative-help-text` slots. When using the `negative-help-text` slot, `<sp-field-group>` will self manage the presence of this content based on the value of the `invalid` property on your `<sp-field-group>` element. Content within the `help-text` slot will be show by default. When your `<sp-field-group>` should receive help text based on state outside of the complexity of `invalid` or not, manage the content addressed to the `help-text` from above to ensure that it displays the right messaging and possesses the right `variant`.
Self managed
    
    <sp-field-group horizontal id="self" label="What are your favorite fruits?">
        <sp-checkbox value="apple">Apple</sp-checkbox>
        <sp-checkbox
            value="not-a-fruit"
            onchange="javascript:this.parentElement.invalid = this.checked"
        >
            Lettuce
        </sp-checkbox>
        <sp-checkbox value="strawberry" checked>Strawberry</sp-checkbox>
        <sp-help-text slot="help-text">One of these is not a fruit.</sp-help-text>
        <sp-help-text slot="negative-help-text" icon>
            Choose actual fruit(s).
        </sp-help-text>
    </sp-field-group>
Managed from above
    
    <sp-field-label for="above">What are your favorite fruits?</sp-field-label>
    <sp-field-group horizontal id="above">
        <sp-checkbox value="apple">Apple</sp-checkbox>
        <sp-checkbox
            value="not-a-fruit"
            onchange="
                const helpText = this.parentElement.querySelector(`[slot='help-text']`);
                helpText.icon = this.checked;
                helpText.textContent = this.checked ? 'Choose actual fruit(s).' : 'One of these is not a fruit.';
                helpText.variant = this.checked ? 'negative' : 'neutral';
            "
        >
            Lettuce
        </sp-checkbox>
        <sp-checkbox value="strawberry" checked>Strawberry</sp-checkbox>
        <sp-help-text slot="help-text">One of these is not a fruit.</sp-help-text>
    </sp-field-group>
### Options
#### Vertical
    
    <sp-field-label for="vertical">
        Choose from vertically placed options
    </sp-field-label>
    <sp-field-group vertical id="vertical">
        <sp-checkbox>Checkbox 1</sp-checkbox>
        <sp-checkbox>Checkbox 2</sp-checkbox>
        <sp-checkbox>Checkbox 3</sp-checkbox>
        <sp-checkbox>Checkbox 4</sp-checkbox>
        <sp-checkbox checked>Checkbox 5</sp-checkbox>
    </sp-field-group>
### Accessibility
#### Include a label
Every field group should have a label. A field without a label is ambiguous and not accessible.
#### Include help text
The description in the help text is flexible and encompasses a range of guidance. Sometimes this guidance is about what to input, and sometime it’s about how to input. This includes information such as:
  -  An overall description of the input field
  -  Hints for what kind of information needs to be input
  -  Specific formatting examples or requirements

Learn more about using help text.
#### Include negative help text
Write error messaging in a human-centered way by guiding a user and showing them a solution — don’t simply state what’s wrong and then leave them guessing as to how to resolve it. Ambiguous error messages can be frustrating and even shame-inducing for users. Also, keep in mind that something that a system may deem an error may not actually be perceived as an error to a user.
Learn more about writing error messages.
#### Do not us a placeholder as a replacement for a label or help-text
Putting instructions for how to complete an input, requirements, or any other essential information into placeholder text is not accessible. Once a value is entered, placeholder text is no longer viewable; if someone is using an automatic form filler, they will never get the information in the placeholder text.
Instead, use the help text description to convey requirements or to show any formatting examples that would help user comprehension. If there's placeholder text and help text at the same time, it becomes redundant and distracting, especially if they're communicating the same thing.
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
      <td>`horizontal`</td>
      <td>`horizontal`</td>
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
      <td>`label`</td>
      <td>`label`</td>
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
      <td>the form controls that make up the group</td>
    </tr>
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
