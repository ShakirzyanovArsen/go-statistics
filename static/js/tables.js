/**
 *
 * @param element where to put table (all tags inside element will be cleared)
 * @param rows table rows
 * @param headers association between data from callback and column header names
 */
function createTable(element, rows, headers) {
    var table = document.createElement('table');
    table.className = 'table table-striped';
    var thead = document.createElement('thead');
    var theadRow = document.createElement('tr');
    headers.forEach(function (header) {
        var th = document.createElement('th');
        th.scope = 'col';
        th.innerText = header.tableName;
        theadRow.appendChild(th);
    });
    thead.appendChild(theadRow);
    table.appendChild(thead);
    var tbody = document.createElement('tbody');
    console.log(rows);
    rows.forEach(function (rowData) {
        var tr = document.createElement('tr');
        headers.forEach(function (header) {
            var td = document.createElement('td');
            td.innerText = rowData[header.name];
            tr.appendChild(td);
        });
        tbody.appendChild(tr)
    });

    table.appendChild(tbody);
    return table;
}