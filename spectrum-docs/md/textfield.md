# sp-textfield
Overview API
## Overview
`sp-textfield` components are text boxes that allow users to input custom text entries with a keyboard. Various decorations can be displayed around the field to communicate the entry requirements.
### Usage
    
    yarn add @spectrum-web-components/textfield
    
Import the side effectful registration of `<sp-textfield>` via:
    
    import '@spectrum-web-components/textfield/sp-textfield.js';
    
When looking to leverage the `Textfield` base class as a type and/or for extension purposes, do so via:
    
    import { Textfield } from '@spectrum-web-components/textfield';
### Anatomy
    
    <sp-textfield id="basic" label="Name"></sp-textfield>
#### Label
A text field must have a label in order to be accessible. A label can be provided either via the `label` attribute, like the previous example or with an `<sp-field-label>` element.
    
    <sp-field-label for="with-field-label">Name</sp-field-label>
    <sp-textfield id="with-field-label"></sp-textfield>
#### Placeholder
Use the `placeholder` attribute to include placeholder text. **Note** : Placeholder text should not be used as a replacement for a label or help help text.
    
    <sp-field-label for="has-placeholder">Name</sp-field-label>
    <sp-textfield id="has-placeholder" placeholder="ex., John Doe"></sp-textfield>
#### Help text
Help text can be accessibly associated with an `<sp-textfield>` element by using the `help-text` or `negative-help-text` slots. When using the `negative-help-text` slot, `<sp-textfield>` will self manage the presence of this content based on the value of the `invalid` property on your `<sp-textfield>` element. Content within the `help-text` slot will be show by default. When your `<sp-textfield>` should receive help text based on state outside of the complexity of `invalid` or not, manage the content addressed to the `help-text` from above to ensure that it displays the right messaging and possesses the right `variant`.
See help text for more information.
Self managed
    
    <sp-field-label for="self">Stay "Positive"</sp-field-label>
    <sp-textfield id="self" pattern="[P][o][s][i][t][i][v][e]" value="Positive">
        <sp-help-text slot="help-text">
            Tell us how you are feeling today.
        </sp-help-text>
        <sp-help-text slot="negative-help-text">Please be "Positive".</sp-help-text>
    </sp-textfield>
Managed from above
    
    <sp-field-label for="managed">Stay "Positive"</sp-field-label>
    <sp-textfield
        id="managed"
        pattern="[P][o][s][i][t][i][v][e]"
        value="Positive"
        oninput='
            const helpText = this.querySelector(`[slot="help-text"]`);
            helpText.textContent = this.invalid ? `Please be "Positive".` : `Tell us how you are feeling today.`;
            helpText.variant = this.invalid ? `negative` : `neutral`;
        '
    >
        <sp-help-text slot="neutral-text">
            Tell us how you're feeling today.
        </sp-help-text>
        <sp-help-text slot="help-text">Please be "Positive".</sp-help-text>
    </sp-textfield>
### Options
#### Sizes
Small
    
    <sp-field-label size="s" for="name-0-s">Name</sp-field-label>
    <sp-textfield
        size="s"
        id="name-0-s"
        placeholder="Enter your name"
    ></sp-textfield>
Medium
    
    <sp-field-label for="name-0-m">Name</sp-field-label>
    <sp-textfield id="name-0-m" placeholder="Enter your name"></sp-textfield>
Large
    
    <sp-field-label size="l" for="name-0-l">Name</sp-field-label>
    <sp-textfield
        size="l"
        id="name-0-l"
        placeholder="Enter your name"
    ></sp-textfield>
Extra Large
    
    <sp-field-label size="xl" for="name-0-xl">Name</sp-field-label>
    <sp-textfield
        size="xl"
        id="name-0-xl"
        placeholder="Enter your name"
    ></sp-textfield>
#### Types
When inputting URLs, telephone numbers, email addresses, or passwords, specify a `type` to provide user affordances like mobile keyboards and obscured characters:
  -  `url`
  -  `tel`
  -  `email`
  -  `password`
  -  `text` (default)

    <sp-field-label for="tel-1">Telephone</sp-field-label>
    <sp-textfield
        id="tel-1"
        type="tel"
        placeholder="Enter your phone number"
    ></sp-textfield>
    <sp-field-label for="password-1">Password</sp-field-label>
    <sp-textfield id="password-1" type="password"></sp-textfield>
If the `type` attribute is not specified, or if it does not match any of these values, the default type adopted is "text."
#### Quiet
The quiet style works best when a clear layout (vertical stack, table, grid) assists in a user's ability to parse the element. Too many quiet components in a small space can be hard to read.
    
    <sp-field-label for="name-3">Name (quietly)</sp-field-label>
    <sp-textfield id="name-3" placeholder="Enter your name" quiet></sp-textfield>
### States
Use the `required` attribute to indicate a textfield value is required. Dictate the validity or invalidity state of the text entry with the `valid` or `invalid` attributes.
    
    <sp-field-label for="name-1" required>Name</sp-field-label>
    <sp-textfield
        id="name-1"
        placeholder="Enter your name"
        valid
        value="My Name"
    ></sp-textfield>
    <br />
    <sp-field-label for="name-2" required>Name</sp-field-label>
    <sp-textfield id="name-2" invalid placeholder="Enter your name"></sp-textfield>
### Accessibility
#### Include a label
Every text field should have a label. A field without a label is ambiguous and not accessible.
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
  </tbody>
</table>
