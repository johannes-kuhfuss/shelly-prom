package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johannes-kuhfuss/shelly-prom/config"
	"github.com/johannes-kuhfuss/shelly-prom/dto"
)

type StatsUiHandler struct {
	Cfg *config.AppConfig
}

func NewStatsUiHandler(cfg *config.AppConfig) StatsUiHandler {
	return StatsUiHandler{
		Cfg: cfg,
	}
}

func (uh *StatsUiHandler) StatusPage(c *gin.Context) {
	configData := dto.GetConfig(uh.Cfg)
	c.HTML(http.StatusOK, "status.page.tmpl", gin.H{
		"configdata": configData,
	})
}

func (uh *StatsUiHandler) AboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.page.tmpl", nil)
}
