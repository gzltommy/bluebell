package settings

type Config struct {
	Name            string `mapstructure:"name"`
	Mode            string `mapstructure:"mode"`
	Version         string `mapstructure:"version"`
	Port            int    `mapstructure:"port"`
	LimitConnection int    `mapstructure:"limit_connection"`
	StartTime       string `mapstructure:"start_time"`
	MachineID       int    `mapstructure:"machine_id"`

	Log   *LogConfig   `mapstructure:"log"`
	MySQL *MySQLConfig `mapstructure:"mysql"`
	Redis *RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
