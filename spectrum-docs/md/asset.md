# sp-asset
Overview API
## Description
Use an `<sp-asset>` element to visually represent a file, folder or image in your application. File and folder representations will center themselves horizontally and vertically in the space provided to the element. Images will be contained to the element, growing to the element's full height while centering itself within the width provided.
### Installation
    
    yarn add @spectrum-web-components/asset
    
Import the side effectful registration of `<sp-asset>` via:
    
    import '@spectrum-web-components/asset/sp-asset.js';
    
When looking to leverage the `Asset` base class as a type and/or for extension purposes, do so via:
    
    import { Asset } from '@spectrum-web-components/asset';
    
## Example
    
    <sp-asset style="height: 128px">
        <img src="https://picsum.photos/500/500" alt="Demo Image" />
    </sp-asset>
### File
    
    <div class="flex">
        <sp-asset variant="file"></sp-asset>
        <sp-asset variant="file" label="Named File Asset"></sp-asset>
    </div>
### Folder
    
    <div class="flex">
        <sp-asset variant="folder"></sp-asset>
        <sp-asset variant="folder" label="Named Folder Asset"></sp-asset>
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
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`'file' | 'folder' | undefined`</td>
      <td>``</td>
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
      <td>content to be displayed in the asset when an acceptable value for `file` is not present</td>
    </tr>
  </tbody>
</table>
