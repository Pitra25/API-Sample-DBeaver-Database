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
type Invoice struct {
	s *service.Service
}

func NewInvoiceHandler(service *service.Service) *Invoice {
	return &Invoice{s: service}
}

// GET_Invoices retrieves all invoices
// @Summary Get all invoices
// @Description Returns a list of all invoices in the system
// @Tags invoices
// @Produce json
// @Success 200 {array} models.Invoice "List of invoices"
// @Failure 500 {object} messages.Message "Internal server error"
// @Router /invoices [get]
func (h *Invoice) GET_Invoices(c *gin.Context) {
	result, err := h.s.Invoice.Get()
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET_InvoiceById retrieves invoice by ID
// @Summary Get invoice by ID
// @Description Returns invoice by specified identifier
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} models.Invoice "Found invoice"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Invoice not found"
// @Router /invoices/{id} [get]
func (h *Invoice) GET_InvoiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	result, err := h.s.Invoice.GetById(id)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST_Invoice creates a new invoice
// @Summary Create new invoice
// @Description Creates a new invoice with provided data
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body models.InvoiceInput true "Invoice data"
// @Success 201 {object} messages.Message "id"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 500 {object} messages.Message "Creation error"
// @Router /invoices [post]
func (h *Invoice) POST_Invoice(c *gin.Context) {
	var invoice models.InvoiceInput
	if err := c.BindJSON(&invoice); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	id, err := h.s.Invoice.Create(&invoice)
	if err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// PUT_Invoice updates an existing invoice
// @Summary Update invoice
// @Description Updates invoice by ID with provided data
// @Tags invoices
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param invoice body models.InvoiceInput true "Updated invoice data"
// @Success 200 {object} Invoice "Updated invoice"
// @Failure 400 {object} messages.Message "Invalid data"
// @Failure 404 {object} messages.Message "Invoice not found"
// @Failure 500 {object} messages.Message "Update error"
// @Router /invoices/{id} [put]
func (h *Invoice) PUT_Invoice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	var invoice models.InvoiceInput
	if err := c.BindJSON(&invoice); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	if err := h.s.Invoice.Put(&invoice, id); err != nil {
		messages.New(c, http.StatusInternalServerError, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}

// DEL_InvoiceById deletes invoice by ID
// @Summary Delete invoice
// @Description Deletes invoice by specified identifier
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 204 "Invoice successfully deleted"
// @Failure 400 {object} messages.Message "Invalid ID"
// @Failure 404 {object} messages.Message "Invoice not found"
// @Failure 500 {object} messages.Message "Deletion error"
// @Router /invoices/{id} [delete]
func (h *Invoice) DEL_InvoiceById(c *gin.Context) {
	invoice_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	err = h.s.Invoice.Delete(invoice_id)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error(), messages.Fatal)
		return
	}

	c.JSON(http.StatusOK, messages.StatusResponse{Status: "ok"})
}
