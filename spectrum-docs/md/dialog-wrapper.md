# sp-dialog-wrapper
Overview API
## Description
`sp-dialog-wrapper` supplies an attribute based interface for the managed customization of an `sp-dialog` element and the light DOM supplied to it. This is paired it with an `underlay` attribute that opts-in to the use of an `sp-underlay` element between your page content and the `sp-dialog` that opens over it.
### Usage
    
    yarn add @spectrum-web-components/dialog
    
Import the side effectful registration of `<sp-dialog-wrapper>` via:
    
    import '@spectrum-web-components/dialog/sp-dialog-wrapper.js';
    
When looking to leverage the `DialogWrapper` base class as a type and/or for extension purposes, do so via:
    
    import { DialogWrapper } from '@spectrum-web-components/dialog';
    
## Example
### Small
    
    <overlay-trigger type="modal">
        <sp-dialog-wrapper
            slot="click-content"
            headline="Dialog title"
            dismissable
            dismiss-label="Close"
            underlay
            footer="Content for footer"
        >
            Content of the dialog
        </sp-dialog-wrapper>
        <sp-button slot="trigger" variant="primary">Toggle Dialog</sp-button>
    </overlay-trigger>
### Fullscreen Mode
    
    <overlay-trigger type="modal">
        <sp-dialog-wrapper
            slot="click-content"
            headline="Dialog title"
            mode="fullscreen"
            confirm-label="Keep Both"
            secondary-label="Replace"
            cancel-label="Cancel"
            footer="Content for footer"
        >
            Content of the dialog
        </sp-dialog-wrapper>
        <sp-button
            slot="trigger"
            variant="primary"
            onClick="
                const overlayTrigger = this.parentElement;
                const dialogWrapper = overlayTrigger.clickContent;
                function handleEvent({type}) {
                    spAlert(this, `<sp-dialog-wrapper> '${type}' event handled.`);
                    dialogWrapper.open = false;
                    dialogWrapper.removeEventListener('confirm', handleEvent);
                    dialogWrapper.removeEventListener('secondary', handleEvent);
                    dialogWrapper.removeEventListener('cancel', handleEvent);
                }
                dialogWrapper.addEventListener('confirm', handleEvent);
                dialogWrapper.addEventListener('secondary', handleEvent);
                dialogWrapper.addEventListener('cancel', handleEvent);
            "
        >
            Toggle Dialog
        </sp-button>
    </overlay-trigger>
### Fullscreen Takeover Mode
    
    <overlay-trigger type="modal">
        <sp-dialog-wrapper
            slot="click-content"
            headline="Dialog title"
            mode="fullscreenTakeover"
            confirm-label="Keep Both"
            secondary-label="Replace"
            cancel-label="Cancel"
            footer="Content for footer"
        >
            Content of the dialog
        </sp-dialog-wrapper>
        <sp-button
            slot="trigger"
            variant="primary"
            onClick="
                const overlayTrigger = this.parentElement;
                const dialogWrapper = overlayTrigger.clickContent;
                function handleEvent({type}) {
                    spAlert(this, `<sp-dialog-wrapper> '${type}' event handled.`);
                    dialogWrapper.open = false;
                    dialogWrapper.removeEventListener('confirm', handleEvent);
                    dialogWrapper.removeEventListener('secondary', handleEvent);
                    dialogWrapper.removeEventListener('cancel', handleEvent);
                }
                dialogWrapper.addEventListener('confirm', handleEvent);
                dialogWrapper.addEventListener('secondary', handleEvent);
                dialogWrapper.addEventListener('cancel', handleEvent);
            "
        >
            Toggle Dialog
        </sp-button>
    </overlay-trigger>
## Accessibility
An `sp-dialog-wrapper` element leverages the `headline` attribute/property to label the dialog content for screen readers. The `headline-visibility` attribute will accept a value of `"none"` to suppress the headline visually.
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
      <td>`cancelLabel`</td>
      <td>`cancel-label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`confirmLabel`</td>
      <td>`confirm-label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`dismissLabel`</td>
      <td>`dismiss-label`</td>
      <td>`string`</td>
      <td>`'Close'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`dismissable`</td>
      <td>`dismissable`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`error`</td>
      <td>`error`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`footer`</td>
      <td>`footer`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`headline`</td>
      <td>`headline`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`headlineVisibility`</td>
      <td>`headline-visibility`</td>
      <td>`'none' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`hero`</td>
      <td>`hero`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`heroLabel`</td>
      <td>`hero-label`</td>
      <td>`string`</td>
      <td>`''`</td>
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
      <td>`noDivider`</td>
      <td>`no-divider`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
      <td>`secondaryLabel`</td>
      <td>`secondary-label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`size`</td>
      <td>`size`</td>
      <td>`'s' | 'm' | 'l' | undefined`</td>
      <td>``</td>
      <td></td>
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
      <td>content for the dialog</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`cancel`</td>
      <td>`Event`</td>
      <td>`Announces that the "cancel" button has been clicked.`</td>
    </tr>
    <tr>
      <td>`close`</td>
      <td>`Event`</td>
      <td>`Announces that the dialog has been closed.`</td>
    </tr>
    <tr>
      <td>`confirm`</td>
      <td>`Event`</td>
      <td>`Announces that the "confirm" button has been clicked.`</td>
    </tr>
    <tr>
      <td>`secondary`</td>
      <td>`Event`</td>
      <td>`Announces that the "secondary" button has been clicked.`</td>
    </tr>
    <tr>
      <td>`undefined`</td>
      <td>`TransitionEvent`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
