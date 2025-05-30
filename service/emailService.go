package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/smtp"
	"os"
)

type EmailService struct {
	smtpHost   string
	smtpPort   string
	from       string
	password   string
	logoBase64 string
}

type StatusOrderData struct {
	Name    string
	Status  string
	OrderId int
	Logo    string
}

type StatusReviewData struct {
	Answer string
	Name   string
}

func NewEmailService(smtpHost, smtpPort, from, password string) (*EmailService, error) {
	logoBytes, err := os.ReadFile("static/logo.png")
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл logo.png: %w", err)
	}

	logoBase64 := "data:image/png;base64," + base64.StdEncoding.EncodeToString(logoBytes)

	return &EmailService{
		smtpHost:   smtpHost,
		smtpPort:   smtpPort,
		from:       from,
		password:   password,
		logoBase64: logoBase64,
	}, nil
}

func (es *EmailService) SendOrderStatusEmail(to, name, status string, orderId int) error {
	tmpl, err := template.ParseFiles("templates/status_order.html")
	if err != nil {
		return fmt.Errorf("ошибка парсинга шаблона: %w", err)
	}

	data := StatusOrderData{
		Status:  status,
		OrderId: orderId,
		Name:    name,
	}

	var htmlBody bytes.Buffer
	if err := tmpl.Execute(&htmlBody, data); err != nil {
		return fmt.Errorf("ошибка выполнения шаблона: %w", err)
	}

	logoBytes, err := os.ReadFile("static/logo.png")
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл logo.png: %w", err)
	}

	boundary := "mixed-" + uuid.New().String()

	var msg bytes.Buffer
	msg.WriteString("From: Next Device <" + es.from + ">\r\n")
	msg.WriteString("To: " + to + "\r\n")
	msg.WriteString("Subject: Обновление статуса заказа\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/related; boundary=" + boundary + "\r\n")
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	msg.WriteString("Content-Transfer-Encoding: 7bit\r\n")
	msg.WriteString("\r\n")
	msg.Write(htmlBody.Bytes())
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: image/png\r\n")
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString("Content-ID: <logo.png>\r\n")
	msg.WriteString("Content-Disposition: inline; filename=\"logo.png\"\r\n")
	msg.WriteString("\r\n")

	b64 := base64.StdEncoding.EncodeToString(logoBytes)
	for i := 0; i < len(b64); i += 76 {
		end := i + 76
		if end > len(b64) {
			end = len(b64)
		}
		msg.WriteString(b64[i:end] + "\r\n")
	}
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "--\r\n")

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err = smtp.SendMail(
		es.smtpHost+":"+es.smtpPort,
		auth,
		es.from,
		[]string{to},
		msg.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("ошибка отправки email: %w", err)
	}

	return nil
}

func (es *EmailService) SendReviewStatusEmail(to, name, answer string) error {
	tmpl, err := template.ParseFiles("templates/status_review.html")
	if err != nil {
		return fmt.Errorf("ошибка парсинга шаблона: %w", err)
	}

	data := StatusReviewData{
		Answer: answer,
		Name:   name,
	}

	var htmlBody bytes.Buffer
	if err := tmpl.Execute(&htmlBody, data); err != nil {
		return fmt.Errorf("ошибка выполнения шаблона: %w", err)
	}

	logoBytes, err := os.ReadFile("static/logo.png")
	if err != nil {
		return fmt.Errorf("не удалось прочитать файл logo.png: %w", err)
	}

	boundary := "mixed-" + uuid.New().String()

	var msg bytes.Buffer
	msg.WriteString("From: Next Device <" + es.from + ">\r\n")
	msg.WriteString("To: " + to + "\r\n")
	msg.WriteString("Subject: Обновление статуса отзыва\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/related; boundary=" + boundary + "\r\n")
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	msg.WriteString("Content-Transfer-Encoding: 7bit\r\n")
	msg.WriteString("\r\n")
	msg.Write(htmlBody.Bytes())
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: image/png\r\n")
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString("Content-ID: <logo.png>\r\n")
	msg.WriteString("Content-Disposition: inline; filename=\"logo.png\"\r\n")
	msg.WriteString("\r\n")

	b64 := base64.StdEncoding.EncodeToString(logoBytes)
	for i := 0; i < len(b64); i += 76 {
		end := i + 76
		if end > len(b64) {
			end = len(b64)
		}
		msg.WriteString(b64[i:end] + "\r\n")
	}
	msg.WriteString("\r\n")

	msg.WriteString("--" + boundary + "--\r\n")

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)

	err = smtp.SendMail(
		es.smtpHost+":"+es.smtpPort,
		auth,
		es.from,
		[]string{to},
		msg.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("ошибка отправки email: %w", err)
	}

	return nil
}
