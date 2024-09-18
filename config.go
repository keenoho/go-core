package core

import (
	"flag"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

// env file field
var (
	FIELD_ENV      = "ENV"
	FIELD_APP_ID   = "APP_ID"
	FIELD_APP_TYPE = "APP_TYPE" // http or grpc
	FIELD_HOST     = "HOST"
	FIELD_PORT     = "PORT"
	FIELD_MODE     = "MODE" // release or test or debug
)

// default data
var (
	DEFAULT_ENV      = "production"
	DEFAULT_APP_ID   = "keenoho-app-0001"
	DEFAULT_APP_TYPE = "http"
	DEFAULT_HOST     = "0.0.0.0"
	DEFAULT_PORT     = "8080"
	DEFAULT_MODE     = "release"
)

type ConfigOption struct {
	Env    string
	EnvDir string
}

func ConfigGet(key string) string {
	return os.Getenv(key)
}

func ConfigSet(key string, value string) {
	os.Setenv(key, value)
}

func ConfigSetMap(envMap map[string]string) {
	for key, value := range envMap {
		os.Setenv(key, value)
	}
}

func ConfigLoad(options ...ConfigOption) {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = ""
	}
	var envDir string = pwd
	var env string = DEFAULT_ENV
	var appId string = DEFAULT_APP_ID
	var appType string = DEFAULT_APP_TYPE
	var host string = DEFAULT_HOST
	var port string = DEFAULT_PORT
	var mode string = DEFAULT_MODE

	flag.StringVar(&env, "env", DEFAULT_ENV, "env usage")
	flag.StringVar(&appId, "appId", "", "appId usage")
	flag.StringVar(&appType, "appType", "", "appType usage")
	flag.StringVar(&host, "host", "", "host usage")
	flag.StringVar(&port, "port", "", "port usage")
	flag.StringVar(&mode, "mode", "", "mode usage")
	flag.Parse()

	if len(options) > 0 {
		for _, opt := range options {
			if len(opt.Env) > 0 {
				env = opt.Env
			}
			if len(opt.EnvDir) > 0 {
				envDir = opt.EnvDir
			}
		}
	}

	os.Setenv(FIELD_ENV, env)
	if len(envDir) > 0 && !strings.HasSuffix(envDir, "/") {
		envDir = envDir + "/"
	}
	filenames := []string{
		path.Join(envDir, ".env"),
		path.Join(envDir, ".env."+env),
	}
	readEnv, _ := godotenv.Read(filenames...)
	for k, v := range readEnv {
		os.Setenv(k, v)
	}

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
	if len(mode) > 0 {
		os.Setenv(FIELD_MODE, mode)
	}
}
