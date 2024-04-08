package internal

import (
	"fmt"
	"os"
	"path"
	"time"

	authcfg "github.com/ronannnn/auth/cfg"
	"github.com/ronannnn/infra/cfg"
)

const (
	// app
	ApplicationName    = "GX-WMS"
	ApplicationVersion = "0.0.1"

	// config
	ConfigDir         = "configs"
	ConfigEnvKey      = "GX_WMS_CONFIG"
	ConfigDefaultFile = "config.default.toml"
	ConfigTestFile    = "config.test.toml"
	ConfigReleaseFile = "config.release.toml"
)

type Cfg struct {
	Sys  cfg.Sys      `mapstructure:"sys"`
	Log  cfg.Log      `mapstructure:"log"`
	Auth cfg.Auth     `mapstructure:"auth"`
	Db   cfg.Db       `mapstructure:"db"`
	Dfs  cfg.Dfs      `mapstructure:"dfs"`
	User authcfg.User `mapstructure:"user"`
}

func ProvideLogCfg() *cfg.Log {
	return &GlCfg.Log
}

func ProvideSysCfg() *cfg.Sys {
	return &GlCfg.Sys
}

func ProvideDbCfg() *cfg.Db {
	return &GlCfg.Db
}

func ProvideAuthCfg() *cfg.Auth {
	return &GlCfg.Auth
}

func ProvideUserCfg() *authcfg.User {
	return &GlCfg.User
}

var GlCfg = func() *Cfg {
	cfg, err := NewCfg(provideConfigFilepath())
	if err != nil {
		panic(err)
	}
	return cfg
}()

func provideConfigFilepath() string {
	var configFilepath string
	if configEnv := os.Getenv(ConfigEnvKey); configEnv == "" { // env not found
		configFilepath = path.Join(ConfigDir, ConfigDefaultFile)
		fmt.Printf("config filepath: %s\n", configFilepath)
	} else { // internal.ConfigEnvKey 常量存储的环境变量不为空 将值赋值于config
		configFilepath = configEnv
		fmt.Printf("env: %s, config filepath: %s\n", ConfigEnvKey, configFilepath)
	}
	return configFilepath
}

func NewCfg(configFilePath string) (configs *Cfg, err error) {
	// reset time zone
	time.Local = time.FixedZone("utc", 0)
	configs = &Cfg{}
	err = cfg.ReadFromFile(configFilePath, configs)
	return
}
