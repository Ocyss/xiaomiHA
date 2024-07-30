package utils

import "github.com/sirupsen/logrus"

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
}
