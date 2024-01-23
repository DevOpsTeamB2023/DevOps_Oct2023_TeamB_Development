function signup(){
    var request = new XMLHttpRequest();
    const form = document.getElementById('signup_form');

    const curl = 'http://localhost:5001/api/v1/accounts';

    const username = form.elements['username'].value;
    const password = form.elements['password'].value;
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

// function deleteAccount(email){
//     var request = new XMLHttpRequest();
//     console.log(screen)
//     request.open('DELETE', 'http://localhost:1765/api/v1/book/'+email);
//     request.send();
// }