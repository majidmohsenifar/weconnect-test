package config

import (
	"flag"

	"github.com/spf13/viper"
)

const (
	EnvConfigKey = "project.environment"
	EnvProd      = "prod"
	EnvTest      = "test"
	EnvDev       = "dev"
)

type Cfg struct {
	viper *viper.Viper
}

func (c *Cfg) GetString(name string) string {
	return c.viper.GetString(name)

}

func (c *Cfg) GetInt(name string) int {
	return c.viper.GetInt(name)
}

func (c *Cfg) GetBool(name string) bool {
	return c.viper.GetBool(name)
}

func (c *Cfg) GetEnv() string {
	return c.GetString(EnvConfigKey)

}

func (c *Cfg) Set(key string, value interface{}) {
	c.viper.Set(key, value)

}

func NewConfigs(viper *viper.Viper) *Cfg {
	c := &Cfg{viper}
	if flag.Lookup("test.v") != nil {
		c.Set(EnvConfigKey, EnvTest)
	}
	return c
}
