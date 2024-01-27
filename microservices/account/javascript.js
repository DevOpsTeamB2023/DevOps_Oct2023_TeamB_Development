function signup(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('signupForm');

    const curl = 'http://localhost:5001/api/v1/accounts';

    const username = form.elements['signup_username'].value;
    const password = form.elements['signup_password'].value;
    console.log(username);
    console.log(password);

    request.open("POST", curl);
    request.send(JSON.stringify({
        "username": username,
        "password": password, 
        "accType": "User", 
        "accStatus": "Pending"
        
    }));
    form.reset();
    alert("Account request sent. Please wait for admin approval.");
    return false //prevent default submission
}

function login(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('loginForm');
    const username = form.elements['login_username'].value;
    const password = form.elements['login_password'].value;
    console.log(username);
    console.log(password);

    const curl = 'http://localhost:5001/api/v1/accounts?username=' + encodeURIComponent(username) + '&password=' + encodeURIComponent(password);
    console.log(curl);

    request.open("GET", curl);
    request.onreadystatechange = function() {
      if (request.readyState === 4) {
        if (request.status === 200) {
          // Successful login, redirect to main page
          location.href = "index.html";
        } else if (request.status === 401) {
          // Login failed, handle error
          form.reset();
          document.getElementById('error-message').innerHTML = 'Incorrect Phone Number or Password.';
        } else {
          // Handle other status codes or network errors
          document.getElementById('error-message').innerHTML = 'An error occurred. Please try again later.';
        }
      }
    };
    request.send();
    return false
}