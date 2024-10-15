package handlers

import (
	"FinalProject/internal/middleware"
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
	h.mux.Use(
		middleware.CORS(),
		middleware.Recovery(),
	)

	registration := h.mux.Group("/register")
	{
		registration.Handle(http.MethodGet, "/check-user", h.GetUserByEmail)
		registration.Handle(http.MethodPost, "", h.Register)
		registration.Handle(http.MethodPost, "/personal-information", h.AddPersonalInfo)
	}

	h.mux.Handle(http.MethodGet, "sign-in")

	h.mux.Handle(http.MethodGet, "/getusers", h.GetUsers)
	h.mux.Handle(http.MethodGet, "/getusers/{id}", h.GetUserByID)
	h.mux.Handle(http.MethodGet, "/delete", h.DeleteUser)

}
