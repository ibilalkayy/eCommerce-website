package routes

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/eCommerce/controllers"
	"github.com/ibilalkayy/eCommerce/middleware"
)

func Routes() {
	// Header
	http.HandleFunc("/", middleware.ErrorHandling(controllers.Home))
	http.HandleFunc("/contact", middleware.ErrorHandling(controllers.Contact))
	http.HandleFunc("/checkout", middleware.ErrorHandling(controllers.Checkout))
	http.HandleFunc("/cart", middleware.ErrorHandling(controllers.Cart))
	http.HandleFunc("/blogs", middleware.ErrorHandling(controllers.Blogs))
	http.HandleFunc("/your-account", middleware.ErrorHandling(controllers.Account))

	// Blogs
	http.HandleFunc("/what-are-the-secrets-to-start-up-success", middleware.ErrorHandling(controllers.FirstBlog))
	http.HandleFunc("/mans-fashion-winter-sale", middleware.ErrorHandling(controllers.SecondBlog))
	http.HandleFunc("/women-fashion-festive", middleware.ErrorHandling(controllers.ThirdBlog))
	http.HandleFunc("/sed-adipiscing-ornare", middleware.ErrorHandling(controllers.ForthBlog))

	// Products
	http.HandleFunc("/products-list", middleware.ErrorHandling(controllers.ProductsList))
	http.HandleFunc("/first-cloth", middleware.ErrorHandling(controllers.FirstCloth))
	http.HandleFunc("/second-cloth", middleware.ErrorHandling(controllers.SecondCloth))
	http.HandleFunc("/third-cloth", middleware.ErrorHandling(controllers.ThirdCloth))
	http.HandleFunc("/forth-cloth", middleware.ErrorHandling(controllers.ForthCloth))
	http.HandleFunc("/fifth-cloth", middleware.ErrorHandling(controllers.FifthCloth))
	http.HandleFunc("/sixth-cloth", middleware.ErrorHandling(controllers.SixthCloth))
	http.HandleFunc("/seventh-cloth", middleware.ErrorHandling(controllers.SeventhCloth))
	http.HandleFunc("/eighth-cloth", middleware.ErrorHandling(controllers.EighthCloth))
	http.HandleFunc("/ninth-cloth", middleware.ErrorHandling(controllers.NinthCloth))
	http.HandleFunc("/awesome-pink-show", middleware.ErrorHandling(controllers.AwesomePinkShow))
	http.HandleFunc("/polo-dress-for-women", middleware.ErrorHandling(controllers.PoloDressForWomen))
	http.HandleFunc("/women-hot-collection", middleware.ErrorHandling(controllers.WomenHotCollection))
	http.HandleFunc("/awesome-cap-for-women", middleware.ErrorHandling(controllers.AwesomeCapForWomen))
	http.HandleFunc("/women-pant-collections", middleware.ErrorHandling(controllers.WomenPantCollections))
	http.HandleFunc("/awesome-bag-collection", middleware.ErrorHandling(controllers.AwesomeBagsCollection1))
	http.HandleFunc("/awesome-bags-collection", middleware.ErrorHandling(controllers.AwesomeBagsCollection2))
	http.HandleFunc("/black-sunglass-for-women", middleware.ErrorHandling(controllers.BlackSunglassForWomen))

	// User
	http.HandleFunc("/password", middleware.ErrorHandling(controllers.Password))
	http.HandleFunc("/confirm-email", middleware.ErrorHandling(controllers.ConfirmEmail))
	http.HandleFunc("/login", middleware.ErrorHandling(controllers.Login))
	http.HandleFunc("/logout", controllers.Logout)

	// Account
	http.HandleFunc("/your-profile", middleware.ErrorHandling(controllers.Profile))
	http.HandleFunc("/update-profile", middleware.ErrorHandling(controllers.UpdateProfile))
	http.HandleFunc("/confirm-account-password", middleware.ErrorHandling(controllers.DeleteProfile))
	http.HandleFunc("/your-payment", middleware.ErrorHandling(controllers.Payment))
	http.HandleFunc("/your-payment-method", middleware.ErrorHandling(controllers.PaymentMethod))
	http.HandleFunc("/your-payment-details", middleware.ErrorHandling(controllers.PaymentDetails))
	http.HandleFunc("/update-payment", middleware.ErrorHandling(controllers.UpdatePayment))
	http.HandleFunc("/confirm-payment-password", middleware.ErrorHandling(controllers.DeletePayment))
	http.HandleFunc("/your-orders", middleware.ErrorHandling(controllers.Orders))

	// Footer
	http.HandleFunc("/about-us", middleware.ErrorHandling(controllers.AboutUs))
	http.HandleFunc("/terms-conditions", middleware.ErrorHandling(controllers.TermsConditions))
	http.HandleFunc("/shipping", middleware.ErrorHandling(controllers.Shipping))
	http.HandleFunc("/privacy-policy", middleware.ErrorHandling(controllers.PrivacyPolicy))
	http.HandleFunc("/faq", middleware.ErrorHandling(controllers.Faq))

	http.HandleFunc("/search", middleware.ErrorHandling(controllers.SearchProducts))

	// To handle the static files
	fileServer := http.FileServer(http.Dir("./views/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
}
