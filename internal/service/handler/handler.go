package handler

import (
	"github.com/ljinf/meet_server/pkg/helper/sid"
	"github.com/ljinf/meet_server/pkg/log"
)

type Handler struct {
	sid    *sid.Sid
	logger *log.Logger
}

func NewHandler(sid *sid.Sid, logger *log.Logger) *Handler {
	return &Handler{
		sid:    sid,
		logger: logger,
	}
}
