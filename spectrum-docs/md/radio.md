# sp-radio
Overview API
## Overview
`<sp-radio>` and `<sp-radio-group>` allow users to select a single option from a list of mutually exclusive options. All possible options are exposed up front for users to compare.
### Usage
    
    yarn add @spectrum-web-components/radio
Import the side effectful registration of `<sp-radio>` or `<sp-radio-group>` via:
    
    import '@spectrum-web-components/radio/sp-radio.js';
    import '@spectrum-web-components/radio/sp-radio-group.js';
When looking to leverage the `Radio` or `RadioGroup` base classes as a type and/or for extension purposes, do so via:
    
    import { Radio, RadioGroup } from '@spectrum-web-components/radio';
### Anatomy
`<sp-radio-group>` holds a list of `<sp-radio>` elements, and is responsible for deselecting radio buttons when a new one is selected, which in turn makes it responsible for keeping track of which one is selected. `<sp-radio>` is responsible for handling user interactions and for visually reflecting if it is the one that is checked or not.
    
    <sp-radio-group label="Choose an option" name="anatomy">
        <sp-radio value="first">Option 1</sp-radio>
        <sp-radio value="second">Option 2</sp-radio>
        <sp-radio value="third">Option 3</sp-radio>
        <sp-radio value="fourth">Option 4</sp-radio>
    </sp-radio-group>
#### Label
The `<sp-radio>` elements are labelled with text in their default slot.
    
    <sp-radio-group label="Choose an option" name="anatomy">
        <sp-radio value="first">Option 1</sp-radio>
        <sp-radio value="second">Option 2</sp-radio>
        <sp-radio value="third">Option 3</sp-radio>
        <sp-radio value="fourth">Option 4</sp-radio>
    </sp-radio-group>
### Options
#### Sizes
Small
    
    <sp-radio-group label="Small" selected="first" name="small">
        <sp-radio value="first" size="s">Option 1</sp-radio>
        <sp-radio value="second" size="s">Option 2</sp-radio>
        <sp-radio value="third" size="s">Option 3</sp-radio>
        <sp-radio value="fourth" size="s">Option 4</sp-radio>
    </sp-radio-group>
Medium
    
    <sp-radio-group label="Medium" selected="first" name="medium">
        <sp-radio value="first" size="m">Option 1</sp-radio>
        <sp-radio value="second" size="m">Option 2</sp-radio>
        <sp-radio value="third" size="m">Option 3</sp-radio>
        <sp-radio value="fourth" size="m">Option 4</sp-radio>
    </sp-radio-group>
Large
    
    <sp-radio-group label="Large" selected="first" name="large">
        <sp-radio value="first" size="l">Option 1</sp-radio>
        <sp-radio value="second" size="l">Option 2</sp-radio>
        <sp-radio value="third" size="l">Option 3</sp-radio>
        <sp-radio value="fourth" size="l">Option 4</sp-radio>
    </sp-radio-group>
Extra Large
    
    <sp-radio-group label="Extra large" selected="first" name="extra-large">
        <sp-radio value="first" size="xl">Option 1</sp-radio>
        <sp-radio value="second" size="xl">Option 2</sp-radio>
        <sp-radio value="third" size="xl">Option 3</sp-radio>
        <sp-radio value="fourth" size="xl">Option 4</sp-radio>
    </sp-radio-group>
#### Styles
Standard radio buttons are the default style for radio buttons. They are optimal for application panels where all visual elements are monochrome in order to direct focus to the content.
**Emphasized** radio buttons are a secondary style for radio buttons. The blue color provides a visual prominence that is optimal for forms, settings, etc. where the radio buttons need to be noticed.
Standard
    
    <div style="display: flex; justify-content: space-between;">
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-1" size="l">
                <h4 class="spectrum-Heading--subtitle1">Default</h4>
            </sp-field-label>
            <sp-radio-group id="example-1" name="example" vertical>
                <sp-radio value="kittens">Kittens</sp-radio>
                <sp-radio value="puppies" checked>Puppies</sp-radio>
            </sp-radio-group>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-2" size="l">
                <h4 class="spectrum-Heading--subtitle1">Invalid</h4>
            </sp-field-label>
            <sp-radio-group invalid id="example-2" name="example" vertical>
                <sp-radio invalid value="kittens">Kittens</sp-radio>
                <sp-radio invalid value="puppies" checked>Puppies</sp-radio>
                 <sp-help-text slot="negative-help-text" icon>
                    This selection is invalid.
                </sp-help-text>
            </sp-radio-group>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-3" size="l">
                <h4 class="spectrum-Heading--subtitle1">Disabled</h4>
            </sp-fieldlabel>
            <sp-radio-group id="example-3" name="example" vertical>
                <sp-radio disabled value="kittens">Kittens</sp-radio>
                <sp-radio disabled value="puppies" checked>Puppies</sp-radio>
            </sp-radio-group>
        </div>
    </div>
Emphasized
    
    <div style="display: flex; justify-content: space-between;">
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-a" size="l">
                <h4 class="spectrum-Heading--subtitle1">Default</h4>
            </sp-field-label>
            <sp-radio-group id="example-a" name="example" vertical>
                <sp-radio emphasized value="kittens">Kittens</sp-radio>
                <sp-radio emphasized value="puppies" checked>Puppies</sp-radio>
            </sp-radio-group>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-b" size="l">
                <h4 class="spectrum-Heading--subtitle1">Invalid</h4>
            </sp-field-label>
            <sp-radio-group invalid id="example-b" name="example" vertical>
                <sp-radio emphasized invalid value="kittens">Kittens</sp-radio>
                <sp-radio emphasized invalid value="puppies" checked>Puppies</sp-radio>
                <sp-help-text slot="negative-help-text" icon>
                    This selection is invalid.
                </sp-help-text>
            </sp-radio-group>
        </div>
    
        <div style="display: flex; flex-direction: column;">
            <sp-field-label for="example-c" size="l">
                <h4 class="spectrum-Heading--subtitle1">Disabled</h4>
            </sp-fieldlabel>
            <sp-radio-group id="example-c" name="example" vertical>
                <sp-radio emphasized disabled value="kittens">Kittens</sp-radio>
                <sp-radio emphasized disabled value="puppies" checked>Puppies</sp-radio>
            </sp-radio-group>
        </div>
    </div>
### Behaviors
#### Handling events
Event handlers for clicks and other user actions can be registered on an `<sp-radio>` similar to a standard `<input type="radio">` element.
    
    <sp-radio id="radio-example" onclick="spAlert(this, '<sp-radio> clicked!')">
        Web component
    </sp-radio>
### Accessibility
Tabbing into a group of radio buttons places the focus on the first radio button selected. If none of the radio buttons are selected, the focus is set on the first one in the group. Space selects the radio button in focus (if not already selected). Using the arrow keys moves focus and selection to the previous or next radio button in the group (last becomes first, and first becomes last). The new radio button in focus gets selected even if the previous one was not.
#### Provide a label
Radio groups and radio items should always have labels.
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
      <td>Represents when the input is checked</td>
    </tr>
    <tr>
      <td>`disabled`</td>
      <td>`disabled`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Uses the disabled style</td>
    </tr>
    <tr>
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`invalid`</td>
      <td>`invalid`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Uses the invalid style</td>
    </tr>
    <tr>
      <td>`readonly`</td>
      <td>`readonly`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td>Identifies this radio button within its radio group</td>
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
      <td>text label of the Radio button</td>
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
      <td>`When the input is interacted with and its state is changed`</td>
    </tr>
  </tbody>
</table>
