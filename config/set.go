package config

import (
	"flag"
	"strings"

	"github.com/spf13/viper"
)

const FileName = "config"
const TestConfigFileName = "config_test"

const FilePath = "./config"
const TestConfigFilePath = "../../config/"

func SetCfg() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName(FileName)
	if !isTestEnv() {
		v.AddConfigPath(FilePath)
	} else {
		v.AddConfigPath(TestConfigFilePath)
	}

	err := v.ReadInConfig()
	if err != nil {
		panic("can not read config file" + err.Error())
	}

	v.AutomaticEnv()
	v.SetEnvPrefix("project")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//rewrite configs for test env
	if isTestEnv() {
		v.SetConfigName(TestConfigFileName)
		err = v.MergeInConfig()
		if err != nil {
			panic("can not merge config file" + err.Error())
		}
	}

	_ = v.AllSettings()
	return v
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}
