// This page contains six functions.

package controllers

import "net/http"

type Footer struct {
	Name       string
	Link       string
	Paragraph1 string
	Paragraph2 string
	Question1  string
	Question2  string
	Question3  string
	Question4  string
	Question5  string
	Answer     string
}

// CommonTemplate()contains different values to be used in all the five footer functions.
// For each value there is a variable that are used in the footer functions
func CommonTemplate(name, link string) Footer {
	note := Footer{
		Name:       name,
		Link:       link,
		Paragraph1: ProductDesc[0],
		Paragraph2: ProductDesc[1],
		Question1:  "- This is the first question",
		Question2:  "- This is the second question",
		Question3:  "- This is the third question",
		Question4:  "- This is the forth question",
		Question5:  "- This is the fifth question",
		Answer:     "He standard Lorem Ipsum passage, used since the 1500s Excepteur sint occaecat",
	}
	return note
}

func AboutUs(w http.ResponseWriter, r *http.Request) error {
	note := CommonTemplate("About Us", "/about-us")
	return FooterDetailsTmpl.Execute(w, note)
}

func Faq(w http.ResponseWriter, r *http.Request) error {
	note := CommonTemplate("Frequently Asked Questions", "/faq")
	return FaqDetailsTmpl.Execute(w, note)
}

func TermsConditions(w http.ResponseWriter, r *http.Request) error {
	note := CommonTemplate("Terms & Conditions", "/terms-conditions")
	return FooterDetailsTmpl.Execute(w, note)
}

func Shipping(w http.ResponseWriter, r *http.Request) error {
	note := CommonTemplate("Shipping", "/shipping")
	return FooterDetailsTmpl.Execute(w, note)
}

func PrivacyPolicy(w http.ResponseWriter, r *http.Request) error {
	note := CommonTemplate("Privacy Policy", "/privacy-policy")
	return FooterDetailsTmpl.Execute(w, note)
}
