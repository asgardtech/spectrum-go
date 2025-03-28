# sp-action-menu
Overview API
## Overview
An `<sp-action-menu>` is an action button that triggers an overlay with `<sp-menu-items>` for activation. Use an `<sp-menu>` element to outline the items that will be made available to the user when interacting with the `<sp-action-menu>` element. By default, `<sp-action-menu>` does not manage a selection. If you'd like for a selection to be managed, use `selects="single"` on the `<sp-menu>` to activate this functionality.
### Usage
    
    yarn add @spectrum-web-components/action-menu
    
Import the side effectful registration of `<sp-action-menu>` via:
    
    import '@spectrum-web-components/action-menu/sp-action-menu.js';
    
The default of `<sp-action-menu>` will load dependencies in `@spectrum-web-components/overlay` asynchronously via a dynamic import. In the case that you would like to import those tranverse dependencies statically, import the side effectful registration of `<sp-action-menu>` as follows:
    
    import '@spectrum-web-components/action-menu/sync/sp-action-menu.js';
    
When looking to leverage the `ActionMenu` base class as a type and/or for extension purposes, do so via:
    
    import { ActionMenu } from '@spectrum-web-components/action-menu';
    
### Options
#### Sizes
Small
    
    <sp-action-menu size="s">
        <span slot="label">More Actions</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
Medium
    
    <sp-action-menu size="m">
        <span slot="label">More Actions</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
Large
    
    <sp-action-menu size="l">
        <span slot="label">More Actions</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
Extra Large
    
    <sp-action-menu size="xl">
        <span slot="label">More Actions</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
## Variants
### No icon
In order to deliver an `<sp-action-menu>` without an icon, use the `label-only` slot. This will supress any icon from being displayed, both the default ellipsis icon or any icon the user might provide to the element.
    
    <sp-action-menu>
        <span slot="label-only">More Actions</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
### No visible label
The visible label that is be provided via the default `<slot>` interface can be omitted in preference of an icon only interface. In this context be sure that the `<sp-action-menu>` continued to be accessible to screen readers by applying the `label` attribute. This will apply an `aria-label` attribute of the same value to the `<button>` element that toggles the menu list.
    
    <sp-action-menu label="More Actions">
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
### Alternate icon
A custom icon can be supplied via the `icon` slot in order to replace the default meatballs icon.
    
    <sp-action-menu>
        <sp-icon-settings slot="icon"></sp-icon-settings>
        <span slot="label">Actions under the gear</span>
        <sp-menu-item>
            Deselect
        </sp-menu-item>
        <sp-menu-item>
            Select inverse
        </sp-menu-item>
        <sp-menu-item>
            Feather...
        </sp-menu-item>
        <sp-menu-item>
            Select and mask...
        </sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>
            Save selection
        </sp-menu-item>
        <sp-menu-item disabled>
            Make work path
        </sp-menu-item>
    </sp-action-menu>
### Selection
When `selects` is set to `single`, the `<sp-action-menu>` element will maintain one selected item after an initial selection is made.
    
    <p>
        The value of the `&lt;sp-action-menu&gt;` element is:
        <span id="single-value"></span>
    </p>
    <sp-action-menu
        selects="single"
        onchange="this.previousElementSibling.querySelector('#single-value').textContent=this.value"
    >
        <span slot="label">Available shapes</span>
        <sp-menu-item value="shape-1-square">Square</sp-menu-item>
        <sp-menu-item value="shape-2-triangle">Triangle</sp-menu-item>
        <sp-menu-item value="shape-3-parallelogram">Parallelogram</sp-menu-item>
        <sp-menu-item value="shape-4-star">Star</sp-menu-item>
        <sp-menu-item value="shape-5-hexagon">Hexagon</sp-menu-item>
        <sp-menu-item value="shape-6-circle" disabled>Circle</sp-menu-item>
    </sp-action-menu>
## Force Popover on Mobile Devices
On mobile, the menu can be exposed in either a `sp-popover` or `sp-tray`. By default, `sp-action-menu` will render an `sp-tray`. If you would like to render `sp-popover` on mobile, add the attribute `force-popover` to the `sp-action-menu`.
Usage Guidance:
  -  Use a tray when a menu’s proximity to its trigger is considered to be less important to the experience, or for showing a volume of content that is too overwhelming for a popover.
  -  Use a popover when a menu’s proximity to its trigger is considered to be important to the experience, or for showing a volume of content that is manageable for a popover.

To see this functionality in action, load this page from your mobile device or use Chrome DevTools (or equivalent) and select a mobile device once the Device Toolbar (the phone/tablet icon) is active.
    
    <sp-action-menu force-popover>
        <span slot="label">Action Menu</span>
        <sp-menu-item>Deselect</sp-menu-item>
        <sp-menu-item>Select Inverse</sp-menu-item>
        <sp-menu-item>Feather...</sp-menu-item>
        <sp-menu-item>Select and Mask...</sp-menu-item>
        <sp-menu-divider></sp-menu-divider>
        <sp-menu-item>Save Selection</sp-menu-item>
        <sp-menu-item disabled>Make Work Path</sp-menu-item>
    </sp-action-menu>
## Adding tootip in action menu
Tooltip in action menu can be attached via adding `<sp-tooltip>` and can be customized by using various parameters (e.g. placement, content, etc) as needed.
    
    <sp-action-menu>
        <sp-tooltip slot="tooltip" self-managed placement="bottom">
            Content
        </sp-tooltip>
        <span slot="label">Available shapes</span>
        <sp-menu-item value="shape-1-square">Square</sp-menu-item>
        <sp-menu-item value="shape-2-triangle">Triangle</sp-menu-item>
        <sp-menu-item value="shape-3-parallelogram">Parallelogram</sp-menu-item>
    </sp-action-menu>
## Accessibility
An `<sp-action-menu>` parent will ensure that the internal `<sp-menu>` features a role of `listbox` and contains children with the role `option`. Upon focusing the `<sp-action-menu>` using `ArrowDown` will also open the menu while throwing focus into first selected (or unselected when none are selected) menu item to assist in selecting of a new value.
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
      <td>`focused`</td>
      <td>`focused`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`forcePopover`</td>
      <td>`force-popover`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Forces the Picker to render as a popover on mobile instead of a tray.</td>
    </tr>
    <tr>
      <td>`icons`</td>
      <td>`icons`</td>
      <td>`'only' | 'none' | undefined`</td>
      <td>``</td>
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
      <td>`string | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>Defines a string value that labels the Picker while it is in pending state.</td>
    </tr>
    <tr>
      <td>`placement`</td>
      <td>`placement`</td>
      <td>`"top" | "top-start" | "top-end" | "right" | "right-start" | "right-end" | "bottom" | "bottom-start" | "bottom-end" | "left" | "left-start" | "left-end"`</td>
      <td>`'bottom-start'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
      <td>`selects`</td>
      <td>`selects`</td>
      <td>`undefined | 'single'`</td>
      <td>`undefined`</td>
      <td>By default `sp-action-menu` does not manage a selection. If you'd like for a selection to be held by the `sp-menu` that it presents in its overlay, use `selects="single" to activate this functionality.</td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | 'black' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string`</td>
      <td>`''`</td>
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
      <td>menu items to be listed in the Action Menu</td>
    </tr>
    <tr>
      <td>`description`</td>
      <td>The description content for the Picker</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>The icon to use for the Action Menu</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>The label to use for the Action Menu</td>
    </tr>
    <tr>
      <td>`label-only`</td>
      <td>The label to use for the Action Menu (no icon space reserved)</td>
    </tr>
    <tr>
      <td>`tooltip`</td>
      <td>Tooltip to be applied to the Action Button</td>
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
      <td>`Announces that the `value` of the element has changed`</td>
    </tr>
    <tr>
      <td>`scroll`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
    <tr>
      <td>`sp-opened`</td>
      <td>`Event`</td>
      <td>`Announces that the overlay has been opened`</td>
    </tr>
  </tbody>
</table>
