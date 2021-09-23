document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.querySelector("#login-form");
  const signupForm = document.querySelector("#signup-form");

  document.querySelector("#linkLogin").addEventListener("click", (e) => {
    e.preventDefault();
    signupForm.classList.add("form-hidden");
    loginForm.classList.remove("form-hidden");
  });

  document.querySelector("#linkSignUp").addEventListener("click", (e) => {
    e.preventDefault();
    loginForm.classList.add("form-hidden");
    signupForm.classList.remove("form-hidden");
  });

  document.querySelectorAll(".input-field").forEach((inputElement) => {
    inputElement.addEventListener("blur", (e) => {
      if (
        e.target.id === "signUpPassword" &&
        e.target.value.length > 0 &&
        e.target.value.length < 10
      ) {
        setInputError(
          inputElement,
          "Password must be at least 10 characters long"
        );
      }
    });

    inputElement.addEventListener("input", (e) => {
      clearInputError(inputElement);
    });
  });

  loginForm.addEventListener("submit", (e) => {
    //e.preventDefault();

    // Perform AJAX/Fetch login

    //setFormMessage(loginForm, "error", "Incorrect username or password!");
  });

});

// how to use:
// setFormMessage(loginForm, "success", "Successfuly logged in!")
// setFormMessage(loginForm, "error", "Incorrect username or password")
function setFormMessage(formElement, type, message) {
  const messageElement = formElement.querySelector(".form-message");

  messageElement.textContent = message;
  messageElement.classList.remove("form-message-error", "form-message-success");
  messageElement.classList.add(`form-message-${type}`);
}

function setInputError(inputElement, message) {
  inputElement.classList.add("input-field-error");
  inputElement.parentElement.querySelector(
    ".input-field-error-message"
  ).textContent = message;
}

function clearInputError(inputElement) {
  inputElement.classList.remove("input-field-error");
  inputElement.parentElement.querySelector(
    ".input-field-error-message"
  ).textContent = "";
}
