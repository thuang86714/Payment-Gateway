package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/service"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

type Controller struct{}

// NewUserController creates a new UserController
func NewController() *Controller {
	return &Controller{}
}

// CreateTransaction
func (ctr *Controller) CreateTransaction(c *gin.Context) {
	//receive request body from the merchant
	var curInvoice models.Invoice
	if err := c.ShouldBindJSON(&curInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create invoiceID with CreateInvocieID
	invoiceID := service.CreateInvoiceID(curInvoice)

	//create transactionWithPSP with NewTransactionWithPSP
	transaction := service.NewTransactionWithPSP(curInvoice, invoiceID)

	//send transactionWithPSP to Bank to process the transaction and receive response
	postResponse, err := service.PostTransactionToBank(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error posting transaction: %v", err)})
		return
	}

	//store PostResponse from Bank in DB
	err = service.CreateTransactionToDB(&postResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postResponse)
}

// GetTransaction handles the HTTP request to retrieve a transaction by invoice ID
func (ctr *Controller) GetTransaction(c *gin.Context) {
    // Retrieve the invoiceID query parameter from the request URL
	invoiceID := c.Query("invoiceID")
	if invoiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoiceID query parameter is required"})
		return
	}

	// Call the service to get the transaction
	getResponse, err := service.GetTransactionByInvoiceID(invoiceID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, getResponse)
}

// UpdateTransaction handles the PATCH request to update an existing transaction record in the database
func (ctr *Controller) UpdateTransaction(c *gin.Context) {
    var updatedResponse models.PostResponse
	if err := c.ShouldBindJSON(&updatedResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    err := service.UpdateTransactionInDB(&updatedResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}
