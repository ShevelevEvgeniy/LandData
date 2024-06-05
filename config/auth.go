package config

type Auth struct {
	ApiKey string `envconfig:"X_API_KEY" env-required:"true"`
}
