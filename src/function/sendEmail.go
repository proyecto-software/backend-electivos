package function

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

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
func SendEmail2(email string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("EMAIL"))

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	Estado := "ACEPTADO"
	Electivo := "REDES INALAMBRICAS"
	m.SetHeader("Subject", "El estado de su solicitud de electivo cambio a : "+Estado)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Has quedado seleccionado en el electivo: "+Electivo)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}