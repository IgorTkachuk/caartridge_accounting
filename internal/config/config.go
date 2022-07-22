package config

import (
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	JWT struct {
		Secret string `yaml:"Secret" env:"JWTSECRET" env-required:"true"`
	} `yaml:"jwt"`
	Listen struct {
		BindIP string `yaml:"bind_ip" env:"BIND" env-default :"127.0.0.1" env-required:"true"`
		Port   string `yaml:"port" env:"PORT" env-default:"3001" env-required:"true"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env:"DBHOST" env-default:"127.0.0.1" env-required:"true"`
	Port     string `yaml:"port" env:"DBPORT" env-default:"5432" env-required:"true"`
	Database string `yaml:"database" env:"DBNAME" env-default:"ctr env-required:"true""`
	Username string `yaml:"username" env:"DBUSER" env-default:"postgres" env-required:"true"`
	Password string `yaml:"password" env:"DBPASS" env-default:"" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
