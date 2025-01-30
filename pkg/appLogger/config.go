package appLogger

type LogConfig struct {
	Level       string
	ServiceName string
	Environment string
	FileOutput  bool
	LogDir      string
	MaxSize     int  // megabytes
	MaxAge      int  // days
	MaxBackups  int  // number of files
	Compress    bool // compress rotated files
}

func DefaultConfig() *LogConfig {
	return &LogConfig{
		Level:      "info",
		FileOutput: true,
		LogDir:     "../../logs",
		MaxSize:    100, // 100MB
		MaxAge:     7,   // 7 days
		MaxBackups: 10,  // 10 files
		Compress:   true,
	}
}
