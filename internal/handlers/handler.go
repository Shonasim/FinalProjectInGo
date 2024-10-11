package handlers

import (
	"FinalProject/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	mux     *gin.Engine
	service *service.Service
	logger  *logrus.Logger
}

func NewHandler(mux *gin.Engine, s *service.Service, log *logrus.Logger) *Handler {
	return &Handler{
		mux:     mux,
		service: s,
		logger:  log,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) InitRoutes() {
	h.mux.Handle(http.MethodPost, "/register", h.Register)
}
