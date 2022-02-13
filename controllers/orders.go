// This page has only one function.

package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
)

// It is used for finding the data and getting the products present in the orders list.
// It will also get the LoginCred and LoginToken to check whether those tokens are found or not and get the credentials also.
// It also uses RefreshToken() function to refresh a token in a cookie and use POST and GET method.
func Orders(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return OrdersTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			Newsletter(w, r)
			return OrdersTmpl.Execute(w, nil)
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}
