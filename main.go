package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"tutorial/paymentapplication/controller"
)

func main(){
	 app := fiber.New()
	 controller, err := controller.NewPaymentController()
	 if(err != nil){
	 	panic(err)
	 }
	 fmt.Println("Starting Payment Application")
	 app.Get("/health-check", func(c *fiber.Ctx) {
		c.Send("healthy")
	 })
	 app.Post("/payment/create-transaction", controller.CreateTransaction)
	 app.Post("/payment/verify-transaction", controller.VerifyTransaction)
	 app.Post("/payment/refund-transaction", controller.Refund)
	 app.Post("/payment/refund-status", controller.RefundStatus)

	app.Listen(10000)
}
