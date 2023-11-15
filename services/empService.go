package services

import (
	"virtual_machine/models"
	"virtual_machine/repositories"

	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
	employeeRepository *repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepository *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (es EmployeeService) GetListEmployee(ctx *gin.Context) ([]*models.Employee, *models.ResponseError) {
	return es.employeeRepository.GetListEmployee(ctx)
}

func (cs EmployeeService) GetEmployee(ctx *gin.Context, id int64) (*models.Employee, *models.ResponseError) {
	return cs.employeeRepository.GetEmployee(ctx, id)
}
