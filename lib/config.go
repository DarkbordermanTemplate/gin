package lib

import (
	"os"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

/*
Service configuration
*/
type Config struct {
	Port    string
	Address string
}

/*
Get environment key as integer.

Default value(defval) will be applied
if the environment key's value is missing or can't convert to desired type.
*/
func getEnvInt(key string, defval int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		logger.Warnf(
			"%s cannot be converted into integer, using default value",
			key,
		)
		return defval
	}
	return value
}

/*
Get environment key as string.

Default value(defval) will be applied
if the environment key's value is missing or can't convert to desired type.
*/
func getEnvStr(key string, defval string) string {
	value := os.Getenv(key)
	if value == "" {
		return defval
	}
	return value
}

var CONFIG = Config{
	Port:    getEnvStr("PORT", "8080"),
	Address: getEnvStr("ADDRESS", "0.0.0.0"),
}
