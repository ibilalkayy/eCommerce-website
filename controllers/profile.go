// The page contains three functions. First will show the profile details,
// second will update and third will delete the profile.

package controllers

// Importing the libraries
import (
	"net/http"
	"time"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
	"github.com/ibilalkayy/eCommerce/models"
)

// Profile() will get login token and login credentials to check if it is logged in and
// then take the signup data from the database and print them in the profile template.
func Profile(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, mongodbData := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			note := models.SignupData{
				Fname:    mongodbData[0],
				Lname:    mongodbData[1],
				Email:    mongodbData[2],
				Password: mongodbData[3],
				Phone:    mongodbData[4],
				Country:  mongodbData[5],
				State:    mongodbData[6],
				Faddress: mongodbData[7],
				Laddress: mongodbData[8],
				Postal:   mongodbData[9],
				Company:  mongodbData[10],
			}
			return ProfileTmpl.Execute(w, note)
		} else if r.Method == "POST" {
			Newsletter(w, r)
			http.Redirect(w, r, "/your-profile", http.StatusFound)
			return nil
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// UpdateProfile() will check the login credentials, take the user information and
// replace the old info with the newly entered info in the database. After that, it will
// delete the old credentials and set the new credentials and new token to automatically login so that a user don't have to login again.
func UpdateProfile(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return UpdateProfileTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			if r.FormValue("update-profile") == "UpdateProfile" {
				pData := models.SignupData{
					Fname:           r.FormValue("fname"),
					Lname:           r.FormValue("lname"),
					Email:           r.FormValue("email"),
					Password:        r.FormValue("password"),
					ConfirmPassword: r.FormValue("confirm-password"),
					Phone:           r.FormValue("phone"),
					Country:         r.FormValue("country"),
					State:           r.FormValue("state"),
					Faddress:        r.FormValue("faddress"),
					Laddress:        r.FormValue("laddress"),
					Postal:          r.FormValue("postal"),
					Company:         r.FormValue("company"),
				}
				// Convert the password into hash values.
				hashPassword := HashPassword([]byte(pData.Password))
				if ComparePasswords(hashPassword, []byte(pData.ConfirmPassword)) {
					newCred := [11]string{pData.Fname, pData.Lname, pData.Email, hashPassword, pData.Phone, pData.Country, pData.State, pData.Faddress, pData.Laddress, pData.Postal, pData.Company}
					// mysql.UpdateProfile(newCred, credData)
					mongodb.UpdateProfile(newCred, credData)
					redis.Del("LoginPass")
					redis.SetCredentials("LoginPass", newCred[2], newCred[3])
					// mysqlDataFound, _ := mysql.FindAccount(newCred[2], newCred[3])
					mongodbDataFound, _ := mongodb.FindAccount(newCred[2], newCred[3])
					if mongodbDataFound && ComparePasswords(hashPassword, []byte(pData.Password)) {
						tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Second))
						if ok {
							if redis.Del("LoginToken") && redis.Del("LoginCred") {
								redis.SetToken("LoginToken", tokenString)
								redis.SetCredentials("LoginCred", newCred[2], newCred[3])
								http.Redirect(w, r, "/your-profile", http.StatusFound)
								return nil
							} else {
								fm := Messages{LoginFailure: "Please login first to view the page"}
								return LoginTmpl.Execute(w, fm)
							}
						} else {
							fm := Messages{LoginFailure: "Please login first to view the page"}
							return LoginTmpl.Execute(w, fm)
						}
					} else {
						fm := Messages{LoginFailure: "Please login first to view the page"}
						return LoginTmpl.Execute(w, fm)
					}
				} else {
					fm := Messages{LoginFailure: "Both passwords are not matched"}
					return LoginTmpl.Execute(w, fm)
				}
			} else if r.FormValue("update-profile") == "Newsletter" {
				Newsletter(w, r)
				return UpdateProfileTmpl.Execute(w, nil)
			}
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}

// DeleteProfile() will check the credentials and then ask for a password to delete a profile.
// If the password is matched with the password present in the database, it will be deleted otherwise it
// will give the failure message.
func DeleteProfile(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return DeleteProfileTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			findPass := models.SignupData{
				Password: r.FormValue("password"),
			}
			if ComparePasswords(credData[1], []byte(findPass.Password)) {
				// mysql.DeleteAccount(credData[0], credData[1])
				mongodb.DeleteAccount(credData[0], credData[1])
				redis.Del("LoginPass")
				redis.Del("LoginToken")
				redis.Del("LoginCred")
				http.Redirect(w, r, "/", http.StatusFound)
				return nil
			} else {
				fm := Messages{PasswordFailure: "Please enter the correct password"}
				return DeleteProfileTmpl.Execute(w, fm)
			}
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}
