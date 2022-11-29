$(document).ready(function (){
	window.step = 1;
	window.maxStep = 5;

	$("#text-step").text(window.step + " / " + window.maxStep);
	$("#div-input-" + window.step).show();

	$("#btn-next").on("click", handlerBtnNext);

	$("#btn-exit").on("click", handlerBtnExit);

	$("#text-input-5").datepicker({
		uiLibrary: "bootstrap4",
		format: "yyyy-mm-dd"
	});

});

const handlerBtnNext = function () {
	if (window.step == maxStep) {
		//remove all and show finish message
		$("#div-message-info").hide();
		$("#div-buttons").hide();
		$("#div-inputs").hide();

		$("#div-loading").show();
		let destAddr = $("#text-input-1").val();
		let busStop = $("#text-input-2-1").val();
		let busNum = $("#text-input-2-2").val();
		let zipName = $("#text-input-3").val();
		let stockName = $("#text-input-4").val();
		let specialDay = $("#text-input-5").val();

		// Send input data to server
		let param = {destAddr: destAddr, busStop: busStop, busNum: busNum, zipName: zipName, stockName: stockName, specialDay: specialDay};
		http_util("post","/userdb", JSON.stringify(param), function (code, data) {
			console.log(code, data);
		}, function (code, err) {
			console.log(code, err);
			return;
		});

		setTimeout(function () {
			$("#div-loading").hide();
			$("#div-message-finish").show();
		}, 1000);
	}
	else {
		window.step = window.step + 1;
		$("#text-step").text(window.step + " / " + window.maxStep);
		$("#div-input-" + window.step).show();
	}
}

const handlerBtnExit = function() {
	window.location.reload();
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
