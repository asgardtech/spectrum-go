# sp-toast
Overview API
## Description
`sp-toast` elements display brief, temporary notifications. They are noticeable but do not disrupt the user experience and do not require an action to be taken.
### Usage
    
    yarn add @spectrum-web-components/toast
    
Import the side effectful registration of `<sp-toast>` via:
    
    import '@spectrum-web-components/toast/sp-toast.js';
    
When looking to leverage the `Toast` base class as a type and/or for extension purposes, do so via:
    
    import { Toast } from '@spectrum-web-components/toast';
    
## Example
    
    <sp-toast open>
        This is important information that you should read, soon.
    </sp-toast>
### With actions
    
    <sp-toast open>
        This is important information that you should read, soon.
        <sp-button
            slot="action"
            static-color="white"
            variant="secondary"
            treatment="outline"
        >
            Do something
        </sp-button>
    </sp-toast>
### Wrapping
    
    <sp-toast open style="width: 300px">
        This is important information that you should read, soon.
        <sp-button
            slot="action"
            static-color="white"
            variant="secondary"
            treatment="outline"
        >
            Do something
        </sp-button>
    </sp-toast>
## Variants
### Negative
    
    <sp-toast open variant="negative">
        This is negative information that you should read, soon.
    </sp-toast>
### Positive
    
    <sp-toast open variant="positive">
        This is positive information that you should read, soon.
    </sp-toast>
### Info
    
    <sp-toast open variant="info">
        This is information that you should read.
    </sp-toast>
## Accessibility
An `<sp-toast>` element is by default rendered with a `role` of `alert`. When rendering the `<sp-toast>` to a page, it should be place in a container with a `role` of `region`.
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
      <td>`iconLabel`</td>
      <td>`icon-label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The `iconLabel` property is used to set the `label` attribute on the icon element. This is used to provide a text alternative for the icon to ensure accessibility.</td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>The `open` property indicates whether the toast is visible or hidden.</td>
    </tr>
    <tr>
      <td>`timeout`</td>
      <td>`timeout`</td>
      <td>`number | null`</td>
      <td>`null`</td>
      <td>When a timeout is provided, it represents the number of milliseconds from when the Toast was placed on the page before it will automatically dismiss itself.</td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`ToastVariants`</td>
      <td>``</td>
      <td>The variant applies specific styling when set to `negative`, `positive`, `info`, `error`, or `warning`.</td>
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
      <td>The toast content</td>
    </tr>
    <tr>
      <td>`action`</td>
      <td>button element surfacing an action in the Toast</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`close`</td>
      <td>`CustomEvent`</td>
      <td>`Announces that the Toast has been closed by the user or by its timeout.`</td>
    </tr>
  </tbody>
</table>
