package config

type Config struct {
	Port      int    `env:"PORT,default=8080"`
	Env       string `env:"ENV,default=development"`
	LogFormat string `env:"LOG_FORMAT,default=text"`
	LogLevel  string `env:"LOG_SEVERITY,default=info"`
}
