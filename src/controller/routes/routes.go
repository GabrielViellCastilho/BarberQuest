package routes

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/appointment_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/availability_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/service_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/special_schedule_controller"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, uc user_controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", uc.FindUserById)
	r.GET("/getUserByEmail/:userEmail", uc.FindUserByEmail)
	r.GET("/getMyUserData", uc.FindMyUserData)
	r.GET("/getAllBarbers", uc.FindAllBarbers)
	r.GET("/getAllUsersByDateOfBirth/:dateOfBirth", uc.FindAllUsersByDateOfBirth)
	r.POST("/createUser", uc.CreateUser)
	r.POST("/createCustomerUser", uc.CreateCustomerUser)
	r.PUT("/updateMyUser", uc.UpdateUser)
	r.PUT("/updatePassword/:token", uc.UpdatePasswordUser)
	r.DELETE("/deleteUser/:userId", uc.DeleteUser)
	r.DELETE("deleteMyUser", uc.DeleteMyUser)

	r.POST("/login", uc.LoginUser)
	r.POST("/sendEmailForgotPassword/:email", uc.SendEmailResetPassword)
}

func InitServiceRoutes(r *gin.RouterGroup, sc service_controller.ServiceController) {
	r.GET("/getServiceById/:serviceId", sc.FindServiceByID)
	r.GET("/getServiceByName/:serviceName", sc.FindServiceByName)
	r.GET("/getAllService", sc.FindAllServices)
	r.GET("/getAllAvailableService", sc.FindAllAvailableServices)
	r.POST("/createService", sc.CreateService)
	r.PUT("/updateService/:serviceId", sc.UpdateService)
	r.DELETE("/deleteService/:serviceId", sc.DeleteService)
}

func InitAvailabilityRoutes(r *gin.RouterGroup, ac availability_controller.AvailabilityController) {
	r.GET("/getAllAvailability", ac.FindAllAvailabilityById)
	r.POST("/createAvailability", ac.CreateAvailability)
	r.PUT("/updateAvailability/:availabilityId", ac.UpdateAvailabilityByIdAndIdBarber)
	r.DELETE("/deleteAvailability/:availabilityId", ac.DeleteAvailabilityById)
}

func InitSpecialScheduleRoutes(r *gin.RouterGroup, sc special_schedule_controller.SpecialScheduleController) {
	r.GET("/getAllSpecialSchedule", sc.FindAllSpecialScheduleById)
	r.POST("/createSpecialSchedule", sc.CreateSpecialSchedule)
	r.DELETE("/deleteSpecialSchedule/:specialScheduleId", sc.DeleteSpecialScheduleById)
}

func InitAppointmentsRoutes(r *gin.RouterGroup, ac appointment_controller.AppointmentController) {
	r.GET("/getAppointmentsByDate/:date", ac.FindAllAppointmentsByDateAndBarberID)
	r.GET("/getAppointmentsByDateAndBarberId/:date/:id", ac.FindAllAppointmentsByDateAndBarberIDByHeader)
	r.GET("/getUserAppointments", ac.FindAllAppointmentsByUserID)
	r.GET("/getUserHistoricAppointments", ac.FindAllHistoricAppointmentsByUserID)
	r.POST("/createAppointment", ac.CreateAppointment)
	r.PUT("/updateCompletedAppointment/:appointmentId", ac.UpdateAppointmentCompletedById)
	r.GET("/checkAvailableSlots/:barberId/:serviceId/:date", ac.CheckAvailableSlots)
	r.DELETE("/deleteAppointments/:appointmentId", ac.DeleteAppointmentByIdAndUserId)
	r.DELETE("/deleteBarberAppointments/:appointmentId", ac.DeleteAppointmentByIdAndBarberId)
}
