package appointment

import (
	"context"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/appointment_domain"
	"go.uber.org/zap"
)

func (ar *appointmentRepository) CreateAppointment(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) (appointment_domain.AppointmentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createAppointment repository",
		zap.String("journey", "createAppointment"))
	query := `
	        INSERT INTO appointments (client_name,client_contact,appointment_date,barber_id,service_id,user_id)
        VALUES ($1, $2, $3,$4,$5,$6)
        RETURNING id;
`
	var id int

	if appointmentDomain.GetUserId() == 0 {
		err := ar.databaseConection.QueryRow(ctx, query, appointmentDomain.GetClientName(), appointmentDomain.GetClientContact(), appointmentDomain.GetAppointmentDate(), appointmentDomain.GetBarberId(), appointmentDomain.GetServiceId(), nil).Scan(&id)
		if err != nil {
			logger.Error("Error create appointment_domain", err, zap.String("journey", "createAppointment"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	} else {
		err := ar.databaseConection.QueryRow(ctx, query, appointmentDomain.GetClientName(), appointmentDomain.GetClientContact(), appointmentDomain.GetAppointmentDate(), appointmentDomain.GetBarberId(), appointmentDomain.GetServiceId(), appointmentDomain.GetUserId()).Scan(&id)
		if err != nil {
			logger.Error("Error create appointment_domain", err, zap.String("journey", "createAppointment"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	appointmentDomain.SetId(id)

	logger.Info("Successful createAppointment repository", zap.String("journey", "createAppointment"),
		zap.String("journey", "createAppointment"))

	return appointmentDomain, nil
}
