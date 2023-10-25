package dispatcher

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

// EmailDispatcher is an interface for sending emails.
type EmailDispatcher interface {
	Compose(to []string, subject, text, html string) *email.Email
	Send(e *email.Email) error
}

// ExchangeEmailDispatcher is an implementation of EmailDispatcher for sending emails through Exchange Server.
type ExchangeEmailDispatcher struct {
	SMTPServer   string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
}

func NewEmailDispatcher(smtpServer string, smtpPort int, smtpUsername string, smtpPassword string) *ExchangeEmailDispatcher {
	emailDispatcher := ExchangeEmailDispatcher{}
	emailDispatcher.SMTPServer = smtpServer
	emailDispatcher.SMTPPort = smtpPort
	emailDispatcher.SMTPUsername = smtpUsername
	emailDispatcher.SMTPPassword = smtpPassword
	return &emailDispatcher
}

// Compose creates a new email message.
func (d *ExchangeEmailDispatcher) Compose(to []string, subject, text, html string) *email.Email {
	e := email.NewEmail()
	e.From = d.SMTPUsername
	e.To = to
	e.Subject = subject
	e.Text = []byte(text)
	e.HTML = []byte(html)
	return e
}

// Send sends the email using the ExchangeEmailDispatcher configuration.
func (d *ExchangeEmailDispatcher) Send(e *email.Email) error {
	auth := smtp.PlainAuth("", d.SMTPUsername, d.SMTPPassword, d.SMTPServer)
	smtpAddr := d.SMTPServer + ":" + string(d.SMTPPort)
	return e.Send(smtpAddr, auth)
}
