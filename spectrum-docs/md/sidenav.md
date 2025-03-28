# sp-sidenav
Overview API
## Description
Side navigation allows users to locate information and features within the UI. It can be used for either hierarchical or flat navigation, and gives the ability to group navigable items categorically. Side navigation is an opportunity to prioritize content or features based on your users’ needs in a way that maintains clear, persistent visibility. Use side navigation within the context of larger elements and mechanisms within the app frame.
`<sp-sidenav>` elements accept both `<sp-sidenav-item>` and `<sp-sidenav-heading>` elements as children in order to construct a hierarchy of navigation elements. `<sp-sidenav-item>` elements will place themselves as a togglable child of their `<sp-sidenav>` element parent. `<sp-sidenav-heading>` elements will create visible structure by grouping their child `<sp-sidenav-item>` children under a non-interactive heading.
### Usage
    
    yarn add @spectrum-web-components/sidenav
    
Import the side effectful registration of `<sp-sidenav>`, `<sp-sidenav-heading>`, or `<sp-sidenav-item>` via:
    
    import '@spectrum-web-components/sidenav/sp-sidenav.js';
    import '@spectrum-web-components/sidenav/sp-sidenav-heading.js';
    import '@spectrum-web-components/sidenav/sp-sidenav-item.js';
    
When looking to leverage the `Sidenav`, `SidenavHeading`, or `SidenavItem` base classes as a type and/or for extension purposes, do so via:
    
    import {
        Sidenav,
        SidenavHeading,
        SidenavItem
    } from '@spectrum-web-components/sidenav';
    
## Example
    
    <sp-sidenav defaultValue="Docs">
        <sp-sidenav-item
            value="Docs"
            href="/components/SideNav"
            label="Docs"
            selected
        ></sp-sidenav-item>
        <sp-sidenav-item
            value="Guides"
            href="/guides/getting_started"
            label="Guides"
        ></sp-sidenav-item>
        <sp-sidenav-item
            value="Community"
            href="/community"
            label="Community"
        ></sp-sidenav-item>
        <sp-sidenav-item
            value="Storybook"
            href="/storybook"
            target="_blank"
            label="Storybook"
        ></sp-sidenav-item>
        <sp-sidenav-item
            value="Releases"
            href="http://git.corp.adobe.com/React/react-spectrum/releases"
            target="_blank"
            label="Releases"
            disabled
        ></sp-sidenav-item>
        <sp-sidenav-item
            value="GitHub"
            href="http://git.corp.adobe.com/React/react-spectrum"
            target="_blank"
            label="Github"
        ></sp-sidenav-item>
    </sp-sidenav>
## Multi-level
Use this variation when you have multiple layers of hierarchical navigation. The headers are styled differently and possess the same interactive behavior as a treeview; clicking the header opens and collapses the children navigation items. In the instances where a top-level navigation item has no children, clicking will send the user to the location of the item.
    
    <sp-sidenav variant="multilevel" defaultValue="Layout">
        <sp-sidenav-item value="Guidelines" label="Guidelines">
        </sp-sidenav-item>
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
        <sp-sidenav-item value="Elements" label="Elements">
        </sp-sidenav-item>
        <sp-sidenav-item value="Patterns" label="Patterns">
        </sp-sidenav-item>
    </sp-sidenav-itm>
## Icon
    
    <sp-sidenav>
        <sp-sidenav-item value="Section Title 1" label="Section Title 1">
            <sp-icon-star slot="icon"></sp-icon-star>
        </sp-sidenav-item>
        <sp-sidenav-item value="Section Title 2" label="Section Title 2">
            <sp-icon-star slot="icon"></sp-icon-star>
        </sp-sidenav-item>
        <sp-sidenav-item value="Section Title 3" label="Section Title 3">
            <sp-icon-star slot="icon"></sp-icon-star>
        </sp-sidenav-item>
    </sp-sidenav>
## Heading
    
    <sp-sidenav variant="multilevel">
        <sp-sidenav-item value="Section 1" label="Section 1"></sp-sidenav-item>
        <sp-sidenav-item value="Section 2" label="Section 2"></sp-sidenav-item>
        <sp-sidenav-heading label="Category 1">
            <sp-sidenav-item value="Section 3" label="Section 3"></sp-sidenav-item>
            <sp-sidenav-item value="Section 4" label="Section 4"></sp-sidenav-item>
        </sp-sidenav-heading>
        <sp-sidenav-heading label="Category 2">
            <sp-sidenav-item value="Section 5" label="Section 5"></sp-sidenav-item>
            <sp-sidenav-item value="Section 6" label="Section 6"></sp-sidenav-item>
        </sp-sidenav-heading>
    </sp-sidenav>
## Accessibility
When the `manage-tab-index` attribute is set on an `sp-sidenav` element then it will presents its child `sp-sidenav-item` children with a single tab-stop. This will leave items beyond the selected item, or first when there is no focusable selected item, will be accessibile via the up and down arrow keys.
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
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined | undefined`</td>
      <td>`undefined`</td>
      <td>An accessible label that describes the component, so that the side navigation can be distinguished from other navigation by screen reader users. It will be applied to aria-label, but not visually rendered.</td>
    </tr>
    <tr>
      <td>`manageTabIndex`</td>
      <td>`manage-tab-index`</td>
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
      <td>`value`</td>
      <td>`value`</td>
      <td>`string | undefined`</td>
      <td>`undefined`</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`'multilevel' | undefined`</td>
      <td>`undefined`</td>
      <td>The multilevel variant will have multiple layers of hierarchical navigation items.</td>
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
      <td>the Sidenav Items to display</td>
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
      <td>`Announces a change in the `value` property of the navigation element. This change can be "canceled" via `event.preventDefault()`.`</td>
    </tr>
  </tbody>
</table>
