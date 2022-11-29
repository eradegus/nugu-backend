$(document).ready(function (){
	window.step = 1;
	window.maxStep = 4;

	$("#text-step").text(window.step + " / 4");
	$("#div-input-" + window.step).show();

	$("#btn-next").on("click", handlerBtnNext);

	$("#btn-exit").on("click", handlerBtnExit);

	$("#text-input-4").datepicker({
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
		let specialDay = $("#text-input-4").val();

		// Send input data to server
		console.log(destAddr + " / " + busStop + " / " + busNum  + " / " + zipName  + " / " + specialDay);

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

