// This page contains 18 functions. 17 of them just execute the template but

package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/products"
)

type ProductDetails struct {
	Name   string
	Image1 string
	Image2 string
	Link   string
	Price  string
	Desc1  string
	Desc2  string
}

var ProductDesc = []string{
	"He standard Lorem Ipsum passage, used since the 1500s Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Section 1.10.32 of de Finibus Bonorum et Malorum, written by Cicero in 45 BC Sed ut perspiciatis unde omnis Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
	"Iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
}

// Common() function is used for fetching the product details from the data.csv file and
// returning them so that it will be used in all the 17 functions.
func Common(productName string) ProductDetails {
	note := ProductDetails{
		Name:   products.AccessData(productName, 0),
		Image1: products.AccessData(productName, 1),
		Image2: products.AccessData(productName, 2),
		Link:   products.AccessData(productName, 3),
		Price:  products.AccessData(productName, 4),
		Desc1:  ProductDesc[0],
		Desc2:  ProductDesc[1],
	}
	return note
}

// For showing the Women Hot Collection page.
func WomenHotCollection(w http.ResponseWriter, r *http.Request) error {
	note := Common("Women Hot Collection")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Awesome Pink Show page.
func AwesomePinkShow(w http.ResponseWriter, r *http.Request) error {
	note := Common("Awesome Pink Show")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the page.
func AwesomeBagsCollection1(w http.ResponseWriter, r *http.Request) error {
	note := Common("Awesome Bag Collection")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Women Pant Collections page.
func WomenPantCollections(w http.ResponseWriter, r *http.Request) error {
	note := Common("Women Pant Collections")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Awesome Bags Collection page.
func AwesomeBagsCollection2(w http.ResponseWriter, r *http.Request) error {
	note := Common("Awesome Bags Collection")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Awesome Cap For Women page.
func AwesomeCapForWomen(w http.ResponseWriter, r *http.Request) error {
	note := Common("Awesome Cap For Women")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Polo Dress For Women page.
func PoloDressForWomen(w http.ResponseWriter, r *http.Request) error {
	note := Common("Polo Dress For Women")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Black Sunglass For Women page.
func BlackSunglassForWomen(w http.ResponseWriter, r *http.Request) error {
	note := Common("Black Sunglass For Women")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the First Cloth page.
func FirstCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("First Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Second Cloth page.
func SecondCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Second Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Third Cloth page.
func ThirdCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Third Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Forth Cloth page.
func ForthCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Forth Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Fifth Cloth page.
func FifthCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Fifth Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Sixth Cloth page.
func SixthCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Sixth Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Seventh Cloth page.
func SeventhCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Seventh Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Eighth Cloth page.
func EighthCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Eighth Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}

// For showing the Ninth Cloth page.
func NinthCloth(w http.ResponseWriter, r *http.Request) error {
	note := Common("Ninth Cloth")
	return ProductDetailsTmpl.Execute(w, note)
}
