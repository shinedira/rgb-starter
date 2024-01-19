package config

import (
	"log"
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Env struct {
	*viper.Viper
}

func BootEnv() *Env {
	v := viper.New()

	v.SetConfigName("env")
	v.SetConfigFile(".env")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file")
		os.Exit(0)
	}

	v.SetEnvPrefix("app")
	v.AutomaticEnv()

	return &Env{v}
}

func (env *Env) Env(envName string, defaultValue ...any) any {
	value := env.Get(envName, defaultValue...)
	if cast.ToString(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}

	return value
}

func (env *Env) Get(path string, defaultValue ...any) any {
	if !env.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return env.Viper.Get(path)
}
