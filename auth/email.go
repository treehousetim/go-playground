package auth

import (
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "log"
)

func SendEmail(toEmail, subject, content string) error {
    from := mail.NewEmail("Your App", "no-reply@yourapp.com")
    to := mail.NewEmail("", toEmail)
    message := mail.NewSingleEmail(from, subject, to, content, content)
    client := sendgrid.NewSendClient("your_sendgrid_api_key")
    response, err := client.Send(message)
    if err != nil {
        log.Println(err)
        return err
    }
    log.Printf("Email sent: %v\n", response.StatusCode)
    return nil
}
