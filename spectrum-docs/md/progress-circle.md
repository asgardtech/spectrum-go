# sp-progress-circle
Overview API
## Description
An `<sp-progress-circle>` shows the progression of a system operation such as downloading, uploading, processing, etc. in a visual way. It can represent determinate or indeterminate progress.
### Usage
    
    yarn add @spectrum-web-components/progress-circle
    
Import the side effectful registration of `<sp-progress-circle>` via:
    
    import '@spectrum-web-components/progress-circle/sp-progress-circle.js';
    
When looking to leverage the `ProgressCircle` base class as a type and/or for extension purposes, do so via:
    
    import { ProgressCircle } from '@spectrum-web-components/progress-circle';
    
## Variants
### Default
An `<sp-progress-circle>` is used to visually show the progression of a system operation such as downloading, uploading, processing, etc.
    
    <div
        style="width: 250px; height: 150px; display: flex; align-items: center; justify-content: space-around;"
    >
        <sp-progress-circle
            label="A small representation of a somewhat completed action"
            progress="71"
            size="s"
        ></sp-progress-circle>
        <sp-progress-circle
            label="A medium representation of a recently started action"
            progress="22"
        ></sp-progress-circle>
        <sp-progress-circle
            label="A large representation of an almost completed action"
            progress="86"
            size="l"
        ></sp-progress-circle>
    </div>
### Static color
If you display your `<sp-progress-cicle>` element over the top of other content, e.g. an image or an alternate background color, it may become appropariate to update the colors with which the circle and loading progress indicator are delivered. To do this, leverage the `static-color` attribute with the value of `white` to ensure the content of your page is being delivered accessibly.
    
    <div
        style="width: 250px; height: 150px; display: flex; align-items: center; justify-content: space-around;  background-color: rgba(0,0,0,0.4);"
    >
        <sp-progress-circle
            label="A small representation of a partially completed action"
            progress="42"
            static-color="white"
            size="s"
        ></sp-progress-circle>
        <sp-progress-circle
            label="A medium representation of a barely started action"
            progress="7"
            static-color="white"
        ></sp-progress-circle>
        <sp-progress-circle
            label="A large representation of a somewhat completed action"
            progress="68"
            static-color="white"
            size="l"
        ></sp-progress-circle>
    </div>
### Indeterminate
A progress circle can be either determinate or indeterminate as signified by `[indeterminate]`. By default, loaders are determinate. Use a determinate loader when progress can be calculated against a specific goal (e.g., downloading a file of a known size). Use an indeterminate loader when progress is happening but the time or effort to completion can’t be determined (e.g., attempting to reconnect to a server).
    
    <div
        style="width: 250px; height: 150px; display: flex; align-items: center; justify-content: space-around;"
    >
        <sp-progress-circle
            label="A small representation of an unclear amount of work"
            indeterminate
            size="s"
        ></sp-progress-circle>
        <sp-progress-circle
            label="A medium representation of an unclear amount of work"
            indeterminate
        ></sp-progress-circle>
        <sp-progress-circle
            label="A large representation of an unclear amount of work"
            indeterminate
            size="l"
        ></sp-progress-circle>
    </div>
### Size
Progress Circles come in 3 sizes: small (`[size="s"]`), medium (`[size="m"]`, default), or large (`[size="l"]`). These are available to fit various contexts. For example, the small loader can be used in place of an icon or in tight spaces, while the large one can be used for full-page loading.
## Accessibility
An `sp-progress-circle` element will register itself as a `role="progressbar"` element in the accessibility tree. Any value applied to the `label` attribute will be set as the `aria-label` attribute on the host. These two attributes can be used interchangably to ensure that the `sp-progress-circle` elements in your UI correctly fulfills its responsibilities to visitors of you site of all abilities.
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
      <td>`indeterminate`</td>
      <td>`indeterminate`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
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
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
  </tbody>
</table>
