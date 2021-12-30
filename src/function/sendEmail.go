package function

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	gomail "gopkg.in/mail.v2"
)

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
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
func SendEmail2(email, electivo1, electivo3, electivo2 string, E1, E2, E3 bool, nombre string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "soporte.electivos.ucn@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	var Estado1 string
	var Estado2 string
	var Estado3 string
	if E1 {
		Estado1 = "ACEPTADA"
	} else {
		Estado1 = "RECHAZADA"
	}
	if E2 {
		Estado2 = "ACEPTADA"
	} else {
		Estado2 = "RECHAZADA"
	}
	if E3 {
		Estado3 = "ACEPTADA"
	} else {
		Estado3 = "RECHAZADA"
	}

	t, _ := TimeIn(time.Now(), "America/Santiago")
	fecha := t.Format("2006-01-02 15:04:05")
	m.SetHeader("Subject", "Postulacion Electivos")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Servicio de postulacion a electivos UCN \n\n"+
		"Estimad@ "+nombre+" le infomamos que su solicitud para la toma de electivos ha cambiado con fecha: "+fecha+"\n"+
		"Las siguientes postulaciones cambiaron su estado a :\n "+
		"\n"+
		"     Prioridad 1:\n"+
		"          Electivo: "+electivo1+"\n"+
		"          Estado  : "+Estado1+" \n "+
		"\n"+

		"     Prioridad 2:\n"+
		"          Electivo: "+electivo2+"\n"+
		"          Estado  : "+Estado2+" \n "+
		"\n"+

		"     Prioridad 3:\n"+
		"          Electivo: "+electivo3+"\n"+
		"          Estado  : "+Estado3+" \n "+
		"\n"+

		"No responder este correo")

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
