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
	RedisDatabase     string
	RedisPassword     string
	RedisHost         string
	RedisPort         string
	SystemceHost      string
}

var saveConfig ConfigData

// 获取全部配置
func GetConfig() ConfigData {
	if len(saveConfig.Env) > 0 {
		return saveConfig
	}
	var conf ConfigData

	portStr := os.Getenv("SERVER_PORT")
	port, _ := strconv.Atoi(portStr)

	conf.Env = os.Getenv("ENV")
	conf.Host = os.Getenv("SERVER_HOST")
	conf.Port = port
	conf.Key = os.Getenv("KEY")
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
	conf.RedisDatabase = os.Getenv("REDIS_DATABASE")
	conf.RedisPassword = os.Getenv("REDIS_PASSWORD")
	conf.RedisHost = os.Getenv("REDIS_HOST")
	conf.RedisPort = os.Getenv("REDIS_PORT")
	conf.SystemceHost = os.Getenv("SYSTEMCE_HOST")

	return conf
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
		os.Setenv("SERVER_HOST", host)
	}
	if len(port) > 0 {
		os.Setenv("SERVER_PORT", port)
	}
	if len(key) > 0 {
		os.Setenv("KEY", key)
	}

	saveConfig = GetConfig()
}
