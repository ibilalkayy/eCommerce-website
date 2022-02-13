// This page contains two functions. Both are used for sending an email. It uses gomail package.

package controllers

// Importing the libraries
import (
	"log"

	"github.com/ibilalkayy/eCommerce/middleware"
	"gopkg.in/gomail.v2"
)

// SendContact() sends the contact information to an admin email address and
// Reply-To will contain an email address of a sender through which an admin can reply to.
func SendContact(value [5]string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", myEmail)
	mail.SetHeader("Reply-To", value[1])
	mail.SetHeader("Subject", value[2])
	mail.SetBody("text/plain", value[0]+",\n\n"+value[3]+"\n\nIt is my mobile/phone number: "+value[4])

	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}

// SendRegister() will send an email to the user for confirmation and it contains a link that will lead to the password entrance page.
func SendRegister(value string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", value)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", "Confirm the email address")
	mail.SetBody("text/html", "Click on this <a href=\"https://application-ecommerce.herokuapp.com/password\">link</a> to give the password")

	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}
