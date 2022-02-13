// This page contains four functions. These functions will find the data in the database.

package mysql

import "log"

type UserCredentials struct {
	Email    string
	Password string
	Cred     [11]string
}

// FindAccount() will find the data in the register table based on the email address
// and password and return eleven string values and boolean also.
func FindAccount(email, password string) (bool, [11]string) {
	db := Connection()
	var uc UserCredentials
	q := "select fname, lname, emails, passwords, phones, country, states, faddress, laddress, postal, company from Register where emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Cred[0], &uc.Cred[1], &uc.Cred[2], &uc.Cred[3], &uc.Cred[4], &uc.Cred[5], &uc.Cred[6], &uc.Cred[7], &uc.Cred[8], &uc.Cred[9], &uc.Cred[10]); err != nil {
		return false, [11]string{}
	}
	return true, uc.Cred
}

// CardDetails() will find the card details in the register table based on the email address
// and password and return the card details.
func CardDetails(email, password string) [4]string {
	db := Connection()
	var uc UserCredentials
	q := "select cardnames, cardnumbers, expmonths, expyears from Register where emails=? and passwords=?"
	if err := db.QueryRow(q, email, password).Scan(&uc.Cred[0], &uc.Cred[1], &uc.Cred[2], &uc.Cred[3]); err != nil {
		return [4]string{}
	}
	values := [4]string{uc.Cred[0], uc.Cred[1], uc.Cred[2], uc.Cred[3]}
	return values
}

// FindAccountForComment() will find the card details in the comment table based
// on the email address and return the bool.
func FindAccountForComment(email string) bool {
	db := Connection()
	var uc UserCredentials
	q := "select emails from Register where emails=?"
	if err := db.QueryRow(q, email).Scan(&uc.Cred[0]); err != nil {
		return false
	}
	return true
}

// FindComment() will find the card details in the comment table based on the email address
// and return all the comments as a UserCredentials with slice.
func FindComment(email string) (res []UserCredentials) {
	db := Connection()
	var uc UserCredentials
	q := "select names, emails, messages, dates, timez from Comment where emails=?"
	rows, err := db.Query(q, email)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&uc.Cred[0], &uc.Cred[1], &uc.Cred[2], &uc.Cred[3], &uc.Cred[4]); err != nil {
			log.Fatal(err)
		}
		res = append(res, uc)
	}
	return res
}
