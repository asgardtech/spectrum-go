# sp-action-bar
Overview API
## Overview
A `<sp-action-bar>` delivers a floating action bar that is a convenient way to deliver stateful actions in cases like selection mode. `<sp-action-bar>` can be deployed in two variants beyond the default: `[varient="fixed"]` to position the element in relation to the page, and `[variant=sticky]` to position the content in relation to content that may scroll.
### Usage
    
    yarn add @spectrum-web-components/action-bar
    
Import the side effectful registration of `<sp-action-bar>` via:
    
    import '@spectrum-web-components/action-bar/sp-action-bar.js';
    
When looking to leverage the `ActionBar` base class as a type and/or for extension purposes, do so via:
    
    import { ActionBar } from '@spectrum-web-components/action-bar';
    
## Example
    
    <sp-action-bar open>
        2 selected
        <sp-action-button slot="buttons" label="Edit">
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button slot="buttons" label="More">
            <sp-icon-more slot="icon"></sp-icon-more>
        </sp-action-button>
    </sp-action-bar>
## Emphasized
Use the `emphasized` attribute to add priority to the information that is delivered within your `<sp-action-bar>` element:
    
    <sp-action-bar emphasized open>
        2 selected
        <sp-action-button slot="buttons" label="Edit">
            <sp-icon-edit slot="icon"></sp-icon-edit>
        </sp-action-button>
        <sp-action-button slot="buttons" label="More">
            <sp-icon-more slot="icon"></sp-icon-more>
        </sp-action-button>
    </sp-action-bar>
## Variants
### Fixed
When using `[variant="fixed"]`, the `<sp-action-bar>` will display by default at the bottom left of the window and can be customized via CSS from the outside.
    
    <h4>Look down and to the left when toggling.</h4>
    <sp-button
        onclick="javascript:this.nextElementSibling.open = !this.nextElementSibling.open;"
    >
        Toggle fixed action bar
    </sp-button>
    <sp-action-bar variant="fixed">2 selected</sp-action-bar>
### Sticky
When using `[variant="sticky"]`, be sure you've spent some time touching up on how `sticky` really works to ensure the most successful delivery of your content.
    
    <section style="position: relative; max-height: 6em; overflow: auto;">
        <h4>Scroll down for toggle button</h4>
        <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
            tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
            veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
            commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
            velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
            occaecat cupidatat non proident, sunt in culpa qui officia deserunt
            mollit anim id est laborum.
        </p>
        <sp-button
            onclick="javascript:this.nextElementSibling.open = !this.nextElementSibling.open;"
        >
            Toggle sticky action bar
        </sp-button>
        <sp-action-bar variant="sticky" style="inset-block: 0px">
            2 selected
            <sp-action-button slot="buttons" label="Edit">
                <sp-icon-edit slot="icon"></sp-icon-edit>
            </sp-action-button>
            <sp-action-button slot="buttons" label="More">
                <sp-icon-more slot="icon"></sp-icon-more>
            </sp-action-button>
        </sp-action-bar>
    </section>
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
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Deliver the Action Bar with additional visual emphasis.</td>
    </tr>
    <tr>
      <td>`flexible`</td>
      <td>`flexible`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>When `flexible` the action bar sizes itself to its content rather than a specific width.</td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`string`</td>
      <td>``</td>
      <td>The variant applies specific styling when set to `sticky` or `fixed`. `variant` attribute is removed when not matching one of the above.</td>
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
      <td>Content to display with the Action Bar</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`close`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
