package utils

import (
	"TeamToDo/global"
	"bytes"
	"crypto/tls"
	"html/template"

	"gopkg.in/gomail.v2"
)

var items = map[string]interface{}{
	"Logo": "",
	"Link": "",
}

// 读取渲染激活邮件
func GenerateActivateMail(uuid string) string {
	link := "http://127.0.0.1:8080/api/users/verify?uuid=" + uuid

	items["Link"] = link
	items["Logo"] = global.Server.Mail.Logo

	return generateMail("./template/activate.tmpl", items, link)
}

// 读取渲染定时题型邮件
func GenerateRemindMail(taskName, taskDesc, startTime, DeadLine string) string {
	link := "http://127.0.0.1:8080"

	items["Link"] = link
	items["Logo"] = global.Server.Mail.Logo
	items["TaskName"] = taskName
	items["TaskDesc"] = taskDesc
	items["StartTime"] = startTime
	items["DeadLine"] = DeadLine

	return generateMail("./template/remind.tmpl", items, link)
}

// 生成一般邮件
func generateMail(filename string, data map[string]interface{}, defaultRet string) string {
	tmpl, err := template.New("Mail").ParseFiles(filename)

	if err != nil {
		global.Logger.Infof("读取模板失败，err: %s", err)
		return defaultRet
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		global.Logger.Infof("渲染模板失败，err: %s", err)
		return defaultRet
	}
	return buf.String()
}

// 发送邮件
func PostEmail(email, text string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", global.Server.Mail.Username)
	m.SetHeader("To", email)
	m.SetHeader("VerifyTeamTodDo")

	// 发送html
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
