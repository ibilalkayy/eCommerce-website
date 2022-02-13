package redis

// Del() will search for the token based on the id and delete it.
// It will return the bool value.
func Del(id string) bool {
	client := CommonClient()
	_, err := client.Del(id).Result()
	if err != nil {
		return false
	} else {
		return true
	}
}
