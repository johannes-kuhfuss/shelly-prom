package config

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/johannes-kuhfuss/services_utils/api_error"
	"github.com/johannes-kuhfuss/services_utils/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Server struct {
		Host                 string `envconfig:"SERVER_HOST"`
		Port                 string `envconfig:"SERVER_PORT" default:"8080"`
		TlsPort              string `envconfig:"SERVER_TLS_PORT" default:"8443"`
		GracefulShutdownTime int    `envconfig:"GRACEFUL_SHUTDOWN_TIME" default:"10"`
		UseTls               bool   `envconfig:"USE_TLS" default:"false"`
		CertFile             string `envconfig:"CERT_FILE" default:"./cert/cert.pem"`
		KeyFile              string `envconfig:"KEY_FILE" default:"./cert/cert.key"`
	}
	Gin struct {
		Mode         string `envconfig:"GIN_MODE" default:"release"`
		TemplatePath string `envconfig:"TEMPLATE_PATH" default:"./templates/"`
	}
	ShellyEM3 struct {
		Host        string `envconfig:"SHELLY_HOST"`
		User        string `envconfig:"SHELLY_USER"`
		Password    string `envconfig:"SHELLY_PASS"`
		IntervalSec int    `envconfig:"SHELLY_INTERVAL_SEC" default:"5"`
	}
	Metrics struct {
	}
	RunTime struct {
		Router        *gin.Engine
		ListenAddr    string
		StartDate     time.Time
		RunShellyPoll bool
	}
}

var (
	EnvFile = ".env"
)

func InitConfig(file string, config *AppConfig) api_error.ApiErr {
	logger.Info(fmt.Sprintf("Initalizing configuration from file %v", file))
	loadConfig(file)
	err := envconfig.Process("", config)
	if err != nil {
		return api_error.NewInternalServerError("Could not initalize configuration. Check your environment variables", err)
	}
	config.RunTime.RunShellyPoll = false
	logger.Info("Done initalizing configuration")
	return nil
}

func loadConfig(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		logger.Info("Could not open env file. Using Environment variable and defaults")
		return err
	}
	return nil
}
