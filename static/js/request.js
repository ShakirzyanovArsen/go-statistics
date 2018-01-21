function makeRequstAndPasteDataToTable(view, dateFrom, dateTo) {
    var userJobsCountDiv = document.getElementById(view.divId);

    userJobsCountDiv.innerHTML = "";
    if (!userJobsCountDiv.classList.contains('loader')) {
        userJobsCountDiv.classList.add('loader');
    }
    var uri = view.uri + '?from=' + dateFrom + ':00' + "&to=" + dateTo + ':00';
    if(view.additionalParams !== null) {
        for (var param in view.additionalParams) {
            if(view.additionalParams.hasOwnProperty(param)) {
                uri += "&" + param + "=" + view.additionalParams[param];
            }
        }
    }
    jQuery.ajax({
        url: uri,
        async: true,
        success: function (result) {
            userJobsCountDiv.classList.remove('loader');
            var table = createTable(userJobsCountDiv, result, view.headers);
            userJobsCountDiv.appendChild(table);
        }
    });
}