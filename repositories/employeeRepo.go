package repositories

import (
	"database/sql"
	"net/http"
	"virtual_machine/models"
	dbcontext "virtual_machine/repositories/dbContext"

	"github.com/gin-gonic/gin"
)

type EmployeeRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEmployeeRepository(dbHandler *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		dbHandler: dbHandler,
	}
}

func (ep EmployeeRepository) GetListEmployee(ctx *gin.Context) ([]*models.Employee, *models.ResponseError) {

	store := dbcontext.New(ep.dbHandler)
	employees, err := store.ListEmployees(ctx)

	listEmployees := make([]*models.Employee, 0)

	for _, v := range employees {
		employee := &models.Employee{
			EmployeeID: v.EmployeeID,
			UserID:     v.UserID,
			Name:       v.Name,
			Position:   v.Position,
		}
		listEmployees = append(listEmployees, employee)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEmployees, nil
}

func (ep EmployeeRepository) GetEmployee(ctx *gin.Context, id int64) (*models.Employee, *models.ResponseError) {

	store := dbcontext.New(ep.dbHandler)
	employee, err := store.GetEmployees(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &employee, nil
}
