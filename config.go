package core

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

// env file field
var (
	FIELD_ENV      = "ENV"
	FIELD_APP_ID   = "APP_ID"
	FIELD_APP_TYPE = "APP_TYPE" // http or grpc
	FIELD_HOST     = "HOST"
	FIELD_PORT     = "PORT"
)

// default data
var (
	DEFAULT_ENV      = "production"
	DEFAULT_APP_ID   = "keenoho-app-0001"
	DEFAULT_APP_TYPE = "http"
	DEFAULT_HOST     = "0.0.0.0"
	DEFAULT_PORT     = "8080"
)

func ConfigGet(key string) string {
	return os.Getenv(key)
}

func ConfigSet(configKey string, envKey string, value string) {
	os.Setenv(envKey, value)
}

func ConfigLoad(targetEnv ...string) {
	var env string = DEFAULT_ENV
	var appId string = DEFAULT_APP_ID
	var appType string = DEFAULT_APP_TYPE
	var host string = DEFAULT_HOST
	var port string = DEFAULT_PORT

	flag.StringVar(&env, "env", DEFAULT_ENV, "env usage")
	flag.StringVar(&appId, "appId", DEFAULT_APP_ID, "appId usage")
	flag.StringVar(&appType, "appType", DEFAULT_APP_TYPE, "appType usage")
	flag.StringVar(&host, "host", DEFAULT_HOST, "host usage")
	flag.StringVar(&port, "port", DEFAULT_PORT, "port usage")
	flag.Parse()

	if len(targetEnv) > 0 {
		env = targetEnv[len(targetEnv)-1]
	}

	os.Setenv(FIELD_ENV, env)
	os.Setenv(FIELD_APP_ID, appId)
	os.Setenv(FIELD_APP_TYPE, appType)
	os.Setenv(FIELD_HOST, host)
	os.Setenv(FIELD_PORT, port)

	envFileName := ".env." + env
	readEnv, _ := godotenv.Read(".env", envFileName)
	for k, v := range readEnv {
		os.Setenv(k, v)
	}
}
