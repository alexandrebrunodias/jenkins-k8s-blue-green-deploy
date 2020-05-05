package routes

import (
	"net/http"

	"github.com/alexandrebrunodias/store/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Load Routes
func Load() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/update", controllers.Update)

	http.HandleFunc("/health", controllers.Health)
	http.Handle("/metrics", promhttp.Handler())

}
