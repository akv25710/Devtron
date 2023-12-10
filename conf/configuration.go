package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"lfetchogger/services/logger_service"
	"os"
)

type Configuration struct {
	Server ServerConfiguration
	Logger LoggerConfiguration
}

type ServerConfiguration struct {
	Port int
}

type LoggerConfiguration struct {
	Cloud  logger_service.CloudProvider
	Bucket string
}

const (
	defaultConfigFile = "assets/config/config.yaml"
)

func ReadConfig() Configuration {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logrus.Warnln("Env [CONFIG_PATH] not defined. Using default config location")
		configPath = defaultConfigFile
	}

	logrus.Info("Configs at: %s", configPath)
	var config Configuration
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		logrus.Fatalln("Unable to read config.. ", err)
	}

	logrus.Info("Configs read successfully...")

	return config
}
