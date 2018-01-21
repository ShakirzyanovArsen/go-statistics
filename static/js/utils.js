function getDate(key) {
    var date =  localStorage.getItem(key);
    date = date === null ? "today" : date;
    return date
}

function makeConfigWithDate(date) {
    return {
        enableTime: true,
        altInput: true,
        maxDate: "today",
        altFormat: "Y-m-d H:i:ss",
        time_24hr: true,
        defaultDate: date
    }
}