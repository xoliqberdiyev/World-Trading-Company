package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/XoliqberdiyevBehruz/wtc_backend/command"
	_ "github.com/XoliqberdiyevBehruz/wtc_backend/docs"
	"github.com/XoliqberdiyevBehruz/wtc_backend/services/about_company"
	"github.com/XoliqberdiyevBehruz/wtc_backend/services/common"
	"github.com/XoliqberdiyevBehruz/wtc_backend/services/common_admin"
	product "github.com/XoliqberdiyevBehruz/wtc_backend/services/product_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/services/user_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
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

	swaggerUrl := fmt.Sprintf("%v://%v/swagger/doc.json", utils.GetString("SWAGGER_SSL", "http"), utils.GetString("SWAGGER_HOST", "localhost:8000"))
	// swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl),
	))
	// for image read
	r.Get("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))).ServeHTTP)

	// user
	userStore := user_admin.NewStore(s.db)
	userHandler := user_admin.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	// common
	commonStore := common.NewStore(s.db)
	commonHandler := common.NewHandler(commonStore)
	commonHandler.RegisterRoutes(r)

	// common-admin
	commonAdminStore := common_admin.NewStore(s.db)
	commonAdminHandler := common_admin.NewHandler(commonAdminStore, userStore)
	commonAdminHandler.RegisterRoutes(r)

	// company-admin
	companyAdminStore := about_company.NewStore(s.db)
	companyAdminHandler := about_company.NewHandler(companyAdminStore, userStore)
	companyAdminHandler.RegisterRoutes(r)

	// product
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegsiterRoutes(r)

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
