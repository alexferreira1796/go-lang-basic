package main

import (
	"api/src/api/routes"
	"net/http"

	_ "github.com/lib/pq"
)


func main () {
	routes.LoadingRoutes()
	
	http.ListenAndServe(":8000", nil)
}
