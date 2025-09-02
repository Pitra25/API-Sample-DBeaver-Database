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
type Customer struct {
	s *service.Service
}

func NewCustomerHandler(service *service.Service) *Customer {
	return &Customer{s: service}
}

// GET_Customers retrieves all customers
// @Summary Get all customers
// @Description Returns a list of all customers in the system
// @Tags customers
// @Produce json
// @Success 200 {array} models.Customer
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /customers [get]
func (h *Customer) GET_Customers(c *gin.Context) {
	result, err := h.s.Customer.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "customers not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_CustomerById retrieves customer by ID
// @Summary Get customer by ID
// @Description Returns customer by specified identifier
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer "Found customer"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Customer not found"
// @Router /customers/{id} [get]
func (h *Customer) GET_CustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	result, err := h.s.Customer.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	} else if result == nil {
		messages.New(c, http.StatusNotFound, "customer not found", messages.Error)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_Customer creates a new customer
// @Summary Create new customer
// @Description Creates a new customer with provided data
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.CustomerInput true "Customer data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /customers [post]
func (h *Customer) POST_Customer(c *gin.Context) {
	var customer models.CustomerInput
	if err := c.BindJSON(&customer); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	id, err := h.s.Customer.Create(&customer)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_Customer updates an existing customer
// @Summary Update customer
// @Description Updates customer by ID with provided data
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param customer body models.CustomerInput true "Updated customer data"
// @Success 200 {object} messages.Message "Updated customer"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Customer not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /customers/{id} [put]
func (h *Customer) PUT_Customer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	var customer models.CustomerInput
	if err := c.BindJSON(&customer); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	if err := h.s.Customer.Put(&customer, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_CustomerById deletes customer by ID
// @Summary Delete customer
// @Description Deletes customer by specified identifier
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 204 "Customer successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Customer not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /customers/{id} [delete]
func (h *Customer) DEL_CustomerById(c *gin.Context) {
	customer_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	err = h.s.Customer.Delete(customer_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Error)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
