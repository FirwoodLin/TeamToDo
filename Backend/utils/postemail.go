package utils

import (
	"TeamToDo/global"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func PostEmail(email, uuid string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "。。。Team")
	m.SetHeader("To", email)
	m.SetHeader("验证你的TeamTodDo账号")

	link := "http://127.0.0.1:8080/users/verify?uuid=" + uuid

	//先直接发送链接
	m.SetBody("text/html", link)

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
