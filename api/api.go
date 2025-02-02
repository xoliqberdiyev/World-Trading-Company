package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/XoliqberdiyevBehruz/wtc_backend/command"
	_ "github.com/XoliqberdiyevBehruz/wtc_backend/docs"
	"github.com/XoliqberdiyevBehruz/wtc_backend/services/admin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	db      *sql.DB
	address string
}

func NewServer(db *sql.DB, address string) *APIServer {
	return &APIServer{db: db, address: address}
}

func (s *APIServer) Run() error {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Frontend URL manzilingiz
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
	))
	// for image read
	r.Get("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))).ServeHTTP)

	// user
	userStore := admin.NewStore(s.db)
	userHandler := admin.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	// command
	if len(os.Args) > 1 && os.Args[1] == "createsuperuser" {
		command.CreateSuperUser(userStore)
		return nil
	}

	handler := corsHandler.Handler(r)
	log.Println("Listen on", s.address)
	address := s.address

	return http.ListenAndServe(address, handler)
}
