package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/service"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

type Controller struct{
	svc *service.Service
}

// NewUserController creates a new UserController
func NewController(svc *service.Service) *Controller {
	return &Controller{svc: svc}
}

// ProcessTransaction will take transactionWithPSP from gateway, process the transaction and respond postResponse
func (ctr *Controller) ProcessTransaction(c *gin.Context) {
	var curTransactionWithPSP models.TransactionWithPSP
	if err := c.ShouldBindJSON(&curTransactionWithPSP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Check if invoice ID already exists
	if ctr.svc.DoesInvoiceExists(curTransactionWithPSP.InvoiceID) {
		initialResponse, err := ctr.svc.CreateResponse(curTransactionWithPSP, "failed: invoiceID already exists")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
			c.Abort()
			return
		}

		c.JSON(http.StatusConflict, initialResponse)
		c.Abort()
		return
	}

	// Store the invoice ID
	if err := ctr.svc.StoreInvoiceID(curTransactionWithPSP.InvoiceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store invoice ID"})
		return
	}

	// Create and send initial "processing" response
	response, err := ctr.svc.CreateResponse(curTransactionWithPSP, "done")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response)
}