# sp-tag
Overview API
## Description
`<sp-tag>` elements represent a category to which the content they are a part of belongs. They can represent keywords or people, and are grouped to describe an item or a search request.
### Usage
    
    yarn add @spectrum-web-components/tags
    
Import the side effectful registration of `<sp-tag>` via:
    
    import '@spectrum-web-components/tags/sp-tag.js';
    
When looking to leverage the `Tag` base class as a type and/or for extension purposes, do so via:
    
    import { Tag } from '@spectrum-web-components/tags';
    
## Sizes
Small
    
    <sp-tags>
        <sp-tag size="s">Tag 1</sp-tag>
        <sp-tag invalid size="s">Tag 2</sp-tag>
        <sp-tag disabled size="s">Tag 3</sp-tag>
    </sp-tags>
Medium
    
    <sp-tags>
        <sp-tag size="m">Tag 1</sp-tag>
        <sp-tag invalid size="m">Tag 2</sp-tag>
        <sp-tag disabled size="m">Tag 3</sp-tag>
    </sp-tags>
Large
    
    <sp-tags>
        <sp-tag size="l">Tag 1</sp-tag>
        <sp-tag invalid size="l">Tag 2</sp-tag>
        <sp-tag disabled size="l">Tag 3</sp-tag>
    </sp-tags>
### Deletable
Use the `deletable` attribute to signify `<sp-tag>` elements that can be removed. The tags will only be focusable and deletable when they have the `deletable` and non-disabled attribute. When the clear button within the `<sp-tab>` is clicked, a `delete` event will be dispatched.
Default
    
    <sp-tags>
        <sp-tag deletable>Tag 1</sp-tag>
        <sp-tag invalid deletable>Tag 2</sp-tag>
        <sp-tag disabled deletable>Tag 3</sp-tag>
    </sp-tags>
With Avatars
    
    <sp-tags>
        <sp-tag deletable>
            Tag 1
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
        <sp-tag invalid deletable>
            Tag 2
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
        <sp-tag disabled deletable>
            Tag 3
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
    </sp-tags>
With Icons
    
    <sp-tags>
        <sp-tag deletable>
            Tag 1
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
        <sp-tag invalid deletable>
            Tag 2
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
        <sp-tag disabled deletable>
            Tag 3
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
    </sp-tags>
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
      <td>`deletable`</td>
      <td>`deletable`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`disabled`</td>
      <td>`disabled`</td>
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
      <td>text content for labeling the tag</td>
    </tr>
    <tr>
      <td>`avatar`</td>
      <td>an avatar element to display within the Tag</td>
    </tr>
    <tr>
      <td>`icon`</td>
      <td>an icon element to display within the Tag</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`delete`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
