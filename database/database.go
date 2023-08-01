package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func Conectar() (*sql.DB, error) {
	erro := godotenv.Load()

	if erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	var (
		password = os.Getenv("MSSQL_DB_PASSWORD")
		user     = os.Getenv("MSSQL_DB_USER")
		database = os.Getenv("MSSQL_DB_DATABASE")
		host     = os.Getenv("MSSQL_DB_HOST")
		port     = os.Getenv("MSSQL_DB_PORT")
	)

	stringConexao := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", host, user, password, port, database)

	return sql.Open("mssql", stringConexao)

}
