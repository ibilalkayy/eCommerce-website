package models

type SendNewsletter struct {
	Email  string
	Record string
}

type ContactData struct {
	Name    string
	Subject string
	Email   string
	Phone   string
	Message string
}

type SignupData struct {
	Fname           string
	Lname           string
	Email           string
	Phone           string
	Country         string
	State           string
	Faddress        string
	Laddress        string
	Postal          string
	Company         string
	Cardname        string
	Cardnumber      string
	Expmonth        string
	Expyear         string
	Password        string
	ConfirmPassword string
}

type AddCard struct {
	Name   string
	Number string
	Month  string
	Year   string
}

type AddComment struct {
	Name    string
	Email   string
	Message string
}
