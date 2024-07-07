package routes

import (
	"api/src/api/controllers"
	"net/http"
)

func LoadingRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novo-produto", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}