# sp-field-label
Overview API
## Description
An `<sp-field-label>` provides accessible labelling for form elements. Use the `for` attribute to outline the `id` of an element in the same DOM tree to which it should associate itself.
### Usage
    
    yarn add @spectrum-web-components/field-label
    
Import the side effectful registration of `<sp-field-label>` via:
    
    import '@spectrum-web-components/field-label/sp-field-label.js';
    
When looking to leverage the `FieldLabel` base class as a type and/or for extension purposes, do so via:
    
    import { FieldLabel } from '@spectrum-web-components/field-label';
    
## Sizes
Small
    
    <sp-field-label for="lifestory-0" size="s">Life Story</sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-0"
    ></sp-textfield>
Medium
    
    <sp-field-label for="lifestory-1" size="m">Life Story</sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-1"
    ></sp-textfield>
Large
    
    <sp-field-label for="lifestory-2" size="l">Life Story</sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-2"
    ></sp-textfield>
Extra Large
    
    <sp-field-label for="lifestory-3" size="xl">Life Story</sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-3"
    ></sp-textfield>
## Examples
    
    <sp-field-label for="lifestory">Life Story</sp-field-label>
    <sp-textfield placeholder="Enter your life story" id="lifestory"></sp-textfield>
    <sp-field-label for="birth-place">Birthplace</sp-field-label>
    <sp-picker id="birth-place">
        <span slot="label">Choose a location:</span>
        <sp-menu-item>Istanbul</sp-menu-item>
        <sp-menu-item>London</sp-menu-item>
        <sp-menu-item>Maputo</sp-menu-item>
        <sp-menu-item>Melbuorne</sp-menu-item>
        <sp-menu-item>New York</sp-menu-item>
        <sp-menu-item>San Fransisco</sp-menu-item>
        <sp-menu-item>Santiago</sp-menu-item>
        <sp-menu-item>Tokyo</sp-menu-item>
    </sp-picker>
## Side Aligned
Using the `side-aligned` attribute will display the `<sp-field-label>` element inline with surrounding elements and the `start` and `end` values outline the alignment of the label text in the width specified.
### Start
Use `side-aligned="start"` to display the `<sp-field-label>` inline and to align the label text to the "start" of the flow of text:
    
    <sp-field-label for="lifestory-1" side-aligned="start" style="width: 120px">
        Life Story
    </sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-1"
    ></sp-textfield>
    <br />
    <br />
    <sp-field-label
        for="birth-place-1"
        side-aligned="start"
        required
        style="width: 120px"
    >
        Birthplace
    </sp-field-label>
    <sp-picker id="birth-place-1">
        <span slot="label">Choose a location:</span>
        <sp-menu-item>Istanbul</sp-menu-item>
        <sp-menu-item>London</sp-menu-item>
        <sp-menu-item>Maputo</sp-menu-item>
        <sp-menu-item>Melbuorne</sp-menu-item>
        <sp-menu-item>New York</sp-menu-item>
        <sp-menu-item>San Fransisco</sp-menu-item>
        <sp-menu-item>Santiago</sp-menu-item>
        <sp-menu-item>Tokyo</sp-menu-item>
    </sp-picker>
### End
Use `side-aligned="end"` to display the `<sp-field-label>` inline and to align the label text to the "end" of the flow of text:
    
    <sp-field-label
        for="lifestory-1"
        side-aligned="end"
        required
        style="width: 120px"
    >
        Life Story
    </sp-field-label>
    <sp-textfield
        placeholder="Enter your life story"
        id="lifestory-1"
    ></sp-textfield>
    <br />
    <br />
    <sp-field-label for="birth-place-1" side-aligned="end" style="width: 120px">
        Birthplace
    </sp-field-label>
    <sp-picker id="birth-place-1">
        <span slot="label">Choose a location:</span>
        <sp-menu-item>Istanbul</sp-menu-item>
        <sp-menu-item>London</sp-menu-item>
        <sp-menu-item>Maputo</sp-menu-item>
        <sp-menu-item>Melbuorne</sp-menu-item>
        <sp-menu-item>New York</sp-menu-item>
        <sp-menu-item>San Fransisco</sp-menu-item>
        <sp-menu-item>Santiago</sp-menu-item>
        <sp-menu-item>Tokyo</sp-menu-item>
    </sp-picker>
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
      <td></td>
    </tr>
    <tr>
      <td>`for`</td>
      <td>`for`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`id`</td>
      <td>`id`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`required`</td>
      <td>`required`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`sideAligned`</td>
      <td>`side-aligned`</td>
      <td>`'start' | 'end' | undefined`</td>
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
      <td>text content of the label</td>
    </tr>
  </tbody>
</table>
