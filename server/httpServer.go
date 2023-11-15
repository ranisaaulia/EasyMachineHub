package server

import (
	"database/sql"
	"log"
	"virtual_machine/controllers"
	"virtual_machine/repositories"
	"virtual_machine/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	employeeController *controllers.EmployeeController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	employeeRepository := repositories.NewEmployeeRepository(dbHandler)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	router := gin.Default()

	router.GET("/employee", employeeController.GetListEmployee)
	router.GET("/employee/:id", employeeController.GetEmployee)

	return HttpServer{
		config:             config,
		router:             router,
		employeeController: employeeController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
