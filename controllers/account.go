package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/database/redis"
)

// Account() function will get the token, and credentials from redis, find the details in the database and
// after applying the conditional statment, it will execute the template.
// It will also check the GET and POST method to see whether they user submitted a form or just visited the page.
func Account(w http.ResponseWriter, r *http.Request) error {
	tokenFound := redis.GetToken("LoginToken")
	credFound, credData := redis.GetCredentials("LoginCred")
	// mysqlDataFound, _ := mysql.FindAccount(credData[0], credData[1])
	mongodbDataFound, _ := mongodb.FindAccount(credData[0], credData[1])
	if tokenFound && credFound && mongodbDataFound && RefreshToken(w, r) {
		if r.Method == "GET" {
			return AccountTmpl.Execute(w, nil)
		} else if r.Method == "POST" {
			Newsletter(w, r)
			return AccountTmpl.Execute(w, nil)
		}
	} else {
		return LoginTmpl.Execute(w, nil)
	}
	return nil
}
