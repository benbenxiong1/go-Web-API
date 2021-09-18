package config

// DBConfig 数据库链接配置项
type DBConfig struct {
	Hostname string
	Username string
	Password string
	Database string
	Port string
}


func init() {
	Config.DBConfig =  DBConfig{
		Hostname: "127.0.0.1",
		Username: "root",
		Password: "123456",
		Database: "blog_laravel",
		Port:     "3305",
	}
}




