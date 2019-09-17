package models

// Configuration struct is used to load configuration from file or environnement
type Configuration struct {
	Port             string `json:"app_port" env:"APP_PORT"`
	DatabaseName     string `json:"database_name" env:"DATABASE_NAME"`
	DatabasePassword string `json:"database_password" env:"DATABASE_PASSWORD"`
	DatabaseEndPoint string `json:"database_endpoint" env:"DATABASE_ENDPOINT"`
	DatabaseUser     string `json:"database_user" env:"DATABASE_USER"`
}
