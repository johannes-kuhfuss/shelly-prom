package dto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johannes-kuhfuss/shelly-prom/config"
)

type ConfigResp struct {
	ServerHost                 string
	ServerPort                 string
	ServerTlsPort              string
	ServerGracefulShutdownTime string
	ServerUseTls               string
	ServerCertFile             string
	ServerKeyFile              string
	GinMode                    string
	StartDate                  string
	AVoltage                   string
	ACurrent                   string
	AActPower                  string
	AAprtPower                 string
	APf                        string
	AFreq                      string
	AErrors                    string
	BVoltage                   string
	BCurrent                   string
	BActPower                  string
	BAprtPower                 string
	BPf                        string
	BFreq                      string
	BErrors                    string
	CVoltage                   string
	CCurrent                   string
	CActPower                  string
	CAprtPower                 string
	CPf                        string
	CFreq                      string
	CErrors                    string
	NCurrent                   string
	NErrors                    string
	TotalCurrent               string
	TotalActPower              string
	TotalAprtPower             string
	UserCalibratedPhase        string
	Errors                     string
}

func GetConfig(cfg *config.AppConfig) ConfigResp {
	resp := ConfigResp{
		ServerHost:                 cfg.Server.Host,
		ServerPort:                 cfg.Server.Port,
		ServerTlsPort:              cfg.Server.TlsPort,
		ServerGracefulShutdownTime: strconv.Itoa(cfg.Server.GracefulShutdownTime),
		ServerUseTls:               strconv.FormatBool(cfg.Server.UseTls),
		ServerCertFile:             cfg.Server.CertFile,
		ServerKeyFile:              cfg.Server.KeyFile,
		GinMode:                    cfg.Gin.Mode,
		StartDate:                  cfg.RunTime.StartDate.Local().Format("2006-01-02 15:04:05 -0700"),
		AVoltage:                   fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.AVoltage),
		ACurrent:                   fmt.Sprintf("%.3f", cfg.RunTime.ShellyCurrentData.ACurrent),
		AActPower:                  fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.AActPower),
		AAprtPower:                 fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.AAprtPower),
		APf:                        fmt.Sprintf("%.2f", cfg.RunTime.ShellyCurrentData.APf),
		AFreq:                      fmt.Sprintf("%.0f", cfg.RunTime.ShellyCurrentData.AFreq),
		AErrors:                    formatErrors(cfg.RunTime.ShellyCurrentData.AErrors),
		BVoltage:                   fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.BVoltage),
		BCurrent:                   fmt.Sprintf("%.3f", cfg.RunTime.ShellyCurrentData.BCurrent),
		BActPower:                  fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.BActPower),
		BAprtPower:                 fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.BAprtPower),
		BPf:                        fmt.Sprintf("%.2f", cfg.RunTime.ShellyCurrentData.BPf),
		BFreq:                      fmt.Sprintf("%.0f", cfg.RunTime.ShellyCurrentData.BFreq),
		BErrors:                    formatErrors(cfg.RunTime.ShellyCurrentData.BErrors),
		CVoltage:                   fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.CVoltage),
		CCurrent:                   fmt.Sprintf("%.3f", cfg.RunTime.ShellyCurrentData.CCurrent),
		CActPower:                  fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.CActPower),
		CAprtPower:                 fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.CAprtPower),
		CPf:                        fmt.Sprintf("%.2f", cfg.RunTime.ShellyCurrentData.CPf),
		CFreq:                      fmt.Sprintf("%.0f", cfg.RunTime.ShellyCurrentData.CFreq),
		CErrors:                    formatErrors(cfg.RunTime.ShellyCurrentData.CErrors),
		TotalCurrent:               fmt.Sprintf("%.3f", cfg.RunTime.ShellyCurrentData.TotalCurrent),
		TotalActPower:              fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.TotalActPower),
		TotalAprtPower:             fmt.Sprintf("%.1f", cfg.RunTime.ShellyCurrentData.TotalAprtPower),
		UserCalibratedPhase:        formatErrors(cfg.RunTime.ShellyCurrentData.UserCalibratedPhase),
		Errors:                     formatErrors(cfg.RunTime.ShellyCurrentData.Errors),
	}
	if cfg.Server.Host == "" {
		resp.ServerHost = "localhost"
	}
	return resp
}

func formatErrors(err []string) string {
	str := strings.Join(err, " ")
	if str == "" {
		str = "None"
	}
	return str
}
