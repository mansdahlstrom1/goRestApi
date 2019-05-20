// Prepare DOM element and attach event Listeners

// Form
const form = document.getElementById("wifi-form");
if (form.attachEvent) {
  form.attachEvent("submit", processForm);
} else {
  form.addEventListener("submit", processForm);
}

// Error Box
const errorBox = document.getElementById("error-box");
const errorTitle = document.getElementById("error-title");
const errorMessage = document.getElementById("error-message");
errorBox.style.display = "none";

function updateErrorMessage(success = false, title = null, message = null, ) {
  if (!success) {
    errorBox.style.display = "block";
    errorTitle.innerHTML = title || "Error...";
    errorMessage.innerHTML = message || "Something went wrong...";
  } else {
    errorBox.style.display = "none";
    errorTitle.innerHTML = "";
    errorMessage.innerHTML = "";
  }
}

// Prepare Http Request and response handler
var xhr = new XMLHttpRequest();
xhr.onload = function () {
	if (xhr.status >= 200 && xhr.status < 300) {
    const res = JSON.parse(xhr.response);
    updateErrorMessage(res.success, "", res.message)
    console.log(res);
	} else {
		console.log('The request failed!');
	}
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
