package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/tintin007k/rss-agg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is not set")
	}
	fmt.Println("port: ", port)

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	v1router.Post("/users", apiCfg.handlerUsersCreate) //new user creation endpoint
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedsCreate))
	v1router.Get("/feeds", apiCfg.handlerFeedsGet)

	v1router.Post("/feedfollows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsCreate))
	v1router.Get("/feedfollows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	v1router.Delete(
		"/feedfollows/{feedFollowID}",
		apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsDelete),
	)

	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("server starting on port 8000")

	srvErr := srv.ListenAndServe()

	if err != nil {
		log.Fatal(srvErr)
	}

}
