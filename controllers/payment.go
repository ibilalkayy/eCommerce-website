// This page contains six function that are related to payment.

package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
	"github.com/ibilalkayy/eCommerce/models"
)

// Payment() will find the token and credentials to login, refresh the token also and
// using the conditional statement, it will execute a template.
func Payment(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return PaymentTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			Newsletter(w, r)
			return PaymentTmpl.Execute(w, nil)
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// PaymentMethod() will also check the login credentials and then take the card information from a user and insert that info into the database.
func PaymentMethod(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return PaymentMethodTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			if r.FormValue("payment-method") == "AddCard" {
				addCard := models.AddCard{
					Name:   r.FormValue("card-name"),
					Number: r.FormValue("card-number"),
					Month:  r.FormValue("month"),
					Year:   r.FormValue("year"),
				}
				values := [4]string{addCard.Name, addCard.Number, addCard.Month, addCard.Year}
				// mysql.UpdateCard(values, credData)
				mongodb.UpdateCard(values, credData)
				http.Redirect(w, r, "/your-payment-details", http.StatusFound)
				return nil
			} else if r.FormValue("payment-method") == "Newsletter" {
				Newsletter(w, r)
				return PaymentMethodTmpl.Execute(w, nil)
			}
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// CardDetails() will get the card details from the database and return that info.
func CardDetails(w http.ResponseWriter, r *http.Request, cred [2]string) models.AddCard {
	// mysqlDetails := mysql.CardDetails(cred[0], cred[1])
	mongodbDetails := mongodb.CardDetails(cred[0], cred[1])
	note := models.AddCard{
		Name:   mongodbDetails[0],
		Number: mongodbDetails[1],
		Month:  mongodbDetails[2],
		Year:   mongodbDetails[3],
	}
	if (len(note.Name) == 0 || note.Name == "card name") && (len(note.Number) == 0 || note.Number == "card number") && (len(note.Month) == 0 || note.Month == "expiry month") && (len(note.Year) == 0 || note.Year == "expiry year") {
		http.Redirect(w, r, "/your-payment-method", http.StatusFound)
	}
	return note
}

// PaymentDetails() will check the login credentials and get the card details and use them in a template.
func PaymentDetails(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			note := CardDetails(w, r, credData)
			return PaymentDetailsTmpl.Execute(w, note)
		} else if r.Method == "POST" {
			Newsletter(w, r)
			note := CardDetails(w, r, credData)
			return PaymentDetailsTmpl.Execute(w, note)
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// UpdatePayment() is same like the PaymentMethod() but it will take the card info to update.
func UpdatePayment(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return UpdatePaymentTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			if r.FormValue("update-payment") == "UpdatePayment" {
				updateCard := models.AddCard{
					Name:   r.FormValue("card-name"),
					Number: r.FormValue("card-number"),
					Month:  r.FormValue("month"),
					Year:   r.FormValue("year"),
				}
				values := [4]string{updateCard.Name, updateCard.Number, updateCard.Month, updateCard.Year}
				// mysql.UpdateCard(values, credData)
				mongodb.UpdateCard(values, credData)
				http.Redirect(w, r, "/your-payment-details", http.StatusFound)
				return nil
			} else if r.FormValue("update-payment") == "Newsletter" {
				Newsletter(w, r)
				return UpdatePaymentTmpl.Execute(w, nil)
			}
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// DeletePayment(). If a clicks on the delete payment button, it will ask for a password.
// If the password matched from the user account, it will delete that payment method.
func DeletePayment(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return DeletePaymentTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			findPass := models.SignupData{
				Password: r.FormValue("password"),
			}
			if ComparePasswords(credData[1], []byte(findPass.Password)) {
				// mysql.DeletePayment(credData)
				mongodb.DeletePayment(credData)
				http.Redirect(w, r, "/", http.StatusFound)
				return nil
			} else {
				fm := Messages{PasswordFailure: "Please enter the correct password"}
				return DeletePaymentTmpl.Execute(w, fm)
			}
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}
