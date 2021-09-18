package config


type config struct {
	DBConfig DBConfig
	Host string
}


var Config config

func init() {
	Config.Host = ""
}
