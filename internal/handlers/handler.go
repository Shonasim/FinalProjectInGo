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

	h.mux.Handle(http.MethodPost, "/sign-in", h.SignIn)

	cities := h.mux.Group("/cities")
	{
		cities.Handle(http.MethodGet, "", h.GetCities)
	}

	v1 := h.mux.Group("/v1")
	v1.Use(
		middleware.Authenticate(),
	)

	car := v1.Group("/cars")
	{
		car.Handle(http.MethodPost, "/add", h.CreateCar)
		car.Handle(http.MethodGet, "/get", h.GetCars)
	}

	seat := car.Group("/seats")
	{
		seat.Handle(http.MethodPost, "/add", h.AddSeats)
	}
	routes := v1.Group("/routes")
	{
		routes.Handle(http.MethodPost, "/add", h.CreateRoute)
		routes.Handle(http.MethodGet, "/get", h.GetRoutes)
		routes.Handle(http.MethodGet, "/get/{route_id}", h.GetRouteById)
	}
}
