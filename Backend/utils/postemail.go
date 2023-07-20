package utils

import (
	"TeamToDo/global"
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func PostEmail(email, uuid string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "TeamTodDo Team <"+global.Server.Mail.Username+">")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "TeamToDo 激活邮件")

	//m.SetHeader("VerifyTeamTodDo")
	emailTemplate :=
		`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>TeamToDo 激活邮件</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      margin: 0;
      padding: 0;
      background-image: url("path/to/your/background-image.jpg");
      background-repeat: no-repeat;
      background-size: cover;
    }
    .container {
      width: 100%%;
      max-width: 600px;
      margin: 0 auto;
      background-color: #fff;
      padding: 20px;
      box-sizing: border-box;
    }
    .logo {
      display: block;
      width: 150px;
      margin: 20px auto;
    }
    .content {
      text-align: center;
    }
    .cta {
      display: inline-block;
      padding: 10px 20px;
      background-color: #1a73e8;
      color: #fff;
      text-decoration: none;
    }
    .copyright {
      font-size: 12px;
      text-align: center;
      color: #333;
      margin-top: 20px;
    }
  </style>
</head>
<body>
  <div class="container">
    <img src="%s" alt="Team To Do Logo" class="logo">
    <div class="content">
      <h1>感谢您的注册！</h1>
      <p>请点击下方的激活链接以完成您的注册过程：</p>
      <a href="%s" class="cta">激活我的帐户</a>
    </div>
    <p class="copyright">© 2023 TeamToDo 保留所有权利。</p>
  </div>
</body>
</html>`
	link := "http://127.0.0.1:8080/api/users/verify?uuid=" + uuid
	text := fmt.Sprintf(emailTemplate, global.Server.Mail.Logo, link)
	//先直接发送链接
	m.SetBody("text/html", text)

	d := gomail.NewDialer(
		global.Server.Mail.Host,
		global.Server.Mail.Post,
		global.Server.Mail.Username,
		global.Server.Mail.Secret,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
