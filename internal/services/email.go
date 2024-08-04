package services

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
)

func sendConfirmationEmail(email string, terms []string, domains []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "BrandMonitor Resultados")

	termsStr := strings.Join(terms, ", ")
	domainsStr := "<ul>"
	for _, domain := range domains {
		domainsStr += fmt.Sprintf("<li>%s</li>", domain)
	}
	domainsStr += "</ul>"

	body := fmt.Sprintf(`
        <html>
        <body>
            <p>Olá,</p>
            <p>Seu diagnóstico para os termos <strong>%s</strong> foi processado.</p>
            <p>Aqui estão os domínios dos concorrentes que estão utilizando seus termos de marca:</p>
            %s
            <p>Obrigado por usar nosso serviço!</p>
        </body>
        </html>
    `, termsStr, domainsStr)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
