package configs

import (
	"os"

	"github.com/lbrulet/API-AWS-RDS/models"
	"github.com/lbrulet/gonfig"
)

// Config application configuration
var Config models.Configuration

func init() {
	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		if os.Getenv("ENVIRONMENT") == "LOCAL" {
			if err := gonfig.GetConf(pwd+"/configs/local/app.config.json", &Config); err != nil {
				panic(err)
			}
		} else {
			if err := gonfig.GetConf(pwd+"/configs/prod/app.config.json", &Config); err != nil {
				panic(err)
			}
		}
	}
}
