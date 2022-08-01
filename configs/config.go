package configs

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
)

var cfg Configuration

type Configuration struct {
	App   AppConfig `toml:"app"`
	Mysql DBConfig  `toml:"mysql"`
	Log   LogConfig `toml:"log"`
}

type DBConfig struct {
	//配置MySQL连接参数
	DSN             string `toml:"dsn"`
	MaxOpenConns    int    `toml:"max_open_conns"`
	MaxIdleConns    int    `toml:"max_idle_conns"`
	ConnMaxLifetime int    `toml:"conn_max_lifetime"`
}

type AppConfig struct {
	Address string `toml:"address"`
}

type LogConfig struct {
	Level  string `toml:"level"`
	Format string `toml:"format"`
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &cfg)
	return err
}

func GetConfig() Configuration {
	return cfg
}

func InitApp() error {
	configPath := flag.String("c", "configs/config.toml", "config file path")
	if err := InitConfig(*configPath); err != nil {
		return fmt.Errorf("加载配置文件出错, %v", err)
	}
	return nil
}
