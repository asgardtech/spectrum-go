# sp-link
Overview API
## Description
An `<sp-link>` allow users to navigate to a different location. They can be presented in-line inside a paragraph or as a standalone text.
### Usage
    
    yarn add @spectrum-web-components/link
    
Import the side effectful registration of `<sp-link>` via:
    
    import '@spectrum-web-components/link/sp-link.js';
    
When looking to leverage the `Link` base class as a type and/or for extension purposes, do so via:
    
    import { Link } from '@spectrum-web-components/link';
    
## Example
    
    This is an <sp-link href="#">example link</sp-link>.
## Variants
### Standard links
Standard links are blue and should be used to call attention to the link or for when the blue color won’t feel too overwhelming in the experience.
    
    This is a <sp-link href="#">standard link</sp-link>.
### Disabled links
Disabled links are blue and should not propagate any events and they are not focussable.
    
    This is a <sp-link disabled href="#">disabled link</sp-link>.
### Secondary links
The secondary variant is the same color as the paragraph text inline of which it appears. Its subdued appearance is optimal for when the primary variant is too overwhelming, such as in blocks of text with several references linked throughout.
    
    This is a <sp-link href="#" variant="secondary">secondary link</sp-link>.
### Static colored links
When a link needs to be placed on top of a colored background or a visual it may be appropriate to ship it with a static color, regardless of the theme settings with which it is delivered. Leverage the `static-color` attribute with its `white` or `black` values to ensure the delivery is the same in all contexts.
White
    
    <div
        style="background-color: #0f797d; padding: 15px 20px; display: inline-block;"
    >
        <p style="color: rgb(240, 240, 240);">
            This
            <sp-link static-color="white" href="#">link</sp-link>
            is over a background.
        </p>
    </div>
Black
    
    <div
        style="background-color: rgb(181, 209, 211); padding: 15px 20px; display: inline-block;"
    >
        <p style="color: rgb(15, 15, 15);">
            This
            <sp-link static-color="black" href="#">link</sp-link>
            is over a background.
        </p>
    </div>
### Quiet links
All links can have a quiet style, which means they don’t have an underline. This style should only be used when the placement and context of the link is explicit enough that a visible underline isn’t necessary. Quiet links are less accessible, so they should not be used for links that are essential to the experience. These are commonly used in website footers, where there are several lists of links that are shortcuts to other pages.
    
    <p>This is a <sp-link quiet href="#">quiet standard link</sp-link>.</p>
    <p>This is a <sp-link quiet variant="secondary" href="#">quiet secondary link</sp-link>.</p>
    <div
        style="background-color: #0f797d; padding: 15px 20px; display: inline-block;"
    >
        <p style="color: rgb(240, 240, 240);">
            This is a
            <sp-link static-color="white" quiet href="#">quiet link</sp-link>
            over a background.
        </p>
    </div>
### Download attribute
The download attribute on an `<a>` tag prompts a user to download a link as opposed to navigating to it. This attribute has been carried forward to `<sp-link>` to function the same.
While it functions this way without assigning a value, actually assigning the value allows custom naming of the download link in accordance with standard `<a>` rules defined by the browser.
    
    This is a <sp-link download="myfile.txt" href="#">download link</sp-link>.
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
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Uses quiet styles or not</td>
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
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'black' | 'white' | undefined`</td>
      <td>``</td>
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
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`'secondary' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
  </tbody>
</table>
