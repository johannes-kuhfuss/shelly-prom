package service

import (
	"errors"
	"fmt"
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
	shellyState, err := GetJsonFromPollUrl(pollUrl.String(), s.Cfg.ShellyEM3.User, s.Cfg.ShellyEM3.Password)
	if err == nil {
		// update metrics
		_ = shellyState
	} else {
		logger.Error("Error while retrieving data from Shelly: ", err)
	}

}

func GetJsonFromPollUrl(pollUrl string, user string, password string) (*domain.ShellyData, error) {
	var shellyData domain.ShellyData

	req, _ := http.NewRequest("GET", pollUrl, nil)
	req.SetBasicAuth(user, password)
	resp, err := httpShellyPollClient.Do(req)
	if err != nil {
		logger.Error("Error while polling Shelly data", err)
		return nil, err
	}
	if resp.StatusCode == 404 {
		err := errors.New("URl not found")
		logger.Error("Error while polling Shelly data", err)
		return nil, err
	}
	defer resp.Body.Close()

	// convert to JSON

	return &shellyData, nil
}
