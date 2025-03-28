# sp-alert-banner
Overview API
## Description
The `<sp-alert-banner>` shows pressing and high-signal messages, such as system alerts. It is meant to be noticed and prompt users to take action.
### Usage
    
    yarn add @spectrum-web-components/alert-banner
    
Import the side effectful registration of `<sp-alert-banner>` via:
    
    import '@spectrum-web-components/alert-banner/sp-alert-banner.js';
    
When looking to leverage the `AlertBanner` base class as a type and/or for extension purposes, do so via:
    
    import { AlertBanner } from '@spectrum-web-components/alert-banner';
    
## Examples
The alert banner accepts text context in the default slot provided:
    
    <sp-alert-banner open>
        All documents in this folder have been archived
    </sp-alert-banner>
### Dismissible
Use the `dismissible` attribute to include an icon-only close button used to dismiss the alert banner:
    
    <sp-alert-banner open dismissible>
        All documents in this folder have been archived
    </sp-alert-banner>
### Actionable
Use the action slot for the contextual action that a user can take in response to the issue described:
    
    <sp-alert-banner open dismissible>
        Your trial has expired
        <sp-button treatment="outline" static-color="white" slot="action">
            Buy now
        </sp-button>
    </sp-alert-banner>
## Variants
### Info
    
    <sp-alert-banner open variant="info" dismissible>
        Your trial will expire in 3 days
        <sp-button treatment="outline" static-color="white" slot="action">
            Buy now
        </sp-button>
    </sp-alert-banner>
### Negative
    
    <sp-alert-banner open variant="negative" dismissible>
        Connection interupted. Check your network to continue
    </sp-alert-banner>
## Closing the alert banner
Alert banners should be used for system-level messages and they should be dismissed only as a result of a user action or if the internal state that triggered the alert has been resolved.
The alert can be dismissed by triggering the close button in case of a dismissible alert. It also exposes a public `close` method to allow consumers to close the alert programmatically.
The component dispatches a `close` event to announce that the alert banner has been closed. This can be prevented by using the `event.preventDefault()` API.
## Accessibility
The `sp-alert-banner` element is rendered with a `role` of `alert`, to inform screen readers and notify users accordingly. When rendering an alert that is dismissable or has actions on a page, it should be placed in a container with a `role` of `region` and that region should be labelled for screen readers users to quickly navigate.
The alert banner supports keyboard interaction as follows:
  -  *Tab* should place focus on the next interactive element, which can be either the actionable button or the close button.
  -  *Tab + Shift* should place focus on the previous interactive element.
  -  *Space* or *Enter* should trigger the interaction if one of the buttons is focused, thus dismissing the alert in case of the close button or triggering the corresponding contextual action.
  -  *Esc* will dismiss an alert banner if it’s a dismissible alert.

Once focus is on the alert banner, arrow keys can be used to navigate between the close button and the slotted action buttons.
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
      <td>`dismissible`</td>
      <td>`dismissible`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether to include an icon-only close button to dismiss the alert banner</td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Controls the display of the alert banner</td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`AlertBannerVariants`</td>
      <td>``</td>
      <td>The variant applies specific styling when set to `negative` or `info`; `variant` attribute is removed when it's passed an invalid variant.</td>
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
      <td>The alert banner text context</td>
    </tr>
    <tr>
      <td>`action`</td>
      <td>Slot for the button element that surfaces the contextual action a user can take</td>
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
      <td>`Announces the alert banner has been closed`</td>
    </tr>
  </tbody>
</table>
