package mailutil

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

var defaultMail *Mail

func Init(conf Config) {
	mail, err := NewMail(conf)
	if err != nil {
		log.Fatalln(err)
	}

	defaultMail = mail
}

func Send(subject, to, body string) error {
	return defaultMail.Send(subject, to, body)
}

type Config struct {
	Host     string `yaml:"host"`
	Sender   string `yaml:"sender"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Mail struct {
	conf Config

	auth smtp.Auth
}

func NewMail(conf Config) (*Mail, error) {
	host, _, err := net.SplitHostPort(conf.Host)
	if err != nil {
		return nil, err
	}

	return &Mail{
		conf: conf,
		auth: smtp.PlainAuth("", conf.Sender, conf.Password, host),
	}, nil
}

func (m *Mail) Send(subject, to, body string) error {
	headers := []string{
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("From: %s <%s>", m.conf.Username, m.conf.Sender),
		fmt.Sprintf("Subject: %s", subject),
	}
	headerStr := strings.Join(headers, "\r\n")
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := fmt.Sprintf("%s\r\n%s\r\n\r\n%s", headerStr, contentType, body)
	return smtp.SendMail(m.conf.Host, m.auth, m.conf.Sender, []string{to}, []byte(msg))
}
