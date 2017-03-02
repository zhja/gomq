package gonoc

import (
	"net/smtp"
    "strings"
    //"fmt"
)

const (
    HOST        = "smtp.126.com"
    SERVER_ADDR = "smtp.126.com:25"
    USER        = "zhja60@126.com" //发送邮件的邮箱
    PASSWORD    = "zhaojian201030" //发送邮件邮箱的密码
)

type Email struct {
    to      string "to"
    subject string "subject"
    mailtype string "mailtype"
    msg     string "msg"
}

func NewEmail(to, subject, mailtype, msg string) *Email {
    return &Email{to: to, subject: subject, mailtype: mailtype, msg: msg}
}

func SendEmail(email *Email) error {
    auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
    sendTo := strings.Split(email.to, ";")
    done := make(chan error, 1024)

    go func() {
        defer close(done)
        for _, v := range sendTo {

        	var content_type string
        	if email.mailtype == "html" {
        		content_type = "Content-Type: text/"+ email.mailtype + "; charset=UTF-8"
	    	}else{
	        	content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	    	}
	 
	    	str := []byte("To: " + v + "\r\nFrom: " + USER + "<"+ USER +">\r\nSubject: " + email.subject + "\r\n" + content_type + "\r\n\r\n" + email.msg)
	    	//send_to := strings.Split(to, ";")
	    	//err := smtp.SendMail(host, auth, user, send_to, msg)
	    	//return err

            //str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

            err := smtp.SendMail(
                SERVER_ADDR,
                auth,
                USER,
                []string{v},
                []byte(str),
            )
            done <- err
        }
    }()

    for i := 0; i < len(sendTo); i++ {
        <-done
    }

    return nil
}

func EmailTemplate(text string) (content string) {
	content = "<div style='width:1024px;margin : auto;font-size: 14px;'><p>您好：<br></p><p>[PANIC] MQ 数据异常,详情为:" + text + "</p><p><strong>如有疑问请联系金山云：</strong><a href='mailto:ksc_idc@kingsoft.com'><strong>ksc_idc@kingsoft.com</strong></a>！</p><p style='float: right;''><em>谢谢!</em><br/><em>金山云NOC平台开发组</em></p></div>"
	return
}

func NocEmail(title string, content string) error {
	// msg := EmailTemplate(content)
	// email := NewEmail("zhaojian3@kingsoft.com", title, "html", msg)
	// _ = SendEmail(email)
	return nil
} 