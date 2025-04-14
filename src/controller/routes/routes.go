package routes

import (
	"fmt"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/appointment_controller"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/availability_controller"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/service_controller"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/special_schedule_controller"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/user_controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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

func getTimestampedFile(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return filePath // Retorna o caminho original se houver erro
	}
	timestamp := fileInfo.ModTime().Unix() // Obtém a data da última modificação
	return fmt.Sprintf("/static/%s?v=%d", filepath.Base(filePath), timestamp)
}

func InitTemplates(r *gin.Engine) {
	// Define o diretório de templates (HTML)
	r.LoadHTMLGlob("src/view/templates/*.html")

	// Define o diretório de arquivos estáticos (CSS, JS, favicon, etc.)
	r.Static("/static", "src/view/templates/static")

	// Rota para a página inicial
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	// Rota para a página inicial
	r.GET("/appointment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "appointment.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	// Rota para a página inicial
	r.GET("/myAppointments", func(c *gin.Context) {
		c.HTML(http.StatusOK, "myAppointments.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/myAppointmentsBarber", func(c *gin.Context) {
		c.HTML(http.StatusOK, "myAppointmentsBarber.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/myHistoric", func(c *gin.Context) {
		c.HTML(http.StatusOK, "historicAppointments.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/myProfile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "myProfile.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/inPersonAppointments", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inPersonAppointment.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/barberAvailability", func(c *gin.Context) {
		c.HTML(http.StatusOK, "barberAvailability.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/barberSpecialSchedule", func(c *gin.Context) {
		c.HTML(http.StatusOK, "specialSchedule.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/appointmentAllBarbers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "appointmentAllBarbers.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	r.GET("/birthdaysOfTheDay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "birthdaysOfTheDay.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})

	// Rota para a de atualizar senha
	r.GET("/updatePassword/:token", func(c *gin.Context) {
		c.HTML(http.StatusOK, "update_password.html", gin.H{
			"CSS": getTimestampedFile("src/view/templates/static/styles.css"),
			"JS":  getTimestampedFile("src/view/templates/static/index.js"),
		})
	})
}
