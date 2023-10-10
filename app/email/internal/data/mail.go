package data

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"kx-boutique/app/email/internal/biz"
	"kx-boutique/app/email/internal/conf"
	"net/smtp"

	"github.com/go-kratos/kratos/v2/log"
)

// Mailer.
type mailer struct {
	SmtpHost    string
	SmtpPort    string
	From        string
	Pwd         string
	TemplateDir string
}

// NewMailer .
func NewMailer(c *conf.Data, logger log.Logger) (biz.Mailer, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the mailer resources")
	}
	return &mailer{
		SmtpHost:    c.Smtp.Host,
		SmtpPort:    c.Smtp.Port,
		From:        c.Smtp.From,
		Pwd:         c.Smtp.Pwd,
		TemplateDir: c.Smtp.TemplatesDir,
	}, cleanup, nil
}

func (m *mailer) SendMail(ctx context.Context, req *biz.SendMailRequest) error {
	// Sender data.
	from := m.From
	password := m.Pwd

	// Receiver email address.
	to := []string{
		req.To,
	}

	// smtp server configuration.
	smtpHost := m.SmtpHost
	smtpPort := m.SmtpPort

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles(m.TemplateDir + "/" + req.Template)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf(req.Subject+" \n%s\n\n", mimeHeaders)))

	// Inject payload data into template
	t.Execute(&body, req.Payload)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// // Sender data.
// from := m.From
// password := m.Pwd

// // Receiver email address.
// to := []string{
// 	req.To,
// }

// // smtp server configuration.
// smtpHost := m.SmtpHost
// smtpPort := m.SmtpPort

// // Authentication.
// auth := smtp.PlainAuth("", from, password, smtpHost)

// t, _ := template.ParseFiles(m.TemplateDir + "/" + req.Template)

// var body bytes.Buffer

// mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

// t.Execute(&body, struct {
// 	Name    string
// 	Message string
// }{
// 	Name:    "Puneet Singh",
// 	Message: "This is a test message in a HTML template",
// })

// // Sending email.
// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// if err != nil {
// 	fmt.Println(err)
// 	return err
// }
// fmt.Println("Email Sent!")
// return nil
