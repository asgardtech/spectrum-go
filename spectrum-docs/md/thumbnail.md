# sp-thumbnail
Overview API
## Description
An `sp-thumbnail` can be used in a variety of locations as a way to display a preview of an image, layer, or effect. `sp-thumbnail` elements are not keyboard-focusable since they're intended to be used inside of a component that a user sets focus to (such as select lists or tree items).
### Usage
    
    yarn add @spectrum-web-components/thumbnail
    
Import the side effectful registration of `<sp-thumbnail>` via:
    
    import '@spectrum-web-components/thumbnail/sp-thumbnail.js';
    
When looking to leverage the `Thumbnail` base class as a type and/or for extension purposes, do so via:
    
    import { Thumbnail } from '@spectrum-web-components/thumbnail';
    
## Sizes
50
    
    <sp-thumbnail size="50">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
75
    
    <sp-thumbnail size="75">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
100
    
    <sp-thumbnail size="100">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
200
    
    <sp-thumbnail size="200">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
300
    
    <sp-thumbnail size="300">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
400
    
    <sp-thumbnail size="400">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
500
    
    <sp-thumbnail size="500">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
600
    
    <sp-thumbnail size="600">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
700
    
    <sp-thumbnail size="700">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
800
    
    <sp-thumbnail size="800">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
900
    
    <sp-thumbnail size="900">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
1000
    
    <sp-thumbnail size="1000">
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
## Focused
When `focused` the `sp-thumbnail` element will be displayed as follows:
    
    <sp-thumbnail focused>
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
## Disabled
Thumbnail should only be displayed as disabled if the entire componet is also disabled. When `disabled` the `sp-thumbnail` element will be displayed as follows:
    
    <sp-thumbnail disabled>
        <img src="https://picsum.photos/100/100" alt="Demo Image" />
    </sp-thumbnail>
## Representing non-square content
By default, an `sp-thumbnail` will ensure that the entirety of the content that it respresents is visible by letterboxing that content with a checkerboard background when its aspect ratio is not square.
    
    <div style="display: flex; gap: var(--spectrum-spacing-100);">
        <sp-thumbnail>
            <img src="https://picsum.photos/300/400" alt="Demo Image" />
        </sp-thumbnail>
    
        <sp-thumbnail>
            <img src="https://picsum.photos/500/100" alt="Demo Image" />
        </sp-thumbnail>
    </div>
The `background` attribute takes a string value of the CSS "background" property in order to customize the letterboxing.
    
    <div style="display: flex; gap: var(--spectrum-spacing-100);">
        <sp-thumbnail background="red">
            <img src="https://picsum.photos/300/400" alt="Demo Image" />
        </sp-thumbnail>
    
        <sp-thumbnail background="#00ff00">
            <img src="https://picsum.photos/500/100" alt="Demo Image" />
        </sp-thumbnail>
    </div>
The `cover` attribute will cause the content to fill the space provided by the `sp-thumbnail` element:
    
    <div style="display: flex; gap: var(--spectrum-spacing-100);">
        <sp-thumbnail cover>
            <img src="https://picsum.photos/300/400" alt="Demo Image" />
        </sp-thumbnail>
    
        <sp-thumbnail cover>
            <img src="https://picsum.photos/500/100" alt="Demo Image" />
        </sp-thumbnail>
    </div>
## Layer and Layer Selected
For when `sp-thumbail` is used in layer management (such as the Compact or Detail Layers panels). The thumbnail is given a thick blue border to indicate its selection when used in layer management.
    
    <div style="display: flex; gap: var(--spectrum-spacing-100);">
        <sp-thumbnail layer>
            <img src="https://picsum.photos/400/400" alt="Demo Image" />
        </sp-thumbnail>
    
        <sp-thumbnail layer selected>
            <img src="https://picsum.photos/500/100" alt="Demo Image" />
        </sp-thumbnail>
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
      <td>`background`</td>
      <td>`background`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`cover`</td>
      <td>`cover`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`layer`</td>
      <td>`layer`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`size`</td>
      <td>`size`</td>
      <td>`ThumbnailSize`</td>
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
      <td>`image`</td>
      <td>image element to present in the Thumbnail</td>
    </tr>
  </tbody>
</table>
