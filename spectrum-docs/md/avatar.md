# sp-avatar
Overview API
## Description
An `<sp-avatar>` is a great way to feature a visual representation of a user.
### Usage
    
    yarn add @spectrum-web-components/avatar
    
Import the side effectful registration of `<sp-avatar>` via:
    
    import '@spectrum-web-components/avatar/sp-avatar.js';
    
When looking to leverage the `Avatar` base class as a type and/or for extension purposes, do so via:
    
    import { Avatar } from '@spectrum-web-components/avatar';
    
## Sizes
50
    
    <sp-avatar
        size="50"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
75
    
    <sp-avatar
        size="75"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
100
    
    <sp-avatar
        size="100"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
200
    
    <sp-avatar
        size="200"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
300
    
    <sp-avatar
        size="300"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
400
    
    <sp-avatar
        size="400"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
500
    
    <sp-avatar
        size="500"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
600
    
    <sp-avatar
        size="600"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
700
    
    <sp-avatar
        size="700"
        label="Demo User"
        src="https://picsum.photos/500/500"
    ></sp-avatar>
## Accessibility
The `label` attribute of the `<sp-avatar>` will be passed into the `<img>` element as the `alt` tag for use in defining a textual representation of the image displayed.
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
      <td>`size`</td>
      <td>`size`</td>
      <td>`AvatarSize`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`src`</td>
      <td>`src`</td>
      <td>`string`</td>
      <td>`''`</td>
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
  </tbody>
</table>
