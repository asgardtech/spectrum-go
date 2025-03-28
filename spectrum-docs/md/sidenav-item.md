# sp-sidenav-item
Overview API
## Description
An `<sp-sidenav-item>` stands as both a child item of an `<sp-sidenav>` element, as well as a parent for further subdivisions of that navigation. An `<sp-sidenav-item>` with further `<sp-sidenav-item>` children will count as a toggle for the visibility of this next level of navigation items, while also updating the parent `<sp-sidenav>` element to its value when selected.
### Usage
    
    yarn add @spectrum-web-components/sidenav
    
Import the side effectful registration of `<sp-sidenav-item>` via:
    
    import '@spectrum-web-components/sidenav/sp-sidenav-item.js';
    
When looking to leverage the `SidenavItem` base classes as a type and/or for extension purposes, do so via:
    
    import { SidenavItem } from '@spectrum-web-components/sidenav';
    
## Example
    
    <sp-sidenav>
        <sp-sidenav-item
            value="Docs"
            label="Docs"
            href="/components/sidenav"
        ></sp-sidenav-item>
    </sp-sidenav>
## Multi-level
    
    <sp-sidenav variant="multilevel">
        <sp-sidenav-item value="Styles" label="Styles" expanded>
            <sp-sidenav-item value="Color" label="Color">
            </sp-sidenav-item>
            <sp-sidenav-item value="Grid" label="Grid" expanded>
                <sp-sidenav-item value="Layout" label="Layout">
                </sp-sidenav-item>
                <sp-sidenav-item value="Responsive" label="Responsive">
                </sp-sidenav-item>
            </sp-sidenav-item>
            <sp-sidenav-item value="Typography" label="Typography">
            </sp-sidenav-item>
        </sp-sidenav-item>
    </sp-sidenav-itm>
## Icon
    
    <sp-sidenav>
        <sp-sidenav-item value="Section Title 1" label="Section Title 1">
            <sp-icon-star slot="icon"></sp-icon-star>
        </sp-sidenav-item>
    </sp-sidenav>
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
      <td>`expanded`</td>
      <td>`expanded`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>`selected`</td>
      <td>`selected`</td>
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
    <tr>
      <td>`target`</td>
      <td>`target`</td>
      <td>`'_blank' | '_parent' | '_self' | '_top' | undefined`</td>
      <td>``</td>
      <td>Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).</td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string | undefined`</td>
      <td>`undefined`</td>
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
      <td>the Sidenav Items to display as children of this item</td>
    </tr>
  </tbody>
</table>
