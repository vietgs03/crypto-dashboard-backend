package settings

type ServerSetting struct {
	Port        int    `mapstructure:"PORT"`
	Mode        string `mapstructure:"MODE"`
	Level       string `mapstructure:"LEVEL"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

type CacheSetting struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
	Database int    `mapstructure:"DATABASE"`
	PoolSize int    `mapstructure:"POOL_SIZE"`
}

type SQLSetting struct {
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"PORT"`
	Username        string `mapstructure:"USERNAME"`
	Password        string `mapstructure:"PASSWORD"`
	DBname          string `mapstructure:"DBNAME"`
	SSLMode         string `mapstructure:"SSL_MODE"`
	MaxConnIdleTime uint32 `mapstructure:"MAX_CONN_IDLE_TIME"`
	MaxConnLifetime uint64 `mapstructure:"MAX_CONN_LIFE_TIME"`
	MaxConns        uint8  `mapstructure:"MAX_CONNS"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"LEVEL"`
	FileOutput bool   `mapstructure:"FILE_OUTPUT"`
	LogDir     string `mapstructure:"LOG_DIR"`
	MaxSize    int    `mapstructure:"MAX_SIZE"`
	MaxAge     int    `mapstructure:"MAX_AGE"`
	MaxBackups int    `mapstructure:"MAX_BACKUPS"`
	Compress   bool   `mapstructure:"COMPRESS"`
}

// JWT Settings
type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
}
