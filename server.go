package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lbrulet/API-AWS-RDS/configs"
)

func main() {
	// Create the MySQL DNS string for the DB connection
	// user:password@protocol(endpoint)/dbname?<params>
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		configs.Config.DatabaseUser, configs.Config.DatabasePassword, configs.Config.DatabaseEndPoint, configs.Config.DatabaseName,
	)

	// Use db to perform SQL operations on database
	db, err := sql.Open("mysql", dnsStr)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
