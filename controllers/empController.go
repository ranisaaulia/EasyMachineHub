package controllers

import (
	"log"
	"net/http"
	"strconv"
	"virtual_machine/services"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService *services.EmployeeService
}

// declare constructor
func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

// method
func (employeeController EmployeeController) GetListEmployee(ctx *gin.Context) {
	response, responseErr := employeeController.employeeService.GetListEmployee(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (employeeController EmployeeController) GetEmployee(ctx *gin.Context) {

	employeeId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := employeeController.employeeService.GetEmployee(ctx, int64(employeeId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
