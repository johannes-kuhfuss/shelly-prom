package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/johannes-kuhfuss/services_utils/logger"
	"github.com/johannes-kuhfuss/shelly-prom/config"
	"github.com/johannes-kuhfuss/shelly-prom/domain"
)

type ShellyPollService interface {
	Poll()
}

type DefaultShellyPollService struct {
	Cfg *config.AppConfig
}

var (
	httpShellyPollTr     http.Transport
	httpShellyPollClient http.Client
)

func NewShellyPollService(cfg *config.AppConfig) DefaultShellyPollService {
	InitShellyPollHttp()
	return DefaultShellyPollService{
		Cfg: cfg,
	}
}

func InitShellyPollHttp() {
	httpShellyPollTr = http.Transport{
		DisableKeepAlives:  false,
		DisableCompression: false,
		MaxIdleConns:       0,
		IdleConnTimeout:    0,
	}
	httpShellyPollClient = http.Client{
		Transport: &httpShellyPollTr,
		Timeout:   5 * time.Second,
	}
}

func (s DefaultShellyPollService) Poll() {
	if s.Cfg.ShellyEM3.Host == "" {
		logger.Warn("No Shelly host given. Not polling Shelly")
		s.Cfg.RunTime.RunShellyPoll = false
	} else {
		logger.Info(fmt.Sprintf("Starting to poll Shelly from host %v", s.Cfg.ShellyEM3.Host))
		s.Cfg.RunTime.RunShellyPoll = true
	}

	for s.Cfg.RunTime.RunShellyPoll {
		ShellyPollRun(s)
		time.Sleep(time.Duration(s.Cfg.ShellyEM3.IntervalSec) * time.Second)
	}
}

func ShellyPollRun(s DefaultShellyPollService) {
	pollUrl := url.URL{
		Scheme:   "http",
		Host:     s.Cfg.ShellyEM3.Host,
		Path:     "/rpc/EM.GetStatus",
		RawQuery: "id=0",
	}
	if s.Cfg.ShellyEM3.UseBasicAuth {
		pollUrl.User = url.UserPassword(s.Cfg.ShellyEM3.User, s.Cfg.ShellyEM3.Password)
	}
	logger.Info(fmt.Sprintf("Url: %v", pollUrl.String()))
	shellyState, err := GetJsonFromPollUrl(pollUrl.String())
	if err == nil {
		// update metrics
		_ = shellyState
	} else {
		logger.Error("Error while retrieving data from Shelly: ", err)
	}

}

func GetJsonFromPollUrl(pollUrl string) (*domain.ShellyData, error) {
	var shellyData domain.ShellyData

	req, _ := http.NewRequest("GET", pollUrl, nil)

	resp, err := httpShellyPollClient.Do(req)
	if err != nil {
		logger.Error("Error while polling Shelly data", err)
		return nil, err
	}
	if resp.StatusCode == 404 {
		err := errors.New("URL not found")
		logger.Error("Error while polling Shelly data", err)
		return nil, err
	}
	if resp.StatusCode == 401 {
		err := errors.New("could not authenticate")
		logger.Error("Error while polling Shelly data", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Error("Error while reading response body with Shelly data", err)
		return nil, err
	}

	err = json.Unmarshal(body, &shellyData)
	if err != nil {
		logger.Error("Error while coverting response body to JSON", err)
		return nil, err
	}

	return &shellyData, nil
}
