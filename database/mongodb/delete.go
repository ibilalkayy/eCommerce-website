// This page contains two functions. Both are for the purpose of deleting the data from the database.

package mongodb

// Importing the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteAccount() will take the email and password and by using DeleteOne(),
// it will delete that document which belongs to the given email and password.
func DeleteAccount(email, password string) {
	collection := Connect.Database("eCommerce").Collection("register")
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"email": email, "password": password}); err != nil {
		log.Fatal(err)
	}
}

// DeletePayment() will filter the document based on the email and password and by using UpdateOne(),
// it will set the document data equal to empty string.
func DeletePayment(cred [2]string) {
	collection := Connect.Database("eCommerce").Collection("register")
	filter := bson.M{"email": cred[0], "password": cred[1]}
	update := bson.M{
		"$set": bson.M{
			"card name":    "",
			"card number":  "",
			"expiry month": "",
			"expiry year":  "",
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
