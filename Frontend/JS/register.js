var button = document.querySelector('button');
button.addEventListener('click',register);

function register () {
    var email = document.getElementById('email').value;
    var pwd1 = document.getElementById('pwd1').value;
    var pwd2 = document.getElementById('pwd2').value;

    if (pwd1 !== pwd2) {
        alert('请重新输入密码');
        document.getElementById("pwd1").value = "";
        document.getElementById("pwd2").value = "";
    }
    else {
        password = pwd1;

        fetch('http://localhost:8080/users/registration', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json'},
            body: JSON.stringify({
                email: email,
                password: password
            })

        })
        .then((response) => {
            if (!response.ok){
                return response.json();
            }else {
                throw new Error('请求失败');
            }
        })
        .then(message => {
            if(message.success) {
                localStorage.setItem('userID', message.data.userID);
                localStorage.setItem('userName', message.data.userName);
                localStorage.setItem('userAvatar', message.data.userAvatar);
                alert('激活链接已发送~ 请注意查看邮箱！');
                window.location.href = "temp_activate.html";
            }
            else {
                alert('message.hint');
                document.getElementById('email').value = '';
                document.getElementById('pwd1').value = '';
                document.getElementById('pwd2').value = '';
            }
        })
        .catch(error => {
            console.error(error);
        });
    }

}