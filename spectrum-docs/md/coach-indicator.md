# sp-coach-indicator
Overview API
## Description
`<sp-coach-indicator>` show the connection between an object and its explanation in a touring mode — for example, the source of  in an onboarding tour.
### Usage
    
    yarn add @spectrum-web-components/coachmark
    
Import the side effectful registration of `<sp-coach-indicator>` via:
    
    import '@spectrum-web-components/coachmark/sp-coach-indicator.js';
    
When looking to leverage the `CoachIndicator` base class as a type and/or for extension purposes, do so via:
    
    import { CoachIndicator } from '@spectrum-web-components/coachmark';
    
## Static color variants
Standard
    
    <sp-coach-indicator></sp-coach-indicator>
    <sp-coach-indicator static-color="dark"></sp-coach-indicator>
    <sp-coach-indicator static-color="light"></sp-coach-indicator>
Quiet
    
    <sp-coach-indicator quiet></sp-coach-indicator>
    <sp-coach-indicator quiet static-color="dark"></sp-coach-indicator>
    <sp-coach-indicator quiet static-color="light"></sp-coach-indicator>
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
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | 'black' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
  </tbody>
</table>
