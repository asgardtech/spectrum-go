# sp-tags
Overview API
## Description
`<sp-tags>` elements contain a collection of `<sp-tag>` elements and allow users to categorize content. They can represent keywords or people, and are grouped to describe an item or a search request.
### Usage
    
    yarn add @spectrum-web-components/tags
    
Import the side effectful registration of `<sp-tags>` or `<sp-tag>` via:
    
    import '@spectrum-web-components/tags/sp-tags.js';
    import '@spectrum-web-components/tags/sp-tag.js';
    
When looking to leverage the `Tags` or `Tag` base classes as a type and/or for extension purposes, do so via:
    
    import {
        Tags,
        Tag
    } from '@spectrum-web-components/tags';
    
## Example
    
    <sp-tags>
        <sp-tag>Tag 1</sp-tag>
        <sp-tag invalid>Tag 2</sp-tag>
        <sp-tag disabled>Tag 3</sp-tag>
    </sp-tags>
    
    <sp-tags>
        <sp-tag>
            Tag 1
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
        <sp-tag invalid>
            Tag 2
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
        <sp-tag disabled>
            Tag 3
            <sp-avatar
                slot="avatar"
                label="Tag 1"
                src=https://picsum.photos/500/500
            ></sp-avatar>
        </sp-tag>
    </sp-tags>
    
    <sp-tags>
        <sp-tag>
            Tag 1
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
        <sp-tag invalid>
            Tag 2
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
        <sp-tag disabled>
            Tag 3
            <sp-icon-magnify slot="icon" size="s"></sp-icon-magnify>
        </sp-tag>
    </sp-tags>
## API
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
      <td>Tag elements to manage as a group</td>
    </tr>
  </tbody>
</table>
