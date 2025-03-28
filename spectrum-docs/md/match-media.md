# match-media
NPM 1.4.0
Overview
## Description
The match media API allows for a developer to query the state of a supplied CSS media query from the JS scope while surfacing an event based API to listen for changes to whether that query is currently matched or not. `MatchMediaController` binds the supplied CSS media query to the supplied Reactive Element and calls for an update in the host element when the query goes between matching and not. This allow for the `matches` property on the reactive controller to be leveraged in your render lifecycle.
A `MatchMediaController` can be bound to any of the growing number of CSS media queries and any number of `MatchMediaControllers` can be bound to a host element. With this in mind the `MatchMediaController` can support a wide array of complex layouts.
### Usage
    
    yarn add @spectrum-web-components/reactive-controllers
    
Import the `MatchMediaController` via:
    
    import { MatchMediaController } from '@spectrum-web-components/reactive-controllers/MatchMedia.js';
    
## Example
A `Host` element that renders a different message depending on the "orientation" of the window in which is it delivered:
    
    import { html, LitElement } from 'lit';
    import { MatchMediaController } from '@spectrum-web-components/reactive-controllers/MatchMedia.js';
    
    class Host extends LitElement {
        orientationLandscape = new MatchMediaController(
            this,
            '(orientation: landscape)'
        );
    
        render() {
            if (this.orientationLandscape.matches) {
                return html`
                    The orientation is landscape.
                `;
            }
            return html`
                The orientation is portrait.
            `;
        }
    }
