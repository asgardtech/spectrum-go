# sp-meter
Overview API
## Description
An `<sp-meter>` is a visual representation of a quantity or achievement. The meter's progress is determined by user actions, rather than system actions.
### Installation
    
    yarn add @spectrum-web-components/meter
    
Import the side-effectful registration of `<sp-meter>` via:
    
    import '@spectrum-web-components/meter/sp-meter.js';
    
When looking to leverage the `Meter` base class as a type and/or for extension purposes, do so via:
    
    import { Meter } from '@spectrum-web-components/meter';
    
## Sizes
Small
    
    <sp-meter size="s" progress="71">Tasks Completed</sp-meter>
Medium
    
    <sp-meter size="m" progress="71">Tasks Completed</sp-meter>
Large
    
    <sp-meter size="l" progress="71">Tasks Completed</sp-meter>
Extra Large
    
    <sp-meter size="xl" progress="71">Tasks Completed</sp-meter>
## Variants
### Informative
By default, the informative variant can be used to represent a neutral or non-semantic value, such as the number of tutorials completed.
    
    <sp-meter progress="50">Storage Space</sp-meter>
### Positive
The positive variant can be used to represent a positive semantic value, such as when there’s a lot of space remaining. Use value `variant="positive"` to define a positive variant.
    
    <sp-meter variant="positive" progress="50">Storage Space</sp-meter>
### Notice
The notice variant can be used to warn users about a situation that may need to be addressed soon, such as when space remaining is becoming limited. Use value `variant="notice"` to define a notice variant.
    
    <sp-meter variant="notice" progress="73">Storage Space</sp-meter>
### Negative
The negative variant can be used to warn users about a critical situation that needs their urgent attention, such as when space remaining is becoming very limited. Use value `variant="negative"` to define a negative variant.
    
    <sp-meter variant="negative" progress="92">Storage Space</sp-meter>
### Side Label
A meter can be delivered with its labeling displayed above its visual indicator or to either side. Use the boolean `[side-label]` attribute to define where this content should appear.
    
    <sp-meter side-label>Side Label</sp-meter>
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
      <td>`label`</td>
      <td>`label`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`progress`</td>
      <td>`progress`</td>
      <td>`number`</td>
      <td>`0`</td>
      <td></td>
    </tr>
    <tr>
      <td>`sideLabel`</td>
      <td>`side-label`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`MeterVariants`</td>
      <td>``</td>
      <td>The variant applies specific styling when set to `negative`, `positive`, `notice` `variant` attribute is removed when not matching one of the above.</td>
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
      <td>text labeling the Meter</td>
    </tr>
  </tbody>
</table>
