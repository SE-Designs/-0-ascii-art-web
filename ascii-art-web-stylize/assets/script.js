// var textarea = document.querySelector('textarea');
// var p = document.querySelector('pre');

// textarea.addEventListener('input', function(event) {
//   p.innerText = event.target.value;
// });


// -----

function clearOutput() {
    	output = document.getElementById("log").innerHTML = ""
}

// result = document.getElementById("result_api")

// var buffer = "";

// function makeRequest() {
// 	font = document.getElementById('font').value;
// 	alignment = document.getElementById('alignment').value;
// 	input = document.getElementById('text').value;
// 	let text = encodeURIComponent(input);
// 	let width = window.innerWidth;
// 	req = 'fonts=' + font + '&alignment=' + alignment + '&width=' + width + '&text=' + text;
// 	httpRequest = new XMLHttpRequest();
// 	httpRequest.onreadystatechange = function() {
// 		document.getElementById("result").innerHTML = httpRequest.responseText;
// 		window.buffer = httpRequest.responseText;
// 		console.log(buffer);
// 	}
// 	httpRequest.open("POST", '/art');
// 	httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
// 	httpRequest.send(req);
// }
