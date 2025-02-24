package config

type Config struct {
	PG_ADDR string `env:"PG_ADDR"`
	PG_USER string `env:"PG_USER"`
	PG_PASSWORD string `env:"PG_PASSWORD"`
	PG_APP_NAME string `env:"PG_APP_NAME"`
}