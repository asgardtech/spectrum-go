# sp-split-view
Overview API
## Description
An `sp-split-view` element delivers its first two direct child elements in a horizontal or vertical (`<sp-split-view vertical>`) orientation that distributes the available page real estate as per the supplied attribute API. When leveraging the resizable attribute a pointer and keyboard accessible affordance is provided for the user to customize the distribution of that area between the available children.
### Usage
    
    yarn add @spectrum-web-components/split-view
    
Import the side effectful registration of `<sp-split-view>` via:
    
    import '@spectrum-web-components/split-view/sp-split-view.js';
    
When looking to leverage the `SplitView` base class as a type and/or for extension purposes, do so via:
    
    import { SplitView } from '@spectrum-web-components/split-view';
    
## Variants
### Horizontal
    
    <sp-split-view>
        <div>Left panel</div>
        <div>Right panel</div>
    </sp-split-view>
### Horizontal Resizable
    
    <sp-split-view
        resizable
        primary-min="50"
        secondary-min="50"
        primary-size="100"
        label="Resize the horizontal panels"
    >
        <div>
            <h1>Left panel</h1>
            <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
            </p>
        </div>
        <div>
            <h2>Right panel</h2>
            <p>
                It is a long established fact that a reader will be distracted by
                the readable content of a page when looking at its layout.
            </p>
        </div>
    </sp-split-view>
### Horizontal Resizable & Collapsible
    
    <sp-split-view resizable label="Resize the horizontal collapsible panels">
        <div>
            <h1>Left panel</h1>
            <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
            </p>
        </div>
        <div>
            <h2>Right panel</h2>
            <p>
                It is a long established fact that a reader will be distracted by
                the readable content of a page when looking at its layout.
            </p>
        </div>
    </sp-split-view>
### Vertical
    
    <sp-split-view vertical>
        <div>Top panel</div>
        <div>Bottom panel</div>
    </sp-split-view>
### Vertical Resizable
    
    <sp-split-view
        vertical
        resizable
        primary-min="50"
        primary-max="150"
        secondary-min="50"
        label="Resize the vertical panels"
    >
        <div>
            <h1>Top panel</h1>
            <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
            </p>
        </div>
        <div>
            <h2>Bottom panel</h2>
            <p>
                It is a long established fact that a reader will be distracted by
                the readable content of a page when looking at its layout.
            </p>
        </div>
    </sp-split-view>
### Vertical Resizable & Collapsible
    
    <sp-split-view
        vertical
        resizable
        style="height: 300px;"
        label="Resize the vertical collapsible panels"
    >
        <div>
            <h1>Top panel</h1>
            <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
            </p>
        </div>
        <div>
            <h2>Bottom panel</h2>
            <p>
                It is a long established fact that a reader will be distracted by
                the readable content of a page when looking at its layout.
            </p>
        </div>
    </sp-split-view>
### Multiple Levels
    
    <sp-split-view
        resizable
        primary-min="50"
        primary-max="200"
        secondary-min="50"
        style="height: 400px; width: 600px;"
    >
        <div>
            <h1>First panel - Level 1</h1>
            <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry. Lorem Ipsum has been the industry's standard dummy text
                ever since the 1500s, when an unknown printer took a galley of type
                and scrambled it to make a type specimen book.
            </p>
        </div>
        <div>
            <h2>Second panel - Level 1</h2>
            <sp-split-view
                vertical
                resizable
                primary-min="50"
                primary-size="100"
                secondary-min="50"
                style="height: 300px;"
            >
                <div>
                    <h3>First panel - Level 2</h3>
                    <p>
                        Lorem Ipsum is simply dummy text of the printing and
                        typesetting industry.
                    </p>
                </div>
                <div>
                    <h4>Second panel - Level 2</h4>
                    <p>
                        It is a long established fact that a reader will be
                        distracted by the readable content of a page when looking at
                        its layout.
                    </p>
                </div>
            </sp-split-view>
        </div>
    </sp-split-view>
## Accessibility
By default, the "separator" element within an `<sp-split-view>` is given the label "Resize the panels". A label is required to surface the interaction correctly to screen readers. You can customize or internationalize this with the `label` attribute.
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
      <td>`collapsible`</td>
      <td>`collapsible`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`primaryMax`</td>
      <td>`primary-max`</td>
      <td>``</td>
      <td>`DEFAULT_MAX_SIZE`</td>
      <td>The maximum size of the primary pane</td>
    </tr>
    <tr>
      <td>`primaryMin`</td>
      <td>`primary-min`</td>
      <td>`number`</td>
      <td>`0`</td>
      <td>The minimum size of the primary pane</td>
    </tr>
    <tr>
      <td>`primarySize`</td>
      <td>`primary-size`</td>
      <td>`number | number`</td>
      <td>``</td>
      <td>+ "px" | number + "%" | "auto"}</td>
    </tr>
    <tr>
      <td>`primarySize`</td>
      <td>`primarySize`</td>
      <td>`number | number`</td>
      <td>``</td>
      <td>+ "px" | number + "%" | "auto"}</td>
    </tr>
    <tr>
      <td>`resizable`</td>
      <td>`resizable`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`secondaryMax`</td>
      <td>`secondary-max`</td>
      <td>``</td>
      <td>`DEFAULT_MAX_SIZE`</td>
      <td>The maximum size of the secondary pane</td>
    </tr>
    <tr>
      <td>`secondaryMin`</td>
      <td>`secondary-min`</td>
      <td>`number`</td>
      <td>`0`</td>
      <td>The minimum size of the secondary pane</td>
    </tr>
    <tr>
      <td>`splitterPos`</td>
      <td>`splitter-pos`</td>
      <td>`number | undefined`</td>
      <td>``</td>
      <td>The current splitter position of split-view</td>
    </tr>
    <tr>
      <td>`vertical`</td>
      <td>`vertical`</td>
      <td>`boolean`</td>
      <td>`false`</td>
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
      <td>`Two`</td>
      <td>sibling elements to be sized by the element attritubes</td>
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
      <td>`Announces the new position of the splitter`</td>
    </tr>
  </tbody>
</table>
