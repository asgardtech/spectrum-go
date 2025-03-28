# sp-switch
Overview API
## Overview
An `<sp-switch>` is used to turn an option on or off. Switches allow users to select the state of a single option at a time. Use a switch rather than a checkbox when activating (or deactivating) an option, instead of selecting.
### Usage
    
    yarn add @spectrum-web-components/switch
    
Import the side effectful registration of `<sp-switch>` via:
    
    import '@spectrum-web-components/switch/sp-switch.js';
    
When looking to leverage the `Switch` base class as a type and/or for extension purposes, do so via:
    
    import { Switch } from '@spectrum-web-components/switch';
    
### Anatomy
A switch consists of a switch input and slotted label.
    
    <sp-switch>Email notifications</sp-switch>
#### Checked
A switch can be checked by setting the `checked` property/attribute.
    
    <sp-field-group vertical>
        <sp-switch>Not checked</sp-switch>
        <sp-switch checked>Checked</sp-switch>
    </sp-field-group>
### Options
#### Sizes
Small
    
    <sp-switch size="s">Small</sp-switch>
Medium
    
    <sp-switch size="m">Medium</sp-switch>
Large
    
    <sp-switch size="l">Large</sp-switch>
Extra Large
    
    <sp-switch size="xl">Extra Large</sp-switch>
#### Emphasized
Emphasized switches, which use the `empahasized` attribute/property are a secondary style for switches. The blue color provides a visual prominence that is optimal for forms, settings, etc. where the switches need to be noticed.
    
    <sp-field-group vertical>
        <sp-switch emphasized>Emphasized</sp-switch>
        <sp-switch emphasized checked>Emphasized and checked</sp-switch>
    </sp-field-group>
### States
A switch can be disabled using the `disabled` property/attribute.
    
    <sp-field-group vertical>
        <sp-switch disabled>Disabled</sp-switch>
        <sp-switch disabled checked>Disabled and checked</sp-switch>
    </sp-field-group>
### Behaviors
#### Handling events
Event handlers for clicks and other user actions can be registered on an `<sp-switch>` similar to a standard `<input type="checkbox">` element.
    
    <sp-switch id="switch-example" onclick="spAlert(this, '<sp-switch> clicked!')">
        Web component
    </sp-switch>
### Accessibility
Switch are rendered in HTML using the `<input type="checkbox">` element with the appropriate accessibility role, `switch`. When the Switch is `checked`, the appropriate ARIA state attribute will automatically be applied.
#### Include a label
A switch is required to have either a visible text label nested inside `<sp-switch>` itself.
    
    <sp-switch>Email notifications</sp-switch>
Standalone switches should be used in situations where the context is clear without an associated text label. For example, a switch located at the top of a panel next to the panel's title makes it clear that the switch will enable/disable the panel options.
In those cases, you can use CSS to visually hide the text label.
    
    <div id="settings">
        <sp-field-label for="notifications-settings">Notifications</sp-field-label>
        <sp-switch id="notify">
            <span class="visually-hidden">Notifications</span>
        </sp-switch>
        <sp-field-group id="notifications-settings" vertical>
            <sp-switch disabled>Email</sp-switch>
            <sp-switch disabled>Telephone</sp-switch>
            <sp-switch disabled>Text</sp-switch>
        </sp-field-group>
    </div>
    
    <style>
        .visually-hidden {
            clip: rect(0 0 0 0);
            clip-path: inset(50%);
            height: 1px;
            overflow: hidden;
            position: absolute;
            white-space: nowrap;
            width: 1px;
        }
        #settings {
            display: grid;
            grid-gap: 10px;
            grid-template-columns: calc(100% - 50px) 50px;
        }
        #notifications-settings {
            grid-column: 1 / 3;
            grid-row: 2;
        }
    </style>
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
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
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
      <td>text label of the Switch</td>
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
      <td>`Announces a change in the `checked` property of a Switch`</td>
    </tr>
  </tbody>
</table>
