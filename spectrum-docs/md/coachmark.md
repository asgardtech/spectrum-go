# sp-coachmark
Overview API
## Description
`<sp-coachmark>` is a temporary message that educates users through new or unfamiliar product experiences. They can be chained into a sequence to form a tour.
### Usage
    
    yarn add @spectrum-web-components/coachmark
    
Import the side effectful registration of `<sp-coachmark>` via:
    
    import '@spectrum-web-components/coachmark/sp-coachmark.js';

When looking to leverage the `Coachmark` base class as a type and/or for extension purposes, do so via:
    
    import { Coachmark } from '@spectrum-web-components/coachmark';
    
## Example
## Default
    
    <sp-coachmark open>
        <div slot="title">Coachmark with Text Only</div>
        <div slot="content">
            This is a Coachmark with nothing but text in it. Kind of lonely in here.
        </div>
    </sp-coachmark>
## Using Action Menu
Coach marks can include an `<sp-action-menu>`, which appears at the top right of the coach mark. The `<sp-action-menu>` should only include ways to interact with the coach mark tour as a whole, with options like “Skip tour” or “Restart tour.”
    
    <sp-coachmark
        open
        current-step="2"
        total-steps="8"
        primary-cta="Next"
        secondary-cta="Previous"
    >
        <div slot="title">Coachmark with Text Only</div>
        <div slot="content">
            This is a Coachmark with nothing but text in it. Kind of lonely in here.
        </div>
        <sp-action-menu
            label="More Actions"
            placement="bottom-end"
            quiet
            slot="actions"
        >
            <sp-menu-item>Skip tour</sp-menu-item>
            <sp-menu-item>Restart tour</sp-menu-item>
        </sp-action-menu>
    </sp-coachmark>
## User Action Dependent
User action-dependent coachmarks are designed to guide users based on their interactions within your application. In such cases, there is no "Next Step" button, as the coachmark progresses when the user takes a specific action. This allows users to learn by doing, rather than simply reading instructions. The coachmark remains until the user performs the required action or takes an alternative route in the tour, such as skipping, restarting, or moving back to a previous step.
Inside the `<sp-coachmark>`, add the content and instructions for the coachmark in the `<sp-coachmark>`. You can also define primary and secondary CTA buttons for user interaction.
**Event Handling:**
The primary and secondary CTA buttons within the coachmark popover can be configured to dispatch events when clicked.
    
    <sp-coachmark
        id="coachmark-action"
        open
        current-step="2"
        total-steps="8"
        primary-cta="Asset added"
        secondary-cta="Previous"
    >
        <div slot="title">Coachmark with user action</div>
        <div slot="content">
            This is a Coachmark with nothing but text in it. Kind of lonely in here.
        </div>
        <sp-action-menu label="More Actions" placement="bottom-end" quiet slot="actions">
            <sp-menu-item>Skip tour</sp-menu-item>
            <sp-menu-item>Restart tour</sp-menu-item>
        </sp-action-menu>
    </sp-coachmark>
    <script type="module">
        const initCoachMark = () => {
            const coachmark = document.querySelector('#coachmark-action');
            coachmark.addEventListener('primary', () => console.log('primary call to action'));
            coachmark.addEventListener('secondary', () => console.log('secondary call to action'));
    
        };
        customElements.whenDefined('sp-coachmark').then(() => {
                    initCoachMark();
        });
    </script>
## Using Images, GIFs
Coach marks can contain images or videos that relate to their content, such as demonstrations of gestures, the feature being used, or illustrations. To use these kinds of media in your rich tooltip, specify a `src`, the type of media, either by using the string or `media-type` object, and an optional `imageAlt` text describing the content.
Media Types allowed: `Images & Gifs`
### Image
    
    <sp-coachmark
        current-step="2"
        total-steps="8"
        open
        primary-cta="Next"
        secondary-cta="Previous"
        src="https://picsum.photos/id/237/200/300"
        media-type="image"
    >
        <div slot="title">Coachmark with 16:9 image</div>
        <div slot="content">This is a Coachmark with some description</div>
        <sp-action-menu
            label="More Actions"
            placement="bottom-end"
            quiet
            slot="actions"
        >
            <sp-menu-item>Skip tour</sp-menu-item>
            <sp-menu-item>Restart tour</sp-menu-item>
        </sp-action-menu>
    </sp-coachmark>
### Custom Image/Gif
A custom media can also be added via `<slot name="cover-photo"></slot>`
    
    <sp-coachmark
        current-step="2"
        total-steps="8"
        open
        primary-cta="Next"
        secondary-cta="Previous"
    >
        <div slot="title">Coachmark with 16:9 image</div>
        <div slot="content">This is a Coachmark with some description</div>
        <img slot="asset" src="https://picsum.photos/id/237/200/300" alt="" />
        <sp-action-menu
            label="More Actions"
            placement="bottom-end"
            quiet
            slot="actions"
        >
            <sp-menu-item>Skip tour</sp-menu-item>
            <sp-menu-item>Restart tour</sp-menu-item>
        </sp-action-menu>
    </sp-coachmark>
## Shortcut keys and modifier keys
Shortcut keys and modifier keys are ways to show users how to trigger a particular tool.
The `shortcutKey` is the primary key used to trigger an interaction and are typically an alphanumeric value (and thus will be rendered as an uppercase character), while the `modifierKeys` are an array of `string`s that represent alternate keys that can be pressed, like `Shift`, `Alt`, `Cmd`, etc.
    
        <sp-coachmark
            open
            current-step="2"
            total-steps="8"
            primary-cta="Next"
            secondary-cta="Previous"
            id="coachmark-keys"
        >
            <sp-action-menu label="More Actions" placement="bottom-end" quiet slot="actions">
                <sp-menu-item>Skip tour</sp-menu-item>
                <sp-menu-item>Restart tour</sp-menu-item>
            </sp-action-menu>
        </sp-coachmark>
    <script type="module">
        const initCoachMark = () => {
            const coachmark = document.querySelector('#coachmark-keys');
            const modifierKeys = ['⇧ Shift', '⌘'];
            const content = {
                title: 'I am a Coachmark with keys',
                description: 'This is a Coachmark with nothing but text in it.'
            };
            coachmark.content= content
            coachmark.modifierKeys = modifierKeys
        };
        customElements.whenDefined('code-example').then(() => {
            customElements.whenDefined('sp-coachmark').then(() => {
                    initCoachMark();
            });
        });
    </script>
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
      <td>`asset`</td>
      <td>`asset`</td>
      <td>`'file' | 'folder' | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`currentStep`</td>
      <td>`current-step`</td>
      <td>`number | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`hasAsset`</td>
      <td>`has-asset`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`item`</td>
      <td>`item`</td>
      <td>`CoachmarkItem | undefined`</td>
      <td>``</td>
      <td></td>
    </tr>
    <tr>
      <td>`modifierKeys`</td>
      <td>`modifierKeys`</td>
      <td>`string[] | undefined`</td>
      <td>`[]`</td>
      <td></td>
    </tr>
    <tr>
      <td>`open`</td>
      <td>`open`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether the popover is visible or not.</td>
    </tr>
    <tr>
      <td>`placement`</td>
      <td>`placement`</td>
      <td>`"top" | "top-start" | "top-end" | "right" | "right-start" | "right-end" | "bottom" | "bottom-start" | "bottom-end" | "left" | "left-start" | "left-end"`</td>
      <td>`'right'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`tip`</td>
      <td>`tip`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`totalSteps`</td>
      <td>`total-steps`</td>
      <td>`number | undefined`</td>
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
      <td>`actions`</td>
      <td>an `sp-action-menu` element outlining actions to take on the represened object</td>
    </tr>
    <tr>
      <td>`cover-photo`</td>
      <td>This is the cover photo for Default and Quiet Coachmark</td>
    </tr>
    <tr>
      <td>`description`</td>
      <td>A description of the card</td>
    </tr>
    <tr>
      <td>`heading`</td>
      <td>HTML content to be listed as the heading</td>
    </tr>
    <tr>
      <td>`step-count`</td>
      <td>Override the default pagination delivery with your own internationalized content</td>
    </tr>
    <tr>
      <td>`default slot`</td>
      <td>content to display within the Popover</td>
    </tr>
  </tbody>
</table>
### Events
<table>
  <thead>
  </thead>
  <tbody>
    <tr>
      <td>`primary`</td>
      <td>`Event`</td>
      <td>`Announces that the "primary" button has been clicked.`</td>
    </tr>
    <tr>
      <td>`secondary`</td>
      <td>`Event`</td>
      <td>`Announces that the "secondary" button has been clicked.`</td>
    </tr>
  </tbody>
</table>
