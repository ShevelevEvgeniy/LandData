package config

type IpInfo struct {
	IpInfoToken  string `env:"IP_INFO_TOKEN" envDefault:""`
	LocalMachine bool   `env:"IP_INFO_LOCAL_MACHINE" envDefault:"true"`
	Country      string `env:"IP_INFO_COUNTRY" envDefault:"RU"`
}
