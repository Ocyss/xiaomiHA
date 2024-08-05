package utils

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	Character string
	User      struct {
		Base  string
		City  string
		Other map[string]string
	}
	Redis asynq.RedisClientOpt
	Gpt   struct {
		ApiKey  string
		ApiBase string
		Tools   struct {
			Weather struct {
				Key      string
				Location string
			}
		}
	}
	FeiShu struct {
		OpenId            string
		AppId             string
		AppSecret         string
		EncryptKey        string
		VerificationToken string

		CalendarId string
	}
}

var (
	C config
)

func init() {
	formatter := logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		DisableQuote:              true,
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&formatter)

	viper.SetDefault("ContentDir", "content")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/xiaomiHA")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		//var configFileNotFoundError viper.ConfigFileNotFoundError
		//if !errors.As(err, &configFileNotFoundError) {
		panic(fmt.Errorf("fatal error config file: %w", err))
		//}
	}
	err := viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}
}
