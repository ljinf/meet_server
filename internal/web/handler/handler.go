package handler

import (
	"github.com/ljinf/meet_server/pkg/jwt"
	"github.com/ljinf/meet_server/pkg/log"
)

type Handler struct {
	jwt    *jwt.JWT
	logger *log.Logger
}

func NewHandler(l *log.Logger, jwt *jwt.JWT) *Handler {
	return &Handler{
		jwt:    jwt,
		logger: l,
	}
}
