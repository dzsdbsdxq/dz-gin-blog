package controller

import (
	"github.com/dzsdbsdxq/dz-gin-blog/app/core/install/service"
)

type InstallHandler struct {
	serv service.IInstallService
}

func NewInstallHandler(serv service.IInstallService) *InstallHandler {
	return &InstallHandler{
		serv: serv,
	}
}

func (h *InstallHandler) SystemAutoReRegister() error {
	err := h.serv.RegisterTables()
	if err != nil {
		return err
	}
	return nil
}
