package config

type Config struct {
	Env *Env
}

func BootConfig() *Config {
	return &Config{
		Env: BootEnv(),
	}
}
