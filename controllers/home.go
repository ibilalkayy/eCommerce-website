// Two functions are present in this page.

package controllers

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/products"
)

type ProductStruct struct {
	Names   []string
	Images1 []string
	Images2 []string
	Links   []string
	Prices  []string
}

// ProductsForHome() function access every products to be used in the home page.
// For accessing the products, they are appended into the pList to be used in the Home template.
func ProductsForHome() ProductStruct {
	var pList ProductStruct
	for i := 0; i < len(products.AccessColumns(0)); i++ {
		pList.Names = append(pList.Names, products.AccessColumns(0)[i])
		pList.Images1 = append(pList.Images1, products.AccessColumns(1)[i])
		pList.Images2 = append(pList.Images2, products.AccessColumns(2)[i])
		pList.Links = append(pList.Links, products.AccessColumns(3)[i])
		pList.Prices = append(pList.Prices, products.AccessColumns(4)[i])
	}
	return pList
}

// Home() function executes the home page with all the products in it.
func Home(w http.ResponseWriter, r *http.Request) error {
	// If a page that has no link is visited, then go to the page error.
	if r.URL.Path != "/" {
		PageError(w, r)
		return nil
	}
	pList := ProductsForHome()
	if r.Method == "GET" {
		return HomeTmpl.Execute(w, pList)
	} else if r.Method == "POST" {
		Newsletter(w, r)
		return HomeTmpl.Execute(w, pList)
	}
	return nil
}
