package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/poornapragnyah/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

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

	db := os.Getenv("DB_URL")
	if db == ""{
		log.Fatal("DB_URL env variable not set")
	}
	
	conn, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Can't connect to database: ",err)
	}

	queries := database.New(conn)
	apiCfg := &apiConfig{
		DB: queries,
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
	v1Router.Post("/user",apiCfg.handleCreateUser)

	router.Mount("/v1",v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Sever is starting on: %v", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}



























