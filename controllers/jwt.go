// This page contains three functions. First two function generate and authorize a token.
// The third function refreshes a token after every five minutes.

package controllers

// Importing the libraries
import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ibilalkayy/eCommerce/database/redis"
	"github.com/ibilalkayy/eCommerce/middleware"
)

// GenerateJWT() sets the expriration time and access the success key present in the Environment Variable
// With the help of that key, a token generates with the expiration time that is returned.
func GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	byteEnv := []byte(middleware.LoadEnvVariable("ACCESS_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(byteEnv)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// IsAuthorized() function stores the token in a cookie and check its validity.
func IsAuthorized(w http.ResponseWriter, r *http.Request, expirationTime time.Time) (string, bool) {
	tokenString, err := GenerateJWT()
	if err != nil {
		return "", false
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})

	byteEnv := []byte(middleware.LoadEnvVariable("ACCESS_SECRET"))
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return byteEnv, err
	})

	if token.Valid {
		return tokenString, true
	} else {
		return "", false
	}
}

// RefreshToken() function calls the IsAuthorized() function and refreshes the token after every
// five minutes and sets the LoginToken in redis if a user refreshes the page.
func RefreshToken(w http.ResponseWriter, r *http.Request) bool {
	if time.Until(time.Now().Add(time.Duration(jwt.StandardClaims{}.ExpiresAt))) > 30*time.Second {
		return false
	} else {
		tokenString, ok := IsAuthorized(w, r, time.Now().Add(5*time.Minute))
		if ok {
			redis.SetToken("LoginToken", tokenString)
			return true
		} else {
			return false
		}
	}
}
