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
	var env string
	var appId string
	var appType string
	var host string
	var port string

	if len(targetEnv) > 0 {
		env = targetEnv[len(targetEnv)-1]
	} else {
		flag.StringVar(&env, "env", DEFAULT_ENV, "env usage")
	}

	os.Setenv(FIELD_ENV, env)
	os.Setenv(FIELD_APP_ID, DEFAULT_APP_ID)
	os.Setenv(FIELD_APP_TYPE, DEFAULT_APP_TYPE)
	os.Setenv(FIELD_HOST, DEFAULT_HOST)
	os.Setenv(FIELD_PORT, DEFAULT_PORT)

	godotenv.Load(".env")
	envFileName := ".env." + env
	readEnv, _ := godotenv.Read(".env", envFileName)
	for k, v := range readEnv {
		os.Setenv(k, v)
	}

	flag.StringVar(&appId, "appId", "", "appId usage")
	flag.StringVar(&appType, "appType", "", "appType usage")
	flag.StringVar(&host, "host", "", "host usage")
	flag.StringVar(&port, "port", "", "port usage")
	flag.Parse()

	if len(appId) > 0 {
		os.Setenv(FIELD_APP_ID, appId)
	}
	if len(appType) > 0 {
		os.Setenv(FIELD_APP_TYPE, appType)
	}
	if len(host) > 0 {
		os.Setenv(FIELD_HOST, host)
	}
	if len(port) > 0 {
		os.Setenv(FIELD_PORT, port)
	}
}
