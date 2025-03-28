# sp-card
Overview API
## Overview
An `<sp-card>` represents a rectangular card that contains a variety of text and image layouts. Cards are typically used to encapsulate units of a data set, such as a gallery of image/caption pairs.
### Usage
    
    yarn add @spectrum-web-components/card
    
Import the side effectful registration of `<sp-card>` via:
    
    import '@spectrum-web-components/card/sp-card.js';
    
When looking to leverage the `Card` base class as a type and/or for extension purposes, do so via:
    
    import { Card } from '@spectrum-web-components/card';
    
### Anatomy
Normal cards can contain a heading, a subheading, a cover photo, and a footer.
    
    <sp-card heading="Card Heading">
        <img alt="" slot="cover-photo" src="https://picsum.photos/250/300" />
        <span slot="subheading">JPG photo</span>
        <div slot="footer">Footer</div>
    </sp-card>
#### Heading
By default, the heading for an `<sp-card>` is applied via the `heading` attribute, which is restricted to string content only. For HTML content, use the `heading` slot instead.
attribute
    
    <sp-card
        heading="Card Heading"
        subheading="JPG Photo"
        style="--spectrum-card-body-header-height: auto;"
    >
        <img alt="" slot="cover-photo" src="https://picsum.photos/250/300" />
        <div slot="footer">Footer</div>
    </sp-card>
slot
    
    <sp-card
        subheading="JPG Photo"
        style="--spectrum-card-body-header-height: auto;"
    >
        <h1 slot="heading">Card Heading</h1>
        <img alt="" slot="cover-photo" src="https://picsum.photos/250/300" />
        <div slot="footer">Footer</div>
    </sp-card>
#### Linking
An `<sp-card>` can be provided with an `href` attribute in order for it to act as one large anchor element. When leveraging the `href` attribute, the `download`, `target` and `rel` attributes customize the card's linking behavior. Use them as follows:
    
    <sp-card
        heading="Card Title"
        subheading="JPG"
        href="https://opensource.adobe.com/spectrum-web-components"
        target="_blank"
    >
        <img
            slot="cover-photo"
            src="https://picsum.photos/200/300"
            alt="Demo Image"
        />
    </sp-card>
#### Cover Photo
Use `slot="cover-photo"` on an image to set it as the card's cover photo.
    
    <sp-card heading="Card Heading" subheading="JPG Photo">
        <img
            slot="cover-photo"
            src="https://picsum.photos/200/300"
            alt="Demo Image"
        />
        <div slot="footer">Footer</div>
    </sp-card>
#### Preview Image
Use `slot="preview"` on an image to set it as the card's preview image.
    
    <sp-card heading="Card Title" subheading="JPG">
        <img slot="preview" src="https://picsum.photos/200/300" alt="Demo Image" />
        <div slot="footer">Footer</div>
    </sp-card>
#### No Preview Image
Cards with no preview image can contain a heading, a subheading, and a footer.
    
    <sp-card heading="Card Title" subheading="No Preview Image">
        <div slot="footer">Footer</div>
    </sp-card>
#### Actions
Cards can be supplied an `actions` via a names slot.
    
    <sp-card heading="Card Heading" subheading="JPG Photo">
        <img
            slot="cover-photo"
            src="https://picsum.photos/200/300"
            alt="Demo Image"
        />
        <div slot="footer">Footer</div>
        <sp-action-menu
            label="More Actions"
            slot="actions"
            placement="bottom-end"
            quiet
        >
            <sp-menu-item>Deselect</sp-menu-item>
            <sp-menu-item>Select Inverse</sp-menu-item>
            <sp-menu-item>Feather...</sp-menu-item>
            <sp-menu-item>Select and Mask...</sp-menu-item>
            <sp-menu-divider></sp-menu-divider>
            <sp-menu-item>Save Selection</sp-menu-item>
            <sp-menu-item disabled>Make Work Path</sp-menu-item>
        </sp-action-menu>
    </sp-card>
### Options
#### Horizontal
Cards with a `horizontal` attribute can contain a heading, a subheading, a cover photo, and a description.
    
    <sp-card horizontal heading="Card Heading" subheading="JPG Photo">
        <img alt="" slot="cover-photo" src="https://picsum.photos/200/250" />
        <div slot="description">10/15/18</div>
    </sp-card>
#### Variant
There are multiple card variants to choose from in Spectrum. The `variant` attribute controls the main variant of the card.
Cards with `variant="quiet"` can contain a heading, a subheading, a cover photo, a description, and a footer. Quiet cards will also accept `actions` via a named slot.
    
    <sp-card variant="quiet" heading="Card Heading" subheading="JPG Photo">
        <img alt="" slot="preview" src="https://picsum.photos/200/350" />
        <div slot="description">10/15/18</div>
        <sp-action-menu
            label="More Actions"
            slot="actions"
            placement="bottom-end"
            quiet
        >
            <sp-menu-item>Deselect</sp-menu-item>
            <sp-menu-item>Select Inverse</sp-menu-item>
            <sp-menu-item>Feather...</sp-menu-item>
            <sp-menu-item>Select and Mask...</sp-menu-item>
            <sp-menu-divider></sp-menu-divider>
            <sp-menu-item>Save Selection</sp-menu-item>
            <sp-menu-item disabled>Make Work Path</sp-menu-item>
        </sp-action-menu>
    </sp-card>
Cards with `variant="gallery"` can contain a heading, a subheading, an image preview, a description, and a footer.
    
    <sp-card variant="gallery" heading="Card Heading" subheading="JPG Photo">
        <img alt="" slot="preview" src="https://picsum.photos/532/192" />
        <div slot="description">10/15/18</div>
        <div slot="footer">Footer</div>
    </sp-card>
#### Asset
When leveraging the `asset` attribute, a card can be declared as representing a `file` or a `folder`:
    
    <sp-card heading="Card Heading" subheading="JPG Photo" asset="file">
        <div slot="heading">File Name</div>
        <div slot="description">10/15/18</div>
        <div slot="footer">Footer</div>
    </sp-card>
    <sp-card subheading="JPG Photo" asset="folder">
        <div slot="heading">Folder Name</div>
        <div slot="description">10/15/18</div>
        <div slot="footer">Footer</div>
    </sp-card>
### Accessibility
#### Be concise
Heading text should be no more than 5-7 words. If the card has an `href`, the heading is used as a link and should ideally be no more than 3 words. For buttons, 1-2 words.
#### Use descriptive heading, link, and button text
Be descriptive. Set expectations on what someone will find and where they will go once they interact with a card. Avoid using the same text on more than one interactive element, unless both elements go to the same place.
#### Make the first word in a heading meaningful
Consider making the first word of links, buttons and headings something an assistive technology user might search for when headings and links are listed alphabetically.
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
      <td>`download`</td>
      <td>`download`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>Causes the browser to treat the linked URL as a download.</td>
    </tr>
    <tr>
      <td>`focused`</td>
      <td>`focused`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`heading`</td>
      <td>`heading`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`horizontal`</td>
      <td>`horizontal`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`href`</td>
      <td>`href`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The URL that the hyperlink points to.</td>
    </tr>
    <tr>
      <td>`label`</td>
      <td>`label`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>An accessible label that describes the component. It will be applied to aria-label, but not visually rendered.</td>
    </tr>
    <tr>
      <td>`referrerpolicy`</td>
      <td>`referrerpolicy`</td>
      <td>`| 'no-referrer' | 'no-referrer-when-downgrade' | 'origin' | 'origin-when-cross-origin' | 'same-origin' | 'strict-origin' | 'strict-origin-when-cross-origin' | 'unsafe-url' | undefined`</td>
      <td>``</td>
      <td>How much of the referrer to send when following the link.</td>
    </tr>
    <tr>
      <td>`rel`</td>
      <td>`rel`</td>
      <td>`string | undefined`</td>
      <td>``</td>
      <td>The relationship of the linked URL as space-separated link types.</td>
    </tr>
    <tr>
      <td>`subheading`</td>
      <td>`subheading`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`target`</td>
      <td>`target`</td>
      <td>`'_blank' | '_parent' | '_self' | '_top' | undefined`</td>
      <td>``</td>
      <td>Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).</td>
    </tr>
    <tr>
      <td>`toggles`</td>
      <td>`toggles`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td></td>
    </tr>
    <tr>
      <td>`value`</td>
      <td>`value`</td>
      <td>`string`</td>
      <td>`''`</td>
      <td></td>
    </tr>
    <tr>
      <td>`variant`</td>
      <td>`variant`</td>
      <td>`'standard' | 'gallery' | 'quiet'`</td>
      <td>`'standard'`</td>
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
      <td>This is the cover photo for Default and Quiet Cards</td>
    </tr>
    <tr>
      <td>`description`</td>
      <td>A description of the card</td>
    </tr>
    <tr>
      <td>`footer`</td>
      <td>Footer text</td>
    </tr>
    <tr>
      <td>`heading`</td>
      <td>HTML content to be listed as the heading</td>
    </tr>
    <tr>
      <td>`preview`</td>
      <td>This is the preview image for Gallery Cards</td>
    </tr>
    <tr>
      <td>`subheading`</td>
      <td>HTML content to be listed as the subheading</td>
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
      <td>`Announces a change in the `selected` property of a card`</td>
    </tr>
    <tr>
      <td>`click`</td>
      <td>`Event`</td>
      <td>``</td>
    </tr>
  </tbody>
</table>
