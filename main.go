package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/XoliqberdiyevBehruz/wtc_backend/api"
	"github.com/XoliqberdiyevBehruz/wtc_backend/config"
	_ "github.com/lib/pq"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description
// @security BearerAuth
func main() {
	cnf := config.Load()
	psqlUrl := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cnf.Postgres.Host, cnf.Postgres.Port, cnf.Postgres.User, cnf.Postgres.Password, cnf.Postgres.Databasse,
	)
	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("DB: connection successfully")
	}
	server := api.NewServer(psqlConn, "0.0.0.0:8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
