package core

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	Env               string
	Host              string
	Port              int
	Key               string
	KeySalt           string
	StaticDir         string
	StaticPath        string
	UploadDir         string
	CorsAllowOrigin   string
	CorsAllowMethods  string
	CorsAllowHeaders  string
	ResponseMode      string
	SessionExpireLong string
	DbDatabase        string
	DbUsername        string
	DbPassword        string
	DbHost            string
	DbPort            string
	DbMode            string
	RedisDatabase     string
	RedisPassword     string
	RedisHost         string
	RedisPort         string
	RedisMode         string
	SystemceHost      string
	AppType           string
	App               string
}

var saveConfig ConfigData

// 获取全部配置
func GetConfig() ConfigData {
	if len(saveConfig.Env) > 0 {
		return saveConfig
	}
	var conf ConfigData

	portStr := os.Getenv("PORT")
	port, _ := strconv.Atoi(portStr)

	conf.Env = os.Getenv("ENV")
	conf.Host = os.Getenv("HOST")
	conf.Port = port
	conf.Key = os.Getenv("KEY")
	conf.KeySalt = os.Getenv("KEY_SALT")
	conf.StaticDir = os.Getenv("STATIC_DIR")
	conf.StaticPath = os.Getenv("STATIC_PATH")
	conf.UploadDir = os.Getenv("UPLOAD_DIR")
	conf.CorsAllowOrigin = os.Getenv("CORS_ALLOW_ORIGIN")
	conf.CorsAllowMethods = os.Getenv("CORS_ALLOW_METHODS")
	conf.CorsAllowHeaders = os.Getenv("CORS_ALLOW_HEADERS")
	conf.ResponseMode = os.Getenv("RESPONSE_MODE")
	conf.SessionExpireLong = os.Getenv("SESSION_EXPIRE_LONG")
	conf.DbDatabase = os.Getenv("DB_DATABASE")
	conf.DbUsername = os.Getenv("DB_USERNAME")
	conf.DbPassword = os.Getenv("DB_PASSWORD")
	conf.DbHost = os.Getenv("DB_HOST")
	conf.DbPort = os.Getenv("DB_PORT")
	conf.DbMode = os.Getenv("DB_MODE")
	conf.RedisDatabase = os.Getenv("REDIS_DATABASE")
	conf.RedisPassword = os.Getenv("REDIS_PASSWORD")
	conf.RedisHost = os.Getenv("REDIS_HOST")
	conf.RedisPort = os.Getenv("REDIS_PORT")
	conf.RedisMode = os.Getenv("REDIS_MODE")
	conf.SystemceHost = os.Getenv("SYSTEMCE_HOST")
	conf.AppType = os.Getenv("APP_TYPE")
	conf.App = os.Getenv("APP")

	return conf
}

// 获取自定义字段
func GetCustomConfig(keys ...string) map[string]string {
	var result = map[string]string{}
	for _, k := range keys {
		result[k] = os.Getenv(k)
	}
	return result
}

// 加载配置
func LoadConfig() {
	var env string
	var host string
	var port string
	var key string
	flag.StringVar(&env, "env", "production", "env usage")
	flag.StringVar(&host, "host", "", "host usage")
	flag.StringVar(&port, "port", "", "port usage")
	flag.StringVar(&key, "key", "", "key usage")
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

	saveConfig = GetConfig()
}
