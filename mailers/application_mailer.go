package mailers

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type BaseMail struct {
	To      []string
	From    string
	Subject string
	Body    string
}

func AppDirectory() (dir string) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return

}

func BuildAuth() (auth smtp.Auth) {

	//Genera la configuracion de authenticacion PLAIN TEXT
	if os.Getenv("SMTP_AUTH") == "PLAIN" {
		user, pass := os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD")
		auth = smtp.PlainAuth("", user, pass, os.Getenv("SMTP_HOST"))
	}

	return

}

func ParseTemplate(filename string, parseData interface{}) (string, error) {

	buf := &bytes.Buffer{}

	t, err := template.ParseFiles(AppDirectory() + "/views/mailers/" + filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = t.Execute(buf, parseData)

	return buf.String(), err
}

func SendMail(content BaseMail) error {

	// Construye la authenticacion

	auth := BuildAuth()

	// Validar destinatarios

	if len(content.To) == 0 {
		log.Fatal("Debe indicar al menos un destinatario")
		return errors.New("el destinatario no puede estar vacio")
	}

	// Build mail's body

	msg := []byte("Content-type: text/html; charset=iso-8859-1;" + "\r\n" + "Subject: " + content.Subject + "\r\n" +
		"\r\n" + content.Body + "\r\n")

	// Join addr

	addr := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")

	// SET FROM (sender)

	var sender = ""

	if content.From == "" {
		sender = "no-contestar@taxky.win"
	} else {
		sender = content.From
	}

	// Send mail

	err := smtp.SendMail(addr, auth, sender, content.To, msg)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Print("Correo enviado.")
		return nil
	}

}
