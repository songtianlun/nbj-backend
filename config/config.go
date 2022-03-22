package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mingin/utils"
	"strings"
)

type Config struct {
	Name string
}

const (
	MINEGIN_DEBUG_SUFFIX       = "_debug"
	MINEGIN_PORT               = "port"
	MINEGIN_MAX_PING_COUNT     = "max_ping_count"
	MINEGIN_RUNMODE            = "runmode"
	MINEGIN_JWT_ACCESS_SECRET  = "jwt.access_secret"
	MINEGIN_JWT_REFRESH_SECRET = "jwt.refresh_secret"
	MINEGIN_DB_TYPE            = "db.type"
	MINEGIN_DB_NAME            = "db.name"
	MINEGIN_DB_ADDR            = "db.addr"
	MINEGIN_DB_USERNAME        = "db.username"
	MINEGIN_DB_PASSWORD        = "db.password"
	MINEGIN_LOG_LEVEL          = "log.level"
	MINEGIN_LOG_FILE_NAME      = "log.file_name"
	MINEGIN_LOG_MAX_SIZE_MB    = "log.max_size_mb"
	MINEGIN_LOG_MAX_FILE_NUM   = "log.max_file_num"
	MINEGIN_LOG_MAX_FILE_DAY   = "log.max_file_day"
	MINEGIN_LOG_COMPRESS       = "log.compress"
	MINEGIN_LOG_STDOUT         = "log.stdout"
	MINEGIN_LOG_ONLY_STDOUT    = "log.only_stdout"

	MINEGIN_DEFAULT_CONFIG_PATH = "./"
	MINEGIN_DEFAULT_CONFIG_NAME = "config"
	MINEGIN_DEFAULT_CONFIG_TYPE = "yaml"
	MINEGIN_DEFAULT_CONFIG_FILE = MINEGIN_DEFAULT_CONFIG_PATH + MINEGIN_DEFAULT_CONFIG_NAME +
		POINT + MINEGIN_DEFAULT_CONFIG_TYPE
	MINEGIN_DEFAULT_PORT               = "6000"
	MINEGIN_DEFAULT_MAX_PING_COUNT     = 3
	MINEGIN_DEFAULT_RUNMODE            = "release"
	MINEGIN_DEFAULT_DB_TYPE            = "sqlite3"
	MINEGIN_DEFAULT_DB_ADDR            = "./minepin.db"
	MINEGIN_DEFAULT_LOG_LEVEL          = "info"
	MINEGIN_DEFAULT_LOG_COMPRESS       = false
	MINEGIN_DEFAULT_LOG_ONLY_STDOUT    = true
	MINEGIN_DEFAULT_JWT_ACCESS_SECRET  = "c4ce87cb12d0d6a65458c0cb38779cec"
	MINEGIN_DEFAULT_JWT_REFRESH_SECRET = "25e0aacbe8e52824bfb397e569f0ab16e1005c397540c95b277df27eaec97ff6"
	POINT                              = "."
)

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
		fmt.Printf("run with abstract config %s\n", c.Name)
	} else if isExist, _ := utils.PathExists(MINEGIN_DEFAULT_CONFIG_FILE); isExist {
		fmt.Printf("run with default config %s\n", MINEGIN_DEFAULT_CONFIG_FILE)
		viper.AddConfigPath(MINEGIN_DEFAULT_CONFIG_PATH)
		viper.SetConfigName(MINEGIN_DEFAULT_CONFIG_NAME)
		viper.SetConfigType(MINEGIN_DEFAULT_CONFIG_TYPE)
	}

	viper.AutomaticEnv()          // 读取匹配的环境变量，环境变量优先级最高
	viper.SetEnvPrefix("MINEGIN") // 读取环境变量的前缀为 MINEGIN
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 设定一些默认值
	viper.SetDefault(MINEGIN_PORT, MINEGIN_DEFAULT_PORT)
	viper.SetDefault(MINEGIN_MAX_PING_COUNT, MINEGIN_DEFAULT_MAX_PING_COUNT)
	viper.SetDefault(MINEGIN_RUNMODE, MINEGIN_DEFAULT_RUNMODE)
	viper.SetDefault(MINEGIN_DB_TYPE, MINEGIN_DEFAULT_DB_TYPE)
	viper.SetDefault(MINEGIN_DB_ADDR, MINEGIN_DEFAULT_DB_ADDR)
	viper.SetDefault(MINEGIN_LOG_LEVEL, MINEGIN_DEFAULT_LOG_LEVEL)
	viper.SetDefault(MINEGIN_LOG_COMPRESS, MINEGIN_DEFAULT_LOG_COMPRESS)
	viper.SetDefault(MINEGIN_LOG_ONLY_STDOUT, MINEGIN_DEFAULT_LOG_ONLY_STDOUT)
	viper.SetDefault(MINEGIN_JWT_ACCESS_SECRET, MINEGIN_DEFAULT_JWT_ACCESS_SECRET)
	viper.SetDefault(MINEGIN_JWT_REFRESH_SECRET, MINEGIN_DEFAULT_JWT_REFRESH_SECRET)

	if isExist, _ := utils.PathExists(MINEGIN_DEFAULT_CONFIG_FILE); !isExist {
		return nil
	}
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

	//c.watchConfig() // 热加载当前无意义，且会在配置文件不存在时阻塞，暂时移除

	return nil
}

func Get(name string) interface{}  { return viper.Get(name) }
func GetString(name string) string { return strings.TrimSpace(viper.GetString(name)) }
func GetInt(name string) int       { return viper.GetInt(name) }
func GetBool(name string) bool     { return viper.GetBool(name) }

// Get MineGin Debug Suffix
func MPDSx() string {
	if GetMineGinRunMode() == gin.DebugMode {
		return MINEGIN_DEBUG_SUFFIX
	} else {
		return ""
	}
}

func GetMineGinPort() string             { return GetString(MINEGIN_PORT + MPDSx()) }
func GetMineGinRunMode() string          { return GetString(MINEGIN_RUNMODE) }
func GetMineGinMaxPingCount() int        { return GetInt(MINEGIN_MAX_PING_COUNT) }
func GetMineGinJwtAccessSecret() string  { return GetString(MINEGIN_JWT_ACCESS_SECRET) }
func GetMineGinJwtRefreshSecret() string { return GetString(MINEGIN_JWT_REFRESH_SECRET) }

func GetMineGinLogLevel() string    { return GetString(MINEGIN_LOG_LEVEL) }
func GetMineGinLogFileName() string { return GetString(MINEGIN_LOG_FILE_NAME) }
func GetMineGinLogMaxSizeMb() int   { return GetInt(MINEGIN_LOG_MAX_SIZE_MB) }
func GetMineGinLogMaxFileNum() int  { return GetInt(MINEGIN_LOG_MAX_FILE_NUM) }
func GetMineGinLogMaxFileDay() int  { return GetInt(MINEGIN_LOG_MAX_FILE_DAY) }
func GetMineGinLogCompress() bool   { return GetBool(MINEGIN_LOG_COMPRESS) }
func GetMineGinLogStdout() bool     { return GetBool(MINEGIN_LOG_STDOUT) }
func GetMineGinLogOnlyStdout() bool { return GetBool(MINEGIN_LOG_ONLY_STDOUT) }

func GetMineGinDbType() string     { return GetString(MINEGIN_DB_TYPE + MPDSx()) }
func GetMineGinDbName() string     { return GetString(MINEGIN_DB_NAME + MPDSx()) }
func GetMineGinDbAddr() string     { return GetString(MINEGIN_DB_ADDR + MPDSx()) }
func GetMineGinDbUserName() string { return GetString(MINEGIN_DB_USERNAME + MPDSx()) }
func GetMineGinDbPassWord() string { return GetString(MINEGIN_DB_PASSWORD + MPDSx()) }
