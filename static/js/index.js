$(document).ready(function (){
	window.logLength = 0;
	setInterval(() => getLogStream(), 1000);

	http_util("post", "/clearlog", {uuid: window.uuid}, function (code, data) {
		console.log(code, data);
	}, function(code, err) {
		console.log(code, err);
		return;
	});
});

const getLogStream = function () {
	let divAccessLog = $("#accessLog");
	http_util("get", "/logstream", {uuid: window.uuid}, function (code, data) {
		if (data.log.length != window.logLength) {
			for (let i = window.logLength; i < data.log.length; i++) {
				divAccessLog.append(data.log[i] + "</br>");
			}
			window.logLength = data.log.length;
			scrollToBottom(divAccessLog);
		}
	}, function(code, err) {
		console.log(code, err);
		return;
	});
}

// HTTP request wrapper
const http_util = function (type, url, params, success_handler, error_handler, base_url) {

	if(base_url) {
		url = base_url + url;
	}

	let success = arguments[3]?arguments[3]:function(){};
	let error = arguments[4]?arguments[4]:function(){};

	$.ajax({
		type: type,
		url: url,
		dataType: "json",
		data: params,
		async: false,
		success: function (data, textStatus, xhr) {
			if(textStatus === "success"){
				success(xhr.status, data);   // there returns the status code
			}
		},
		error: function (xhr, error_text, statusText) {
			error(xhr.status, xhr);  // there returns the status code
		}
	});
}


const scrollToBottom = function (targetDiv) {
	targetDiv.scrollTop(targetDiv.height() + 999999999999);
}

