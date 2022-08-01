package configs

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"rt-msg-carrier/pkg/env"
	"rt-msg-carrier/pkg/file"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	MySQL struct {
		Read struct {
			Addr   string `toml:"addr"`
			User   string `toml:"user"`
			Pass   string `toml:"pass"`
			Name   string `toml:"name"`
			DbName string `toml:"dbname"`
		} `toml:"read"`
		Write struct {
			Addr   string `toml:"addr"`
			User   string `toml:"user"`
			Pass   string `toml:"pass"`
			Name   string `toml:"name"`
			DbName string `toml:"dbname"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
			Timeout         int           `toml:"timeout"`
		} `toml:"base"`
	} `toml:"mysql"`

	Redis struct {
		Addr         string `toml:"addr"`
		Pass         string `toml:"pass"`
		Db           int    `toml:"db"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	} `toml:"redis"`

	LogConfig struct {
		Level  string `toml:"level"`
		Format string `toml:"format"`
	}

	HashIds struct {
		Secret string `toml:"secret"`
		Length int    `toml:"length"`
	} `toml:"hashids"`

	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`

	Server struct {
		Addr         string `toml:"addr"`
		ReadTimeout  int    `toml:"read_timeout"`
		WriteTimeout int    `toml:"write_timeout"`
		DebugMode    bool   `toml:debug_mode`
	} `toml:"server"`
}

var (
	devConfigs []byte

	fatConfigs []byte

	uatConfigs []byte

	proConfigs []byte
)

func init() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfigs)
	case "fat":
		r = bytes.NewReader(fatConfigs)
	case "uat":
		r = bytes.NewReader(uatConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(fatConfigs)
	}

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.AddConfigPath("./configs")

	configFile := "./configs/" + env.Active().Value() + "_configs.toml"
	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}

func (c *Config) DebugMode() bool {
	return c.Server.DebugMode
}

func (c *Config) GinMode() string {
	if c.DebugMode() {
		return gin.DebugMode
	} else {
		return gin.ReleaseMode
	}
}
