package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alexandrebrunodias/store/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Index index
func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", models.FindAllProducts())
}

// New new
func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

// Insert insert
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)

		if err != nil {
			log.Print("Error while converting price: ", err)
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))

		if err != nil {
			log.Print("Error while converting quantity: ", err)
		}

		models.CreateProduct(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}

// Delete delete
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}

// Update update
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method == "POST" {
		models.Update(r.FormValue("id"),
			r.FormValue("name"),
			r.FormValue("description"),
			r.FormValue("price"),
			r.FormValue("quantity"))

		http.Redirect(w, r, "/", 301)
		return
	}

	product := models.FindByID(id)
	templates.ExecuteTemplate(w, "Update", product)
}
