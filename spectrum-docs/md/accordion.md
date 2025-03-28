# sp-accordion
Overview API
## Description
The `<sp-accordion>` element contains a list of items that can be expanded or collapsed to reveal additional content or information associated with each item. There can be zero expanded items, exactly one expanded item, or more than one item expanded at a time, depending on the configuration. This list of items is defined by child `<sp-accordion-item>` elements that are targetted to the default slot of their `<sp-accordion>` parent.
### Usage
    
    yarn add @spectrum-web-components/accordion
Import the side effectful registration of `<sp-accordion>` and `<sp-accordion-item>` via:
    
    import '@spectrum-web-components/accordion/sp-accordion.js';
    import '@spectrum-web-components/accordion/sp-accordion-item.js';
When looking to leverage the `Accordion` and `AccordionItem` base class as a type and/or for extension purposes, do so via:
    
    import { Accordion, AccordionItem } from '@spectrum-web-components/accordion';
## Example
    
    <sp-accordion>
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 5">
            <div>Item 5</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 6">
            <div>Item 6</div>
        </sp-accordion-item>
    </sp-accordion>
## Allow Multiple
    
    <sp-accordion allow-multiple>
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 5">
            <div>Item 5</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 6">
            <div>Item 6</div>
        </sp-accordion-item>
    </sp-accordion>
## Sizes
Small
    
    <sp-accordion size="s">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Medium
    
    <sp-accordion size="m">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Large
    
    <sp-accordion size="l">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Extra Large
    
    <sp-accordion size="xl">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
## Density
The `density` property, when applied, accepts the values of `compact` or `spacious`.
### Compact
Small
    
    <sp-accordion density="compact" size="s">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Medium
    
    <sp-accordion density="compact" size="m">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Large
    
    <sp-accordion density="compact" size="l">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Extra Large
    
    <sp-accordion density="compact" size="xl">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
### Spacious
Small
    
    <sp-accordion density="spacious" size="s">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Medium
    
    <sp-accordion density="spacious" size="m">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Large
    
    <sp-accordion density="spacious" size="l">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
Extra Large
    
    <sp-accordion density="spacious" size="xl">
        <sp-accordion-item label="Heading 1">
            <div>Item 1</div>
        </sp-accordion-item>
        <sp-accordion-item disabled label="Heading 2">
            <div>Item 2</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 3">
            <div>Item 3</div>
        </sp-accordion-item>
        <sp-accordion-item label="Heading 4">
            <div>Item 4</div>
        </sp-accordion-item>
    </sp-accordion>
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
      <td>`allowMultiple`</td>
      <td>`allow-multiple`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Allows multiple accordion items to be opened at the same time</td>
    </tr>
    <tr>
      <td>`density`</td>
      <td>`density`</td>
      <td>`'compact' | 'spacious' | undefined`</td>
      <td>``</td>
      <td>Sets the spacing between the content to borders of an accordion item</td>
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
      <td>The sp-accordion-item children to display.</td>
    </tr>
  </tbody>
</table>
