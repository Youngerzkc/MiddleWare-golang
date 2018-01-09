package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	fmt.Println("Hello World")
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)

	send_to := strings.Split(to, ";")
	
	err := smtp.SendMail(host, auth, user, send_to, msg)
	fmt.Println("err",err)
	return err
}

func main() {
	// user := "yang**@yun*.com"
	// password := "***"
	user:="young_zkc@163.com"
	password:="13462012314ZKC"
	host := "smtp.163.com:25"
	to := "984888712@qq.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		" 你好，这是我的golang邮件  "
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}