package main

import (
	"fmt"
	"log"

	"github.com/XoliqberdiyevBehruz/wtc_backend/api"
	"github.com/XoliqberdiyevBehruz/wtc_backend/config"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/jmoiron/sqlx"
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
	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Println(err)
	}
	server := api.NewServer(psqlConn.DB, utils.GetString("SERVER_PORT", "localhost:8000"))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
