package controllers

// Importing the libraries
import (
	"net/http"
	"strings"

	"github.com/ibilalkayy/eCommerce/products"
)

// SearchProducts() will access the product names and links from the products file and
// and run a loop to access all of them. With the help of conditional statement, it will check
// whether the search products are equal to the products in the file.
func SearchProducts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		if r.FormValue("operation") == "search" {
			note := struct{ Search string }{Search: r.FormValue("searchProducts")}
			productNames := products.AccessColumns(0)
			productLinks := products.AccessColumns(3)
			for i := 0; i < 17; i++ {
				if productNames[i] == note.Search || strings.ToLower(productNames[i]) == strings.ToLower(note.Search) {
					http.Redirect(w, r, productLinks[i], http.StatusFound)
					break
				}
			}
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	return nil
}
