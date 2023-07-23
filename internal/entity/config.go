package entity

type ServerConfig struct {
	DatabaseDsn string `env:"DB_DSN" envDefault:"postgres://postgres@localhost:5432/postgres"`
	ServerAddr  string `env:"SRV_ADDR" envDefault:"0.0.0.0:8080"`
}
