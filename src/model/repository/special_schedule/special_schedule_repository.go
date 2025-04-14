package special_schedule

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/special_schedule_domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type specialScheduleRepository struct {
	databaseConection *pgxpool.Pool
}

func NewSpecialScheduleRepository(databaseConection *pgxpool.Pool) *specialScheduleRepository {
	return &specialScheduleRepository{databaseConection}
}

type SpecialScheduleRepository interface {
	CreateSpecialSchedule(ctx context.Context, specialScheduleDomain special_schedule_domain.SpecialScheduleInterface) (special_schedule_domain.SpecialScheduleInterface, *rest_err.RestErr)
	FindAllSpecialScheduleByID(ctx context.Context, barberID int) ([]special_schedule_domain.SpecialScheduleInterface, *rest_err.RestErr)
	DeleteSpecialScheduleById(ctx context.Context, id int, barber_id int) *rest_err.RestErr
}
