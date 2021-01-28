function getWebServiceStatus(url, callback) {
    $.ajax({
        type: "GET",
        url: url,
        complete: function (a) {
            callback(a.status, a.responseJSON);
        }
    })
}

function getWebService(url, callback) {
    $.ajax({
        type: "GET",
        url: url,
        complete: function (a) {
            if (a.status == 200) {
                callback(a.responseJSON);
            } else {
                console.log("call api fail, status:" + a.status)
            }
        }
    })
}

function deleteWebService(url, callback) {
    $.ajax({
        type: "DELETE",
        url: url,
        complete: function (a) {
            if (a.status == 200) {
                callback(a.responseJSON);
            } else {
                console.log("call api fail, status:" + a.status)
            }
        }
    })
}

function postWebService(url, data, callback) {
    $.ajax({
        type: "POST",
        url: url,
        data: JSON.stringify(data),
        dataType: "json",
        complete: function (a) {
            callback(a.status, a.responseJSON);
        }
    })
}

function postFormWebService(url, form, callback) {
    $.ajax({
        url: url,
        type: "POST",
        data: form,
        processData: false,
        contentType: false,
        complete: function (a) {
            if (a.status == 200) {
                callback(a.responseJSON);
            } else {
                console.log("call api fail, status:" + a.status)
            }
        }
    })
}

function postFormWebServiceStatus(url, form, callback) {
    $.ajax({
        url: url,
        type: "POST",
        data: form,
        processData: false,
        contentType: false,
        complete: function (a) {
            callback(a.status, a.responseJSON);
        }
    })
}

String.format = function () {
    if (arguments.length == 0)
        return null;
    var str = arguments[0];
    for (var i = 1; i < arguments.length; i++) {
        var re = new RegExp('\\{' + (i - 1) + '\\}', 'gm');
        str = str.replace(re, arguments[i]);
    }
    return str;
};