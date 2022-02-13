// This file contains six functions. These are the single functions that are used in many other
// functions.

package controllers

// Importing the libraries
import (
	"log"
	"net/http"
	"time"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/models"
	"github.com/ibilalkayy/eCommerce/products"
	"golang.org/x/crypto/bcrypt"
)

type DataList struct {
	Names  string
	Images string
	Links  string
	Prices string
}

// Cart() is used for showing the cart items. It uses GET and POST method.
// GET is used when you visit the page and POST is used when you fill a form.
func Cart(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return CartTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		Newsletter(w, r)
		return CartTmpl.Execute(w, nil)
	}
	return nil
}

// ProductsList() is used for accessing the products information and printing them in the template.
func ProductsList(w http.ResponseWriter, r *http.Request) error {
	var pList []DataList
	for i := 0; i < 17; i++ {
		pList = append(pList, DataList{
			products.AccessColumns(0)[i],
			products.AccessColumns(1)[i],
			products.AccessColumns(3)[i],
			products.AccessColumns(4)[i],
		})
	}
	if r.Method == "GET" {
		return ProductsListTmpl.Execute(w, pList)
	} else if r.Method == "POST" {
		Newsletter(w, r)
		return ProductsListTmpl.Execute(w, pList)
	}
	return nil
}

// HashPassword() will generate the hash values from a byte and convert that into string.
func HashPassword(value []byte) string {
	hash, err := bcrypt.GenerateFromPassword(value, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// Newsletter() will take an email, attach the current date time to it and then save it in the database.
func Newsletter(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	send := models.SendNewsletter{
		Email:  r.FormValue("EMAIL"),
		Record: currentTime.Format("2006-01-02 15:04:05"),
	}
	// mysql.InsertNewsletter(send.Email, send.Record)
	note := [2]string{send.Email, send.Record}
	mongodb.InsertNewsletter(note)
}

// ComparePasswords() will compare two byte values and then return true or false.
func ComparePasswords(hashPass string, plainPass []byte) bool {
	hashByte := []byte(hashPass)

	if err := bcrypt.CompareHashAndPassword(hashByte, plainPass); err != nil {
		return false
	}
	return true
}

// PageError() will execute the page error template. It will be executed if some body visits unknown page link.
func PageError(w http.ResponseWriter, r *http.Request) {
	PageErrorTmpl.Execute(w, nil)
}
