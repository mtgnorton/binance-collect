package shared

import (
	"net/smtp"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

var Email = email{}

type email struct {
}

// 异步通过qq邮箱服务发送邮件

func (e *email) SendByQQAsync(subject, body, to string, isHtml bool) {
	go e.SendByQQSync(subject, body, to, isHtml)
}

// 通过qq邮箱服务发送邮件

func (e *email) SendByQQSync(subject, body, to string, isHtml bool) error {
	from := "851426308@qq.com"
	password := "mrnjlfbtvhobbdcd"
	host := "smtp.qq.com:587"
	sendUserName := "gf-admin" //发送邮件的人名称

	return e.sendOrigin(from, sendUserName, password, host, to, subject, body, isHtml)

}

/**
 * 发送邮件
 * @param from 发送人
 * @param sendUserName 发送人名称
 * @param password 发送人密码 不是邮箱密码,需要登陆你的邮箱，在设置，账号，启用IMAP/SMTP服务，会发送一段身份验证符号给你，用这个登陆
 * @param host 发送人邮箱服务器地址
 * @param to 接收人,多个以;号隔开
 * @param subject 主题
 * @param body 内容
 * @param isHtml 是否是html格式
 */
func (e *email) sendOrigin(from, sendUserName, password, host, to, subject, body string, isHtml bool) error {

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", from, password, hp[0])
	var content_type string

	if isHtml {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + from + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

	targets := strings.Split(to, ";")
	g.Dump(host, auth, from, targets, msg)
	err := smtp.SendMail(host, auth, from, targets, msg)
	return err
}
