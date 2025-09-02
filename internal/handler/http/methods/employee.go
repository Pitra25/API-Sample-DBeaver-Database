package http_m

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository/models"
	service "Educational-API-DBeaver-Sample-Database/internal/servise"
	messages "Educational-API-DBeaver-Sample-Database/pkg/messages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:ignore
type Employee struct {
	s *service.Service
}

func NewEmployeeHandler(service *service.Service) *Employee {
	return &Employee{s: service}
}

// GET_Employees retrieves all employees
// @Summary Get all employees
// @Description Returns a list of all employees in the system
// @Tags employees
// @Produce json
// @Success 200 {array} models.Employee "List of employees"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /employees [get]
func (h *Employee) GET_Employees(c *gin.Context) {
	result, err := h.s.Employee.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_EmployeeById retrieves employee by ID
// @Summary Get employee by ID
// @Description Returns employee by specified identifier
// @Tags employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} models.Employee "Found employee"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Employee not found"
// @Router /employees/{id} [get]
func (h *Employee) GET_EmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Employee.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_Employee creates a new employee
// @Summary Create new employee
// @Description Creates a new employee with provided data
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body models.EmployeeInput true "Employee data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /employees [post]
func (h *Employee) POST_Employee(c *gin.Context) {
	var employee models.EmployeeInput
	if err := c.BindJSON(&employee); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Employee.Create(&employee)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_Employee updates an existing employee
// @Summary Update employee
// @Description Updates employee by ID with provided data
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body models.EmployeeInput true "Updated employee data"
// @Success 200 {object} Employee "Updated employee"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Employee not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /employees/{id} [put]
func (h *Employee) PUT_Employee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	var employee models.EmployeeInput
	if err := c.BindJSON(&employee); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	if err := h.s.Employee.Put(&employee, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_EmployeeById deletes employee by ID
// @Summary Delete employee
// @Description Deletes employee by specified identifier
// @Tags employees
// @Produce json
// @Param id path string true "Employee ID"
// @Success 204 "Employee successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Employee not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /employees/{id} [delete]
func (h *Employee) DEL_EmployeeById(c *gin.Context) {
	employee_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	err = h.s.Employee.Delete(employee_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
