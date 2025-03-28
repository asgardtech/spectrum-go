# sp-breadcrumbs
Overview API
## Description
### Usage
    
    yarn add @spectrum-web-components/breadcrumbs
    
Import the side effectful registration of `<sp-breadcrumbs>` and `<sp-breadcrumb-item>` via:
    
    import '@spectrum-web-components/breadcrumbs/sp-breadcrumbs.js';
    import '@spectrum-web-components/breadcrumbs/sp-breadcrumb-item.js';
    
When looking to leverage the `Breadcrumbs` or `BreadcrumbItem` base class as a type and/or for extension purposes, do so via:
    
    import { Breadcrumbs, BreadcrumbItem } from '@spectrum-web-components/breadcrumbs';
    
## Example
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item value="1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item value="2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Disabled
`sp-breadcrumb-item` can have a `disabled` state which disables the events from the disabled item.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item disabled value="1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item value="2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Compact
When needing to optimize for functional space of `sp-breadcrumbs`, the compact option is useful for reducing the height of the breadcrumbs while still maintaining the proper user context.
    
    <sp-breadcrumbs compact>
        <sp-breadcrumb-item value="1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item value="2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Links
By default, `sp-breadcrumbs` emits a `change` event when clicking on one of its children. However, there may be cases in which these should redirect to another page. This can be achieved by using the `href` attribute instead of `value`. Please note that the `change` event will no longer be triggered in this case.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item href="https://opensource.adobe.com/home">
            Home
        </sp-breadcrumb-item>
        <sp-breadcrumb-item href="https://opensource.adobe.com/trend">
            Trend
        </sp-breadcrumb-item>
        <sp-breadcrumb-item href="https://opensource.adobe.com/assets">
            March 2019 Assets
        </sp-breadcrumb-item>
    </sp-breadcrumbs>
## Overflowing
When the space is limited or the maximum number of visible items is reached, the component will render the first breadcrumbs inside an action menu. If needed, the last breadcrumb item will be truncated and will render a tooltip with the full text.
As recommended by Spectrum Design, by default the maximum visible breadcrumbs is 4. If you want to override this, you can use the `max-visible-items` attribute.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item value="your_stuff">Your stuff</sp-breadcrumb-item>
        <sp-breadcrumb-item value="team">Team</sp-breadcrumb-item>
        <sp-breadcrumb-item value="in_progress">In progress</sp-breadcrumb-item>
        <sp-breadcrumb-item value="files">Files</sp-breadcrumb-item>
        <sp-breadcrumb-item value="trend">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="winter">Winter</sp-breadcrumb-item>
        <sp-breadcrumb-item value="assets">Assets</sp-breadcrumb-item>
        <sp-breadcrumb-item value="18x24">18x24</sp-breadcrumb-item>
    </sp-breadcrumbs>
### Show root
Use the "root" slot to always render the first breadcrumb item, even if the breadcrumbs are overflowing.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item slot="root" value="your_stuff">
            Your stuff
        </sp-breadcrumb-item>
        <sp-breadcrumb-item value="team">Files</sp-breadcrumb-item>
        <sp-breadcrumb-item value="trend">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="winter">Winter</sp-breadcrumb-item>
        <sp-breadcrumb-item value="assets">Assets</sp-breadcrumb-item>
        <sp-breadcrumb-item value="18x24">18x24</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Custom Action Menu
The component offers the possibility to replace the action menu's icon with a custom one using the `icon` slot. Moreover, for accesibility purposes you can provide an internationalized string for the menu label using the `menu-label` attribute.
    
    <sp-breadcrumbs menu-label="Settings">
        <sp-icon-settings slot="icon"></sp-icon-settings>
    
        <sp-breadcrumb-item value="displays">Displays</sp-breadcrumb-item>
        <sp-breadcrumb-item value="main">Main display</sp-breadcrumb-item>
        <sp-breadcrumb-item value="brightness">Brightness</sp-breadcrumb-item>
        <sp-breadcrumb-item value="presets">Presets</sp-breadcrumb-item>
        <sp-breadcrumb-item value="1">Preset #1</sp-breadcrumb-item>
    </sp-breadcrumbs>
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
      <td>`compact`</td>
      <td>`compact`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>compact option is useful for reducing the height of the breadcrumbs</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td>Accessible name for the Breadcrumbs component</td>
    </tr>
    <tr>
      <td>`maxVisibleItems`</td>
      <td>`max-visible-items`</td>
      <td>`number`</td>
      <td>`4`</td>
      <td>Override the maximum number of visible items</td>
    </tr>
    <tr>
      <td>`menuLabel`</td>
      <td>`menu-label`</td>
      <td>`string`</td>
      <td>`'More items'`</td>
      <td>Change the default label of the action menu</td>
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
      <td>`icon`</td>
      <td>change the default icon of the action menu</td>
    </tr>
    <tr>
      <td>`root`</td>
      <td>Breadcrumb item to always display</td>
    </tr>
    <tr>
      <td>`default slot`</td>
      <td>Breadcrumb items</td>
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
      <td>`Announces the selected breadcrumb item`</td>
    </tr>
  </tbody>
</table>
