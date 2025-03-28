# sp-popover
Overview API
## Overview
An `<sp-popover>` is used to display transient content (menus, options, additional actions etc.) and appears when clicking/tapping on a source (tools, buttons, etc.) It stands out via its visual style (stroke and drop shadow) and floats on top of the rest of the interface. This component does not implement the actual overlay behavior and interactions. This is handled by the `Overlay` system.
### Usage
    
    yarn add @spectrum-web-components/popover
    
Import the side effectful registration of `<sp-popover>` via:
    
    import '@spectrum-web-components/popover/sp-popover.js';
    
When looking to leverage the `Popover` base class as a type and/or for extension purposes, do so via:
    
    import { Popover } from '@spectrum-web-components/popover';
    
## Anatomy
    
    <div
        style="
            position: relative;
            height: 100px;
        "
    >
        <sp-popover open>
            Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
        </sp-popover>
    </div>
### Options
#### Default with no tip
Default popover with no tip and no placement. Popovers will fill up the space of their containing element by default. The default popover has no padding.
    
    <div
        style="
            position: relative;
            height: 180px;
            max-width: 320px;
        "
    >
        <sp-popover variant="default" open>
            <h2>Popover title</h2>
            <p>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </p>
        </sp-popover>
    </div>
#### Dialog popovers
To apply a managed amount of padding within your `<sp-popover>`, you may choose to wrap your slotted content in an `<sp-dialog>` element, as seen below:
    
    <div
        style="
            position: relative;
            height: 250px;
            max-width: 320px;
        "
    >
        <sp-popover open>
            <sp-dialog>
                <h3 slot="heading">Popover title</h3>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </sp-dialog>
        </sp-popover>
    </div>
#### Popover with tip
The `placement` attribute can be used to customize how the `<sp-popover>` points to its related content. `placement="top"` will point down to the related content from the top, etc.
Top
    
    <div
        style="
            position: relative;
            height: 250px;
            max-width: 320px;
        "
    >
        <sp-popover placement="top" tip open>
            <sp-dialog>
                <h3 slot="heading">Popover title</h3>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </sp-dialog>
        </sp-popover>
    </div>
Right
    
    <div
        style="
            position: relative;
            height: 200px;
            max-width: 320px;
        "
    >
        <sp-popover placement="right" tip open>
            <sp-dialog>
                <h3 slot="heading">Popover title</h3>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </sp-dialog>
        </sp-popover>
    </div>
Bottom
    
    <div
        style="
            position: relative;
            height: 200px;
            max-width: 320px;
        "
    >
        <sp-popover placement="bottom" tip open>
            <sp-dialog>
                <h3 slot="heading">Popover title</h3>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </sp-dialog>
        </sp-popover>
    </div>
Left
    
    <div
        style="
            position: relative;
            height: 200px;
            max-width: 320px;
        "
    >
        <sp-popover placement="left" tip open>
            <sp-dialog>
                <h3 slot="heading">Popover title</h3>
                Cupcake ipsum dolor sit amet jelly beans. Chocolate jelly caramels.
                Icing soufflé chupa chups donut cheesecake. Jelly-o chocolate cake
                sweet roll cake danish candy biscuit halvah
            </sp-dialog>
        </sp-popover>
    </div>
### Accessibility
For components used with a popover, see the accessibility guidelines of the particular component.
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
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the popover is visible or not.</td>
    </tr>
    <tr>
      <td>`placement`</td>
      <td>`placement`</td>
      <td>`"top" | "top-start" | "top-end" | "right" | "right-start" | "right-end" | "bottom" | "bottom-start" | "bottom-end" | "left" | "left-start" | "left-end"`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`tip`</td>
      <td>`tip`</td>
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
      <td>content to display within the Popover</td>
    </tr>
  </tbody>
</table>
