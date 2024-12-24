// notification/notification.go

package notification

import (
    "fmt"
    "net/smtp"
    "github.com/pk1151222/bug-scanners/internal/models"
)

func SendEmailAlert(result *models.ScanResult) {
    if len(result.HTTPVulnerabilities) > 0 || len(result.SSLVulnerabilities) > 0 {
        msg := fmt.Sprintf("Subject: Critical Vulnerabilities Detected on %s\n\n", result.URL)
        msg += fmt.Sprintf("HTTP Vulnerabilities: %v\n", result.HTTPVulnerabilities)
        msg += fmt.Sprintf("SSL Vulnerabilities: %v\n", result.SSLVulnerabilities)

        auth := smtp.PlainAuth("", "your_email@example.com", "your_password", "smtp.example.com")
        err := smtp.SendMail("smtp.example.com:587", auth, "your_email@example.com", []string{"recipient@example.com"}, []byte(msg))
        if err != nil {
            fmt.Println("Error sending email:", err)
        } else {
            fmt.Println("Alert email sent successfully.")
        }
    }
}
