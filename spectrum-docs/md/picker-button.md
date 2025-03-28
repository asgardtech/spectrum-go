# sp-picker-button
Overview API
## Overview
An `<sp-picker-button>` is used as a sub-component of patterns like the `<sp-combobox>` (release pending) to pair a button interface with a text input. With its many custom states and alterations, it isn't likely to be leveraged directly outside of more complex UIs.
### Usage
    
    yarn add @spectrum-web-components/picker-button
    
Import the side effectful registration of `<sp-picker-button>` via:
    
    import '@spectrum-web-components/picker-button/sp-picker-button.js';
    
When looking to leverage the `PickerButton` base class as a type and/or for extension purposes, do so via:
    
    import { PickerButton } from '@spectrum-web-components/picker-button';
    
### Anatomy
#### Text and icon
With the use of the `label` slot, you can deliver an `<sp-picker-button>` with both an icon and a text label:
    
    <sp-picker-button><span slot="label">All</span></sp-picker-button>
#### Icon only
Without content addressed to the `label` slot, the `<sp-picker-button>` with both an icon and a text label:
    
    <sp-picker-button></sp-picker-button>
#### Custom icon
You can customize the icon in an `<sp-picker-button>` by addressing a new icon to the `icon` slot:
    
    <sp-picker-button><sp-icon-add slot="icon"></sp-icon-add></sp-picker-button>
### Options
#### Sizes
Small
    
    <sp-picker-button size="s"></sp-picker-button>
    <br />
    <sp-picker-button size="s"><span slot="label">All</span></sp-picker-button>
Medium
    
    <sp-picker-button size="m"></sp-picker-button>
    <br />
    <sp-picker-button size="m"><span slot="label">All</span></sp-picker-button>
Large
    
    <sp-picker-button size="l"></sp-picker-button>
    <br />
    <sp-picker-button size="l"><span slot="label">All</span></sp-picker-button>
Extra Large
    
    <sp-picker-button size="xl"></sp-picker-button>
    <br />
    <sp-picker-button size="xl"><span slot="label">All</span></sp-picker-button>
#### Rounded
When delivered as part of the `express` system, an `<sp-picker-button>` with the `rounded` attribute will be given deeply rounded corners:
    
    <sp-picker-button rounded></sp-picker-button>
    <br />
    <sp-picker-button rounded><span slot="label">All</span></sp-picker-button>
#### Quiet
When delivered with the `quiet` attribute, the `<sp-picker-button>` will take a less pronounced visual delivery:
    
    <sp-picker-button quiet></sp-picker-button>
    <br />
    <sp-picker-button quiet><span slot="label">All</span></sp-picker-button>
#### Position
By default the `<sp-picker-button>` will be given a `position` attribute with the value `right`, which is best leveraged at the right edge of an associated `<sp-textfield>` element. If your UI desires that the `<sp-picker-button>` be placed to the left of the related input, use the `position` attribute and set it to `left`:
    
    <sp-picker-button position="left"></sp-picker-button>
    <br />
    <sp-picker-button position="left">
        <span slot="label">All</span>
    </sp-picker-button>
### States
#### Open
When paired with an expanded UI, e.g. an `<sp-combobox>` with its autocomplete options visible, an `<sp-picker-button>` should be given the `open` attribute to visual match the delivered state in the larger UI:
    
    <sp-picker-button open></sp-picker-button>
    <br />
    <sp-picker-button open><span slot="label">All</span></sp-picker-button>
#### Disabled
Leveraging the `disabled` attribute will dim the `<sp-picker-button>` and alter its presentation in the accessibility tree:
    
    <sp-picker-button disabled></sp-picker-button>
    <br />
    <sp-picker-button disabled><span slot="label">All</span></sp-picker-button>
#### Invalid
When delivered as part of the `spectrum` system, an `<sp-picker-button>` with the `invalid` attribute will be given a red border:
    
    <sp-picker-button invalid></sp-picker-button>
    <br />
    <sp-picker-button invalid><span slot="label">All</span></sp-picker-button>
### Accessibility
The example below is for demonstration purposes. For an example implementation of `<sp-picker-button>` view `Combobox.ts`. For comprehensive information on menu button accessibility, see WAI ARIA Authoring Practice Guide's Menu Button Pattern.
    
    <sp-field-label for="color">Color</sp-field-label>
    <sp-textfield id="color"></sp-textfield>
    <overlay-trigger type="modal">
        <sp-picker-button
            aria-controls="colors-menu"
            aria-expanded="false"
            aria-haspopup="menu"
            aria-describedby="color"
            slot="trigger"
        ></sp-picker-button>
        <sp-tray slot="click-content">
            <sp-menu id="colors-menu">
                <sp-menu-item>Red</sp-menu-item>
                <sp-menu-item>Blue</sp-menu-item>
            </sp-menu>
        </sp-tray>
    </overlay-trigger>
#### Include a label
To ensure menu items can be read by assistive technology, do *one* of the following:
  -  Place visible text in the component's `label` slot.
  -  Use `aria-label` attribute.
  -  Set the `aria-labelledby` attribute to the ID reference of the menu element.

#### Set aria properties correctly
To indicate to assistive technology what the button does, do *all* of the following:
  -  Set the `aria-controls` property to the ID reference of the menu element.
  -  Set the `aria-haspopup` property to `"menu"` or `"true"`.
  -  Set the `aria-expanded` property to `"menu"` or `"true"` or `"false"` depending on whether the menu is displayed.

### Add keyboard interaction
Ensure that picker button can be operated by keyboard users:
  -  Required: Open the menu and focus on first menu item when `Enter` or `Space` is pressed.
  -  *Optional* : Open the menu and focus on first menu item when `Down Arrow` is pressed.
  -  *Optional* : Open the menu and focus on last menu item when `Up Arrow` is pressed.

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
      <td>`active`</td>
      <td>`active`</td>
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
      <td>`download`</td>
      <td>`download`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Causes the browser to treat the linked URL as a download.</td>
    </tr>
    <tr>
      <td>`href`</td>
      <td>`href`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The URL that the hyperlink points to.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.</td>
    </tr>
    <tr>
      <td>`referrerpolicy`</td>
      <td>`referrerpolicy`</td>
      <td>`| 'no-referrer' | 'no-referrer-when-downgrade' | 'origin' | 'origin-when-cross-origin' | 'same-origin' | 'strict-origin' | 'strict-origin-when-cross-origin' | 'unsafe-url' | undefined`</td>
      <td>``</td>
      <td>How much of the referrer to send when following the link.</td>
    </tr>
    <tr>
      <td>`rel`</td>
      <td>`rel`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The relationship of the linked URL as space-separated link types.</td>
    </tr>
    <tr>
      <td>`tabIndex`</td>
      <td>`tabIndex`</td>
      <td>`number`</td>
      <td>``</td>
      <td>The tab index to apply to this control. See general documentation about the tabindex HTML property</td>
    </tr>
    <tr>
      <td>`target`</td>
      <td>`target`</td>
      <td>`'_blank' | '_parent' | '_self' | '_top' | undefined`</td>
      <td>``</td>
      <td>Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).</td>
    </tr>
    <tr>
      <td>`type`</td>
      <td>`type`</td>
      <td>`'button' | 'submit' | 'reset'`</td>
      <td>`'button'`</td>
      <td>The default behavior of the button. Possible values are: `button` (default), `submit`, and `reset`.</td>
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
      <td>text content to be displayed in the Button element</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>icon element(s) to display at the start of the button</td>
    </tr>
  </tbody>
</table>
