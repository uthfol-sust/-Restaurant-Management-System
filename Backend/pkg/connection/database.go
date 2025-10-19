package connection

import (
	"database/sql"
	"fmt"
	"log"
	"restaurant-system/pkg/config"

	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func ConnectDB() {
	dbCfg := config.LocalConfig

	dbServer := mysql.NewConfig()

	dbServer.User = dbCfg.DBUser
	dbServer.Passwd = dbCfg.DBPassword
	dbServer.Net = dbCfg.DBNet
	dbServer.Addr = dbCfg.DBHost + ":" + dbCfg.DBPort
	dbServer.DBName = dbCfg.DBName

	d, err := sql.Open("mysql", dbServer.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	if err := d.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database..!")
	db = d
}


func Migrate() {

}

func GetDB() *sql.DB {
	if db == nil {
		ConnectDB()
		
	}
	return db
}
