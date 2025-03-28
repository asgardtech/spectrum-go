# sp-action-group
Overview API
## Description
`sp-action-group` delivers a set of action buttons in horizontal or vertical orientation while ensuring the appropriate spacing between those buttons. The `compact` attribute merges these buttons so that they are visually joined to clarify their relationship to each other and their distance from other buttons/groups.
### Usage
    
    yarn add @spectrum-web-components/action-group
    
Import the side effectful registration of `<sp-action-group>` via:
    
    import '@spectrum-web-components/action-group/sp-action-group.js';
    
When looking to leverage the `ActionGroup` base class as a type and/or for extension purposes, do so via:
    
    import { ActionGroup } from '@spectrum-web-components/action-group';
    
## Sizes
Extra Small
    
    <sp-action-group size="xs">
        <sp-action-button>Extra Small 1</sp-action-button>
        <sp-action-button>Extra Small 2</sp-action-button>
    </sp-action-group>
Small
    
    <sp-action-group size="s">
        <sp-action-button>Small 1</sp-action-button>
        <sp-action-button>Small 2</sp-action-button>
    </sp-action-group>
Medium
    
    <sp-action-group size="m">
        <sp-action-button>Medium 1</sp-action-button>
        <sp-action-button>Medium 2</sp-action-button>
    </sp-action-group>
Large
    
    <sp-action-group size="l">
        <sp-action-button>Large 1</sp-action-button>
        <sp-action-button>Large 2</sp-action-button>
    </sp-action-group>
Extra Large
    
    <sp-action-group size="xl">
        <sp-action-button>Extra Large 1</sp-action-button>
        <sp-action-button>Extra Large 2</sp-action-button>
    </sp-action-group>
## Selects
An `<sp-action-group selects="single|multiple">` will manage a `selected` property that consists on an array of the `<sp-action-button>` children that are currently selected. A `change` event is dispatched from the `<sp-action-group>` element when the value of `selected` is updated. This event can be canceled via `event.preventDefault()`, after which the value of `selected` will be returned to what it was previously.
When a selection can be made, it is a good idea to supply the group of options with accessible text that names the group of buttons. This can be done in a non-visual way via the `label` attribute of the `<sp-action-group>` element. You can also associate the `<sp-action-group>` to a second, visible DOM element via the `aria-labelledby` attribute or, when available, via the `for` attribute on the second element to make the association in the other direction.
### Single
An `<sp-action-group selects="single">` will manage its `<sp-action-button>` children as "radio buttons" allowing the user to select a *single* one of the buttons presented. The `<sp-action-button>` children will only be able to turn their `selected` value on while maintaining a single selection after an initial selection is made.
    
    <sp-action-group
        selects="single"
        emphasized
        label="Single Selection Demo Group"
    >
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button selected>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
### Multiple
An `<sp-action-group selects="multiple">` will manage its `<sp-action-button>` children as "checkboxes" allowing the user to select a *multiple* of the buttons presented. The `<sp-action-button>` children will toggle their `selected` value on and off when clicked successively.
    
    <sp-action-group
        selects="multiple"
        emphasized
        label="Multiple Selection Demo Group"
    >
        <sp-action-button selected>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button selected>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
## Selected
The `selected` property represents the selection state within a button group. This property can be managed either by the component or by the user. Passing in an array of button values will make `<sp-action-group>` a controllable component. Though `selected` would more commonly be set via Javascript expressions (i.e. `<sp-action-group .selected=${["first"]}>`), it is also possible to set `selected` as a JSON string.
    
    <sp-action-group selects="single" selected='["first"]'>
        <sp-action-button value="first">First</sp-action-button>
        <sp-action-button value="second">Second</sp-action-button>
    </sp-action-group>
By default, an `<sp-action-group>` will select any button passed into `selected`. Afterwards, `.selects` controls how button values are added to the selection state. For example, if `.selects` is not specified when `selected` is set, any further interaction will result in no change to the selection.
    
    <sp-action-group selected='["first", "second"]'>
        <sp-action-button value="first">First</sp-action-button>
        <sp-action-button value="second">Second</sp-action-button>
        <sp-action-button value="third">Third</sp-action-button>
    </sp-action-group>
Similarly, if `selected` contains more than one button value, but `selects = "single"`, then those initial buttons will be highlighted, but further interaction will result in radio-button functionality.
    
    <sp-action-group selects="single" selected='["first", "second"]'>
        <sp-action-button value="first">First</sp-action-button>
        <sp-action-button value="second">Second</sp-action-button>
        <sp-action-button value="third">Third</sp-action-button>
    </sp-action-group>
## Horizontal
By default, an `<sp-action-group>` will organize its child buttons horizontally and the delivery of those buttons can be modified with the `compact`, `emphasized`, or `quiet` attributes.
    
    <sp-action-group>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group compact>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group quiet>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group compact emphasized>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </sp-action-group>
## Vertical
The use of the `vertical` attribute instructs the `<sp-action-group>` element to organize its child buttons vertically, while accepting the same `compact`, `emphasized`, and `quiet` attributes as modifiers.
    
    <div style="display: flex; justify-content: space-around;">
        <sp-action-group vertical>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Button 1
            </sp-action-button>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Longer Button 2
            </sp-action-button>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Short 3
            </sp-action-button>
        </sp-action-group>
        <sp-action-group vertical compact>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Button 1
            </sp-action-button>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Longer Button 2
            </sp-action-button>
            <sp-action-button>
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
                Short 3
            </sp-action-button>
        </sp-action-group>
        <sp-action-group vertical quiet>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
        </sp-action-group>
        <sp-action-group compact vertical>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
            <sp-action-button label="Zoom in">
                <sp-icon-magnify slot="icon"></sp-icon-magnify>
            </sp-action-button>
        </sp-action-group>
    </div>
## Justified
The `justified` attribute will cause the `<sp-action-group>` element to fill the available horizontal space and evenly distribute that space across its child button elements.
    
    <sp-action-group justified>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group justified compact>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Button 1
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Longer Button 2
        </sp-action-button>
        <sp-action-button>
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
            Short 3
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group justified quiet>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </sp-action-group>
    <br />
    <sp-action-group compact justified>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
        <sp-action-button label="Zoom in">
            <sp-icon-magnify slot="icon"></sp-icon-magnify>
        </sp-action-button>
    </sp-action-group>
## Accessibility
The accessibility `role` for an `<sp-action-group>` element depends on the manner in which items are selected. By default, `<sp-action-group selects="single">` will have `role="radiogroup"`, because it manages its children as a "radio group", while `<sp-action-group>` or `<sp-action-group selects="multiple">` will have `role="toolbar"`, which makes sense for a group of buttons or checkboxes between which one can navigate using the arrow keys.
When more than one `<sp-action-group>` elements are combined together with in a toolbar, the `role` attribute for `<sp-action-group>` or `<sp-action-group selects="multiple">` should be overwritten using `role="group"` or `role="presentation"`, so that toolbars are not nested, as demonstrated in the following example of a hypothetical toolbar for formatting text within a rich text editor:
    
    <div
        aria-label="Text Formatting"
        role="toolbar"
        style="height: 32px; display: flex; gap: 6px"
    >
        <sp-action-group
            aria-label="Text Style"
            selects="multiple"
            role="group"
            compact
            emphasized
        >
            <sp-action-button label="Bold" value="bold">
                <sp-icon-text-bold slot="icon"></sp-icon-text-bold>
            </sp-action-button>
            <sp-action-button label="Italic" value="italic">
                <sp-icon-text-italic slot="icon"></sp-icon-text-italic>
            </sp-action-button>
            <sp-action-button label="Underline" value="underline">
                <sp-icon-text-underline slot="icon"></sp-icon-text-underline>
            </sp-action-button>
        </sp-action-group>
        <sp-divider
            size="s"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-group
            aria-label="Text Align"
            selects="single"
            compact
            emphasized
        >
            <sp-action-button label="Left" value="left" selected>
                <sp-icon-text-align-left slot="icon"></sp-icon-text-align-left>
            </sp-action-button>
            <sp-action-button label="Center" value="center">
                <sp-icon-text-align-center slot="icon"></sp-icon-text-align-center>
            </sp-action-button>
            <sp-action-button label="Right" value="right">
                <sp-icon-text-align-right slot="icon"></sp-icon-text-align-right>
            </sp-action-button>
            <sp-action-button label="Justify" value="justify">
                <sp-icon-text-align-justify
                    slot="icon"
                ></sp-icon-text-align-justify>
            </sp-action-button>
        </sp-action-group>
        <sp-divider
            size="s"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-group
            aria-label="List Style"
            selects="multiple"
            role="group"
            compact
            emphasized
        >
            <sp-action-button
                label="Bulleted"
                value="bulleted"
                onclick="
                    /* makes mutually exclusive checkbox */
                    this.selected &&
                        requestAnimationFrame(() => this.parentElement.selected = []);
                    this.parentElement.selected = [];
                "
            >
                <sp-icon-text-bulleted slot="icon"></sp-icon-text-bulleted>
            </sp-action-button>
            <sp-action-button
                label="Numbering"
                value="numbering"
                onclick="
                    /* makes mutually exclusive checkbox */
                    this.selected && 
                        requestAnimationFrame(() => this.parentElement.selected = []);
                    this.parentElement.selected = [];
                "
            >
                <sp-icon-text-numbered slot="icon"></sp-icon-text-numbered>
            </sp-action-button>
        </sp-action-group>
        <sp-divider
            size="s"
            style="align-self: stretch; height: auto;"
            vertical
        ></sp-divider>
        <sp-action-group role="presentation" compact>
            <sp-action-button disabled label="Copy" value="copy">
                <sp-icon-copy slot="icon"></sp-icon-copy>
            </sp-action-button>
            <sp-action-button disabled label="Paste" value="paste">
                <sp-icon-paste slot="icon"></sp-icon-paste>
            </sp-action-button>
            <sp-action-button disabled label="Cut" value="cut">
                <sp-icon-cut slot="icon"></sp-icon-cut>
            </sp-action-button>
        </sp-action-group>
    </div>
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
      <td>`compact`</td>
      <td>`compact`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`justified`</td>
      <td>`justified`</td>
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
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`selects`</td>
      <td>`selects`</td>
      <td>`undefined | 'single' | 'multiple'`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`staticColor`</td>
      <td>`static-color`</td>
      <td>`'white' | 'black' | undefined`</td>
      <td>``</td>
      <td></td>
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
      <td>`default slot`</td>
      <td>the sp-action-button elements that make up the group</td>
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
      <td>`Announces that selection state has been changed by user`</td>
    </tr>
  </tbody>
</table>
