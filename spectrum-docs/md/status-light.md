# sp-status-light
Overview API
## Description
An `<sp-status-light>` is a great way to convey semantic meaning, such as statuses and categories.
### Usage
    
    yarn add @spectrum-web-components/status-light
    
Import the side effectful registration of `<sp-status-light>` via:
    
    import '@spectrum-web-components/status-light/sp-status-light.js';
    
When looking to leverage the `StatusLight` base class as a type and/or for extension purposes, do so via:
    
    import { StatusLight } from '@spectrum-web-components/status-light';
    
## Sizes
Small
    
    <sp-status-light size="s" variant="positive">approved</sp-status-light>
Medium
    
    <sp-status-light size="m" variant="positive">approved</sp-status-light>
Large
    
    <sp-status-light size="l" variant="positive">approved</sp-status-light>
Extra Large
    
    <sp-status-light size="xl" variant="positive">approved</sp-status-light>
## Variants
There are many variants to choose from in Spectrum. The `variant` attribute controls the main variant of the status light, and `neutral` being the default. Following are the supported variants:
  -  positive
  -  negative
  -  notice
  -  info
  -  neutral
  -  yellow
  -  fuchsia
  -  indigo
  -  seafoam
  -  chartreuse
  -  magenta
  -  celery
  -  purple

## Disabled
    
    <sp-status-light variant="positive" disabled>disabled</sp-status-light>
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
      <td>A status light in a disabled state shows that a status exists, but is not available in that circumstance. This can be used to maintain layout continuity and communicate that a status may become available later.</td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`| 'negative' | 'notice' | 'positive' | 'info' | 'neutral' | 'yellow' | 'fuchsia' | 'indigo' | 'seafoam' | 'chartreuse' | 'magenta' | 'celery' | 'purple'`</td>
      <td>`'info'`</td>
      <td>The visual variant to apply to this status light.</td>
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
      <td>text label of the Status Light</td>
    </tr>
  </tbody>
</table>
