# sp-tabs
Overview API
## Description
The `<sp-tabs>` displays a list of `<sp-tab>` element children as `role="tablist"`. An `<sp-tab>` element is associated with a sibling `<sp-tab-panel>` element via their `value` attribute. When an `<sp-tab>` element is `selected`, the associated `<sp-tab-panel>` will also be selected, showing that panel and hiding the others.
### Usage
    
    yarn add @spectrum-web-components/tabs
    
Import the side effectful registration of `<sp-tabs>`, `<sp-tab>` or `<sp-tab-panel>` via:
    
    import '@spectrum-web-components/tabs/sp-tabs.js';
    import '@spectrum-web-components/tabs/sp-tab.js';
    import '@spectrum-web-components/tabs/sp-tab-panel.js';
    
When looking to leverage the `Tabs`, `Tab`, or `TabPanel` base class as a type and/or for extension purposes, do so via:
    
    import {
        Tabs,
        Tab,
        TabPanel
    } from '@spectrum-web-components/tabs';
    
## Sizes
Small
    
    <sp-tabs selected="1" size="s">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Medium
    
    <sp-tabs selected="1" size="m">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Large
    
    <sp-tabs selected="1" size="l">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Extra Large
    
    <sp-tabs selected="1" size="xl">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
## Horizontal Tabs
An `<sp-tabs>` element will display horizontally by default. It can be modified with states like `compact`, `disabled`, and `quiet`, or with content like icons, etc.
Compact
Compact tabs should never be used without the quiet variation. Please use Quiet + Compact Tabs instead.
    
    <sp-tabs selected="1" compact>
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Disabled
When an `<sp-tabs>` element is given the `disabled` attribute its `<sp-tab>` children will be disabled as well, preventing a visitor from changing the selected tab. By default, `<sp-tab-panel>` children will not be addressed and the available content of the currently selected tab will be fully visible.
    
    <sp-tabs selected="2" disabled>
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1 is not selectable</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2 is selected</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3 is not selectable</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4 is not selectable</sp-tab-panel>
    </sp-tabs>
Icons
    
    <sp-tabs selected="1">
        <sp-tab label="Tab 1" value="1">
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab label="Tab 2" value="2">
            <sp-icon-close slot="icon"></sp-icon-close>
        </sp-tab>
        <sp-tab label="Tab 3" value="3">
            <sp-icon-chevron-down slot="icon"></sp-icon-chevron-down>
        </sp-tab>
        <sp-tab label="Tab 4" value="4">
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Quiet
    
    <sp-tabs selected="1" quiet>
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Quiet + Compact
    
    <sp-tabs selected="1" quiet compact>
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
## Vertical Tabs
An `<sp-tabs>` element will display horizontally by default. It can be modified with states like `compact`, `disabled`, and `quiet`, or with content like icons, etc.
Compact
Compact tabs should never be used without the quiet variation. Please use Quiet + Compact Tabs instead.
    
    <sp-tabs selected="1" compact direction="vertical">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Disabled
When an `<sp-tabs>` element is given the `disabled` attribute its `<sp-tab>` children will be disabled as well, preventing a visitor from changing the selected tab. By default, `<sp-tab-panel>` children will not be addressed and the available content of the currently selected tab will be fully visible.
    
    <sp-tabs selected="2" disabled direction="vertical">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1 is not selectable</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2 is selected</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3 is not selectable</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4 is not selectable</sp-tab-panel>
    </sp-tabs>
Icons
    
    <sp-tabs selected="1" direction="vertical">
        <sp-tab label="Tab 1" value="1">
            <sp-icon-checkmark slot="icon"></sp-icon-checkmark>
        </sp-tab>
        <sp-tab label="Tab 2" value="2">
            <sp-icon-close slot="icon"></sp-icon-close>
        </sp-tab>
        <sp-tab label="Tab 3" value="3">
            <sp-icon-chevron-down slot="icon"></sp-icon-chevron-down>
        </sp-tab>
        <sp-tab label="Tab 4" value="4">
            <sp-icon-help slot="icon"></sp-icon-help>
        </sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Quiet
    
    <sp-tabs selected="1" quiet direction="vertical">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
Quiet + Compact
    
    <sp-tabs selected="1" quiet compact direction="vertical">
        <sp-tab label="Tab 1" value="1"></sp-tab>
        <sp-tab label="Tab 2" value="2"></sp-tab>
        <sp-tab label="Tab 3" value="3"></sp-tab>
        <sp-tab label="Tab 4" value="4"></sp-tab>
        <sp-tab-panel value="1">Content for Tab 1</sp-tab-panel>
        <sp-tab-panel value="2">Content for Tab 2</sp-tab-panel>
        <sp-tab-panel value="3">Content for Tab 3</sp-tab-panel>
        <sp-tab-panel value="4">Content for Tab 4</sp-tab-panel>
    </sp-tabs>
## Accessibility
When an `<sp-tabs>` has a `selected` value, the `<sp-tab>` child of that `value` will be given `[tabindex="0"]` and will receive initial focus when tabbing into the `<sp-tabs>` element. When no `selected` value is present, the first `<sp-tab>` child will be treated in this way. When focus is currently within the `<sp-tabs>` element, the left and right arrows will move that focus back and forth through the available `<sp-tab>` children.
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
      <td>`auto`</td>
      <td>`auto`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether to activate a tab on keyboard focus or not.</td>
    </tr>
    <tr>
      <td>`compact`</td>
      <td>`compact`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>The tab items are displayed closer together.</td>
    </tr>
    <tr>
      <td>`direction`</td>
      <td>`direction`</td>
      <td>`'vertical' | 'vertical-right' | 'horizontal'`</td>
      <td>`'horizontal'`</td>
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
      <td>`enableTabsScroll`</td>
      <td>`enableTabsScroll`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>The tab list is displayed without a border.</td>
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
      <td>Tab elements to manage as a group</td>
    </tr>
    <tr>
      <td>`tab-panel`</td>
      <td>Tab Panel elements related to the listed Tab elements</td>
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
      <td>`The selected Tab child has changed.`</td>
    </tr>
  </tbody>
</table>
