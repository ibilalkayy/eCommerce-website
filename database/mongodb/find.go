// This page contains four functions. These functions will find the data in the database.

package mongodb

// Importing the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserCredentials struct {
	Fname      string `bson:"fname"`
	Lname      string `bson:"lname"`
	Email      string `bson:"email"`
	Password   string `bson:"password"`
	Phone      string `bson:"phone"`
	Country    string `bson:"country"`
	State      string `bson:"state"`
	Faddress   string `bson:"faddress"`
	Laddress   string `bson:"laddress"`
	Postal     string `bson:"postal"`
	Company    string `bson:"company"`
	Cardname   string `bson:"card name"`
	Cardnumber string `bson:"card number"`
	Expmonth   string `bson:"expiry month"`
	Expyear    string `bson:"expiry year"`
}

type CommentVariables struct {
	Name    []byte `bson:"name"`
	Email   []byte `bson:"email"`
	Message []byte `bson:"message"`
	Date    []byte `bson:"date"`
	Time    []byte `bson:"time"`
}

// FindAccount() will find the data in the register collection based on the email address
// and password and return eleven string values and boolean also.
func FindAccount(email, password string) (bool, [11]string) {
	var uc UserCredentials
	collection := Connect.Database("eCommerce").Collection("register")
	if err := collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&uc); err != nil {
		return false, [11]string{}
	}
	values := [11]string{uc.Fname, uc.Lname, uc.Email, uc.Password, uc.Phone, uc.Country, uc.State, uc.Faddress, uc.Laddress, uc.Postal, uc.Company}
	return true, values
}

// CardDetails() will find the card details in the register collection based on the email address
// and password and return the card details.
func CardDetails(email, password string) [4]string {
	var uc UserCredentials
	collection := Connect.Database("eCommerce").Collection("register")
	if err := collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&uc); err != nil {
		return [4]string{}
	}
	values := [4]string{uc.Cardname, uc.Cardnumber, uc.Expmonth, uc.Expyear}
	return values
}

// FindAccountForComment() will find the card details in the comment collection based
// on the email address and return the bool.
func FindAccountForComment(email string) bool {
	var uc UserCredentials
	collection := Connect.Database("eCommerce").Collection("register")
	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&uc); err != nil {
		return false
	}
	return true
}

// FindComment() will find the card details in the comment collection based on the email address
// and return all the comments as a CommentVariable with slice.
// It will run the for loop with rows.Next() to print all comments and append that data into res variable.
func FindComment(email string) (res []CommentVariables) {
	var cv CommentVariables
	collection := Connect.Database("eCommerce").Collection("comment")
	rows, err := collection.Find(context.TODO(), bson.M{"email": email})
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close(context.TODO())

	for rows.Next(context.TODO()) {
		if err := rows.Decode(&cv); err != nil {
			log.Fatal(err)
		}
		res = append(res, cv)
	}
	return res
}
