# sp-underlay
Overview API
## Description
An `<sp-underlay>` is used primarily in concert with elements similar to `<sp-dialog>` that lay over the rest of your page to deliver a blocking layer between those two contexts. See this element in action as part of the `<sp-dialog-wrapper>` element.
### Usage
    
    yarn add @spectrum-web-components/underlay
    
Import the side effectful registration of `<sp-underlay>` via:
    
    import '@spectrum-web-components/underlay/sp-underlay.js';
    
When looking to leverage the `Underlay` base class as a type and/or for extension purposes, do so via:
    
    import { Underlay } from '@spectrum-web-components/underlay';
    
## Example
When leveraging an `<sp-underlay>` in conjunction with overlay content, place it as a sibling prior to your overlay content.
    
    <style>
        sp-underlay:not([open]) + sp-dialog {
            display: none;
        }
        sp-underlay + sp-dialog {
            position: fixed;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            z-index: 1;
            background: var(--spectrum-gray-100);
        }
    </style>
    
    <sp-button
        onclick="
            console.log(this.nextElementSibling);
            this.nextElementSibling.open = true;
        "
    >
        Open dialog with underlay element
    </sp-button>
    
    <sp-underlay></sp-underlay>
    <sp-dialog size="s">
        <h1 slot="heading">Hello, I'm an overlay!</h1>
        <p>Enjoy your day...</p>
        <sp-button
            slot="button"
            onclick="
                this.parentElement.previousElementSibling.open = false;
            "
        >
            Close
        </sp-button>
    </sp-dialog>
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
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
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
      <td>`When the underlay is "clicked" and the consuming pattern should chose whether to close based on that interaction`</td>
    </tr>
  </tbody>
</table>
