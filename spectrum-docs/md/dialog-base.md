# sp-dialog-base
Overview API
## Description
`sp-dialog-base` accepts slotted dialog content (often an `<sp-dialog>`) and presents that content in a container that is animated into place when toggling the `open` attribute. In concert with the Overlay API or Overlay Trigger, the provided dialog content will be displayed over the rest of the page. Leverage the `interaction = modal` and `receivesFocus = 'auto'` settings in the Overlay API to ensure that focus is thrown into the dialog content when opened and that the tab order will be trapped within it while open.
### Usage
    
    yarn add @spectrum-web-components/dialog
    
Import the side effectful registration of `<sp-dialog-base>` via:
    
    import '@spectrum-web-components/dialog/sp-dialog-base.js';
    
When looking to leverage the `DialogBase` base class as a type and/or for extension purposes, do so via:
    
    import { DialogBase } from '@spectrum-web-components/dialog';
    
## Example
    
    <overlay-trigger type="modal">
        <sp-dialog-base underlay slot="click-content">
            <sp-dialog size="s">
                <h2 slot="heading">A thing is about to happen</h2>
                <p>Something that might happen a lot is about to happen.</p>
                <p>
                    The click events for the "OK" button are bound to the story not
                    to the components in specific.
                </p>
                <sp-button
                    variant="secondary"
                    treatment="fill"
                    slot="button"
                    onclick="this.dispatchEvent(new Event('close', { bubbles: true, composed: true }));"
                >
                    Ok
                </sp-button>
                <sp-checkbox slot="footer">Don't show me this again</sp-checkbox>
            </sp-dialog>
        </sp-dialog-base>
        <sp-button slot="trigger" variant="primary">Toggle Dialog</sp-button>
    </overlay-trigger>
## Dialog
`sp-dialog-base` expects a single slotted child element to play the role of the dialog that it will deliver within your application. When leveraging it as a base class be sure to customize the `dialog` getter to ensure that it acquires the appropriate element for your use case in order to correctly pass focus into your content when the dialog is opened.
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
      <td>`dismissable`</td>
      <td>`dismissable`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`mode`</td>
      <td>`mode`</td>
      <td>`'fullscreen' | 'fullscreenTakeover' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`responsive`</td>
      <td>`responsive`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>When set to true, fills screens smaller than 350px high and 400px wide with the full dialog.</td>
    </tr>
    <tr>
      <td>`underlay`</td>
      <td>`underlay`</td>
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
      <td>A Dialog element to display.</td>
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
      <td>`Event`</td>
      <td>`Announces that the dialog has been closed.`</td>
    </tr>
    <tr>
      <td>`undefined`</td>
      <td>`TransitionEvent`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
