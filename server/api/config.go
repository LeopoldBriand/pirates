package api

type Config struct {
	HTTPAddress string
}

func LoadConfig() Config {
	return Config{
		HTTPAddress: "0.0.0.0:3000", //TODO: Load from env
	}
}
