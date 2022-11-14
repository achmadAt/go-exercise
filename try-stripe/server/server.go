package server

import (
	"context"
	"net/http"
	"os"
	"try-stripe/model"
	"try-stripe/service"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/charge"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Server(ctx context.Context, e *echo.Echo, service service.Service) {
	root := e.Group("strype")
	root.POST("", func(c echo.Context) error {
		var payment model.Models
		payment.ID = primitive.NewObjectID()
		if err := c.Bind(&payment); err != nil {
			return err
		}
		apiKey := os.Getenv("STRIPE_KEY")
		stripe.Key = apiKey
		_, err := charge.New(&stripe.ChargeParams{
			Amount:       stripe.Int64(payment.Amount),
			Currency:     stripe.String(string(stripe.CurrencyUSD)),
			Source:       &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_visa")},
			ReceiptEmail: stripe.String(payment.ReceiptEmail),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &echo.Map{"data": err})
		}
		data, err := service.AddPayment(ctx, &payment)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &echo.Map{"data": data})
	})
}
