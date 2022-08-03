package setting

type Config struct {
	AppName         string `mapstructure:"app_name"`
	Mode            string `mapstructure:"mode"`
	Version         string `mapstructure:"version"`
	Port            int    `mapstructure:"port"`
	LimitConnection int    `mapstructure:"limit_connection"`
	StartTime       string `mapstructure:"start_time"`
	MachineID       int    `mapstructure:"machine_id"`

	Auth  *AuthConfig  `mapstructure:"auth"`
	Log   *LogConfig   `mapstructure:"log"`
	MySQL *MySQLConfig `mapstructure:"mysql"`
	Redis *RedisConfig `mapstructure:"redis"`
}

type (
	AuthConfig struct {
		JwtSecret string `mapstructure:"jwt_secret"`
		JwtExpire int64  `mapstructure:"jwt_expire"`
	}

	LogConfig struct {
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	}

	MySQLConfig struct {
		Host         string `mapstructure:"host"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Port         int    `mapstructure:"port"`
		Dbname       string `mapstructure:"dbname"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
	}

	RedisConfig struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
		Port     int    `mapstructure:"port"`
		DB       int    `mapstructure:"db"`
		PoolSize int    `mapstructure:"pool_size"`
	}
)
