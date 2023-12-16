package component

import (
	"database/sql"
	"e-wallet/internal/config"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

func GetDbConnection(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Database.Host,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.DbName,
		cnf.Database.Port,
	)

	connection, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error when open connection %s", err.Error())
	}

	return connection

}
