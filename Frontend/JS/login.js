var button = document.querySelector("button");
button.addEventListener('click', login);

import { token, currentUserID, currentUserName, currentUserEmail, currentUserAvatar, UserID } from "./global.js";

function login() {
    var email = document.getElementById("email").value;
    var password = document.getElementById("pwd").value;

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/api/users/login", true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                var response = JSON.parse(xhr.responseText);
                if (response.success) {
                    token = response.data.token;
                    localStorage.setItem('token', token);

                    var user = response.data.user;
                    UserID = user.userID.toString();
                    currentUserID = user.userID.toString();
                    currentUserName = user.userName;
                    currentUserAvatar = user.userAvatar;
                    currentUserEmail = email;

                    localStorage.setItem('userID', user.userID.toString());
                    localStorage.setItem('userName', user.userName);
                    localStorage.setItem('userEmail', user.userEmail);
                    localStorage.setItem('userAvatar', user.userAvatar);

                    window.location.href = "team_01.html";
                }
                else {
                    alert(response.hint);
                }
            }
            else {
                alert("请求失败 : 用户不存在！请注册账号！");
                window.location.href = "register.html";
            }
        }
    };
    var requestBody = JSON.stringify({
        email: email,
        password: password
    });
    xhr.send(requestBody);

}

