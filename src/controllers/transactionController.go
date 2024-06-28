package controllers

import (
	"server/src/configs"
	"server/src/helpers"
	"server/src/middlewares"
	"server/src/models"
	"server/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/veritrans/go-midtrans"
)

func GetAllTransaction(c *fiber.Ctx) error {
	res := models.GetAllTransaction()
	count := helpers.CountData("Transactions")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully retrieved all transactions",
		"data":    res,
		"count":   count,
	})
}

func GetUserTransaction(c *fiber.Ctx) error {
	claims := middlewares.GetUserClaims(c)
	id := claims["ID"].(float64)
	status := c.Query("status")
	res := models.GetTransactionUser(uint(id), status)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "OK",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})
}

func CreateTransaction(c *fiber.Ctx) error {
	claims := middlewares.GetUserClaims(c)
	id := claims["ID"].(float64)

	var newTransaction models.Transaction
	if err := c.BodyParser(&newTransaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	tx := configs.DB.Begin()
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to begin transaction",
		})
	}

	var totalAmount float64
	var items []midtrans.ItemDetail
	var totalQuantity uint64
	var productId []uint64
	var newStock []uint64
	productQuantities := make(map[uint64]uint64)
	for _, detail := range newTransaction.Details {
		productDetail := models.GetDetailProduct(int(detail.ProductID))
		if productDetail == nil {
			tx.Rollback()
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}

		productQuantities[detail.ProductID] += uint64(detail.ProductQuantity)
		for productID, totalQuantity := range productQuantities {
			productDetail := models.GetDetailProduct(int(productID))
			if productDetail == nil {
				tx.Rollback()
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "Product not found",
				})
			}

			newStockValue := uint64(productDetail.Stock) - totalQuantity
			newStock = append(newStock, uint64(newStockValue))
			productId = append(productId, productID)
		}

		newStockValue := productDetail.Stock - uint(detail.ProductQuantity)
		newStock = append(newStock, uint64(newStockValue))
		productId = append(productId, uint64(productDetail.ID))

		productPrice := productDetail.Price
		detailAmount := productPrice * float64(detail.ProductQuantity)
		totalAmount += detailAmount
		totalQuantity += detail.ProductQuantity

		detail.TransactionID = uint64(newTransaction.ID)
		detail.Product = *productDetail

		items = append(items, midtrans.ItemDetail{
			Price: int64(productPrice),
			Qty:   int32(detail.ProductQuantity),
			Name:  productDetail.Name,
		})
	}
	transactionNumber := helpers.GenerateTransactionNumber()
	createTransaction := models.Transaction{

		UserID:            uint64(id),
		TransactionNumber: transactionNumber,
		Quantity:          totalQuantity,
		TotalAmount:       totalAmount,
		ShippingAddress:   newTransaction.ShippingAddress,
		Status:            "waiting payment",
		PaymentMethod:     newTransaction.PaymentMethod,
		Details:           newTransaction.Details,
	}

	if err := tx.Create(&createTransaction).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create transaction",
		})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to commit transaction",
		})
	}

	foundUser := models.GetDetailUser(id)
	snapGateway := midtrans.SnapGateway{
		Client: services.MidtransClient,
	}

	var totalItemPrice int64
	for _, item := range items {
		totalItemPrice += item.Price * int64(item.Qty)
	}

	req := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionNumber,
			GrossAmt: totalItemPrice,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: foundUser.Name,
			Email: foundUser.Email,
		},
		Items: &items,
	}

	snapResp, err := snapGateway.GetToken(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create transaction with Midtrans",
			"error":   err.Error(),
		})
	}
	res, _ := models.GetActiveCartDetail(uint(id))
	var cartId []uint
	for _, detail := range res {
		cartId = append(cartId, detail.ID)
	}

	for i, id := range productId {
		if err := configs.DB.Model(&models.Product{}).Where("id = ?", id).Update("stock", newStock[i]).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update stock",
			})
		}
	}
	configs.DB.Where("id IN ?", cartId).Delete(&models.CartDetail{})
	cart, _ := models.GetActiveCartByUserId(uint(id))
	if cart != nil {
		var total_amount float64
		if len(cart.CartDetail) >= 1 {
			for _, item := range cart.CartDetail {
				if item.IsChecked {
					total_amount += item.TotalPrice
				} else {
					total_amount = float64(0)
				}
			}
		} else {
			if err := models.DeleteCart(cart.ID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Successfully created transaction",
		"data":         createTransaction,
		"token":        snapResp.Token,
		"redirect_url": snapResp.RedirectURL,
	})
}

func PaymentCallback(c *fiber.Ctx) error {
	var callbackData struct {
		VaNumbers []struct {
			VaNumber string `json:"va_number"`
			Bank     string `json:"bank"`
		} `json:"va_numbers"`
		TransactionTime   string `json:"transaction_time"`
		TransactionStatus string `json:"transaction_status"`
		TransactionID     string `json:"transaction_id"`
		StatusMessage     string `json:"status_message"`
		StatusCode        string `json:"status_code"`
		SignatureKey      string `json:"signature_key"`
		SettlementTime    string `json:"settlement_time"`
		PaymentType       string `json:"payment_type"`
		PaymentAmounts    []struct {
		} `json:"payment_amounts"`
		OrderID     string `json:"order_id"`
		MerchantID  string `json:"merchant_id"`
		GrossAmount string `json:"gross_amount"`
		FraudStatus string `json:"fraud_status"`
		ExpiryTime  string `json:"expiry_time"`
		Currency    string `json:"currency"`
	}
	if err := c.BodyParser(&callbackData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	tx := configs.DB.Begin()
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to begin transaction",
		})
	}

	var paymentMethod string
	if len(callbackData.VaNumbers) > 0 {
		paymentMethod = callbackData.PaymentType + "-" + callbackData.VaNumbers[0].Bank
	} else {
		paymentMethod = callbackData.PaymentType
	}
	updateFieldsPending := map[string]interface{}{
		"status":         "waiting payment",
		"payment_method": paymentMethod,
	}
	if callbackData.TransactionStatus == "pending" {
		if err := tx.Model(&models.Transaction{}).Where("transaction_number = ?", callbackData.OrderID).Updates(&updateFieldsPending).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update transaction status",
			})
		}
	}

	updateFields := map[string]interface{}{
		"status":         "packed",
		"payment_method": paymentMethod,
	}
	if callbackData.TransactionStatus == "settlement" {
		if err := tx.Model(&models.Transaction{}).Where("transaction_number = ?", callbackData.OrderID).Updates(updateFields).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update transaction status",
			})
		}

	}
	updateFieldsError := map[string]interface{}{
		"status":         "canceled",
		"payment_method": paymentMethod,
	}
	if callbackData.TransactionStatus == "expire" {
		if err := tx.Model(&models.Transaction{}).Where("transaction_number = ?", callbackData.OrderID).Updates(updateFieldsError).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update transaction status",
			})
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to commit transaction",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Transaction status updated to packed",
	})
}
