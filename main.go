package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func requestLogger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		log.Printf("%s %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func main(){
	
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("PORT env variable not set")
	}
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://", "https://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Use(requestLogger)


	v1Router.Get("/healthz",handleReadiness)
	v1Router.Get("/err",handleErr)

	router.Mount("/v1",v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Sever is starting on %v", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}



























