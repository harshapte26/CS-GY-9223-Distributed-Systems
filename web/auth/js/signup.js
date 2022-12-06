function registerUser() {

  var firstName = document.getElementById("firstname").value;
  var lastName = document.getElementById("lastname").value;
  var email = document.getElementById("email").value;
  var password = document.getElementById("pwd").value;

  if (!firstName) {
      toastr.error("Please enter First Name!")
      return
  }
  if (!lastName) {
      toastr.error("Please enter Last Name!")
      return
  }
  if (!email) {
      toastr.error("Please enter Email!")
      return
  }
  if (!password) {
      toastr.error("Please enter Password!")
      return
  }

  var registerData = {
      "firstName": firstName,
      "lastName": lastName,
      "email": email,
      "password": password
  }
  $.ajax({
      url: "http://localhost:8000/register/",
      method: "POST",
      data: JSON.stringify(registerData),
      dataType: 'json',
      contentType: "application/json",
      success: function (data) {
          if (data.Status == 200) {
              toastr.success(data.Message);
          } else {
              toastr.error(data.Message)
          }

      },
  });
}