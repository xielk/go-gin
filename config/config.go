package config

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Jwt `mapstructure:"jwt"`
}

type Server struct {
	Port  int  `mapstructure:"port"`
	Debug bool `mapstructure:"debug"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Jwt struct {
	Key  string  `mapstructure:"key"`
}