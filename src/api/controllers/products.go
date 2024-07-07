package controllers

import (
	"api/src/api/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFormatted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na convesão do preço: ", err)
		}
		quantityFormatted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na convesão da quantidade: ", err)
		}

		models.SaveProduct(name, description, priceFormatted, quantityFormatted)
	}

	redirect(w, r)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)

	redirect(w, r)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	if idProduct == "" {
		fmt.Println("ID não informado")
		redirect(w, r)
		return
	}

	productSelected := models.GetProductSelected(idProduct)
	templates.ExecuteTemplate(w, "EditProduct", productSelected)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idFormatted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int", err.Error())
		}

		priceFormatted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na convesão do preço: ", err)
		}
		quantityFormatted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na convesão da quantidade: ", err)
		}

		models.UpdateProduct(idFormatted, name, description, priceFormatted, quantityFormatted)
		
		redirect(w, r)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	const verbHttp = 301
	http.Redirect(w, r, "/", verbHttp)
}