package ports

import (
	"AzubiTool/internal/domain"
	"time"
)

type BcsService interface {
	GetOid(username, password string) (string, error)
	GetEfforts(oid, username, password string, from, to time.Time) ([]domain.EffortRecord, error)
	GetAppointments(oid, username, password string, from, to time.Time) ([]domain.Appointment, error)
}
