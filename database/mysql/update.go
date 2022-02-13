// This page contains three functions that update the data into the database.

package mysql

import "log"

// UpdateRegisterPassword() will filter the data based on email and password and
// update a password with a new one.
func UpdateRegisterPassword(newPass, email, oldPass string) {
	db := Connection()
	q := "UPDATE Register SET passwords=? WHERE emails=? and passwords=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()

	_, err = update.Exec(newPass, email, oldPass)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateProfile() will filter the data based on email and password and
// update multiple values in the database.
func UpdateProfile(value [11]string, cred [2]string) {
	db := Connection()
	q := "UPDATE Register SET fname=?, lname=?, emails=?, passwords=?, phones=?, country=?, states=?, faddress=?, laddress=?, postal=?, company=? WHERE emails=? and passwords=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()

	_, err = update.Exec(value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7], value[8], value[9], value[10], cred[0], cred[1])
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateCard() will filter the data based on email and password and
// update the card values in the database.
func UpdateCard(value [4]string, cred [2]string) {
	db := Connection()
	q := "UPDATE Register SET cardnames=?, cardnumbers=?, expmonths=?, expyears=? WHERE emails=? and passwords=?"
	update, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()

	_, err = update.Exec(value[0], value[1], value[2], value[3], cred[0], cred[1])
	if err != nil {
		log.Fatal(err)
	}
}
