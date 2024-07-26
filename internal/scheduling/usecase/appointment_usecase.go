package usecase

import (
	"helthmed-scheduling/internal/scheduling/domain"
	"helthmed-scheduling/internal/scheduling/interfaces/messaging"
)

type AppointmentUseCase struct {
	repo domain.AppointmentRepository
}

func NewAppointmentUseCase(repo domain.AppointmentRepository) *AppointmentUseCase {
	return &AppointmentUseCase{repo: repo}
}

func (u *AppointmentUseCase) Create(appointment *domain.Appointment) error {
	return u.repo.Create(appointment)

	// Publish the event
	messaging.PublishAppointmentCreated(appointment)
	return nil
}

func (u *AppointmentUseCase) Get(id uint) (*domain.Appointment, error) {
	return u.repo.FindByID(id)
}

func (u *AppointmentUseCase) Update(appointment *domain.Appointment) error {
	return u.repo.Update(appointment)
}

func (u *AppointmentUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}
