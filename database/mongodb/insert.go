// This page contains four functions that will insert the data in the database.

package mongodb

// Importing the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// InsertNewsletter() will insert the email address and time record in the newsletter collection of database
func InsertNewsletter(value [2]string) {
	collection := Connect.Database("eCommerce").Collection("newsletter")
	if len(value[0]) != 0 {
		if _, err := collection.InsertOne(context.TODO(), bson.M{"email": value[0], "record": value[1]}); err != nil {
			log.Fatal(err)
		}
	}
}

// InsertContact() will insert the name, subject, email, phone and message in the contact collection of database
func InsertContact(value [5]string) {
	collection := Connect.Database("eCommerce").Collection("contact")
	if len(value[2]) != 0 {
		if _, err := collection.InsertOne(context.TODO(), bson.M{"name": value[0], "subject": value[1], "email": value[2], "phone": value[3], "message": value[4]}); err != nil {
			log.Fatal(err)
		}
	}
}

// InsertSignup() will insert the registration data in the register collection of database
func InsertSignup(value [15]string) {
	collection := Connect.Database("eCommerce").Collection("register")
	if len(value[2]) != 0 && len(value[3]) != 0 {
		if _, err := collection.InsertOne(context.TODO(), bson.M{"fname": value[0], "lname": value[1], "email": value[2], "password": value[3], "phone": value[4], "country": value[5], "state": value[6], "faddress": value[7], "laddress": value[8], "postal": value[9], "company": value[10], "card name": value[11], "card number": value[12], "expiry month": value[13], "expiry year": value[14]}); err != nil {
			log.Fatal(err)
		}
	}
}

// InsertComment() will insert the name, email, message, date and time in the comment collection of database
func InsertComment(value [5]string) {
	collection := Connect.Database("eCommerce").Collection("comment")
	if len(value[0]) != 0 && len(value[1]) != 0 && len(value[2]) != 0 {
		if _, err := collection.InsertOne(context.TODO(), bson.M{"name": value[0], "email": value[1], "message": value[2], "date": value[3], "time": value[4]}); err != nil {
			log.Fatal(err)
		}
	}
}
