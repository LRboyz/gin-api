package config

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	LogsAddress string      `mapstructure:"logsAddress"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"Port"`
	Name     string `mapstructure:"Name"`
	Password string `mapstructure:"Password"`
	DBName   string `mapstructure:"DBName"`
}
