package handlers

import (
	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type HandlerImpl struct {
	Config				models.HandlerConfig
}

func NewHandlerImpl(config models.HandlerConfig) Handler {
	return &HandlerImpl{
		Config: config,
	}
}

func (h *HandlerImpl) ValidHost(host string) bool {

	for _, v := range h.Config.HostsRegs {
		match := v.FindString(host)
		if match != "" {
			return true
		}
	}

	return false
}
