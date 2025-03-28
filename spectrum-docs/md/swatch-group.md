# sp-swatch-group
Overview API
## Description
An `<sp-swatch-group>` group is a grouping of `<sp-swatch>` elements that are related to each other.
### Usage
    
    yarn add @spectrum-web-components/swatch
    
Import the side effectful registration of `<sp-swatch-group>` via:
    
    import '@spectrum-web-components/swatch/sp-swatch-group.js';
    
When looking to leverage the `SwatchGroup` base class as a type and/or for extension purposes, do so via:
    
    import { SwatchGroup } from '@spectrum-web-components/swatch';
    
## Sizes
Extra Small
    
    <sp-swatch-group size="xs">
        <sp-swatch color="var(--spectrum-gray-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-red-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-orange-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-yellow-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-chartreuse-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-celery-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-green-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-seafoam-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
Small
    
    <sp-swatch-group size="s">
        <sp-swatch color="var(--spectrum-gray-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-red-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-orange-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-yellow-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-chartreuse-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-celery-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-green-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-seafoam-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
Medium
    
    <sp-swatch-group size="m">
        <sp-swatch color="var(--spectrum-gray-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-red-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-orange-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-yellow-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-chartreuse-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-celery-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-green-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-seafoam-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
Large
    
    <sp-swatch-group size="l">
        <sp-swatch color="var(--spectrum-gray-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-red-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-orange-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-yellow-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-chartreuse-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-celery-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-green-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-seafoam-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
## Density
The `density` attribute/property is not required and when applied accepts the values of `compact` or `spacious`.
### Compact
    
    <sp-swatch-group density="compact">
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
### Spacious
    
    <sp-swatch-group density="spacious">
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
## Selection
An `<sp-swatch-group>` element can carry a selection of a `single` swatch or of `multiple` swatches. Then the `selects` property is set to one of these values, the `selected` property will surface an array the represents the string values that have been selected in the UI.
When the value of `selected` is updated via user input, the `change` event will be dispatched on the `<sp-swatch-group>` element to announce that interaction. Calling `preventDefault()` on the `chagne` event will prevent both the `<sp-swatch-group>` and the `<sp-swatch>` that initiated the `change` interaction from updating their `selected` values.
The value of `selected` can also be provided directly from the `<sp-swatch>` children. Child `<sp-swatch>` elements with their own `selected` attribute will be gathered and merged with any other selection data on the `<sp-swatch-group>` parent to populate `selected`.
### Single
The `selected` property is always represented as an array, and as such an application leveraging an `<sp-swatch-group>` element can apply more than one selection, regardless of the vaue of `selects`, however all future interactions will force the interace to a single selection.
    
    <sp-swatch-group
        selects="single"
        onchange="this.nextElementSibling.textContent = `Selected: ${JSON.stringify(this.selected, null, ' ')}`"
    >
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)" selected></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)" selected></sp-swatch>
    </sp-swatch-group>
    <div>
        Selected: [ "var(--spectrum-purple-500)", "var(--spectrum-magenta-500)" ]
    </div>
### Multiple
`<sp-swatch>` children of an `<sp-swatch-group selects="mutiple">` parent will toggle their selection.
    
    <sp-swatch-group
        selects="multiple"
        onchange="this.nextElementSibling.textContent = `Selected: ${JSON.stringify(this.selected, null, ' ')}`"
    >
        <sp-swatch color="var(--spectrum-blue-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-indigo-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-fuchsia-500)"></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)"></sp-swatch>
    </sp-swatch-group>
    <div>Selected: [ ]</div>
### Value
When available, the value of `selected` will be constructed from the `value` attributes/properties of the child `<sp-swatch>` elements. This can be useful when you would like the swatch data needs to correlate to a hash rather than the raw color string.
    
    <sp-swatch-group
        selects="multiple"
        selected='["color-2"]'
        onchange="this.nextElementSibling.textContent = `Selected: ${JSON.stringify(this.selected, null, ' ')}`"
    >
        <sp-swatch color="var(--spectrum-blue-500)" value="color-0"></sp-swatch>
        <sp-swatch
            color="var(--spectrum-indigo-500)"
            value="color-1"
            selected
        ></sp-swatch>
        <sp-swatch color="var(--spectrum-purple-500)" value="color-2"></sp-swatch>
        <sp-swatch
            color="var(--spectrum-fuchsia-500)"
            value="color-3"
            selected
        ></sp-swatch>
        <sp-swatch color="var(--spectrum-magenta-500)" value="color-4"></sp-swatch>
    </sp-swatch-group>
    <div>Selected: [ "color-2", "color-1", "color-3" ]</div>
## Swatch modifying attributes
An `<sp-swatch-group>` element can be modified by the following attributes/properties to customize its delivery as desired for your use case: `border`, `disabled`, `mixedValue` (accepted as the `mixed-value` attribute), `nothing`, `rounding`, `shape`, and `size`. Use these in concert with each other for a variety of final visual deliveries. Applying a value for one of these attributes/properties to an `<sp-swatch-group>` element will have it forward the value to all of the `<sp-swatch>` elements that are a direct child of the group, overriding any value that may be applied directly to those children.
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
      <td>`border`</td>
      <td>`border`</td>
      <td>`SwatchBorder`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`density`</td>
      <td>`density`</td>
      <td>`'compact' | 'spacious' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`rounding`</td>
      <td>`rounding`</td>
      <td>`SwatchRounding`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`string[]`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`selects`</td>
      <td>`selects`</td>
      <td>`SwatchSelects`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`shape`</td>
      <td>`shape`</td>
      <td>`SwatchShape`</td>
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
      <td>`default slot`</td>
      <td>Swatch elements to manage as a group</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`change`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
