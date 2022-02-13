// This page uses four functions that are related to user information and
// performing login logout functionality.

package controllers

// Importing the libraries
import (
	"net/http"
	"time"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
	"github.com/ibilalkayy/eCommerce/models"
)

type Messages struct {
	PasswordFailure string
	LoginFailure    string
}

// ConfirmEmail() will execute the confirm email template. It will be used when a user registers,
// then this page will be shown to visit the email and confirm the email.
func ConfirmEmail(w http.ResponseWriter, r *http.Request) error {
	return ConfirmEmailTmpl.Execute(w, nil)
}

// Password() will take the password and confirm from a user, make a the signup credential,
// convert that into hash password and then insert that hash password in the database. After that,
// it will set the login credentials and token and generate the JWT token by using the IsAuthorized() function and
// making it able to automatically login without logging in again.
func Password(w http.ResponseWriter, r *http.Request) error {
	getPass := models.SignupData{
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirm-password"),
	}
	signupCredFound, signupCred := redis.GetCredentials("SignupCred")
	hashPassword := HashPassword([]byte(getPass.Password))
	if len(getPass.Password) != 0 {
		if signupCredFound && ComparePasswords(hashPassword, []byte(getPass.ConfirmPassword)) {
			// mysql.UpdateRegisterPassword(hashPassword, signupCred[0], signupCred[1])
			mongodb.UpdateRegisterPassword(signupCred, "password", hashPassword)
			// mysqlMatched, _ := mysql.FindAccount(signupCred[0], hashPassword)
			mongodbMatched, _ := mongodb.FindAccount(signupCred[0], hashPassword)
			redis.SetCredentials("LoginCred", signupCred[0], hashPassword)
			if mongodbMatched {
				tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Minute))
				if ok {
					redis.SetCredentials("LoginPass", signupCred[0], hashPassword)
					redis.SetToken("LoginToken", tokenString)
					redis.GetCredentials("LoginCred")
					http.Redirect(w, r, "/", http.StatusFound)
					return nil
				} else {
					note := Messages{PasswordFailure: "Both passwords are not matched."}
					return PasswordTmpl.Execute(w, note)
				}
			} else {
				note := Messages{PasswordFailure: "Both passwords are not matched."}
				return PasswordTmpl.Execute(w, note)
			}
		} else {
			note := Messages{PasswordFailure: "Both passwords are not matched."}
			return PasswordTmpl.Execute(w, note)
		}
	}
	return PasswordTmpl.Execute(w, nil)
}

// Login() will take the email and password. It will get the login credentials and token and
// find the account information in the database based on the login credentials. After comparing the passwords,
// it will delete the previous credentials and generate new ones to login again.
func Login(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return LoginTmpl.Execute(w, nil)
	}
	userLogin := models.SignupData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	passFound, loginPass := redis.GetCredentials("LoginPass")
	// mysqlMatched, _ := mysql.FindAccount(userLogin.Email, loginPass[1])
	mongodbMatched, _ := mongodb.FindAccount(userLogin.Email, loginPass[1])
	if passFound && mongodbMatched && ComparePasswords(loginPass[1], []byte(userLogin.Password)) {
		tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Minute))
		if ok {
			if redis.Del("LoginToken") && redis.Del("LoginCred") {
				redis.SetCredentials("LoginCred", userLogin.Email, loginPass[1])
				redis.SetToken("LoginToken", tokenString)
				http.Redirect(w, r, "/", http.StatusFound)
				return nil
			} else {
				http.Redirect(w, r, "/", http.StatusFound)
				return nil
			}
		} else {
			fm := Messages{LoginFailure: "Enter the correct email or password"}
			return LoginTmpl.Execute(w, fm)
		}
	} else {
		fm := Messages{LoginFailure: "Enter the correct email or password"}
		return LoginTmpl.Execute(w, fm)
	}
}

// Logout() will delete the login credentials and
// set IsAuthorized() time limit to zero to delete token from the cookie.
func Logout(w http.ResponseWriter, r *http.Request) {
	if redis.Del("LoginToken") && redis.Del("LoginCred") {
		_, _ = IsAuthorized(w, r, time.Now().Add(0*time.Minute))
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
