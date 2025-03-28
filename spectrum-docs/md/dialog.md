# sp-dialog
Overview API
## Description
`sp-dialog` displays important information that users need to acknowledge. They appear over the interface and block further interactions. When used directly the `sp-dialog` element surfaces a `slot` based API for deep customization of the content to be included in the overlay.
### Usage
    
    yarn add @spectrum-web-components/dialog
    
Import the side effectful registration of `<sp-dialog>` via:
    
    import '@spectrum-web-components/dialog/sp-dialog.js';
    
When looking to leverage the `Dialog` base class as a type and/or for extension purposes, do so via:
    
    import { Dialog } from '@spectrum-web-components/dialog';
    
## Sizes
Small
    
    <sp-dialog size="s">
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
    </sp-dialog>
Medium
    
    <sp-dialog size="m">
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
    </sp-dialog>
Large
    
    <sp-dialog size="l">
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
    </sp-dialog>
## Variants
### Dismissable
When supplied with the `dissmissable` attribute an `<sp-dialog>` element will surface a "close" button afordance that will dispatch a DOM event with the name of `close` when pressed.
Note: the `dissmissable` attribute will not be followed when `mode="fullscreen"` or `mode="fullscreenTakeover"` are applies in accordance with the Spectrum specification.
    
    <sp-dialog size="m" dismissable>
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
    </sp-dialog>
### No Divider
    
    <sp-dialog size="m" dismissable no-divider>
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
    </sp-dialog>
### Hero
    
    <sp-dialog size="medium" dismissable no-divider>
        <div
            slot="hero"
            style="background-image: url(https://picsum.photos/1400/260)"
        ></div>
        <h2 slot="heading">Disclaimer</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Auctor augue mauris
        augue neque gravida. Libero volutpat sed ornare arcu. Quisque egestas diam
        in arcu cursus euismod quis viverra. Posuere ac ut consequat semper viverra
        nam libero justo laoreet. Enim ut tellus elementum sagittis vitae et leo
        duis ut. Neque laoreet suspendisse interdum consectetur libero id faucibus
        nisl. Diam volutpat commodo sed egestas egestas. Dolor magna eget est lorem
        ipsum dolor. Vitae suscipit tellus mauris a diam maecenas sed. Turpis in eu
        mi bibendum neque egestas congue. Rhoncus est pellentesque elit ullamcorper
        dignissim cras lobortis.
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
      <td>`size`</td>
      <td>`size`</td>
      <td>`'s' | 'm' | 'l' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`AlertDialogVariants`</td>
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
      <td>`button`</td>
      <td>Button elements addressed to this slot may be placed below the content when not delivered in a fullscreen mode</td>
    </tr>
    <tr>
      <td>`footer`</td>
      <td>Content addressed to the `footer` will be placed below the main content and to the side of any `[slot='button']` content</td>
    </tr>
    <tr>
      <td>`heading`</td>
      <td>Acts as the heading of the dialog. This should be an actual heading tag `</td>
    </tr>
    <tr>
      <td>`hero`</td>
      <td>Accepts a hero image to display at the top of the dialog</td>
    </tr>
    <tr>
      <td>`default slot`</td>
      <td>Content not addressed to a specific slot will be interpreted as the main content of the dialog</td>
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
  </tbody>
</table>
