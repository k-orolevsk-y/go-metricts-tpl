package env

import "github.com/caarlos0/env/v6"

var Config struct {
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

func Init() error {
	if err := env.Parse(&Config); err != nil {
		return err
	}

	return nil
}
