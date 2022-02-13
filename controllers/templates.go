package controllers

import "html/template"

// Attaching the base.html file to each template.
func MakeTemplate(path string) *template.Template {
	files := []string{path, "views/templates/single-tmpl/base.html"}
	return template.Must(template.ParseFiles(files...))
}

// Header Templates
var (
	HomeTmpl     = MakeTemplate("views/templates/header/index.html")
	ContactTmpl  = MakeTemplate("views/templates/header/contact.html")
	CheckoutTmpl = MakeTemplate("views/templates/header/checkout.html")
	CartTmpl     = MakeTemplate("views/templates/header/cart.html")
	BlogsTmpl    = MakeTemplate("views/templates/header/blogs.html")
	AccountTmpl  = MakeTemplate("views/templates/header/account.html")
)

// Blog Templates
var (
	FirstBlogTmpl  = MakeTemplate("views/templates/blogs/first-blog.html")
	SecondBlogTmpl = MakeTemplate("views/templates/blogs/second-blog.html")
	ThirdBlogTmpl  = MakeTemplate("views/templates/blogs/third-blog.html")
	ForthBlogTmpl  = MakeTemplate("views/templates/blogs/forth-blog.html")
)

// Product Details Templates
var (
	ProductDetailsTmpl = MakeTemplate("views/templates/single-tmpl/product-details.html")
)

// User Templates
var (
	PasswordTmpl     = MakeTemplate("views/templates/user/password.html")
	ConfirmEmailTmpl = MakeTemplate("views/templates/user/confirm-email.html")
	LoginTmpl        = MakeTemplate("views/templates/user/login.html")
)

// Account Templates
var (
	ProfileTmpl        = MakeTemplate("views/templates/account/profile.html")
	UpdateProfileTmpl  = MakeTemplate("views/templates/account/update-profile.html")
	DeleteProfileTmpl  = MakeTemplate("views/templates/account/delete-profile.html")
	PaymentTmpl        = MakeTemplate("views/templates/account/payment.html")
	PaymentMethodTmpl  = MakeTemplate("views/templates/account/payment-method.html")
	PaymentDetailsTmpl = MakeTemplate("views/templates/account/payment-details.html")
	UpdatePaymentTmpl  = MakeTemplate("views/templates/account/update-payment.html")
	DeletePaymentTmpl  = MakeTemplate("views/templates/account/delete-payment.html")
	OrdersTmpl         = MakeTemplate("views/templates/account/orders.html")
)

// Footer Templates
var (
	FooterDetailsTmpl = MakeTemplate("views/templates/footer/footer-details.html")
	FaqDetailsTmpl    = MakeTemplate("views/templates/footer/faq-details.html")
)

// Single Templates
var (
	PageErrorTmpl    = MakeTemplate("views/templates/single-tmpl/page-error.html")
	ProductsListTmpl = MakeTemplate("views/templates/single-tmpl/products-list.html")
)
