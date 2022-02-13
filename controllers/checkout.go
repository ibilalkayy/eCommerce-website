package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
	"github.com/ibilalkayy/eCommerce/middleware"
	"github.com/ibilalkayy/eCommerce/models"
)

// Checkout() takes the information from a user, generate a default password and store all the data
// in the database. Redis will also store the credentials with a default password.
// After clicking on the Register button, SendRegister() function will send an
// email to the user email address to verify and click on the link to open the password for entrance of a new password.
func Checkout(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return CheckoutTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		if r.FormValue("operation") == "Register" {
			signup := models.SignupData{
				Fname:      r.FormValue("fname"),
				Lname:      r.FormValue("lname"),
				Email:      r.FormValue("email"),
				Phone:      r.FormValue("phone"),
				Country:    r.FormValue("country"),
				State:      r.FormValue("state"),
				Faddress:   r.FormValue("faddress"),
				Laddress:   r.FormValue("laddress"),
				Postal:     r.FormValue("postal"),
				Company:    r.FormValue("company"),
				Cardname:   "card name",
				Cardnumber: "card number",
				Expmonth:   "expiry month",
				Expyear:    "expiry year",
			}
			defaultPass := HashPassword([]byte(middleware.LoadEnvVariable("DEFAULT_PASSWORD")))
			note := [15]string{signup.Fname, signup.Lname, signup.Email, defaultPass, signup.Phone, signup.Country, signup.State, signup.Faddress, signup.Laddress, signup.Postal, signup.Company, signup.Cardname, signup.Cardnumber, signup.Expmonth, signup.Expyear}
			// mysql.InsertSignup(note)
			mongodb.InsertSignup(note)
			redis.SetCredentials("SignupCred", signup.Email, defaultPass)
			SendRegister(signup.Email)
			http.Redirect(w, r, "/confirm-email", http.StatusFound)
			return nil
		} else if r.FormValue("operation") == "Newsletter" {
			Newsletter(w, r)
			return CheckoutTmpl.Execute(w, nil)
		}
	}
	return nil
}
