package config

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug     bool
	Port      int
	Secret    string
	JwtSecret string
	BaseUrl   string
	FrontUrl  string

	Search   SearchConfig
	Mailer   MailerConfig
	Postgres PostgresConfig
	Mongo    MongoConfig
	Redis    RedisConfig
	Bucket   BucketConfig
	FileSys  FileSysConfig
}

type PostgresConfig struct {
	Host   string
	Port   int
	DbUser string
	DbPass string
	DbName string
}

type MongoConfig struct {
	Host   string
	DBName string
}

type RedisConfig struct {
	Host     string
	Username string
	Password string
}

type SearchConfig struct {
	Host   string
	ApiKey string
}

type MailerConfig struct {
	ApiKey   string
	MailFrom string
}

type BucketConfig struct {
	GCPBucketName  string
	GCPUploadPath  string
	GCPAuthURL     string
	GCPClientEmail string
	GCPBaseFolder  string
	GCPPrivateKey  string
}

type FileSysConfig struct {
	RelativePath string
}

var config *Config

func Get() *Config {
	if config == nil {
		config = loadConfig()
	}
	return config
}

func loadConfig() *Config {
	// Cargar el archivo .env
	_ = godotenv.Load()

	debug := os.Getenv("DEBUG") == "true"

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000 // valor por defecto
	}

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		postgresPort = 5432
	}

	// Get the base64 encoded private key and decode it to UTF-8
	gcpPrivateKeyBase64 := os.Getenv("GCP_PRIVATE_KEY_BASE64")
	gcpPrivateKeyDecoded := ""
	if gcpPrivateKeyBase64 != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(gcpPrivateKeyBase64)
		if err == nil {
			gcpPrivateKeyDecoded = string(decodedBytes)
		}
	}

	return &Config{
		Debug:     debug,
		Port:      port,
		Secret:    os.Getenv("SECRET"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		BaseUrl:   os.Getenv("BASE_URL"),
		FrontUrl:  os.Getenv("FRONT_URL"),

		Search: SearchConfig{
			Host:   os.Getenv("SEARCH_HOST"),
			ApiKey: os.Getenv("SEARCH_API_KEY"),
		},
		Mailer: MailerConfig{
			ApiKey:   os.Getenv("MAILER_API_KEY"),
			MailFrom: os.Getenv("MAILER_FROM"),
		},
		Postgres: PostgresConfig{
			Host:   os.Getenv("POSTGRES_HOST"),
			Port:   postgresPort,
			DbUser: os.Getenv("POSTGRES_USER"),
			DbPass: os.Getenv("POSTGRES_PASS"),
			DbName: os.Getenv("POSTGRES_DB"),
		},
		Mongo: MongoConfig{
			Host:   os.Getenv("MONGO_HOST"),
			DBName: os.Getenv("MONGO_DB"),
		},
		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Bucket: BucketConfig{
			GCPBucketName:  os.Getenv("GCP_BUCKET_NAME"),
			GCPUploadPath:  os.Getenv("GCP_UPLOAD_PATH"),
			GCPAuthURL:     os.Getenv("GCP_AUTH_URL"),
			GCPClientEmail: os.Getenv("GCP_CLIENT_EMAIL"),
			GCPPrivateKey:  gcpPrivateKeyDecoded,
			GCPBaseFolder:  os.Getenv("GCP_BASE_FOLDER"),
		},
	}
}
