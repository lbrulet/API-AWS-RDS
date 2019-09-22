package main

import (
	"github.com/lbrulet/API-AWS-RDS/configs"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/routes"
)

func main() {
	database.InitDB()
	srv := routes.InitRouter()

	srv.Run(configs.Config.Port)
}
