package config

import "time"

type RetryConfig struct {
	Attempts uint          `env:"RETRY_ATTEMPTS" envDefault:"5"`
	Delay    time.Duration `env:"RETRY_DELAY" envDefault:"5s"`
}
