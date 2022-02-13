// This page contains two functions that will delete the data from the database.

package mysql

import "log"

// It will take the email and password from the database and after finding that data,
// delete a specific a table from the database.
func DeleteAccount(email, password string) {
	db := Connection()
	q := "DELETE from Register WHERE emails=? and passwords=?"
	delete, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer delete.Close()

	_, err = delete.Exec(email, password)
	if err != nil {
		log.Fatal(err)
	}
}

// It will take the email and password from the database and after finding that data,
// update a specific data in a table in the database.
func DeletePayment(cred [2]string) {
	db := Connection()
	q := "UPDATE Register SET cardnames=?, cardnumbers=?, expmonths=?, expyears=? WHERE emails=? and passwords=?"
	delete, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer delete.Close()

	_, err = delete.Exec("", "", "", "", cred[0], cred[1])
	if err != nil {
		log.Fatal(err)
	}
}
