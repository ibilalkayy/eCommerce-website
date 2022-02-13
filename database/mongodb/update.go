// This page contains three functions that update the data into the database.

package mongodb

// Importing the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateRegisterPassword() will filter the data based on email and password and
// update a key with a value.
func UpdateRegisterPassword(cred [2]string, key, value string) {
	collection := Connect.Database("eCommerce").Collection("register")
	filter := bson.M{"email": cred[0], "password": cred[1]}
	update := bson.M{
		"$set": bson.M{
			key: value,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateProfile() will filter the data based on email and password and
// update the values based on the given keys.
func UpdateProfile(value [11]string, cred [2]string) {
	collection := Connect.Database("eCommerce").Collection("register")
	filter := bson.M{"email": cred[0], "password": cred[1]}
	update := bson.M{
		"$set": bson.M{
			"fname":    value[0],
			"lname":    value[1],
			"email":    value[2],
			"password": value[3],
			"phone":    value[4],
			"country":  value[5],
			"state":    value[6],
			"faddress": value[7],
			"laddress": value[8],
			"postal":   value[9],
			"company":  value[10],
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateCard() will filter the data based on email and password and
// update the card values based on the card keys.
func UpdateCard(value [4]string, cred [2]string) {
	collection := Connect.Database("eCommerce").Collection("register")
	filter := bson.M{"email": cred[0], "password": cred[1]}
	update := bson.M{
		"$set": bson.M{
			"card name":    value[0],
			"card number":  value[1],
			"expiry month": value[2],
			"expiry year":  value[3],
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
