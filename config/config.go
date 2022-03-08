package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

const (
	MINEPIN_DEBUG_SUFFIX     = "_debug"
	MINEPIN_PORT             = "port"
	MINEPIN_MAX_PING_COUNT   = "max_ping_count"
	MINEPIN_RUNMODE          = "runmode"
	MINEPIN_DB_TYPE          = "db.type"
	MINEPIN_DB_NAME          = "db.name"
	MINEPIN_DB_ADDR          = "db.addr"
	MINEPIN_DB_USERNAME      = "db.username"
	MINEPIN_DB_PASSWORD      = "db.password"
	MINEPIN_LOG_LEVEL        = "log.level"
	MINEPIN_LOG_FILE_NAME    = "log.file_name"
	MINEPIN_LOG_MAX_SIZE_MB  = "log.max_size_mb"
	MINEPIN_LOG_MAX_FILE_NUM = "log.max_file_num"
	MINEPIN_LOG_MAX_FILE_DAY = "log.max_file_day"
	MINEPIN_LOG_COMPRESS     = "log.compress"
	MINEPIN_LOG_STDOUT       = "log.stdout"
	MINEPIN_LOG_ONLY_STDOUT  = "log.only_stdout"

	MINEPIN_DEFAULT_PORT           = "8080"
	MINEPIN_DEFAULT_MAX_PING_COUNT = 3
	MINEPIN_DEFAULT_RUNMODE        = "release"
)

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()          // 读取匹配的环境变量，环境变量优先级最高
	viper.SetEnvPrefix("MINEPIN") // 读取环境变量的前缀为 MINEPIN
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 设定一些默认值
	viper.SetDefault(MINEPIN_PORT, MINEPIN_DEFAULT_PORT)
	viper.SetDefault(MINEPIN_MAX_PING_COUNT, MINEPIN_DEFAULT_MAX_PING_COUNT)
	viper.SetDefault(MINEPIN_RUNMODE, MINEPIN_DEFAULT_RUNMODE)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {})
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func Get(name string) interface{}  { return viper.Get(name) }
func GetString(name string) string { return strings.TrimSpace(viper.GetString(name)) }
func GetInt(name string) int       { return viper.GetInt(name) }
func GetBool(name string) bool     { return viper.GetBool(name) }

// Get MinePin Debug Suffix
func MPDSx() string {
	if GetMinePinRunMode() == gin.DebugMode {
		return MINEPIN_DEBUG_SUFFIX
	} else {
		return ""
	}
}

func GetMinePinPort() string      { return GetString(MINEPIN_PORT + MPDSx()) }
func GetMinePinRunMode() string   { return GetString(MINEPIN_RUNMODE) }
func GetMinePinMaxPingCount() int { return GetInt(MINEPIN_MAX_PING_COUNT) }

func GetMinePinLogLevel() string    { return GetString(MINEPIN_LOG_LEVEL) }
func GetMinePinLogFileName() string { return GetString(MINEPIN_LOG_FILE_NAME) }
func GetMinePinLogMaxSizeMb() int   { return GetInt(MINEPIN_LOG_MAX_SIZE_MB) }
func GetMinePinLogMaxFileNum() int  { return GetInt(MINEPIN_LOG_MAX_FILE_NUM) }
func GetMinePinLogMaxFileDay() int  { return GetInt(MINEPIN_LOG_MAX_FILE_DAY) }
func GetMinePinLogCompress() bool   { return GetBool(MINEPIN_LOG_COMPRESS) }
func GetMinePinLogStdout() bool     { return GetBool(MINEPIN_LOG_STDOUT) }
func GetMinePinLogOnlyStdout() bool { return GetBool(MINEPIN_LOG_ONLY_STDOUT) }

func GetMinePinDbType() string     { return GetString(MINEPIN_DB_TYPE + MPDSx()) }
func GetMinePinDbName() string     { return GetString(MINEPIN_DB_NAME + MPDSx()) }
func GetMinePinDbAddr() string     { return GetString(MINEPIN_DB_ADDR + MPDSx()) }
func GetMinePinDbUserName() string { return GetString(MINEPIN_DB_USERNAME + MPDSx()) }
func GetMinePinDbPassWord() string { return GetString(MINEPIN_DB_PASSWORD + MPDSx()) }
