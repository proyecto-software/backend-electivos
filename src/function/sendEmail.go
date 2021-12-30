package function

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	gomail "gopkg.in/mail.v2"
)

func SendEmail() {
	from := "soporte.inscripcion.alumnos@gmail.com"
	password := "proyecto2021"

	// Receiver email address.
	to := []string{
		"dionisio.olivares@alumnos.ucn.cl",
		"ssp013@alumnos.ucn.cl",
		"jose.flores02@alumnos.ucn.cl",
		"nicolas.garcia@alumnos.ucn.cl",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("probando probando..... dionisio gei .")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
func SendEmail2(email, electivo string, E bool) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "soporte.electivos.ucn@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	var Estado string
	if E {
		Estado = "ACEPTADA"
	} else {
		Estado = "RECHAZADA"
	}

	m.SetHeader("Subject", "SOLICITUD: "+Estado)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "El estado de su solicitud pertenecitente al electivo: "+electivo+" cambio a : "+Estado)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "soporte.electivos.ucn@gmail.com", "proyecto2021")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
