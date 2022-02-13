package redis

// Importing the libraries
import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

type MyCredentials struct {
	Email    string
	Password string
	Token    string
}

// CommonClient() is used for setting the connetion.
// You can either set localhost or cloud address in the Addr
func CommonClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "<cloud address> OR localhost:6379",
		Password: "<password>",
		DB:       0,
	})
	return client
}

// SetCredentials() will convert the string into json format and
// give/set the id to that json email and password.
func SetCredentials(id, email, password string) {
	client := CommonClient()
	json, err := json.Marshal(MyCredentials{Email: email, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Set(id, json, 0).Err(); err != nil {
		log.Fatal(err)
	}
}

// SetToken() will convert the string into json format and
// give/set the id to that token string.
func SetToken(id, tokenString string) {
	client := CommonClient()
	json, err := json.Marshal(MyCredentials{Token: tokenString})
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Set(id, json, 0).Err(); err != nil {
		log.Fatal(err)
	}
}
