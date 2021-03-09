package handlers

import (
	"net/http"
	"udemyProject/pkg/render"
)

// Homepage resresents / page
func Homepage(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html")

}

// Aboutpage resresents /about page
func Aboutpage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
