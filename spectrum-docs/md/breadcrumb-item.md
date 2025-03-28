# sp-breadcrumb-item
Overview API
## Overview
For use within an `<sp-breadcrumbs>` element, an `<sp-breadcrumb-item>` represents a single item in a list.
### Usage
    
    yarn add @spectrum-web-components/breadcrumbs
    
Import the side effectful registration of `<sp-breadcrumb-item>` as follows:
    
    import '@spectrum-web-components/breadcrumbs/sp-breadcrumb-item.js';
    
When looking to leverage the `BreadcrumbItem` base class as a type and/or for extension purposes, do so via:
    
    import { BreadcrumbItem } from '@spectrum-web-components/breadcrumbs';
    
## Example
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item value="1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item value="2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Links
When using the `href` attribute instead of the `value` attribute, the breadcrumb item behaves as a regular anchor link.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item href="/1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item href="/2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item href="/3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
## Disabled
Disabled breadcrumb items no longer receive focus and keyboard events.
    
    <sp-breadcrumbs>
        <sp-breadcrumb-item disabled value="1">Home</sp-breadcrumb-item>
        <sp-breadcrumb-item disabled value="2">Trend</sp-breadcrumb-item>
        <sp-breadcrumb-item value="3">March 2019 Assets</sp-breadcrumb-item>
    </sp-breadcrumbs>
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
      <td>`value`</td>
      <td>`value`</td>
      <td>`string | undefined`</td>
      <td>`undefined`</td>
      <td></td>
    </tr>
  </tbody>
</table>
