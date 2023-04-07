package common

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var configKeyMap map[string]string = map[string]string{
	"App":              "APP",
	"AppType":          "APP_TYPE",
	"Key":              "KEY",
	"Env":              "ENV",
	"Host":             "HOST",
	"Port":             "PORT",
	"RegisterHost":     "REGISTER_HOST",
	"RegisterPort":     "REGISTER_PORT",
	"StaticDir":        "STATIC_DIR",
	"StaticPath":       "STATIC_PATH",
	"UploadDir":        "UPLOAD_DIR",
	"TrustedProxies":   "TRUSTED_PROXIES",
	"CorsAllowOrigin":  "CORS_ALLOW_ORIGIN",
	"CorsAllowMethods": "CORS_ALLOW_METHODS",
	"CorsAllowHeaders": "CORS_ALLOW_HEADERS",
	"DbDatabase":       "DB_DATABASE",
	"DbUsername":       "DB_USERNAME",
	"DbPassword":       "DB_PASSWORD",
	"DbHost":           "DB_HOST",
	"DbPort":           "DB_PORT",
	"DbMode":           "DB_MODE",
	"RedisDatabase":    "REDIS_DATABASE",
	"RedisPassword":    "REDIS_PASSWORD",
	"RedisHost":        "REDIS_HOST",
	"RedisPort":        "REDIS_PORT",
	"RedisMode":        "REDIS_MODE",
	"InternalIp":       "INTERNAL_IP",
	"PublicIp":         "PUBLIC_IP",
}

var cacheConfig map[string]string = nil

func AddConfigKey(configKey string, envName string) {
	configKeyMap[configKey] = envName
}

func AddConfigKeys(configMap map[string]string) {
	for k, v := range configMap {
		configKeyMap[k] = v
	}
}

func GetConfig(keys ...string) map[string]string {
	if len(keys) > 0 {
		conf := map[string]string{}
		for _, k := range keys {
			envKey, isExist := configKeyMap[k]
			if isExist {
				conf[k] = os.Getenv(envKey)
			}
		}
		return conf
	} else if cacheConfig != nil {
		return cacheConfig
	} else {
		cacheConfig = map[string]string{}
		for k, ek := range configKeyMap {
			cacheConfig[k] = os.Getenv(ek)
		}
		return cacheConfig
	}
}

func SetConfig(configKey string, envKey string, value string) {
	AddConfigKey(configKey, envKey)
	os.Setenv(envKey, value)
}

func LoadConfig() {
	var env string
	var host string
	var port string
	var key string
	var app string
	var appType string
	var registerHost string
	var registerPort string
	flag.StringVar(&env, "env", "production", "env usage")
	flag.StringVar(&host, "host", "", "host usage")
	flag.StringVar(&port, "port", "", "port usage")
	flag.StringVar(&key, "key", "", "key usage")
	flag.StringVar(&app, "app", "", "app usage")
	flag.StringVar(&appType, "appType", "", "app_type usage")
	flag.StringVar(&registerHost, "registerHost", "", "registerHost usage")
	flag.StringVar(&registerPort, "registerPort", "", "registerPort usage")
	flag.Parse()

	envFileName := ".env." + env

	readEnv, _ := godotenv.Read(".env", envFileName)
	for k, v := range readEnv {
		os.Setenv(k, v)
	}
	os.Setenv("ENV", env)

	if len(host) > 0 {
		os.Setenv("HOST", host)
	}
	if len(port) > 0 {
		os.Setenv("PORT", port)
	}
	if len(key) > 0 {
		os.Setenv("KEY", key)
	}
	if len(app) > 0 {
		os.Setenv("APP", app)
	}
	if len(appType) > 0 {
		os.Setenv("APP_TYPE", appType)
	}
	if len(registerHost) > 0 {
		os.Setenv("REGISTER_HOST", registerHost)
	}
	if len(registerPort) > 0 {
		os.Setenv("REGISTER_PORT", registerPort)
	}
}
