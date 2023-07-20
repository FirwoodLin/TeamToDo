package utils

import (
	"TeamToDo/global"
	"bytes"
	"crypto/tls"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

var items = map[string]interface{}{
	"Logo": "",
	"Link": "",
}

// GenerateActivateMail 读取渲染 激活邮件
func GenerateActivateMail(uuid string) string {
	link := global.Server.Server.Host + ":" + global.Server.Server.Port + "/api/users/verify?uuid=" + uuid

	items["Link"] = link
	items["Logo"] = global.Server.Mail.Logo

	return generateMail("./template/activate.tmpl", items, link)
}

// GenerateRemindMail 读取渲染 定时型邮件
func GenerateRemindMail(taskName, taskDesc, startTime, DeadLine string) string {
	link := global.Server.Server.Host + ":" + global.Server.Server.Port

	items["Link"] = link
	items["Logo"] = global.Server.Mail.Logo
	items["TaskName"] = taskName
	items["TaskDesc"] = taskDesc
	items["StartTime"] = startTime
	items["DeadLine"] = DeadLine
	//global.Logger.Debugf("GenerateRemindMail, items: %v", items)
	return generateMail("./template/remind.tmpl", items, link)
}

// 生成一般邮件
func generateMail(filename string, data map[string]interface{}, defaultRet string) string {
	file, err := os.ReadFile("./template/remind.tmpl")
	//fmt.Println(file)
	// 将file 转换为string
	str := string(file)
	tmpl, err := template.New("Mail").Parse(str)

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

// PostEmail 发送邮件
func PostEmail(email, text string) error {
	global.Logger.Debugf("PostEmail ready to 发送邮件给：%s", email)
	m := gomail.NewMessage()

	m.SetHeader("From", "TeamToDo 团队 <"+global.Server.Mail.Username+">")
	m.SetHeader("To", email)
	m.SetBody("text/html", text) // 发送html

	d := gomail.NewDialer(
		global.Server.Mail.Host,
		global.Server.Mail.Post,
		global.Server.Mail.Username,
		global.Server.Mail.Secret,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		global.Logger.Errorf("发送邮件to %s失败，err: %s", email, err.Error())
		return err
	}
	return nil
}
