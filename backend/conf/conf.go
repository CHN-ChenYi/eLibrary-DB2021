package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	logrus.Info("Configuration file loaded")

	var confItems = map[string][]string{
		"sql": {"user", "password", "protocol", "host", "port", "db_name"},
		"app": {"port", "cors"},
	}

	for k, v := range confItems {
		checkConfIsSet(k, v)
	}

	logrus.Info("All required values in configuration file are set")
}

func checkConfIsSet(name string, keys []string) {
	for i := range keys {
		wholeKey := name + "." + keys[i]
		if !viper.IsSet(wholeKey) {
			logrus.WithField(wholeKey, nil).
				Fatal("The following item of your configuration file hasn't been set properly: ")
		}
	}
}
