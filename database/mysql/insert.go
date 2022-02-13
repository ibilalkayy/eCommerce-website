package mysql

import "log"

// InsertNewsletter() will insert the email address and time record in the newsletter table of database
func InsertNewsletter(email, record string) {
	db := CreateTable("db.SQL", 0)
	q := "INSERT INTO Newsletter(emails, records) VALUES(?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	if len(email) != 0 {
		_, err = insert.Exec(email, record)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// InsertContact() will insert the name, subject, email, phone and message in the contact table of database
func InsertContact(value [5]string) {
	db := CreateTable("db.SQL", 1)
	q := "INSERT INTO Contact(names, subjects, emails, phones, messages) VALUES(?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	if len(value[2]) != 0 {
		_, err = insert.Exec(value[0], value[1], value[2], value[3], value[4])
		if err != nil {
			log.Fatal(err)
		}
	}
}

// InsertSignup() will insert the registration data in the register table of database
func InsertSignup(value [15]string) {
	db := CreateTable("db.SQL", 2)
	q := "INSERT INTO Register(fname, lname, emails, passwords, phones, country, states, faddress, laddress, postal, company, cardnames, cardnumbers, expmonths, expyears) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	if len(value[2]) != 0 && len(value[3]) != 0 {
		_, err = insert.Exec(value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7], value[8], value[9], value[10], value[11], value[12], value[13], value[14])
		if err != nil {
			log.Fatal(err)
		}
	}
}

// InsertComment() will insert the name, email, message, date and time in the comment table of database
func InsertComment(value [5]string) {
	db := CreateTable("db.SQL", 3)
	q := "INSERT INTO Comment(names, emails, messages, dates, timez) VALUES(?, ?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	if len(value[0]) != 0 && len(value[1]) != 0 && len(value[2]) != 0 {
		_, err = insert.Exec(value[0], value[1], value[2], value[3], value[4])
		if err != nil {
			log.Fatal(err)
		}
	}
}
