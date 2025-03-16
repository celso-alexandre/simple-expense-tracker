package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/celso-alexandre/api/docs"
	"github.com/celso-alexandre/api/router"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title My API
// @version 1.0
// @description This is a sample API using Swagger.
// @host localhost:8080
// @BasePath /api/v1

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ServeSwagger(w http.ResponseWriter, r *http.Request) {
	httpSwagger.WrapHandler(w, r)
}

func main() {
	mux := router.SetupRouter()
	port := "8080"

	mux.HandleFunc("/swagger", httpSwagger.WrapHandler)
	mux.HandleFunc("/", httpSwagger.WrapHandler)

	server := http.Server{
		Addr:    ":" + port,
		Handler: corsMiddleware(mux),
	}

	fmt.Println("🚀 Server running on http://localhost:" + port)
	log.Fatal(server.ListenAndServe())
}
