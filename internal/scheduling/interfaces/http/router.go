package http

import (
	"helthmed-scheduling/internal/scheduling/infrastructure/repository"
	"helthmed-scheduling/internal/scheduling/usecase"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	appointmentRepo := repository.NewAppointmentRepository()
	appointmentUsecase := usecase.NewAppointmentUseCase(appointmentRepo)
	appointmentHandler := NewAppointmentHandler(*appointmentUsecase)

	api := router.Group("/api")
	{
		api.POST("/appointments", appointmentHandler.Create)
		api.GET("/appointments/:id", appointmentHandler.Get)
		api.PUT("/appointments", appointmentHandler.Update)
		api.DELETE("/appointments/:id", appointmentHandler.Delete)
	}
}
