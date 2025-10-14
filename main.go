package main

import (
	"fmt"
	"log"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/database/postgre"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/appointment_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/availability_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/routes"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/service_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/special_schedule_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/user_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/appointment"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/baber_service"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/barber_availability"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/special_schedule"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/user"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/appointment_service"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/availability_service"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/barber_service_service"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/special_schedule_service"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	conn, err := postgre.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database connection: %v", err))
	}

	ur := user.NewUserRepository(conn)
	service := user_service.NewUserDomainService(ur)
	controller := user_controller.NewUserController(service)

	routes.InitUserRoutes(&router.RouterGroup, controller)

	ur1 := baber_service.NewServiceRepository(conn)
	service1 := barber_service_service.NewServiceDomainService(ur1)
	controller1 := service_controller.NewServiceController(service1)

	routes.InitServiceRoutes(&router.RouterGroup, controller1)

	ur2 := barber_availability.NewAvailabilityRepository(conn)
	service2 := availability_service.NewAvailabilityDomainService(ur2)
	controller2 := availability_controller.NewAvailabilityController(service2)

	routes.InitAvailabilityRoutes(&router.RouterGroup, controller2)

	ur3 := special_schedule.NewSpecialScheduleRepository(conn)
	service3 := special_schedule_service.NewSpecialScheduleDomainService(ur3)
	controller3 := special_schedule_controller.NewSpecialScheduleController(service3)

	routes.InitSpecialScheduleRoutes(&router.RouterGroup, controller3)

	ur4 := appointment.NewAppointmentRepository(conn)
	service4 := appointment_service.NewAppointmentDomainService(ur4)
	controller4 := appointment_controller.NewAppointmentController(service4, service)

	routes.InitAppointmentsRoutes(&router.RouterGroup, controller4)
	
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}

}
