# sp-top-nav
Overview API
## Description
`<sp-top-nav>` delivers site navigation, particularly for when that navigation will change the majority of the page's content and/or the page's URL when selected. All primary elements of an `<sp-top-nav>` should be directly accessible in the tab order.
### Usage
    
    yarn add @spectrum-web-components/top-nav
    
Import the side effectful registration of `<sp-top-nav>` and `<sp-top-nav-item>` as follows:
    
    import '@spectrum-web-components/top-nav/sp-top-nav.js';
    import '@spectrum-web-components/top-nav/sp-top-nav-item.js';
    
When looking to leverage the `TopNav` or `TopNavItem` base classes as a type and/or for extension purposes, do so via:
    
    import { TopNav, TopNavItem } from '@spectrum-web-components/top-nav';
    
### Example
    
    <sp-top-nav>
        <sp-top-nav-item href="#">Site Name</sp-top-nav-item>
        <sp-top-nav-item href="#page-1" style="margin-inline-start: auto;">
            Page 1
        </sp-top-nav-item>
        <sp-top-nav-item href="#page-2">Page 2</sp-top-nav-item>
        <sp-top-nav-item href="#page-3">Page 3</sp-top-nav-item>
        <sp-top-nav-item href="#page-4">Page with Really Long Name</sp-top-nav-item>
        <sp-action-menu
            label="Account"
            placement="bottom-end"
            style="margin-inline-start: auto;"
            quiet
        >
            <sp-menu-item>Account Settings</sp-menu-item>
            <sp-menu-item>My Profile</sp-menu-item>
            <sp-menu-divider></sp-menu-divider>
            <sp-menu-item>Share</sp-menu-item>
            <sp-menu-divider></sp-menu-divider>
            <sp-menu-item>Help</sp-menu-item>
            <sp-menu-item>Sign Out</sp-menu-item>
        </sp-action-menu>
    </sp-top-nav>
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
      <td>`ignoreURLParts`</td>
      <td>`ignore-url-parts`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td>A space separated list of part of the URL to ignore when matching for the "selected" Top Nav Item. Currently supported values are `hash` and `search`, which will remove the `#hash` and `?search=value` respectively.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>The Top Nav is displayed without a border.</td>
    </tr>
    <tr>
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`selectionIndicatorStyle`</td>
      <td>`selectionIndicatorStyle`</td>
      <td>``</td>
      <td>`noSelectionStyle`</td>
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
      <td>Nav Items to display as a group</td>
    </tr>
  </tbody>
</table>
