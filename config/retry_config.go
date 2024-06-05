package config

import "time"

type RetryConfig struct {
	Attempts uint          `envconfig:"RETRY_ATTEMPTS" envDefault:"5"`
	Delay    time.Duration `envconfig:"RETRY_DELAY" envDefault:"5s"`
}
