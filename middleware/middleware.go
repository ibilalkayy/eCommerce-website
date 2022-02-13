package middleware

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func LoadEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal(err)
	}
	return value
}

// Type to be used as a parameter in a function
type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandling handles the error and returns http.Handler
func ErrorHandling(h MyHandlerFunc) http.HandlerFunc {
	// The reason for returning a HandlerFunc is to use it as a middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Fatal(err)
		}
	})
}
