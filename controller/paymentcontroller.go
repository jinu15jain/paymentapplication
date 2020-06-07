package controller

import (
	"github.com/gofiber/fiber"
	"tutorial/paymentapplication/paymentschema"
)

type PaymentController struct {
	PaymentService *paymentschema.PaymentService
	RefundService *paymentschema.RefundService
}

func NewPaymentController() (*PaymentController, error) {
	paymentService, err := paymentschema.NewPaymentService("Transaction")
	if err != nil {
		return nil, err
	}

	refundService, err := paymentschema.NewRefundService("Refund")
	if err != nil {
		return nil, err
	}
	return &PaymentController{
		PaymentService: paymentService,
		RefundService:  refundService}, nil
}


func (p *PaymentController) CreateTransaction(c *fiber.Ctx){
	c.Send("Create Transaction")
}

func (p *PaymentController) VerifyTransaction(c *fiber.Ctx){
	c.Send("Verify Transaction")
}

func (p *PaymentController) Refund(c *fiber.Ctx){
	c.Send("Refund Transaction")
}

func (p *PaymentController) RefundStatus(c *fiber.Ctx){
	c.Send("Refund Status")
}