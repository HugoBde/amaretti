const passwordField = document.getElementsByName("password")[0]

function togglePasswordVisibility(event) {
  if (passwordField.type == "password") {
    revealPassword(event.target, passwordField);
  } else {
    hidePassword(event.target, passwordField);
  }
}

function hidePassword(toggleBtn, passwordField) {
  toggleBtn.innerHTML = "visibility";
  passwordField.type = "password";
}

function revealPassword(toggleBtn, passwordField) {
  toggleBtn.innerHTML = "visibility_off";
  passwordField.type = "text";
}

