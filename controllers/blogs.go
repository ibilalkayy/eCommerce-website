// This page is for blogs. It contains six functions.
// Four functions execute the four different blog templates by using GET and POST method.

package controllers

// Importing the libraries
import (
	"fmt"
	"net/http"
	"time"

	"github.com/ibilalkayy/eCommerce/database/mongodb"
	"github.com/ibilalkayy/eCommerce/models"
	"github.com/ibilalkayy/eCommerce/products"
)

type CommentData struct {
	Name     string
	Email    string
	Message  string
	Date     string
	Time     string
	Comments string
}

var cData []CommentData

// The ShowComments() function contains the code to get, find the comments and print them on the page.
// giveComment variable get a comment and find the email in the database. If it is matched then
// insert and find that comment in the database to show it on the blog template.
func ShowComments(w http.ResponseWriter, r *http.Request) {
	giveComment := models.AddComment{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}
	dt := time.Now()
	date := dt.Format("02-Jan-2006")
	time := dt.Format("15:04:05")
	// mysqlAccountFound := mysql.FindAccountForComment(giveComment.Email)
	mongodbAccountFound := mongodb.FindAccountForComment(giveComment.Email)
	if mongodbAccountFound {
		values := [5]string{giveComment.Name, giveComment.Email, giveComment.Message, date, time}
		// mysql.InsertComment(values)
		mongodb.InsertComment(values)

		// Finding and getting MySQL data
		// mysqlValues := mysql.FindComment(giveComment.Email)
		// x := fmt.Sprint(len(mysqlValues))
		// for i := 0; i < len(mysqlValues); i++ {
		// 	getMysqlComment := CommentData{
		// 		Name:     mysqlValues[i].Cred[0],
		// 		Email:    mysqlValues[i].Cred[1],
		// 		Message:  mysqlValues[i].Cred[2],
		// 		Date:     "On " + mysqlValues[i].Cred[3],
		// 		Time:     "At " + mysqlValues[i].Cred[4],
		// 		Comments: "Comments (" + x + ")",
		// 	}
		// 	fmt.Println(getMysqlComment)
		// }

		// Finding and getting Mongodb data
		mongodbValues := mongodb.FindComment(giveComment.Email)
		y := fmt.Sprint(len(mongodbValues))
		for i := 0; i < len(mongodbValues); i++ {
			getMongodbComment := CommentData{
				Name:     string(mongodbValues[i].Name),
				Email:    string(mongodbValues[i].Email),
				Message:  string(mongodbValues[i].Message),
				Date:     "On " + string(mongodbValues[i].Date),
				Time:     "At " + string(mongodbValues[i].Time),
				Comments: "Comments (" + y + ")",
			}
			cData = append(cData, getMongodbComment)
			fmt.Println(getMongodbComment)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// The Blogs() function contains the main page that has all the list of the blogs.
// Names, Images and Links is appended to the bList that uses the products to be accessed and used in the blog page.
// The blog page prints all the products details.
func Blogs(w http.ResponseWriter, r *http.Request) error {
	var bList []DataList
	for i := 17; i < 21; i++ {
		bList = append(bList, DataList{
			products.AccessColumns(0)[i],
			products.AccessColumns(1)[i],
			products.AccessColumns(3)[i],
			"",
		})
	}
	if r.Method == "GET" {
		return BlogsTmpl.Execute(w, bList)
	} else if r.Method == "POST" {
		Newsletter(w, r)
		return BlogsTmpl.Execute(w, bList)
	}
	return nil
}

func FirstBlog(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return FirstBlogTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		ShowComments(w, r)
		Newsletter(w, r)
		return FirstBlogTmpl.Execute(w, cData)
	}
	return nil
}

func SecondBlog(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return SecondBlogTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		ShowComments(w, r)
		Newsletter(w, r)
		return SecondBlogTmpl.Execute(w, cData)
	}
	return nil
}

func ThirdBlog(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return ThirdBlogTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		ShowComments(w, r)
		Newsletter(w, r)
		return ThirdBlogTmpl.Execute(w, cData)
	}
	return nil
}

func ForthBlog(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return ForthBlogTmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		ShowComments(w, r)
		Newsletter(w, r)
		return ForthBlogTmpl.Execute(w, cData)
	}
	return nil
}
