package config

type (
	ServerConfig struct {
		Port         uint16 `mapstructure:"port"`
		Host         string `mapstructure:"host"`
		SymmetricKey string `mapstructure:"sign_key"`
	}
	DatabaseConfig struct {
		Host     string `mapstructure:"host"`
		Port     uint16 `mapstructure:"port"`
		User     string `mapstructure:"username"`
		Password string `mapstructure:"user_password"`
		Name     string `mapstructure:"database_name"`
	}
	AppConfig struct {
		Server   ServerConfig `mapstructure:"server"`
		Database struct {
			Abz1 DatabaseConfig `mapstructure:"abz_1"`
			Xyc2 DatabaseConfig `mapstructure:"xyc_2"`
		} `mapstructure:"databases"`
	}
)
