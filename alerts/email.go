package alerts

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/wecredit/prometheus-sdk/config"
)

func SendEmailAlert(cfg config.EmailAlert, subject, body string) error {
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTPHost)
	msg := "From: " + cfg.From + "\n" +
		"To: " + cfg.To[0] + "\n" +
		"Subject: " + subject + "\n\n" + body

	return smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.From, cfg.To, []byte(msg))
}

func SendEmailAlertWithRetry(cfg config.EmailAlert, subject, body string, retries int) {
	for i := 0; i < retries; i++ {
		err := SendEmailAlert(cfg, subject, body)
		if err == nil {
			fmt.Println("Email alert sent.")
			return
		}
		fmt.Printf("Failed to send email (try %d/%d): %v\n", i+1, retries, err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	fmt.Println("Alert email failed after retries.")
}
