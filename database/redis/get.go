package redis

// Importing the libraries
import (
	"encoding/json"
	"regexp"
)

// GetCredentials() will get the data based on an id that you give and
// it will convert that data from a json to a regular string to be returned.
func GetCredentials(id string) (bool, [2]string) {
	client := CommonClient()
	val, err := client.Get(id).Result()
	if err != nil {
		return false, [2]string{"", ""}
	}

	note := MyCredentials{}
	err = json.Unmarshal([]byte(val), &note)
	if err != nil {
		return false, [2]string{"", ""}
	}
	return true, [2]string{note.Email, note.Password}
}

// GetToken() will get the data based on an id that you give and
// it will apply the regular expression to give only the token and
// return true if it is found.
func GetToken(id string) bool {
	client := CommonClient()
	val, err := client.Get(id).Result()
	if err != nil {
		return false
	}

	re, err := regexp.Compile(`.*:"|".*`)
	if err != nil {
		return false
	}

	_ = re.ReplaceAllString(val, "")
	return true
}
