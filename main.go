package main

import (
	"net/http"

	"github.com/alexandrebrunodias/store/db"
	"github.com/alexandrebrunodias/store/routes"
)

func main() {
	db.CreateDatabase()
	routes.Load()
	http.ListenAndServe(":8080", nil
}
