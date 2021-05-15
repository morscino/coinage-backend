package config

type Config struct {
	DB  DatabaseConfig
	App AppConfig
}

type DatabaseConfig struct {
	Host     string `envconfig:"GIGO_DB_HOST"`
	Name     string `envconfig:"GIGO_DB_NAME"`
	Dialect  string `envconfig:"GIGO_DB_DIALECT"`
	User     string `envconfig:"GIGO_DB_USER"`
	Password string `envconfig:"GIGO_DB_PASSWORD"`
	Port     string `envconfig:"GIGO_DB_PORT"`
	SSLMode  string `envconfig:"GIGO_DB_SSL_MODE"`
}

type AppConfig struct {
	Port string `envconfig:"GIGO_PORT"`
}
