package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/krapie/showbox/server/handlers"
)

func Endpoints(ctx context.Context) (*gin.Engine, error) {
	router := gin.Default()

	handler, err := handlers.New(ctx)
	if err != nil {
		return nil, err
	}

	v1 := router.Group("/api/v1")
	{
		v1.POST("/shows", handler.RegisterShow)
		v1.GET("/shows/:id", handler.GetShow)

		v1.GET("/tickets", handler.GetTickets)
		v1.POST("/tickets/:id/reserve", handler.ReserveTicket)
		v1.POST("/tickets/:id/buy", handler.BuyTicket)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	return router, nil
}
