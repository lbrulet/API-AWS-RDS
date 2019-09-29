package main

import (
	"fmt"

	"github.com/lbrulet/API-AWS-RDS/configs"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/routes"
)

func main() {
	if err := database.InitDB(); err != nil {
		fmt.Println(err)
	} else {
		srv := routes.InitRouter()
		srv.Run(configs.Config.Port)
	}
}
