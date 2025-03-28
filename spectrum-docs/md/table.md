# sp-table
Overview API
## Description
An `<sp-table>` is used to create a container for displaying information. It allows users to sort, compare, and take action on large amounts of data.
### Usage
    
    yarn add @spectrum-web-components/table
    
Import the side effectful registration of `<sp-table>`, `<sp-table-body>`, `<sp-table-cell>`, `<sp-table-checkbox-cell>`, `<sp-table-head>`, `<sp-table-head-cell>`. and `<sp-table-row>` via:
    
    import '@spectrum-web-components/table/elements.js';
    
Or individually via:
    
    import '@spectrum-web-components/table/sp-table.js';
    import '@spectrum-web-components/table/sp-table-body.js';
    import '@spectrum-web-components/table/sp-table-cell.js';
    import '@spectrum-web-components/table/sp-table-checkbox-cell.js';
    import '@spectrum-web-components/table/sp-table-head.js';
    import '@spectrum-web-components/table/sp-table-head-cell.js';
    import '@spectrum-web-components/table/sp-table-row.js';
    
When looking to leverage the `Table`, `TableBody`, `TableCell`, `TableCheckboxCell`, `TableHead`, `TableHeadCell`, or `TableRow` base classes as a type and/or for extension purposes, do so via:
    
    import {
        Table,
        TableBody,
        TableCell,
        TableCheckboxCell,
        TableHead,
        TableHeadCell,
        TableRow
    } from '@spectrum-web-components/table';
    
## Example
    
    <sp-table>
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
            <sp-table-row>
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
            </sp-table-row>
            <sp-table-row>
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Selection
To manage selection on an `<sp-table>`, utilise the `selects` attribute on `<sp-table>`. Each `<sp-table-row>` has a `value` attribute which, by default, corresponds to its index in the table, and these `value`s tell `<sp-table>` which `<sp-table-row>`s are selected. The selected items can be manually applied via the `selected` property on the table.
### `selects="single"`
When `selects="single"`, the `<sp-table>` will manage a *single* selection in the array value of `selected`.
    
    <sp-table
        selects="single"
        selected='["row1"]'
        onchange="spAlert(this, `Selected: ${JSON.stringify(this.selected)}`)"
    >
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row4">
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row5">
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
### `selects="multiple"`
When `selects="multiple"`, the `<sp-table>` manages selections via a presence toggle and adds them to the `selected` array. Additionally, an `<sp-table-checkbox-cell>` will be made available in the `<sp-table-head>` in order to select/deselect all items in the `<sp-table>`.
    
    <sp-table
        selects="multiple"
        selected='["row1", "row2"]'
        onchange="spAlert(this, `Selected: ${JSON.stringify(this.selected)}`)"
    >
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row4">
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
                <sp-table-cell>Row Item Delta</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row5">
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
                <sp-table-cell>Row Item Echo</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Emphasized
Use the `emphasized` attribute to add priority to the information that is delivered within your `<table>` element. In particular, this affects the appearance of selected rows, and will set the emphasized style for the checkboxes within `sp-table-checkbox-cell`. Leaving off the `emphasized` attribute will display the non-emphasized colors.
    
    <sp-table emphasized selects="multiple" selected='["row1"]'>
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Density
The optional `density` property changes the spacing around table cell content from the "regular" default. It accepts the values of `compact` or `spacious`.
Compact
    
    <sp-table density="compact">
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
Spacious
    
    <sp-table density="spacious">
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Sizes
Small
    
    <sp-table size="s">
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
Medium (Default)
    
    <sp-table>
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
Large
    
    <sp-table size="l">
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
Extra Large
    
    <sp-table size="xl">
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Quiet
When using the `quiet` property, the overall look of the table will change. The quiet variant of Table has a transparent background and no borders on the left and right.
    
    <sp-table quiet>
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
        <sp-table-body>
            <sp-table-row value="row1">
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
                <sp-table-cell>Row Item Alpha</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row2">
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
                <sp-table-cell>Row Item Bravo</sp-table-cell>
            </sp-table-row>
            <sp-table-row value="row3">
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
                <sp-table-cell>Row Item Charlie</sp-table-cell>
            </sp-table-row>
        </sp-table-body>
    </sp-table>
## Virtualized Table
For large amounts of data, the `<sp-table>` can be virtualised to easily add table rows by using properties.
    
    <sp-table
        id="table-virtualized-demo"
        style="height: 200px"
        scroller="true"
    >
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
    </sp-table>
    <script type="module">
        const initItems = (count) => {
            const total = count;
            const items = [];
            while (count) {
                count--;
                items.push({
                    name: String(total - count),
                    date: count,
                });
            }
            return items;
        };
        const initTable = () => {
            const table = document.querySelector('#table-virtualized-demo');
            table.items = initItems(50);
    
            table.renderItem = (item, index) => {
                const cell1 = document.createElement('sp-table-cell');
                const cell2 = document.createElement('sp-table-cell');
                const cell3 = document.createElement('sp-table-cell');
                cell1.textContent = `Row Item Alpha ${item.name}`;
                cell2.textContent = `Row Item Alpha ${index}`;
                cell3.textContent = `Last Thing`;
                return [cell1, cell2, cell3];
            }
        };
        customElements.whenDefined('sp-table').then(() => {
            initTable();
        });
    </script>
### How to use it
The virtualised table takes `items` as either a property or a JSON-encoded string, an array of type `Record`, where the key is a `string` and the value can be whatever you'd like. `items` is then fed into the `renderItem` method, which takes an `item` and its `index` as parameters and renders the `<sp-table-row>` for each item. An example is as follows:
    
    const renderItem = (item: Item, index: number): TemplateResult => {
        return html`
            <sp-table-cell>Rowsaa Item Alpha ${item.name}</sp-table-cell>
            <sp-table-cell>Row Item Alpha ${item.date}</sp-table-cell>
            <sp-table-cell>Row Item Alpha ${index}</sp-table-cell>
        `;
    };
`renderItem` is then included as a property of `<sp-table>`, along with the `items`, to render a full `<sp-table>` without excessive manual HTML writing.
You can also render a different cell at a particular index by doing something like below:
    
    const renderItem = (item: Item, index: number): TemplateResult => {
        if (index === 15) {
            return html`
                <sp-table-cell style="text-align: center">Custom Row</sp-table-cell>
            `;
        }
        return html`
            <sp-table-cell>Row Item ${item.name}</sp-table-cell>
            <sp-table-cell>Row Item ${item.date}</sp-table-cell>
            <sp-table-cell>Row Item ${index}</sp-table-cell>
        `;
    };
Please note that there is a bug when attempting to select all virtualised elements. The items are selected programatically, it's just not reflected visually.
### Selection
By default the `selected` property will surface an array of item indexes that are currently selected. However, when making a selection on a virtualized table, it can be useful to track selection as something other than indexes. To do so, set a custom method for the `itemValue` property. The `itemValue` method accepts an item and its index as arguments and should return the value you would like to track in the `selected` property.
    
    <sp-table
        id="table-item-value-demo"
        style="height: 200px"
        scroller="true"
        selects="multiple"
    >
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
    </sp-table>
    <div class="selection">Selected: [ ]</div>
    <script type="module">
        const initItems = (count) => {
            const total = count;
            const items = [];
            while (count) {
                count--;
                items.push({
                    id: crypto.randomUUID(),
                    name: String(total - count),
                    date: count,
                });
            }
            return items;
        };
        const initTable = () => {
            const table = document.querySelector('#table-item-value-demo');
            table.items = initItems(50);
    
            table.renderItem = (item, index) => {
                const cell1 = document.createElement('sp-table-cell');
                const cell2 = document.createElement('sp-table-cell');
                const cell3 = document.createElement('sp-table-cell');
                cell1.textContent = `Row Item Alpha ${item.name}`;
                cell2.textContent = `Row Item Alpha ${index}`;
                cell3.textContent = `Last Thing`;
                return [cell1, cell2, cell3];
            };
    
            table.addEventListener('change', (event) => {
                const selected = event.target.nextElementSibling;
                selected.textContent = `Selected: ${JSON.stringify(event.target.selected, null, ' ')}`;
            });
        };
        customElements.whenDefined('sp-table').then(() => {
            initTable();
        });
    </script>
### Row Types
All values in the item array are assumed to be homogenous by default. This means all of the rendered rows are either delivered as provided, or, in the case you are leveraging `selects`, rendered with an `<sp-table-checkbox-cell>`. However, when virtualizing a table with selection, it can sometimes be useful to surface rows with additional interactions, e.g. "Load more data" links. To support that, you can optionally include the `_$rowType$` brand in your item. The values for this are outlined by the `RowType` enum and include `ITEM` (0) and `INFORMATION` (1). When `_$rowType$: RowType.INFORMATION` is provided, it instructs the `<sp-table>` not to deliver an `<sp-table-checkbox-cell>` in that row.
    
    <sp-table
        id="table-row-type-demo"
        style="height: 200px"
        scroller="true"
        selects="single"
    >
        <sp-table-head>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
            <sp-table-head-cell>Column Title</sp-table-head-cell>
        </sp-table-head>
    </sp-table>
    <script type="module">
        const initItems = (count) => {
            const total = count;
            const items = [];
            while (count) {
                count--;
                items.push({
                    name: String(total - count),
                    date: count,
                });
            }
            return items;
        };
        const initTable = () => {
            const table = document.querySelector('#table-row-type-demo');
            const items = initItems(50);
            items.splice(3, 0, {
                _$rowType$: 1,
            });
            table.items = items;
    
            table.renderItem = (item, index) => {
                if (item._$rowType$ === 1) {
                    const infoCell = document.createElement('sp-table-cell');
                    infoCell.textContent = 'Use this row type for non-selectable content.';
                    return [infoCell];
                }
                const cell1 = document.createElement('sp-table-cell');
                const cell2 = document.createElement('sp-table-cell');
                const cell3 = document.createElement('sp-table-cell');
                cell1.textContent = `Row Item Alpha ${item.name}`;
                cell2.textContent = `Row Item Alpha ${index}`;
                cell3.textContent = `Last Thing`;
                return [cell1, cell2, cell3];
            };
        };
        customElements.whenDefined('sp-table').then(() => {
            initTable();
        });
    </script>
### The `scroller` property
By default, the virtualized table doesn't contain a scroll bar and will display the entire length of the table body. Use the `scroller` property and specify an inline style for the height to get a `Table` of your desired height that scrolls.
## Sorting on the Virtualized Table
The virtualized table supports sorting its elements.
For each table column you want to sort, use the `sortable` attribute in its respective `<sp-table-head-cell>`. `sort-direction="asc"|"desc"` specifies the direction the sort goes, in either ascending or descending order, respectively. The `@sorted` event listener on `<sp-table>` can be utilised to specify a method to fire when the `<sp-table-head-cell>` dispatches the `sorted` event. To specify which aspect of an item you'd like to sort by, use the `sort-key` attribute.
    
    <sp-table
        id="sorted-virtualized-table"
        style="height: 200px"
        scroller="true"
    >
        <sp-table-head>
            <sp-table-head-cell sortable sort-direction="desc" sort-key="name">
                Sortable Column
            </sp-table-head-cell>
            <sp-table-head-cell>Non-sortable Column</sp-table-head-cell>
            <sp-table-head-cell>Non-sortable Column</sp-table-head-cell>
        </sp-table-head>
    </sp-table>
    <script type="module">
        const initItems = (count) => {
            const total = count;
            const items = [];
            while (count) {
                count--;
                items.push({
                    name: String(total - count),
                    date: count,
                });
            }
            return items;
        }
    
        let items = initItems(50);
    
        const initTable = () => {
            const table = document.querySelector('#sorted-virtualized-table');
    
            table.items = items;
    
            table.renderItem = (item, index) => {
                const cell1 = document.createElement('sp-table-cell');
                const cell2 = document.createElement('sp-table-cell');
                const cell3 = document.createElement('sp-table-cell');
                cell1.textContent = `Row Item Alpha ${item.name}`;
                cell2.textContent = `Row Item Beta ${item.date}`;
                cell3.textContent = `Index: ${index}`;
                return [cell1, cell2, cell3];
            }
            table.addEventListener('sorted', (event) => {
                const { sortDirection, sortKey } = event.detail;
                items = items.sort((a, b) => {
                    const first = String(a[sortKey]);
                    const second = String(b[sortKey]);
                    return sortDirection === 'asc'
                        ? first.localeCompare(second)
                        : second.localeCompare(first);
                });
                table.items = [...items];
            });
        };
    
        customElements.whenDefined('sp-table').then(() => {
            initTable();
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
      <td>`density`</td>
      <td>`density`</td>
      <td>`'compact' | 'spacious' | undefined`</td>
      <td>``</td>
      <td>Changes the spacing around table cell content.</td>
    </tr>
    <tr>
      <td>`emphasized`</td>
      <td>`emphasized`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Deliver the Table with additional visual emphasis to selected rows.</td>
    </tr>
    <tr>
      <td>`itemValue`</td>
      <td>`itemValue`</td>
      <td>``</td>
      <td>``</td>
      <td>The value of an item. By default, it is set to the index of the sp-table-row.</td>
    </tr>
    <tr>
      <td>`items`</td>
      <td>`items`</td>
      <td>`Record<string, unknown>[]`</td>
      <td>`[]`</td>
      <td>The content of the rows rendered by the virtualized table. The key is the value of the sp-table-row, and the value is the sp-table-row's content (not the row itself).</td>
    </tr>
    <tr>
      <td>`quiet`</td>
      <td>`quiet`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Display with "quiet" variant styles.</td>
    </tr>
    <tr>
      <td>`role`</td>
      <td>`role`</td>
      <td>`string`</td>
      <td>`'grid'`</td>
      <td></td>
    </tr>
    <tr>
      <td>`scroller`</td>
      <td>`scroller`</td>
      <td>`boolean`</td>
      <td>`false`</td>
      <td>Whether or not the virtualized table has a scroll bar. If this is set to true, make sure to specify a height in the sp-table's inline styles.</td>
    </tr>
    <tr>
      <td>`selected`</td>
      <td>`selected`</td>
      <td>`string[]`</td>
      <td>`[]`</td>
      <td>An array of</td>
    </tr>
    <tr>
      <td>`selects`</td>
      <td>`selects`</td>
      <td>`undefined | 'single' | 'multiple'`</td>
      <td>``</td>
      <td>Whether the Table allows users to select a row or rows, and thus controls whether or not the Table also renders checkboxes.</td>
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
      <td>`Announces a change in the `selected` property of a table row`</td>
    </tr>
    <tr>
      <td>`undefined`</td>
      <td>`RangeChangedEvent`</td>
      <td>``</td>
    </tr>
    <tr>
      <td>`rangeChanged`</td>
      <td>`Event`</td>
      <td>`Announces a change in the range of visible cells on the table body`</td>
    </tr>
  </tbody>
</table>
