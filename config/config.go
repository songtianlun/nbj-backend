package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	Name string
}

const (
	MINEPIN_PORT           = "port"
	MINEPIN_MAX_PING_COUNT = "max_ping_count"
	MINEPIN_RUNMODE        = "runmode"

	MINEPIN_DEFAULT_PORT           = "8080"
	MINEPIN_DEFAULT_MAX_PING_COUNT = 3
	MINEPIN_DEFAULT_RUNMODE        = "release"
)

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()          // 读取匹配的环境变量
	viper.SetEnvPrefix("MINEPIN") // 读取环境变量的前缀为 MINEPIN
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 优先采用 PORT 环境变量的值
	if port := os.Getenv("PORT"); port != "" {
		viper.SetDefault(MINEPIN_PORT, port)
		viper.Set(MINEPIN_PORT, port) // 使用优先级最高的显示设置固化
	} else {
		viper.SetDefault(MINEPIN_PORT, MINEPIN_DEFAULT_PORT)
	}
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
func GetString(name string) string { return viper.GetString(name) }
func GetInt(name string) int       { return viper.GetInt(name) }
func GetBool(name string) bool     { return viper.GetBool(name) }
