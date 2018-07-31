package repo

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

//variavel singleton que armazena a conexao
var db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil
	var (
		password      = flag.String("password", "paftug", "the database password")
		port     *int = flag.Int("port", 1433, "the database port")
		server        = flag.String("server", "localhost", "the database server")
		user          = flag.String("user", "sa", "the database user")
		database      = flag.String("database", "DEALER", "the database name")
	)
	flag.Parse()

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	Db, err := sql.Open("mssql", connString)
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

//GetDBConnection Obtem a conexao com o banco de dados
func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexaoComBancoDeDadosSQL()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}
