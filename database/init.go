package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lbrulet/API-AWS-RDS/configs"
)

type DatabaseManager struct {
	DB *sql.DB
}

var DBManager DatabaseManager

func InitDB() {
	var err error
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		configs.Config.DatabaseUser, configs.Config.DatabasePassword, configs.Config.DatabaseEndPoint, configs.Config.DatabaseName,
	)

	DBManager.DB, err = sql.Open("mysql", dnsStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DBManager.DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[DB] Connected to: %s\n", configs.Config.DatabaseEndPoint)
}
