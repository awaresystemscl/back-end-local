package main

import (
    "fmt"
    "log"
    "net/smtp"
    "net/mail"
    "strings"
    "encoding/base64"
)

func encodeRFC2047(String string) string{
    // use mail's rfc2047 to encode any string
    addr := mail.Address{String, ""}
    return strings.Trim(addr.String(), " <>")
}

func send(body string, nombre string, email string) {
    smtpServer := "smtp.gmail.com"
    auth := smtp.PlainAuth(
        "",
        "awaresystemscl@gmail.com",
        "awarelala123",
        smtpServer,
    )

    from := mail.Address{"AwareSystems", "notificaciones@awaresystems.cl"}
    to := mail.Address{nombre, email}
    title := "Alerta de violacion de requisitos "

    header := make(map[string]string)
    header["From"] = from.String()
    header["To"] = to.String()
    header["Subject"] = title
    header["MIME-Version"] = "1.0"
    header["Content-Type"] = "text/plain; charset=\"utf-8\""
    header["Content-Transfer-Encoding"] = "base64"

    message := ""
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    err := smtp.SendMail(
        smtpServer + ":587",
        auth,
        from.Address,
        []string{to.Address},
        []byte(message),
    )
    if err != nil {
        log.Fatal(err)
    }
    
}