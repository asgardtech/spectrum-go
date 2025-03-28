# textarea
Overview API
## Description
`sp-textfield[multiline]` components are text areas that allow users to input custom multiline text entries with a keyboard. Various decorations can be displayed around the field to communicate the entry requirements.
### Usage
    
    yarn add @spectrum-web-components/textfield
    
Import the side effectful registration of `<sp-textfield>` via:
    
    import '@spectrum-web-components/textfield/sp-textfield.js';
    
When looking to leverage the `Textfield` base class as a type and/or for extension purposes, do so via:
    
    import { Textfield } from '@spectrum-web-components/textfield';
    
## Sizes
Small
    
    <sp-field-label size="s" for="story-0-s">Background</sp-field-label>
    <sp-textfield
        size="s"
        id="story-0-s"
        multiline
        placeholder="Enter your life story"
    ></sp-textfield>
Medium
    
    <sp-field-label for="story-0-m">Background</sp-field-label>
    <sp-textfield
        id="story-0-m"
        multiline
        placeholder="Enter your life story"
    ></sp-textfield>
Large
    
    <sp-field-label size="l" for="story-0-l">Background</sp-field-label>
    <sp-textfield
        size="l"
        id="story-0-l"
        multiline
        placeholder="Enter your life story"
    ></sp-textfield>
Extra Large
    
    <sp-field-label size="xl" for="story-0-xl">Background</sp-field-label>
    <sp-textfield
        size="xl"
        id="story-0-xl"
        multiline
        placeholder="Enter your life story"
    ></sp-textfield>
## Variants
### Valid
Dictate the validity state of the text entry with the `valid` attribute.
    
    <sp-field-label for="story-1" required>Background</sp-field-label>
    <sp-textfield
        id="story-1"
        multiline
        placeholder="Enter your name"
        valid
    ></sp-textfield>
### Invalid
Dictate the invalidity state of the text entry with the `invalid` attribute.
    
    <sp-field-label for="story-2" required>Background</sp-field-label>
    <sp-textfield
        id="story-2"
        invalid
        multiline
        placeholder="Enter your name"
    ></sp-textfield>
### Quiet
The quiet style works best when a clear layout (vertical stack, table, grid) assists in a user's ability to parse the element. Too many quiet components in a small space can be hard to read.
    
    <sp-field-label for="story-3">Background (quietly)</sp-field-label>
    <sp-textfield
        id="story-3"
        multiline
        placeholder="Enter your name"
        quiet
    ></sp-textfield>
### Grows
By default the text area has a fixed height and will scroll when text entry goes beyond the available space. With the use of the `grows` attribute the text area will grow to accommodate the full content of the element.
    
    <div style="display: flex; flex-wrap: wrap;">
        <div>
            <sp-field-label for="story-4">Background</sp-field-label>
            <sp-textfield
                id="story-4"
                multiline
                placeholder="Enter your name"
                value="By default the text area has a fixed height and will scroll when text entry goes beyond the available space. With the use of the `grows` attribute the text area will grow to accommodate the full content of the element."
            ></sp-textfield>
        </div>
        <div>
            <sp-field-label for="story-5">Background</sp-field-label>
            <sp-textfield
                id="story-5"
                grows
                multiline
                placeholder="Enter your name"
                value="By default the text area has a fixed height and will scroll when text entry goes beyond the available space. With the use of the `grows` attribute the text area will grow to accommodate the full content of the element."
            ></sp-textfield>
        </div>
        <div>
            <sp-field-label for="story-6">Background (quietly)</sp-field-label>
            <sp-textfield
                id="story-6"
                grows
                multiline
                placeholder="Enter your name"
                value="By default the text area has a fixed height and will scroll when text entry goes beyond the available space. With the use of the `grows` attribute the text area will grow to accommodate the full content of the element."
                quiet
            ></sp-textfield>
        </div>
    </div>
## Help text
Help text can be accessibly associated with an `<sp-textfield multiline>` element by using the `help-text` or `negative-help-text` slots. When using the `negative-help-text` slot, `<sp-textfield multiline>` will self manage the presence of this content based on the value of the `invalid` property on your `<sp-textfield multiline>` element. Content within the `help-text` slot will be show by default. When your `<sp-textfield multiline>` should receive help text based on state outside of the complexity of `invalid` or not, manage the content addressed to the `help-text` from above to ensure that it displays the right messaging and possesses the right `variant`.
Self managed
    
    <sp-field-label for="self">Stay "Positive"</sp-field-label>
    <sp-textfield
        multiline
        id="self"
        pattern="[P][o][s][i][t][i][v][e]"
        value="Positive"
    >
        <sp-help-text slot="help-text">
            Tell us how you are feeling today.
        </sp-help-text>
        <sp-help-text slot="negative-help-text">Please be "Positive".</sp-help-text>
    </sp-textfield>
Managed from above
    
    <sp-field-label for="managed">Stay "Positive"</sp-field-label>
    <sp-textfield
        multiline
        id="managed"
        pattern="[P][o][s][i][t][i][v][e]"
        value="Positive"
        oninput='
            const helpText = this.querySelector(`[slot="help-text"]`);
            helpText.textContent = this.invalid ? `Please be "Positive".` : `Tell us how you are feeling today.`;
            helpText.variant = this.invalid ? `negative` : `neutral`;
        '
    >
        <sp-help-text slot="help-text">
            Tell us how you're feeling today.
        </sp-help-text>
    </sp-textfield>
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
