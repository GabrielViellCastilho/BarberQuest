package special_schedule_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/repository/special_schedule"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/special_schedule_domain"
)

type specialScheduleDomainService struct {
	specialScheduleRepository special_schedule.SpecialScheduleRepository
}

func NewSpecialScheduleDomainService(ssr special_schedule.SpecialScheduleRepository) *specialScheduleDomainService {
	return &specialScheduleDomainService{
		specialScheduleRepository: ssr,
	}
}

type SpecialScheduleService interface {
	CreateSpecialSchedule(ctx context.Context, specialScheduleDomain special_schedule_domain.SpecialScheduleInterface) (*response.SpecialScheduleResponse, *rest_err.RestErr)
	FindAllSpecialSchedulesByID(ctx context.Context, barberID int) ([]*response.SpecialScheduleResponse, *rest_err.RestErr)
	DeleteSpecialSchedule(ctx context.Context, id int, barber_id int) *rest_err.RestErr
}
