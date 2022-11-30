$(document).ready(function (){
	window.step = 1;
	window.maxStep = 5;

	$("#text-step").text(window.step + " / " + window.maxStep);
	$("#div-input-" + window.step).show();

	$("#btn-skip").on("click", handlerBtnSkip);
	$("#btn-next").on("click", handlerBtnNext);

	$("#btn-exit").on("click", handlerBtnExit);

	$("#text-input-5").datepicker({
		uiLibrary: "bootstrap4",
		format: "yyyy-mm-dd"
	});

	let calendarIcon = $(".gj-icon");
	calendarIcon.text("");
	calendarIcon.addClass("fa-solid fa-calendar-check");
	calendarIcon.removeClass("gj-icon");
});

const handlerBtnSkip = function () {
	if (window.step != maxStep) {
		$("#div-input-" + window.step).hide();
		window.step = window.step + 1;
		$("#text-step").text(window.step + " / " + window.maxStep);
		$("#div-input-" + window.step).show();
	}
}

const handlerBtnNext = function () {
	if (window.step == maxStep) {
		//remove all and show finish message
		$("#div-message-info").hide();
		$("#div-buttons").hide();
		$("#div-inputs").hide();

		$("#div-loading").show();

		let destAddr = !$("#text-input-1").val() ? "서초구" : $("#text-input-1").val();
		let busStop = !$("#text-input-2-1").val() ? "관악경찰서" : $("#text-input-2-1").val();
		let busNum = !$("#text-input-2-2").val() ? "5511" : $("#text-input-2-2").val();
		let zipName = !$("#text-input-3").val() ? "한남더힐" : $("#text-input-3").val();
		let stockName = !$("#text-input-4").val() ? "SK텔레콤" : $("#text-input-4").val();
		let specialDay = !$("#text-input-5").val() ? "2021-12-04" : $("#text-input-5").val();

		// Send input data to server
		let param = {destAddr: destAddr, busStop: busStop, busNum: busNum, zipName: zipName, stockName: stockName, specialDay: specialDay};
		http_util("post","/userdb", JSON.stringify(param), function (code, data) {
			console.log(code, data);
		}, function (code, err) {
			console.log(code, err);
			return;
		});

		// Generate example message
		let exampleGoodmorning = "\"좋은 아침이에요!</br> 현재 관악구 날씨는 15도로 맑아요</br>"
								+ destAddr + "에는 오늘 비가 올 수 있어요</br>"
								+ stockName + "의 어제 종가는 100,000원 이에요</br>"
								+ "어제 " + zipName + "에 새로운 실거래가 발생했어요</br>"
								+ "실거래가는 12억 3456만원 이에요</br>"
								+ "좋은 하루 되세요!\"";
		let exampleSeeya = "\"네! " + busNum + " 버스는 " + busStop + " 정류장에 </br>"
								+ "12분 후 도착예정이에요 </br>"
								+ "다음 버스는 34분 후 도착예정이에요</br>"
								+ "잘 다녀오세요!\"";

		$("#text-example-goodmorning").html(exampleGoodmorning);
		$("#text-example-seeya").html(exampleSeeya);

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
