package settings

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Sql    SQLSetting    `mapstructure:"sql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  CacheSetting  `mapstructure:"redis"`
	JWT    JWTSetting    `mapstructure:"jwt"`
}

type ServerSetting struct {
	Port        int    `mapstructure:"port"`
	Mode        string `mapstructure:"mode"`
	Level       string `mapstructure:"level"`
	ServiceName string `mapstructure:"service_name"`
	Environment string `mapstructure:"environment"`
}

type CacheSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type SQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	FileOutput bool   `mapstructure:"file_output"`
	LogDir     string `mapstructure:"log_dir"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

// JWT Settings
type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
}
