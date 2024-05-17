package monitorHandlers

import (
	_pkgConfig "github.com/MarkTBSS/053_Health_Check_Service/config"
	_pkgModulesMonitor "github.com/MarkTBSS/053_Health_Check_Service/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg _pkgConfig.IConfig
}

func MonitorHandler(cfg _pkgConfig.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &_pkgModulesMonitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
