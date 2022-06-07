package config

type (
	ServerConfig struct {
		Port          uint16 `mapstructure:"port" toml:"port"`
		Host          string `mapstructure:"host" toml:"host"`
		SymmetricKey  string `mapstructure:"sign_key" toml:"sign_key"`
		IsDevelopment bool   `mapstructure:"is_debug" toml:"is_debug"`
	}
	DatabaseConfig struct {
		Host     string `mapstructure:"host" toml:"host"`
		Port     uint16 `mapstructure:"port" toml:"port"`
		User     string `mapstructure:"username" toml:"username"`
		Password string `mapstructure:"user_password" toml:"user_password"`
		Name     string `mapstructure:"database_name" toml:"database_name"`
	}
	AppConfig struct {
		Server   ServerConfig `mapstructure:"server" toml:"server"`
		Database struct {
			Abz1 DatabaseConfig `mapstructure:"abz_1" toml:"abz_1"`
			Xyc2 DatabaseConfig `mapstructure:"xyc_2" toml:"xyc_2"`
		} `mapstructure:"databases" toml:"databases"`
	}
)
