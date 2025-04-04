# shared
NPM 1.4.0
Overview
## Description
Shared mixins, tools, etc. that support developing Spectrum Web Components.
### Usage
    
    npm install @spectrum-web-components/shared
Individual base classes and mixins can be imported as follows:
    
    import {
        Focusable,
        FocusVisiblePolyfillMixin,
        getActiveElement,
        LikeAnchor,
        ObserveSlotText,
    } from '@spectrum-web-components/shared';
### getDeepElementFromPoint
The `getDeepElementFromPoint` method allows you to obtain the deepest possible element at a given coordinates on the current page. The method will step into any available `shadowRoot`s until it reaches the first element with no `shadowRoot` or no children available at the given coordinates.
### Focusable
The `Focusable` subclass of `LitElement` adds some helpers method and lifecycle coverage in order to support passing focus to a container element inside of a custom element. The Focusable base class handles tabindex setting into shadowed elements automatically and is based heavily on the aybolit delegate-focus-mixin.
    
    import { Focusable } from '@spectrum-web-components/shared';
    import { html } from 'lit-element';
    
    class FocusableButton extends Focusable {
        public static override get styles(): CSSResultArray {
            return [...super.styles];
        }
        public get focusElement(): HTMLElement {
            return this.shadowRoot.querySelector('#button') as HTMLElement;
        }
    
        protected override render(): TemplateResult {
            return html`
                <button
                    id="button"
                >
                    Focus for this button is being managed by the focusable base class.
                </button>
            `;
        }
    }
### FocusVisiblePolyfillMixin
Use this mixin if you would like to leverage `:focus-visible` based selectors in your CSS. Learn more about the polyfill that powers this.
### getActiveElement
Use this helper to find an `activeElement` in your component. Learn more about tracking active elements over shadow DOM boundaries.
### LikeAnchor
Mix `download`, `label`, `href`, and `target` properties into your element to allow it to act more like an `HTMLAnchorElement`.
### ObserveSlotPresence
When working with styles that are driven by the conditional presence of `<slot>`s in a component's shadow DOM, you will need to track whether light DOM that is meant for that slot exists. Use the `ObserveSlotPresence` mixin to target specific light DOM to observe the presence of and trigger `this.requestUpdate()` calls when content fulfilling that selector comes in and out of availability.
    
    import { ObserveSlotPresence } from '@spectrum-web-components/shared';
    import { LitElement, html } from 'lit-element';
    class ObserveSlotPresenceElement extends ObserveSlotPresence(LitElement, '[slot="conditional-slot"]') {
        // translate the mixin properties into locally understandable language
        protected get hasConditionalSlotContent() {
            return this.slotContentIsPresent;
        }
        protected override render(): TemplateResult {
            return html`
                <button
                    id="button"
                >
                    ${this.hasConditionalSlotContent
                        ? html`
                            <slot
                                name="conditional-slot"
                            ></slot>
                        `
                        : html``
                    }
                </button>
            `;
        }
        protected updated(): void {
            console.log(this.slotContentIsPresent); // => true when <observing-slot-presence-element><div slot="conditional-slot"></div></observing-slot-presence-element>
        }
    }
    customElements.define('observing-slot-presence-element', ObserveSlotPresenceElement);
### ObserveSlotText
When working with `<slot>`s and their `slotchange` event, you will have the opportunity to capture when the nodes and/or elements in your element are added or removed. However, if the `textContent` of a text node changes, you will not receive the `slotchange` event because the slot hasn't actually received new nodes and/or elements in the exchange. When working with a lit-html binding `<your-element>${text}</your-element>` that means you will not receive a `slotchange` event when the value of `text` goes from `text = ''` to `text = 'something'` or the other way. In these cases the `ObserveSlotText` can be leverages to apply a mutation observe onto your element that tracks `characterData` mutations so that you can resspond as desired.
    
    import { ObserveSlotText } from '@spectrum-web-components/shared';
    import { LitElement, html } from 'lit-element';
    
    class ObserveSlotTextElement extends ObserveSlotText(LitElement, '#observing-slot') {
        protected override render(): TemplateResult {
            return html`
                <button
                    id="button"
                >
                    <slot
                        id="observing-slot"
                        @slotchange=${this.manageObservedSlot}
                    ></slot>
                </button>
            `;
        }
        protected updated(): void {
            console.log(this.slotHasContent); // => true when <observing-slot-text-element>Text</observing-slot-text-element>
        }
    }
    
    customElements.define('observing-slot-text-element', ObserveSlotTextElement);
