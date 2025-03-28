# sp-divider
Overview API
## Description
`sp-divider` brings clarity to a layout by grouping and dividing content that exists in close proximity. It can also be used to establish rhythm and hierarchy.
### Usage
    
    yarn add @spectrum-web-components/divider
    
Import the side effectful registration of `<sp-divider>` via:
    
    import '@spectrum-web-components/divider/sp-divider.js';
    
When looking to leverage the `Divider` base class as a type and/or for extension purposes, do so via:
    
    import { Divider } from '@spectrum-web-components/divider';
    
## Horizontal
Small
    
    <h2 class="spectrum-Heading spectrum-Heading--sizeXS">Small</h2>
    <sp-divider size="s"></sp-divider>
    <p class="spectrum-Body">
        Divide like-elements (tables, tool groups, elements within a panel, etc.)
    </p>
Medium
    
    <h2 class="spectrum-Heading spectrum-Heading--sizeS">Medium</h2>
    <sp-divider size="m"></sp-divider>
    <p class="spectrum-Body">
        Divide subsections, or divide different groups of elements (between panels,
        rails, etc.)
    </p>
Large
    
    <h2 class="spectrum-Heading spectrum-Heading--sizeM">Large</h2>
    <sp-divider size="l"></sp-divider>
    <p class="spectrum-Body">Page or Section Titles.</p>
## Vertical
When a vertical Divider is used inside of a flex container, use `align-self: stretch; height: auto;` on the Divider.
Small
    
    <div style="height: 32px; display: flex;">
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-divider
            size="s"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </div>
Medium
    
    <div style="height: 32px; display: flex;">
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-divider
            size="m"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </div>
Large
    
    <div style="height: 32px; display: flex;">
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-divider
            size="l"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-button quiet label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </div>
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
      <td>`vertical`</td>
      <td>`vertical`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
  </tbody>
</table>
