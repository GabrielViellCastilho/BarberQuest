package appointment

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/appointment_domain"
	"go.uber.org/zap"
)

func (ar *appointmentRepository) UpdateAppointmentCompletedById(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateAppointmentCompletedById repository",
		zap.String("journey", "updateAppointmentCompletedById"))

	query := `
	UPDATE appointments
	SET completed = $1
	WHERE id = $2 AND barber_id = $3;
	`

	result, err := ar.databaseConection.Exec(ctx, query, appointmentDomain.GetCompleted(), appointmentDomain.GetId(), appointmentDomain.GetBarberId())
	if err != nil {
		logger.Error("Error updating appointment_domain", err, zap.String("journey", "updateAppointmentCompletedById"))
		return rest_err.NewInternalServerError("Database error when updating appointment_domain")
	}

	noRow := result.RowsAffected()
	if noRow == 0 {
		logger.WarnWithoutError("No appointment_domain found with given ID", zap.Int("id", appointmentDomain.GetId()), zap.String("journey", "updateAppointmentCompletedById"))
		return rest_err.NewNotFoundError("Appointment not found")
	}

	logger.Info("Successful updateAppointmentCompletedById repository",
		zap.Int64("rowsAffected", noRow),
		zap.String("journey", "updateAppointmentCompletedById"))

	return nil
}
