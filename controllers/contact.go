package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/models"
)

// Contact() function takes the name, subject, email, phone and message from a person and
// send that data to my email using the SendContact() function. The data will also be stored in the database for record.
func Contact(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return ContactTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		send := models.ContactData{
			Name:    r.FormValue("name"),
			Subject: r.FormValue("subject"),
			Email:   r.FormValue("email"),
			Phone:   r.FormValue("phone"),
			Message: r.FormValue("message"),
		}
		values := [5]string{send.Name, send.Email, send.Subject, send.Message, send.Phone}
		// mysql.InsertContact(values)
		mongodb.InsertContact(values)
		SendContact(values)
		Newsletter(w, r)
		return ContactTmpl.Execute(w, nil)
	}
	return nil
}
