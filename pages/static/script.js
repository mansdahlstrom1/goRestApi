// Prepare form element and attach event Listeners
const form = document.getElementById("wifi-form");
if (form.attachEvent) {
  form.attachEvent("submit", processForm);
} else {
  form.addEventListener("submit", processForm);
}


// Prepare Http Request and response handler
var xhr = new XMLHttpRequest();
xhr.onload = function () {

	// Process our return data
	if (xhr.status >= 200 && xhr.status < 300) {
		// What do when the request is successful
    console.log('success!', xhr);
    const res = JSON.parse(xhr.response);
    console.log(res);
	} else {
		// What do when the request fails
		console.log('The request failed!');
	}

	// Code that should run regardless of the request status
	console.log('This always runs...');
};

// Submit event handler
function processForm(e) {
  if (e.preventDefault) {
    e.preventDefault();
  }

  const params = {
    ssid: document.querySelector('#ssid').value,
    passphrase: document.querySelector('#passphrase').value
  }

  // Validate params
  if (!params.ssid || params.ssid.length < 1) {
    console.error('invalid ssid ?')
  }

  xhr.open('POST', '/api', true);
  xhr.setRequestHeader('Content-type', 'application/json');
  xhr.send(JSON.stringify(params));

  return false;
}
